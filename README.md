# goexport

[Documentation online](http://godoc.org/github.com/fern4lvarez/goexport)

**goexport** reads a given file path and exports all key-value bash formatlines to the env

## Install (with GOPATH set on your machine)
----------

* Step 1: Get the `goexport` package

```
go get github.com/fern4lvarez/goexport
```

* Step 2 (Optional): Run tests

```
$ go test -v -cover ./...
```

##Usage
----------
```
package main

import (
  "fmt"
  "os"
  "path/filepath"

  export "github.com/fern4lvarez/goexport"
)

func main() {
  exampleFile := filepath.Join(os.Getenv("PWD"), "example", "example.sh")
  export.Do(exampleFile)

  fmt.Println("K", os.Getenv("K"))
  fmt.Println("FOO", os.Getenv("FOO"))
}

  // Output:
  // K V
  // FOO BAR
}
```

##License
----------
goexport is MIT licensed.