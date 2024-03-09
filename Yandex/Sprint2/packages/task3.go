package main

import (
	"fmt"
	"os"
	"strconv"
)

type test struct {
	args       []string
	expectFail bool
}

func run() error {
	if len(os.Args) != 4 {
		return fmt.Errorf("Usage: %s <grid_size_x> <grid_size_y> <initial_fill_percentage>", os.Args[0])
	}

	gridSizeX, err := strconv.Atoi(os.Args[1])
	if err != nil || gridSizeX <= 0 {
		return fmt.Errorf("grid_size_x must be a positive integer")
	}

	gridSizeY, err := strconv.Atoi(os.Args[2])
	if err != nil || gridSizeY <= 0 {
		return fmt.Errorf("grid_size_y must be a positive integer")
	}

	fillPercentage, err := strconv.Atoi(os.Args[3])
	if err != nil || fillPercentage <= 0 || fillPercentage > 100 {
		return fmt.Errorf("initial_fill_percentage must be an integer between 0 and 100")
	}

	file, err := os.Create("config.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%dx%d %d%%", gridSizeX, gridSizeY, fillPercentage)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	tests := []test{
		{
			args:       []string{"10", "20", "30"},
			expectFail: false,
		},
		{
			args:       []string{"-1", "-3", "-6"},
			expectFail: true,
		},
		{
			args:       []string{"1", "-3", "6"},
			expectFail: true,
		},
		{
			args:       []string{"-1", "3", "6"},
			expectFail: true,
		},
		{
			args:       []string{"1", "3", "-6"},
			expectFail: true,
		},
		{
			args:       []string{"0", "3", "6"},
			expectFail: true,
		},
		{
			args:       []string{"1", "0", "6"},
			expectFail: true,
		},
		{
			args:       []string{"44", "3", "0"},
			expectFail: true,
		},
		{
			args:       []string{"5", "5", "lala"},
			expectFail: true,
		},
		{
			args:       []string{"laa", "5", "10"},
			expectFail: true,
		},
		{
			args:       []string{"5", "lala", "10"},
			expectFail: true,
		},
	}

	for _, tc := range tests {
		os.Args = append([]string{"main"}, tc.args...)
		err := run()
		if tc.expectFail && err == nil {
			fmt.Printf("Test failed. Expected failure but got success for args: %v\n", tc.args)
		} else if !tc.expectFail && err != nil {
			fmt.Printf("Test failed. Expected success but got failure for args: %v. Error: %v\n", tc.args, err)
		}
	}
}
