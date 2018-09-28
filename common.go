package commonCLI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

// PrintJSON pretty print JSON string
func PrintJSON(str string) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(str), "", "    ")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	fmt.Println(string(prettyJSON.Bytes()))
	return
}

// OutputJSON displays output of query for alerts in JSON format
func OutputJSON(input interface{}) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintJSON(string(b))
}

// ErrorCheck default error check
func ErrorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// HomeDir set default config location
func HomeDir() string {
	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	return dir
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// inCLI detect if cli runs from akamai framework
func inCLI(appShortName string) string {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	appName := "akamai-" + appShortName
	if inCLI {
		appName = "akamai " + appShortName
	}

	return appName
}

func VerifyArgumentByName(c *cli.Context, argName string) {
	if c.String(argName) == "" {
		log.Fatal(fmt.Sprintf("Please provide required argument(s)! [ %s ]", argName))
	}
}

// SetStringId value
func SetStringId(c *cli.Context, errMessage string) string {
	var id string
	if c.NArg() == 0 {
		log.Fatal(errMessage)
	}

	id = c.Args().Get(0)
	return id
}

// SetIntID value and verify that it is int
func SetIntID(c *cli.Context, errMessage string) string {
	var id string
	if c.NArg() == 0 {
		log.Fatal(errMessage)
	}

	id = c.Args().Get(0)
	isStringInt(id)
	return id
}

func isStringInt(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		errStr := fmt.Sprintf("ID should be integer, you provided: %q\n", id)
		log.Fatal(errStr)
	}
}
