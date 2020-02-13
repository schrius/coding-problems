package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func insert(head *Node, val int) *Node {
	ptr := head
	temp := &Node{val, nil}

	if head == nil {
		head = temp
	} else {
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = temp
	}
	return head
}

func removeKthLastElement(head *Node, k int) *Node {
	len := 0
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.next
	}

	if k > len {
		fmt.Printf("k %d is larger than the length of the list", k)
		return head
	} else if k == len {
		return head.next
	} else {
		dif := len - k
		var prev *Node
		cur := head
		for i := 0; i < dif; i++ {
			prev = cur
			cur = cur.next
		}
		prev.next = cur.next
		return head
	}
}

func dispaly(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main() {
	var head *Node
	head = insert(head, 1)
	head = insert(head, 2)
	head = insert(head, 4)
	head = insert(head, 5)
	head = insert(head, 8)

	dispaly(head)
	// remove the second last element "5"
	head = removeKthLastElement(head, 2)

	dispaly(head)
}
