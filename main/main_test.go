package main

import (
	"errors"
	"fmt"
	"testing"
)

type testConfig struct {
	args []string
	err error
	config
}


func TestParseArgs(t *testing.T) {
	// TODO -> Insert definition tests[] array
tests := []testConfig{
    {
		args: []string{"-h"},
		err: nil,
		config: config{
			printUsage: true, numTimes: 0,
		},
	},
    {
		args: []string{"10"},
		err: nil,
		config: config {
			printUsage: false,
			numTimes: 10,
		},
	},
    {
		args: []string{"abc"},
		err: errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
		config: config{
			printUsage: false,
			numTimes: 0,
		},
	},
    {
		args: []string{"1", "foo"},
		err: errors.New("Invalid number of arguments"),
		config: config{
			printUsage: false,
			numTimes: 0,
		},
	},
}

	fmt.Print(tests)
}
