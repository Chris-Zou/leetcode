#include<iostream>

bool canWinNim(int n) {
    if (n % 4 == 0){
        return false;
    }
    else{
        return true;
    }
}

int main(){
    std::cout << canWinNim(1) << std::endl;
    std::cout << canWinNim(2) << std::endl;
    std::cout << canWinNim(3) << std::endl;
    std::cout << canWinNim(4) << std::endl;
    std::cout << canWinNim(5) << std::endl;
    std::cout << canWinNim(6) << std::endl;
    std::cout << canWinNim(7) << std::endl;
    std::cout << canWinNim(8) << std::endl;
}
