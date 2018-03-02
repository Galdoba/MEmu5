package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/utils"
)

func iStr(i int) string {
	return strconv.Itoa(i)
}

func refreshHelpWin() {
	WMap["Help"].WClear()
	WMap["Help"].WPrintLn("{default}BASIC CONTROL:")
	WMap["Help"].WPrintLn("{default}<UP>        {green}- Scroll active window up")
	WMap["Help"].WPrintLn("{default}<DOWN>      {green}- Scroll active window down")
	WMap["Help"].WPrintLn("{default}<BACKSPACE> {green}- delete last character from User Input window")
	WMap["Help"].WPrintLn("{default}<ENTER>     {green}- initiate command typed in User Input window")
	WMap["Help"].WPrintLn("")
	WMap["Help"].WPrintLn("{default}COMMANDS:")
	WMap["Help"].WPrintLn("{green}All commands folowing the same patern. Input is not cASe seNseTIve")
	WMap["Help"].WPrintLn(" {default}|MATRIX_ACTION| > |TARGET| > |SPECIFICATIONS|")
	WMap["Help"].WPrintLn("")
	WMap["Help"].WPrintLn("{default}SPECIFICATIONS:")
	WMap["Help"].WPrintLn(" |-PL| {green} - Spend point of Edge and use 'Push the Limit' rules for this roll")
	WMap["Help"].WPrintLn("{yellow}Example: {Default}Matrix Perception>File 14>-pl")
	WMap["Help"].WPrintLn(" {default}|-2M|{green} ({default}|-3M|{green}) - Atempt to place 2(3) marks with 1 BF or HotF action with respective modifier")
	WMap["Help"].WPrintLn("{yellow}Example: {default}Brute Force>File 14>-2m")
	WMap["Help"].WPrintLn("{yellow}Example: {default}Hack on the Fly>File 14>-3m")
	WMap["Help"].WPrintLn("")
	WMap["Help"].WPrintLn("{default}EXAMPLES:")
	WMap["Help"].WPrintLn(" ")
	WMap["Help"].WPrintLn(" {default}Login>Tom - {green}will search {default}PlayerDB.txt{green} file and load persona with alias {default}'Tom'{green} if found")
	WMap["Help"].WPrintLn(" ")
	WMap["Help"].WPrintLn(" {default}Matrix Search>Host>Dante's Inferno - {green}will initiate search for host with name {default}'Dante's Inferno'")
	WMap["Help"].WPrintLn(" {yellow}Warning: {green}Matrix search might take up to 20 combat turns to complete depending on roll result. Check {default}Shadowrun Core Rulebook (p.241){GREEN} for dicepool and time calculations")
}

func refreshSystemWin() {
	WMap["System"].WClear()
	printToWindow("System", "{yellow}31313213{green}n[21")
	printToWindow("System", "Great Text "+iStr(5+9))
	programs := player.GetDevice().GetCyberSoftwareList()
	for i := range programs {
		WMap["System"].WPrintLn("Program " + strconv.Itoa(i+1) + ": " + programs[i].programName)
		WMap["Persona"].WPrintLn("Program "+strconv.Itoa(i+1)+": "+programs[i].programName, congo.ColorYellow)
		if programs[i].programRating > 0 {
			WMap["System"].WPrintLn("--Rating: " + strconv.Itoa(programs[i].programRating))
			WMap["Persona"].WPrintLn("--Rating: "+strconv.Itoa(programs[i].programRating), congo.ColorYellow)
		}
		WMap["Persona"].WPrintLn("--Status: "+programs[i].programStatus, congo.ColorYellow)
	}
}

func refreshProgramsWin() {
	WMap["Programs"].WClear()
	programs := player.GetDevice().GetCyberSoftwareList()
	for i := range programs {
		drawLineInWindow("Programs")
		WMap["Programs"].WPrint("{green}Program " + strconv.Itoa(i+1) + ": {white}" + programs[i].programName)
		if programs[i].programRating > 0 {
			WMap["Programs"].WPrint(" (Rating: " + strconv.Itoa(programs[i].programRating) + ")")
		}
		WMap["Programs"].WPrintLn("")
		tag := ""
		switch programs[i].programStatus {
		case "Stored":
			tag = "{default}"
		case "Running":
			tag = "{green}"
		}
		WMap["Programs"].WPrintLn("{green}Status   : " + tag + programs[i].programStatus + "{default}")
	}
}

func refreshActionsWin() {
	WMap["Actions"].WClear()
	srmod := 0
	if player.GetSilentRunningMode() {
		srmod = -2
	}
	hotVR := 0
	if player.GetSimSence() == "HOT-SIM" {
		hotVR = 2
	}
	WMap["Actions"].WPrintLn("{green} Possible actions:")
	WMap["Actions"].WPrintLn(" ")
	WMap["Actions"].WPrintLn("{default}BRUTE FORCE - Cybercombat + Logic [Attack] v. Willpower + Firewall")
	WMap["Actions"].WPrintLn("Current Dicepool :")
	WMap["Actions"].WPrintLn("{green}Cybercombat Skill:  " + iStr(player.GetCyberCombatSkill()))
	WMap["Actions"].WPrintLn("Logic            :  " + iStr(player.GetLogic()))

	if player.GetSilentRunningMode() {
		WMap["Actions"].WPrintLn("Running Silent   : " + iStr(srmod))
	}
	if player.GetSimSence() == "HOT-SIM" {
		WMap["Actions"].WPrintLn("HOT-SIM VR       :  " + iStr(hotVR))

	}
	WMap["Actions"].WPrintLn("{yellow}Total            :  " + iStr(hotVR+srmod+player.GetCyberCombatSkill()+player.GetLogic()))
}

