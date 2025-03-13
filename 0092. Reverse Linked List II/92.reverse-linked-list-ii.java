// Given the head of a singly linked list and two integers left and right where
// left <= right, reverse the nodes of the list from position left to position
// right, and return the reversed list.
//
//
//  Example 1:
//
//
// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]
//
//
//  Example 2:
//
//
// Input: head = [5], left = 1, right = 1
// Output: [5]
//
//
//
//  Constraints:
//
//
//  The number of nodes in the list is n.
//  1 <= n <= 500
//  -500 <= Node.val <= 500
//  1 <= left <= right <= n
//
//
//
// Follow up: Could you do it in one pass?
//
//  Related Topics Linked List 👍 12102 👎 688

// leetcode submit region begin(Prohibit modification and deletion)

import java.util.List;

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
    private ListNode successor;

    // Method 1. 遍歷
    // public ListNode reverseBetween(ListNode head, int left, int right) {
    //     ListNode dummy = new ListNode(0);
    //     dummy.next = head;
    //     ListNode p = dummy;
    //     for (int i = 0; i < left - 1; i++) {
    //         p = p.next;
    //     }

    //     ListNode prev = null;
    //     ListNode curr = p.next;
    //     for (int i = 0; i < right - left + 1; i++) {
    //         ListNode next = curr.next;
    //         curr.next = prev;
    //         prev = curr;
    //         curr = next;
    //     }
    //     p.next.next = curr;
    //     p.next = prev;
    //     return dummy.next;
    // }

    // Method 2. 遞迴
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (left == 1) {
            return reverseN(head, right);
        }
        head.next = reverseBetween(head.next, left - 1, right - 1);
        return head;
    }

    private ListNode reverseN(ListNode head, int n) {
        if (n == 1) {
            successor = head.next;
            return head;
        }
        ListNode last = reverseN(head.next, n - 1);
        head.next.next = head;
        head.next = successor;
        return last;
    }
}
// leetcode submit region end(Prohibit modification and deletion)
