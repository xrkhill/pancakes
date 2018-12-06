package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xrkhill/pancakes/flipit"
)

func main() {
	lines, err := getLinesFromStdin()

	if err != nil {
		fmt.Errorf("Encountered an error: %s.", err)
		os.Exit(1)
	}

	for i, line := range lines {
		// skip the first line of the file (test count)
		if i == 0 {
			continue
		}

		flips, err := flipit.FlipCount(line)

		if err != nil {
			fmt.Errorf("Encountered an error: %s.", err)
			os.Exit(1)
		}

		fmt.Printf("Case #%d: %d\n", i, flips)
	}
}

func getLinesFromStdin() ([]string, error) {
	var list []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	err := scanner.Err()

	if err != nil {
		fmt.Println("Error reading stdin:", err)
	}

	return list, err
}
