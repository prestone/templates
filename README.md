# Templates Reloader Engine
Very cool templates reloader. Its really help in development websites. Just add your html files, css, js anything and make development local easy.

```go
//create a template loader engine
//it render all templates after any changes
//looks like Vue node.js local server...
r := templates.New()

//add your templates
r.Template("home", "html/index.html")
r.Template("login", "html/login.html")
r.Template("product", "html/product.html")

//add components 
//its injects to virtual version of templates
//original file still original
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
```

Just a simple examples of engine. You can create any var names, I like @@name style;)

Your original html template. With vue for example
```html
@@header
<p>Welcome to @@company<p>
@@footer
```

Component Header
```html
<div>Header</div>
```

Component Footer
```html
<div>Footer</div>
```

Dict
```go
r.Dict("@@company", "Apple Inc.")
```

Get page
```go
r.Get("home")
```

Compiled
```html
<div>Header</div>
<p>Welcome to Apple Inc.<p>
<div>Footer</div>
```
