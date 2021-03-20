#include "stdc++.h"
using namespace std;

// 剑指07: 重建二叉树 (根据前序加中序)

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return buildTreeRecur(preorder, inorder, 0, preorder.size()-1, 0, inorder.size()-1);
    }

    TreeNode* buildTreeRecur(vector<int>& preorder, vector<int>& inorder, int prel, int prer, int inl, int inr) {
        if (inl > inr || prel > prer) {
            return NULL;
        }

        TreeNode* root = new TreeNode(preorder[prel]);
        for (int i = inl; i <= inr; ++i) {
            if (inorder[i] == preorder[prel]) {
                root->left = buildTreeRecur(preorder, inorder, prel+1, prer, inl, i-1);
                // 左子树个数是i-inl个
                root->right = buildTreeRecur(preorder, inorder, prel+1+i-inl, prer, i+1, inr);
                break;
            }
        }
        return root;
    }

};