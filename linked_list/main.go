package main

import "fmt"

type queue struct {
	head *node
    tail *node
    length int
}

type node struct {
	cargo int
	next *node
}

func (list *queue) Traverse() {
    n := list.head
    for  i := 0; i < list.length + 1; i++ {
        if n == nil {
            fmt.Println("")
            return
        }
        fmt.Printf("%d: %d - ", i, n.cargo)
        n = n.next
    }
}


func (list *queue) push(cargo int) {
    n := &node{cargo: cargo}
    list.length++
    if list.tail == list.head && list.tail == nil{
        list.tail = n
        list.head = list.tail
        n.next = nil 
    } else {
        tail := list.tail
        list.tail = n
        tail.next = list.tail
    }
}

func (list *queue) pop() *node {
    if list.head == nil {
        return nil
    }
    list.length--
    h := list.head
    list.head = h.next
    return h

}

func main(){
	list := queue{}
	for i := 0; i <= 9; i++ {
		list.push(i)
	}
    list.Traverse()
    fmt.Println(list.pop().cargo)
    list.Traverse()
}
