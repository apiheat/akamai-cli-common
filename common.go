package commonCLI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
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
func inCLI(appName string) string {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	if inCLI {
		appName = strings.Replace(appName, "-cli-", " ", -1)
	}

	return appName
}

//GetArgumentStr retrieves argument value for string
func GetArgumentStr(c *cli.Context, errMessage string) string {
	id := getArgumentValue(c, errMessage)

	return id
}

//GetArgumentInt retrieves argument value for int
func GetArgumentInt(c *cli.Context, errMessage string) int {
	id := getArgumentValue(c, errMessage)

	ok, val := isStringInt(id)
	if ok != true {
		log.Fatal(errMessage)
	}

	return val
}

//getArgumentValue retrieves argument value
func getArgumentValue(c *cli.Context, errMessage string) string {
	var id string

	if c.NArg() == 0 {
		log.Fatal(errMessage)
	}

	id = c.Args().Get(0)
	return id
}
