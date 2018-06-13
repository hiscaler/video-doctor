package vd

import (
	"path"
	"os"
	"errors"
)

type Video struct {
	originalFile  string
	filename      string
	fileExtension string
	runtimeDir    string
	resultDir     string
	tempFile      string
	information   VideoInformation
}

type VideoInformation struct {
	size     int
	width    int
	height   int
	duration int
}

func (v *Video) SetFile(file string) Video {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			v.originalFile = file
			v.filename = path.Base(v.originalFile)
			v.fileExtension = path.Ext(v.originalFile)
		}
		errors.New("File not exists.")
	}

	return *v
}

func (v *Video) getFile() string {
	if len(v.tempFile) > 0 {
		return v.tempFile
	} else {
		return v.originalFile
	}
}
