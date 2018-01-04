package main

import (
	"strconv"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/ConGo/utils"
)

//Turn -
var Turn int

//CombatTurn -
var CombatTurn int

//InitiativePass -
var InitiativePass int

//TPassOrder -
type TPassOrder struct {
	iconActed []bool
	iconID    []int
	iconInit  []int
}

//CombatRooster -
var CombatRooster TPassOrder

func addIconToCombatRooster(icon IIcon) {
	CombatRooster.iconID = append(CombatRooster.iconID, icon.GetID())
	CombatRooster.iconInit = append(CombatRooster.iconInit, icon.GetInitiative())
	CombatRooster.iconActed = append(CombatRooster.iconActed, false)
}

func endActionPhase(icon IIcon) bool {
	if CombatRooster.iconActed == nil {
		printLog("Creating CR...", congo.ColorDefault)
		buildInitiativePassOrder()
	}
	for i := range CombatRooster.iconID { //Check combat rooster. mark Icon as Acted and break
		if icon.GetID() == CombatRooster.iconID[i] {
			CombatRooster.iconActed[i] = true
			//return false
			break
		}
	}
	return true
}

func isActionExecutedBy(icon IIcon) bool {
	for i := range CombatRooster.iconID {
		if icon.GetID() == CombatRooster.iconID[i] {
			return CombatRooster.iconActed[i]
		}
	}
	return false
}

func approveDeletion(icon interface{}) bool {
	if ic, ok := icon.(IIC); ok {
		if ic.GetMatrixCM() < 1 {

			return true
		}
	}
	if host, ok := icon.(IHost); ok {
		if host.GetGridName() == "Matrix00064" {

			return true
		}
	}
	/*if file, ok := icon.(IFile); ok {
		if Turn > 1 {
			file.SetInitiative(-1)
			return true
		}
	}*/
	return false
}

func rollInitiativeALL() {
	InitiativePass = 1
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			icon.RollInitiative()
			if icon.GetType() == "File" {
				icon.SetInitiative(-1000)
			}
			if icon.GetType() == "Host" {
				icon.SetInitiative(-1000)
			}
		}
	}
}

func sortMovementOrder(icons []IIcon, maxInit int) []IIcon {
	var sortedOrder []IIcon
	for i := maxInit; i > 0; i-- {
		for j := range icons {
			if i == icons[j].GetInitiative() {
				sortedOrder = append(sortedOrder, icons[j])
			}
		}
	}
	return sortedOrder
}

func trackCombat() {
	congo.WindowsMap.ByTitle["Process"].WClear()
	congo.WindowsMap.ByTitle["Process"].WPrintLn("Info: Combat Turn # "+strconv.Itoa(CombatTurn), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Process"].WPrintLn("Info: Initiative Pass # "+strconv.Itoa(InitiativePass), congo.ColorDefault)
	for i := range CombatRooster.iconID {
		id := CombatRooster.iconID[i]
		in := CombatRooster.iconInit[i]
		status := "false"
		if CombatRooster.iconActed[i] {
			status = "true"
		}
		if pickObjByID(id) != nil {
			congo.WindowsMap.ByTitle["Process"].WPrintLn("Icon: "+pickObjByID(id).GetName()+" / "+strconv.Itoa(in)+" / "+status, congo.ColorDefault)
		}
		//congo.WindowsMap.ByTitle["Process"].WPrintLn("Icon: "+pickObjByID(id).GetName()+" / "+strconv.Itoa(in)+" / "+status, congo.ColorDefault)
	}
}

func buildInitiativeOrder() ([]IIcon, int) {
	var movemetOrder []IIcon
	var maxInit int
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetType() != "File" && icon.GetType() != "Host" {
				maxInit = utils.Max(maxInit, icon.GetInitiative())
				movemetOrder = append(movemetOrder, icon)
			}
		}
	}
	movemetOrder = sortMovementOrder(movemetOrder, maxInit)
	return movemetOrder, maxInit
}

func endInitiativePass() {
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			icon.SetInitiative(icon.GetInitiative() - 10)
			icon.ResetActionsCount()
		}
	}
	buildInitiativePassOrder()
}

func clearICes() {
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if approveDeletion(icon) {
				if icon.GetType() == "IC" {
					host := icon.GetHost()
					host.DeleteIC(icon.(*TIC))
				}
			}
		}
	}
}

func intruderInHost(host IHost) bool {
	if host == Matrix {
		return false
	}
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetHost() == host && icon.GetFaction() != host.GetFaction() {
				return true
			}
		}
	}
	return false
}

func resetCombatRooster() {
	CombatRooster.iconActed = CombatRooster.iconActed[:0]
	CombatRooster.iconID = CombatRooster.iconID[:0]
	CombatRooster.iconInit = CombatRooster.iconInit[:0]
}

