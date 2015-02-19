/*
goexport reads a file and exports
all key-value bash format lines to the env
*/
package goexport

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Do reads the given file path and exports
// all key-value bash format lines to the env
//
// e.g.
// FOO=BAR
// export BAZ=BOG
func Do(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		handleLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func handleLine(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	if strings.HasPrefix(line, "#") {
		return
	}

	line = strings.TrimPrefix(line, "export ")
	line = strings.TrimSpace(line)

	validKV := regexp.MustCompile(`^[(\w|\W)]+\=[(\w|\W)]+$`)
	if validKV.MatchString(line) {
		kv := strings.SplitN(line, "=", 2)
		k, v := kv[0], strings.Trim(kv[1], `"`)
		os.Setenv(k, v)
	}
}