func refreshPersonaWin() {
	WMap["Persona"].(*congo.TWindow).WClear()
	if player.GetName() == "Unknown" {
		WMap["Persona"].WPrintLn("{green}User Name: {yellow}<<UNREGISTRATED>>(GM-mode)", congo.ColorGreen)
	} else {
		WMap["Persona"].WPrintLn("{green}User Name: {default}"+player.GetName(), congo.ColorGreen)
	}
	if player.GetName() != "Unknown" {
		WMap["Persona"].WPrintLn("{green}Device type: {default}"+player.GetDevice().deviceType, congo.ColorGreen)
		WMap["Persona"].WPrintLn("{green}Device model: {default}"+player.GetDevice().model, congo.ColorGreen)
		WMap["Persona"].WPrintLn("{green}Persona User Mode: {default}"+player.GetSimSence(), congo.ColorGreen)
	}
	WMap["Persona"].WPrintLn("{green}Grid: {default}"+player.GetGrid().GetGridName(), congo.ColorGreen)
	if checkLinkLock(player) == true {
		WMap["Persona"].WPrintLn("{red}WARNING: LINK-LOCK DETECTED!", congo.ColorRed)
	}
	if player.GetPhysicalLocation() == true {
		WMap["Persona"].WPrintLn("{red}WARNING: Physical location tracked!{default}", congo.ColorYellow)
	}
	drawLineInWindow("Persona")
	WMap["Persona"].WPrintLn("{green}Running Programs: ", congo.ColorGreen)
	var rPrgLst []string
	var rPrgLstRat []int
	loadedPrgsQty := 0
	programs := player.GetDevice().GetCyberSoftwareList()
	for j := range programs {
		if programs[j].programStatus == "Running" {
			rPrgLst = append(rPrgLst, programs[j].programName)
			rPrgLstRat = append(rPrgLstRat, programs[j].programRating)
			loadedPrgsQty++
		}
	}
	if loadedPrgsQty > player.GetDevice().GetMaxRunningPrograms() {
		player.CrashRandomProgram()
	}
	for i := 0; i < player.GetDevice().GetMaxRunningPrograms(); i++ {
		if len(rPrgLst) < player.GetDevice().GetMaxRunningPrograms() {
			rPrgLst = append(rPrgLst, "--EMPTY--")
			rPrgLstRat = append(rPrgLstRat, 0)
		}
		congo.WindowsMap.ByTitle["Persona"].WPrint("{green} Slot " + strconv.Itoa(i+1) + ": {default}" + rPrgLst[i])
		if rPrgLstRat[i] > 0 {
			congo.WindowsMap.ByTitle["Persona"].WPrint("(Rating: " + strconv.Itoa(rPrgLstRat[i]) + ")")
		}
		WMap["Persona"].WPrintLn("{default}")
	}
	WMap["Persona"].WPrintLn("{green}Attribute Array: ")
	WMap["Persona"].WPrintLn("{green} Attack        : {default}" + iStr(player.GetAttack()))
	WMap["Persona"].WPrintLn("{green} Sleaze        : {default}" + iStr(player.GetSleaze()))
	WMap["Persona"].WPrintLn("{green} DataProcess   : {default}" + iStr(player.GetDataProcessing()))
	WMap["Persona"].WPrintLn("{green} Firewall      : {default}" + iStr(player.GetFirewall()))
	if pl, ok := player.(ITechnom); ok {
		WMap["Persona"].WPrintLn(" {green}RESONANCE     : {default}" + iStr(pl.GetResonance()))
	} else {
		WMap["Persona"].WPrintLn("{green} Device Rating : {default}" + iStr(player.GetDeviceRating()))
	}
	drawLineInWindow("Persona")
	WMap["Persona"].WPrintLn("{green}Condition Monitor:", congo.ColorGreen)
	if player.GetDevice().model != "Living Persona" {
		congo.WindowsMap.ByTitle["Persona"].WPrint("Matrix  : {WHITE}", congo.ColorDefault)
		if player.GetMatrixCM() < 1 {
			congo.WindowsMap.ByTitle["Persona"].WPrint("{red}DEVICE BRICKED!!!")
		} else {
			for i := 0; i < player.GetMatrixCM(); i++ {
				congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorGreen)
			}
		}
		WMap["Persona"].WPrintLn(" {green}")
	}
	congo.WindowsMap.ByTitle["Persona"].WPrint("Stun    : {WHITE}", congo.ColorYellow)
	if player.GetStunCM() < 1 {
		congo.WindowsMap.ByTitle["Persona"].WPrint("{RED}UNCONSCIOUS!!!", congo.ColorRed)
	} else {
		for i := 0; i < player.GetStunCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorYellow)
		}
	}
	WMap["Persona"].WPrintLn(" {green}")
	congo.WindowsMap.ByTitle["Persona"].WPrint("Physical: {WHITE}", congo.ColorYellow)
	if player.GetPhysCM() < player.GetBody() {
		congo.WindowsMap.ByTitle["Persona"].WPrint("DEAD!!!", congo.ColorRed)
	} else if player.GetPhysCM() < 1 {
		congo.WindowsMap.ByTitle["Persona"].WPrint("CRITICAL!!!", congo.ColorRed)
	} else {
		for i := 0; i < player.GetPhysCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorRed)
		}
	}
	WMap["Persona"].WPrintLn(" {default}")
	drawLineInWindow("Persona")
	var silentStatus string
	if player.GetSilentRunningMode() {
		silentStatus = "{GREEN}TRUE"
	} else {
		silentStatus = "{yellow}FALSE{green}"
	}
	WMap["Persona"].WPrintLn("{GREEN}Running Silent: " + silentStatus)
	if player.GetInitiative() > 9000 {
		WMap["Persona"].WPrintLn("Persona Initiative: {red}null")
	} else {
		WMap["Persona"].WPrintLn("Persona Initiative: {WHITE}" + iStr(player.GetInitiative()))
	}
	if player.IsConnected() == false {
		WMap["Persona"].WPrintLn("Persona disconnected...", congo.ColorRed)
	}
	if player.GetEdge() > 0 {
		WMap["Persona"].WPrintLn("{green}Edge : {default}" + iStr(player.GetEdge()) + "/" + iStr(player.GetMaxEdge()))
	}
	if player.GetFullDeffenceFlag() == true {
		WMap["Persona"].WPrintLn("Full Defence = "+strconv.FormatBool(player.GetFullDeffenceFlag()), congo.ColorYellow)
	}
	drawLineInWindow("Persona")
	WMap["Persona"].WPrintLn("--DEBUG--Total Objects: "+strconv.Itoa(len(ObjByNames)), congo.ColorYellow)
	WMap["Persona"].WPrintLn("--DEBUG--waitFlag: "+strconv.FormatBool(player.GetWaitFlag()), congo.ColorYellow)
	for i := range player.GetSpecializationList() {
		WMap["Persona"].WPrintLn(player.GetSpecializationList()[i])
	}
	WMap["Persona"].WPrintLn("--DEBUG--Free/Simple Actions: "+strconv.Itoa(player.GetFreeActionsCount())+"/"+strconv.Itoa(player.GetSimpleActionsCount()), congo.ColorGreen)
	totalMarks := player.CountMarks()
	WMap["Persona"].WPrintLn("Confirmed Marks on Persona: "+strconv.Itoa(totalMarks), congo.ColorYellow)
	WMap["Persona"].WPrintLn("Matrix Search in: "+strconv.Itoa(player.GetSearchResultIn()), congo.ColorYellow)
	for i := range player.GetSearchProcess().SearchIconName {
		name := player.GetSearchProcess().SearchIconName[i]
		objType := player.GetSearchProcess().SearchIconType[i]
		timeTotal := player.GetSearchProcess().SearchTime[i]
		if timeTotal == 0 {
			//player.UpdateSearchProcess()
		}
		timeSpent := player.GetSearchProcess().SpentTurns[i]
		WMap["Persona"].WPrintLn("Search: "+objType+" "+name, congo.ColorGreen)
		currentPer := 0
		//turnsPart := 0
		if timeSpent != 0 {
			//turnsPart = (100 / timeTotal)
			//r := player.GetInitiative()/10 + 1
			currentPer = utils.Min(((100 / timeTotal) * (timeSpent)), 100)
		}
		WMap["Persona"].WPrintLn(" Progress: "+strconv.Itoa(currentPer)+"%", congo.ColorGreen)
	}
	for i := range player.GetDownloadProcess().DownloadIconName {
		name := player.GetDownloadProcess().DownloadIconName[i]
		downloaded := player.GetDownloadProcess().DownloadedData[i]
		size := player.GetDownloadProcess().FileSize[i]
		WMap["Persona"].WPrintLn("Downloading file: "+name, congo.ColorGreen)
		WMap["Persona"].WPrintLn("Progress: "+strconv.Itoa(downloaded)+" of "+strconv.Itoa(size)+" Mp", congo.ColorGreen)
	}
	drawLineInWindow("Persona")
	WMap["Persona"].WPrintLn(fmt.Sprintf("CFDB: %v", CFDBMap), congo.ColorYellow)
	//for i := 0; i < len(CFDBMap); i++ {
	//	WMap["Persona"].WPrintLn(fmt.Sprintf("%v", CFDBMap[i+1]), congo.ColorYellow)
	//}
	//fow := player.GetFieldOfView()
	//WMap["Persona"].WPrintLn(fmt.Sprintf("FoW: %v", fow), congo.ColorYellow)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("0", congo.ColorDefault)

}

