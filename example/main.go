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
