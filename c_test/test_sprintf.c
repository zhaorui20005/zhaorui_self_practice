#include <stdio.h>
int main()
{
	int n = 0;
	char buf[100] = {0};
	n = sprintf(buf, "%s", "name");
	printf("buf is %s, n is %d\n", buf, n);

	return 0;
}
