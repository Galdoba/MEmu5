package main

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/ConGo/congo"
)

//SPACES -
const (
	SPACES = " \x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0A\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1A\x1b\x1c\x1d\x1e\x1f"
)

//DeviseMap -
var (
	width            int
	height           int
	activeBorderName string
	canClose         bool
	info             interface{}
	key              string
	id               int
	windowList       []congo.IWindow
	//Window           map[string]congo.IWindow

	gridList  []IGrid
	hostList  []IObj
	DeviseMap interface{}
	//CFDBMap   interface{}
	//MActionsMap      interface{}
	SourceIcon  IObj
	TargetIcon  IObj
	TargetIcon2 IObj
	player      IPersona
	Matrix      *THost
	STime       string
	SrTime      time.Time
	TimeMarker  string
	command     string
	GREEN       congo.TColor
)

//ObjByNames -
/*var Window = map[string]congo.IWindow{
//"null": nil,
}*/

//ObjByNames -
var ObjByNames = map[string]IObj{
//"null": nil,
}

//CFDBMap - Complex Forms Storing Map (global)
var CFDBMap = map[int]ComplexForm{
//null
}

//WMap - Complex Forms Storing Map (global)
var WMap = map[string]congo.IWindow{
//null
}

//SPowerMap - Active Sprite Powers Storing Map (global)
var SPowerMap = map[int]SpritePower{
//null
}

func init() {
	//w1 := congo.NewWindow(width - (width*3/10),0,(width*3/10),height, "Log", "Block")
}

func hold() {
	dur := time.Second / 4
	time.Sleep(dur)
	draw()
}

