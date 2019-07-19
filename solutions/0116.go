package solutions

import (
	"reflect"

	"pokman/bulbasaur/leetcode/ds"
)

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() {}

    Node(int _val, Node* _left, Node* _right, Node* _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
class Solution {
public:
    Node* connect(Node* root) {
        recurTree(root);
        return root;
    }

    void recurTree(Node* node) {
        if (node == NULL || node->left == NULL || node->right == NULL) {
            return;
        }

        node->left->next = node->right;
        if (node->next != NULL) {
            node->right->next = node->next->left;
        }

        recurTree(node->left);
        recurTree(node->right);
    }
};
*/

func connectBinaryTree(root *ds.TreeNode) {
}

func init() {
	desc := `
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.
	`
	sol := Solution{
		Title:  "Populating Next Right Pointers in Each Node",
		Desc:   desc,
		Method: reflect.ValueOf(connectBinaryTree),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0116"] = sol
}