func buildInitiativePassOrder() {
	var iconPool []IIcon
	var maxInit int
	resetCombatRooster()
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetType() != "File" && icon.GetType() != "Host" {
				maxInit = utils.Max(maxInit, icon.GetInitiative())
				iconPool = append(iconPool, icon)
			}
		}
	}
	for init := maxInit; init > 0; init-- {
		for j := range iconPool {
			if init == iconPool[j].GetInitiative() {
				addIconToCombatRooster(iconPool[j])
			}
		}
	}
}

func checkTurn() bool {
	//	printLog("Start checkTurn()", congo.ColorDefault)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("/Check END PASS", congo.ColorDefault)
	if itIsEndPass() {
		//	congo.WindowsMap.ByTitle["Log"].WPrintLn("/TRUE", congo.ColorDefault)
		endInitiativePass()
		InitiativePass++
		if itIsEndCombatTurn() {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn(" end Turn by " + icon.GetName(), congo.ColorDefault)
			CombatTurn++
			Turn++
			InitiativePass = 1
			hostAction()
			rollInitiativeALL()
			//	resetCombatRooster()
			buildInitiativePassOrder()
			STime = forwardShadowrunTime()
			if player.GetName() != "Unknown" {
				drawLineInWindow("Log")
				congo.WindowsMap.ByTitle["Log"].WPrintLn("SYSTEM TIME: "+STime, congo.ColorDefault)
			}
			if STime == TimeMarker {
				player.SetWaitFlag(false)
			}
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("//SYSTEM TIME: "+STime, congo.ColorDefault)
			return true
		}

		return true
	}
	congo.WindowsMap.ByTitle["Process"].WClear()
	if player.isOnline() == false {
		refreshEnviromentWin()
		refreshPersonaWin()
		refreshGridWin()
		congo.Flush()
		player.SetInitiative(999999)
		//os.Exit(1)
	}
	if CombatRooster.iconActed == nil {
		//addIconToCombatRooster(player)
		//	printLog("Creating CR...", congo.ColorDefault)
		buildInitiativePassOrder()
	}
	//trackCombat()
	//maxInit := 0
	//var movemetOrder []IIcon
	var icon IIcon
	mActionName := "ICWAIT"
	//outIndex := 0
	clearICes()
	//	_, maxInit = buildInitiativeOrder()
	refreshEnviromentWin()
	for i := range CombatRooster.iconActed {
		if CombatRooster.iconActed[i] {
			//printLog("3", congo.ColorDefault)
			continue
		} else {
			if tryIcon, ok := pickObjByID(CombatRooster.iconID[i]).(IIcon); ok {
				icon = tryIcon
				break
			}
			//icon = pickObjByID(CombatRooster.iconID[i]).(IIcon)
			//break
		}
		//panic(1) //не заходит
	}
	if icon == nil {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("CR is empty/all true = end pass", congo.ColorDefault)
		endInitiativePass()
	}
	if icon != nil {
		if icon.GetID() == player.GetID() {
			if player.GetWaitFlag() {
				SourceIcon = pickObjByID(player.GetID())
				doAction("WAIT")
			}
			return true
		}
		if ic, ok := icon.(IIC); ok {
			if ic.GetInitiative() > 0 {
				//выбираем источник
				SourceIcon = ic
				//выбираем цель
				TargetIcon = player
				//выбираем действие
				mActionName = icDecide(ic) //нужен целеуказывающий механизм для айсов

				/////////////////////////////////////////////////////////////////////
				//printLog("--point DoAction", congo.ColorDefault)
				//printLog(ic.GetName()+" / "+mActionName, congo.ColorDefault)
				doAction(mActionName)
				//printLog("Go ednActionforIC()", congo.ColorDefault)
				endActionPhase(ic)
				//printLog("end checkTurn() 2", congo.ColorDefault)
				return false
			}

		}
		if agent, ok := icon.(IAgent); ok {
			if agent.GetInitiative() > 0 {
				//выбираем источник
				SourceIcon = agent
				actStr, tarStr := agent.RunActionProtocol()
				actStr = formatString(actStr)
				//tarStr = formatString(tarStr)
				//выбираем цель
				TargetIcon = ObjByNames[tarStr]
				//выбираем действие
				printLog("Agent Protocol: "+agent.GetActionProtocol(), congo.ColorDefault)
				mActionName = actStr
				/////////////////////////////////////////////////////////////////////
				command = formatString(agent.GetName() + ">" + mActionName + ">" + tarStr)
				printLog(command, congo.ColorDefault)
				doAction(mActionName)
				//printLog("Agent: "+mActionName+", target: "+tarStr, congo.ColorDefault)
				endActionPhase(agent)
				//printLog("end checkTurn() 2", congo.ColorDefault)
				return false
			}
		}
	}
	return true
}

