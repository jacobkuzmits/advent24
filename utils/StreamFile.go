package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func StreamFile(relativePath string, callback func(string)) {
	// Get the caller's directory (day1/, day2/, etc)
	_, callerFile, _, _ := runtime.Caller(1)
	callerDir := filepath.Dir(callerFile)

	// Resolve full path relative to caller
	fullPath := filepath.Join(callerDir, relativePath)

	f, err := os.OpenFile(fullPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		callback(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}
