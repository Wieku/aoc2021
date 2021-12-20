package util

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadLines(path string) []string {
	file, err := os.Open(filepath.Base(path))
	if err != nil {
		file, err = os.Open(path)
		if err != nil {
			panic(err)
		}
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
