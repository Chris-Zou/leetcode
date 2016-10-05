//
//  rightSideView.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 10/5/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <vector>
#include <queue>
#include <iostream>

using namespace std;


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
class Solution {
public:
    vector<int> rightSideView(TreeNode* root) {
        vector<int> ret;
        
        if (root == NULL)
            return ret;
        
        queue<TreeNode*> qu;
        qu.push(root);
        int num = 1;
        int nextNum = 0;
        
        while(!qu.empty()){
            while(num > 0){
                TreeNode* tmp = qu.front();
                qu.pop();
                if(tmp->left != NULL){
                    qu.push(tmp->left);
                    nextNum++;
                }
                if(tmp->right != NULL){
                    qu.push(tmp->right);
                    nextNum++;
                }
                num--;
                
                if(num == 0){
                    ret.push_back(tmp->val);
                    num = nextNum;
                    nextNum = 0;
                }
            }
        }
        
        return ret;
        
    }
};

int main(int argc, char *argv[]){
    TreeNode* t1 = new TreeNode(1);
    TreeNode* t2 = new TreeNode(2);
    TreeNode* t3 = new TreeNode(3);
    //TreeNode* t4 = new TreeNode(4);
    TreeNode* t5 = new TreeNode(5);
    
    t1->left = t2;t1->right = t3;
    t2->right = t5;
    //t3->right = t4;
    
    Solution so;
    
    auto ret = so.rightSideView(t1);
    
    for(auto val : ret){
        cout << val << " ";
    }
}