//Given the root of a binary tree and an integer targetSum, return true if the
//tree has a root-to-leaf path such that adding up all the values along the path
//equals targetSum.
//
// A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//Output: true
//Explanation: The root-to-leaf path with the target sum is shown.
//
//
// Example 2:
//
//
//Input: root = [1,2,3], targetSum = 5
//Output: false
//Explanation: There are two root-to-leaf paths in the tree:
//(1 --> 2): The sum is 3.
//(1 --> 3): The sum is 4.
//There is no root-to-leaf path with sum = 5.
//
//
// Example 3:
//
//
//Input: root = [], targetSum = 0
//Output: false
//Explanation: Since the tree is empty, there are no root-to-leaf paths.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 10
//102 👎 1164


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
    private static class Node {
        TreeNode node;
        int sum;

        public Node(TreeNode node, int sum) {
            this.node = node;
            this.sum = sum;
        }
    }

    public boolean hasPathSum(TreeNode root, int targetSum) {
        if (root == null) {
            return false;
        }

        // Method 1. DFS
        // if (root.left == null && root.right == null) {
        //     return root.val == targetSum;
        // }
        // return hasPathSum(root.left, targetSum - root.val) || hasPathSum(root.right, targetSum - root.val);

        // Method 2. BFS
        Queue<Node> queue = new LinkedList<>();
        queue.add(new Node(root, targetSum - root.val));

        while (!queue.isEmpty()) {
            Node current = queue.poll();
            TreeNode node = current.node;
            int currSum = current.sum;

            if (node.left == null && node.right == null && currSum == 0) {
                return true;
            }
            if (node.left != null) {
                queue.add(new Node(node.left, currSum - node.left.val));
            }
            if (node.right != null) {
                queue.add(new Node(node.right, currSum - node.right.val));
            }
        }

        return false;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
