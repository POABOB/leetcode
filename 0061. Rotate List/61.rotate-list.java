//Given the head of a linked list, rotate the list to the right by k places.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5], k = 2
//Output: [4,5,1,2,3]
//
//
// Example 2:
//
//
//Input: head = [0,1,2], k = 4
//Output: [2,0,1]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 500].
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
//
//
// Related Topics Linked List Two Pointers 👍 10264 👎 1491


//leetcode submit region begin(Prohibit modification and deletion)

/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

public class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null) {
            return null;
        }
        int nums = 0;
        ListNode p = head;
        while (p != null) {
            nums++;
            p = p.next;
        }
        k = k % nums;
        if (k == 0 || nums == 1) {
            return head;
        }

        ListNode fast = head, slow = head;
        for (int i = 0; i < k; i++) {
            fast = fast.next;
        }
        while (fast.next != null) {
            slow = slow.next;
            fast = fast.next;
        }
        ListNode newHead = slow.next;
        slow.next = null;
        fast.next = head;
        return newHead;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
