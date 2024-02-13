package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T){
	t.Run("wahyu", func(t *testing.T) {
		result := HelloWorld("wahyu")
		require.Equal(t, "hello wahyu", result, "output harus hello wahyu")
	})
	t.Run("aprilliandhika", func(t *testing.T) {
		result := HelloWorld("aprilliandhika")
		require.Equal(t, "hello aprilliandhika", result, "output harus hello aprilliandhika")
	})

}

// func TestMain1(t *testing.T){
// 	fmt.Println("sebelum testing")
// 	t.Run()
// 	fmt.Println("setelah testing")
// }

func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin"{
		t.Skip("tidak bisa di run di Mac Os")
	}
	result := HelloWorld("wahyu")
	require.Equal(t, "hello wahyu", result, "output harus hello wahyu")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("wahyu")
	require.Equal(t, "hello wahyu", result, "output harus hello wahyu")
	fmt.Println("test is done")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("wahyu")
	assert.Equal(t, "hello wahyu", result, "output harus hello wahyu")
	fmt.Println("test is done")
}

func TestHelloWorld(t *testing.T){
	result := HelloWorld("wahyu")

	if result != "hello wahyu"{
		t.Fail()
	}
	fmt.Println("done")
}
