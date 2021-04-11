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

		ls.SetRect(0, 5, 25, y)
		termui.Render(ls)

	} else {

		ls.SetRect(0, 5, 25, y)
		d.SetRect(25, 5, x, y)

		termui.Render(ls, d)
	}
}

func (screen *BaseScreen) Create() {

	x, y := termui.TerminalDimensions()

	if screen.Header == nil {
		screen.Header = BuildHeader()
	}

	if screen.UIList == nil {
		screen.Screen = "apps"
		screen.UIList = AppList()
	}

	// header
	h := screen.Header
	h.SetRect(0, 0, x, 5)

	// menu list
	ls := screen.UIList
	ls.TextStyle = termui.NewStyle(termui.ColorYellow)
	ls.WrapText = false

	if screen.Display == nil {
		ls.SetRect(0, 5, x, y)

	} else {
		ls.SetRect(0, 5, 25, y)

	}

	termui.Render(h, ls)
}

func (screen *BaseScreen) HandleSelectItem() {

	selectedItem := screen.UIList.Rows[screen.UIList.SelectedRow]

	switch selectedItem {
	case " ---- Home ---- ":
		items := AppList()
		screen.Screen = "apps"
		screen.UIList = items
		screen.Previous = nil
		screen.Display = nil

	case "<---- Return":
		*screen = *screen.Previous

	case "App info":
		screen.Display = AppInfo(screen.UIList.Title)
		screen.Screen = "appInfo"
		screen.Display.Title = "App info"

	case "Dyno info":
		screen.Display = DynoInfo(screen.UIList.Title, selectedItem)
		screen.Screen = "dynoInfo"
		screen.Display.Title = "Dyno info"

	case "Dynos":
		var previousScreen BaseScreen
		previousScreen = *screen

		items := AppDynos(screen.UIList.Title)
		screen.Screen = "dynos"
		screen.UIList = items
		screen.Display = nil
		screen.Previous = &previousScreen

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
			items := DynoOptions(screen.UIList.Title)
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
