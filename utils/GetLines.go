package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
)

func GetLines(relativePath string) ([]string, error) {
	// Get the caller's directory (day1/, day2/, etc)
	_, callerFile, _, _ := runtime.Caller(1)
	callerDir := filepath.Dir(callerFile)

	// Resolve full path relative to caller
	fullPath := filepath.Join(callerDir, relativePath)

	f, err := os.OpenFile(fullPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
