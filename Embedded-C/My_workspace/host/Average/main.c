/*
 * main.c
 *
 *  Created on: Jul 4, 2026
 *      Author: jmetzg11
 */

#include <stdio.h>

int main(void)
{
	float num1, num2, num3;
	float average;

	printf("Enter the first number:");
	fflush(stdout);
	scanf("%f",&num1);

	printf("\nEnter the second number:");
	fflush(stdout);
	scanf("%f",&num2);

	printf("\nEnter the third number:");
	fflush(stdout);
	scanf("%f",&num3);

	average = (num1 + num2 + num3) / 3;

	printf("\nAverage is: %f\n",average);

	printf("\nPress any key to exit the application\n");
	while(getchar() != '\n')
	{

	}
	getchar();
	return 0;
}
