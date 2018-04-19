// dir project dir.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

var fileNames []string

func main() {
	f := filepath.Dir("/f/golang/src/degree")
	fmt.Println(f)
	f1, _ := ioutil.ReadDir("..")
	fmt.Println(reflect.TypeOf(f1))
	storeFiles("..")
	fmt.Println(fileNames, len(fileNames))

}

func storeFiles(dir string) {
	if tmp, _ := os.Stat(dir); tmp.IsDir() {
		dirs, _ := ioutil.ReadDir(dir)
		for _, d := range dirs {
			if d.IsDir() {
				storeFiles(dir + "/" + d.Name())
			} else {
				fileNames = append(fileNames, d.Name())
			}
		}
	} else {
		fileNames = append(fileNames, dir)
	}

	return
}
