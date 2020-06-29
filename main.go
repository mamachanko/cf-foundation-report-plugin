package main

import (
	"encoding/json"
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
)

type PluginDemonstratingParams struct {
	uppercase *bool
}

func main() {
	plugin.Start(new(PluginDemonstratingParams))
}

type Report struct {
	Apps   []App   `json:"apps"`
	Orgs   []Org   `json:"orgs"`
	Spaces []Space `json:"spaces"`
}

type App struct {
	Name      string `json:"name"`
	Instances int    `json:"instances"`
}

type Org struct {
	Name string `json:"name"`
}

type Space struct {
	Name string `json:"name"`
}

func (pluginDemo *PluginDemonstratingParams) Run(cliConnection plugin.CliConnection, args []string) {
	report := Report{
		Apps:   GetApps(cliConnection),
		Orgs:   GetOrgs(cliConnection),
		Spaces: GetSpaces(cliConnection),
	}

	json, err := json.Marshal(report)
	if err != nil {
		fmt.Println("failed to serialize report")
	}

	fmt.Println(string(json))
}

func GetSpaces(cfCli plugin.CliConnection) []Space {
	cfSpaces, err := cfCli.GetSpaces()
	if err != nil {
		fmt.Println("failed to get spaces")
	}
	var spaces []Space

	for _, cfSpace := range cfSpaces {
		space := Space{Name: cfSpace.Name}
		spaces = append(spaces, space)
	}

	return spaces
}

func GetOrgs(cfCli plugin.CliConnection) []Org {
	cfOrgs, err := cfCli.GetOrgs()
	if err != nil {
		fmt.Println("failed to get orgs")
	}
	var orgs []Org

	for _, cfOrg := range cfOrgs {
		org := Org{Name: cfOrg.Name}
		orgs = append(orgs, org)
	}

	return orgs
}

func GetApps(cfCli plugin.CliConnection) []App {
	cfApps, err := cfCli.GetApps()
	if err != nil {
		fmt.Println("failed to get apps")
	}
	var apps []App

	for _, cfApp := range cfApps {
		app := App{Name: cfApp.Name, Instances: cfApp.TotalInstances}
		apps = append(apps, app)
	}

	return apps
}

func (pluginDemo *PluginDemonstratingParams) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "FoundationReport",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "foundation-report",
				HelpText: "provides foundation metadata",
			},
		},
	}
}
