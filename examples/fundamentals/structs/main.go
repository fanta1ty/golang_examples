package main

import "fmt"

type TestStruct struct {
	Sample    string `json:"sample"`
	SampleInt int    `json:"sample_int"`
}

//func (t TestStruct) ChangeSample(s string) {
//	t.Sample = s
//	fmt.Println("Change sample: ", t)
//}

func (t *TestStruct) ChangeSample(s string) {
	t.Sample = s
	fmt.Println("Change sample: ", t)
}

func main() {
	testVar := TestStruct{"aa", 1}
	fmt.Println("testVar: ", testVar)
	testVar.ChangeSample("sample")
	fmt.Println("testVar: ", testVar)

	testVar2 := TestStruct{Sample: "aa", SampleInt: 1}
	fmt.Println("testVar2: ", testVar2)

}
