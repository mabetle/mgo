package mcore

import "strings"

// GetUnixPath change windows path to unix like format
func GetUnixPath(location string) string {
	return String(location).ReplaceAll(`\`, "/").String()
}

// GetLocation
func GetLocation(dir string, file string) string {
	if String(dir).IsEndWith("/") {
		return dir + file
	}
	return dir + "/" + file
}

// GetFileName returns path filename.
// path seperate with "/"
func GetFileName(path string) string {
	if !String(path).IsContains("/") {
		return path
	}
	return String(path).SepEnd("/").String()
}

// GetFileExt returns file extendf
func GetFileExt(path string) string {
	if !String(path).IsContains(".") {
		return ""
	}
	return String(path).SepEnd(".").String()
}

// GetFilePath returns file path
func GetFilePath(file string) string {
	file = GetUnixPath(file)
	ps := strings.Split(file, "/")
	return strings.Join(ps[:len(ps)-1], "/")
}
