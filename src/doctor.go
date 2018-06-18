package main

import (
	"fmt"
	"vd"
	"os"
	"path/filepath"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}
func main() {
	workDirectory, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	identity := "test"
	file := filepath.Join(workDirectory, "videos", "test", "demo.mp4")
	v := new(vd.Video)
	v.Init(identity).SetFile(file)
	fmt.Println(v.TempFile)
}
