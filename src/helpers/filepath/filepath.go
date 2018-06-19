// 助手类
package filepath

import (
	"os"
	"path/filepath"
	"os/exec"
	"strings"
	"io"
)

// Check file is exist
func IsExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

// Copy file
func CopyFile(dstName, srcName string) (Written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
