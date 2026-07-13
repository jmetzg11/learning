#include<stdio.h>

int main(void)
{
	float chareE = 1.60217662e-19;
	printf("Number = %0.8f\n", chareE);
	printf("Number = %0.8e\n", chareE);

	double chareE2 = 1.60217662e-19;
	printf("Number = %0.28lf\n", chareE2);
	printf("Number = %0.8el\n", chareE2);

	return 0;
}
