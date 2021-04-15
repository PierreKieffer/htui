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
	// header.TextStyle.Fg = termui.ColorMagenta
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

	///////// DEMO
	// demo_slice := []string{"db-system-api-production", "db-system-api-staging", "load-balancer-production", "load-balancer-staging", "sync-worker-production", "sync-worker-staging"}
	// appList.Rows = demo_slice
	// appList.Rows = append(appList.Rows, utils...)
	///////// DEMO

	return appList
}

func AppOptions(appName string, withoutReturn ...bool) *widgets.List {
	options := widgets.NewList()
	options.Title = appName

	if len(withoutReturn) > 0 {
		options.Rows = []string{"App info", "Dynos formation", "Logs"}
		utils := []string{" ---- Home ---- "}
		options.Rows = append(options.Rows, utils...)
		return options
	}

	options.Rows = []string{"App info", "Dynos formation", "Logs"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)
	return options
}

func AppFormation(appName string) *widgets.List {
	formationList := widgets.NewList()

	formationList.Title = appName

	appFormations, err := core.GetAppFormation(appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, appFormation := range appFormations {
		formationList.Rows = append(formationList.Rows, appFormation.Type)
	}

	utils := []string{"<---- Return", " ---- Home ---- "}

	formationList.Rows = append(formationList.Rows, utils...)

	return formationList
}

func FormationOptions(appName string, formationType string) *widgets.List {
	options := widgets.NewList()
	options.Title = fmt.Sprintf("%v / %v", appName, formationType)
	options.Rows = []string{"Dynos formation info", "Scale"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)
	return options
}

func FormationInfo(selectedFormation string) *widgets.Paragraph {

	selectedFormationSplit := strings.Split(selectedFormation, " / ")

	appName := selectedFormationSplit[0]
	formationType := selectedFormationSplit[1]

	x, y := termui.TerminalDimensions()

	infoScreen := widgets.NewParagraph()

	appFormation, err := core.GetFormationInfo(appName, formationType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonInfo, err := json.MarshalIndent(appFormation, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	infoScreen.Text = string(jsonInfo)

	infoScreen.SetRect(40, 5, x, y)

	return infoScreen

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

	infoScreen := widgets.NewParagraph()

	app, err := core.GetAppInfo(appName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonInfo, err := json.MarshalIndent(app, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	infoScreen.Text = string(jsonInfo)
	infoScreen.SetRect(40, 5, x, y)

	return infoScreen
}

func DynoInfo(selectedDyno string) *widgets.Paragraph {

	selectedDynoSplit := strings.Split(selectedDyno, " / ")

	appName := selectedDynoSplit[0]
	dynoName := selectedDynoSplit[1]

	x, y := termui.TerminalDimensions()

	infoScreen := widgets.NewParagraph()

	dyno, err := core.GetDynoInfo(appName, dynoName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonInfo, err := json.MarshalIndent(dyno, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	infoScreen.Text = string(jsonInfo)
	infoScreen.SetRect(40, 5, x, y)

	return infoScreen
}

func AppLogs(screen *BaseScreen, cache *CacheStorage, signal chan bool) {

	var streamSignal = make(chan bool)

	// x, y := termui.TerminalDimensions()

	go core.StreamLogs(cache.AppName, cache.LogsBuffer, streamSignal)

	for {
		select {
		case <-signal:
			streamSignal <- true

		case log := <-cache.LogsBuffer:

			log = strings.Replace(log, "\n", "", -1)

			// Split log to handle max lenght of log line
			screenMaxSize := screen.UIList.Rectangle.Max.X - 5

			if len(log) >= screenMaxSize {
				for len(log) >= screenMaxSize {
					if log[:1] == " " {
						log = log[1:]
					}
					subLog := log[:screenMaxSize]
					log = log[screenMaxSize:]
					screen.UIList.Rows = append(screen.UIList.Rows, subLog)
				}

				screen.UIList.Rows = append(screen.UIList.Rows, log)

			} else {
				screen.UIList.Rows = append(screen.UIList.Rows, log)

			}

		}
	}

}

func Help() *widgets.Paragraph {
	help := widgets.NewParagraph()
	help.Text = `

 Heroku Terminal User Interface

 Move around : 
 	- go up               ▲  or 'k'
 	- go down             ▼  or 'j'
 	- go to the top       'gg'
 	- go to the bottom    'G'
 	- select item         'enter'
 	- Quit htui           'q'
`

	return help
}

func RestartSelectedDyno(selectedDyno string) (string, *widgets.Paragraph) {

	selectedDynoSplit := strings.Split(selectedDyno, " / ")

	appName := selectedDynoSplit[0]
	dynoName := selectedDynoSplit[1]

	x, y := termui.TerminalDimensions()

	infoScreen := widgets.NewParagraph()

	err := core.RestartDyno(appName, dynoName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dyno, err := core.GetDynoInfo(appName, dynoName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dynoState := dyno.State

	jsonInfo, err := json.MarshalIndent(dyno, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	infoScreen.Text = string(jsonInfo)
	infoScreen.SetRect(40, 5, x, y)

	return dynoState, infoScreen
}

func HelpList() *widgets.List {

	helpList := widgets.NewList()
	helpList.Title = "Help"

	utils := []string{" ---- Home ---- "}
	helpList.Rows = append(helpList.Rows, utils...)

	return helpList
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
