package handlers

import (
	"fmt"
	"github.com/nathanjsharpe/trakit/output"
	"github.com/nathanjsharpe/trakit/trakitapi"
)

var landmarkHeaders = []string{
	"Id",
	"Name",
	"Address",
	"Description",
	"Icon",
	"ShowOnMap",
	"Lat",
	"Lng",
	"Status",
}

func Landmark(args []string) {
	if len(args) == 0 {
		landmarkHelp()
	} else {
		first, args := args[0], args[1:]
		switch first {
		case "get", "g":
			handleLandmarkGet(args)
		}
	}
}

func handleLandmarkGet(args []string) {
	if len(args) == 0 {
		fmt.Println("Get one or all landmarks")
		fmt.Println("Usage: trakit landmark get <'all' or id>")
	} else {
		first, args := args[0], args[1:]
		switch first {
		case "all", "a":
			showAllLandmarks(args)
		default:
			fmt.Println("fetching landmark with id", first)
		}
	}
}

func showAllLandmarks(args []string) {
	landmarks, err := trakitapi.GetAllLandmarks()
	if err != nil {
		output.PrintErrorAndExit(err.Error())
	}
	rows := make([][]string, len(landmarks))
	for i := range landmarks {
		rows[i] = landmarks[i].ToRow()
	}
	output.PrintResourceTable(rows, landmarkHeaders)
}

func landmarkHelp() {
	data := [][]string{
		[]string{"trakit landmark get", "Get one or all landmarks (alias: g)"},
		[]string{"trakit landmark create", "Create a landmark (alias: c)"},
		[]string{"trakit landmark delete", "Delete a landmark (alias: d)"},
		[]string{"trakit landmark update", "Update a landmark (alias: u)"},
	}

	fmt.Println("Perform actions on landmark resource")
	fmt.Println("\nUsage:")

	output.PrintBorderlessTable(data)
	fmt.Println("\nAliases: landmark, landmarks, lm")
}
