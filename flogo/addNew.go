package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"os"
	//"strings"

	"github.com/TIBCOSoftware/flogo-cli/builder"
	"github.com/TIBCOSoftware/flogo-cli/cli"
	//"github.com/TIBCOSoftware/flogo-cli/util"
	//	"io/ioutil"
	//	"net/http"
	//	"net/url"
	//	"strconv"
)

var optAddN = &cli.OptionInfo{
	Name:      "add_new",
	UsageLine: "add <url>",
	Short:     "Add an activity, flow, model, trigger or palette to a flogo project",
	Long:      `Add an activity, flow, model, trigger or palette to a flogo project`,
}

func init() {
	commandRegistry.RegisterCommand(&cmdAddN{option: optAddN})
}

type cmdAddN struct {
	option *cli.OptionInfo
}

func (c *cmdAddN) OptionInfo() *cli.OptionInfo {
	return c.option
}

func (c *cmdAddN) AddFlags(fs *flag.FlagSet) {
}

func (c *cmdAddN) Exec(args []string) error {

	_ = loadProjectDescriptor()

	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "Error: item path not specified\n\n")
		cmdUsage(c)
	}

	if len(args) > 1 {
		fmt.Fprint(os.Stderr, "Error: Too many arguments given\n\n")
		cmdUsage(c)
	}

	itemPath := args[0]

	if len(itemPath) == 0 {
		fmt.Fprint(os.Stderr, "Error: item path not specified\n\n")
		cmdUsage(c)
	}

	err := builder.DoGoGet(itemPath)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error getting dependencies\n\n")
	}

	return nil
}
