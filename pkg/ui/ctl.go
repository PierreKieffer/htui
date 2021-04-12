package ui

import (
	"encoding/json"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/core"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"os"
	"strings"
)

func BuildHeader() *widgets.Paragraph {
	header := widgets.NewParagraph()
	header.Text = `
  _   _        _ 
 | |_| |_ _  _(_)
 | ' \  _| || | |
 |_||_\__|\_,_|_|

 Heroku Terminal User Interface
`

	return header
}

func HomeList() *widgets.List {
	options := widgets.NewList()
	options.Title = "Home"
	options.Rows = []string{"Apps", "Addons", "Help"}
	return options
}

func AppList() *widgets.List {

	appList := widgets.NewList()
	appList.Title = "Apps"

	apps, err := core.GetApps()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, app := range apps {
		appList.Rows = append(appList.Rows, app.Name)
	}

	utils := []string{"<---- Return", " ---- Home ---- "}
	appList.Rows = append(appList.Rows, utils...)

	return appList
}

func AppOptions(appName string) *widgets.List {
	options := widgets.NewList()
	options.Title = appName
	options.Rows = []string{"App info", "Dynos", "Logs"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)
	return options
}

func DynoOptions(appName string, dynoName string) *widgets.List {
	options := widgets.NewList()
	options.Title = fmt.Sprintf("%v / %v", appName, dynoName)
	options.Rows = []string{"Dyno info", "Restart", "Stop"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)
	return options
}

func AppDynos(appName string) *widgets.List {
	dynosList := widgets.NewList()
	dynosList.Title = appName

	dynos, err := core.GetAppDynos(appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, dyno := range dynos {
		dynosList.Rows = append(dynosList.Rows, dyno.Name)
	}

	utils := []string{"<---- Return", " ---- Home ---- "}

	dynosList.Rows = append(dynosList.Rows, utils...)

	return dynosList
}

func AppInfo(appName string) *widgets.Paragraph {

	x, y := termui.TerminalDimensions()

	var info string
	infoScreen := widgets.NewParagraph()

	apps, err := core.GetApps()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, app := range apps {
		if app.Name == appName {
			jsonInfo, err := json.MarshalIndent(app, "", "    ")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			info = string(jsonInfo)
			break
		}
	}

	infoScreen.Text = info
	infoScreen.SetRect(40, 5, x, y)

	return infoScreen
}

func DynoInfo(selectedDyno string) *widgets.Paragraph {

	selectedDynoSplit := strings.Split(selectedDyno, " / ")

	appName := selectedDynoSplit[0]
	dynoName := selectedDynoSplit[1]

	x, y := termui.TerminalDimensions()

	var info string
	infoScreen := widgets.NewParagraph()

	dynos, err := core.GetAppDynos(appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, dyno := range dynos {
		if dyno.Name == dynoName {
			jsonInfo, err := json.MarshalIndent(dyno, "", "    ")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			info = string(jsonInfo)
			break
		}
	}

	infoScreen.Text = info
	infoScreen.SetRect(40, 5, x, y)

	return infoScreen
}

func AppNodes() []*widgets.TreeNode {

	apps, err := core.GetApps()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var appNodes []*widgets.TreeNode

	for _, app := range apps {
		var appNode widgets.TreeNode
		appNode.Value = nodeValue(app.Name)
		appNode.Nodes = BuildAppNodes(app)

		appNodes = append(appNodes, &appNode)
	}

	return appNodes
}

func BuildAppNodes(app core.App) []*widgets.TreeNode {

	var nodes []*widgets.TreeNode

	var infoNode widgets.TreeNode
	infoNode.Value = nodeValue("info")
	infoNode.Nodes = nil

	nodes = append(nodes, &infoNode)

	var dynosNode widgets.TreeNode
	dynosNode.Value = nodeValue("dynos")
	dynosNode.Nodes = nil

	nodes = append(nodes, &dynosNode)

	var logsNode widgets.TreeNode
	logsNode.Value = nodeValue("logs")
	logsNode.Nodes = nil

	nodes = append(nodes, &logsNode)

	return nodes
}

type nodeValue string

func (value nodeValue) String() string {
	return string(value)
}
