package main

import "fmt"

type name interface {
	GetInput_name() string
	do_inputName() string
}

type Input_name struct {
	name string
}

func (i *Input_name) SetInput_name(name string) {
	i.name = name
}

func (i *Input_name) GetInput_name() string {
	return i.name
}

func (i *Input_name) do_inputName() string {
	fmt.Scan(&i.name)
	return i.name
}
