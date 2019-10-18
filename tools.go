package templates

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	ppp "path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func dirs(dir string, h func(f os.FileInfo)) {
	if dir == "" {
		dir = "."
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, f := range files {
		h(f)
	}
}

func list(dir string, ext ...string) (list []string) {
	switch len(ext) > 0 {
	case true:
		dirs(dir, func(f os.FileInfo) {
			if f.IsDir() {
				return
			}
			for _, x := range ext {
				if strings.HasSuffix(f.Name(), x) {
					list = append(list, filepath.Join(dir, f.Name()))
					return
				}
			}
		})
	default:
		dirs(dir, func(f os.FileInfo) {
			if f.IsDir() {
				return
			}
			list = append(list, filepath.Join(dir, f.Name()))
		})
	}
	return
}

func notify(filename string, onUpdate func()) (err error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("File %s updated\n", filename)
					onUpdate()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println(err)
			}
		}
	}()
	err = watcher.Add(filename)
	if err != nil {
		return
	}
	<-done
	return
}

func save(filename string, body []byte) (err error) {
	dir, _, full := path(filename)
	err = os.MkdirAll(dir, os.ModePerm)
	f, err := os.Create(full)
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	_, err = w.Write(body)
	return
}

func path(filename string) (dir, name, full string) {
	name = ppp.Base(filename)
	dir = ppp.Dir(filename)
	full = ppp.Join(dir, name)
	return
}
