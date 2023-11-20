package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T){
	result := HelloWorld("wahyu")

	if result != "hello wahyu"{
		t.Fail()
	}
	fmt.Println("done")
}
