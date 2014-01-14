package glogbridge

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// captureStderr captures the bytes written to os.Stderr during the execution of f.
func captureStderr(f func()) []byte {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	oldStderr := os.Stderr
	os.Stderr = w
	defer func() {
		os.Stderr = oldStderr
	}()

	c := make(chan []byte, 1)
	go func() {
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			panic(err)
		}
		c <- bytes
	}()

	f()
	w.Close()
	return <-c
}

func TestRedirection(t *testing.T) {
	flag.Set("logtostderr", "true")

	output := string(captureStderr(func() {
		log.Print("Hello, log!")
	}))
	pattern := "W%02d%02d %02d:%02d:%02d.%06d %05d glogbidge.go:%d] glogbridge_test.go:%d: Hello, log!\n"
	var month, day, hour, minute, second, micros, pid, line1, line2 int
	n, err := fmt.Sscanf(output, pattern, &month, &day, &hour, &minute, &second, &micros, &pid, &line1, &line2)
	if n != 9 || err != nil {
		t.Errorf("Logged line in unexpected format:\n%s", output)
	}
}
