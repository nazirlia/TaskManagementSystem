package wraperr

import (
	"fmt"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file not found %w", err)
	}
	return f, nil
}
