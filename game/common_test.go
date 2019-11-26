package game

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"strings"
)

const (
	TestFilePath  = "/invasion/assets/small_world.txt"
	TestNumAliens = 5
)

// LoadFilePath : loads the sample test file world.txt for testing
func LoadFilePath() string {
	_, currentDirPath, _, _ := runtime.Caller(0)
	splitPath := strings.Split(currentDirPath, "/invasion/")
	filePath := splitPath[0] + TestFilePath
	return filePath
}

// CaptureOutput
func CaptureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
