//
//  reverseList.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 8/14/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

ListNode* reverseList(ListNode* head) {
    if (head == NULL)
        return NULL;
    if (head->next == NULL)
        return head;
    
    ListNode* pre = NULL;
    ListNode* p = head;
    //ListNode* n = p->next;
    
    while(p != NULL){
        ListNode *tmp = p->next;
        p->next = pre;
        pre = p;
        p = tmp;
    }
    
    return pre;
}

void display(ListNode* head){
    ListNode *p = head;
    while(p != NULL){
        std::cout << p->val << " ";
        p = p->next;
    }
    std::cout << std::endl;
}

int main(int argc, char *argv[]){
    ListNode* l1 = new ListNode(1);
    ListNode* l2 = new ListNode(2);
    ListNode* l3 = new ListNode(3);
    l1->next = l2;l2->next = l3;
    display(l1);
    ListNode* ret = reverseList(l1);
    display(ret);
    
    
    
    return 0;
}