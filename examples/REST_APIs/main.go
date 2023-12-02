package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type HandleViaStruct struct {
}

type MemoryStore struct {
	items map[string]string
}

type Store interface {
	Add(shortenedURL, longURL string) error
	Remove(shortenedURL string) error
	Get(shortenedURL string) (string, error)
}

type AddPath struct {
	domain string
	store  Store
}

func (*HandleViaStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world received a request")
	defer log.Println("End hello world request")
	fmt.Fprintf(w, "Hello world via struct")
}

func (m *MemoryStore) Add(shortenedURL, longURL string) error {
	if m.items[shortenedURL] != "" {
		return fmt.Errorf("value already exists here")
	}
	m.items[shortenedURL] = longURL
	log.Println(m.items)
	return nil
}

func (m *MemoryStore) Remove(shortenedURL string) error {
	if m.items[shortenedURL] == "" {
		return fmt.Errorf("value does not exist here")
	}
	delete(m.items, shortenedURL)
	return nil
}

func (m *MemoryStore) Get(shortendURL string) (string, error) {
	longURL, ok := m.items[shortendURL]
	if !ok {
		return "", fmt.Errorf("no mapped url available here")
	}
	return longURL, nil
}

func NewMemoryStore() MemoryStore {
	return MemoryStore{
		items: make(map[string]string),
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request")
	defer log.Print("End hello world request")
	fmt.Fprintf(w, "Hello World")
}

func (a *AddPath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type addPathRequest struct {
		URL string `json:"url"`
	}

	var parsed addPathRequest
	err := json.NewDecoder(r.Body).Decode(&parsed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
	}
	h := sha1.New()
	h.Write([]byte(parsed.URL))
	sum := h.Sum(nil)
	hash := hex.EncodeToString(sum)[:10]
	err = a.store.Add(hash, parsed.URL)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
	}

	type addPathResponse struct {
		ShortenedURL string `json:"shortened_url"`
		LongURL      string `json:"long_url"`
	}

	pathResp := addPathResponse{
		ShortenedURL: fmt.Sprintf("%v/%v", a.domain, hash),
		LongURL:      parsed.URL,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pathResp)
}

type DeletePath struct {
	store Store
}

func (p *DeletePath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty hash"))
		return
	}

	err := p.store.Remove(hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}

type RedirectPath struct {
	store Store
}

func (p *RedirectPath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty hash"))
		return
	}

	longURL, err := p.store.Get(hash)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

type internalStore struct {
	Version string            `json:"version"`
	Items   map[string]string `json:"items"`
}

type FileStore struct {
	filename string
}

func NewFileStore(filename string) (FileStore, error) {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		is := internalStore{
			Version: "v1",
			Items:   make(map[string]string),
		}
		raw, err := json.Marshal(is)
		if err != nil {
			return FileStore{}, fmt.Errorf("unable to generate json representation for file")
		}

		err = os.WriteFile(filename, raw, 0644)
		if err != nil {
			return FileStore{}, fmt.Errorf("unable to persist file")
		}
	}

	return FileStore{filename: filename}, nil
}

func (f *FileStore) Add(shortenedURL, longURL string) error {
	raw, err := os.ReadFile(f.filename)
	if err != nil {
		return err
	}
	var is internalStore
	err = json.Unmarshal(raw, &is)
	if err != nil {
		return fmt.Errorf("unable to parse incoming json store data. Err: %v", err)
	}
	_, ok := is.Items[shortenedURL]
	if ok {
		return fmt.Errorf("shortened url already stored")
	}
	is.Items[shortenedURL] = longURL
	modRaw, err := json.Marshal(is)
	if err != nil {
		return fmt.Errorf("unable to convert data to json representation")
	}
	err = os.WriteFile(f.filename, modRaw, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStore) Remove(shortenedURL string) error {
	raw, err := os.ReadFile(f.filename)
	if err != nil {
		return err
	}
	var is internalStore
	err = json.Unmarshal(raw, &is)
	if err != nil {
		return fmt.Errorf("unable to parse incoming json store data. Err: %v", err)
	}
	delete(is.Items, shortenedURL)
	modRaw, err := json.Marshal(is)
	if err != nil {
		return fmt.Errorf("unable to convert data to json representation")
	}
	err = os.WriteFile(f.filename, modRaw, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStore) Get(shortenedURL string) (string, error) {
	raw, err := os.ReadFile(f.filename)

	if err != nil {
		return "", err
	}

	var is internalStore
	err = json.Unmarshal(raw, &is)
	if err != nil {
		return "", fmt.Errorf("unable to parse incoming json store data. Err: %v", err)
	}
	longURL, ok := is.Items[shortenedURL]

	if !ok {
		return "", fmt.Errorf("no url unavailable for that shortened url")
	}

	return longURL, nil
}

func main() {
	log.Print("Hello world sample started.")
	r := mux.NewRouter()
	redirectPath := "http://localhost:8080/r"
	//mem := NewMemoryStore()
	fs, err := NewFileStore("testing.json")
	if err != nil {
		panic("unable to create filestore appropriately")
	}
	r.Handle("/", &HandleViaStruct{}).Methods("GET")
	r.Handle("/add", &AddPath{domain: redirectPath, store: &fs}).Methods("POST")
	r.Handle("/r/{hash}", &DeletePath{store: &fs}).Methods("DELETE")
	r.Handle("/r/{hash}", &RedirectPath{store: &fs}).Methods("GET")
	_ = http.ListenAndServe(":8080", r)
}
