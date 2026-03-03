package wraperr

import (
	"fmt"
	"os"
)

func OpenFile2(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		// Since loop breaks here when file is not exist, I just created fake file object for demonstration.
		f = os.NewFile(0, filename)
	}

	if errRf := ReadFile(f); errRf != nil {
		return fmt.Errorf("failed to open file: %w", errRf)
	}

	return nil
}

func ReadFile(f *os.File) error {
	_, err := os.ReadFile(f.Name())
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	return nil
}
