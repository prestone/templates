package templates

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

type compile struct {
	files    []string
	name     string
	body     []byte
	optimize func(v []byte) []byte
}

func (a *compile) render() {
	var b bytes.Buffer
	for _, x := range a.files {
		f, _ := ioutil.ReadFile(x)
		if f == nil {
			continue
		}
		b.Write(f)
	}

	a.body = b.Bytes()
	if a.optimize != nil {
		a.body = a.optimize(a.body)
	}
	save(a.name, a.body)
	fmt.Println("render", a.name)
}

func (a *Engine) Compile(to string, files ...string) {
	c := new(compile)
	c.files = make([]string, len(files))
	c.name = to
	for pos, path := range files {
		c.files[pos] = path
		go notify(path, func() {
			c.render()
			a.render()
		})
	}
	c.render()
}

func (a *Engine) CSS(to string, files ...string) {
	c := new(compile)
	c.name = to
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	c.optimize = func(v []byte) []byte {
		r, _ := m.Bytes("text/css", v)
		return r
	}
	for _, path := range files {
		if !exists(path) {
			fmt.Println("CSS not found", path)
			continue
		}
		c.files = append(c.files, path)
		go notify(path, func() {
			c.render()
			a.render()
		})
	}
	c.render()
}

func (a *Engine) CSSDir(to string, dir string) {
	a.CSS(to, list(dir, "css")...)
}
