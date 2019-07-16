import java.util.HashSet;
import java.util.Set;

public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null) return false;
        ListNode slow, fast;
        slow = fast = head;
        while (slow.next != null && fast.next != null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (fast == slow) return true;
        }
        return false;
    }
    public boolean hasCycleSet(ListNode head) {
        Set<ListNode> nodeMap = new HashSet<>();
        while (head != null) {
            if (nodeMap.contains(head) == true) return true;
            else nodeMap.add(head);
            head = head.next;
        }
        return false;
    }
}
class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
        next = null;
    }
}


