## groups example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/wxio/wxcli"
)

type Config struct {
	Fizz       string
	Buzz       bool
	Foo        `wxcli:"group=Foo"`
	Ping, Pong int `wxcli:"group=More"`
}

type Foo struct {
	Bar  int
	Bazz int
}

func main() {
	c := Config{}
	wxcli.Parse(&c)
	fmt.Printf("%+v\n", c)
}
```
<!--/tmpl-->

Group order in the help text is first-use order

```
$ eg-groups --help
```

<!--tmpl,code=plain:go build -o eg-groups && ./eg-groups --help ; rm eg-groups -->
``` plain 

  Usage: eg-groups [options]

  Options:
  --fizz, -f
  --buzz, -b
  --help, -h  display help

  Foo options:
  --bar
  --bazz

  More options:
  --ping, -p
  --pong

```
<!--/tmpl-->
