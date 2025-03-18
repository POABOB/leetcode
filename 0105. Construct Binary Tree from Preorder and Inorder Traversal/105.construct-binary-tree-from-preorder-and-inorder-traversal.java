//Given two integer arrays preorder and inorder where preorder is the preorder
//traversal of a binary tree and inorder is the inorder traversal of the same tree,
// construct and return the binary tree.
//
//
// Example 1:
//
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//
//
// Example 2:
//
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//
//
// Constraints:
//
//
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder and inorder consist of unique values.
// Each value of inorder also appears in preorder.
// preorder is guaranteed to be the preorder traversal of the tree.
// inorder is guaranteed to be the inorder traversal of the tree.
//
//
// Related Topics Array Hash Table Divide and Conquer Tree Binary Tree 👍 15703
//👎 561


//leetcode submit region begin(Prohibit modification and deletion)

import java.util.HashMap;
import java.util.Map;

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

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        Map<Integer, Integer> hashmap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            hashmap.put(inorder[i], i);
        }
        return build(preorder, inorder, hashmap, 0, preorder.length - 1, 0, inorder.length - 1);
    }

    public TreeNode build(int[] preorder, int[] inorder, Map<Integer, Integer> hashmap, int preStart, int preEnd, int inStart, int inEnd) {
        if (preStart > preEnd) {
            return null;
        }
        int rootValue = preorder[preStart];
        int i = hashmap.get(rootValue);
        TreeNode root = new TreeNode(rootValue);
        root.left = build(preorder, inorder, hashmap, preStart + 1, preStart + i - inStart, inStart, i - 1);
        root.right = build(preorder, inorder, hashmap, preStart + i - inStart + 1, preEnd, i + 1, inEnd);
        return root;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
