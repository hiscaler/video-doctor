package main

import (
	"fmt"
	"vd"
	"os"
	"path/filepath"
	"helpers"
)

func main() {
	workDirectory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("Work Directory : " + workDirectory)

	identity := "test"
	dirs := [2]string{"runtime", "output"}
	for _, dir := range dirs {
		path := filepath.Join(workDirectory, identity, dir)
		if !helpers.IsExist(path) {
			os.Mkdir(path, os.ModePerm)
		}
	}
	file := filepath.Join(workDirectory, "videos", "test", "demo.mp4")
	v := new(vd.Video)
	v.Init(identity).SetFile(file)
	fmt.Println(v.TempFile)
}