func refreshGridWin() {
	if player.GetMatrixCM() < 1 {
		WMap["Persona"].WPrintLn("YOU ARE BRICKED!!!", congo.ColorRed)
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("YOU ARE BRICKED!!!", congo.ColorRed)
	}
	WMap["Grid"].(*congo.TWindow).WClear()
	congo.WindowsMap.ByTitle["Grid"].WPrintLn("Grid:", congo.ColorGreen)
	congo.WindowsMap.ByTitle["Grid"].WPrintLn(player.GetGrid().GetGridName(), congo.ColorGreen)
	if player.CheckRunningProgram("Baby Monitor") {
		warningColor := congo.ColorGreen
		if player.GetGrid().GetOverwatchScore() < 20 {
			warningColor = congo.ColorGreen
		} else if player.GetGrid().GetOverwatchScore() < 31 {
			warningColor = congo.ColorYellow
		} else {
			warningColor = congo.ColorRed
		}
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("Overwatch Score: "+strconv.Itoa(player.GetGrid().GetOverwatchScore()), warningColor)
	} else {
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("Overwatch Score: "+strconv.Itoa(player.GetGrid().GetLastSureOS())+" or more", congo.ColorYellow)
	}

	congo.WindowsMap.ByTitle["Grid"].WPrintLn("Host:", congo.ColorGreen)
	host := player.GetHost().name
	if host == "Matrix" {
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("--not in Host--", congo.ColorGreen)
	} else {
		congo.WindowsMap.ByTitle["Grid"].WPrintLn(" "+host, congo.ColorYellow)
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("Host Alert Status: ", congo.ColorGreen)
		statusColor := congo.ColorDefault
		alert := player.GetHost().GetHostAlertStatus()
		switch alert {
		case "No Alert":
			statusColor = congo.ColorGreen
		case "Passive Alert":
			statusColor = congo.ColorYellow
		case "Active Alert":
			statusColor = congo.ColorRed
		default:
		}
		congo.WindowsMap.ByTitle["Grid"].WPrintLn(" "+alert, statusColor)
	}

}

func getSortedKeysByType(objType string) []int {
	var keys []int
	for _, obj := range ObjByNames {
		if obj, ok := obj.(IObj); ok {
			key := obj.GetID()
			if obj.GetType() == objType {
				keys = append(keys, key)
			}
		}
	}
	sort.Ints(keys)
	return keys
}

func drawLineInWindow(windowName string) {
	for i := 0; i < congo.WindowsMap.ByTitle[windowName].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle[windowName].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle[windowName].WPrintLn("", congo.ColorDefault)
}

func treePrefix(n int, cap int, depth int) string {
	prefix := ""
	for i := 0; i < depth; i++ {

		if depth-1 > i {
			prefix = prefix + /* "├─"*/ "│ "
			continue
		}
		if n+1 == cap {
			prefix = prefix + "└─"
		} else {
			prefix = prefix + "├─"
		}
	}
	return prefix
}

func getIconsInGrid(g IGrid) []int {
	var keys []int
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok && icon.GetGrid() == g {
			if icon.GetType() != "Host" && icon.GetHost() != Matrix {
				continue
			}
			if icon.GetType() == "Host" || icon.GetType() == "Persona" || icon.GetType() == "Device" || icon.GetType() == "Sprite" || icon.GetType() == "Agent" {
				keys = append(keys, icon.GetID())
			}
		}
	}
	sort.Ints(keys)
	return keys
}

func getIconsInHost(h IHost) []int {
	var keys []int
	var pKeys []int
	var iKeys []int
	var fKeys []int
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IPersona); ok && icon.GetHost() == h {
			pKeys = append(pKeys, icon.GetID())
		}
	}
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIC); ok && icon.GetHost() == h {
			iKeys = append(iKeys, icon.GetID())
		}
	}
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IFile); ok && icon.GetHost() == h {
			fKeys = append(fKeys, icon.GetID())
		}
	}
	sort.Ints(pKeys)
	sort.Ints(iKeys)
	sort.Ints(fKeys)
	keys = append(keys, pKeys...)
	keys = append(keys, iKeys...)
	keys = append(keys, fKeys...)
	return keys
}

