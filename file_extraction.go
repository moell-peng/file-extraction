package main

import (
	"flag"
	"fmt"
	"time"
	"path/filepath"
	"os"
	"github.com/moell-peng/file-extraction/config"
	"strings"
	"io"
	"errors"
)

var dateFormat = "2006-01-02 15:04:05"

func main() {

	startTime, endTime, confPath, flagErr := flagParse()
	if flagErr != nil {
		fmt.Println(flagErr)
		return
	}

	if err := config.Load(confPath); err != nil {
		fmt.Println("Failed to load confuration", err)
		return
	}

	conf := config.Get()

	if (conf.SaveDir == "" || conf.Dir == "" || conf.SaveDir == conf.Dir) {
		fmt.Println("Please confure the correct operating directory and save directory")
		return
	}

	filepath.Walk(conf.Dir,
		func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return err
			}

			path = strings.Replace(path, "\\", "/", -1)

			if (inExclude(path, conf.ExcludeList)) {
				if (info.IsDir()) {
					return filepath.SkipDir
				} else {
					return nil
				}
			}

			if (info.ModTime().Unix() > startTime.Unix() && info.ModTime().Unix() < endTime.Unix()) {
				if (!info.IsDir()) {
					des := strings.Replace(path, conf.Dir, conf.SaveDir, -1)
					copyFile(path, des)
				}
			}
			return nil
		})

}

func flagParse() (start time.Time, end time.Time, configPath string, err error) {
	now := time.Now()
	h, _ := time.ParseDuration("-1h")

	startTime := flag.String("start_time", now.Add(h).Format(dateFormat), "Last modification time start time, default is one hour ago")
	endTime := flag.String("end_time", now.Format(dateFormat), "Last modified end time, the default is the current time")
	confPath := flag.String("conf_path", "config/config.yaml", "Configuration file path")

	flag.Parse()

	parseStartTime, _:= time.ParseInLocation(dateFormat, *startTime, time.Local)
	parseEndTime, _:= time.ParseInLocation(dateFormat, *endTime, time.Local)

	if (parseStartTime.Unix() < 0 || parseEndTime.Unix() < 0) {
		return parseStartTime, parseEndTime, *confPath, errors.New("Please enter the correct time format such as 2018-01-01 01:01:01")
	}

	return parseStartTime, parseEndTime, *confPath, nil
}

func copyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	os.MkdirAll(filepath.Dir(des), os.ModePerm)

	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

func inExclude(path string, excludeList []string) (exists bool) {
	for _, exclude := range excludeList {
		if (path == exclude) {
			return true
		}
	}
	return false
}

