package main

import "fmt"

type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

func Oposite(nodeValue int) int {
	return 1 - nodeValue
}

func SingleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	if rootNode == nil || rootNode.LinkedNodes[nodeValue] == nil {
		return rootNode
	}
	saveNode := rootNode.LinkedNodes[nodeValue]
	if saveNode == nil {
		return rootNode
	}

	rootNode.LinkedNodes[nodeValue] = saveNode.LinkedNodes[Oposite(nodeValue)]
	saveNode.LinkedNodes[Oposite(nodeValue)] = rootNode

	return saveNode
}

func DoubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	if rootNode == nil || rootNode.LinkedNodes[Oposite(nodeValue)] == nil {
		return rootNode
	}

	child := rootNode.LinkedNodes[Oposite(nodeValue)]
	if child == nil || child.LinkedNodes[nodeValue] == nil {
		return rootNode
	}

	rootNode.LinkedNodes[Oposite(nodeValue)] = SingleRotation(rootNode.LinkedNodes[Oposite(nodeValue)], nodeValue)
	return SingleRotation(rootNode, Oposite(nodeValue))
}

func AdjustBalance(rootNode *TreeNode, nodeValue int, balanceValue int) {
	node := rootNode.LinkedNodes[nodeValue]
	if node == nil {
		return
	}
	oppIndex := Oposite(nodeValue)
	oppNode := node.LinkedNodes[oppIndex]
	if oppNode == nil {
		return
	}

	rootNode.BalanceValue, node.BalanceValue = 0, 0

	switch oppNode.BalanceValue {
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
	case -balanceValue:
		node.BalanceValue = balanceValue
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	}

	oppNode.BalanceValue = 0
}

func BalanceTree(rootNode *TreeNode, nodeValue int) *TreeNode {
	if rootNode == nil || rootNode.LinkedNodes[nodeValue] == nil {
		return rootNode
	}

	node := rootNode.LinkedNodes[nodeValue]
	if node == nil {
		return rootNode
	}
	balance := 2*nodeValue - 1

	if rootNode.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return SingleRotation(rootNode, nodeValue)
	}

	AdjustBalance(rootNode, nodeValue, balance)
	return DoubleRotation(rootNode, Oposite(nodeValue))
}

func InsertRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}
	dir := 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}

	var done bool
	rootNode.LinkedNodes[dir], done = InsertRNode(rootNode.LinkedNodes[dir], key)

	if done {
		return rootNode, true
	}

	rootNode.BalanceValue += Oposite(dir)*2 - 1

	if rootNode.BalanceValue == 0 {
		return rootNode, true
	}

	if rootNode.BalanceValue == 1 || rootNode.BalanceValue == -1 {
		return rootNode, false
	}

	return BalanceTree(rootNode, dir), true
}

func InsertNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = InsertRNode(*treeNode, key)
}

func RemoveBalance(rootNode *TreeNode, nodeValue int) (*TreeNode, bool) {
	oppNodeValue := Oposite(nodeValue)
	node := rootNode.LinkedNodes[oppNodeValue]
	if node == nil {
		return rootNode, true
	}

	balance := 2*nodeValue - 1

	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return SingleRotation(rootNode, nodeValue), false

	case balance:
		AdjustBalance(rootNode, oppNodeValue, -balance)
		return DoubleRotation(rootNode, nodeValue), false
	}

	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return SingleRotation(rootNode, nodeValue), true
}

func RemoveRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}

	if rootNode.KeyValue.EqualTo(key) {
		if rootNode.LinkedNodes[0] == nil {
			return rootNode.LinkedNodes[1], false
		}
		if rootNode.LinkedNodes[1] == nil {
			return rootNode.LinkedNodes[0], false
		}
		heirNode := rootNode.LinkedNodes[0]
		parent := rootNode
		for heirNode.LinkedNodes[1] != nil {
			parent = heirNode
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		if parent == rootNode {
			parent.LinkedNodes[0] = heirNode.LinkedNodes[0]
		} else {
			parent.LinkedNodes[1] = heirNode.LinkedNodes[0]
		}
	}
	dir := 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}

	var done bool
	rootNode.LinkedNodes[dir], done = RemoveRNode(rootNode.LinkedNodes[dir], key)

	if done {
		return rootNode, true
	}

	rootNode.BalanceValue += 1 - 2*dir

	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}

	return RemoveBalance(rootNode, dir)
}

func RemoveNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = RemoveRNode(*treeNode, key)
}

type integerKey int

func (k integerKey) LessThan(k1 KeyValue) bool { return k < k1.(integerKey) }
func (k integerKey) EqualTo(k1 KeyValue) bool  { return k == k1.(integerKey) }

// Fungsi untuk mencetak AVL Tree secara inorder
func PrintInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PrintInOrder(root.LinkedNodes[0])
	fmt.Print(root.KeyValue, " ")
	PrintInOrder(root.LinkedNodes[1])
}

func main() {
	var root *TreeNode

	// Insert beberapa nilai
	values := []integerKey{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		InsertNode(&root, v)
	}

	fmt.Println("AVL Tree setelah insert (inorder traversal):")
	PrintInOrder(root)
	fmt.Println()

	// Hapus beberapa node
	toDelete := []integerKey{30, 40}
	for _, v := range toDelete {
		RemoveNode(&root, v)
		fmt.Printf("AVL Tree setelah menghapus %d:\n", v)
		PrintInOrder(root)
		fmt.Println()
	}
}