func initialize() {
	//winMap := make([]interface{}, 1)
	setSeed()
	congo.InitBorders()
	width, height = congo.GetSize()
	congo.SetTBorder("Default")
	activeBorderName = congo.GetTBorderName()
	congo.InitWindowsMap()
	InitDeviceDatabase()
	DeviseMap := AddDevice()
	if DeviseMap.DeviceDB == nil {
		panic("INITIATE ERROR: Device map must not be NIL...")
	}

	InitMatrixActionMap()
	//MActions.MActionMap["xxx"] = 1354
	/*MActions.MActionMap =*/ //AddMAction()
	/*if MActions.MActionMap == nil {
		panic("INITIATE ERROR: Matrix Action map must not be NIL...")
	}*/
	//deviceDB := make(map[string]*TDevice)
	//deviceDB["Camera"].SetRunningMode(false)

	//objectList = append(objectList, NewFile("hh"))

	congo.NewKeyboardAction("Choose_next_window", "<TAB>", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		activeWindow := 1
		for i := range windowList {
			if WMap[windowList[i]].InFocus() == true {
				activeWindow = i
			}
			WMap[windowList[i]].SetFocus(false)
		}
		if activeWindow == len(windowList)-1 {
			activeWindow = 0
		} else {
			activeWindow++
		}

		WMap[windowList[activeWindow]].SetFocus(true)
		return true
	})

	congo.NewKeyboardAction("Exit_Programm", "<esc>", "", func(ev *congo.KeyboardEvent) bool {
		canClose = true
		return true
	})

	congo.NewKeyboardAction("Add Line", "<space>", "", func(ev *congo.KeyboardEvent) bool {
		WMap["User Input"].WPrint(" ")
		return true
	})

	congo.NewKeyboardAction("Input", "<ENTER>", "", func(ev *congo.KeyboardEvent) bool {
		input := WMap["User Input"].WRead()
		UserInput(player.GetName() + ">" + input)
		WMap["User Input"].WClear()
		draw()
		return true
	})

	congo.NewKeyboardAction("Delete_input_char", "<BACKSPACE>", "", func(ev *congo.KeyboardEvent) bool {
		content := WMap["User Input"].WGetContent()
		slContent := congo.SplitSubN(content, 1)
		if len(slContent) >= 1 {
			slContent = slContent[:len(slContent)-1]
		}
		content = strings.Join(slContent, "")
		WMap["User Input"].WClear()
		WMap["User Input"].WSetContent(content)
		draw()
		return true
	})

	congo.NewKeyboardAction("Expand Help Window", "<F1>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		for _, w := range WMap {
			w.SetFocus(false)
		}
		WMap["Help"].SetSize(width/2, height-height/5)
		//SET POSITION
		WMap["Help"].SetPosition(width*20/100, 1)
		WMap["Help"].SetBorderVisibility(true)
		WMap["Help"].SetFocus(true)
		WMap["Enviroment"].SetPosition(width+1, 1)
		WMap["Programs"].SetPosition(width+1, 1)
		WMap["System"].SetPosition(width+1, 1)
		WMap["Actions"].SetPosition(width+1, 1)
		WMap["Top"].WClear()
		multiPrintTo("Top", "███ {DEFAULT}[F1]HELP{DEFAULT} ███ {GREEN}[F2]ENVIROMENT{DEFAULT} ███ {GREEN}[F3]PROGRAMS{DEFAULT} ███ {GREEN}[F4]System{DEFAULT} ███ {GREEN}[F5]ACTIONS{DEFAULT} ████████████████████")
		refreshHelpWin()
		draw()
		return true
	})

	congo.NewKeyboardAction("Expand Enviroment Window", "<F2>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		WMap["Enviroment"].SetSize(width/2, height-height/5)
		//SET POSITION
		WMap["Enviroment"].SetPosition(width*20/100, 1)
		WMap["Help"].SetPosition(width+1, 1)
		WMap["Programs"].SetPosition(width+1, 1)
		WMap["System"].SetPosition(width+1, 1)
		WMap["Actions"].SetPosition(width+1, 1)
		for _, w := range WMap {
			w.SetFocus(false)
		}
		WMap["Enviroment"].SetFocus(true)
		WMap["Top"].WClear()
		multiPrintTo("Top", "███ {GREEN}[F1]HELP{DEFAULT} ███ {DEFAULT}[F2]ENVIROMENT{DEFAULT} ███ {GREEN}[F3]PROGRAMS{DEFAULT} ███ {GREEN}[F4]System{DEFAULT} ███ {GREEN}[F5]ACTIONS{DEFAULT} ████████████████████")
		draw()
		return true
	})

	congo.NewKeyboardAction("Expand Programs Window", "<F3>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		WMap["Programs"].SetSize(width/2, height-height/5)
		//SET POSITION
		WMap["Programs"].SetPosition(width*20/100, 1)
		WMap["Help"].SetPosition(width+1, 1)
		WMap["Enviroment"].SetPosition(width+1, 1)
		WMap["System"].SetPosition(width+1, 1)
		WMap["Actions"].SetPosition(width+1, 1)
		for _, w := range WMap {
			w.SetFocus(false)
		}
		WMap["Programs"].SetFocus(true)
		WMap["Top"].WClear()
		multiPrintTo("Top", "███ {GREEN}[F1]HELP{DEFAULT} ███ {GREEN}[F2]ENVIROMENT{DEFAULT} ███ {DEFAULT}[F3]PROGRAMS{DEFAULT} ███ {GREEN}[F4]System{DEFAULT} ███ {GREEN}[F5]ACTIONS{DEFAULT} ████████████████████")
		refreshProgramsWin()
		draw()
		return true
	})

	congo.NewKeyboardAction("Expand System Window", "<F4>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		WMap["System"].SetSize(width/2, height-height/5)
		//SET POSITION
		WMap["System"].SetPosition(width*20/100, 1)
		WMap["Help"].SetPosition(width+1, 1)
		WMap["Enviroment"].SetPosition(width+1, 1)
		WMap["Programs"].SetPosition(width+1, 1)
		WMap["Actions"].SetPosition(width+1, 1)
		for _, w := range WMap {
			w.SetFocus(false)
		}
		WMap["System"].SetFocus(true)
		WMap["Top"].WClear()
		multiPrintTo("Top", "███ {GREEN}[F1]HELP{DEFAULT} ███ {GREEN}[F2]ENVIROMENT{DEFAULT} ███ {GREEN}[F3]PROGRAMS{DEFAULT} ███ {DEFAULT}[F4]System{DEFAULT} ███ {GREEN}[F5]ACTIONS{DEFAULT} ████████████████████")
		refreshSystemWin()
		draw()
		return true
	})

	congo.NewKeyboardAction("Expand Actions Window", "<F5>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		WMap["Actions"].SetSize(width/2, height-height/5)
		//SET POSITION
		WMap["Actions"].SetPosition(width*20/100, 1)
		WMap["Help"].SetPosition(width+1, 1)
		WMap["Enviroment"].SetPosition(width+1, 1)
		WMap["Programs"].SetPosition(width+1, 1)
		WMap["System"].SetPosition(width+1, 1)
		for _, w := range WMap {
			w.SetFocus(false)
		}
		WMap["Actions"].SetFocus(true)
		WMap["Top"].WClear()
		multiPrintTo("Top", "███ {GREEN}[F1]HELP{DEFAULT} ███ {GREEN}[F2]ENVIROMENT{DEFAULT} ███ {GREEN}[F3]PROGRAMS{DEFAULT} ███ {GREEN}[F4]System{DEFAULT} ███ {DEFAULT}[F5]ACTIONS{DEFAULT} ████████████████████")
		refreshActionsWin()
		draw()
		return true
	})

	congo.NewKeyboardAction("Defauil View", "<F12>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)

		WMap["Persona"].SetSize(width*20/100, height-height/5+1)
		WMap["Grid"].SetSize(width*20/100, height*2/10)
		WMap["Enviroment"].SetSize(width/2, height-height/5)
		//SET POSITION

		WMap["Persona"].SetPosition(0, 0)
		WMap["Grid"].SetPosition(0, height-height*2/10)
		WMap["Enviroment"].SetPosition(width*20/100, 1)
		WMap["Process"].SetPosition(width*20/100, height*8/10+1)
		WMap["Log"].SetFocus(true)
		//NewFile("File ")
		//NewDevice("Camera ", 2)
		//addDevice("Camera3")
		//addDevice("Camera5")
		//addDevice("Erika MCD-1")
		//NewDevice("Erika MCD-1", 4)
		//WMap["User Input"].WPrint("matrix search>host", congo.ColorGreen)
		draw()
		return true
	})

	congo.NewKeyboardAction("Move_selector_down", "<down>", "", func(ev *congo.KeyboardEvent) bool { //KeyboardEvent
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if WMap[windowList[i]].InFocus() == true {
				index := WMap[windowList[i]].GetScrollIndex()
				/*WMap[windowList[i]].SetScrollIndex(utils.Min(index+1, WMap[windowList[i]].GetStoredRows()-WMap[windowList[i]].GetPrintableHeight()+2))
				if WMap[windowList[i]].GetStoredRows() < WMap[windowList[i]].GetPrintableHeight() {
					WMap[windowList[i]].SetScrollIndex(0)

				}*/
				WMap[windowList[i]].SetScrollIndex(index - 1)
				WMap[windowList[i]].WDraw()

			}
		}
		//curs++
		//curs = utils.Min(curs, len(styles)-1)
		return true
	})

	congo.NewKeyboardAction("Move_selector_up", "<up>", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if WMap[windowList[i]].InFocus() == true {
				//WMap[windowList[i]].SetAutoScroll(false)
				index := WMap[windowList[i]].GetScrollIndex()
				//WMap[windowList[i]].SetScrollIndex(utils.Max(index-1, 0))
				WMap[windowList[i]].SetAutoScroll(false)
				WMap[windowList[i]].SetScrollIndex(index + 1)
				WMap[windowList[i]].WDraw()
			}
		}
		return true
	})

	congo.NewResizeAction("Resize_Window", "<Resize>", "", func(ev *congo.ResizeEvent) bool { //KeyboardEvent
		congo.Flush()
		width, height = congo.GetSize()
		if width <= 101 {
			width = 101
		}
		if height <= 11 {
			height = 11
		}
		congo.ClearScreen(' ', congo.GetFgColor(), congo.GetBgColor())
		//w5 := congo.NewWindow(width*20/100 +1 ,0,width/2,height, "Enviroment", "Block")
		//SET SIZE
		WMap["Log"].SetSize((width * 3 / 10), height-3)
		WMap["Persona"].SetSize(width*20/100, height-height/5+1)
		WMap["Grid"].SetSize(width*20/100, height*2/10)
		WMap["User Input"].SetSize((width * 3 / 10), 5)
		WMap["Enviroment"].SetSize(width/2, height-height/5)
		WMap["Process"].SetSize(width/2, height*2/10)
		WMap["Process"].SetAutoScroll(true)
		//SET POSITION
		WMap["Log"].SetPosition(width-width*30/100, 0)
		WMap["Persona"].SetPosition(0, 0)
		WMap["Grid"].SetPosition(0, height-height*2/10)
		WMap["User Input"].SetPosition(width-(width*3/10), height-5)
		WMap["Enviroment"].SetPosition(width*20/100, 1)
		WMap["Help"].SetPosition(width*20/100, height*8/10+1)
		WMap["Programs"].SetPosition(width*20/100, height*8/10+1)
		WMap["System"].SetPosition(width*20/100, height*8/10+1)
		WMap["Actions"].SetPosition(width*20/100, height*8/10+1)
		WMap["Process"].SetPosition(width*20/100, height*8/10+1)

		draw()
		return true
	})

	congo.ActionMap.Apply()
}

