#include <string.h>
#include <stdio.h>
int main()
{
	char *a = "hello";
	char *b = "hel";
	for(int i = 0; i < 8;++i)
	printf("The result of print strncmp(a,b,%d) is %d\n", i,strncmp(a,b,i));
	return 0;
}
