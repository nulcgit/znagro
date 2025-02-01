package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		return
	}
	programDir := filepath.Dir(exePath)

	lockFile := filepath.Join(programDir, "znagro.lock")

	file, err := os.OpenFile(lockFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Failed to create lock file:", err)
		return
	}
	defer file.Close()

	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		fmt.Println("Another instance of znagro is already running.")
		return
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	fmt.Println("Znagro - knowledge aggregator for the Myaoogle system")

	defer os.Remove(lockFile)
}
