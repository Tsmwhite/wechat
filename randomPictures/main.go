package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
	"wechat/config"
	"wechat/core/encrypt"
	model "wechat/models"
)

const apiUrl = "https://picsum.photos/%d/%d.jpg"
const rootPath = "./web/src/assets/avatar/"
const extension = ".jpg"

var width = flag.Int("w", 200, "图片宽度")
var height = flag.Int("h", 200, "图片高度")
var number = flag.Int("n", 500, "图片数量")

func main() {
	flag.Parse()
	err := config.SetupServer()
	if err != nil {
		panic(err)
	}
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum", cpuNum)
	runtime.GOMAXPROCS(cpuNum)
	files := getPicture()
	insertFileData(files)
}

func getPicture() []string {
	api := fmt.Sprintf(apiUrl, *width, *height)
	fmt.Println("api", api)
	var files []string
	now := time.Now().Format("20060102150405")
	wait := sync.WaitGroup{}
	wait.Add(*number)
	for i := 0; i < *number; i++ {
		filename := fmt.Sprintf("%s%d%s", now, i, extension)
		go func() {
			defer wait.Done()
			err := downloadImg(api, filename)
			if err == nil {
				files = append(files, filename)
			} else {
				fmt.Println(err)
			}
		}()
	}
	wait.Wait()
	return files
}

func downloadImg(api, filename string) error {
	res, err := http.Get(api)
	if err != nil {
		return errors.New("Http Get error:" + err.Error())
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Io ReadAll error:" + err.Error())
	}
	file := fmt.Sprintf("%s%s", rootPath, filename)
	fmt.Println("file", file)
	err = ioutil.WriteFile(file, content, fs.ModePerm)
	if err != nil {
		return errors.New("Io WriteFile error:" + err.Error())
	}
	return nil
}

func insertFileData(files []string) {
	var fileData []model.File
	now := time.Now().Unix()
	for _, file := range files {
		fileData = append(fileData, model.File{
			Uuid:       encrypt.CreateUuid(),
			Path:       file,
			CreateTime: now,
			UpdateTime: now,
		})
	}
	model.DB().Create(&fileData)
}
