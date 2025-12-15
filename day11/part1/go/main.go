package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Device struct {
	name    string
	outputs []string
}

var devices = make(map[string]*Device)
var visiting = make(map[string]bool)

func getOrAddDevice(name string) *Device {
	if dev, exists := devices[name]; exists {
		return dev
	}
	dev := &Device{name: name, outputs: []string{}}
	devices[name] = dev
	return dev
}

func countPaths(current string) int64 {
	if current == "out" {
		return 1
	}

	dev, exists := devices[current]
	if !exists || len(dev.outputs) == 0 {
		return 0
	}

	if visiting[current] {
		return 0
	}

	visiting[current] = true
	var total int64 = 0
	for _, output := range dev.outputs {
		total += countPaths(output)
	}
	visiting[current] = false

	return total
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open list.txt:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		rest := strings.TrimSpace(parts[1])

		dev := getOrAddDevice(name)

		outputs := strings.Fields(rest)
		dev.outputs = append(dev.outputs, outputs...)
	}

	result := countPaths("you")
	fmt.Println(result)
}
