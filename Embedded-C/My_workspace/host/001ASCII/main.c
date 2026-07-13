/*
 * main.c
 *
 *  Created on: Jul 4, 2026
 *      Author: jmetzg11
 */


#include <stdio.h>

int main(void)
{
	char a, b, c, d, e, f;
	printf("Write 6 characters:");
	scanf("%c %c %c %c %c %c",&a,&b,&c,&d,&e,&f);

	printf("first %c: %d\n",a,a);
	printf("second %c: %d\n",b,b);
	printf("third: %d\n",c);
	printf("fourth: %d\n",d);
	printf("fifth: %d\n",e);
	printf("sixth: %d\n",f);

	printf("Press any key to exit\n");
	while(getchar() != '\n')
	{

	}
	getchar();

	return 0;
}
