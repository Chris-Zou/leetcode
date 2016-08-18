//
//  isAnagram.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 8/18/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <map>
#include <string>

using namespace std;

bool isAnagram(string s, string t) {
    size_t len = s.size();
    if(len != t.size())
        return false;
    map<char, int> ms, mt;
    for(size_t i = 0; i < len; i++){
        ms[s[i]]++;
        mt[t[i]]++;
    }
    
    if (ms.size() != mt.size())
        return false;
    
    for(map<char, int>::iterator it = ms.begin(); it != ms.end(); ++it){
        if((*it).second != mt[(*it).first])
            return false;
    }
    
    return true;
}

int main(){
    cout << isAnagram("anagram", "nagaram") << endl;
    cout << isAnagram("rat", "car") << endl;
    
    return 0;
}
