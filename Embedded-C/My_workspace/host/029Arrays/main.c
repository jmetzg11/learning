#include <stdint.h>
#include <stdio.h>

int main() {

  // arrays are just references to data
  uint8_t studentsAge8[100];
  printf("Size of an array = %zu\n", sizeof(studentsAge8));

  uint32_t studentsAge32[100];
  printf("Size of an array = %zu\n", sizeof(studentsAge32));
  printf("Base address of the array = %p\n", studentsAge32);

  // initialize an array
  int len = 10;
  uint8_t dataLength[len];

  // writing to the value of array 
  uint8_t someData[10] = {0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff};
  printf("Before 2nd data item = %x\n",*(someData+1));
  *(someData+1) = 0x9;
  printf("After 2nd data item = %x\n",*(someData+1));
  someData[1] = 0x6;
  printf("After 2nd data item = %x\n",*(someData+1));
  for (uint32_t i = 0; i < sizeof(someData); i++ ){
      printf("%x\t",someData[i]);
  }
  
  printf("\n");
  return 0;
}
