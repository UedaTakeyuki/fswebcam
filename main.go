package fw

/*
#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>
#include <locale.h>
#include <gd.h>
#include <errno.h>
#include <signal.h>
#include <sys/types.h>
#include <sys/stat.h>
#include "fswebcam.h"
#include "log.h"
#include "src.h"
#include "dec.h"
#include "effects.h"
#include "parse.h"

#cgo LDFLAGS: -lgd
#cgo CFLAGS: -g -O2 -DHAVE_CONFIG_H

extern int cmain(int argc, char *argv[]);

*/
import "C"

import "strings"

//func Capture(commandLine string) {
func Capture(args []string) {
//	args := strings.Split(commandLine, " ") // fswebcam -d v4l2:/dev/video0 kerokero.jpg
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	C.cmain(argc, &argv[0])
}
