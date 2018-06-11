#include<stdio.h>
int main(){
    printf("fib 5 is %d \n", fib(5));
    printf("fib 6 is %d \n", fib(6));
    printf("fib 7 is %d \n", fib(7));
    printf("fib 8 is %d \n", fib(8));
}
int fib(int n){
    if(n == 1 || n == 2 ){
	return 1;
    }
    return fib(n-1) + fib(n-2);
}
