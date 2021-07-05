package main

import (
  "io/ioutil"
  "log"
  "os"
  "path/filepath"
  "strings"
)

func main() {
  log.Println("开始获取并设置CDN")
	inputDir := os.Getenv("STATIC_FILE_PATH")
	if inputDir == "" {
    inputDir = "/app/html"
  }
  cdnReplaceHolder := os.Getenv("CDN_REPLACE_HOLDER")
  if cdnReplaceHolder == "" {
    cdnReplaceHolder = "https://cdn_url_place_holder/"
  }
	err := filepath.Walk(inputDir,
		func(fPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !!info.IsDir() {
				return nil
			}
			ext := filepath.Ext(fPath)
			isVendor := false
			if ext == ".js" {
				_, filename := filepath.Split(fPath)
				beginStr := filename[0:7]
				if beginStr == "vendor." {
					isVendor = true
				}
			}
			if !isVendor && (ext == ".html" || ext == ".js" || ext == ".css") {
        log.Println(fPath)
        read, err := ioutil.ReadFile(fPath)
        if err != nil {
          panic(err)
        }
        new := os.Getenv("CDN_URL")
        if new == "" {
          new = "/"
        }
        newContents := strings.Replace(string(read), cdnReplaceHolder, new, -1)
        err = ioutil.WriteFile(fPath, []byte(newContents), 0)
        if err != nil {
          panic(err)
        }
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	} else {
    log.Println("CDN设置完毕！")
  }
}
