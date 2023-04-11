package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Assertion assert()
func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Luna")
	assert.Equal(t, "Hello Luna", result, "Result must be 'Hello Luna'")
	fmt.Println("TestHelloWorldAssert Done")
}

// Assertion require()
func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Luna")
	require.Equal(t, "Hello Luna", result, "Result must be 'Hello Luna'")
	fmt.Println("TestHelloWorldRequire Done")
}

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