func itIsEndPass() bool {

	//printLog("End Pass?", congo.ColorYellow)
	end := true
	for i := range CombatRooster.iconActed {
		if CombatRooster.iconActed[i] { //} && pickObjByID(CombatRooster.iconID[i]) != nil {
			//icon acted
			end = true
		} else {
			if pickObjByID(CombatRooster.iconID[i]) != nil {
				//icon not acted
				end = false
				break
			}
		}
		end = end && CombatRooster.iconActed[i] //if all icons acted end = true
	}
	return end
}

func itIsEndCombatTurn() bool {
	_, maxInit := buildInitiativeOrder()
	if maxInit < 1 {
		return true
	}
	return false
}

func icDecide(ic IIC) string {
	if player.isOnline() == true {
		mActionName := "ICWAIT"
		attackPotential := false
		intruderPresence := false

		SourceIcon = ic
		host := ic.GetHost()
		data := ic.GetFieldOfView() //.KnownData[obj.GetID()]
		if host.alert == "Passive Alert" || host.alert == "Active Alert" {
			for _, obj := range ObjByNames {
				if intruder, ok := obj.(IIcon); ok {
					marks := intruder.GetMarkSet() //смотрим какие марки есть на вторженце
					for id, qty := range marks.MarksFrom {
						if id == host.GetID() && qty > 0 && qty != 4 && ic.GetFaction() != intruder.GetFaction() {
							if data.KnownData[intruder.GetID()][0] != "Spotted" {
								ic.ChangeFOWParametr(ic.GetID(), 0, "Spotted") // to change 1 FOWParametr use : (int id, key, string newValue)
							}
							attackPotential = true
						}
					}
					if intruder.GetFaction() != host.GetFaction() && intruder.GetHost() == host { // && ic.GetName() == "Patrol IC" {
						intruderPresence = true
					}
					if intruderPresence == true {
						data = ic.GetFieldOfView()
						for id, value := range data.KnownData {
							if id == intruder.GetID() && value[0] == "Spotted" {
								attackPotential = true
							}
						}

					}
				}
			}
		}
		if intruderPresence == true && ic.GetName() == "Patrol IC" {
			mActionName = "EXECUTE_SCAN"
		}
		if attackPotential {
			if ic.(*TIC).icChoseTarget() != nil {
				switch ic.GetName() {
				case "Patrol IC":
					mActionName = "PATROL_IC_ACTION"
				case "Acid IC":
					mActionName = "ACID_IC_ACTION"
				case "Binder IC":
					mActionName = "BINDER_IC_ACTION"
				case "Jammer IC":
					mActionName = "JAMMER_IC_ACTION"
				case "Marker IC":
					mActionName = "MARKER_IC_ACTION"
				case "Killer IC":
					mActionName = "KILLER_IC_ACTION"
				case "Sparky IC":
					mActionName = "SPARKY_IC_ACTION"
				case "Tar Baby IC":
					mActionName = "TAR_BABY_IC_ACTION"
				case "Black IC":
					mActionName = "BLACK_IC_ACTION"
				case "Blaster IC":
					mActionName = "BLASTER_IC_ACTION"
				case "Probe IC":
					mActionName = "PROBE_IC_ACTION"
				case "Scramble IC":
					mActionName = "SCRAMBLE_IC_ACTION"
				case "Catapult IC":
					mActionName = "CATAPULT_IC_ACTION"
				case "Shoker IC":
					mActionName = "SHOKER_IC_ACTION"
				case "Track IC":
					mActionName = "TRACK_IC_ACTION"
				case "Bloodhound IC":
					if ic.(*TIC).GetLastTargetName() == "" {
						mActionName = "BLOODHOUND_IC_SCAN"
					} else {
						mActionName = "BLOODHOUND_IC_ACTION"
					}
				case "Crash IC":
					mActionName = "CRASH_IC_ACTION"
				default:
					if mActionName == "EXECUTE_SCAN" {
						mActionName = "EXECUTE_SCAN"
					} else {
						mActionName = "ICWAIT"
					}
				}
			}
		} else {
			ic.(*TIC).TakeFOWfromHost()
		}
		return mActionName
		//doAction(mActionName)
	}
	return "ICWAIT"
}

func getICAttack() string {
	if src, ok := SourceIcon.(*TIC); ok {
		switch src.GetName() {
		case "Patrol IC":
			return "PATROL_IC_ACTION"
		default:
			return "unknown IC action"
		}

		/*congo.WindowsMap.ByTitle["Log"].WPrintLn("ACTION + "+src.GetName(), congo.ColorDefault)
		return src.GetName()*/
	}
	return "none"
}

