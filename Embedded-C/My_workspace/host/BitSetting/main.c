
#include<stdio.h>
#include<stdint.h>

int main(void){
    int32_t num, output;

    printf("Enter number:\n");
    scanf("%d",&num);

    output = num | 0x90;
    printf("[input] [output] : 0x%x 0x%x\n",num,output);

    return 1;
}