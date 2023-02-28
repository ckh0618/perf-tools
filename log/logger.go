package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"time"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

var (
	log *logrus.Logger
)

func init() {

	log = logrus.New()
	log.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        time.RFC3339, // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)
	log.SetReportCaller(true)
}

// Info ...
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

var (

	// ConfigError ...
	ConfigError = "%v type=config.error"

	// HTTPError ...
	HTTPError = "%v type=http.error"

	// HTTPWarn ...
	HTTPWarn = "%v type=http.warn"

	// HTTPInfo ...
	HTTPInfo = "%v type=http.info"
)
