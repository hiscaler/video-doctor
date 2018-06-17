package main

import (
	"fmt"
	"vd"
	"os"
	"path/filepath"
)

func main() {
	workDirectory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("Work Directory : " + workDirectory)

	identity := "test"
	file := filepath.Join(workDirectory, "videos", "test", "demo.mp4")
	v := new(vd.Video)
	v.Init(identity).SetFile(file)
	fmt.Println(v.TempFile)
}
