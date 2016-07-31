//
//  countPrimes.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 7/31/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <vector>
#include <math.h>

using namespace std;

int countPrimes(int n) {
    bool *Del = new bool[n];
    Del[2] = false;
    
    for(int i=3; i<n; ++i)
        if(i%2 == 0)
            Del[i] = true;
        else
            Del[i] = false;
    
    for(int i=3; i<n; i+=2)
        if(!Del[i])
        {
            if(i*i>n)  break;
            for(int j=2; i*j<n; ++j)
                Del[i*j] = true;
        }
    
    int count = 0;
    for(int i=2; i<n; ++i)
        if(!Del[i])
            ++count;
    
    delete [] Del;
    return count;
}

int main(int argc, char *argv[]){
    cout << countPrimes(15000000) << endl;
}
