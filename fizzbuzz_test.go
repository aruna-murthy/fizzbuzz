package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/aruna-murthy/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %v should not have returned true", 1)
		t.Fail()
	}
	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %v should have returned true", 3)
		t.Fail()
	}
	if result != "fizz" {
		t.Logf("Result should have been fizz")
		t.Fail()
	}
	result, ok = fizzbuzz.Fizzbuzz(30)
	if !ok {
		t.Logf("The value %v should have returned true", 30)
		t.Fail()
	}
	result, ok = fizzbuzz.Fizzbuzz(10)
	if !ok {
		t.Logf("The value %v should have returned true", 10)
		t.Fail()
	}
}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// Output: 3
}