func refreshEnviromentWin() {
	congo.WindowsMap.ByTitle["Enviroment"].WClear()

	WMap["Enviroment"].WPrintLn("{default}Combat Turn №:" + iStr(CombatTurn))
	WMap["Enviroment"].WPrintLn("{default}Initiative Pass №:" + iStr(InitiativePass))
	WMap["Enviroment"].WPrintLn("{default}Time: " + STime)
	drawLineInWindow("Enviroment")
	//prefix := ""
	keysForHost := getSortedKeysByType("Host")
	keysForFile := getSortedKeysByType("File")
	keysForIC := getSortedKeysByType("IC")
	keysForGrids := getSortedKeysByType("Grid")

	//colorTag := "{default}"
	//host := player.GetHost()
	WMap["Enviroment"].WPrintLn("{default}" + "Matrix:")
	//WMap["Enviroment"].WPrintLn("│─├─└─")
	//colorCode := "{default}"
	for g := range keysForGrids {
		var prefix []string
		if len(prefix) < 1 {
			prefix = append(prefix, "  ")
		}
		prefix[0] = "├─"
		if g+1 == len(keysForGrids) {
			prefix[0] = "└─"
		}
		grid := pickObjByID(keysForGrids[g]).(IGrid)
		if host, ok := grid.(IHost); ok { //skip Host
			host.GetName()
			continue
		}
		WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + grid.GetGridName())
		iconsInGridID := getIconsInGrid(grid)
		for h := range iconsInGridID {
			if len(prefix) < 2 {
				prefix = append(prefix, "")
			}
			if prefix[0] == "├─" {
				prefix[0] = "│ "
			}
			if prefix[0] == "└─" {
				prefix[0] = "  "
			}
			icon := pickObjByID(iconsInGridID[h]).(IIcon)
			prefix[1] = "├─"
			if h+1 == len(iconsInGridID) {
				prefix[1] = "└─"
			}
			colorCode := ""
			switch icon.GetType() {
			case "Persona":
				colorCode = "{default}"
			case "Host":
				colorCode = "{yellow}"
			default:
				colorCode = "{default}"
			}
			iconType := icon.GetType()

			WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "" + icon.GetName() + playerMarks(icon) + colorCode + " (" + iconType + "){default}")
			if prefix[1] == "└─" {
				prefix[1] = "  "
			}
			if prefix[1] == "├─" {
				prefix[1] = "│ "
			}
			WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + icon.GetDevice().GetModel() + " (" + knownStats(icon) + ")")
			if icon.GetSilentRunningMode() {
				WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "{red}Running Silent{default}")
			}

			/////////////GO INSIDE THE HOST
			if host, ok := icon.(IHost); ok {
				silentIcons := 0
				alert := host.GetHostAlertStatus()
				switch alert {
				case "No Alert":
					alert = "{green}No Alert{default}"
				case "Passive Alert":
					alert = "{Yellow}Passive Alert{default}"
				case "Active Alert":
					alert = "{Red}Active Alert{default}"
				}
				iconsID := getIconsInHost(host)
				for ic := range iconsID {
					if len(prefix) < 3 {
						prefix = append(prefix, "")
					}
					prefix[2] = "├─"
					if len(iconsID) == ic+1 {
						prefix[2] = "└─"
					}
					if prefix[1] == "└─" {
						prefix[1] = "  "
					}
					if prefix[1] == "├─" {
						prefix[1] = "│ "
					}
					icon := pickObjByID(iconsID[ic]).(IIcon)
					if host, ok := icon.(IHost); ok { //skip Host
						host.GetName()
						continue
					}
					if h+1 == len(iconsInGridID) {
						prefix[1] = "  "
					}
					if persona, ok := icon.(IPersona); ok {
						whatCanSee := player.GetFieldOfView().KnownData[persona.GetID()]
						WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + persona.GetName() + playerMarks(persona) + " (" + knownStats(persona) + ")")
						prefix[2] = "│   "
						if ic+1 == len(iconsID) {
							//	prefix[1] = "  "
							prefix[2] = "  "
						}
						if whatCanSee[2] != "Unknown" || persona.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							cm := ""
							for i := 0; i < persona.GetMatrixCM(); i++ {
								cm = cm + "X "
							}
							WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "└Matrix Condition: " + cm)
							if persona.GetSilentRunningMode() {
								WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "└{red}Running Silent{default}")
							}
						}
						prefix[2] = ""
						if persona.GetSilentRunningMode() {
							silentIcons++
						}
					}
					//WMap["Enviroment"].WPrintLn(strings.Join(prefix, "")+icon.GetName())
					if ice, ok := icon.(IIC); ok {
						whatCanSee := player.GetFieldOfView().KnownData[ice.GetID()]
						//WMap["Enviroment"].WPrintLn(strings.Join(prefix, "")+ice.GetName()+" (внутри)")
						prefix[2] = "├─"
						/*if ic+1 == len(iconsID) {
							//	prefix[1] = "  "
							prefix[2] = "    "
						}*/
						if whatCanSee[0] != "Unknown" || ice.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "{yellow}" + ice.GetName() + "{default}")
						}
						prefix[2] = "│ "
						if whatCanSee[2] != "Unknown" || ice.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							cm := ""
							for i := 0; i < ice.GetMatrixCM(); i++ {
								cm = cm + "X "
							}
							WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "  └Matrix Condition: " + cm)
						}
						prefix[2] = ""
					}
					if file, ok := icon.(IFile); ok {
						whatCanSee := player.GetFieldOfView().KnownData[file.GetID()]
						//WMap["Enviroment"].WPrintLn(strings.Join(prefix, "")+file.GetName()+": "+file.GetFileName())
						prefix[2] = "├─"
						/*if ic+1 == len(iconsID) {
							//	prefix[1] = "  "
							prefix[2] = "  "
						}*/
						if whatCanSee[0] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + file.GetName() + ": " + file.GetFileName())
							prefix[2] = "│ "
							if whatCanSee[12] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
								WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "  └Encryption Rating: " + iStr(file.GetEncryptionRating()))
							}
							if whatCanSee[3] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
								WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "  └Data Bomb Rating: " + iStr(file.GetDataBombRating()))
							}
							if whatCanSee[3] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
								WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "  └File Size: " + iStr(file.GetSize()) + " MP")
							}
							if whatCanSee[1] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
								WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "  └Last Edited: " + file.GetLastEditDate())
							}
							if file.GetSilentRunningMode() {
								WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "  └{red}File is in silent mode{default}")
							}
						}
						/*if prefix[2] = "  " {

						}*/

						prefix[2] = ""
						if file.GetSilentRunningMode() {
							silentIcons++
						}
					}
				}
				if player.GetHost() == host {
					WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "├─Alert Status: " + alert)
					WMap["Enviroment"].WPrintLn(strings.Join(prefix, "") + "└─Silent Icons: " + iStr(silentIcons))
				}
			}

		}

	}

	/*for g := range keysForGrids {
		depth := 1
		colorTag := "{default}"
		grid := pickObjByID(keysForGrids[g]).(IGrid)
		if host, ok := grid.(IHost); ok { //skip Host
			host.GetName()
			continue
		}
		if player.GetGrid() == grid {
			colorTag = "{green}"
		}
		prefix = treePrefix(g, len(keysForGrids), depth)
		WMap["Enviroment"].WPrintLn(colorTag+prefix+grid.GetGridName())
		for h := range keysForHost {
			depth = 2
			host := pickObjByID(keysForHost[h]).(IHost)
			marksOnHost := host.GetMarkSet().MarksFrom[player.GetID()]
			marksTag := ""
			for m := 0; m < marksOnHost; m++ {
				marksTag = marksTag + "*"
			}

			if host.GetGrid() == grid {
				depth = 2
				prefix = treePrefix(h, len(keysForHost), depth)
				WMap["Enviroment"].WPrintLn(colorTag+prefix+"Host: "+host.GetName()+marksTag)
				alertCol := ""
				alertStatus := host.GetHostAlertStatus()
				switch alertStatus {
				case "No Alert":
					alertCol = "{green}"
				case "Passive Alert":
					alertCol = "{Yellow}"
				case "Active Alert":
					alertCol = "{Red}"
				}
				depth = 3
				prefix = treePrefix(h, len(keysForHost), depth)
				WMap["Enviroment"].WPrintLn(colorTag+prefix+"Host Alert Status: "+alertCol+alertStatus)
				whatCanSee := player.GetFieldOfView().KnownData[host.GetID()]
				if whatCanSee[5] != "Unknown" {
					WMap["Enviroment"].WPrintLn("{default}| | | └Rating:         "+iStr(host.GetDeviceRating()))
				}
				if whatCanSee[7] != "Unknown" {
					WMap["Enviroment"].WPrintLn("{default}| | | └Attack:         "+iStr(host.GetAttack()))
				}
				if whatCanSee[8] != "Unknown" {
					WMap["Enviroment"].WPrintLn("{default}| | | └Sleaze:         "+iStr(host.GetSleaze()))
				}
				if whatCanSee[9] != "Unknown" {
					WMap["Enviroment"].WPrintLn("{default}| | | └Firewall:       "+iStr(host.GetDataProcessing()))
				}
				if whatCanSee[10] != "Unknown" {
					WMap["Enviroment"].WPrintLn("{default}| | | └Data Processing:"+iStr(host.GetFirewall()))
				}
				for i := range keysForIC {
					ic := pickObjByID(keysForIC[i]).(IIC)
					whatCanSee := player.GetFieldOfView().KnownData[ic.GetID()]

					if (whatCanSee[0] == "Spotted" && player.GetHost() == host) || ic.GetMarkSet().MarksFrom[player.GetID()] == 4 {
						depth = 3
						WMap["Enviroment"].WPrintLn(treePrefix(i, len(keysForIC), depth))
						WMap["Enviroment"].WPrintLn("{default}| | └Icon: "+ic.GetName())
					}
				}
				for i := range keysForFile {
					file := pickObjByID(keysForFile[i]).(IFile)
					whatCanSee := player.GetFieldOfView().KnownData[file.GetID()]
					if (whatCanSee[0] == "Spotted") || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
						WMap["Enviroment"].WPrintLn("{default}| | └Icon: "+file.GetName()+" ("+file.GetFileName()+")")
						if whatCanSee[12] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							WMap["Enviroment"].WPrintLn("{default}| |   └Encryption Rating:"+iStr(file.GetEncryptionRating()))
						}
						if whatCanSee[3] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							WMap["Enviroment"].WPrintLn("{default}| |   └Data Bomb Rating:"+iStr(file.GetDataBombRating()))
						}
						if whatCanSee[3] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							WMap["Enviroment"].WPrintLn("{default}| |   └File Size:"+iStr(file.GetSize())+" MP")
						}
						if whatCanSee[1] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
							WMap["Enviroment"].WPrintLn("{default}| |   └Last Edited: "+file.GetLastEditDate())
						}

					}
				}

			}

			colorTag = "{default}"
		}

	}*/

	drawLineInWindow("Enviroment")
	drawLineInWindow("Enviroment")
	//keysForHost := getSortedKeysByType("Host")
	drawLine := false
	for i := range keysForHost {
		if drawLine {
			drawLineInWindow("Enviroment")
		}
		host := pickObjByID(keysForHost[i]).(IHost)
		var sampleCode [30]string
		sampleCode[0] = "Spotted" //[0]
		sampleCode[1] = "Unknown" //[1]
		var checkFoW [30]string
		//marks := host.GetMarkSet()
		playerMarks := host.GetMarkSet().MarksFrom[player.GetID()]
		checkFoW = sampleCode
		whatCanSee := player.GetFieldOfView().KnownData[host.GetID()]
		if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrint("Icon: "+host.GetName(), congo.ColorGreen)
			for i := 0; i < playerMarks; i++ {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint("*", congo.ColorGreen)
			}
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Type: Host", congo.ColorGreen)
		}
		if whatCanSee[5] != "Unknown" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Rating: "+strconv.Itoa(host.GetDeviceRating()), congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Rating: Unknown", congo.ColorYellow)
		}
		Att := "Unknown"
		Slz := "Unknown"
		DtPrc := "Unknown"
		Frw := "Unknown"
		if whatCanSee[7] != "Unknown" {
			Att = strconv.Itoa(host.GetAttack())
		}
		if whatCanSee[8] != "Unknown" {
			Slz = strconv.Itoa(host.GetSleaze())
		}
		if whatCanSee[9] != "Unknown" {
			DtPrc = strconv.Itoa(host.GetDataProcessing())
		}
		if whatCanSee[10] != "Unknown" {
			Frw = strconv.Itoa(host.GetFirewall())
		}
		congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("--Host Attribute Array--", congo.ColorGreen)
		//Show Host Attack
		if whatCanSee[7] != "Unknown" {
			Att = strconv.Itoa(host.GetAttack())
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Attack: "+Att, congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Attack: "+Att, congo.ColorYellow)
		}
		//Show Host Sleaze
		if whatCanSee[8] != "Unknown" {
			Att = strconv.Itoa(host.GetSleaze())
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Sleaze: "+Slz, congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Sleaze: "+Slz, congo.ColorYellow)
		}
		//Show Host DataProcessing
		if whatCanSee[9] != "Unknown" {
			Att = strconv.Itoa(host.GetDataProcessing())
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Data Processing: "+DtPrc, congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Data Processing: "+DtPrc, congo.ColorYellow)
		}
		//Show Host Firewall
		if whatCanSee[10] != "Unknown" {
			Att = strconv.Itoa(host.GetFirewall())
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Firewall: "+Frw, congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Firewall: "+Frw, congo.ColorYellow)
		}
		//Show Host Grid
		if whatCanSee[13] != "Unknown" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Grid: "+host.GetGridName(), congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Grid: Unknown", congo.ColorYellow)
		}
		//Show Host Alert
		if host.GetHostAlertStatus() == "No Alert" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Alert Status: "+host.GetHostAlertStatus(), congo.ColorGreen)
		} else if host.GetHostAlertStatus() == "Passive Alert" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Alert Status: "+host.GetHostAlertStatus(), congo.ColorYellow)
		} else if host.GetHostAlertStatus() == "Active Alert" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Alert Status: "+host.GetHostAlertStatus(), congo.ColorRed)
		}
		//Show Host IC Statuses
		if whatCanSee[4] != "Unknown" {
			for i := 0; i < host.GetDeviceRating(); i++ {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint(host.GetICState().icName[i]+": ", congo.ColorGreen)
				if host.GetICState().icStatus[i] == true {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("is Active", congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("is Passive", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
			}
		}
		drawLine = true
		if drawLine {
			drawLineInWindow("Enviroment")
		}
	}

	drawLine = false
	/////////////////////////////////////
	//keysForIC := getSortedKeysByType("IC")
	for i := range keysForIC {
		ic := pickObjByID(keysForIC[i]).(IIC)
		var sampleCode [30]string
		sampleCode[0] = "Spotted" //[0]
		sampleCode[1] = "Unknown" //[1]
		var checkFoW [30]string
		checkFoW = sampleCode
		whatCanSee := player.GetFieldOfView().KnownData[ic.GetID()]
		whatKnowAboutHost := player.GetFieldOfView().KnownData[ic.GetHost().GetID()]
		playerMarks := ic.GetMarkSet().MarksFrom[player.GetID()]
		if ic.GetHost().name == player.GetHost().name && whatCanSee[0] == "Spotted" {
			drawLine = true
			if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint("Icon: "+ic.GetName(), congo.ColorGreen)
				for i := 0; i < playerMarks; i++ {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("*", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)

			}
			/*if whatCanSee[11] != "Unknown" || whatCanSee[11] != "UNSCANNABLE" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Icon Name: "+ic.GetName(), congo.ColorGreen)
			}*/
			icMCM := " _ "
			if whatCanSee[2] != "Unknown" {
				icMCM = strconv.Itoa(ic.GetMatrixCM())
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Matrix Condition Monitor: "+icMCM, congo.ColorGreen)
				//marks := ic.GetMarkSet()
				//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("Marks on IC: %v", marks), congo.ColorYellow)
			} else {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Matrix Condition Monitor: Unknown", congo.ColorYellow)
			}

			if whatCanSee[5] != "Unknown" || whatKnowAboutHost[5] != "Unknown" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Rating: "+strconv.Itoa(ic.GetDeviceRating()), congo.ColorGreen)
			}
			Att := "Unknown"
			Slz := "Unknown"
			DtPrc := "Unknown"
			Frw := "Unknown"
			showAttArray := false
			if whatCanSee[7] != "Unknown" || whatKnowAboutHost[7] != "Unknown" {
				Att = strconv.Itoa(ic.GetAttack())
				showAttArray = true
			}
			if whatCanSee[8] != "Unknown" || whatKnowAboutHost[8] != "Unknown" {
				Slz = strconv.Itoa(ic.GetSleaze())
				showAttArray = true
			}
			if whatCanSee[9] != "Unknown" || whatKnowAboutHost[9] != "Unknown" {
				DtPrc = strconv.Itoa(ic.GetDataProcessing())
				showAttArray = true
			}
			if whatCanSee[10] != "Unknown" || whatKnowAboutHost[10] != "Unknown" {
				Frw = strconv.Itoa(ic.GetFirewall())
				showAttArray = true
			}
			if showAttArray == true {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("---IC Attribute Array---", congo.ColorGreen)

				//Show Host Attack
				if whatCanSee[7] != "Unknown" || whatKnowAboutHost[7] != "Unknown" {
					Att = strconv.Itoa(ic.GetAttack())
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Attack: "+Att, congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Attack: "+Att, congo.ColorYellow)
				}
				//Show Host Sleaze
				if whatCanSee[8] != "Unknown" || whatKnowAboutHost[8] != "Unknown" {
					Att = strconv.Itoa(ic.GetSleaze())
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Sleaze: "+Slz, congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Sleaze: "+Slz, congo.ColorYellow)
				}
				//Show Host DataProcessing
				if whatCanSee[9] != "Unknown" || whatKnowAboutHost[9] != "Unknown" {
					Att = strconv.Itoa(ic.GetDataProcessing())
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Data Processing: "+DtPrc, congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Data Processing: "+DtPrc, congo.ColorYellow)
				}
				//Show Host Firewall
				if whatCanSee[10] != "Unknown" || whatKnowAboutHost[10] != "Unknown" {
					Att = strconv.Itoa(ic.GetFirewall())
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Firewall: "+Frw, congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Firewall: "+Frw, congo.ColorYellow)
				}
				/*if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("IC Name: "+ic.GetName(), congo.ColorGreen)
				}*/
			}
			if drawLine {
				drawLineInWindow("Enviroment")
			}
		}

	}
	///////////////////////////////////
	drawLine = false

	//keysForFile := getSortedKeysByType("File")
	for i := range keysForFile {
		file := pickObjByID(keysForFile[i]).(IFile)
		var sampleCode [30]string
		sampleCode[0] = "Spotted" //[0]
		sampleCode[1] = "Unknown" //[1]
		var checkFoW [30]string
		checkFoW = sampleCode
		whatCanSee := player.GetFieldOfView().KnownData[file.GetID()]
		playerMarks := file.GetMarkSet().MarksFrom[player.GetID()]
		if file.GetHost() == player.GetHost() {
			drawLine = true
			if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint("Icon: "+file.GetType()+" "+strconv.Itoa(file.GetID()), congo.ColorGreen)
				for i := 0; i < playerMarks; i++ {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("*", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Name: "+file.GetFileName(), congo.ColorGreen)

				if whatCanSee[3] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
					b := file.GetDataBombRating()
					if b > 0 {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File DataBomb Rating: "+strconv.Itoa(file.GetDataBombRating()), congo.ColorYellow)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File DataBomb Rating: "+strconv.Itoa(file.GetDataBombRating()), congo.ColorGreen)
					}
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File DataBomb Rating: Unknown", congo.ColorYellow)
				}
				if whatCanSee[12] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
					e := file.GetEncryptionRating()
					if e > 0 {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Encryption Rating: "+strconv.Itoa(file.GetEncryptionRating()), congo.ColorYellow)
					} else {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Encryption Rating: "+strconv.Itoa(file.GetEncryptionRating()), congo.ColorGreen)
					}
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Encryption Rating: Unknown", congo.ColorYellow)
				}
				if whatCanSee[15] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Size: "+strconv.Itoa(file.GetSize())+" Mp", congo.ColorGreen)
				} else {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("File Size: Unknown", congo.ColorGreen)
				}
				if whatCanSee[1] != "Unknown" || file.GetMarkSet().MarksFrom[player.GetID()] == 4 {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Last Edit Time: "+file.GetLastEditDate(), congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("DEBUG: "+file.GetName()+" value = "+strconv.Itoa(file.GetValue()), congo.ColorGreen)
				if file.GetSilentRunningMode() == true {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(file.GetName()+" is silent running", congo.ColorRed)
				}
				if drawLine {
					drawLineInWindow("Enviroment")
				}
			}
		}
	}
	drawLine = false
	keysForAgent := getSortedKeysByType("Agent")
	for i := range keysForAgent {
		if drawLine {
			drawLineInWindow("Enviroment")
		}
		agent := pickObjByID(keysForAgent[i]).(IAgent)
		var sampleCode [30]string
		sampleCode[0] = "Spotted" //[0]
		sampleCode[1] = "Unknown" //[1]
		var checkFoW [30]string
		checkFoW = sampleCode
		whatCanSee := player.GetFieldOfView().KnownData[agent.GetID()]
		playerMarks := agent.GetMarkSet().MarksFrom[player.GetID()]
		if agent.GetHost() == player.GetHost() {
			drawLine = true
			if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint("Icon: "+agent.GetType()+" "+strconv.Itoa(agent.GetID()), congo.ColorGreen)
				for i := 0; i < playerMarks; i++ {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("*", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
				if whatCanSee[18] != "Unknown" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Owner: "+agent.GetOwner().GetName(), congo.ColorGreen)
					if agent.GetOwner().GetName() == player.GetName() {
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Action Protocol: "+agent.GetActionProtocol(), congo.ColorGreen)
					}
				}
				if whatCanSee[5] != "Unknown" || whatCanSee[7] != "Unknown" || whatCanSee[8] != "Unknown" || whatCanSee[9] != "Unknown" || whatCanSee[10] != "Unknown" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Agent Rating: "+strconv.Itoa(agent.GetRating()), congo.ColorGreen)
				}
				if whatCanSee[2] != "Unknown" {
					curMCM := agent.GetMatrixCM()
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("Matrix Condition:", congo.ColorGreen)
					for i := 0; i < curMCM; i++ {
						congo.WindowsMap.ByTitle["Enviroment"].WPrint(" X", congo.ColorGreen)
					}
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
				}
				if whatCanSee[13] != "Unknown" || whatCanSee[13] != "UNSCANNABLE" {
					var position string
					if agent.GetHost() != Matrix {
						position = agent.GetHost().GetName()
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host: "+position, congo.ColorGreen)
					} else {
						position = agent.GetGrid().GetGridName()
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Grid: "+position, congo.ColorGreen)
					}
				}

				if drawLine {
					drawLineInWindow("Enviroment")
				}
			}
		}
	}
	drawLine = false
	keysForSprite := getSortedKeysByType("Sprite")
	for i := range keysForSprite {
		if drawLine {
			drawLineInWindow("Enviroment")
		}
		sprite := pickObjByID(keysForSprite[i]).(ISprite)
		var sampleCode [30]string
		sampleCode[0] = "Spotted" //[0]
		sampleCode[1] = "Unknown" //[1]
		var checkFoW [30]string
		checkFoW = sampleCode
		whatCanSee := player.GetFieldOfView().KnownData[sprite.GetID()]
		playerMarks := sprite.GetMarkSet().MarksFrom[player.GetID()]
		if sprite.GetHost() == player.GetHost() {
			drawLine = true
			if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint("Icon: "+sprite.GetType()+" "+strconv.Itoa(sprite.GetID()), congo.ColorGreen)
				for i := 0; i < playerMarks; i++ {
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("*", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
				if whatCanSee[11] != "Unknown" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Sprite type: "+sprite.GetUDevice()+" (level "+strconv.Itoa(sprite.GetSprLevel())+")", congo.ColorGreen)
				}
				if whatCanSee[18] != "Unknown" {
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Owner: "+sprite.GetOwner().GetName(), congo.ColorGreen)
					if sprite.GetOwner().GetName() == player.GetName() {
						//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Action Protocol: "+agent.GetActionProtocol(), congo.ColorGreen)
					}
				}
				if whatCanSee[5] != "Unknown" || whatCanSee[7] != "Unknown" || whatCanSee[8] != "Unknown" || whatCanSee[9] != "Unknown" || whatCanSee[10] != "Unknown" {
					//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Agent Rating: "+strconv.Itoa(agent.GetRating()), congo.ColorGreen)
				}
				if whatCanSee[2] != "Unknown" {
					curMCM := sprite.GetMatrixCM()
					congo.WindowsMap.ByTitle["Enviroment"].WPrint("Matrix Condition:", congo.ColorGreen)
					congo.WindowsMap.ByTitle["Process"].WPrint("0", congo.ColorGreen)
					for i := 0; i < curMCM; i++ {
						congo.WindowsMap.ByTitle["Process"].WPrint("1", congo.ColorGreen)
						congo.WindowsMap.ByTitle["Enviroment"].WPrint(" X", congo.ColorGreen)
					}
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
				}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Attack: "+strconv.Itoa(sprite.GetAttack()), congo.ColorGreen)
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Sleaze: "+strconv.Itoa(sprite.GetSleaze()), congo.ColorGreen)
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Data Proc: "+strconv.Itoa(sprite.GetDataProcessing()), congo.ColorGreen)
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Firewall: "+strconv.Itoa(sprite.GetFirewall()), congo.ColorGreen)
				if whatCanSee[13] != "Unknown" || whatCanSee[13] != "UNSCANNABLE" {
					var position string
					if sprite.GetHost() != Matrix {
						position = sprite.GetHost().GetName()
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host: "+position, congo.ColorGreen)
					} else {
						position = sprite.GetGrid().GetGridName()
						congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Grid: "+position, congo.ColorGreen)
					}
				}
			}
			if drawLine {
				drawLineInWindow("Enviroment")
			}
		}

	}

}

