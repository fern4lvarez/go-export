# go-export [![GoDoc](https://godoc.org/github.com/fern4lvarez/go-export?status.png)](http://godoc.org/github.com/fern4lvarez/go-export)

**go-export** reads a given file path and exports all key-value bash format lines to the env

## Install

* Step 1: Get the `go-export` package

```
go get github.com/fern4lvarez/go-export
```

* Step 2 (Optional): Run tests

```
$ go test -v -cover ./...
```

##Usage

```go
package main

import (
  "fmt"
  "os"
  "path/filepath"

  "github.com/fern4lvarez/go-export"
)

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
}


func main() {
	ExampleDo()
  	// Output:
  	// K V
  	// FOO BAR
  	// user fa@csdada*%^&*%&
  	// pwd vshugvshcfvhscf
}
```

## License

go-export is MIT licensed.
