/*
 * main.c
 *
 *  Created on: Jul 4, 2026
 *      Author: jmetzg11
 */


#include<stdio.h>

int main(void) {
	char a, b, c, d, e, f;

	printf("Write 6 characters:");

	a = getchar();
	getchar();

	b = getchar();
	getchar();

	c = getchar();
	getchar();

	d = getchar();
	getchar();

	e = getchar();
	getchar();

	f = getchar();
	getchar();

	printf("\nASCII codes : %u,%u,%u,%u,%u,%u",a,b,c,d,e,f);

	while(getchar() != '\n')
	{

	}
	getchar();

	return 0;
}
