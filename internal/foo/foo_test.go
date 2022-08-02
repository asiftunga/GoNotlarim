package foo

import (
	"fmt"
	"os"
	"testing"
)

func TestFoo(t *testing.T){
	expected := "Foo"
	actual := Foo()
	if expected != actual{ //bu kisim assertion kismidir
		t.Errorf("expected %s do not match actual %s",expected,actual)
	}
}

func TestFoo2(t *testing.T) {
   env := os.Getenv("MYENV")
   fmt.Println(env)
   //..
}

/* func TestArgs(t *testing.T) {
    arg1 := os.Args[1]
    if arg1 != "Foo" {
        t.Errorf("Expected baz do not match actual %s", arg1)
    }
} */