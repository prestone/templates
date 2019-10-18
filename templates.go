package templates

import (
	"bytes"
	"fmt"
	"sync"
)

func New() (a *Engine) {
	a = new(Engine)
	a.db.templates = make(map[string]*template)
	return
}

type Engine struct {
	db struct {
		templates  map[string]*template
		components []*component
		dict       []*dict
		sync.Mutex
		minify bool
	}
}

func (a *Engine) render() {
	a.db.Lock()
	defer a.db.Unlock()

	//components
	for _, template := range a.db.templates {
		template.reset()
		for _, r := range a.db.components {
			switch r.f != nil {
			case true:
				template.body = bytes.ReplaceAll(template.body, r.key, r.f())
			default:
				template.body = bytes.ReplaceAll(template.body, r.key, r.body)
			}
		}
	}
	//dict
	for _, template := range a.db.templates {
		for _, x := range a.db.dict {
			template.body = bytes.ReplaceAll(template.body, x.key, x.value)
		}
	}

	//optimize
	for _, r := range a.db.templates {
		if r.optimize != nil {
			r.body = r.optimize(r.body)
		}
	}
}

func (a *Engine) Follow(path string) {
	go notify(path, func() {
		a.render()
	})
}

type body []byte

func (a *body) Bytes() []byte {
	return []byte(*a)
}
func (a *body) String() string {
	return string(*a)
}
func (a *body) Replace(k, v string) *body {
	if a == nil {
		return a
	}
	(*a) = bytes.ReplaceAll((*a), []byte(k), []byte(v))
	return a
}
func (a *body) ReplaceBytes(k string, v []byte) *body {
	if a == nil {
		return a
	}
	(*a) = bytes.ReplaceAll((*a), []byte(k), v)
	return a
}

func (a *Engine) Get(id string) *body {
	a.db.Lock()
	defer a.db.Unlock()
	if a.db.templates[id] != nil {
		c := new(body)
		(*c) = a.db.templates[id].body
		return c
	}
	fmt.Println("ID not found", id)
	return &body{}
}

func (a *Engine) GetBytes(id string) []byte {
	a.db.Lock()
	defer a.db.Unlock()
	if a.db.templates[id] != nil {
		return a.db.templates[id].body
	}
	fmt.Println("ID not found", id)
	return nil
}

func (a *Engine) GetString(id string) string {
	a.db.Lock()
	defer a.db.Unlock()
	if a.db.templates[id] != nil {
		return string(a.db.templates[id].body)
	}
	fmt.Println("ID not found", id)
	return ""
}
