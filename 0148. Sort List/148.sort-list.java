//Given the head of a linked list, return the list after sorting it in
//ascending order.
//
//
// Example 1:
//
//
//Input: head = [4,2,1,3]
//Output: [1,2,3,4]
//
//
// Example 2:
//
//
//Input: head = [-1,5,3,4,0]
//Output: [-1,0,3,4,5]
//
//
// Example 3:
//
//
//Input: head = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 5 * 10⁴].
// -10⁵ <= Node.val <= 10⁵
//
//
//
// Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.
//e. constant space)?
//
// Related Topics Linked List Two Pointers Divide and Conquer Sorting Merge
//Sort 👍 12214 👎 384


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

class Pair<T1, T2> {
    final T1 a;
    final T2 b;

    Pair(T1 a, T2 b) {
        this.a = a;
        this.b = b;
    }
}

class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null) return null;

        // 計算整體 list 長度
        ListNode curr = head;
        int count = 0;
        while (curr != null) {
            curr = curr.next;
            count++;
        }

        // 每次進行分組，分組次數為 O(logn)，n 為 list 長度
        ListNode dummy = new ListNode();
        dummy.next = head;
        for (int i = 1; i <= count; i *= 2) {
            // 將組內 list 進行排序，O(n)
            curr = dummy.next;
            ListNode tail = dummy;
            while (curr != null) {
                ListNode left = curr;
                ListNode right = nextN(i, left);
                curr = nextN(i, right); // 剩下的 list
                Pair<ListNode, ListNode> sortedPartialList = merge(left, right);
                // tail 主要是用來連接 "組與組" 之間的
                tail.next = sortedPartialList.a;
                tail = sortedPartialList.b;
            }
        }
        return dummy.next;
    }

    private ListNode nextN(int n, ListNode root) {
        if (root == null) return null;
        while (n > 1 && root.next != null) {
            root = root.next;
            n--;
        }
        ListNode next = root.next;
        root.next = null;
        return next;
    }

    // sort 並合併兩 list，並 return head 與 tail
    private Pair<ListNode, ListNode> merge(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode();
        ListNode curr = dummy;
        while (l1 != null && l2 != null) {
            if (l1.val > l2.val) {
                curr.next = l2;
                l2 = l2.next;
            } else {
                curr.next = l1;
                l1 = l1.next;
            }
            curr = curr.next;
        }
        if (l1 != null) curr.next = l1;
        if (l2 != null) curr.next = l2;

        while (curr.next != null) {
            curr = curr.next;
        }
        return new Pair<>(dummy.next, curr);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
