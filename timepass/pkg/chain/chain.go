package chain

import (
	"fmt"
	"linked-lists-2/pkg/link"
)

type Chain struct {
	head *link.Link
	tail *link.Link
}

func NewChain() *Chain {
	return &Chain{}
}

func (c *Chain) AddLink(data int) {
	if c.tail == nil {
		c.head = link.NewLink(data)
		c.tail = c.head
		return
	}

	for c.tail.GetNextLink() != nil {
		c.tail = c.tail.GetNextLink()
		continue
	}
	c.tail.AddNewLink(link.NewLink(data))
	c.tail = c.tail.GetNextLink()
}

func (c *Chain) GetLinks() {
	c.tail = c.head
	for c.tail.GetNextLink() != nil {
		fmt.Println(c.tail.GetData())
		c.tail = c.tail.GetNextLink()
	}

	fmt.Println(c.tail.GetData())

}
