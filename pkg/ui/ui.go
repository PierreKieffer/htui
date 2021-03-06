package ui

import (
	"fmt"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"strconv"
	"strings"
	"time"
)

type BaseScreen struct {
	Screen  string
	Header  *widgets.Paragraph
	UIList  *widgets.List
	Display *widgets.Paragraph

	Previous *BaseScreen
}

type CacheStorage struct {
	AppName    string
	DynoName   string
	DynoState  string
	LogsBuffer chan string
}

var cache *CacheStorage

var signal = make(chan bool)

func (screen *BaseScreen) Update() {

	x, y := termui.TerminalDimensions()

	ls := screen.UIList
	d := screen.Display

	if d == nil {

		ls.SetRect(0, 10, x, y)
		termui.Render(ls)

		switch screen.Screen {
		case "appLogs":

			go AppLogs(screen, cache, signal)

			go func() {
				for screen.Screen == "appLogs" {
					termui.Render(ls)
					time.Sleep(1 * time.Second)
				}
			}()
		}

	} else {

		switch screen.Screen {
		default:
			d.TitleStyle.Fg = termui.ColorYellow
			ls.TitleStyle.Fg = termui.ColorYellow
			ls.SetRect(0, 10, 40, y)
			d.SetRect(40, 10, x, y)

			termui.Render(ls, d)
		}
	}
}

func (screen *BaseScreen) Create() {

	x, y := termui.TerminalDimensions()

	if screen.Header == nil {
		screen.Header = BuildHeader()
	}

	if screen.UIList == nil {
		screen.Screen = "home"
		items, details := Home()
		screen.UIList = items
		screen.Display = details
	}

	// header
	h := screen.Header
	h.SetRect(0, 0, x, 10)

	// menu list
	ls := screen.UIList
	ls.SelectedRowStyle = termui.NewStyle(termui.ColorMagenta)
	ls.TitleStyle.Fg = termui.ColorYellow
	ls.WrapText = false

	if screen.Display == nil {
		ls.SetRect(0, 10, x, y)
		termui.Render(h, ls)

	} else {
		d := screen.Display
		d.TitleStyle.Fg = termui.ColorYellow

		ls.SetRect(0, 10, 40, y)
		d.SetRect(40, 10, x, y)
		termui.Render(h, ls, d)

	}

}

