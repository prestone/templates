package templates

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type component struct {
	path string
	key  []byte
	body []byte
	f    func() []byte
}

func (a *component) reset() (err error) {
	a.body, err = ioutil.ReadFile(a.path)
	if err != nil {
		fmt.Println(a.path, err)
	}
	return
}

func (a *Engine) Component(key string, path string) {
	if !exists(path) {
		fmt.Println("Component not found", path)
		return
	}
	i := &component{
		key:  []byte(key),
		path: path,
	}
	i.reset()
	a.db.components = append(a.db.components, i)
	a.render()

	go notify(path, func() {
		i.reset()
		a.render()
	})
}

func (a *Engine) ComponentTime(key string) {
	i := &component{
		key: []byte(key),
		f:   func() []byte { return []byte(fmt.Sprint(time.Now().Unix())) },
	}
	a.db.components = append(a.db.components, i)
	a.render()
}

func (a *Engine) ComponentRand(key string) {
	i := &component{
		key: []byte(key),
		f:   func() []byte { return []byte(fmt.Sprint(rand.Uint64())) },
	}
	a.db.components = append(a.db.components, i)
	a.render()
}
