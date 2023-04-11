package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"runtime"
)

func TestMain(m *testing.M){
	// Before
	fmt.Println("Before Unit Test")

	// Running Unit Test 
	m.Run()

	// After
	fmt.Println("After Unit Test")
}

// Skip Test
func TestSkip(t *testing.T){
	if runtime.GOOS == "linux"{
		t.Skip("Cannot run on linux")
	}
	result := HelloWorld("Luna")
	require.Equal(t, "Hello Luna", result, "Result must be 'Hello Luna'")
}

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