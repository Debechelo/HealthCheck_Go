package main

import (
	"fmt"
	"strings"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	items []Checkable
}

func (c *Checker) Add(items ...Checkable) {
	for _, item := range items {
		c.items = append(c.items, item)
	}
}

func (c Checker) String() string {
	str := strings.Builder{}
	for i := 0; i < len(c.items); i++ {
		str.WriteString(c.items[i].GetID() + " ")
	}

	return str.String()
}

func (c *Checker) Check() {
	for i := 0; i < len(c.items); i++ {
		if !c.items[i].Health() {
			fmt.Println(c.items[i].GetID() + " не работает")
		}
	}
}
