package log_view

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type File struct {
	Name    string
	Path    string
	Size    int64
	ModTime time.Time
	Body    []byte
}

// Чтение файла
func (f *File) Read() {
	file, err := os.Open(f.Path + f.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, 32*1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("read %d bytes: %v", n, err)
		}
		if n > 0 {
			f.Body = buf[:n]
		}
	}
}

type LogFiles struct {
	Name_category string
	Files         []*File
}

func (lf *LogFiles) Init() (bool, error) {
	files, err := ioutil.ReadDir("/var/log/")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	lf.Name_category = "standart"
	for _, f := range files {
		if !f.IsDir() {
			file := File{Name: f.Name(), Path: "/var/log/", Size: f.Size(), ModTime: f.ModTime()}
			go file.Read()
			lf.Files = append(lf.Files, &file)
		}
	}
	return true, nil
}
