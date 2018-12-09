package common

import (
	"fmt"
	"strings"
	"time"
)

//const char *Buildtime(void);
//const char *FixVersion(void);
import "C"

var markStartTime time.Time = time.Now()

// Largest time is 2540400:10:10
func RunningTime() string {
	// get second
	u := uint64(time.Since(markStartTime) / time.Second)

	second := u % 60
	u /= 60
	minutes := u % 60
	u /= 60

	return fmt.Sprintf("%02d:%02d:%02d", u, minutes, second)
}

/*format : 2018-12-09 15:26:26*/
func BuildTime() string {
	return C.GoString(C.Buildtime())
}

/*format : major.minor.buildDate - 1.0.20181209 */
func Version(major, minor string) string {
	version := []string{
		major,
		minor,
		C.GoString(C.FixVersion()),
	}

	return strings.Join(version, ".")
}
