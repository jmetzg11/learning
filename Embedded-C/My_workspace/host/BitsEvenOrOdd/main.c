#include<stdio.h>
#include<stdint.h>

int main(void){
    int32_t num;

    printf("Enter number:\n");
    scanf("%d",&num);

    if (num & 1) {
        printf("Number is odd\n");
    } else {
        printf("number is even\n");
    }

    return 1;
}