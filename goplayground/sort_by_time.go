package main

import (
	"fmt"
	"io/ioutil"
	//"os"
	//"syscall"
	"sort"
	"time"
)

type fileInfo struct {
	name    string
	modtime time.Time
}

type fileSlice []fileInfo

func (f fileSlice) Len() int {
	return len(f)
}

func (f fileSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f fileSlice) Less(i, j int) bool {
	return f[i].modtime.Before(f[j].modtime)
}

func (f fileSlice) PrintInfo() {
	for _, v := range f {
		fmt.Printf("%s ", v.name)
	}
	fmt.Println()
}

func getAllfileName(dirname string) fileSlice {
	var ret fileSlice
	files, _ := ioutil.ReadDir(dirname)
	for _, f := range files {
		var tmp fileInfo
		tmp.name = f.Name()
		tmp.modtime = f.ModTime()
		ret = append(ret, tmp)
	}
	return ret
}

func main() {
	ret := getAllfileName(".")
	ret.PrintInfo()
	sort.Sort(ret)
	ret.PrintInfo()
}