func draw() {

	WMap["Persona"].WDraw()
	WMap["Top"].WDraw()
	WMap["Grid"].WDraw()
	WMap["Help"].WDraw()
	//WMap["Enviroment"].WDraw()
	WMap["Programs"].WDraw()
	WMap["System"].WDraw()
	WMap["Actions"].WDraw()
	WMap["Enviroment"].WDraw()
	WMap["Log"].WDraw()
	WMap["User Input"].WDraw()
	WMap["Process"].SetAutoScroll(true)
	WMap["Process"].WDraw()

	//WMap["Top"].WDraw()
	//windowList[0].(*congo.TWindow).WPrintLn("sdkjf11111111111h", congo.ColorDefault)
	//congo.PrintText(1, 1, fmt.Sprintf("%v-%T-%s ", info, info, info))
	//congo.PrintText(1, 5, string(key))
	//WMap["Enviroment"].WPrintLn(fmt.Sprintf("%v - ", MActions.MActionMap), congo.ColorGreen)
	//refreshEnviromentWin()
	//refreshPersonaWin()
	//refreshGridWin()
	//WMap["Enviroment"].WPrintLn(fmt.Sprintf("%v - ", MActions.MActionMap), congo.ColorGreen)
	congo.Flush()
}

func main() {
	err := congo.Init()
	if err != nil {
		panic(err)
	}
	defer congo.Close()
	//////////////////////////////////
	initialize()
	//ObjDB := map[string]IObj{}
	//var objDB = map[int]IObj{}
	GREEN = congo.ColorGreen
	width, height = congo.GetSize()
	STime = generateCurrentTime()
	WMap["Log"] = congo.NewWindow(width-(width*3/10), 0, (width * 3 / 10), height-4, "Log", "Block")
	WMap["Persona"] = congo.NewWindow(0, 0, width*20/100, height-height/5+1, "Persona", "Block")
	WMap["Grid"] = congo.NewWindow(0, height-height*2/10, width*20/100, height*2/10, "Grid", "Block")
	WMap["User Input"] = congo.NewWindow(width-(width*3/10), height-5, (width * 3 / 10), 5, "User Input", "Block") //width*3/10)+1, height, (width*3/10)-2, 3
	WMap["Actions"] = congo.NewWindow(width*20/100, 1, width/2, height-height/5, "Actions", "Block")
	WMap["Enviroment"] = congo.NewWindow(width*20/100, 1, width/2, height-height/5, "Enviroment", "Block")
	WMap["Help"] = congo.NewWindow(width*20/100, 1, width/2, height-height/5, "Help", "Block")
	WMap["Programs"] = congo.NewWindow(width*20/100, 1, width/2, height-height/5, "Programs", "Block")
	WMap["System"] = congo.NewWindow(width*20/100, 1, width/2, height-height/5, "System", "Block")
	WMap["Process"] = congo.NewWindow(width*20/100, height*8/10+1, width/2, height*2/10, "Process", "Block")
	WMap["Enviroment"] = congo.NewWindow(width*20/100, 1, width/2, height-height/5, "Enviroment", "Block")
	WMap["Top"] = congo.NewWindow(width*20/100, 0, width/2+4, 3, "Top", "Block")
	WMap["Top"].SetBorderVisibility(false)
	multiPrintTo("Top", "███ {GREEN}[F1]HELP{DEFAULT} ███ {DEFAULT}[F2]ENVIROMENT{DEFAULT} ███ {GREEN}[F3]PROGRAMS{DEFAULT} ███ {GREEN}[F4]System{DEFAULT} ███ {GREEN}[F5]ACTIONS{DEFAULT} ████████████████████")

	//multPrint("Log", "{GREEN}Some text 0")

	//w4.SetBorderVisibility(false)
	/*windowList = append(windowList, w1)
	windowList = append(windowList, w2)
	windowList = append(windowList, w3)
	windowList = append(windowList, w4)
	windowList = append(windowList, WMap["Enviroment"])
	windowList = append(windowList, w6)
	windowList = append(windowList, w7)
	w1.SetAutoScroll(true) */ //

	createDefaultGrids()

	player = NewPersona("Unknown", "<UNREGISTRATED>")
	Matrix = player.GetGrid().NewHost("Matrix", 0)
	ObjByNames[player.GetName()] = player
	if o, ok := ObjByNames["player"]; ok {
		WMap["Log"].WPrintLn("Player Exist: " + o.GetName())
	}
	colortag := "{GREEN}"
	WMap["Log"].WPrintLn(colortag + "Connection to Matrix established...")
	WMap["Log"].WPrintLn(colortag + "...Identity spoofed")
	WMap["Log"].WPrintLn(colortag + "...Encryption keys generated")
	WMap["Log"].WPrintLn(colortag + "...Connected to onion routers")
	WMap["Log"].WPrintLn(colortag + "...Enter Login:")

	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
	kbd := congo.CreateKeyboard()
	kbd.StartKeyboard()
	for !canClose {
		//draw()
		if congo.IsChanged() {
			draw()
			congo.ResetUpdate()
		}
		if kbd.KeyPressed() {
			ev := kbd.ReadEvent()
			if ev.GetEventType() == "Keyboard" {
				var char string
				key := ev.(*congo.KeyboardEvent).GetRune()
				if key != 0 {
					congo.PostUpdate()
					char = string(key)
					WMap["User Input"].WPrint(char)
					nom := strconv.Itoa(int(key))
					noms := strings.Split(nom, "")
					s := 0
					for i := range noms {
						g, _ := strconv.Atoi(noms[i])
						s = s + g
						if s > 6 {
							s = s - 6
						}
					}
				}
			}
			info = ev
			congo.HandleEvent(ev)
		} else {
			time.Sleep(1)
		}
		checkTurn()
	}
}
