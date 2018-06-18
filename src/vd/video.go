package vd

import (
	"strings"
	"github.com/go-ozzo/ozzo-config"
	"path"
	"errors"
	"path/filepath"
	"helpers"
	"os"
	"os/exec"
	"fmt"
	"io/ioutil"
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
	workflow  string // 处理流程设置
	watermark []string
	voice struct {
		file      string
		startTime int
		duration  int
	}
	concatFiles               string
	cutSecondsPer             int
	removeHeaderFooterSeconds map[string]int
}

func (v *Video) Init(identity string) *Video {
	v.WorkDir = filepath.Join(helpers.GetCurrentDirectory(), "videos", identity)
	identity = strings.Trim(identity, "")
	c := config.New()
	configFile, _ := filepath.Abs("src/configs/" + identity + ".json")
	if helpers.IsExist(configFile) {
		c.Load(configFile)
		cfg := new(Config)
		cfg.workflow = c.GetString("workflow", "")
		cfg.watermark = []string{c.GetString("watermark")}
		voice := struct {
			file      string
			startTime int
			duration  int
		}{}
		voice.file = c.GetString("voice.file")
		voice.startTime = c.GetInt("voice.startTime")
		voice.duration = c.GetInt("voice.duration")
		cfg.voice = voice
		cfg.concatFiles = c.GetString("concatFiles", "header#footer")
		cfg.cutSecondsPer = c.GetInt("cutSecondsPer", 90)
		m := make(map[string]int)
		m["header"] = c.GetInt("removeHeaderFooterSeconds.header", 0)
		m["footer"] = c.GetInt("removeHeaderFooterSeconds.footer", 0)
		cfg.removeHeaderFooterSeconds = m
		v.ProcessConfig = *cfg

		v.RuntimeDir = filepath.Join(v.WorkDir, "runtime")
		v.OutputDir = filepath.Join(v.WorkDir, "output")

		// 创建 runtime 和 output 目录
		dirs := [2]string{"runtime", "output"}
		for _, dir := range dirs {
			path := filepath.Join(v.WorkDir, dir)
			println(path)
			if !helpers.IsExist(path) {
				os.Mkdir(path, os.ModePerm)
			}
		}
	} else {
		errors.New(configFile + " is not exists.")
	}

	return v
}

// 生成临时文件名
func (v *Video) tmpFile() string {
	return filepath.Join(v.RuntimeDir, helpers.GenerateUniqueId()+v.FileExtension)
}

func (v *Video) SetFile(file string) *Video {
	if helpers.IsExist(file) {
		v.OriginalFile = file
		v.Filename = path.Base(v.OriginalFile)
		v.FileExtension = path.Ext(v.Filename)
		dstFile := v.tmpFile()
		if _, err := helpers.CopyFile(dstFile, v.OriginalFile); err == nil {
			v.TempFile = dstFile
		} else {
			errors.New("Copy original file to runtime directory error." + err.Error())
		}
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

func (v *Video) Mute() *Video {
	file := v.TempFile
	println(file)

	return v
}

// 执行 ffmpeg 命令处理音视频文件
func (v *Video) ffmpegCommand(cmd string) {
	exec.Command(cmd)
}

// 遮盖水印
func (v *Video) RemoveWatermark() *Video {
	config := v.ProcessConfig.watermark
	tempFile := v.tmpFile()
	cmd := fmt.Sprintf("ffmpeg -i %s -vf delogo=%s %s", v.TempFile, config, tempFile)
	v.ffmpegCommand(cmd)

	return v
}

func (v *Video) Clean() {
	files, _ := ioutil.ReadDir(v.RuntimeDir)
	for _, file := range files {
		if file.IsDir() {
			os.RemoveAll(file.Name())
		} else {
			os.Remove(file.Name())
		}
	}
}

func (v *Video) Done() {
	v.Clean()
}