func (screen *BaseScreen) HandleSelectItem() {

	var selectedItem string

	if screen.Screen == "appLogs" {

		signal <- true

		items, details := AppOptions(strings.Split(screen.UIList.Title, " - ")[0])
		screen.Screen = "appOptions"
		screen.UIList = items
		screen.Display = details

		previousItems := AppList()
		screen.Previous.Screen = "apps"
		screen.Previous.UIList = previousItems
		screen.Previous.Display = nil

		screen.Update()
		return
	}

	selectedItem = screen.UIList.Rows[screen.UIList.SelectedRow]

	switch selectedItem {
	case " ---- Home ---- ":
		/*
			Return to Home page
		*/

		items, details := Home()
		screen.Screen = "home"
		screen.UIList = items
		screen.Display = details
		screen.Previous = nil

	case "<---- Return":
		/*
			Point to screen.Previous address
		*/

		*screen = *screen.Previous

	case "Help":
		/*
			Go to Help page
		*/
		var previousScreen BaseScreen
		previousScreen = *screen

		items := HelpList()
		screen.Screen = "help"
		screen.UIList = items
		screen.Display = Help()
		screen.Previous = &previousScreen

	/*
		Apps
	*/
	case "Apps":
		var previousScreen BaseScreen
		previousScreen = *screen

		items := AppList()
		screen.Screen = "apps"
		screen.UIList = items
		screen.Previous = &previousScreen
		screen.Display = nil

	case "App info":
		var previousScreen BaseScreen
		previousScreen = *screen

		screen.Display = AppInfo(screen.UIList.Title)
		screen.Screen = "appInfo"
		screen.Display.Title = "App info"
		screen.Previous = &previousScreen

	case "Logs":
		var previousScreen BaseScreen
		previousScreen = *screen

		var logsBuffer = make(chan string)

		cache.LogsBuffer = logsBuffer
		cache.AppName = screen.UIList.Title

		screen.Screen = "appLogs"

		screen.UIList.Rows = []string{}
		screen.UIList.Title = fmt.Sprintf("%v - %v", cache.AppName, "Logs | Return : 'enter' | Top : 'gg' | Bottom 'G'")
		screen.Display = nil

		screen.Previous = &previousScreen

	/*
		Dynos
	*/
	case "Dynos formations":
		var previousScreen BaseScreen
		previousScreen = *screen

		items := AppFormation(screen.UIList.Title)
		screen.Screen = "formation"
		screen.UIList = items
		screen.Display = nil
		screen.Previous = &previousScreen

	case "Dynos formation info":
		screen.Display = FormationInfo(screen.UIList.Title)
		screen.Screen = "formationInfo"
		screen.Display.Title = "Dynos formation info"

	case "Scale dynos formation":
		var previousScreen BaseScreen
		previousScreen = *screen

		screen.Screen = "scaleFormation"
		screen.Display = nil
		items := FormationScalingOptions(screen.UIList.Title)
		screen.UIList = items
		screen.Previous = &previousScreen

	case "Update dynos size":
		var previousScreen BaseScreen
		previousScreen = *screen

		screen.Screen = "updateDynosType"
		screen.Display = nil
		items := DynoSizesOptions(screen.UIList.Title)
		screen.UIList = items
		screen.Previous = &previousScreen

	case "Restart":
		dynoState, display := RestartSelectedDyno(screen.UIList.Title)
		screen.Display = display
		screen.Screen = "restartDyno"
		screen.Display.Title = fmt.Sprintf("Restart dyno %v", screen.UIList.Title)

		selectedDynoSplit := strings.Split(screen.UIList.Title, " / ")
		cache.AppName = selectedDynoSplit[0]
		cache.DynoName = selectedDynoSplit[1]
		cache.DynoState = dynoState

	/*
		Addons
	*/
	case "Addons":
		var previousScreen BaseScreen
		previousScreen = *screen

		items := AddonList()
		screen.Screen = "addons"
		screen.UIList = items
		screen.Previous = &previousScreen
		screen.Display = nil

	case "Addon info":
		var previousScreen BaseScreen
		previousScreen = *screen

		screen.Display = AddonInfo(screen.UIList.Title)
		screen.Screen = "addonInfo"
		screen.Display.Title = "Addon info"
		screen.Previous = &previousScreen

	default:
		var previousScreen BaseScreen
		previousScreen = *screen

		switch screen.Screen {
		case "apps":
			items, details := AppOptions(selectedItem)
			screen.Screen = "appOptions"
			screen.UIList = items
			screen.Display = details
			screen.Previous = &previousScreen

		case "formation":
			items, details := FormationOptions(screen.UIList.Title, selectedItem)
			screen.Screen = "formationOptions"
			screen.UIList = items
			screen.Display = details
			screen.Previous = &previousScreen

		case "scaleFormation":
			quantity, _ := strconv.Atoi(selectedItem)
			screen.Display = UpdateFormationQuantity(screen.UIList.Title, quantity)
			screen.Display.Title = fmt.Sprintf("%v : Scale to %v", screen.UIList.Title, selectedItem)
			screen.UIList.Rows = []string{"<---- Return", " ---- Home ---- "}
			screen.UIList.ScrollTop()

			screen.Previous.Display = nil

		case "updateDynosType":
			screen.Display = UpdateFormationSize(screen.UIList.Title, selectedItem)
			screen.Display.Title = fmt.Sprintf("%v : Update size to %v", screen.UIList.Title, selectedItem)
			screen.UIList.Rows = []string{"<---- Return", " ---- Home ---- "}
			screen.UIList.ScrollTop()

			screen.Previous.Display = nil

		case "addons":
			items, _ := AddonOptions(selectedItem)
			screen.Screen = "addonOptions"
			screen.UIList = items
			screen.Display = nil
			screen.Previous = &previousScreen
		}

	}

	screen.Update()
}

var baseScreen BaseScreen

func App() {

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	if cache == nil {
		cache = &CacheStorage{}
	}

	baseScreen.Create()

	previousKey := ""

	uiEvents := termui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			baseScreen.UIList.ScrollDown()
		case "k", "<Up>":
			baseScreen.UIList.ScrollUp()
		case "<C-d>":
			baseScreen.UIList.ScrollHalfPageDown()
		case "<C-u>":
			baseScreen.UIList.ScrollHalfPageUp()
		case "<C-f>":
			baseScreen.UIList.ScrollPageDown()
		case "<C-b>":
			baseScreen.UIList.ScrollPageUp()
		case "<Enter>":
			baseScreen.HandleSelectItem()
		case "g":
			if previousKey == "g" {
				baseScreen.UIList.ScrollTop()
			}
		case "G", "<End>":
			baseScreen.UIList.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		baseScreen.Create()

	}
}
