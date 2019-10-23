package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage %v <old> <new>", os.Args[0])
		return
	}

	old := os.Args[1]
	new := os.Args[2]

	if _, err := os.Stat(old); os.IsNotExist(err) {
		fmt.Printf("OLD not found")
		return
	} else if err != nil {
		fmt.Printf("error when accessing OLD: %s", err)
		return
	}

	if _, err := os.Stat(new); err != nil && !os.IsNotExist(err) {
		fmt.Printf("error when accessing NEW: %s", err)
		return
	} else if !os.IsNotExist(err) {
		if err := os.Remove(new); err != nil {
			fmt.Printf("error when deleting NEW: %s", err)
			return
		}
	}

	if err := os.Symlink(old, new); err != nil {
		fmt.Printf("error when symlinking: %s", err)
	} else {
		fmt.Printf("symlink created")
	}
}
