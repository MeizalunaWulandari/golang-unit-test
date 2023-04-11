package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"runtime"
)

func BenchmarkSub(b *testing.B){
	b.Run("Luna", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			HelloWorld("Luna")
		}
	})
	b.Run("Andini", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			HelloWorld("Andini")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Luna")
	}  
}

func BenchmarkHelloWorldAndini(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Andini")
	}  
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name : "Luna",
			request : "Luna",
			expected : "Hello Luna",
		},
		{
			name : "Andini",
			request : "Andini",
			expected : "Hello Andini",
		},
		{
			name : "Rizka",
			request : "Rizka",
			expected : "Hello Rizka",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T){
	t.Run("Luna", func (t *testing.T){
		result := HelloWorld("Luna")
		require.Equal(t, "Hello Luna", result, "Result must be `Hello Luna`")
	})
	t.Run("Andini", func (t *testing.T){
		result := HelloWorld("Andini")
		require.Equal(t, "Hello Andini", result, "Result must be `Hello Andini`")
	})
}

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