func refreshProcessWin() {
	congo.WindowsMap.ByTitle["Process"].WDraw()

}

func printLog(data interface{}) {
	WMap["Log"].SetAutoScroll(true)
	//if SourceIcon != nil {
	//	if SourceIcon.(IObj).GetFaction() == player.GetFaction() { //вылетает при выборе хода - что-то связанное с тем что оно берет объект но получает нил
	congo.WindowsMap.ByTitle["Log"].WPrintLn(data)
	hold()
	//	}
	//}
}

type mPrnt struct {
	text  interface{}
	color congo.TColor
}

func printToWindow(winName string, text string) {
	tpt := unTag(text)
	for pt := range tpt {
		congo.WindowsMap.ByTitle[winName].WPrint(tpt[pt].text, tpt[pt].color)
	}
	congo.WindowsMap.ByTitle[winName].WPrintLn("", congo.ColorGreen)
}

func multiPrintTo(winName string, data ...string) {
	for i := range data {
		congo.WindowsMap.ByTitle[winName].WPrint(data[i])
	}
	congo.WindowsMap.ByTitle[winName].WPrintLn("")
}

func multPrint(key, text string) {

	color := congo.ColorGreen
	WMap[key].WPrint(text, color)

}

func testPrint(win string, tVal ...interface{}) {
	col := congo.ColorDefault
	for _, i := range tVal {
		switch val := i.(type) {
		case string, int:
			congo.WindowsMap.ByTitle[win].WPrint(val, col)
		case congo.TColor:
			col = val
		default:
			WMap[win].WPrint("UNSUPPORTED", congo.ColorRed)
		}
	}
}

