package main

import "github.com/urfave/cli/v2"

var (
	CloudCommands = []*cli.Command{
		{
			Name:     "cloud",
			Usage:    "LiveKit Cloud commands",
			Category: "Project Management",
			Subcommands: []*cli.Command{
				{
					Name:  "login",
					Usage: "Login to LiveKit Cloud",
				},
				{
					Name:  "logout",
					Usage: "Logout from LiveKit Cloud",
				},
				{
					Name:  "launch-dashboard",
					Usage: "Launch LiveKit Cloud dashboard for the active project",
					Flags: []cli.Flag{
						projectFlag,
					},
				},
			},
		},
	}
)
