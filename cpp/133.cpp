/*
// Definition for a Node.

*/
#include <iostream>
#include <vector>
using namespace std;

class Node {
public:
    int val;
    vector<Node*> neighbors;

    Node() {}

    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

class Solution {
    Node* visited[101] = {nullptr};
public:
    Node* cloneGraph(Node* node) {
        int len = node->neighbors.size();
        Node* res = new Node(node->val, vector<Node*>() );
        visited[node->val] = res;
        
        for (int i=0;i < len; i ++) {
            if (!visited[node->neighbors[i]->val])
                res->neighbors.push_back(cloneGraph(node->neighbors[i]));
            else 
                res->neighbors.push_back(visited[node->neighbors[i]->val]);
        }     
        return res;  
    }
};