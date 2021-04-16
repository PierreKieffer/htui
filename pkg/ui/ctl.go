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

func Home() (*widgets.List, *widgets.Paragraph) {
	options := widgets.NewList()
	options.Title = "Home"
	options.Rows = []string{"Apps", "Addons", "Help"}

	details := widgets.NewParagraph()
	details.Text = `

     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit htui           'q'

`
	return options, details
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

func AppOptions(appName string, withoutReturn ...bool) (*widgets.List, *widgets.Paragraph) {

	details := widgets.NewParagraph()
	details.Text = `

    App info : Get app details. 

    Dynos formations : Get the formations of processes that should be maintained for an app.
                     Update the formations to scale processes or change dyno sizes. 

    Logs : Real time logs tailing for an app.
         Browse through the logs 

`
	options := widgets.NewList()
	options.Title = appName

	if len(withoutReturn) > 0 {
		options.Rows = []string{"App info", "Dynos formations", "Logs"}
		utils := []string{" ---- Home ---- "}
		options.Rows = append(options.Rows, utils...)
		return options, details
	}

	options.Rows = []string{"App info", "Dynos formations", "Logs"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)

	return options, details
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

func FormationOptions(appName string, formationType string) (*widgets.List, *widgets.Paragraph) {
	options := widgets.NewList()
	options.Title = fmt.Sprintf("%v / %v", appName, formationType)
	options.Rows = []string{"Dynos formation info", "Scale dynos formation", "Update dynos size"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)

	details := widgets.NewParagraph()
	details.Text = `

    Formation : The formation of processes that should be maintained for an app.
                Update the formation to scale processes or change dyno sizes. 

    Dynos formation info : Get formation details. 

    Scale dynos formation : Update the number of process to maintain, between 0 to 10. 

    Update dynos size : Update dynos type to support app size : 
        Dyno Type      Memory (RAM)    CPU Share    Compute    Dedicated    Sleeps
        --------------------------------------------------------------------------
        free           512 MB          1x           1x-4x      no           yes
        hobby          512 MB          1x           1x-4x      no           no
        standard-1x    512 MB          1x           1x-4x      no           no
        standard-2x    1024 MB         2x           4x-8x      no           no
        performance-m  2.5 GB          100%%        12x        yes          no
        performance-l  14 GB           100%%        50x        yes          no

`

	return options, details
}

func UpdateFormationQuantity(selectedFormation string, quantity int) *widgets.Paragraph {

	selectedFormationSplit := strings.Split(selectedFormation, " / ")

	appName := selectedFormationSplit[0]
	formationType := selectedFormationSplit[1]

	x, y := termui.TerminalDimensions()

	infoScreen := widgets.NewParagraph()

	formationUpdateResp, err := core.UpdateFormationQuantity(appName, formationType, quantity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	infoScreen.Text = string(formationUpdateResp)

	infoScreen.SetRect(40, 5, x, y)

	return infoScreen

}

func UpdateFormationSize(selectedFormation string, size string) *widgets.Paragraph {

	selectedFormationSplit := strings.Split(selectedFormation, " / ")

	appName := selectedFormationSplit[0]
	formationType := selectedFormationSplit[1]

	x, y := termui.TerminalDimensions()

	infoScreen := widgets.NewParagraph()

	formationUpdateResp, err := core.UpdateFormationSize(appName, formationType, size)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	infoScreen.Text = string(formationUpdateResp)

	infoScreen.SetRect(40, 5, x, y)

	return infoScreen

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

func FormationScalingOptions(selectedFormation string) *widgets.List {
	options := widgets.NewList()
	options.Title = selectedFormation
	options.Rows = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	utils := []string{"<---- Return", " ---- Home ---- "}
	options.Rows = append(options.Rows, utils...)
	return options
}

func DynoSizesOptions(selectedFormation string) *widgets.List {
	options := widgets.NewList()
	options.Title = selectedFormation
	options.Rows = []string{"free", "hobby", "standard-1x", "standard-2x", "performance-m", "performance-l"}
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
            _   _        _ 
           | |_| |_ _  _(_)                 
           | ' \  _| || | |
           |_||_\__|\_,_|_|
    
     Heroku Terminal User Interface

     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit htui           'q'

     -----------------------------
     -    htui authentication    -
     -----------------------------
     htui uses API token mechanism for authentication to Heroku, with HEROKU_API_KEY environment variable. 
     If ~/.netrc file exists (UNIX), HEROKU_API_KEY is set automatically. 
     If ~/.netrc doesn't exist, you need to set HEROKU_API_KEY manually : 
     - Retrieve the API token : 
     	- heroku CLI : heroku auth:token
     	- heroku account setting web page : API Key
     - export HEROKU_API_KEY="api token" 




     -----------------------------
     -          Author           -
     -----------------------------
     https://github.com/PierreKieffer
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
