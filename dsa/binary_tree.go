package dsa

import "strconv"

var visits []int

// 重置访问列表
func resetVisits() {
	visits = []int{}
}

// 数组转字符串
func visitsStr() string {
	str := ""
	for _, v := range visits {
		str += strconv.Itoa(v) + ","
	}
	str = str[:len(str)-1]
	return str
}

type BTNode struct {
	data  int
	left  *BTNode
	right *BTNode
}

func NewBTNode(data int) *BTNode {
	return &BTNode{
		data: data,
	}
}

func (n *BTNode) AddLeft(ln *BTNode) {
	n.left = ln
}

func (n *BTNode) AddRight(rn *BTNode) {
	n.right = rn
}

// 先序遍历
func (n *BTNode) PreOrder() {
	if n == nil {
		return
	}
	n.Visit()
	n.left.PreOrder()
	n.right.PreOrder()
}

// TODO: 非递归先序遍历
func (n *BTNode) PreOrderNonrecursion() {

}

// 中序遍历
func (n *BTNode) InOrder() {
	if n == nil {
		return
	}
	n.left.InOrder()
	n.Visit()
	n.right.InOrder()
}

// TODO: 非递归中序遍历
func (n *BTNode) InOrderNonrecursion() {

}

// 后序遍历
func (n *BTNode) PostOrder() {
	if n == nil {
		return
	}
	n.left.PostOrder()
	n.right.PostOrder()
	n.Visit()
}

// TODO: 非递归后序遍历
func (n *BTNode) PostOrderNonrecursion() {

}

// 层序遍历
func (n *BTNode) Level(qSize int) {
	if n == nil {
		return
	}
	var front, rear int
	queue := make([]*BTNode, qSize)
	rear++
	queue[rear] = n
	for front != rear { // 队列非空，表示还有未遍历元素
		front = (front + 1) % qSize
		n := queue[front] // 队头出栈
		n.Visit()
		if n.left != nil { // 左孩子进栈
			rear = (rear + 1) % qSize
			queue[rear] = n.left
		}
		if n.right != nil { // 右孩子进栈
			rear = (rear + 1) % qSize
			queue[rear] = n.right
		}
	}
}

func (n *BTNode) Visit() {
	visits = append(visits, n.data)
}
