package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/otiai10/dockerhubapi"
	"github.com/urfave/cli"
)

// Fetch ...
var Fetch = cli.Command{
	Name:    "fetch",
	Aliases: []string{"f"},
	Usage:   "Fetch resource information",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "repo,r",
			Usage: "Repository name to fetch",
		},
		cli.StringFlag{
			Name:  "repos",
			Usage: "Repositories owner's name to fetch",
		},
		cli.StringFlag{
			Name:  "user,u",
			Usage: "User name to fetch",
		},
	},
	Action: func(ctx *cli.Context) error {

		var resource dockerhubapi.Resource
		if repo := ctx.String("repo"); repo != "" {
			namespace := strings.Split(ctx.String("repo"), "/")
			if len(namespace) != 2 {
				return fmt.Errorf("repo name must be {user}/{name}")
			}
			user, name := namespace[0], namespace[1]
			resource = &dockerhubapi.Repository{User: user, Name: name}
		} else if uname := ctx.String("repos"); uname != "" {
			resource = &dockerhubapi.UserRepositories{User: uname}
		} else if user := ctx.String("user"); user != "" {
			resource = &dockerhubapi.User{Username: user}
		}

		if resource == nil {
			return fmt.Errorf("no resource to fetch is specified")
		}

		api := dockerhubapi.New()
		if err := api.Fetch(resource); err != nil {
			return err
		}

		enc := json.NewEncoder(ctx.App.Writer)
		enc.SetIndent("", "\t")
		return enc.Encode(resource)
	},
}
