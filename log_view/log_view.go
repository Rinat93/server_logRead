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

// LogFiles Структура с списком лог файлов
type LogFiles struct {
	NameCategory string
	Files        []*File
}

// Init охраняем список логов
func (lf *LogFiles) Init() (bool, error) {
	// go lf.EventRegister()
	return lf.Add("/var/log/")
}

// func (lf *LogFiles) EventRegister() {
// 	for {

// 	}
// }

// Add Добавить путь логов
func (lf *LogFiles) Add(path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	lf.NameCategory = "standart"
	for _, f := range files {
		if !f.IsDir() {
			file := File{Name: f.Name(), Path: "/var/log/", Size: f.Size(), ModTime: f.ModTime()}
			go file.Read()
			lf.Files = append(lf.Files, &file)
		}
	}
	return true, nil
}