func unTag(s string) []mPrnt {
	var output []mPrnt
	i1 := strings.Index(s, "{")
	color := congo.ColorGreen
	tag := ""
	for i1 != -1 {
		text := s[:i1]
		output = append(output, mPrnt{text, color})
		s = s[i1:]
		i2 := strings.Index(s, "}")
		if i2 == -1 {
			break
		}
		tag = s[:i2+1]
		s = s[i2+1:]
		switch strings.ToUpper(tag) {
		case "{YELLOW}":
			color = congo.ColorYellow
		case "{RED}":
			color = congo.ColorRed
		case "{GREEN}":
			color = congo.ColorGreen
		case "{DEFAULT}":
			color = congo.ColorDefault
		default:
			color = congo.ColorGreen
		}
		i1 = strings.Index(s, "{")
	}
	output = append(output, mPrnt{s, color})
	//printLog("1087: "+tag, color)
	return output
}

func testPrint2(s string) {
	//start := 0
	i1 := strings.Index(s, "[")
	color := congo.TColor{}
	tag := ""
	for i1 != -1 {
		text := s[:i1]
		testPrint("Log", color, text)
		s = s[i1:]
		i2 := strings.Index(s, "]")
		if i2 == -1 {
			panic(s)
		}
		tag = s[:i2+1]
		s = s[i2+1:]
		//congo.WindowsMap.ByTitle["Log"].WPrint(tag, congo.ColorRed)
		switch strings.ToUpper(tag) {
		case "{yellow}":
			color = congo.ColorYellow
		case "{red}":
			color = congo.ColorRed
		default:
		}
		i1 = strings.Index(s, "[")

	}
	congo.WindowsMap.ByTitle["Log"].WPrint(s, color)
}

