package streamzip

import (
	"archive/zip"
	"io"
)

/**
 * compress
 *
 * @date 18/5/30
 */

// ZipFileInfo the file's info which will be used for zipping
type ZipFileInfo struct {
	Name       string
	ReadStream io.Reader
}

// Compress the files
func Compress(files []ZipFileInfo, writeStream io.Writer) error {
	// wrapper the writeStream
	zipWriter := zip.NewWriter(writeStream)
	defer zipWriter.Close()

	// set the buf size to up the speed for samll files
	buf := make([]byte, 1000000)

	for _, fileInfo := range files {
		// create the writeStream for each file
		writer, err := zipWriter.Create(fileInfo.Name)
		if err != nil {
			return err
		}

		// zip stream
		_, err = io.CopyBuffer(writer, fileInfo.ReadStream, buf)
		if err != nil {
			return err
		}
	}
	return nil
}
