## help example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import "github.com/wxio/wxcli"

type HelpConfig struct {
	Zip  string `wxcli:"mode=arg" help:"<zip> is a required arg which lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	Foo  string `help:"this is help for foo"`
	Bar  int    `help:"and help for bar"`
	Fizz string `help:"lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus at commodo odio. Sed id tincidunt purus. Cras vel felis dictum, lobortis metus a, tempus tellus, and fizz"`
	Buzz bool   `help:"and help for buzz"`
}

func main() {

	c := HelpConfig{
		Foo: "42",
	}

	wxcli.New(&c).
		Name("help").
		Summary("The help program demonstrates how to customise the help text").
		Version("1.0.0").
		Repo("https://github.com/wxio/foo").
		Parse()
}
```
<!--/tmpl-->

```
$ eg-help --help
```

<!--tmpl,code=plain:go build -o eg-help && ./eg-help --help ; rm eg-help -->
``` plain 

  Usage: help [options] <zip>

  The help program demonstrates how to customise the help text

  <zip> is a required arg which lorem ipsum dolor sit amet, consectetur adipiscing elit

  Options:
  --foo, -f      this is help for foo (default 42)
  --bar, -b      and help for bar
  --fizz         lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus at commodo odio.
                 Sed id tincidunt purus. Cras vel felis dictum, lobortis metus a, tempus tellus,
                 and fizz
  --buzz         and help for buzz
  --version, -v  display version
  --help, -h     display help

  Version:
    1.0.0

  Read more:
    https://github.com/wxio/foo

```
<!--/tmpl-->
