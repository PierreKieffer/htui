package ui

import (
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

type BaseScreen struct {
	Screen   string
	Header   *widgets.Paragraph
	UIList   *widgets.List
	Display  *widgets.Paragraph
	Previous *BaseScreen
}

func (screen *BaseScreen) Update() {

	x, y := termui.TerminalDimensions()

	ls := screen.UIList
	d := screen.Display

	if d == nil {

		ls.SetRect(0, 10, x, y)
		termui.Render(ls)

	} else {

		ls.SetRect(0, 10, 40, y)
		d.SetRect(40, 10, x, y)

		termui.Render(ls, d)
	}
}

func (screen *BaseScreen) Create() {

	x, y := termui.TerminalDimensions()

	if screen.Header == nil {
		screen.Header = BuildHeader()
	}

	if screen.UIList == nil {
		screen.Screen = "home"
		screen.UIList = HomeList()
	}

	// header
	h := screen.Header
	h.SetRect(0, 0, x, 10)

	// menu list
	ls := screen.UIList
	ls.TextStyle = termui.NewStyle(termui.ColorYellow)
	ls.WrapText = false

	if screen.Display == nil {
		ls.SetRect(0, 10, x, y)

	} else {
		ls.SetRect(0, 10, 40, y)

	}

	termui.Render(h, ls)
}

func (screen *BaseScreen) HandleSelectItem() {

	selectedItem := screen.UIList.Rows[screen.UIList.SelectedRow]

	switch selectedItem {
	case " ---- Home ---- ":
		/*
			Return to Home page
		*/
		items := HomeList()
		screen.Screen = "home"
		screen.UIList = items
		screen.Previous = nil
		screen.Display = nil

	case "<---- Return":
		/*
			Point to screen.Previous address
		*/
		*screen = *screen.Previous

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

	case "Dynos":
		var previousScreen BaseScreen
		previousScreen = *screen

		items := AppDynos(screen.UIList.Title)
		screen.Screen = "dynos"
		screen.UIList = items
		screen.Display = nil
		screen.Previous = &previousScreen
		// case "Logs":
		// var previousScreen BaseScreen
		// previousScreen = *screen

	/*
		Dynos
	*/
	case "Dyno info":
		screen.Display = DynoInfo(screen.UIList.Title)
		screen.Screen = "dynoInfo"
		screen.Display.Title = "Dyno info"

	default:
		var previousScreen BaseScreen
		previousScreen = *screen

		switch screen.Screen {
		case "apps":
			items := AppOptions(selectedItem)
			screen.Screen = "appOptions"
			screen.UIList = items
			screen.Previous = &previousScreen

		case "dynos":
			items := DynoOptions(screen.UIList.Title, selectedItem)
			screen.Screen = "dynoOptions"
			screen.UIList = items
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
