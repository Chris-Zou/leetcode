//
//  intersection.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 8/21/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <vector>
#include <set>
#include <map>

using namespace std;

vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
    set<int> st;
    for(vector<int>::iterator it = nums1.begin(); it != nums1.end(); it++){
        st.insert(*it);
    }
    
    set<int> ret;
    
    for(vector<int>::iterator it = nums2.begin(); it != nums2.end(); it++){
        set<int>::iterator bResult = st.find(*it);
        if (bResult != st.end()){
            ret.insert(*it);
        }
    }
    
    vector<int> finalRet;
    for(set<int>::iterator it = ret.begin(); it != ret.end(); it++){
        finalRet.push_back(*it);
    }
    
    return finalRet;
}

int main(int argc, char *argv[]){
    vector<int> v1{1, 2, 2, 1};
    vector<int> v2{2, 2};
    
    vector<int> ret = intersection(v1, v2);
    for(const int &a : ret){
         cout << a << " ";
    }
    cout << endl;
    
    return 0;
}