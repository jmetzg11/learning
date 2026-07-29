#include <stdint.h>
#include <stdio.h>

void array_display(uint8_t const *const pArray, uint32_t nItems);

int main(void) {
  uint8_t someData[10] = {0xff, 0xff, 0xff, 0xff, 0xff,
                          0xff, 0xff, 0xff, 0xff, 0xff};

  for (uint32_t i = 0; i < 10; i++) {
    someData[i] = i;
  };

  uint32_t nItems = sizeof(someData) / sizeof(uint8_t);
  array_display(someData, nItems);
  printf("\n");

  array_display((someData + 2), nItems-2);
  printf("\n");
  return 0;
}

void array_display(uint8_t const *const pArray, uint32_t nItems) {
  for (uint32_t i = 0; i < nItems; i++) {
    printf("%x\t", *(pArray + i));
    printf("%x\t",pArray[i]);
  }
}
