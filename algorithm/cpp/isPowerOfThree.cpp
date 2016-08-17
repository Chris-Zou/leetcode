#include <iostream>

bool isPowerOfThree(int n) {
    if (n <= 0)
        return false;
                        
    return 1162261467 % n == 0;
}

int main(int argc, char *argv[]){
    
    bool ret = false;
    ret = isPowerOfThree(3);
    std::cout << ret << std::endl;


    ret = isPowerOfThree(4);
    std::cout << ret << std::endl;

    ret = isPowerOfThree(5);
    std::cout << ret << std::endl;

    return 0;
}
