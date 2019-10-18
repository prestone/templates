package templates

import (
	"fmt"
	"strconv"
)

type dict struct {
	key   []byte
	value []byte
}

func (a *Engine) Dict(key string, v interface{}) {
	a.db.dict = append(a.db.dict, createdict(key, v))
	a.render()
}

func createdict(k string, v interface{}) *dict {
	var vv []byte
	switch v.(type) {
	case []byte:
		vv = v.([]byte)
	case string:
		vv = []byte(v.(string))
	case int:
		vv = []byte(strconv.Itoa(v.(int)))
	default:
		vv = []byte(fmt.Sprint(v))
	}

	return &dict{[]byte(k), vv}
}
