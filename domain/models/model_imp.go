package models

import "fmt"

type ModelImpl interface {
	ToString() string
}

type Model struct {
	Id   int
	Name string
}

func (this *Model) SetId(id int) {
	this.Id = id
}

func (this *Model) SetName(name string) {
	this.Name = name
}

func (this *Model) ToString() string {
	return fmt.Sprint("Model is %s,id is %d", this.Name, this.Id)
}
