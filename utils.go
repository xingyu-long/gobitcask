package gobitcask

import "os"

func createFolder(dirPath string) error {
	// create the folder
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return nil
}

func deleteFolder(dirPath string) error {
	if _, err := os.Stat(dirPath); err == nil {
		err = os.RemoveAll(dirPath)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
