package mcore

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

// GetFileExt returns file extend
func GetFileExt(path string) string {
	if !String(path).IsContains(".") {
		return ""
	}
	return String(path).SepEnd(".").String()
}
