
public class _0_find_Kth_to_tail {
    public static void main(String[] args) {
    }

    public FindKthToTail(ListNode pListHead, int k) {
        if (pListHead == null) {
            return null;
        }

        ListNode pAhead = pListHead;
        ListNode pBhead = null;

        for (int i=0; i<k; i++) {
            if (pAhead == null) {
                return null;
            }
            pAhead = pAhead.m_pNext;
        }

        pBhead = pListHead;

        while (pAhead != null) {
            pAhead = pAhead.m_pNext;
            pBhead = pBhead.m_pNext;
        }
        return pBhead;

    }
}

class ListNode {
    public int val;
    public ListNode m_pNext;

    public ListNode(int val, ListNode m_pNext) {
        this.val = val;
        this.m_pNext = m_pNext;
    }

    public ListNode() {
        this(0, null);
    }
}
