package main

type _ListNode struct {
	val			int
	m_pNext		*_ListNode
}

/**
	desc：	求链表中的倒数第k个节点

	tips：	对象的创建、指针
 */


func (self *_ListNode) newListNode() *_ListNode {
	return &_ListNode{val:0, m_pNext:nil}
}

func FindKthToTail(pListHead *_ListNode, k int) *_ListNode {
	if pListHead == nil {
		return nil
	}

	pHead := pListHead
	var pBehind *_ListNode = pListHead

	for i:=0; i<k-1;  {
		if pHead == nil {
			return nil
		}
		pHead = pHead.m_pNext
	}

	for pHead != nil {
		pHead = pHead.m_pNext
		pBehind = pBehind.m_pNext
	}
	return pBehind
}



