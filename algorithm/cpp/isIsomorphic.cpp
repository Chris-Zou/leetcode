//
//  isIsomorphic.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 7/31/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <set>

using namespace std;

void init(string s, set<vector<int> > &st){
    map<char, vector<int> > m;
    
    for(int i = 0; i < s.size(); i++){
        m[s[i]].push_back(i);
    }
    
    for(map<char, vector<int> >::iterator it = m.begin(); it != m.end(); it++){
        st.insert((*it).second);
    }
}

bool isIsomorphic(string s, string t) {
    set<vector<int> > ss, st;
    init(s, ss);
    init(t, st);
    
    if(ss.size() != st.size()){
        return false;
    }
    
    for(auto it = ss.begin(); it != ss.end(); ++it){
        auto tmpIt = st.find((*it));
        if (tmpIt == st.end())
            return false;
    }
    
    return true;
}

int main(int argc, char *argv[]){
    cout << isIsomorphic("egg", "add") << endl;
    cout << isIsomorphic("foo", "bar") << endl;
    cout << isIsomorphic("paper", "title") << endl;
}
