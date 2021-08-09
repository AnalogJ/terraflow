package main

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg/actions/init"
	"github.com/analogj/terraflow/pkg/actions/project"
	"github.com/analogj/terraflow/pkg/version"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var goos string
var goarch string

func main() {

	cli.CommandHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}
USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}
CATEGORY:
   {{.Category}}{{end}}{{if .Description}}
DESCRIPTION:
   {{.Description}}{{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

	app := &cli.App{
		Name:     "terraflow",
		Usage:    "terraform, but with opinionated configuration management",
		Version:  version.VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Jason Kulatunga",
				Email: "jason@thesparktree.com",
			},
		},
		Before: func(c *cli.Context) error {

			terraflow := "github.com/AnalogJ/terraflow"

			var versionInfo string
			if len(goos) > 0 && len(goarch) > 0 {
				versionInfo = fmt.Sprintf("%s.%s-%s", goos, goarch, version.VERSION)
			} else {
				versionInfo = fmt.Sprintf("dev-%s", version.VERSION)
			}

			subtitle := terraflow + utils.LeftPad2Len(versionInfo, " ", 65-len(terraflow))

			color.New(color.FgGreen).Fprintf(c.App.Writer, fmt.Sprintf(utils.StripIndent(
				`
			 ____  ____  ____  ____   __   ____  __     __   _  _ 
			(_  _)(  __)(  _ \(  _ \ / _\ (  __)(  )   /  \ / )( \
			  )(   ) _)  )   / )   //    \ ) _) / (_/\(  O )\ /\ /
			 (__) (____)(__\_)(__\_)\_/\_/(__)  \____/ \__/ (_/\_)
			%s

			`), subtitle))

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-level",
				Usage:   "specify the log level",
				Value:   "INFO",
				EnvVars: []string{"TF_LOG"},
			},
			&cli.StringFlag{
				Name:    "state-bucket-name",
				Usage:   "provide the bucket name where terraform state is stored. Required for all components, excluding 'bootstrap'",
				EnvVars: []string{"_STATE_BUCKET_NAME"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "project",
				Usage: "Create a terraflow project folder structure",
				//UsageText:   "doo - does the dooing",
				Action: func(c *cli.Context) error {
					fmt.Fprintln(c.App.Writer, c.Command.Usage)
					if c.Bool("debug") {
						log.SetLevel(log.DebugLevel)
					} else {
						log.SetLevel(log.InfoLevel)
					}

					return project.Start(c.String("component"), c.String("environment"))
				},

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "component",
						Usage: "specify the component to create",
					},
					&cli.StringFlag{
						Name:  "environment",
						Usage: "specify the environment to create",
					},
				},
			},

			{
				Name:  "init",
				Usage: "Initialize a Terraflow working directory",
				Action: func(c *cli.Context) error {
					fmt.Fprintln(c.App.Writer, c.Command.Usage)
					if c.Bool("debug") {
						log.SetLevel(log.DebugLevel)
					} else {
						log.SetLevel(log.InfoLevel)
					}

					return init.Start(c.String("component"), c.String("environment"))
				},

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "component",
						Usage: "specify the component to use",

						Required: true,
					},
					&cli.StringFlag{
						Name:     "environment",
						Usage:    "specify the environment to use",
						Required: true,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(color.HiRedString("ERROR: %v", err))
	}

}
