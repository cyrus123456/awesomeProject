package main

import "fmt"

type Node struct {
	num  int
	name string
	next *Node
}

func addNode(head *Node, newNode *Node) {
	if head.next == nil {
		head.num = newNode.num
		head.name = newNode.name
		head.next = head
		fmt.Printf("%v添加到环形列表\n", newNode.name)
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = head

}

func delNode(head *Node, name string) {
	temp := head
	for {
		if temp.next.name == name {
			temp.next = temp.next.next
			break
		}
	}
}

func show(head *Node) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	fmt.Println("该链表存储的内容")
	for {
		fmt.Printf("%v\n", *temp)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	head := &Node{}

	tom_Node := &Node{
		num:  1,
		name: "tom",
	}
	addNode(head, tom_Node)

	jerry_Node := &Node{
		num:  2,
		name: "jerry",
	}
	addNode(head, jerry_Node)

	kobe_Node := &Node{
		num:  3,
		name: "kobe",
	}
	addNode(head, kobe_Node)

	kyrei_Node := &Node{
		num:  4,
		name: "kyrei",
	}
	addNode(head, kyrei_Node)

	show(head)
	delNode(head, "kobe")
	show(head)

}
