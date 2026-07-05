/*
 * main_new.c
 *
 *  Created on: Jul 4, 2026
 *      Author: jmetzg11
 */


#include <stdio.h>

int main(void)
{
	float num1, num2, num3;
	float average;

	printf("Enter 3 numbers:");
	scanf("%f %f %f",&num1,&num2,&num3);

	average = (num1 + num2 + num3) / 3;

	printf("\nAverage is: %f\n",average);

	printf("\nPress any key to exit the application\n");
	while(getchar() != '\n')
	{

	}
	getchar();
	return 0;
}
