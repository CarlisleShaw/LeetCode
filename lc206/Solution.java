/****************************************
Reverse Linked List
When a chain is being reversed, it will be disconnected into 2 chains.
The head node of 2 chains needs to be recorded as 'prev' & 'cur'.
Before disconnecting 'cur' & 'cur.next', 'cur.next' needs to be recorded as 'next'.
           prev  cur   next
             |    |   /
origin:     NULL  1->2->3->4->5->NULL
processing: NULL<-1<-2  3->4->5->NULL
                    /   |   \
                prev   cur   next
Operations：
1.cur.next->next
2.connect prev & cur (disconnect cur & next)
3.move prev & cur 1 position to the right 
****************************************/
public class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode cur = head;
        ListNode prev = null;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}
