//batch replace

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mabetle/mgo/mcore"
	"github.com/mabetle/mgo/mcore/mcon"
)

// args
var (
	rootDir    string
	skipDirs   string
	skipFiles  string
	searchStr  string
	replaceStr string
	exts       string
	verbose    bool
	confirm    bool
	test       bool
	recursive  bool
	help       bool
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"Usage:%s [flags] \"search content\" \"replace content\"\n\n",
		os.Args[0])

	flag.PrintDefaults()
}

func DoFlag() {
	flag.Usage = usage
	wd, _ := os.Getwd()
	flag.StringVar(&rootDir, "d", wd, "Root dir")
	flag.StringVar(&exts, "e", "", "Extends, separate by comma")
	flag.StringVar(&skipDirs, "sd", "", "Skip Dirs, separate by comma")
	flag.StringVar(&skipFiles, "sf", "", "Skip Files, separate by comma")
	flag.BoolVar(&confirm, "y", false, "Confirm do search")
	flag.BoolVar(&test, "t", false, "Test replace, no write back")
	flag.BoolVar(&recursive, "r", true, "Recursive dir")
	flag.BoolVar(&verbose, "V", true, "Print more info")
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
}

func ShowArgs() {
	fmt.Println("App Arguments")
	fmt.Println("     Root Dir: ", rootDir)
	fmt.Println(" File Extends: ", exts)
	fmt.Println("    Skip Dirs: ", skipDirs)
	fmt.Println("   Skip Files: ", skipFiles)
	fmt.Println("      Verbose: ", verbose)
	fmt.Println("         Test: ", test)
	fmt.Println("      Confirm: ", confirm)
	fmt.Println("    Recursive: ", recursive)
	fmt.Println("       Search: ", searchStr)
	fmt.Println("      Replace: ", replaceStr)
}

func ScanSearchContent() {
	searchStr = mcore.ReadNotBlankLineWithMsg("Input Search Content")
}

func ScanReplaceConent() {
	if confirm {
		return
	}
	if replaceStr == "" {
		yn := mcore.ReadBool(true, "Blank replace content, are you sure?")
		// not sure
		if !yn {
			replaceStr = mcore.ReadLineWithMsg("Input Replace Content")
			ScanReplaceConent()
		}
	}
}

func replace() {
	files := mcore.GetSubFiles(rootDir, recursive, exts, skipDirs, skipFiles)
	fmt.Printf("Found %d files.\n", len(files))
	for _, file := range files {
		fileReplace(file, searchStr, replaceStr)
	}
}

func fileReplace(file, searchStr, replaceStr string) {
	show(file, searchStr)
	text, err := mcore.ReadFileAll(file)
	// cannot read file
	if nil != err {
		fmt.Printf("Read file error,File:%s Error: %v \n", file, err)
		return
	}
	// not contains search content, skip
	if !strings.Contains(text, searchStr) {
		return
	}
	nums := strings.Count(text, searchStr)
	fmt.Printf("File: %s found %d matches.\n", file, nums)

	// confirm replace
	if !confirm && !mcore.ReadBool(true, "Confirm to replace in file") {
		return
	}
	// test not write back
	if test {
		return
	}
	//do replace
	text = strings.Replace(text, searchStr, replaceStr, -1)

	if _, err := mcore.WriteFile(file, text); err != nil {
		fmt.Printf("Write file error. File %s Error: %v \n", file, err)
	} else {
		fmt.Println("Write file :", file)
	}
}

func show(file string, search string) {
	if !verbose {
		return
	}
	fmt.Printf("Processing File:%s Search: %s\n", file, search)
	//found
	data, err := mcore.ReadFileAll(file)
	if err != nil {
		fmt.Printf("Read file error. File:%s Error:%v\n", file, err)
		return
	}
	num := strings.Count(data, search)
	// found match
	if num > 0 {
		fmt.Printf("Found %d matches in file: %s\n", num, file)
	} else {
		fmt.Printf("not found matches in file.\n")
		return
	}
	lines, _ := mcore.ReadFileLines(file)
	for lineNum, line := range lines {
		if strings.Contains(line, searchStr) {
			fmt.Printf("%d ", lineNum+1)
			rows := strings.Split(line, search)
			for i, v := range rows {
				fmt.Print(v)
				if i != len(rows)-1 {
					mcon.PrintGreen(search)
				}
			}
			fmt.Println()
		}
	}
}

func confirmArgs() {
	ShowArgs()
	// provide -y
	if confirm {
		return
	}
	doReplace := mcore.ReadBool(true, "Confirm do replace")
	if !doReplace {
		scanArgs()
	}
}

func scanArgs() {
	rootDir = mcore.ReadLineWithDefaultAndMsg(rootDir, "Root Dir")
	skipDirs = mcore.ReadLineWithDefaultAndMsg(skipDirs, "Skip Dirs")
	skipFiles = mcore.ReadLineWithDefaultAndMsg(skipFiles, "Skip Files")
	searchStr = mcore.ReadLineWithDefaultAndMsg(searchStr, "Search Content")
	replaceStr = mcore.ReadLineWithDefaultAndMsg(replaceStr, "Replace Content")
	exts = mcore.ReadLineWithDefaultAndMsg(exts, "File Extends")
	verbose = mcore.ReadBool(verbose, "Verbose")
	recursive = mcore.ReadBool(recursive, "Recursive")
	confirmArgs()
}

func main() {
	DoFlag()

	if help {
		fmt.Println("Help about command")
		usage()
		return
	}

	switch flag.NArg() {
	case 0:
		ScanSearchContent()
		ScanReplaceConent()
	case 1:
		searchStr = flag.Args()[0]
		ScanReplaceConent()
	case 2:
		searchStr = flag.Args()[0]
		replaceStr = flag.Args()[1]
	default:
	}
	confirmArgs()
	replace()
}
