#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>
#include <string.h>
#include <assert.h>
void print_green(const char *content) {
	printf("\033[32m %s\033[0m\n", content);
}
void print_red(const char *content) {
	printf("\033[31m %s\033[0m\n", content);
}
void print_purple(const char *content) {
	printf("\033[35m %s\033[0m\n", content);
}
#define col_num 3
#define str_num 5
#define max_height 40
#define max_line_len 80
#define max_print_len 200
void (*a[col_num])(const char *) = { print_green, print_red, print_purple };
char str_content[str_num] = {'#', '&', '8', '!', '@'};
char *constr_random_line(int count) {
	char array[max_line_len];
	memset(array, 0 , max_line_len);
	int i = 0;
	if(count > max_line_len || count < 0) exit(1);
	for (i = 0; i < count; ++i) {
		int index = rand() % str_num;
		array[i] = str_content[index];
	}
	array[i] = '\0';
	return strdup(array);
}
void print_christmas_tree(int height) {
	int i = 0;
	char print_buffer[max_print_len];
	char space_skip[max_height] = {0};
	memset(space_skip, ' ', max_height - height);
	for(;i < height; ++i){
		memset(print_buffer, 0 , max_print_len);
		int left_space = (height - i);
		int content_len = i * 2 + 1;
		int right_space = left_space;
		assert(height*2+1 < max_print_len);
		char space_buffer[max_height] = {0};
		memset(space_buffer, ' ', left_space);
		int length = 0;
		length += snprintf(print_buffer, max_print_len, "%s", space_skip);
		length += snprintf(print_buffer + length, max_print_len - length, "%s", space_buffer);
		char *content = constr_random_line(content_len);
		length += snprintf(print_buffer + length, max_print_len - length, "%s", content);
		length += snprintf(print_buffer + length, max_print_len - length, "%s", space_buffer);
		free(content);
		int color = rand() % col_num;
		a[color](print_buffer);
	}
}
void print_rect() {
	char print_buffer[max_print_len];
	char space_skip[max_height] = {0};
	memset(space_skip, ' ', 33);
	for(int i = 0; i < 5;++i) {
		memset(print_buffer, 0 , max_print_len);
		char * content = constr_random_line(15);
		int length = 0;
		length += snprintf(print_buffer, max_print_len, "%s", space_skip);
		length += snprintf(print_buffer + length, max_print_len - length, "%s", content);
		int color = i % col_num;
		a[color](print_buffer);
		free(content);
	}
}
int main()
{
	srand(time(NULL));
	while(1){
		for(int i =0; i < 1; ++i) {
			print_christmas_tree(10);
			print_christmas_tree(15);
			print_rect();
		}
		usleep(1000000);
		system("clear");
	}
	return 0;
}
