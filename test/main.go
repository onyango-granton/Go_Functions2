package main

import "fmt"

type Node struct {
	data            int
	nextNodePointer *Node
}

type linkedList struct {
	headNodePointer *Node
	listLength      int
}

func (l *linkedList) addNode(num int) {
	newNode := &Node{data: num}
	if l.headNodePointer == nil {
		l.headNodePointer = newNode
	} else {
		currentHeadNode := l.headNodePointer
		l.headNodePointer = newNode
		l.headNodePointer.nextNodePointer = currentHeadNode
	}
	l.listLength++
}

func (l *linkedList) display() {
	currentHead := l.headNodePointer
	for currentHead != nil {
		fmt.Println(currentHead.data)
		currentHead = currentHead.nextNodePointer
	}
}

func main() {
	linkList := linkedList{}

	linkList.addNode(4)
	linkList.addNode(6)
	linkList.addNode(8)
	linkList.addNode(10)

	linkList.display()
}
