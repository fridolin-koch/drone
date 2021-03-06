package main

import (
	"os"
	"text/template"

	"github.com/urfave/cli"
)

var userListCmd = cli.Command{
	Name:   "ls",
	Usage:  "list all users",
	Action: userList,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: tmplUserList,
		},
	},
}

func userList(c *cli.Context) error {
	client, err := newClient(c)
	if err != nil {
		return err
	}

	users, err := client.UserList()
	if err != nil || len(users) == 0 {
		return err
	}

	tmpl, err := template.New("_").Parse(c.String("format") + "\n")
	if err != nil {
		return err
	}
	for _, user := range users {
		tmpl.Execute(os.Stdout, user)
	}
	return nil
}

// template for user list items
var tmplUserList = `{{ .Login }}`
