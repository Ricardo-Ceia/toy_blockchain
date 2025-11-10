package files_oper

import (
	"os"
)

func ReadFromFile(path string) []byte {
	fileData, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return fileData
}
