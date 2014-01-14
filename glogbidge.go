// Package glogbridge redirects the output from the builtin log package to the glog package.
package glogbridge

import (
	"log"

	"github.com/golang/glog"
)

// glogger implements io.Writer and redirects all writes to glog.Warning.
type glogger struct{}

func (g glogger) Write(b []byte) (int, error) {
	glog.Warning(string(b))
	return len(b), nil
}

func init() {
	log.SetOutput(glogger{})
	log.SetFlags(log.Lshortfile)
}
