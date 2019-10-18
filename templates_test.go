package templates

//testing

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()

	fmt.Println(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	r := New()
	r.Template("index", "index.html")
	r.Template("empty", "empty.html")
	r.Component("@@menu", "menu.html")
	r.Dict("@@content", "paste content")
	r.Dict("_xxxx", "'so nice'")

	fmt.Println(r.GetString("index"))
}

func TestHTML(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	fmt.Println(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	r := New()
	r.HTML("index", "index.html")
	r.HTML("empty", "empty.html")
	r.Component("@@menu", "menu.html")
	r.Dict("@@content", "paste content")
	r.Dict("_xxxx", "'so nice'")

	fmt.Println(r.GetString("index"))
}

func TestCompile(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	fmt.Println(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	r := New()
	r.CSS("compiled.css", "1.css", "2.css")

	b, err := ioutil.ReadFile("compiled.css")
	a.Nil(err)
	fmt.Println(string(b))
	os.Remove("compiled.css")
}

func TestCompileDir(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()

	fmt.Println(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	r := New()
	r.CSSDir("compiled.css", "")

	b, err := ioutil.ReadFile("compiled.css")
	a.Nil(err)
	fmt.Println(string(b))
	os.Remove("compiled.css")
}

/*
//create a template loader engine
//it render all templates after any changes
//looks like Vue node.js local server...
r := templates.New()

//templates is your files
r.Template("home", "html/index.html")
r.Template("login", "html/login.html")
r.Template("product", "html/product.html")

//components is inquire to virtual version files
r.Component("@header", "html/header.html")
r.Component("@footer", "html/footer.html")

//dict is a constants
r.Dict("$company", "Apple Inc.")
r.Dict("$today", "")

//css will be optimized
r.CSS("compiled.css", "text.css", "colors.css", "buttons.css")

//all css files from directory will be compiled
r.CSSDir("compile.css", "html/css")
//get page
r.Get("home").
	Replace("$today", time.Now().String()).
	Replace("@good", "good").
	Bytes()


*/
