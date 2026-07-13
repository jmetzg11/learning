#include <stdint.h>
#include <stdio.h>

struct carModel {
  uint32_t carNumber;
  uint32_t carPrice;
  uint16_t carMaxSpeed;
  float carWeight;
};

int main(void) {
  struct carModel carBMW = {2021, 15000, 220, 1300};
  struct carModel carFord = {.carNumber = 4031,
                             .carPrice = 35000,
                             .carMaxSpeed = 160,
                             .carWeight = 1900.06};

  printf("Details of BMW is as follows:\n");
  printf("carnumber = %u\n", carBMW.carNumber);
  printf("carprice = %u\n", carBMW.carPrice);
  printf("carmax speed = %u\n", carBMW.carMaxSpeed);
  printf("car weight = %f\n", carBMW.carWeight);
  carBMW.carWeight = 12;
  printf("new car weight = %f\n", carBMW.carWeight);


  printf("\nsize of carFord %zu\n", sizeof(carFord));
  return 0;
}
