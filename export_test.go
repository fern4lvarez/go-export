package export

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var (
	msgFail  = "%v method fails. Expects %v, returns %v"
	filePath = filepath.Join(os.Getenv("PWD"), "example", "example.sh")
)

func TestDo(t *testing.T) {
	if err := Do(filePath); err != nil {
		t.Errorf(msgFail, "Do", nil, err)
	}

	if v := os.Getenv("K"); v != "V" {
		t.Errorf(msgFail, "Do", "V", v)
	}

	if v := os.Getenv("FOO"); v != "BAR" {
		t.Errorf(msgFail, "Do", "BAR", v)
	}

	if v := os.Getenv("user"); v != "fa@csdada*%^&*%&" {
		t.Errorf(msgFail, "Do", "fa@csdada*%^&*%&", v)
	}

	if v := os.Getenv("pwd"); v != "vshugvshcfvhscf" {
		t.Errorf(msgFail, "Do", "vshugvshcfvhscf", v)
	}

}

func TestDoError(t *testing.T) {
	if err := Do("donotexist"); err == nil {
		t.Errorf(msgFail, "Do", "file does not exist", nil)
	}
}

func ExampleDo() {
	exampleFile := filepath.Join(os.Getenv("PWD"), "example", "example.sh")
	err := Do(exampleFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("K", os.Getenv("K"))
	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Println("user", os.Getenv("user"))
	fmt.Println("pwd", os.Getenv("pwd"))
	// Output:
	// K V
	// FOO BAR
	// user fa@csdada*%^&*%&
	// pwd vshugvshcfvhscf

}
