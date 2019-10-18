package templates

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

type template struct {
	id       string
	path     string
	body     []byte
	dict     []*dict
	optimize func([]byte) []byte
}

func (a *template) Dict(k string, v interface{}) {
	a.dict = append(a.dict, createdict(k, v))
}

func (a *template) reset() (err error) {
	a.body, err = ioutil.ReadFile(a.path)
	if err != nil {
		fmt.Println(a.path, err)
	}
	for _, x := range a.dict {
		a.body = bytes.ReplaceAll(a.body, x.key, x.value)
	}
	return
}

func (a *Engine) Template(id string, path string) (x *template) {
	if !exists(path) {
		fmt.Println("Not exist", path)
		return
	}
	i := &template{
		id:   id,
		path: path,
	}
	i.reset()
	a.db.Lock()
	a.db.templates[id] = i
	a.db.Unlock()
	go notify(path, func() {
		a.render()
	})
	a.render()
	return i
}

//has optimisation but! very carefully, Vue is not working after optimisation;)
func (a *Engine) HTML(id string, path string) {
	if !exists(path) {
		fmt.Println("Not exist", path)
		return
	}

	i := &template{
		id:   id,
		path: path,
	}
	m := minify.New()
	m.AddFunc("text/html", css.Minify)
	i.optimize = func(v []byte) []byte {
		r, _ := m.Bytes("text/html", v)
		return r
	}
	i.reset()
	a.db.Lock()
	a.db.templates[id] = i
	a.db.Unlock()
	go notify(path, func() {
		a.render()
	})
	a.render()
}
