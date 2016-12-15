package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	go_path "path"

	"github.com/TIBCOSoftware/flogo-cli/builder"
	"github.com/TIBCOSoftware/flogo-cli/cli"
	"github.com/TIBCOSoftware/flogo-cli/util"
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

	projectDescriptor := loadProjectDescriptor()

	_ = builder.DoGoGet(itemPath)
	// Ignoring error because go get sometimes errors but all goes fine.
	// If the get failed, it will be caught later on in the code

	id := builder.RandId()
	// TODO check if the id already exist and regenerate until unique

	item, err := addItem(projectDescriptor, id, itemPath)
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Error adding item of path '%s': %s\n\n", itemPath, err.Error()))
		os.Exit(2)
	}
	fmt.Fprintf(os.Stdout, "Added Item with id: '%s' and path '%s'\n", item.Id, item.Ref)
	return nil
}

func addItem(projectDescriptor *FlogoProjectDescriptor, id, itemPath string) (*Item, error) {
	dir, err := builder.GetContributionDir(itemPath)
	if err != nil {
		return nil, err
	}
	// TODO do this check better for Trigger and Activity

	if exists(go_path.Join(dir, "service.json")) {
		item := &Item{Id: id, Ref: itemPath, Data: "My_data"}
		services := append(projectDescriptor.Services, item)
		projectDescriptor.Services = services
		fgutil.WriteJSONtoFile(fileDescriptor, projectDescriptor)
		return item, nil
	}

	return nil, errors.New("Item not supported")
}

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
