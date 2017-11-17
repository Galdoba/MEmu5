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
	printLog("Info: Combat Turn # "+strconv.Itoa(CombatTurn), congo.ColorDefault)
	printLog("Info: Initiative Pass # "+strconv.Itoa(InitiativePass), congo.ColorDefault)
	printLog("Info: player Init: "+strconv.Itoa(player.GetInitiative()), congo.ColorDefault)
}

func buildInitiativeOrder() ([]IIcon, int) {
	var movemetOrder []IIcon
	var maxInit int
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			maxInit = utils.Max(maxInit, icon.GetInitiative())
			movemetOrder = append(movemetOrder, icon)
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
}

func checkTurn() {
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
			}
		}
		sortMovementOrder(movemetOrder, maxInit)
		refreshEnviromentWin()
		for i := range movemetOrder {
			if icon, ok := movemetOrder[i].(IIcon); ok {
				if icon.GetType() != "File" {
					//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Initiative = "+strconv.Itoa(icon.GetInitiative()), congo.ColorDefault)
				}
			}
		}
		for i := range movemetOrder {
			if icon, ok := movemetOrder[i].(IIC); ok {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("Checking = "+icon.GetName()+": Initiative = "+strconv.Itoa(icon.GetInitiative()), congo.ColorGreen)
				if icon.GetInitiative() == maxInit && maxInit > 0 {
					//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Initiative = "+strconv.Itoa(icon.GetInitiative()), congo.ColorDefault)
					icDecide(icon) //нужен целеуказывающий механизм для айсов - перевести из Листа в Мап
					//continue
				}

			}
			if len(movemetOrder)-1 < i { //костыль от Index Out of Range
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Force brake", congo.ColorDefault)
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
						mActionName := "ICWAIT"
						attackPotential := false
						intruderPresence := false
						ic := obj
						SourceIcon = ic
						host := obj.(IIC).GetHost()
						data := ic.GetFieldOfView() //.KnownData[obj.GetID()]
						if host.alert == "Passive Alert" || host.alert == "Active Alert" {
							for _, obj := range ObjByNames {
								if intruder, ok := obj.(IPersona); ok {

									//		}
									//	}

									//	for j := range objectList {
									//		if intruder, ok := objectList[j].(IIcon); ok {
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
							printLog("Attack of "+ic.GetName()+" detected...", congo.ColorYellow)
						}
						doAction(mActionName)
					}
				}
			}
			//	movemetOrder = movemetOrder[:i]
		}
		refreshPersonaWin()
		if maxInit <= 0 {
			rollInitiative()
			hostAction()
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
		}
	}
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
}

