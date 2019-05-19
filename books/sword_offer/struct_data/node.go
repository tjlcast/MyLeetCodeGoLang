package struct_data

import "fmt"

type ListNode struct {
	m_nValue		int
	m_pNext			*ListNode
}


/**
	在传入的链表的尾部插入一个值，
	并返回改变后链表的root
 */
func AddToTail(pHead *ListNode, value int) *ListNode {
	pRoot := pHead
	pNew := &ListNode{}
	pNew.m_nValue = value
	pNew.m_pNext = nil

	if pHead == nil {
		pRoot = pNew
	} else {
		for pHead.m_pNext != nil {
			pHead = pHead.m_pNext
		}
		pHead.m_pNext = pNew
	}

	return pRoot
}

/**
	在链表中找到第一个含有输入值的节点并删除该节点
 */
func RemoveNode(pHead *ListNode, value int) {
	if pHead == nil {
		return
	}

	pRoot := &ListNode{}
	pRoot.m_pNext = pHead

	for pRoot.m_pNext != nil {
		if pRoot.m_pNext.m_nValue == value {
			pLast := pRoot.m_pNext.m_pNext
			pRoot.m_pNext.m_pNext = nil
			pRoot.m_pNext = pLast
			break
		}
		pRoot = pRoot.m_pNext
	}
}

/**
	从头到尾打印链表
 */

func PrintNodes(pHead *ListNode) {
	if pHead == nil {
		fmt.Println("")
	}

	for pHead != nil {
		fmt.Println(pHead.m_nValue)
		pHead = pHead.m_pNext
	}
}