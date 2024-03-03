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

void set_nth_bit(byte *data, byte n, bool value) {
  if (n < 8) {
    if (value) {
      *data = *data | (1 << n);
    } else {
      *data = *data ^ (1 << n);
    }
    return;
  }
}

int main(void) {
  byte data = 0;
  printf("Initialized data value:   ");
  byte_binary_representation(data);
  printf("---------------------------------\n");

  for (byte i = 0; i < 8; i++) {
    data = data | (1 << i);
    printf("Setting %d. bool to true:  ", i + 1);
    byte_binary_representation(data);
  }
  printf("---------------------------------\n");

  for (byte i = 0; i < 8; i++) {
    data = data ^ (1 << i);
    printf("Setting %d. bool to false: ", i + 1);
    byte_binary_representation(data);
  }
  printf("---------------------------------\n");

  printf("Setting 0. bit to high: ");
  set_nth_bit(&data, 0, true);
  byte_binary_representation(data);

  printf("Setting 3. bit to high: ");
  set_nth_bit(&data, 3, true);
  byte_binary_representation(data);

  printf("Setting 7. bit to high: ");
  set_nth_bit(&data, 7, true);
  byte_binary_representation(data);

  printf("Setting 3. bit to low:  ");
  set_nth_bit(&data, 3, false);
  byte_binary_representation(data);

  printf("Setting 7. bit to low:  ");
  set_nth_bit(&data, 7, false);
  byte_binary_representation(data);

  printf("Setting 0. bit to low:  ");
  set_nth_bit(&data, 0, false);
  byte_binary_representation(data);

  return 0;
}