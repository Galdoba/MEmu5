package main

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/ConGo/utils"
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
	windowList       []interface{}
	//objectList       []IObj

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
	command     string
)

//ObjByNames -
var ObjByNames = map[string]IObj{
//"null": nil,
}

var CFDBMap = map[int]ComplexForm{
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
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				activeWindow = i
			}
			congo.WindowsMap.ByTitle[windowList[i]].SetFocus(false)
		}
		if activeWindow == 4 {
			activeWindow = 0
		} else {
			activeWindow++
		}

		congo.WindowsMap.ByTitle[windowList[activeWindow]].SetFocus(true)
		return true
	})

	congo.NewKeyboardAction("Exit_Programm", "<esc>", "", func(ev *congo.KeyboardEvent) bool {
		canClose = true
		return true
	})

	congo.NewKeyboardAction("Add Line", "<space>", "", func(ev *congo.KeyboardEvent) bool {
		congo.WindowsMap.ByTitle["User Input"].WPrint(" ", congo.ColorGreen)
		return true
	})

	congo.NewKeyboardAction("Input", "<ENTER>", "", func(ev *congo.KeyboardEvent) bool {
		input := congo.WindowsMap.ByTitle["User Input"].WRead()
		if len(input) >= 2 {
			sl := []byte(input)
			sl = sl[:len(sl)-0]
			input = string(sl)
		} else {
			input = ""
		}
		if input != "" {
			congo.WindowsMap.ByTitle["Log"].SetAutoScroll(true)
			UserInput(player.GetName() + ">" + input)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn(input)
		}
		congo.WindowsMap.ByTitle["User Input"].WClear()
		draw()
		return true
	})

	congo.NewKeyboardAction("Delete_input_char", "<BACKSPACE>", "", func(ev *congo.KeyboardEvent) bool {
		input := congo.WindowsMap.ByTitle["User Input"].WRead()
		if len(input) > 2 {
			sl := []byte(input)
			sl = sl[:len(sl)-3]
			input = string(sl)
		} else {
			congo.WindowsMap.ByTitle["User Input"].WClear()
		}
		congo.WindowsMap.ByTitle["User Input"].WClear()
		congo.WindowsMap.ByTitle["User Input"].WPrint(input, congo.ColorGreen)
		draw()
		return true
	})

	congo.NewKeyboardAction("Expand Persona Window", "<F1>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		congo.WindowsMap.ByTitle["Persona"].SetPosition(0, 0)
		congo.WindowsMap.ByTitle["Persona"].SetSize((width*70)/100, height)
		congo.WindowsMap.ByTitle["Grid"].SetPosition(width+1, 0)
		congo.WindowsMap.ByTitle["Enviroment"].SetPosition(width, 0)
		congo.WindowsMap.ByTitle["Process"].SetPosition(width, 0)
		congo.WindowsMap.ByTitle["Persona"].WDraw()
		return true
	})

	congo.NewKeyboardAction("Expand Grid Window", "<F2>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		congo.WindowsMap.ByTitle["Grid"].SetPosition(0, 0)
		congo.WindowsMap.ByTitle["Grid"].SetSize((width*70)/100, height)
		congo.WindowsMap.ByTitle["Grid"].WDraw()
		congo.WindowsMap.ByTitle["Persona"].SetPosition(width, 0)
		congo.WindowsMap.ByTitle["Enviroment"].SetPosition(width, 0)
		congo.WindowsMap.ByTitle["Process"].SetPosition(width, 0)
		simpleTest(-1, 3, 4, 2)
		return true
	})

	congo.NewKeyboardAction("Expand Enviroment Window", "<F3>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)
		congo.WindowsMap.ByTitle["Enviroment"].SetPosition(0, 0)
		congo.WindowsMap.ByTitle["Enviroment"].SetSize((width*70)/100, height)
		congo.WindowsMap.ByTitle["Grid"].SetPosition(0, 0)
		congo.WindowsMap.ByTitle["Persona"].SetPosition(width, 0)
		congo.WindowsMap.ByTitle["Grid"].SetSize((width*70)/100, height)
		congo.WindowsMap.ByTitle["Grid"].WDraw()
		congo.WindowsMap.ByTitle["Persona"].SetPosition(width, 0)
		congo.WindowsMap.ByTitle["Process"].SetPosition(width, 0)
		return true
	})

	congo.NewKeyboardAction("Defauil View", "<F12>", "", func(ev *congo.KeyboardEvent) bool {
		congo.FillRect(0, 0, width/100*70, height, ' ', congo.ColorBlack, congo.ColorBlack)

		congo.WindowsMap.ByTitle["Persona"].SetSize(width*20/100, height-height/5+1)
		congo.WindowsMap.ByTitle["Grid"].SetSize(width*20/100, height*2/10)
		congo.WindowsMap.ByTitle["Enviroment"].SetSize(width/2, height-height/5+1)
		//SET POSITION

		congo.WindowsMap.ByTitle["Persona"].SetPosition(0, 0)
		congo.WindowsMap.ByTitle["Grid"].SetPosition(0, height-height*2/10)
		congo.WindowsMap.ByTitle["Enviroment"].SetPosition(width*20/100, 0)
		congo.WindowsMap.ByTitle["Process"].SetPosition(width*20/100, height*8/10+1)
		//NewFile("File ")
		//NewDevice("Camera ", 2)
		//addDevice("Camera3")
		//addDevice("Camera5")
		//addDevice("Erika MCD-1")
		//NewDevice("Erika MCD-1", 4)
		congo.WindowsMap.ByTitle["User Input"].WPrint("matrix search>host", congo.ColorGreen)
		draw()
		return true
	})

	congo.NewKeyboardAction("Move_selector_down", "<down>", "", func(ev *congo.KeyboardEvent) bool { //KeyboardEvent
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				index := congo.WindowsMap.ByTitle[windowList[i]].GetScrollIndex()
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(utils.Min(index+1, congo.WindowsMap.ByTitle[windowList[i]].GetStoredRows()-congo.WindowsMap.ByTitle[windowList[i]].GetPrintableHeight()+2))
				if congo.WindowsMap.ByTitle[windowList[i]].GetStoredRows() < congo.WindowsMap.ByTitle[windowList[i]].GetPrintableHeight() {
					congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(0)
				}
				congo.WindowsMap.ByTitle[windowList[i]].WDraw()

			}
		}

		return true
	})

	congo.NewKeyboardAction("Move_selector_up", "<up>", "", func(ev *congo.KeyboardEvent) bool {
		windowList := congo.WindowsMap.GetNames()
		sort.Strings(windowList)
		for i := range windowList {
			if congo.WindowsMap.ByTitle[windowList[i]].InFocus() == true {
				congo.WindowsMap.ByTitle[windowList[i]].SetAutoScroll(false)
				index := congo.WindowsMap.ByTitle[windowList[i]].GetScrollIndex()
				congo.WindowsMap.ByTitle[windowList[i]].SetScrollIndex(utils.Max(index-1, 0))
				congo.WindowsMap.ByTitle[windowList[i]].WDraw()
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
		congo.WindowsMap.ByTitle["Log"].SetSize((width * 3 / 10), height)
		congo.WindowsMap.ByTitle["Persona"].SetSize(width*20/100, height-height/5+1)
		congo.WindowsMap.ByTitle["Grid"].SetSize(width*20/100, height*2/10)
		congo.WindowsMap.ByTitle["User Input"].SetSize((width * 3 / 10), 3)
		congo.WindowsMap.ByTitle["Enviroment"].SetSize(width/2, height-height/5+1)
		congo.WindowsMap.ByTitle["Process"].SetSize(width/2, height*2/10)
		congo.WindowsMap.ByTitle["Process"].SetAutoScroll(true)
		//SET POSITION
		congo.WindowsMap.ByTitle["Log"].SetPosition(width-width*30/100, 0)
		congo.WindowsMap.ByTitle["Persona"].SetPosition(0, 0)
		congo.WindowsMap.ByTitle["Grid"].SetPosition(0, height-height*2/10)
		congo.WindowsMap.ByTitle["User Input"].SetPosition(width-(width*3/10), height-3)
		congo.WindowsMap.ByTitle["Enviroment"].SetPosition(width*20/100, 0)
		congo.WindowsMap.ByTitle["Process"].SetPosition(width*20/100, height*8/10+1)

		draw()
		return true
	})

	congo.ActionMap.Apply()
}

func draw() {

	congo.WindowsMap.ByTitle["Persona"].WDraw()
	congo.WindowsMap.ByTitle["Grid"].WDraw()
	congo.WindowsMap.ByTitle["Enviroment"].WDraw()
	congo.WindowsMap.ByTitle["Log"].WDraw()
	congo.WindowsMap.ByTitle["User Input"].WDraw()
	congo.WindowsMap.ByTitle["Process"].SetAutoScroll(true)
	congo.WindowsMap.ByTitle["Process"].WDraw()
	//windowList[0].(*congo.TWindow).WPrintLn("sdkjf11111111111h", congo.ColorDefault)
	//congo.PrintText(1, 1, fmt.Sprintf("%v-%T-%s ", info, info, info))
	//congo.PrintText(1, 5, string(key))
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("%v - ", MActions.MActionMap), congo.ColorGreen)
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("%v - ", MActions.MActionMap), congo.ColorGreen)
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

	width, height = congo.GetSize()
	dur := time.Second / 3
	STime = generateCurrentTime()
	w1 := congo.NewWindow(width-(width*3/10), 0, (width * 3 / 10), height, "Log", "Block")
	w2 := congo.NewWindow(0, 0, width*20/100, height-height/5+1, "Persona", "Block")
	w3 := congo.NewWindow(0, height-height*2/10, width*20/100, height*2/10, "Grid", "Block")
	w4 := congo.NewWindow(width-(width*3/10), height-3, (width * 3 / 10), 3, "User Input", "Block") //width*3/10)+1, height, (width*3/10)-2, 3
	w5 := congo.NewWindow(width*20/100, 0, width/2, height-height/5+1, "Enviroment", "Block")
	w6 := congo.NewWindow(width*20/100, height*8/10+1, width/2, height*2/10, "Process", "Block")
	//w4.SetBorderVisibility(false)
	windowList = append(windowList, w1)
	windowList = append(windowList, w2)
	windowList = append(windowList, w3)
	windowList = append(windowList, w4)
	windowList = append(windowList, w5)
	windowList = append(windowList, w6)
	w1.SetAutoScroll(true) //

	createDefaultGrids()

	//congo.WindowsMap.ByTitle["Grid"].SetPosition(0, height-height*2/10)

	player = NewPersona("Unknown", "<UNREGISTRATED>")
	//player.SetName("sdfjkh")
	//player, _ = ImportPlayerFromDB("Unknown") //.(*TPersona)
	Matrix = player.GetGrid().NewHost("Matrix", 0)
	ObjByNames[player.GetName()] = player
	if o, ok := ObjByNames["player"]; ok {
		w1.WPrintLn("Player Exist: "+o.GetName(), congo.ColorGreen)
	}

	w1.WPrintLn("Connection to Matrix established...", congo.ColorGreen)
	time.Sleep(dur)
	draw()
	w1.WPrintLn("...Identity spoofed", congo.ColorGreen)
	time.Sleep(dur)
	draw()
	w1.WPrintLn("...Encryption keys generated", congo.ColorGreen)
	time.Sleep(dur)
	draw()
	w1.WPrintLn("...Connected to onion routers", congo.ColorGreen)
	time.Sleep(dur)
	draw()
	w1.WPrintLn("...Enter Login:", congo.ColorGreen)
	time.Sleep(dur)
	draw()
	//startCombatTurn()

	/*	w1.WPrintLn("Begin Session:", congo.ColorGreen)
		w1.WPrintLn(generateCurrentTime(), congo.ColorGreen)*/

	//gridList[0].(*TGrid).NewHost("Ares Host", 0)

	//gridList[0].(*TGrid).NewHost("Ares Host", 0)

	/*for i:= 0; i<5; i++ {
		w1.WPrintLn("ksdfhaksjhfasdfh" + strconv.Itoa(i))
	}
	w1.WPrintLn("15")*/
	//congo.WindowsMap.ByTitle["W1"].WDraw()

	kbd := congo.CreateKeyboard()
	kbd.StartKeyboard()
	//SourceIcon = pickObjByID(player.GetID())
	//doAction("WAIT")

	for !canClose {
		draw()
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
					w4.WPrint(char, congo.ColorGreen)

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
					//w6.WPrint(fmt.Sprintf("%v ", s), congo.ColorGreen)
				}

			}
			info = ev
			congo.HandleEvent(ev)
		} else {
			time.Sleep(1)
		}
		//checkTurn()
		/*	if player.GetWaitFlag() {
			SourceIcon = pickObjByID(player.GetID())
			doAction("WAIT")
		}*/
		checkTurn() // Рекурсия
	}

}
