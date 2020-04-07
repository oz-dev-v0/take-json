package cmd

import (
	"os"

	"github.com/pkg/errors"
)

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func getFile() (*os.File, error) {
	if flags.filePath == "" {
		return nil, errors.New("please input a file")
	}
	if !fileExists(flags.filePath) {
		return nil, errors.New("the file provided does not exist")
	}
	file, e := os.Open(flags.filePath)
	if e != nil {
		return nil, errors.Wrapf(e,
			"unable to read the file %s", flags.filePath)
	}
	return file, nil
}

func fileExists(filePath string) bool {
	info, e := os.Stat(filePath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}
