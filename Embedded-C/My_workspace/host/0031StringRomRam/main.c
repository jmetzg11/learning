#include <stdio.h>

int main(void) {
  // string
  char msg1[] = "Hello how are you?";
  msg1[0] = 'b';
  printf("Message is: %s\n", msg1);
  printf("Address of 'msg1' variable = %p\n", &msg1);
  printf("Value of 'msg1' variable = %p\n", msg1);

  // literal string 
  char const *msg2 = "fastbitlab.com";
  // msg2[0] = 'b';
  printf("Message is: %s\n", msg2);
  printf("Address of 'msg2' variable = %p\n", &msg2);
  printf("Value of 'msg2' variable = %p\n", msg2);
  
  return 0;
}
