package t707

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	Node *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Node: &Node{val: -1, next: nil},
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.Size || index < 0 {
		return -1
	}

	i := 0
	p := this.Node
	for p != nil {
		if i == index {
			return p.val
		}

		p = p.next
		i++
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := &Node{
		val:  val,
		next: this.Node,
	}
	this.Node = head
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	p := this.Node
	if this.Size == 0 {
		this.Node = &Node{
			val:  val,
			next: this.Node,
		}
		this.Size++
		return
	}
	for i := 0; i < this.Size-1; i++ {
		p = p.next
	}
	p.next = &Node{val: val, next: p.next}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index > this.Size {
		return
	}

	p := this.Node
	for i := 0; i < index-1; i++ {
		p = p.next
	}
	p.next = &Node{
		val:  val,
		next: p.next,
	}
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}
	if index == 0 {
		this.Node = this.Node.next
		this.Size--
		return
	}

	i := 0
	p := this.Node
	for p != nil {
		if i == index-1 {
			p.next = p.next.next
			this.Size--
			return
		}

		p = p.next
		i++
	}
}

func ExampleConstructor() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1))

	obj2 := Constructor()
	obj2.AddAtTail(1)
	fmt.Println(obj2.Get(0))

	// ["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
	// [[],            [2],        [1],            [2],         [7],        [3],       [2],        [5],        [5],        [5],   [6],            [4]]
	obj3 := Constructor()
	obj3.AddAtHead(2)
	obj3.DeleteAtIndex(1)
	obj3.AddAtHead(2)
	obj3.AddAtHead(7)
	obj3.AddAtHead(3)
	obj3.AddAtHead(2)
	obj3.AddAtHead(5)
	obj3.AddAtTail(5)
	fmt.Println(obj3.Get(5))
	obj3.DeleteAtIndex(6)
	obj3.DeleteAtIndex(4)

	// Output:
	// 2
	// 3
	// 1
	// 2
}
