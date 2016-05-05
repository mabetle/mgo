package mio

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// open file for read.
// file implements io.Reader and io.Writer interface.
func NewFileReader(location string) (io.Reader, error) {
	file, err := os.OpenFile(location, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}

// file may not exist.
// if file not exist, create one.
func NewFileWriter(location string) (io.Writer, error) {
	fs, err := os.Create(location)
	if err == os.ErrNotExist {
		fs.Write([]byte{})
		log.Printf("File:%s not exist, create one.")
	} else if err != nil {
		return nil, err
	}
	return bufio.NewWriter(fs), nil
}

// OpenFile
func OpenFile(location string) (*os.File, error) {
	return os.Open(location)
}

// ReadBinneryFile
func ReadBinneryFile(location string) ([]byte, error) {
	return ioutil.ReadFile(location)
}

// put string to file
// if file not exist, create one.
func WriteBinneryFile(location string, content []byte) (int, error) {
	fs, e := os.Create(location) // Create can overide exsts file.
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	log.Printf("Write File To: %s \n", location)
	return fs.Write(content)
}
