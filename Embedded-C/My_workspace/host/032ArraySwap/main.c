#include <stdint.h>
#include <stdio.h>

void wait_for_user_input(void);
void display_array(uint32_t *pArray, uint32_t nItem);
void swap_arrays(uint32_t *array1, uint32_t *array2, uint32_t nitem1,
                 uint32_t nitem2);

int main(void) {
  int32_t nItem1, nItem2;

  printf("Array swapping program\n");
  printf("Enter no of elements of Array-1 and Array-2: ");
  scanf("%d %d", &nItem1, &nItem2);

  if ((nItem1 < 0) || (nItem2 < 0)) {
    printf("Number of elements cannot be negative\n");
    wait_for_user_input();
    return 0;
  }

  uint32_t array1[nItem1];
  uint32_t array2[nItem2];

  for (int32_t i = 0; i < nItem1; i++) {
    printf("Enter %d element of array1:", i + 1);
    scanf("%u", &array1[i]);
  }
  for (int32_t i = 0; i < nItem2; i++) {
    printf("Enter %d element of array2:", i + 1);
    scanf("%u", &array2[i]);
  }

  printf("Contents before swap\n");
  display_array(array1, (uint32_t)nItem1);
  printf("\n");
  display_array(array2, (uint32_t)nItem2);
  printf("\n");

  swap_arrays(array1, array2, nItem1, nItem2);

  printf("Contents after swap\n");
  display_array(array1, (uint32_t)nItem1);
  printf("\n");
  display_array(array2, (uint32_t)nItem2);
  printf("\n");

  return 0;
}

void wait_for_user_input(void) { getchar(); }

void display_array(uint32_t *pArray, uint32_t nItem) {
  for (uint32_t i = 0; i < nItem; i++) {
    printf("%4u\t", pArray[i]);
  }
}

void swap_arrays(uint32_t *array1, uint32_t *array2, uint32_t nitem1,
                 uint32_t nitem2) {
  uint32_t len = nitem1 < nitem2 ? nitem1 : nitem2;
  for (uint32_t i = 0; i < len; i++) {
    int32_t temp = array1[i];
    array1[i] = array2[i];
    array2[i] = temp;
  }
}
