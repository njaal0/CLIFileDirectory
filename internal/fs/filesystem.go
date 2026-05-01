package fs

import (
	"fmt"
	"log"
	"os"
)

func ListEntries(path string) (entries []os.DirEntry, err error) {
	entries, err = os.ReadDir(path)
	if err != nil {
		log.Fatal(fmt.Println(err))
		return []os.DirEntry{}, err
	}

	return entries, nil
}

func CreateDir(path string) error {
	return os.Mkdir(path, 0755)
}

func RenameEntry(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
