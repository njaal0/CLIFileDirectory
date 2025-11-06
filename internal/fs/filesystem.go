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
