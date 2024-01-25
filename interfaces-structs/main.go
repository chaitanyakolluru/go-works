package main

import "fmt"

type book struct {
	author, title string
}

type magazine struct {
	title, issue string
}

func (b *book) assign(a, t string) {
	b.author = a
	b.title = t
}

func (m *magazine) assign(t, i string) {
	m.title = t
	m.issue = i
}

func (b *book) print() {
	fmt.Println(b.author, b.title)
}

func (m *magazine) print() {
	fmt.Println(m.title, m.issue)
}

type printer interface {
	print() // interface created with one function print() that is satisfied by both structs book and magazine
}

func main() {
	var b book
	var m magazine

	b.assign("chai", "time")
	m.assign("adventure", "v12.01")

	fmt.Println("just regularly", b, m)
	fmt.Println("calling b and m from print functions from objects.")
	b.print()
	m.print()

	var i printer
	i = &b
	i.print()
	i = &m
	i.print()
}
