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

func endActionPhase(icon IIcon) {
	for i := range CombatRooster.iconID {
		if icon.GetID() == CombatRooster.iconID[i] {
			CombatRooster.iconActed[i] = true
			break
		}
	}
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
	/*if persona, ok := icon.(IPersona); ok {
		if persona.GetMatrixCM() < 1 {

			return true
		}
	}*/
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
		//printLog(strconv.Itoa(i), congo.ColorDefault)
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
	congo.WindowsMap.ByTitle["Process"].WPrintLn("Info: player Init: "+strconv.Itoa(player.GetInitiative()), congo.ColorDefault)
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

func declareAction(icon IIcon) {
	printLog(icon.GetName()+" is acting", congo.ColorDefault)
	if icon.IsPlayer() {
		trackCombat()
		waitEvent := false
		i := 0
		for !waitEvent {
			draw()
			printLog("command: "+command, congo.ColorRed)
			printLog(strconv.Itoa(i), congo.ColorRed)
			i++
			if i == 10 {
				waitEvent = true
			}
		}
		waitEvent = true

	}
	//trackCombat()
}

func startCombatTurn() {
	CombatTurn++
	var maxInit int
	var movemetOrder []IIcon
	rollInitiativeALL()
	movemetOrder, maxInit = buildInitiativeOrder()
	for maxInit > 0 {
		for i := range movemetOrder {
			if movemetOrder[i].GetInitiative() > 0 {
				checkTurn()
				declareAction(movemetOrder[i])
				endAction()
			}
			//declareAction(movemetOrder[i])
		}
		endInitiativePass()
		movemetOrder, maxInit = buildInitiativeOrder()
	}
}

func endInitiativePass() {
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			icon.SetInitiative(icon.GetInitiative() - 10)
		}
	}
	InitiativePass++
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
				printLog(iconPool[j].GetName()+" // "+strconv.Itoa(iconPool[j].GetInitiative()), congo.ColorDefault)
			}
		}
	}
}

