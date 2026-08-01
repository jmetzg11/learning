#include<stdio.h>

int main(){
    // scanf does not care about spaces
    // int a, b, c;
    // printf("Enter 3 numbers;\n");
    // scanf("%d%d%d",&a,&b,&c);
    // printf("Numbers are: %d %d %d\n",a,b,c);


    // tedious method
    // char fname[30], lname[30];
    // printf("Enter your full name:\n");
    // scanf("%s%s",fname,lname);
    // printf("Your name is: %s %s\n",fname, lname);

    char name[30];
    printf("Enter your full name:\n");
    scanf("%[^\n]",name);
    printf("Your name is: %s\n",name);
    return 0;
}