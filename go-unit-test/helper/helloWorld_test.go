package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {

	benchmark := []struct {
		name string
		request string
	}{
		{
			name: "helloWorld(Yazid)",
			request: "Yazid",
		},
		{
			name: "helloWorld(Baihaqi)",
			request: "Baihaqi",
		},
		{
			name: "helloWorld(Muhammad Yazid Baihaqi)",
			request: "Muhammad Yazid Baihaqi",
		},
	}

	for _, bench := range benchmark {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helloWorld(bench.request)
			}
		})
	}
}

func TestTableTest(t *testing.T) {

	datas := []struct {
		name, request, expected string
	}{
		{
			name:     "HelloWorld(Muhammad)",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name:     "HelloWorld(Yazid)",
			request:  "Yazid",
			expected: "Hello Yazid",
		},
		{
			name:     "HelloWorld(Baihaqi)",
			request:  "Baihaqi",
			expected: "Hello Baihaqi",
		},
	}

	for _, data := range datas {
		t.Run(data.name, func(t *testing.T) {
			result := helloWorld(data.request)
			require.Equal(t, data.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	fmt.Println("")

	m.Run()

	fmt.Println("")
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Muhammad", func(t *testing.T) {
		result := helloWorld("Muhammad")
		require.Equal(t, "Hello Muhammad", result, "Result must be 'Hello Muhammad'")
	})
	t.Run("Yazid", func(t *testing.T) {
		result := helloWorld("Yazid")
		require.Equal(t, "Hello Yazid", result, "Result must be 'Hello Yazid'")
	})
}

func TestSkip(t *testing.T) { //skip

	if runtime.GOOS == "linux" {
		t.Skip("Can't run on Linux")
	}

	result := helloWorld("Baihaqi")
	assert.Equal(t, "Hello Baihaqi", result, "Result must be 'Hello Baihaqi'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldAssert(t *testing.T) { //Fail()
	result := helloWorld("Baihaqi")
	assert.Equal(t, "Hello Baihaqi", result, "Result must be 'Hello Baihaqi'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) { //FailNow()
	result := helloWorld("Baihaqi")
	require.Equal(t, "Hello Baihaqi", result, "Result must be 'Hello Baihaqi'")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorld(t *testing.T) { //Fail()
	result := helloWorld("Baihaqi")

	if result != "Hello Baihaqi" {
		t.Error("Result must be 'Hello Baihaqi'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldBaihaqi(t *testing.T) { //FailNow()
	result := helloWorld("Baihaqi")

	if result != "Hello Baihaqi" {
		t.Fatal("Result must be 'Hello Baihaqi'")
	}

	fmt.Println("TestHelloWorldBaihaqi Done")
}
