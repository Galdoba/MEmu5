package main

import (
	"sort"
	"strconv"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/utils"
)

func refreshPersonaWin() {

	windowList[1].(*congo.TWindow).WClear()
	//player = *objectList[0].(*TPersona)
	if player.GetName() == "Unknown" {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("User Name: <<UNREGISTRATED>>", congo.ColorGreen)
	} else {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("User Name: "+player.GetName(), congo.ColorGreen)
	}
	if player.GetName() != "Unknown" {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Device type: "+player.GetDevice().deviceType, congo.ColorGreen)
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Device model: "+player.GetDevice().model, congo.ColorGreen)
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona User Mode: "+player.GetSimSence(), congo.ColorGreen)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Grid: "+player.GetGrid().GetGridName(), congo.ColorGreen)
	if checkLinkLock(player) == true {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("WARNING: LINK-LOCK DETECTED!", congo.ColorRed)
	}
	if player.GetPhysicalLocation() == true {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("WARNING: Physical location tracked!", congo.ColorYellow)
	}
	//device := player.GetDevice()
	for i := 0; i < congo.WindowsMap.ByTitle["Persona"].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle["Persona"].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Running Programs: ", congo.ColorGreen)
	var rPrgLst []string
	loadedPrgsQty := 0
	for j := range player.GetDevice().GetSoftwareList().programStatus {
		//for j := range player.GetDeviceSoft().programStatus {
		//congo.WindowsMap.ByTitle["Persona"].WPrintLn("Program: "+player.GetDeviceSoft().programName[j]+" | Status: "+player.GetDeviceSoft().programStatus[j], congo.ColorGreen)
		if player.GetDevice().GetSoftwareList().programStatus[j] == "Running" {
			rPrgLst = append(rPrgLst, player.GetDeviceSoft().programName[j])
			loadedPrgsQty++
		}
	}
	if loadedPrgsQty > player.GetDevice().GetMaxRunningPrograms() {
		player.CrashRandomProgram()
	}
	for i := 0; i < player.GetDevice().GetMaxRunningPrograms(); i++ {
		if len(rPrgLst) < player.GetDevice().GetMaxRunningPrograms() {
			rPrgLst = append(rPrgLst, "--EMPTY--")
		}
		congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Slot "+strconv.Itoa(i+1)+": "+rPrgLst[i], congo.ColorGreen)
	}
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn("Total loaded programs: "+strconv.Itoa(loadedPrgsQty), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Attribute Array: ", congo.ColorGreen)
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn(strconv.Itoa(player.GetAttack())+" "+strconv.Itoa(player.GetSleaze())+" "+strconv.Itoa(player.GetDataProcessing())+" "+strconv.Itoa(player.GetFirewall()), congo.ColorGreen)
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona MCM: " + strconv.Itoa(objectList[0].(IPersona).GetMatrixCM()), congo.ColorYellow)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Att  : "+strconv.Itoa(player.GetAttack()), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Slz  : "+strconv.Itoa(player.GetSleaze()), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" DtPr : "+strconv.Itoa(player.GetDataProcessing()), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Fwll : "+strconv.Itoa(player.GetFirewall()), congo.ColorGreen)
	for i := 0; i < congo.WindowsMap.ByTitle["Persona"].GetPrintableWidth(); i++ {
		congo.WindowsMap.ByTitle["Persona"].WPrint("-", congo.ColorDefault)
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("", congo.ColorDefault)
	col := congo.ColorDefault
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Condition Monitor:", congo.ColorGreen)
	congo.WindowsMap.ByTitle["Persona"].WPrint("Matrix  : ", congo.ColorDefault)
	if player.GetMatrixCM() < 1 {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("DEVICE BRICKED!!!", congo.ColorRed)

	} else {
		for i := 0; i < player.GetMatrixCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorGreen)
			col = congo.ColorGreen
			if i < 6 {
				col = congo.ColorYellow
			}
			if i < 3 {
				col = congo.ColorRed
			}

		}
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" ", col)
	congo.WindowsMap.ByTitle["Persona"].WPrint("Stun    : ", congo.ColorYellow)
	if player.GetStunCM() < 1 {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("UNCONSCIOUS!!!", congo.ColorRed)
		//congo.WindowsMap.ByTitle["Log"].WPrint("Emergency logging  terminated...", congo.ColorGreen)
	} else {
		for i := 0; i < player.GetStunCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorYellow)
			col = congo.ColorGreen
			if i < 6 {
				col = congo.ColorYellow
			}
			if i < 3 {
				col = congo.ColorRed
			}
		}
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" ", col)
	congo.WindowsMap.ByTitle["Persona"].WPrint("Physical: ", congo.ColorYellow)

	if player.GetPhysCM() < player.GetBody() {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("DEAD!!!", congo.ColorRed)
	} else if player.GetPhysCM() < 1 {
		col = congo.ColorRed
		congo.WindowsMap.ByTitle["Persona"].WPrint("CRITICAL!!!", congo.ColorRed)
	} else {
		for i := 0; i < player.GetPhysCM(); i++ {
			congo.WindowsMap.ByTitle["Persona"].WPrint("X ", congo.ColorRed)
			col = congo.ColorGreen
			if i < 6 {
				col = congo.ColorYellow
			}
			if i < 3 {
				col = congo.ColorRed
			}
		}
	}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn(" ", col)
	drawLineInWindow("Persona")
	if player.GetSilentRunningMode() {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Running Silent: TRUE", congo.ColorGreen)
	} else {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Running Silent: FALSE", congo.ColorYellow)
	}

	if player.GetInitiative() > 9000 {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona Initiative: null", congo.ColorRed)
	} else {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona Initiative: "+strconv.Itoa(player.GetInitiative()), congo.ColorGreen)
	}
	if player.IsConnected() == false {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Persona disconnected...", congo.ColorRed)

	}
	if player.GetEdge() > 0 {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Edge = "+strconv.Itoa(player.GetEdge())+"/"+strconv.Itoa(player.GetMaxEdge()), congo.ColorGreen)
	}
	if player.GetFullDeffenceFlag() == true {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Full Defence = "+strconv.FormatBool(player.GetFullDeffenceFlag()), congo.ColorYellow)
	}
	drawLineInWindow("Persona")
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("--DEBUG--Total Objects: "+strconv.Itoa(len(ObjByNames)), congo.ColorYellow)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("--DEBUG--waitFlag: "+strconv.FormatBool(player.GetWaitFlag()), congo.ColorYellow)
	//for i := range player.specialization {
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn(player.specialization[i], congo.ColorYellow)
	//}
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("--DEBUG--Free/Simple Actions: "+strconv.Itoa(player.GetFreeActionsCount())+"/"+strconv.Itoa(player.GetSimpleActionsCount()), congo.ColorGreen)

	totalMarks := player.CountMarks()
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Confirmed Marks on Persona: "+strconv.Itoa(totalMarks), congo.ColorYellow)
	congo.WindowsMap.ByTitle["Persona"].WPrintLn("Matrix Search in: "+strconv.Itoa(player.GetSearchResultIn()), congo.ColorYellow)
	for i := range player.GetSearchProcess().SearchIconName {
		name := player.GetSearchProcess().SearchIconName[i]
		objType := player.GetSearchProcess().SearchIconType[i]
		timeTotal := player.GetSearchProcess().SearchTime[i]
		if timeTotal == 0 {
			//player.UpdateSearchProcess()
		}
		timeSpent := player.GetSearchProcess().SpentTurns[i]
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Search: "+objType+" "+name, congo.ColorGreen)
		currentPer := 0
		//turnsPart := 0
		if timeSpent != 0 {
			//turnsPart = (100 / timeTotal)
			//r := player.GetInitiative()/10 + 1
			currentPer = utils.Min(((100 / timeTotal) * (timeSpent)), 100)
		}
		congo.WindowsMap.ByTitle["Persona"].WPrintLn(" Progress: "+strconv.Itoa(currentPer)+"%", congo.ColorGreen)
	}
	for i := range player.GetDownloadProcess().DownloadIconName {
		name := player.GetDownloadProcess().DownloadIconName[i]
		downloaded := player.GetDownloadProcess().DownloadedData[i]
		size := player.GetDownloadProcess().FileSize[i]
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Downloading file: "+name, congo.ColorGreen)
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("Progress: "+strconv.Itoa(downloaded)+" of "+strconv.Itoa(size)+" Mp", congo.ColorGreen)
	}
	drawLineInWindow("Persona")
	//fow := player.GetFieldOfView()
	//congo.WindowsMap.ByTitle["Persona"].WPrintLn(fmt.Sprintf("FoW: %v", fow), congo.ColorYellow)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("0", congo.ColorDefault)

}

