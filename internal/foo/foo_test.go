package foo

import "testing"

func TestFoo(t *testing.T){
	expected := "Foo"
	actual := Foo()
	if expected != actual{ //bu kisim assertion kismidir
		t.Errorf("expected %s do not match actual %s",expected,actual)
	}
}