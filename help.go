package main

import (
	"github.com/nathanjsharpe/trakit/output"
)

func Help(args []string) {
	data := [][]string{
		[]string{"trakit environments", "List possible trakit environments"},
		[]string{"trakit landmark", "Landmarks resource"},
		[]string{"trakit login", "Log into Trakit"},
		[]string{"trakit logout", "Log out of Trakit"},
		[]string{"trakit session", "Display details of current session"},
		[]string{"trakit user", "Users resource"},
		[]string{"trakit version", "Print version of Trakit CLI and services for selected environment"},
	}

	output.PrintBorderlessTable(data)
}
