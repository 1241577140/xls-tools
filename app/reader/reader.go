package reader

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/xuri/excelize/v2"
)

type File struct {
	FileName string
	FileDir  string
	data     *excelize.File
	fileList []string
	cache    []*excelize.File
}

func NewFile() *File {
	return &File{}
}

func (f *File) SetName(n *string) *File {
	if n != nil {
		f.FileName = *n
	}
	return f
}

func (f *File) SetDir(d *string) *File {
	if d != nil {
		f.FileDir = *d
	}
	return f
}

func (f *File) Load() error {
	f.parseDir()
	wg := sync.WaitGroup{}
	wg.Add(len(f.fileList))
	for i := 0; i < len(f.fileList); i++ {
		go func(i int) {
			file, err := excelize.OpenFile(f.fileList[i])
			if err != nil {
				log.Fatalln(err)
				return
			}
			f.cache = append(f.cache, file)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return nil
}

func (f *File) parseDir() {
	dir := make([]string, 0)
	if f.FileName != "" {
		dir = append(dir, strings.Split(f.FileName, ",")...)
	}
	if f.FileDir != "" {
		readDir, err := os.ReadDir(f.FileDir)
		if err != nil {
			return
		}
		for _, e := range readDir {
			if !e.IsDir() && strings.Contains(e.Name(), "xlsx") {
				dir = append(dir, e.Name())
			}
		}
	}
	f.fileList = dir
}

func (f *File) AllExcel() []*excelize.File {
	return f.cache
}
