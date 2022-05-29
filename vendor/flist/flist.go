package flist

import (
	"io/ioutil"
	"strings"
)

const csvRes string = ".csv"

// список .CSV фалов в директории
func GetList(directory string) map[int]string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	filesList := make(map[int]string)
	var i int = 0

	for _, file := range files {
		if strings.Contains(file.Name(), csvRes) {
			filesList[i] = file.Name()
			i++
		}
	}

	return filesList
}
