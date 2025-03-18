//Given the root of a binary tree, check whether it is a mirror of itself (i.e.,
// symmetric around its center).
//
//
// Example 1:
//
//
//Input: root = [1,2,2,3,4,4,3]
//Output: true
//
//
// Example 2:
//
//
//Input: root = [1,2,2,null,3,null,3]
//Output: false
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
//
//
//
//Follow up: Could you solve it both recursively and iteratively?
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 15
//911 👎 408


//leetcode submit region begin(Prohibit modification and deletion)

import java.util.LinkedList;
import java.util.Queue;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    public boolean isSymmetric(TreeNode root) {
        // Method 1. DFS
        // if (root == null) return true;
        // return isMirror(root.left, root.right);

        // Method 2. BFS
        if (root == null) return true;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root.left);
        queue.add(root.right);

        while (!queue.isEmpty()) {
            TreeNode n1 = queue.poll();
            TreeNode n2 = queue.poll();

            if (n1 == null && n2 == null) continue;
            if (n1 == null || n2 == null || n1.val != n2.val) return false;

            queue.add(n1.left);
            queue.add(n2.right);
            queue.add(n1.right);
            queue.add(n2.left);
        }
        return true;
    }

    private boolean isMirror(TreeNode n1, TreeNode n2) {
        if (n1 == null && n2 == null) return true;
        if (n1 == null || n2 == null) return false;
        return (n1.val == n2.val) && isMirror(n1.left, n2.right) && isMirror(n1.right, n2.left);

    }
}
//leetcode submit region end(Prohibit modification and deletion)

