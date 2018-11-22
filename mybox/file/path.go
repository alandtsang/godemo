package main

import (
	"fmt"
	"os"
)

// Check if the path exists, if it does not exist, create it.
func CheckPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	path := "/tmp/test"
	if err := CheckPath("/tmp/test"); err != nil {
		fmt.Printf("%s not ok\n", path)
	} else {
		fmt.Printf("%s ok\n", path)
	}
}
