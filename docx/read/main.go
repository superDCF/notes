package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	docxFilePath := "./1857.translated.v7.docx"

	// 打开docx文件
	r, err := zip.OpenReader(docxFilePath)
	if err != nil {
		log.Fatalf("Error opening docx file: %v\n", err)
		return
	}
	defer r.Close()
	date := time.Now().Format("20060102-150405")
	err = os.MkdirAll(date, 0777)
	if err != nil {
		log.Fatalf("Error MkdirAll date %s: %v\n", date, err)
		return
	}
	// 遍历压缩包中的文件
	for _, f := range r.File {
		// 打开文件
		file, err := f.Open()
		if err != nil {
			log.Fatalf("Error opening file %s: %v\n", f.Name, err)
			return
		}
		defer file.Close()

		// 读取文件内容
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v\n", f.Name, err)
			return
		}

		// 打印文件名和内容
		fmt.Printf("File: %s\n", f.Name)
		fmt.Printf("Content:\n%s\n", content)
		name := strings.ReplaceAll(f.Name, "/", "_")
		nf, err := os.Create(fmt.Sprintf("./%s/%s", date, name))
		if err != nil {
			log.Fatalf("Error Create file %s: %v\n", f.Name, err)
			return
		}

		_, err = nf.Write(content)
		if err != nil {
			log.Fatalf("Error Write file %s: %v\n", f.Name, err)
			return
		}
		err = nf.Sync()
		if err != nil {
			log.Fatalf("Error Sync file %s: %v\n", f.Name, err)
			return
		}
	}
}
