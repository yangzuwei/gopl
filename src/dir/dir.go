// dir project dir.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

func main() {
	f := filepath.Dir("/f/golang/src/degree")
	fmt.Println(f)
	f1, _ := os.Stat("..")
	fmt.Println(reflect.TypeOf(f1))
	var fileNames []string
	var dirNames = []os.FileInfo{f1}
	for tmp := range dirNames {
		if err != nil {
			break
		}
		if tmp.IsDir() {
			inside, _ := ioutil.ReadDir(tmp.Name())
			for i, _ := range inside {
				if i.IsDir() {
					st.Push(i)
				} else {
					fileNames = append(fileNames, i.Name())
				}
			}
		} else {
			fileNames = append(fileNames, tmp.Name())
		}

		fmt.Println(tmp)
	}

}