func (ic *TIC) icChoseTarget() interface{} {
	var potentialTargets []int
	host := ic.GetHost()
	ic.SetFieldOfView(host.GetFieldOfView())
	for _, obj := range ObjByNames {
		if trg, ok := obj.(IIcon); ok {

			//	}
			//}

			//for i := range objectList {
			//	obj := objectList[i]
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("try = "+strconv.Itoa(obj.(IObj).GetID()), congo.ColorRed)
			//	if trg, ok := obj.(IIcon); ok {
			if trg.GetFaction() != ic.GetFaction() && trg.GetHost().name == ic.GetHost().name {
				canSee := ic.canSee.KnownData[trg.GetID()]
				//congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" is VALID", congo.ColorRed)
				//congo.WindowsMap.ByTitle["Log"].WPrintLn(canSee[0], congo.ColorRed)
				if canSee[0] == "Spotted" {
					//congo.WindowsMap.ByTitle["Log"].WPrintLn(ic.GetName()+" can see "+trg.GetName(), congo.ColorYellow)
					//congo.WindowsMap.ByTitle["Log"].WPrintLn(" len(potentialTargets) = "+strconv.Itoa(len(potentialTargets)), congo.ColorRed)
					potentialTargets = append(potentialTargets, trg.GetID())
					//congo.WindowsMap.ByTitle["Log"].WPrintLn(" len(potentialTargets) = "+strconv.Itoa(len(potentialTargets)), congo.ColorRed)
					shuffleInt(potentialTargets)
				} else {
					//congo.WindowsMap.ByTitle["Log"].WPrintLn("Host can see:: "+host.canSee.KnownData[trg.GetID()][0], congo.ColorDefault)
					if host.canSee.KnownData[trg.GetID()][0] == "Spotted" {
						canSee[0] = "Spotted"
						//	congo.WindowsMap.ByTitle["Log"].WPrintLn("************* "+host.canSee.KnownData[trg.GetID()][0], congo.ColorDefault)
						ic.SetFieldOfView(host.GetFieldOfView())
					}
				}

			}
		}
		shuffleInt(potentialTargets)
	}
	shuffleInt(potentialTargets)
	for _, obj := range ObjByNames {
		//obj := objectList[j]
		if len(potentialTargets) == 0 {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn(" len(potentialTargets) = 0 ", congo.ColorRed)
			return nil
		}
		if obj.(IObj).GetID() == potentialTargets[0] {
			trg := obj
			TargetIcon = trg
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("return potentialTargets[0]", congo.ColorRed)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("id = "+strconv.Itoa(TargetIcon.(IObj).GetID()), congo.ColorRed)
			return trg
		}
	}

	//congo.WindowsMap.ByTitle["Log"].WPrintLn("End icChoseTargets()", congo.ColorRed)
	return nil
}

func rollInitiative() {
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			icon.RollInitiative()
		}
	}
}

func hostAction() {
	keysForPersona := getSortedKeysByType("Persona")
	for i := range keysForPersona {
		persona := pickObjByID(keysForPersona[i]).(IPersona)
		persona.SetFullDeffenceFlag(false)
		persona.UpdateSearchProcess()
		persona.UpdateDownloadProcess()
		//persona.RollInitiative()
	}

	keysForHost := getSortedKeysByType("Host")
	for i := range keysForHost {
		host := pickObjByID(keysForHost[i]).(*THost)
		host.GatherMarks()
		////////////////////////////////////////обнуляем к фолсу айсы которых по факту нет - костыль

		/////////////////////////////////////////конец костыля
		if host.checkAlert() == "Active Alert" {
			host.LoadNextIC()
		} else {
			if host.PickPatrolIC() != nil {

				patrolIC := host.PickPatrolIC()
				patrolIC.SetActionReady(patrolIC.GetActionReady() - 1)
				//printLog("ActionReady = "+strconv.Itoa(patrolIC.GetActionReady()), congo.ColorDefault)
				//patrolIC.actionReady = patrolIC.actionReady - 1
				if host.alert == "Passive Alert" || host.alert == "Active Alert" {
					patrolIC.SetActionReady(-1)
					//patrolIC.actionReady = -1
				}
				if patrolIC.GetActionReady() == 0 {
					if player.GetHost() == patrolIC.GetHost() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Host: Rotine scan initiated...", congo.ColorYellow)
					}
					SourceIcon = patrolIC
					//TargetIcon = "someone"
					mActionName := "EXECUTE_SCAN"
					doAction(mActionName)

					patrolIC.SetActionReady(calculatePartolScan(patrolIC.GetDeviceRating()))
				}
			}
		}
	}
}

func (host *THost) checkAlert() string {
	return host.GetHostAlertStatus()
}
