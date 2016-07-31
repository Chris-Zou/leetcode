//
//  isHappy.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 7/31/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <set>

using namespace std;

int calcSum(int n){
    int ret = 0;
    
    while (n != 0) {
        int tmp = n % 10;
        ret += tmp * tmp;
        n = n / 10;
    }
    
    return ret;
}

bool isHappy(int n) {
    set<int> s;
    s.insert(n);
    
    while(n != 1){
        int tmp = calcSum(n);
        set<int>::iterator it = s.find(tmp);
        if (it != s.end()){
            return false;
        }
        s.insert(tmp);
        n = tmp;
    }
    
    return true;
}

int main(int argc, char *argv[]){
    cout << isHappy(19) << endl;
}