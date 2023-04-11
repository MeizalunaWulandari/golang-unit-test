package helper

import (
	"testing"
	"fmt"
)

func TestHelloWorldLuna(t *testing.T){
	result := HelloWorld("Luna")

	if result != "Hello Luna" {
		// Error
		// t.Fail()
		t.Error("Result must be 'Hello Luna' ")
	}

	fmt.Println("TestHelloWorldLuna Done")
}

func TestHelloWorldAndini(t *testing.T){
	result := HelloWorld("Andini")

	if result != "Hello Andini" {
		// Error
		// t.FailNow()
		t.Fatal("Result must be 'Hello Andini' ")


	}
	fmt.Println("TestHelloWorldAndini Done") 
}