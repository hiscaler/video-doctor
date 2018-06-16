package vd

import (
	"strings"
	"github.com/go-ozzo/ozzo-config"
	"path"
	"errors"
	"path/filepath"
	"helpers"
)

type Video struct {
	OriginalFile  string
	Filename      string
	FileExtension string
	WorkDir       string
	RuntimeDir    string
	OutputDir     string
	TempFile      string
	Information   VideoInformation
	ProcessConfig Config
}

type VideoInformation struct {
	size     int
	width    int
	height   int
	duration int
}

type Config struct {
	workflow                  string // 处理流程设置
	audioFile                 string
	concatFiles               string
	cutSecondsPer             int
	removeHeaderFooterSeconds []int
}

func (v *Video) Init(identity string) *Video {
	v.WorkDir = helpers.GetCurrentDirectory()
	identity = strings.Trim(identity, "")
	c := config.New()
	configFile, _ := filepath.Abs("src/configs/" + identity + ".json")
	if helpers.IsExist(configFile) {
		c.Load(configFile)
		cfg := new(Config)
		cfg.workflow = c.GetString("workflow", "")
		cfg.audioFile = c.GetString("audioFile")
		cfg.concatFiles = c.GetString("concatFiles", "header#footer")
		cfg.cutSecondsPer = c.GetInt("cutSecondsPer", 90)
		cfg.removeHeaderFooterSeconds = []int{c.GetInt("removeHeaderFooterSeconds.header"), c.GetInt("removeHeaderFooterSeconds.footer")}
		v.ProcessConfig = *cfg

		v.RuntimeDir = filepath.Join(v.WorkDir, "runtime")
		v.OutputDir = filepath.Join(v.WorkDir, "output")
	} else {
		errors.New(configFile + " is not exists.")
	}

	return v

}

func (v *Video) SetFile(file string) *Video {
	if helpers.IsExist(file) {
		v.OriginalFile = file
		v.TempFile = file
		v.Filename = path.Base(v.OriginalFile)
		v.FileExtension = path.Ext(v.Filename)
	} else {
		errors.New(file + "File not exists.")
	}

	return v
}

func (v *Video) getFile() string {
	if len(v.TempFile) > 0 {
		return v.TempFile
	} else {
		return v.OriginalFile
	}
}
