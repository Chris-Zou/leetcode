#include <stdio.h>
#include <stdlib.h>

int getBits(int n){
	int ret = 0;
	while (n != 0){
		n = n & (n-1);
		ret++;
	}
	return ret;
}

int* countBits(int num, int* returnSize) {
	int *ret = (int*)malloc((num+1) * sizeof(int));
	*returnSize = num+1;
    for(int i = 0; i <= num; i++){
		ret[i] = getBits(i);
	}
	
	return ret;
}

int main(){
	int *ret, nums;
	ret = countBits(5, &nums);
	printf("%d\n", nums);
	for (int i = 0; i < nums; i++){
		printf("%d ", ret[i]);
	}
	printf("\n");
}