func refreshGridWin() {
	if player.GetMatrixCM() < 1 {
		congo.WindowsMap.ByTitle["Persona"].WPrintLn("YOU ARE BRICKED!!!", congo.ColorRed)
		congo.WindowsMap.ByTitle["Grid"].WPrintLn("YOU ARE BRICKED!!!", congo.ColorRed)
	}
	windowList[2].(*congo.TWindow).WClear()
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
		if icon, ok := obj.(IIcon); ok {
			key := icon.GetID()
			if icon.GetType() == objType {
				keys = append(keys, key)
			}
			//keys = append(keys, key)

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

func refreshEnviromentWin() {
	congo.WindowsMap.ByTitle["Enviroment"].WClear()

	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(fmt.Sprintf("ObjByName: %v", ObjByNames), congo.ColorYellow)
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("------------------------------", congo.ColorDefault)

	keys := getSortedKeysByType("Host")
	for k := range keys {
		//	host := pickObjByID(keys[k]).(IIcon)
		//	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(" host is: "+host.GetName()+" "+host.GetType()+" id = "+strconv.Itoa(host.GetID()), congo.ColorDefault)
		for _, obj := range ObjByNames {
			if icon, ok := obj.(IIcon); ok {
				if icon.GetID() == keys[k] {
					//				congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(" Icon is: "+icon.GetName()+" "+icon.GetType()+" id = "+strconv.Itoa(icon.GetID()), congo.ColorDefault)
				} else {
					//	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(" Icon id: "+strconv.Itoa(k)+" not found", congo.ColorDefault)
				}
			}
		}
	}

	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("------------------------------", congo.ColorDefault)

	//var row string
	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Combat Turn №: "+strconv.Itoa(CombatTurn), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Initiative Pass №: "+strconv.Itoa(InitiativePass), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(STime, congo.ColorDefault)
	//drawLineInWindow("Enviroment")
	//for i := range CombatRooster.iconID {
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("id: "+strconv.Itoa(CombatRooster.iconID[i])+"///"+strconv.FormatBool(CombatRooster.iconActed[i])+"///"+strconv.Itoa(CombatRooster.iconInit[i]), congo.ColorDefault)
	//}
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn(generateShadowrunTime(), congo.ColorDefault)
	//congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("------------------------------", congo.ColorDefault)
	keysForHost := getSortedKeysByType("Host")
	drawLine := false
	for i := range keysForHost {
		drawLineInWindow("Enviroment")
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
		if checkFoW[5] == whatCanSee[5] && checkFoW[5] != "Unknown" {
			congo.WindowsMap.ByTitle["Enviroment"].WPrintLn("Host Rating: "+whatCanSee[5], congo.ColorGreen)
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
	keysForIC := getSortedKeysByType("IC")
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
			/*if whatCanSee[11] != "Unknown" {
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

	keysForFile := getSortedKeysByType("File")
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
}

func refreshProcessWin() {
	congo.WindowsMap.ByTitle["Process"].WDraw()

}

func printLog(s string, color congo.TColor) {
	//if SourceIcon != nil {
	//	if SourceIcon.(IObj).GetFaction() == player.GetFaction() { //вылетает при выборе хода - что-то связанное с тем что оно берет объект но получает нил
	congo.WindowsMap.ByTitle["Log"].WPrintLn(s, color)
	hold()
	//	}
	//}
}
