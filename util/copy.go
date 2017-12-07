package util

import (
	"io/ioutil"
)

// CopyFile will create a copy of the specified file in the specified
// location. The destination file should include the filename.
func CopyFile(sourceFile, destFile string) error {
	from, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destFile, from, 0666)
	return err
}
