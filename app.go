package commonCLI

// CreateNewApp creates new app description
import (
	"time"

	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/urfave/cli/v2"
)

//CreateNewApp returns new application
func CreateNewApp(appShortName, usage, version string) *cli.App {
	
	app := &cli.App{
		Name:     inCLI(appShortName),
		Usage:    usage,
		Version:  version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Rafal Pieniazek",
				Email: "https://mailhide.io/e/XZ3bESkk",
			},
			&cli.Author{
				Name:  "Petr Artamonov",
				Email: "-",
			},
		},
		Copyright: "(c) 2020 Automation Ninjas",
		HelpName:  inCLI(appShortName),
	}

	return app
}

//CreateFlags inits default flags used across all CLIs
func CreateFlags() []cli.Flag {

	appFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "section",
			Value: "default",
			Aliases: []string{
				"s",
			},
			Usage:   "`NAME` of section to use from credentials file",
			EnvVars: []string{string(edgegrid.EnvVarEdgercSection)},
		},
		&cli.StringFlag{
			Name:  "config",
			Value: HomeDir(),
			Aliases: []string{
				"c",
			},
			Usage:   "Location of the credentials `FILE`",
			EnvVars: []string{string(edgegrid.EnvVarEdgercPath)},
		},
		&cli.StringFlag{
			Name:    "debug",
			Value:   "",
			Usage:   "(Deprecated) Debug Level",
			EnvVars: []string{string(edgegrid.EnvVarDebugLevelSection)},
		},
		&cli.StringFlag{
			Name:  "verbosity",
			Value: "",
			Aliases: []string{
				"v",
			},
			Usage:   "Defines level of messages being displayed ( info,warning,error,trace,debug )",
			EnvVars: []string{string(edgegrid.EnvVarDebugLevelSection)},
		},
		&cli.StringFlag{
			Name:  "account-switch-key",
			Value: "",
			Aliases: []string{
				"ask",
			},
			Usage: "Account Switch Key (ASK)",
		},
	}

	return appFlags
}
