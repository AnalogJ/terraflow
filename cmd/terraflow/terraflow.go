package main

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	applyAction "github.com/analogj/terraflow/pkg/actions/apply"
	cleanAction "github.com/analogj/terraflow/pkg/actions/clean"
	initAction "github.com/analogj/terraflow/pkg/actions/init"
	planAction "github.com/analogj/terraflow/pkg/actions/plan"
	projectAction "github.com/analogj/terraflow/pkg/actions/project"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/analogj/terraflow/pkg/version"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
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
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}
					appLogger := logrus.WithFields(logrus.Fields{
						"type": "project",
					})

					appConfig := config.New()
					appConfig.Set("component", c.String("component"))
					appConfig.Set("environment", c.String("environment"))

					return projectAction.Start(appLogger, appConfig)
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
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}
					appLogger := logrus.WithFields(logrus.Fields{
						"type": "init",
					})

					appConfig := config.New()
					appConfig.Set("component", c.String("component"))
					appConfig.Set("environment", c.String("environment"))

					return initAction.Start(appLogger, appConfig)
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
			{
				Name:  "clean",
				Usage: "Clean a Terraflow working directory",
				Action: func(c *cli.Context) error {
					fmt.Fprintln(c.App.Writer, c.Command.Usage)

					if c.Bool("debug") {
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}
					appLogger := logrus.WithFields(logrus.Fields{
						"type": "init",
					})

					appConfig := config.New()

					return cleanAction.Start(appLogger, appConfig)
				},
			},
			{
				Name:  "plan",
				Usage: "Terraform Plan",
				Action: func(c *cli.Context) error {
					fmt.Fprintln(c.App.Writer, c.Command.Usage)

					if c.Bool("debug") {
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}
					appLogger := logrus.WithFields(logrus.Fields{
						"type": "plan",
					})

					appConfig := config.New()
					appConfig.Set("component", c.String("component"))
					appConfig.Set("environment", c.String("environment"))
					appConfig.Set("target", c.String("target"))
					appConfig.Set("var", c.StringSlice("var"))

					return planAction.Start(appLogger, appConfig)
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
					&cli.StringFlag{
						Name:     "target",
						Usage:    "specify the component to target",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:     "var",
						Usage:    "key=value pairs to pass to terraform",
						Required: true,
					},
				},
			},
			{
				Name:  "apply",
				Usage: "Terraform Apply",
				Action: func(c *cli.Context) error {
					fmt.Fprintln(c.App.Writer, c.Command.Usage)

					if c.Bool("debug") {
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}
					appLogger := logrus.WithFields(logrus.Fields{
						"type": "apply",
					})

					appConfig := config.New()
					appConfig.Set("component", c.String("component"))
					appConfig.Set("environment", c.String("environment"))
					appConfig.Set("target", c.String("target"))
					appConfig.Set("var", c.StringSlice("var"))

					return applyAction.Start(appLogger, appConfig)
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
					&cli.StringFlag{
						Name:     "target",
						Usage:    "specify the component to target",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:     "var",
						Usage:    "key=value pairs to pass to terraform",
						Required: true,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(color.HiRedString("ERROR: %v", err))
	}

}
