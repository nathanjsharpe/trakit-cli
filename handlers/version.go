package handlers

import (
	"github.com/nathanjsharpe/trakit/output"
	"github.com/nathanjsharpe/trakit/trakitapi"
)

func Version(args []string, cliVersion string) {
	versions := trakitapi.GetVersions()
	data := [][]string{
		[]string{"Trakit CLI", cliVersion},
		[]string{"Web App", versions.Web},
		[]string{"App Api", versions.App},
		[]string{"Data Api", versions.Data},
		[]string{"Auth Api", versions.Auth},
		[]string{"Event Api", versions.Event},
	}
	output.PrintSimpleTable(data)
}
