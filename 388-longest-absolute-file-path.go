package main

import (
	"fmt"
	"strings"
)

type file struct {
	deepth int
	name string
}

func lengthLongestPath(path string) int {
	if path == "" {
		return 0
	}
	var lastLongestPath string
	fs := getFilesOrFolder(path)
	pathes := map[int]string{}
	for i:=0; i<len(fs); i++ {
		pathes[fs[i].deepth] = fs[i].name
		if i<len(fs)-1 {
			if fs[i].deepth >= fs[i+1].deepth {
				if isAFile(fs[i].name) {
					p := getFilePath(fs[i].deepth, pathes)
					if len(p) > len(lastLongestPath) {
						lastLongestPath = p
					}
				}
			}
		} else {
			if isAFile(fs[i].name) {
				p := getFilePath(fs[i].deepth, pathes)
				if len(p) > len(lastLongestPath) {
					lastLongestPath = p
				}
			}
		}
	}
	fmt.Println(lastLongestPath)
	return len(lastLongestPath)
}

func getFilePath(deepth int, pathes map[int]string) string {
	var path string
	for i:=0; i<=deepth; i++ {
		p, ok := pathes[i]
		if !ok {
			fmt.Println(i)
			fmt.Println(pathes[i])
			fmt.Println("something is wrong !!!")
		}
		if path == "" {
			path = p
		} else {
			path = path + "/" + p
		}
	}

	return path
}

func isAFile(name string) bool {
	si := strings.Index(name, ".")
	return si != -1 && si != len(name)-1
}

func getFilesOrFolder(path string) []file {
	var fs []file
	ss := strings.Split(path, "\n")
	for _, s := range ss {
		var tc = 0
		for i:=0; i<=len(s)-1; i++ {
			if string(s[i]) != "\t" {
				fs = append(fs, file{tc, s[i:len(s)]})
				break
			}
			tc += 1
		}
	}
	fmt.Println(fs)
	return fs
}

func main() {
    fss := []string{
        // "",
        // "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
        // "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
        //"dir\n    file.txt",
        // "a\n\tb1\n\t\tf1.txt\n\taaaaa\n\t\tf2.txt",
        // "a\n\tb.txt\na2\n\tb2.txt",
    }
    for _, fs := range fss {
        fmt.Printf("--------------\nfor \n%v\nit's %v\n--------------\n", fs, lengthLongestPath(fs))
    }
}