package main

import "fmt"


type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
	length int
}


func (l *LinkedList) addNode(num int){
	newNode := Node{data: num}
	if l.head == nil {
		l.head = &newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil{
			currentNode = currentNode.next
		}
		currentNode.next = &newNode
	}
	l.length ++
}

func (l *LinkedList) Display(){
	current := l.head
	for i:=0; i<l.length;i++{
		fmt.Println(current.data)
		current = current.next
	}
}


func main(){
	list := LinkedList{}
	list.addNode(4)
	list.addNode(5)
	list.addNode(6)
	list.addNode(7)

	fmt.Println(list.head, list.length)
	list.Display()
}