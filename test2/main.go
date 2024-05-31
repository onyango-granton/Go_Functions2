package main

import "fmt"

type Node struct {
	data         int
	nextNodeAddr *Node
}

type linkedList struct {
	headNodeAddress *Node
	length          int
}

func (l *linkedList) addNode(num int) {
	newNodeAddress := &Node{data: num}
	if l.headNodeAddress == nil {
		l.headNodeAddress = newNodeAddress
	} else {
		currentNodeAddress := l.headNodeAddress
		l.headNodeAddress = newNodeAddress
		l.headNodeAddress.nextNodeAddr = currentNodeAddress
	}
	l.length++
}

func (l *linkedList) display() {
	currentNodeAddress := l.headNodeAddress
	for currentNodeAddress != nil {
		fmt.Println(currentNodeAddress.data)
		currentNodeAddress = currentNodeAddress.nextNodeAddr
	}
}

func main() {
	listLink := &linkedList{}

	listLink.addNode(5)
	listLink.addNode(6)
	listLink.addNode(10)
	listLink.addNode(15)

	listLink.display()
}
