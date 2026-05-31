package main

import "github.com/tr3nb0lone/chaos-client/runner"

func main() {
	opts := runner.ParseOptions()
	runner.RunEnumeration(opts)
}
