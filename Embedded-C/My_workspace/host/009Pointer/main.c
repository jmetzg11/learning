
#include<stdio.h>

int main(void)
{
    char data = 100;
    printf("Value of the vairable is: %d\n", data);
    printf("Address of the vairable is: %p\n", &data);
   
   char*  pAddress = &data;
   char value = *pAddress;
   printf("read value is: %d\n", value); 

   *pAddress = 65;
   printf("Value of the vairable is: %d\n", data);
}