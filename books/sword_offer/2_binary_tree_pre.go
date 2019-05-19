package main

import "fmt"

/**
		desc:	1、链表的创建、链表三序的访问
				2、链表的重建（依赖前、中序数组）

		tips:	对切片进行拆包【append([]int{0}, nums...)】
 */

type BinaryTreeNode struct {
	m_nValue			int
	m_pLeft				*BinaryTreeNode
	m_pRight			*BinaryTreeNode
}

func newBinaryTreeNode(n int) *BinaryTreeNode {
	return &BinaryTreeNode{n, nil, nil}
}

func buildBinaryTree(nums []int, pos int) *BinaryTreeNode {
	// validate.
	if nums==nil || len(nums)==0 || pos>=len(nums) {
		return nil
	}

	// make a new node for first ele.
	val := nums[pos]
	node := newBinaryTreeNode(val)

	// build subtrees.
	node.m_pLeft = buildBinaryTree(nums, pos*2)
	node.m_pRight = buildBinaryTree(nums,pos*2 + 1)

	return node
}

func BinaryTreePreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d\t", root.m_nValue)
	BinaryTreePreOrder(root.m_pLeft)
	BinaryTreePreOrder(root.m_pRight)
}

func BinaryTreeInOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	BinaryTreeInOrder(root.m_pLeft)
	fmt.Printf("%d\t", root.m_nValue)
	BinaryTreeInOrder(root.m_pRight)
}


func BinaryTreeAfOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	BinaryTreeAfOrder(root.m_pLeft)
	BinaryTreeAfOrder(root.m_pRight)
	fmt.Printf("%d\t", root.m_nValue)
}

func BinaryTreeRebuild(pNums []int, lstart, llen int, iNums []int, istart, ilen int) *BinaryTreeNode {
	if llen == 0 {
		return nil
	}

	val := pNums[lstart]
	node := &BinaryTreeNode{val, nil, nil}

	i:=0
	for ; i<ilen; i++ {
		if val == iNums[istart+i] {
			break
		}
	}

	node.m_pLeft = BinaryTreeRebuild(pNums, lstart+1, i, iNums, istart, i)
	node.m_pRight = BinaryTreeRebuild(pNums, lstart+1+i, llen-1-i, iNums, istart+i+1, llen-1-i)

	return node
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := buildBinaryTree(append([]int{0}, nums...), 1)

	fmt.Println("PreOrder: ")
	BinaryTreePreOrder(root)

	fmt.Println("InOrder: ")
	BinaryTreeInOrder(root)

	fmt.Println("AfOrder: ")
	BinaryTreeAfOrder(root)

	nums1 := []int{1, 2, 4, 7, 3, 5, 6, 8}
	nums2 := []int{4, 7, 2, 1, 5, 3, 8, 6}
	root = BinaryTreeRebuild(nums1, 0, len(nums1), nums2, 0, len(nums2))
	fmt.Printf("\nthe result:\n")
	BinaryTreePreOrder(root)
}

