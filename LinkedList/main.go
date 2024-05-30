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
	linkList := linkedList{}

	linkList.addNode(4)
	linkList.addNode(5)
	linkList.addNode(6)
	linkList.addNode(10)

	linkList.display()
}