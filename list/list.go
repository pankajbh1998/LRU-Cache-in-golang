package list

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func (n *Node) RemoveNode()*Node{
	left := n.left
	right := n.right
	if left != nil {
		left.right=right
	}
	if right != nil {
		right.left = left
	}
	return left
}

func (n *Node) AddNodeToHead(new *Node){
	right := n.right
	if right != nil {
		right.left = new
	}
	new.right=right
	new.left= n
	n.right=new
}

func NewNode(val int)*Node{
	return &Node{
		 val:   val,
		 left:  nil,
		 right: nil,
	}
}

func (n *Node) GetValue()int{
	return n.val
}

func(n *Node)Print(){
	head := n.right
	for head != nil {
		fmt.Print(head.val, ", ")
		head = head.right
	}
}
