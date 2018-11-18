#include <stdlib.h>
#include "include/capi/cef_app_capi.h"

int main(int argc, char **argv) {
	cef_main_args_t *args = (cef_main_args_t *)calloc(1, sizeof(cef_main_args_t));
	args->argc = argc;
	args->argv = argv;
	return cef_execute_process(args, NULL, NULL);
}
