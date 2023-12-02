package privateSample

import "fmt"

type unExposedStruct struct {
	Sample string
}

type ExposedStruct struct {
	ExposedSample   string
	unExposedSample string
}

func unExposedFunc() {
	fmt.Println("unExposedFunc")
}

func ExposedFunc() {
	fmt.Println("ExposedFunc")
}

func NewUnExposedStruct() unExposedStruct {
	return unExposedStruct{"Sample"}
}
