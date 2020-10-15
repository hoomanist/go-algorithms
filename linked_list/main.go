package main

import "fmt"

type queue struct {
	head *node
}

type node struct {
	cargo int
	next *node
}

func (list *queue) First() *node {
	return list.head
}

func (list *queue) push(cargo int) {
	newNode := node{cargo: cargo, next:nil}
	if list.head != nil {
		list.head.next = &newNode
		list.head = &newNode
	}else {
		list.head = &newNode
	}
}

func (current_node *node) Next() *node {
	return current_node.next
}

func main(){
	list := queue{}
	for i := 0; i <= 9; i++ {
		list.push(i)
	}
	fmt.Println(list.head.cargo)
}
