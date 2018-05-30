package main

import (
	"os"

	"github.com/wujohns/streamzip/streamzip"
)

/**
 * example to use
 *
 * @date 18/5/30
 */

func main() {
	fileReadStream1, _ := os.Open("example_files/go.png")
	fileReadStream2, _ := os.Open("example_files/tosaka.jpg")
	writeStream, _ := os.Create("example_files/pack.zip")

	fileReaders := make([]streamzip.ZipFileInfo, 0)
	fileReaders = append(fileReaders, streamzip.ZipFileInfo{
		Name:       "go.png",
		ReadStream: fileReadStream1,
	})
	fileReaders = append(fileReaders, streamzip.ZipFileInfo{
		Name:       "tosaka.png",
		ReadStream: fileReadStream2,
	})
	streamzip.Compress(fileReaders, writeStream)
}
