package common

//const char *Buildtime(void);
import "C"

/*format : 2018-12-09 15:26:26*/
func BuildTime() string {
	return C.GoString(C.Buildtime())
}
