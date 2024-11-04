package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

// Run a Go file in the current directory
func runGoFile() {
    cmd := exec.Command("sh", "-c", "go run .")
	output, err := cmd.Output()
	if err != nil {
		color.Red("error executing command: %s", err)
		return
	}
	fmt.Println(string(output))
}

// Compile and run a C file
func runCFile(fileName string) {
	cmd := exec.Command("gcc", fileName, "-o", "output")
	if err := cmd.Run(); err != nil {
		color.Red("error compiling C file: %s", err)
		return
	}
	cmd = exec.Command("./output")
	output, err := cmd.Output()
	if err != nil {
		color.Red("error executing C program: %s", err)
		return
	}
	fmt.Println(string(output))
}

// Compile and run a C++ file
func runCppFile(fileName string) {
	cmd := exec.Command("g++", fileName, "-o", "output")
	if err := cmd.Run(); err != nil {
		color.Red("error compiling C++ file: %s", err)
		return
	}
	cmd = exec.Command("./output")
	output, err := cmd.Output()
	if err != nil {
		color.Red("error executing C++ program: %s", err)
		return
	}
	fmt.Println(string(output))
}

// Run a Python file
func runPythonFile(fileName string) {
	cmd := exec.Command("python", fileName)
	output, err := cmd.Output()
	if err != nil {
		color.Red("error executing Python file: %s", err)
		return
	}
	fmt.Println(string(output))
}

func main() {
	currentDir, err := os.Getwd() // Get the current working directory
	if err != nil {
		color.Red("error:", err)
		panic(err)
	}

	watcher, err := fsnotify.NewWatcher() // Create a new file watcher
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close() // Ensure the watcher is closed when done

	err = watcher.Add(currentDir) // Watch the current directory
	if err != nil {
		log.Fatal(err)
	}

	color.Green("Watching \"%s\"...", currentDir)

	var lastEventTime time.Time
	const debounceDuration = 200 * time.Millisecond // Debounce duration to avoid rapid triggers

	for {
		select {
		case ev := <-watcher.Events: // Handle file system events
			now := time.Now()
			fileName := ev.Name

            // Remove backup file suffix
			if fileName[len(fileName)-1] == '~' {
				fileName = fileName[:len(fileName)-1]
			}

			// Check if the file exists and debounce the event
			_, err := os.Stat(fileName)
			if os.IsNotExist(err) || now.Sub(lastEventTime) < debounceDuration {
				continue
			}

			lastEventTime = now
			log.Printf("event on file: %s", fileName)

			// Run the appropriate function based on file extension
			if filepath.Ext(fileName) == ".go" {
				runGoFile()
			} else if filepath.Ext(fileName) == ".c" {
				runCFile(fileName)
			} else if filepath.Ext(fileName) == ".cpp" {
				runCppFile(fileName)
			} else if filepath.Ext(fileName) == ".py" {
				runPythonFile(fileName)
			} else {
				color.Red("file support coming soon ....") // Unsupported file type
				panic("")
			}

		case err := <-watcher.Errors: // Handle errors from the watcher
			log.Println("error:", err)
		}
	}
}
