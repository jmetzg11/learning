
#include <stdint.h>
#include <stdio.h>

int main(void) {
  int32_t num1, num2;
  printf("Enter two numbers:");
  scanf("%d %d", &num1, &num2);

  printf("And &: %d\n", (num1 & num2));
  printf("Or |: %d\n", (num1 | num2));
  printf("XOR ^: %d\n", (num1 ^ num2));
  printf("Negate ~: num1: %d, num2: %d\n", ~num1, ~num2);
}
