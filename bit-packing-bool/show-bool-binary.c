#include <stdbool.h>
#include <stdio.h>

typedef unsigned char byte;
void byte_binary_representation(byte val) {
  for (byte i = 1 << 7; i > 0; i = i / 2) {
    if ((val & i) != 0) {
      printf("1");
    } else {
      printf("0");
    }
  }
  printf("\n");
}

int main(void) {
  printf("Binary representation for false: ");
  byte_binary_representation(false);

  printf("Binary representation for true:  ");
  byte_binary_representation(true);

  // for (byte i = 0; i < 255; i++) {
  //   byte_binary_representation(i);
  // }
  return 0;
}