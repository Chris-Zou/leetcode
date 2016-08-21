//
//  singleNumber3.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 8/21/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <vector>

using namespace std;


vector<int> singleNumber(vector<int>& nums) {
    int tmp = 0;
    size_t len = nums.size();
    
    vector<int> ret;
    if(len == 0)
        return ret;
    
    tmp = nums[0];
    for(size_t i = 1; i < len; i++){
        tmp = tmp ^ nums[i];
    }
    
    if(tmp == 0)
        return ret;
    
    int one = 1;
    for(int i =  0; i < 32; i++){
        if((tmp&one) != 0)
            break;
        one = one << 1;
    }
    
    int ret1 = 0, ret2 = 0;
    for(size_t i = 0; i < len; i++){
        if((nums[i] & one) == 0)
            ret1 = ret1 ^ nums[i];
        else
            ret2 = ret2 ^ nums[i];
    }
    
    ret.push_back(ret1);
    ret.push_back(ret2);
    
    return ret;
}

int main(int argc, char *argv[]){
    vector<int> v{1, 2, 1, 3, 2, 5};
    
    vector<int> ret = singleNumber(v);
    
    for(const int &i : ret)
        cout << i << " ";
    cout << endl;
    
    return 0;
}
