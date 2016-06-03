package mcore

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// PrepareFile prepare file
func PrepareFile(location string) error {
	location = ProcessDir(location)
	if IsFileExist(location) {
		return nil
	}
	_, err := WriteFile(location, "")
	return err
}

// NewFileWriter file may not exist.
// if file not exist, create one.
func NewFileWriter(location string) (io.Writer, error) {
	location = ProcessDir(location)
	fs, err := os.Create(location)
	if err == os.ErrNotExist {
		fs.Write([]byte{})
		logger.Debugf("File:%s not exist, create one.")
	} else if err != nil {
		return nil, err
	}
	return bufio.NewWriter(fs), nil
}

// GetFileModifyTime returns file modified time
func GetFileModifyTime(file string) (int64, error) {
	file = ProcessDir(file)
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// GetFileSize return files bytes
func GetFileSize(file string) (int64, error) {
	file = ProcessDir(file)
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// RemoveFile delete file
func RemoveFile(file string) error {
	file = ProcessDir(file)
	return os.Remove(file)
}

// RenameFile rename file name
func RenameFile(file string, to string) error {
	file = ProcessDir(file)
	to = ProcessDir(to)
	return os.Rename(file, to)
}

// MoveFile move file
// if the newPath not exist, create one
// make sure move can success.
func MoveFile(oldPath, newPath string) error {
	//TODO not finished yet.
	println("MoveFile not implement")
	return nil
}

// WriteFileLines write file lines
func WriteFileLines(file string, lines []string) (int, error) {
	content := StringArray(lines).MergeWithOsNewLine()
	return WriteFile(file, content)
}

// WriteFileBytes write file bytes
func WriteFileBytes(file string, v []byte) error {
	r := bytes.NewBuffer(v)
	return WriteFileReader(file, r)
}

// WriteFileReader write file header
func WriteFileReader(file string, r io.Reader) error {
	file = ProcessDir(file)
	if err := PrepareFileDir(file); err != nil {
		return err
	}
	w, errW := NewFileWriter(file)
	if errW != nil {
		return errW
	}
	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	return nil
}

// PrepareFileDir prepare file directory
func PrepareFileDir(file string) error {
	file = ProcessDir(file)
	dir := GetFileDir(file)
	// exist, skip
	if IsFileExist(dir) {
		return nil
	}
	if err := MakeDir(dir); err != nil {
		logger.Debugf("Make directory error: Directory: %s Error: %v", dir, err)
		return err
	}
	return nil
}

// ProcessDir process directory
func ProcessDir(dir string) string {
	// process dir sperator
	if IsWindows() {
		dir = strings.Replace(dir, "/", "\\", -1)
	} else {
		dir = strings.Replace(dir, "\\", "/", -1)
	}
	// process not expand vars
	dir = ExpandPath(dir)
	return dir
}

// WriteFile put string to file
func WriteFile(file string, content string) (n int, err error) {
	// after write, logger result
	defer func() {
		logger.Debugf("Write to file: %s, bytes: %d, ErrMsg: %v \n", file, n, err)
	}()
	file = ProcessDir(file)
	//prepare dir
	if err := PrepareFileDir(file); err != nil {
		return 0, err
	}
	fs, e := os.Create(file)
	if e != nil {
		logger.Debugf("Create file error: File: %s, Error: %v", file, e)
		return 0, e
	}
	defer fs.Close()
	n, err = fs.WriteString(content)
	return
}

// WriteFileGBK write file
func WriteFileGBK(file, text string) (int, error) {
	return WriteFile(file, EncodeGBK(text))
}

// AppendFile append file
func AppendFile(file string, appendContent string) (int, error) {
	return AppendFileByLocation(file, appendContent, false)
}

// AppendFileBefore append file from beginning
func AppendFileBefore(file string, appendContent string) (int, error) {
	return AppendFileByLocation(file, appendContent, true)
}

// AppendFileByLocation append file by location
func AppendFileByLocation(file string, appendContent string, before bool) (int, error) {
	file = ProcessDir(file)
	if IsFileExist(file) {
		content, err := ReadFileAll(file)
		if err != nil {
			return 0, err
		}
		if before {
			content = appendContent + content
		} else {
			content = content + appendContent
		}
		return WriteFile(file, content)
	}
	//file not exist
	return WriteFile(file, appendContent)
}

// ReadFileAll same to GetContent
func ReadFileAll(file string) (string, error) {
	return GetFileContent(file)
}

// GetFileContent returns string from text file
func GetFileContent(file string) (string, error) {
	file = ProcessDir(file)
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

// IsFile returns false when it's a directory or does not exist.
func IsFile(file string) bool {
	file = ProcessDir(file)
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsFileExist returns whether a file or directory exists.
func IsFileExist(path string) bool {
	path = ProcessDir(path)
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// TouchFile touch file
func TouchFile(file string) error {
	if IsFileExist(file) {
		// TODO update modify time
		println("not implement TouchFile when exists.")
		return nil
	}
	//not exist, create blank file
	_, err := WriteFile(file, "")
	return err
}

// GetRealPath returns file real path
func GetRealPath(file string) string {
	f, err := os.Stat(file)
	if err != nil {
		return file
	}
	return f.Name()
}

// GetFileDir return file directory
func GetFileDir(location string) string {
	vs := strings.Split(location, "/")
	sb := NewStringBuffer()
	sb.Append("/")
	for index, v := range vs {
		if v == "" {
			continue
		}
		if index < len(vs)-1 {
			sb.Append(v, "/")
		}
	}
	return sb.String()
}

// MakeDir make dir
func MakeDir(dir string) error {
	dir = ProcessDir(dir)
	if IsFileExist(dir) {
		return fmt.Errorf("%s is exist", dir)
	}
	if err := os.MkdirAll(dir, 0777); err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("No enough permission to create dir")
		}
		return err
	}
	return nil
}

// GetPathFiles returns all file in dir
func GetDirSubFiles(dir string) (fs []string) {
	dir = ProcessDir(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range files {
		if f.IsDir() {
			// skip dir
			continue
		}
		fs = append(fs, dir+"/"+f.Name())
	}
	return
}

// GetSubFiles returns dir sub files
// exts: format .xx format.
// r: recursive
func GetSubFiles(
	dir string,
	r bool,
	exts string,
	skipDirs string,
	skipFiles string,
) []string {
	dir = ProcessDir(dir)
	var result []string
	err := GetSubFilesImpl(&result, dir, r, exts, skipDirs, skipFiles)
	if err != nil {
		logger.Warn(err)
	}
	return result
}

// GetSubFilesImp returns sub files
func GetSubFilesImpl(
	result *[]string,
	dir string,
	r bool,
	exts string,
	skipDirs string,
	skipFiles string,
) error {
	dir = String(dir).TrimEnd("/").String()
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return err
	}

	extsArray := strings.Split(exts, ",")

	for _, file := range files {
		// skip start with . dir or file
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		// full match
		if skipDirs != "" && String(file.Name()).IsContainsInSepStringIgnoreCase(skipDirs, ",") {
			continue
		}
		if file.IsDir() && !file.Mode().IsRegular() {
			if r {
				GetSubFilesImpl(result, dir+"/"+file.Name(), r, exts, skipDirs, skipFiles)
			}
		} else {
			name := dir + "/" + file.Name()
			if skipFiles != "" && String(file.Name()).IsContainsInSepStringIgnoreCase(skipFiles, ",") {
				continue
			}
			// has exts
			if len(extsArray) > 0 {
				for _, ext := range extsArray {
					if String(name).IsEndWith(ext) {
						*result = append(*result, name)
					}
				}
			} else {
				// no exts
				*result = append(*result, name)
			}
		}
	}
	return nil
}

// WriteFileWithError get content from somewhere may be occur error
// when error is not nil, skip write
func WriteFileWithError(location, content string, err error) error {
	if err != nil {
		return err
	}
	_, err = WriteFile(location, content)
	return err
}

// DirSubs include dir and files
func DirSubs(dir string) ([]string, error) {
	dir = ProcessDir(dir)
	f, err := os.Open(dir)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	return f.Readdirnames(-1)
}

// DirSubFiles returns dir sub files.
func DirSubFiles(dir string) ([]string, error) {
	dir = ProcessDir(dir)
	f, err := os.Open(dir)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	infos, err := f.Readdir(-1)
	if err != nil {
		return []string{}, err
	}
	var names []string
	for _, fi := range infos {
		if !fi.IsDir() {
			names = append(names, fi.Name())
		}
	}
	return names, nil
}
