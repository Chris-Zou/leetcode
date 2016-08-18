//
//  ImplQueueUsedStack.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 8/18/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <stack>

using namespace std;

class Queue {
public:
    // Push element x to the back of queue.
    void push(int x) {
        s1.push(x);
    }
    
    // Removes the element from in front of queue.
    void pop(void) {
        if(!s2.empty())
            s2.pop();
        else{
            while(!s1.empty()){
                s2.push(s1.top());
                s1.pop();
            }
            s2.pop();
        }
    }
    
    // Get the front element.
    int peek(void) {
        if(!s2.empty())
            return s2.top();
        else{
            while(!s1.empty()){
                s2.push(s1.top());
                s1.pop();
            }
            return s2.top();
        }
    }
    
    // Return whether the queue is empty.
    bool empty(void) {
        return s1.empty()&&s2.empty();
    }
    
private:
    stack<int> s1;
    stack<int> s2;
};

int main(int argc, char *argv[]){
    Queue q;
    q.push(1);
    q.push(2);
    q.push(3);
    q.pop();
    cout << q.peek() << endl;
    q.push(4);
    q.pop();
    cout << q.peek() << endl;
    
    return 0;
}
