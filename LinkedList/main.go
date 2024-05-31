package main

import (
	"fmt"
)

type Node struct{
	data int
	nextNode *Node
}

type linkedList struct{
	head *Node
}

func (l *linkedList) addNode(num int){
	newNode := Node{data: num}
	if l.head == nil{
		l.head = &newNode
	} else {
		currentNode := l.head
		l.head = &newNode
		l.head.nextNode = currentNode
	}
}

func (l *linkedList) display() {
	currentNode := l.head
	for currentNode != nil{
		fmt.Println(currentNode.data)
		currentNode = currentNode.nextNode
	}
}

func main() {
	linkListA := linkedList{}

	linkListA.addNode(4)
	linkListA.addNode(5)
	linkListA.addNode(6)
	linkListA.addNode(10)

	linkListA.display()
}