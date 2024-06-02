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

func (l *linkedList) deleteHead() {
	headNode := l.headNodeAddress
	l.headNodeAddress = headNode.nextNodeAddr
}

func (l *linkedList) selfDel(num int) {
	currentNodeAddr := l.headNodeAddress
	for currentNodeAddr != nil {
		//if currentNodeAddr.nextNodeAddr.nextNodeAddr.nextNodeAddr == nil{
		//if currentNodeAddr.nextNodeAddr.nextNodeAddr.data == num {

		//	l.deleteTail()

		//	break
		//}
		if currentNodeAddr.data == num {
			currentNodeAddr.data = currentNodeAddr.nextNodeAddr.data
			currentNodeAddr.nextNodeAddr = currentNodeAddr.nextNodeAddr.nextNodeAddr
		}
		currentNodeAddr = currentNodeAddr.nextNodeAddr
	}
}

//Detailed commit message


func (l *linkedList) deleteTail() {
	currentNode := l.headNodeAddress
	for currentNode != nil {
		if currentNode.nextNodeAddr.nextNodeAddr.nextNodeAddr == nil {
			currentNode.nextNodeAddr.nextNodeAddr = nil
			break
		}
		currentNode = currentNode.nextNodeAddr
	}
}

func main() {
	listLink := &linkedList{}

	listLink.addNode(5)
	listLink.addNode(6)
	listLink.addNode(10)
	listLink.addNode(15)
	//listLink.deleteTail()
	//fmt.Println(listLink.headNodeAddress.nextNodeAddr.nextNodeAddr.nextNodeAddr)
	//listLink.selfDel(15)
	listLink.deleteHead()
	listLink.display()
}
