package goexport

import (
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

	if v1 := os.Getenv("K"); v1 != "V" {
		t.Errorf(msgFail, "Do", "V", v1)
	}

	if v1 := os.Getenv("FOO"); v1 != "BAR" {
		t.Errorf(msgFail, "Do", "BAR", v1)
	}

}

func TestDoError(t *testing.T) {
	if err := Do("donotexist"); err == nil {
		t.Errorf(msgFail, "Do", "file does not exist", nil)
	}
}
