package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Adi",
			request: "Adi",
		},
		{
			name:    "AdiSiswanto",
			request: "Adi Siswanto",
		},
		{
			name:    "Sizuwano",
			request: "Sizuwano",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			HelloWorld(benchmark.request)
		})
	}
}
func BenchmarkSub(b *testing.B) {
	b.Run("Adi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adi")
		}
	})

	b.Run("Bilek", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bilek")
		}
	})
}
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adi")
	}
}
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Adi",
			request:  "Adi",
			expected: "Hello Adi",
		},
		{
			name:     "Hayasaka",
			request:  "Hayasaka",
			expected: "Hello Hayasaka",
		},
		{
			name:     "Roxanne",
			request:  "Roxanne",
			expected: "Hello Roxanne",
		},
		{
			name:     "Bimlek",
			request:  "Bimlek",
			expected: "Hello Bimlek",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
func TestSubTest(t *testing.T) {
	t.Run("Adi", func(t *testing.T) {
		result := HelloWorld("Adi")
		assert.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	})

	t.Run("Ujang", func(t *testing.T) {
		result := HelloWorld("Ujang")
		assert.Equal(t, "Hello Ujang", result, "Result must be 'Hello Ujang'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Testing")

	m.Run()

	fmt.Println("After Unit Testing")
}
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on windows")
	}

	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	fmt.Println("TestSkip done")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	fmt.Println("TestHelloWorld with Require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Adi")
	assert.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	fmt.Println("TestHelloWorld with Assert done")
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Adi")
	if result != "Hello Adi" {
		t.Fatal("result isn't Hello Adi")
	}

	fmt.Println("TestHelloWorld Done")
}
