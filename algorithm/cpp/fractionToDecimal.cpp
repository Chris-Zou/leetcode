//
//  fractionToDecimal.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 7/31/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <string>
#include <vector>

using namespace std;

struct frac{
    long long result;
    long long remain;
};

string fractionToDecimal(int numerator, int denominator) {
    if (denominator == 0){
        return "";
    }
    
    bool isNegtive = false;
    
    long long numer = numerator;
    long long denomin = denominator;
    
    if (numer >0 && denomin < 0){
        isNegtive = true;
        denomin = -denomin;
    }else if(numer < 0 && denomin > 0){
        isNegtive = true;
        numer = -numer;
    }else if(numer <0 || denomin < 0){
        numer = -numer;
        denomin = -denomin;
    }
    
    long long beforDot = numer / denomin;
    numer = numer - beforDot * denomin;
    
    vector<frac> tmp;
    
    string ret;
    string tmpRet("");
    
    bool isFinished = false;
    
    while(numer != 0){
        long long tmpResult = (int)(10*numer / denomin);
        long long tmpRemain = (int)(10*numer - tmpResult * denomin);
        for(int i = (int)tmp.size() - 1; i >= 0; i--){
            if(tmp[i].remain == tmpRemain && tmp[i].result == tmpResult){
                for(int j = 0; j < i; j++){
                    tmpRet += to_string(tmp[j].result);
                }
                tmpRet += '(';
                for (int j = i; j < tmp.size(); j++){
                    tmpRet += to_string(tmp[j].result);
                }
                tmpRet += ')';
                isFinished = true;
                break;
            }
        }
        if (isFinished){
            break;
        }
        frac tmpf;
        tmpf.result = tmpResult;
        tmpf.remain = tmpRemain;
        tmp.push_back(tmpf);
        numer = tmpRemain;
    }
    
    if (isNegtive){
        ret += '-';
    }
    ret += to_string(beforDot);
    if (tmpRet != ""){
        ret += "." + tmpRet;
    }else if(numer == 0 && tmp.size() != 0){
        ret += ".";
        for(int i = 0; i < tmp.size(); i++){
            ret += to_string(tmp[i].result);
        }
    }
    
    return ret;
    
}

int main(int argc, char *argv[]){
    cout << fractionToDecimal(1, 3) << endl;
    cout << fractionToDecimal(-1, 3) << endl;
    cout << fractionToDecimal(-1, -3) << endl;
    cout << fractionToDecimal(31, 99) << endl;
    cout << fractionToDecimal(47, 99) << endl;
    cout << fractionToDecimal(2, 1) << endl;
    cout << fractionToDecimal(1, 5) << endl;
    cout << fractionToDecimal(1, 6) << endl;
    cout << fractionToDecimal(-1, -2147483648) << endl;
}
