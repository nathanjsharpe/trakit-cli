package handlers

import (
	"github.com/nathanjsharpe/trakit/config"
	"github.com/nathanjsharpe/trakit/output"
	"strconv"
)

func Session(args []string) {
	session := config.LoadSession()
	data := [][]string{
		[]string{"Environment", session.Environment},
		[]string{"User", session.User.Username},
		[]string{"User Id", strconv.Itoa(session.User.Id)},
		[]string{"Expires", session.Token.Expiration},
	}
	output.PrintSimpleTable(data)
}
