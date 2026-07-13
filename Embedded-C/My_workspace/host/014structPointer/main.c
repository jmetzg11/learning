#include <stdio.h>

struct DataSet {
  char data1;
  int data2;
  char data3;
  short data4;
};

void displayMemberElements(struct DataSet *pdata);

int main(void) {
  struct DataSet data;

  data.data1 = 0x11;
  data.data2 = 0xFFFFEEEE;
  data.data3 = 0x22;
  data.data4 = 0xABCD;

  struct DataSet *pData;
  pData = &data;

  printf("Before data.data1 = %d\n", data.data1);
  pData->data1 = 0x55;
  printf("After data.data1 = %d\n", data.data1);

  displayMemberElements(&data);

  return 0;
}

void displayMemberElements(struct DataSet *pData) {
  printf("data1 from helper = %X\n", pData->data1);
  printf("data2 from helper = %X\n", pData->data2);
}