func checkTurn() bool {
	if CombatRooster.iconActed == nil {
		//addIconToCombatRooster(player)
		printLog("7", congo.ColorDefault)
		buildInitiativePassOrder()
	}
	trackCombat()
	maxInit := 0
	//var movemetOrder []IIcon
	var icon IIcon
	mActionName := "ICWAIT"
	//outIndex := 0
	clearICes()
	_, maxInit = buildInitiativeOrder()
	refreshEnviromentWin()
	for i := range CombatRooster.iconActed {
		if CombatRooster.iconActed[i] {
			printLog("3", congo.ColorDefault)
			continue
		} else {
			icon = pickObjByID(CombatRooster.iconID[i]).(IIcon)
			break
		}
		panic(1) //не заходит
	}
	if icon != nil {
		congo.WindowsMap.ByTitle["Process"].WPrintLn(icon.GetName()+"'s Turn", congo.ColorYellow)
		if icon.GetID() == player.GetID() {
			return true
		}
		if ic, ok := icon.(IIC); ok {
			if ic.GetInitiative() == maxInit && maxInit > 0 {
				//выбираем источник
				SourceIcon = ic
				congo.WindowsMap.ByTitle["Log"].WPrintLn(ic.GetName()+" Init: "+strconv.Itoa(ic.GetInitiative()), congo.ColorDefault)
				//выбираем действие
				mActionName = icDecide(ic) //нужен целеуказывающий механизм для айсов
				//break                        //continue
				congo.WindowsMap.ByTitle["Log"].WPrintLn(ic.GetName()+" decided "+mActionName, congo.ColorDefault)
				//выбираем цель
				TargetIcon = player
				doAction(mActionName)
				endActionPhase(ic)
				printLog("1", congo.ColorDefault)
				return false
			}

		}

	} else {
		printLog("End Pass?", congo.ColorYellow)
		endInitiativePass()
	}
	//printLog(icon.GetName()+"'s Turn", congo.ColorYellow)

	if maxInit < 1 {
		rollInitiative()
		/*for _, o := range ObjByNames {
			if icon, ok := o.(IIcon); ok {
				if icon.GetInitiative() > 0 {
					//	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" init: "+strconv.Itoa(icon.GetInitiative()), congo.ColorYellow)
				}
			}
		}*/
		hostAction()
		STime = forwardShadowrunTime()
		congo.WindowsMap.ByTitle["Log"].WPrintLn("//SYSTEM TIME: "+STime, congo.ColorDefault)
		Turn++

	}
	printLog("2", congo.ColorDefault)
	return true
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

func checkTurn1() {
	//trackCombat()
	maxInit := 0
	turnGo := true
	lap := 0
	//outIndex := 0
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

	for turnGo == true && player.isOnline() == true {
		maxInit = 0
		var movemetOrder []IIcon
		for _, obj := range ObjByNames {
			if icon, ok := obj.(IIcon); ok {
				if icon.GetType() == "File" {
					icon.SetInitiative(-1)
				}
				maxInit = utils.Max(maxInit, icon.GetInitiative())
				movemetOrder = append(movemetOrder, icon)
				printLog("append"+icon.GetName(), congo.ColorGreen)
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("*/*/*/*/*/**/*/*//*/**/", congo.ColorDefault)
				/*congo.WindowsMap.ByTitle["Log"].WPrintLn("movementOrder:=", congo.ColorDefault)
				for i := range movemetOrder {
					congo.WindowsMap.ByTitle["Log"].WPrintLn(strconv.Itoa(i)+" / "+movemetOrder[i].GetName()+" / Init = "+strconv.Itoa(movemetOrder[i].GetInitiative()), congo.ColorDefault)
				}*/
			}

		}
		sortMovementOrder(movemetOrder, maxInit)

		refreshEnviromentWin()

		/*	for j := range objectList {
			//	for _, obj := range ObjByNames {
			//	if icon, ok := obj.(IIcon); ok {
			if obj, ok := objectList[j].(IIcon); ok {
				if obj.GetType() == "File" {
					objectList[j].(IIcon).SetInitiative(-1)
					//	congo.WindowsMap.ByTitle["Log"].WPrintLn(obj.GetName()+" is file...", congo.ColorDefault)
				} else {
					//congo.WindowsMap.ByTitle["Log"].WPrintLn(obj.GetType()+" Check Initiative...", congo.ColorDefault)
				}
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("object: "+obj.GetName()+" have "+strconv.Itoa(obj.GetInitiative())+" initiative", congo.ColorDefault)

				maxInit = utils.Max(maxInit, obj.GetInitiative())
				/*	if obj.GetName() != player.GetName() {
					objectList[i].(IIcon).SetInitiative(-1)

				}*/
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("object: " + obj.GetName() + " have " + strconv.Itoa(obj.GetInitiative()) + " initiative", congo.ColorDefault)
		//	}

		//}
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("MaxINit "+strconv.Itoa(maxInit), congo.ColorDefault)
		//refreshEnviromentWin()
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("Turn "+strconv.Itoa(Turn)+" order", congo.ColorDefault)
		for i := range movemetOrder {
			if icon, ok := movemetOrder[i].(IIcon); ok {
				if icon.GetType() != "File" {
					//	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Initiative = "+strconv.Itoa(icon.GetInitiative()), congo.ColorDefault)
				}

			}
		}

		for i := range movemetOrder {
			if icon, ok := movemetOrder[i].(IIC); ok {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Initiative = "+strconv.Itoa(icon.GetInitiative()), congo.ColorDefault)
				icon.SetInitiative(icon.GetInitiative() - 10)
			}

			if len(movemetOrder)-1 < i { //костыль от Index Out of Range
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Force brake", congo.ColorDefault)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("######################################", congo.ColorDefault)
				break
			}
			if obj, ok := movemetOrder[i].(IIcon); ok {
				if persona, ok := movemetOrder[i].(IPersona); ok {
					persona.CheckConvergence()
					//congo.WindowsMap.ByTitle["Log"].WPrintLn("persona turn", congo.ColorDefault)
				}
				if obj.GetInitiative() == maxInit && maxInit > 0 {
					/*	congo.WindowsMap.ByTitle["Log"].WPrintLn("try: "+obj.GetName(), congo.ColorYellow)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Obj init: "+strconv.Itoa(obj.GetInitiative()), congo.ColorYellow)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Obj type: "+obj.GetType(), congo.ColorYellow)*/
					if obj.IsPlayer() == true {
						/*congo.WindowsMap.ByTitle["Log"].WPrintLn("Player's Turn: ", congo.ColorYellow)
						congo.WindowsMap.ByTitle["Process"].WPrint(".", congo.ColorGreen)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("lap: "+strconv.Itoa(lap), congo.ColorYellow)*/

						if lap > 10 {
							turnGo = false
						}
						turnGo = false
					}
					if obj.GetType() == "IC" && player.isOnline() == true {

						mActionName := "ICWAIT"
						attackPotential := false
						intruderPresence := false
						ic := obj
						SourceIcon = ic
						host := obj.(IIC).GetHost()
						data := ic.GetFieldOfView() //.KnownData[obj.GetID()]
						if host.alert == "Passive Alert" || host.alert == "Active Alert" {
							for _, obj := range ObjByNames {
								if intruder, ok := obj.(IIcon); ok {
									marks := intruder.GetMarkSet() //смотрим какие марки есть на вторженце
									for id, qty := range marks.MarksFrom {
										if id == host.GetID() && qty > 0 && qty != 4 && ic.GetFaction() != obj.GetFaction() {
											if data.KnownData[intruder.GetID()][0] != "Spotted" {
												ic.ChangeFOWParametr(obj.GetID(), 0, "Spotted") // to change 1 FOWParametr use : (int id, key, string newValue)
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
						doAction(mActionName)
					}
				}
			}
			//persona in MOVORD
			//	movemetOrder = movemetOrder[:i]
			//panic(1)
		}

		refreshPersonaWin()
		if maxInit <= 0 {
			hostAction()
			rollInitiative()
			STime = forwardShadowrunTime()
			printLog("System time: "+STime, congo.ColorGreen)
			Turn++
			turnGo = false
		}
		if lap > 10 {
			turnGo = false
		}
		if player.isOnline() == false {
			refreshEnviromentWin()
			refreshPersonaWin()
			refreshGridWin()
			congo.Flush()
			player.SetInitiative(999999)
			//os.Exit(1)
		}
	}
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
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
		persona.RollInitiative()
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

/*
func checkTurn() bool {
	//trackCombat()
	maxInit := 0
	var movemetOrder []IIcon
	turnGo := true
	lap := 0
	autoWait := false
	mActionName := "ICWAIT"
	//outIndex := 0
	clearICes()
	for turnGo == true && player.isOnline() == true {

		movemetOrder, maxInit = buildInitiativeOrder()
		refreshEnviromentWin()

		for i := range movemetOrder {
			if persona, ok := movemetOrder[i].(IPersona); ok {
				if persona.GetID() == player.GetID() {
					return true
				}
			}
			if icon, ok := movemetOrder[i].(IIC); ok {
				if icon.GetInitiative() == maxInit && maxInit > 0 {
					//выбираем источник
					SourceIcon = icon
					congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" Init: "+strconv.Itoa(icon.GetInitiative()), congo.ColorDefault)
					//выбираем действие
					mActionName = icDecide(icon) //нужен целеуказывающий механизм для айсов
					//break                        //continue
					congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" decided "+mActionName, congo.ColorDefault)
					//выбираем цель
					TargetIcon = player
					doAction(mActionName)
					return false
				}
				if maxInit < 1 {
					rollInitiative()
					for _, o := range ObjByNames {
						if icon, ok := o.(IIcon); ok {
							if icon.GetInitiative() > 0 {
								//	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" init: "+strconv.Itoa(icon.GetInitiative()), congo.ColorYellow)
							}
						}
					}

					hostAction()
					STime = forwardShadowrunTime()

					//drawLineInWindow("Log")
					congo.WindowsMap.ByTitle["Log"].WPrintLn("//SYSTEM TIME: "+STime, congo.ColorDefault)
					//drawLineInWindow("Log")
					Turn++

					turnGo = false
				}
				break
			}
			if len(movemetOrder)-1 < i { //костыль от Index Out of Range
				congo.WindowsMap.ByTitle["Log"].WPrintLn("--DEBUG__ERROR:  Force brake", congo.ColorDefault)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("######################################", congo.ColorDefault)
				break
			}
			if obj, ok := movemetOrder[i].(IIcon); ok {
				if persona, ok := movemetOrder[i].(IPersona); ok {
					persona.CheckConvergence()
				}
				if obj.GetInitiative() == maxInit && maxInit > 0 {
					if obj.IsPlayer() == true {

						if lap > 10 {
							turnGo = false
						}
						turnGo = false
					}
					if obj.GetType() == "IC" && player.isOnline() == true {
						//mActionName := "ICWAIT"
						attackPotential := false
						intruderPresence := false
						ic := obj
						SourceIcon = ic
						host := obj.(IIC).GetHost()
						data := ic.GetFieldOfView() //.KnownData[obj.GetID()]
						if host.alert == "Passive Alert" || host.alert == "Active Alert" {
							for _, obj := range ObjByNames {
								if intruder, ok := obj.(IPersona); ok {

									//congo.WindowsMap.ByTitle["Log"].WPrintLn(ic.GetName()+" try "+intruder.GetName(), congo.ColorDefault)

									marks := intruder.GetMarkSet() //смотрим какие марки есть на вторженце
									for id, qty := range marks.MarksFrom {
										if id == host.GetID() && qty > 0 && qty != 4 && ic.GetFaction() != intruder.GetFaction() {
											if data.KnownData[intruder.GetID()][0] != "Spotted" {
												ic.ChangeFOWParametr(obj.GetID(), 0, "Spotted") // to change 1 FOWParametr use : (int id, key, string newValue)
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
						if mActionName != "ICWAIT" && ic.GetHost() == player.GetHost() {
							//printLog("Attack of "+ic.GetName()+" detected...", congo.ColorYellow)
							//printLog("Action: "+mActionName, congo.ColorYellow)
						}
						doAction(mActionName)
					}
				}
			}
			//	movemetOrder = movemetOrder[:i]
		}
		refreshPersonaWin()
		if maxInit < 1 {
			rollInitiative()
			for _, o := range ObjByNames {
				if icon, ok := o.(IIcon); ok {
					if icon.GetInitiative() > 0 {
						//	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" init: "+strconv.Itoa(icon.GetInitiative()), congo.ColorYellow)
					}
				}
			}

			hostAction()
			STime = forwardShadowrunTime()

			//drawLineInWindow("Log")
			congo.WindowsMap.ByTitle["Log"].WPrintLn("SYSTEM TIME: "+STime, congo.ColorDefault)
			//drawLineInWindow("Log")
			Turn++

			turnGo = false
		}
		if lap > 1000 {
			turnGo = false
		}
		if player.isOnline() == false {
			refreshEnviromentWin()
			refreshPersonaWin()
			refreshGridWin()
			congo.Flush()
			player.SetInitiative(999999)
		} else {
			//SourceIcon = pickObjByID(player.GetID())
			//doAction("WAIT")
		}
		lap++
	}
	if autoWait {
		if player.GetWaitFlag() {
			autoWait = false
			SourceIcon = pickObjByID(player.GetID())
			//doAction("WAIT")
		}
	}
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
	return true
}*/
