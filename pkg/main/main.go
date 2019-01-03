package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juxtin/gore/pkg/debug"
	"github.com/juxtin/gore/pkg/files"
	"github.com/juxtin/gore/pkg/gore"
	cli "gopkg.in/urfave/cli.v1"
)

func die(msg string) {
	fmt.Fprintf(os.Stderr, msg+"\n")
	os.Exit(1)
}

func graph(c *cli.Context) error {
	rootDirArg := c.String("dir")
	gopath := c.String("gopath")
	if len(gopath) == 0 {
		die("You must either set $GOPATH in your shell or use the --gopath flag")
	}
	debug.Print("Gopath:", gopath)
	output := c.String("output")
	debug.Print("Root Dir:", rootDirArg)
	debug.Print("Output:", output)
	rootDir, err := filepath.Abs(rootDirArg)
	if err != nil {
		die("Unable to find absolute path to '" + rootDirArg + "': " + err.Error())
	}
	graph := gore.BuildGraph(gopath, rootDir)
	if len(output) == 0 {
		// defaulting to stdout
		fmt.Println(graph)
	} else {
		debug.Print("writing to: " + output)
		if err := files.WriteFile(output, graph); err != nil {
			die("Error writing '" + output + "': " + err.Error())
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gore"
	app.Usage = "Visualize a Golang project's internal import graph"
	app.UsageText = "gore [global options] [root directory]"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "gopath, g",
			Usage:  "GOPATH to use while scanning for files",
			EnvVar: "GOPATH",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Usage: "directory under which the .go source files are located",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "destination filepath for the output graph (default: stdout)",
		},
	}
	app.Action = graph

	err := app.Run(os.Args)
	if err != nil {
		die(err.Error())
	}
}
