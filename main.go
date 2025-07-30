package main

import (
	"fmt"
	"os"
	"time"

	"math/rand"
)

// Create the file if it does not exists, or truncate the
// existing one before write the content.
// (in place)
func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err

	}

	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync()

}


func main() {
	SaveData1("test.txt", []byte("Hello, World!"))

}
