#include <stdio.h>

int fabonaci(int i){
    if(i==0){
        return 0;
    }
    if(i==1){
        return 1;
    }
    return fabonaci(i-2)+fabonaci(i-1);
}

int main(int argc, int *argv[]){
    int *i,j;
    i=argv[1];
    printf("%s,%d", argv[1],i);
/*    for (j=0;j<i;j++){
        printf("%d\n",fabonaci(j));
    }*/
    return 0;
}
