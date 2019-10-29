/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    map<int,int> vis;
    int kk = 0;
    int r = -1;
    void traversal(TreeNode* root,int k)
    {
        if (root==NULL)
            return;
        vis[root->val] = 1;
        traversal(root->left,k);
        if (vis[root->val] == 1)
            kk++;
        if (kk==k)
            r = root->val;
        traversal(root->right,k);
        return;
    }
    int kthSmallest(TreeNode* root, int k) {
        traversal(root,k);
        return r;
    }
};