/*keysForSprite := getSortedKeysByType("Sprite")
for i := range keysForSprite {
	if drawLine {
		drawLineInWindow("Enviroment")
	}
	agent := pickObjByID(keysForSprite[i]).(ISprite)
	var sampleCode [30]string
	sampleCode[0] = "Spotted" //[0]
	sampleCode[1] = "Unknown" //[1]
	var checkFoW [30]string
	checkFoW = sampleCode
	whatCanSee := player.GetFieldOfView().KnownData[agent.GetID()]
	playerMarks := agent.GetMarkSet().MarksFrom[player.GetID()]
	if agent.GetHost() == player.GetHost() {
		drawLine = true
		if checkFoW[0] == whatCanSee[0] && checkFoW[0] == "Spotted" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrint("Icon: "+agent.GetType()+" "+strconv.Itoa(agent.GetID()), congo.ColorGreen)
			for i := 0; i < playerMarks; i++ {
				congo.WindowsMap.ByTitle["Enviroment"].WPrint("*", congo.ColorGreen)
			}
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
			if whatCanSee[18] != "Unknown" {
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Owner: "+agent.GetOwner().GetName(), congo.ColorGreen)
				if agent.GetOwner().GetName() == player.GetName() {
					//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Action Protocol: "+agent.GetActionProtocol(), congo.ColorGreen)
				}
			}
			if whatCanSee[5] != "Unknown" || whatCanSee[7] != "Unknown" || whatCanSee[8] != "Unknown" || whatCanSee[9] != "Unknown" || whatCanSee[10] != "Unknown" {
				//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Agent Rating: "+strconv.Itoa(agent.GetRating()), congo.ColorGreen)
			}
			if whatCanSee[2] != "Unknown" {
				//curMCM := agent.GetMatrixCM()
				//congo.WindowsMap.ByTitle["Enviroment"].WPrint("Matrix Condition:", congo.ColorGreen)
				//for i := 0; i < curMCM; i++ {
				//	congo.WindowsMap.ByTitle["Enviroment"].WPrint(" X", congo.ColorGreen)
				//}
				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("", congo.ColorGreen)
			}
			if whatCanSee[13] != "Unknown" || whatCanSee[13] != "UNSCANNABLE" {
				var position string
				if agent.GetHost() != Matrix {
					position = agent.GetHost().GetName()
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host: "+position, congo.ColorGreen)
				} else {
					position = agent.GetGrid().GetGridName()
					congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Grid: "+position, congo.ColorGreen)
				}
			}*/

func playerMarks(icon IObj) string {
	var s string
	marks := icon.GetMarkSet().MarksFrom[player.GetID()]
	for i := 0; i < marks; i++ {
		s = s + "*"
	}
	return s
}

func knownStats(icon IIcon) string {
	s := ""
	data := player.GetFieldOfView().KnownData[icon.GetID()]
	/*if data[5] != "Unknown" {
		s = s + "Rating=" + strconv.Itoa(icon.GetDeviceRating()) + " "
	}*/
	if data[7] != "Unknown" {
		s = s + strconv.Itoa(icon.GetAttack()) + "/"
	} else {
		s = s + "?/"
	}

	if data[8] != "Unknown" {
		s = s + strconv.Itoa(icon.GetSleaze()) + "/"
	} else {
		s = s + "?/"
	}
	if data[9] != "Unknown" {
		s = s + strconv.Itoa(icon.GetDataProcessing()) + "/"
	} else {
		s = s + "?/"
	}
	if data[10] != "Unknown" {
		s = s + strconv.Itoa(icon.GetFirewall())
	} else {
		s = s + "?"
	}
	return "ASDF: " + s
}
