package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	fileMask    = "test_file_"
	numberFiles = 10
	fileSize    = 25e4
	filePath    = "files"
)

func main() {
	start := time.Now()

	for i := 1; i < numberFiles+1; i++ {
		file, err := os.Create(fmt.Sprintf("%v/%v%v", filePath, fileMask, i))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		data := make([]byte, int(fileSize), int(fileSize))
		file.Write(data)
	}

	log.Println(time.Since(start))
}