func icDecide(ic IIC) {
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

			/*for j := range objectList {
				if intruder, ok := objectList[j].(IIcon); ok {
					marks := intruder.GetMarkSet() //смотрим какие марки есть на вторженце
					for id, qty := range marks.MarksFrom {
						if id == host.GetID() && qty > 0 && qty != 4 && ic.GetFaction() != objectList[j].(IIcon).GetFaction() {
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
			}*/
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

func checkTurn0() {
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
							for j := range objectList {
								if intruder, ok := objectList[j].(IIcon); ok {
									marks := intruder.GetMarkSet() //смотрим какие марки есть на вторженце
									for id, qty := range marks.MarksFrom {
										if id == host.GetID() && qty > 0 && qty != 4 && ic.GetFaction() != objectList[j].(IIcon).GetFaction() {
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
			//	movemetOrder = movemetOrder[:i]
			//panic(1)
		}

		refreshPersonaWin()
		if maxInit <= 0 {
			rollInitiative()
			hostAction()
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
	for j := range objectList {
		obj := objectList[j]
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

func icAct(maxInit int) {
	for i := range objectList {
		if obj, ok := objectList[i].(IIcon); ok {
			if obj.GetInitiative() == maxInit {
				if obj.GetType() == "IC" {
					/*SourceIcon = objectList[i].(IIC)
					TargetIcon = obj.GetOwner() //temp
					doAction("ICACTION")*/

					congo.WindowsMap.ByTitle["Log"].WPrintLn("Simulate IC Action: ", congo.ColorRed)
					congo.WindowsMap.ByTitle["Log"].WPrintLn("IC old Initiative: "+strconv.Itoa(obj.GetInitiative()), congo.ColorRed)
					objectList[i].(IIcon).SetInitiative(obj.GetInitiative() - 10)
					congo.WindowsMap.ByTitle["Log"].WPrintLn("IC new Initiative: "+strconv.Itoa(obj.GetInitiative()), congo.ColorRed)
					congo.WindowsMap.ByTitle["Log"].WPrintLn(obj.GetName()+">IC_ACTION>GALDOBA ", congo.ColorDefault)
				}
			}
		}
	}
}

func rollInitiative() {

	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {

			sms := icon.GetSimSence()
			dice := 0
			switch sms {
			case "AR":
				dice = 1
			case "COLD-SIM":
				dice = 3
			case "HOT-SIM":
				dice = 4
			default:
				dice = 4
				//panic(0)
			}
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("object: "+obj.GetName()+" rolling initiative", congo.ColorYellow)
			icon.SetInitiative(icon.GetDataProcessing() + icon.GetDeviceRating() + xd6Test(dice))
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("object: " + obj.GetName() + " have " + strconv.Itoa(obj.GetInitiative()) + " initiative", congo.ColorDefault)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("object: " + player.name + " have " + strconv.Itoa(player.GetInitiative()) + " initiative", congo.ColorDefault)
		}
	}
	/*for i := range objectList {
		if obj, ok := objectList[i].(IPersona); ok {
			sms := obj.GetSimSence()
			dice := 0
			switch sms {
			case "AR":
				dice = 1
			case "COLD-SIM VR":
				dice = 3
			case "HOT-SIM":
				dice = 4
			default:
				dice = 4
				//panic(0)
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn("End Round "+strconv.Itoa(Turn), congo.ColorDefault)
			objectList[i].(IPersona).SetInitiative(objectList[i].(IPersona).GetIntuition() + objectList[i].(IPersona).GetDataProcessing() + xd6Test(dice))
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("object: " + player.name + " have " + strconv.Itoa(player.GetInitiative()) + " initiative", congo.ColorDefault)
		}
	}*/

}

func hostAction() {
	//startCombatTurn()
	//panic("s")
	//for _, obj := range ObjByNames {
	//	if icon, ok := obj.(IIcon); ok {
	/*for _, icon := range ObjByNames {
	if persona, ok := icon.(*TPersona); ok {
		persona.UpdateSearchProcess()
		/*if persona.GetSearchResultIn() > 0 {
			persona.SetSearchResultIn(persona.GetSearchResultIn() - 1)
			if persona.GetSearchResultIn() == 0 {
				printLog("gogopowerrangers", congo.ColorDefault)
			}
		}*/
	//		}
	//	}

	keysForPersona := getSortedKeysByType("Persona")
	for i := range keysForPersona {
		persona := pickObjByID(keysForPersona[i]).(IPersona)
		persona.SetFullDeffenceFlag(false)
		persona.UpdateSearchProcess()
		persona.UpdateDownloadProcess()
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
				patrolIC.actionReady = patrolIC.actionReady - 1
				if host.alert == "Passive Alert" || host.alert == "Active Alert" {
					patrolIC.actionReady = -1
				}
				if patrolIC.actionReady == 0 {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Host rotine scan initiated...", congo.ColorYellow)
					SourceIcon = patrolIC
					//TargetIcon = "someone"
					mActionName := "EXECUTE_SCAN"
					doAction(mActionName)

					patrolIC.actionReady = calculatePartolScan(patrolIC.deviceRating)
				}
			}
		}
	}

	/*for i := range gridList {

		if host, ok := gridList[i].(*THost); ok {
			host.GatherMarks()
			////////////////////////////////////////обнуляем к фолсу айсы которых по факту нет - костыль

			/////////////////////////////////////////конец костыля
			if host.checkAlert() == "Active Alert" {
				host.LoadNextIC()
			} else {
				if host.PickPatrolIC() != nil {

					patrolIC := host.PickPatrolIC()
					patrolIC.actionReady = patrolIC.actionReady - 1
					if host.alert == "Passive Alert" || host.alert == "Active Alert" {
						patrolIC.actionReady = -1
					}
					if patrolIC.actionReady == 0 {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Host rotine scan initiated...", congo.ColorYellow)
						SourceIcon = patrolIC
						//TargetIcon = "someone"
						mActionName := "EXECUTE_SCAN"
						doAction(mActionName)

						patrolIC.actionReady = calculatePartolScan(patrolIC.deviceRating)
					}
				}
			}
		}
	}*/
}

/*if patrolIC, ok := objectList[i].(*TIC); ok {
	host := patrolIC.GetHost()
	if patrolIC.GetName() == "Patrol IC" {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("LOWER ACTION READY!!!!!!", congo.ColorYellow)
		patrolIC.actionReady = patrolIC.actionReady - 1
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("NEW ACTION READY: "+strconv.Itoa(patrolIC.actionReady), congo.ColorYellow)
		if host.alert == "Passive Alert" || host.alert == "Active Alert" {
			patrolIC.actionReady = -1
			/*	congo.WindowsMap.ByTitle["Log"].WPrintLn("Patrol IC Scan Host...", congo.ColorYellow)
				SourceIcon = patrolIC
				TargetIcon = "someone"
				mActionName := "EXECUTE_SCAN"
				doAction(mActionName)
		}
		if patrolIC.actionReady == 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Patrol IC Scan Host...", congo.ColorYellow)
			SourceIcon = patrolIC
			TargetIcon = "someone"
			mActionName := "EXECUTE_SCAN"
			doAction(mActionName)

			patrolIC.actionReady = calculatePartolScan(patrolIC.deviceRating)
		}
	}
}*/

func (host *THost) checkAlert() string {

	return host.GetHostAlertStatus()
}
