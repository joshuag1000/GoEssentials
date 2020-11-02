package goessentials

import (
	"io"
	"os"
)

//
// File related Subroutines and Functions.
//

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//
// String Subroutines and Functions.
//

// this can be used to reverse a string
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
