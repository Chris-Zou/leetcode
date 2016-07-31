//
//  removeElements.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 7/31/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

ListNode* removeElements(ListNode* head, int val) {
    if (head == NULL)
        return NULL;
    
    ListNode* mockHead = new ListNode(0);
    mockHead->next = head;
    ListNode* pre = mockHead;
    ListNode* p = head;
    
    while(p != NULL){
        ListNode* tmp = p->next;
        if(p->val == val){
            pre->next = tmp;
            delete p;
            p = tmp;
        }else{
            pre = p;
            p = p->next;
        }
    }
    
    return mockHead->next;
}

void Display(ListNode* head){
    ListNode* p = head;
    
    while (p) {
        std::cout << p->val << " " ;
        p = p->next;
    }
    
    std::cout << std::endl;
}

int main(int argc, char *argv[]){
    ListNode *l1 = new ListNode(1);
    ListNode *l2 = new ListNode(1);
    
    l1->next = l2;
    Display(l1);
    Display(removeElements(l1, 1));
}
