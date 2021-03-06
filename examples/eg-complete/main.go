package main

import (
	"fmt"
	"strings"

	"github.com/wxio/wxcli"
)

type Config struct {
	Alpha string
	Beta  myEnum
	Delta `wxcli:"mode=cmd"`
	Echo  `wxcli:"mode=cmd"`
}

type Delta struct {
	Zip bool
	Zop string
	Dir string
}

type Echo struct {
	Ping  bool
	Pong  string
	Files []string
}

func main() {
	c := Config{}
	wxcli.New(&c).
		Complete().
		Parse()
	fmt.Printf("%+v\n", c)
}

//myEnum implements completer and setter
//to provide an enum-like field
type myEnum string

var valid = []string{"a", "b", "c"}

func (myEnum) Complete(s string) []string {
	return valid
}

func (e *myEnum) Set(s string) error {
	for _, v := range valid {
		if s == v {
			*e = myEnum(s)
			return nil
		}
	}
	return fmt.Errorf("must be one of: %v", strings.Join(valid, ", "))
}
