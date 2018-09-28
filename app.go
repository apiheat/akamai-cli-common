package commonCLI

// CreateNewApp creates new app description
import (
	edgegrid "github.com/apiheat/go-edgegrid"
	"github.com/urfave/cli"
)

func CreateNewApp(appShortName, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = inCLI(appShortName)
	app.HelpName = inCLI(appShortName)
	app.Usage = usage
	app.Version = version
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	return app
}

func CreateFlags() []cli.Flag {
	appFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "section, s",
			Value:  "default",
			Usage:  "`NAME` of section to use from credentials file",
			EnvVar: string(edgegrid.EnvVarEdgercSection),
		},
		cli.StringFlag{
			Name:   "config, c",
			Value:  HomeDir(),
			Usage:  "Location of the credentials `FILE`",
			EnvVar: string(edgegrid.EnvVarEdgercPath),
		},
		cli.StringFlag{
			Name:   "debug",
			Value:  "",
			Usage:  "Debug Level",
			EnvVar: string(edgegrid.EnvVarDebugLevelSection),
		},
	}

	return appFlags
}
