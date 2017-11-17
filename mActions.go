package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/ConGo/congo"
)

//import "math/rand"

//MActions -
var MActions TMActionMap

//MAttributes -
//var MAttributes TMAttributesMap

//TMAttributesMap -
/*type TMAttributesMap struct {
	MAttributes map[string]interface{}
}*/

//TMActionMap -
type TMActionMap struct {
	MActionMap map[string]interface{}
}

//InitMatrixActionMap -
func InitMatrixActionMap() {
	MActions = TMActionMap{}
	MActions.MActionMap = map[string]interface{}{}
	MActions.MActionMap["BRUTE_FORCE"] = BruteForce
	MActions.MActionMap["SWITCH_INTERFACE_MODE"] = SwitchInterfaceMode
	MActions.MActionMap["CHECK_OVERWATCH_SCORE"] = CheckOverwatchScore
	MActions.MActionMap["CRACK_FILE"] = CrackFile
	MActions.MActionMap["DATA_SPIKE"] = DataSpike
	MActions.MActionMap["DISARM_DATABOMB"] = DisarmDataBomb
	MActions.MActionMap["EDIT"] = Edit
	MActions.MActionMap["ENTER_HOST"] = EnterHost
	MActions.MActionMap["EXIT_HOST"] = ExitHost
	MActions.MActionMap["ERASE_MARK"] = EraseMark
	MActions.MActionMap["GRID_HOP"] = GridHop
	MActions.MActionMap["HACK_ON_THE_FLY"] = HackOnTheFly
	MActions.MActionMap["MATRIX_PERCEPTION"] = MatrixPerception
	MActions.MActionMap["MATRIX_SEARCH"] = MatrixSearch
	MActions.MActionMap["SCAN_ENVIROMENT"] = ScanEnviroment
	MActions.MActionMap["SWAP_ATTRIBUTES"] = SwapAttributes
	MActions.MActionMap["LOAD_PROGRAM"] = LoadProgram
	MActions.MActionMap["LOGIN"] = Login
	MActions.MActionMap["UNLOAD_PROGRAM"] = UnloadProgram
	MActions.MActionMap["SET_DATABOMB"] = SetDatabomb
	MActions.MActionMap["SWAP_PROGRAMS"] = SwapPrograms
	MActions.MActionMap["LONGACT"] = LongAct
	MActions.MActionMap["WAIT"] = Wait
	MActions.MActionMap["FULL_DEFENCE"] = FullDefence
	/////////////////////////////////////
	MActions.MActionMap["EXECUTE_SCAN"] = PatrolICActionArea
	MActions.MActionMap["PATROL_IC_ACTION"] = PatrolICActionTarget
	MActions.MActionMap["ACID_IC_ACTION"] = AcidICActionTarget
	MActions.MActionMap["BINDER_IC_ACTION"] = BinderICActionTarget
	MActions.MActionMap["JAMMER_IC_ACTION"] = JammerICActionTarget
	MActions.MActionMap["MARKER_IC_ACTION"] = MarkerICActionTarget
	MActions.MActionMap["KILLER_IC_ACTION"] = KillerICActionTarget
	MActions.MActionMap["SPARKY_IC_ACTION"] = SparkyICActionTarget
	MActions.MActionMap["TAR_BABY_IC_ACTION"] = TarBabyICActionTarget
	MActions.MActionMap["BLACK_IC_ACTION"] = BlackICActionTarget
	MActions.MActionMap["BLASTER_IC_ACTION"] = BlasterICActionTarget
	MActions.MActionMap["PROBE_IC_ACTION"] = ProbeICActionTarget
	MActions.MActionMap["SCRAMBLE_IC_ACTION"] = ScrambleICActionTarget
	MActions.MActionMap["CATAPULT_IC_ACTION"] = CatapultICActionTarget
	MActions.MActionMap["SHOKER_IC_ACTION"] = ShokerICActionTarget
	MActions.MActionMap["TRACK_IC_ACTION"] = TrackICActionTarget
	MActions.MActionMap["BLOODHOUND_IC_ACTION"] = BloodhoundICActionTarget
	MActions.MActionMap["BLOODHOUND_IC_SCAN"] = BloodhoundICActionArea
	MActions.MActionMap["CRASH_IC_ACTION"] = CrashICActionTarget
	MActions.MActionMap["ICWAIT"] = ICWait

	/*DB = TDeviceDB{}
	DB.DeviceDB = map[string]*TDevice{}*/
}

//LongAct -
func LongAct(src IObj, trg IObj) {
	icon := src.(IPersona)
	//trg = TargetIcon.(*THost)

	congo.WindowsMap.ByTitle["Log"].WPrintLn("Start LongAct by "+icon.GetName(), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+"`s searchLen = "+strconv.Itoa(icon.GetLongAct()), congo.ColorDefault)

	endAction()
}

//PatrolICActionArea -
func PatrolICActionArea(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon
	patrolIC := src.(*TIC)
	host := patrolIC.GetHost()
	//hostName := host.GetName()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetDataProcessing()
	suc1, gl, cgl := simpleTest(patrolIC.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		suc1 = 0
	}
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 4)
	if len(comm) < 4 {
		comm = append(comm, "")
		comm = append(comm, "")
		comm = append(comm, "")
		comm = append(comm, "")
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Patrol succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("command: "+text, congo.ColorRed)
	if targ, ok := trg.(IIcon); ok {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("target name = "+targ.GetName(), congo.ColorRed)
	}
	//iconInSilentMode := false
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "No Alert" && patrolIC.actionReady == 0 {
		for _, obj := range ObjByNames {
			if icon, ok := obj.(IIcon); ok {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Target id: "+strconv.Itoa(icon.GetID()), congo.ColorDefault)
				//obj.GetName()
				//icon := *objectList[o].(*TPersona)
				if icon.GetHost().name == src.(*TIC).GetHost().name {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Target id: "+strconv.Itoa(icon.GetID()), congo.ColorRed)
					if icon.GetSilentRunningMode() == true {
						if icon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+icon.GetName(), congo.ColorGreen)
						}
						dp2 := icon.GetSleaze() //+ icon.GetLogic()
						suc2, dgl, dcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
						if dgl {
							suc1++
						}
						if dcgl {
							suc1++
						}
						netHits := suc1 - suc2
						if icon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
							if dgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Encryption glitch detected", congo.ColorYellow)
							}
							if dcgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Encryption critical failure", congo.ColorRed)
							}
						}
						congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
						if netHits > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" affected", congo.ColorRed)
							}
							patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" detected", congo.ColorYellow)
							}
							if icon.GetFaction() != host.GetFaction() {
								host.alert = "Active Alert"
							}
							host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
						}
					}
					if icon.GetSilentRunningMode() == false {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Target is not silent: ", congo.ColorRed)
						if suc1 > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" affected", congo.ColorRed)
							}
							patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" detected", congo.ColorYellow)
							}
							if icon.GetFaction() != host.GetFaction() {
								host.alert = "Active Alert"
							}
						}
					}
				}
			}

		}
	} else {
		//for o := range objectList {
		for _, obj := range ObjByNames {
			if icon, ok := obj.(IPersona); ok {
				//if obj, ok := objectList[o].(IPersona); ok {
				obj.GetName()
				//icon := *objectList[o].(*TPersona)
				if icon.GetHost().name == src.(*TIC).GetHost().name {
					if icon.GetSilentRunningMode() == true {
						dp2 := icon.GetLogic() + icon.GetSleaze()
						suc2, dgl, dcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
						if dgl {
							suc1++
						}
						if dcgl {
							suc1++
						}
						netHits := suc1 - suc2
						if icon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
							if dgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Encryption glitch detected", congo.ColorYellow)
							}
							if dcgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Encryption critical failure", congo.ColorRed)
							}
						}
						congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
						if netHits > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" affected", congo.ColorRed)
							}
							patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" detected", congo.ColorYellow)
							}
							if icon.GetFaction() != host.GetFaction() {
								host.alert = "Active Alert"
							}
							host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
						}
					}
					if icon.GetSilentRunningMode() == false {

						if suc1 > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" affected", congo.ColorRed)
							}
							patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" detected", congo.ColorYellow)
							}
							if icon.GetFaction() != host.GetFaction() {
								host.alert = "Active Alert"
							}
						}
					}
				}
			}
		}
	}
	isComplexAction()
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker
			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender
			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host
			}
		}
	}

	/*if checkMarks(0) == false {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Begin Action", congo.ColorRed)
	} else { //выполняем само действие
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Begin Action", congo.ColorRed)
	}*/
	endAction()
}

//PatrolICActionTarget -
func PatrolICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon
	//hostName := ""

	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetDataProcessing()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Patrol succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	//iconInSilentMode := false
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetSilentRunningMode() == true {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
					}
					dp2 := focusIcon.GetLogic() + focusIcon.GetSleaze()
					suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
					if dgl {
						//suc1++
						addOverwatchScoreToTarget(suc1)
					}
					if dcgl {
						suc1++
						addOverwatchScoreToTarget(10)
					}
					netHits := suc1 - suc2
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
						if dgl {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
						}
						if dcgl {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
						}
					}
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
					if netHits > 0 {
						if focusIcon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
						}
						canSee := src.(*TIC).canSee.KnownData[focusIcon.GetID()]
						canSee[0] = "Spotted"
						src.(*TIC).canSee.KnownData[focusIcon.GetID()] = canSee
						hostSee := host.canSee.KnownData[focusIcon.GetID()]
						hostSee[0] = "Spotted"
						host.canSee.KnownData[focusIcon.GetID()] = hostSee
						host.alert = "Active Alert"
						if focusIcon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...detected", congo.ColorRed)
						}
					}
				}
			}

		}

	}
	/*if checkMarks(0) == false {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Begin Action", congo.ColorRed)
	} else { //выполняем само действие
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Begin Action", congo.ColorRed)
	}*/
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//AcidICActionTarget -
func AcidICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Acid IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					//suc1++
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.GetFirewall() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetDeviceFirewallMod(focusIcon.GetFirewallMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall Rating reduced to "+strconv.Itoa(focusIcon.GetFirewall()), congo.ColorYellow)
						}
						//congo.WindowsMap.ByTitle["Log"].WPrintLn("new FocusIcon.GetFirewall =  "+strconv.Itoa(focusIcon.GetFirewall()), congo.ColorRed)
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}

	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//BinderICActionTarget -
func BinderICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Binder IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetDataProcessing()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					//suc1++
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.GetDataProcessing() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetDeviceDataProcessingMod(focusIcon.GetDataProcessingMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Data Processing Rating reduced to "+strconv.Itoa(focusIcon.GetDataProcessing()), congo.ColorYellow)
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//JammerICActionTarget -
func JammerICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Jammer IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetAttack()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					//suc1++
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.GetAttack() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetDeviceAttackMod(focusIcon.GetAttackMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Attack Rating reduced to "+strconv.Itoa(focusIcon.GetAttack()), congo.ColorYellow)
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//MarkerICActionTarget -
func MarkerICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Marker IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetSleaze()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					//suc1++
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.GetSleaze() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetDeviceSleazeMod(focusIcon.GetSleazeMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Sleaze Rating reduced to "+strconv.Itoa(focusIcon.GetSleaze()), congo.ColorYellow)
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//KillerICActionTarget -
func KillerICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Killer IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - "+src.(*TIC).GetName()+" have "+strconv.Itoa(m)+" marks on "+focusIcon.GetName(), congo.ColorDefault)
					fullDamage := host.GetAttack() + netHits + 2*m
					realDamage := focusIcon.ResistMatrixDamage(fullDamage)
					focusIcon.ReceiveMatrixDamage(realDamage)

				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//SparkyICActionTarget -
func SparkyICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Sparky IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - "+src.(*TIC).GetName()+" have "+strconv.Itoa(m)+" marks on "+focusIcon.GetName(), congo.ColorDefault)
					fullDamage := host.GetAttack() + netHits + 2*m
					realDamage := focusIcon.ResistMatrixDamage(fullDamage)
					focusIcon.ReceiveMatrixDamage(realDamage)
					//////////BIOFEEDBACK DAMAGE//////////////
					biofeedbackDamage := 0
					if focusIcon.GetSimSence() == "HOT-SIM" || focusIcon.GetSimSence() == "COLD-SIM" {
						biofeedbackDamage = realDamage
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "HOT-SIM" {
						focusIcon.ReceivePhysBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "COLD-SIM" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else {
						biofeedbackDamage = 0
					}
					///////////////////////////////////////////

				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//TarBabyICActionTarget -
func TarBabyICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Tar Baby IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetLogic() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected!", congo.ColorYellow)
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure!!!", congo.ColorRed)
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2

				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.GetLinkLockStatus().LockedByID[src.(*TIC).GetID()] == true {
						focusIcon.markSet.MarksFrom[src.(*TIC).GetID()] = focusIcon.markSet.MarksFrom[src.(*TIC).GetID()] + 1
						if focusIcon.markSet.MarksFrom[src.(*TIC).GetID()] > 3 {
							focusIcon.markSet.MarksFrom[src.(*TIC).GetID()] = 3
						}
					} else {
						src.(*TIC).LockIcon(focusIcon)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...connection locked", congo.ColorRed)
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//BlackICActionTarget -
func BlackICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Black IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - "+src.(*TIC).GetName()+" have "+strconv.Itoa(m)+" marks on "+focusIcon.GetName(), congo.ColorDefault)
					fullDamage := host.GetAttack() + netHits + 2*m
					realDamage := focusIcon.ResistMatrixDamage(fullDamage)
					focusIcon.ReceiveMatrixDamage(realDamage)
					//////////BIOFEEDBACK DAMAGE//////////////
					biofeedbackDamage := 0
					if focusIcon.GetSimSence() == "HOT-SIM" || focusIcon.GetSimSence() == "COLD-SIM" {
						biofeedbackDamage = realDamage
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "HOT-SIM" {
						focusIcon.ReceivePhysBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "COLD-SIM" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else {
						biofeedbackDamage = 0
					}
					//////////LINKLOCK////////////////////////
					src.(*TIC).LockIcon(focusIcon)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...connection locked", congo.ColorRed)
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//BlasterICActionTarget -
func BlasterICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Black IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - "+src.(*TIC).GetName()+" have "+strconv.Itoa(m)+" marks on "+focusIcon.GetName(), congo.ColorDefault)
					fullDamage := host.GetAttack() + netHits + 2*m
					realDamage := focusIcon.ResistMatrixDamage(fullDamage)
					focusIcon.ReceiveMatrixDamage(realDamage)
					//////////BIOFEEDBACK DAMAGE//////////////
					biofeedbackDamage := 0
					if focusIcon.GetSimSence() == "HOT-SIM" || focusIcon.GetSimSence() == "COLD-SIM" {
						biofeedbackDamage = realDamage
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "HOT-SIM" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "COLD-SIM" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else {
						biofeedbackDamage = 0
					}
					//////////LINKLOCK////////////////////////
					src.(*TIC).LockIcon(focusIcon)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...connection locked", congo.ColorRed)
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//ProbeICActionTarget -
func ProbeICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Probe IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					focusIcon.markSet.MarksFrom[host.GetID()] = focusIcon.markSet.MarksFrom[host.GetID()] + 1
					if focusIcon.markSet.MarksFrom[host.GetID()] > 3 {
						focusIcon.markSet.MarksFrom[host.GetID()] = 3
					}
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG -"+focusIcon.GetName()+" marked by "+host.GetName(), congo.ColorDefault)
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//ScrambleICActionTarget -
func ScrambleICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Scramble IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.markSet.MarksFrom[host.GetID()] > 3 {
						focusIcon.markSet.MarksFrom[host.GetID()] = 3
					}
					if focusIcon.markSet.MarksFrom[host.GetID()] > 2 {
						focusIcon.Dumpshock()
						focusIcon.SetInitiative(999999)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Connection terminated", congo.ColorGreen)
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//CatapultICActionTarget -
func CatapultICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Catapult IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				altDp := focusIcon.GetLogic() + focusIcon.GetFirewall()
				if altDp > dp2 {
					dp2 = altDp
				}
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.GetFirewall() < 1 {
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - Firewall reduced to 0 ", congo.ColorYellow)
						}
					} else {
						focusIcon.SetDeviceFirewallMod(focusIcon.GetFirewallMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...firewall rating reduced to "+strconv.Itoa(focusIcon.GetFirewall()), congo.ColorYellow)
						}
					}
					//////////BIOFEEDBACK DAMAGE//////////////
					biofeedbackDamage := 0
					if focusIcon.GetSimSence() == "HOT-SIM" || focusIcon.GetSimSence() == "COLD-SIM" {
						biofeedbackDamage = netHits + focusIcon.markSet.MarksFrom[host.GetID()]
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "HOT-SIM" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "COLD-SIM" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else {
						biofeedbackDamage = 0
					}

				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//ShokerICActionTarget -
func ShokerICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Acid IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					focusIcon.SetInitiative(focusIcon.GetInitiative() - 5)
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...initiative reduced by 5", congo.ColorYellow)
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker
			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender
			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host
			}
		}
	}
	endAction()
}

//TrackICActionTarget -
func TrackICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Probe IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			//locks := focusIcon.GetLinkLockStatus()
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetSleaze()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.markSet.MarksFrom[host.GetID()] > 1 {
						focusIcon.SetPhysicalLocation(true)
						if focusIcon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...physical location tracked", congo.ColorGreen)
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...local authorities reported by "+host.GetName()+" administration", congo.ColorGreen)
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//BloodhoundICActionTarget -
func BloodhoundICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	//trg = TargetIcon.(*TPersona)
	for i := range objectList {
		if objectList[i].(IIcon).GetName() == src.(*TIC).GetLastTargetName() {
			trg = objectList[i].(IIcon)
		}
	}
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Bloodhound IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetWillpower() + focusIcon.GetSleaze()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.markSet.MarksFrom[host.GetID()] > 1 {
						focusIcon.SetPhysicalLocation(true)
						if focusIcon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...physical location tracked", congo.ColorGreen)
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...local authorities reported by "+host.GetName()+" administration", congo.ColorGreen)
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker
			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender
			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host
			}
		}
	}
	endAction()
}

//BloodhoundICActionArea -
func BloodhoundICActionArea(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon
	bloodhoundIC := src.(*TIC)
	host := bloodhoundIC.GetHost()
	//hostName := host.GetName()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetDataProcessing()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		suc1 = 0
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Bloodhound succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	//iconInSilentMode := false
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		for o := range objectList {
			if obj, ok := objectList[o].(IPersona); ok {
				obj.GetName()
				icon := *objectList[o].(*TPersona)
				if icon.GetHost().name == src.(*TIC).GetHost().name {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Target id: "+strconv.Itoa(icon.GetID()), congo.ColorRed)
					if icon.GetSilentRunningMode() == true {
						dp2 := icon.GetLogic() + icon.GetSleaze()
						suc2, dgl, dcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
						if dgl {
							suc1++
						}
						if dcgl {
							suc1++
						}
						netHits := suc1 - suc2
						if icon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
							if dgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Encryption glitch detected", congo.ColorYellow)
							}
							if dcgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+": Encryption critical failure", congo.ColorRed)
							}
						}
						congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
						if netHits > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" affected", congo.ColorRed)
							}
							bloodhoundIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" detected", congo.ColorYellow)
							}
							if icon.GetFaction() != host.GetFaction() {
								host.alert = "Active Alert"
								bloodhoundIC.SetLastTargetName(icon.GetName())
							}
						}
					}
					if icon.GetSilentRunningMode() == false {

						if suc1 > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" affected", congo.ColorRed)
							}
							bloodhoundIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" detected", congo.ColorYellow)
							}
							if icon.GetFaction() != host.GetFaction() {
								host.alert = "Active Alert"
								bloodhoundIC.SetLastTargetName(icon.GetName())
							}
						}
					}
				}
			}
		}
	}
	isComplexAction()
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker
			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender
			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host
			}
		}
	}

	/*if checkMarks(0) == false {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Begin Action", congo.ColorRed)
	} else { //выполняем само действие
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Begin Action", congo.ColorRed)
	}*/
	endAction()
}

//CrashICActionTarget -
func CrashICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(*TPersona)
	host := src.(*TIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attacking "+trg.(IObj).GetName(), congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Crash IC succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(*TPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+focusIcon.GetName(), congo.ColorGreen)
				}
				dp2 := focusIcon.GetIntuition() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2
				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall glitch detected", congo.ColorYellow)
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName()+": Firewall critical failure", congo.ColorRed)
					}
				}
				congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - NetHits: "+strconv.Itoa(netHits), congo.ColorDefault)
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected", congo.ColorRed)
					}
					if focusIcon.markSet.MarksFrom[host.GetID()] > 0 {
						focusIcon.CrashRandomProgram()
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded", congo.ColorGreen)
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker

			}
		}
	}
	for i := range objectList {
		if defender, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = defender

			}
		}
	}
	for i := range objectList {
		if host, ok := objectList[i].(IHost); ok {
			if objectList[i].(IHost).GetID() == src.(IHost).GetID() {
				objectList[i] = host

			}
		}
	}
	endAction()
}

//ICWait -
func ICWait(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon
	if ic, ok := SourceIcon.(*TIC); ok {
		ic.SetInitiative(ic.GetInitiative() - 10)
	}
	endAction()

}

/////////////////////////////////////////////////////////////

//AddMAction -
func AddMAction() {
	//MActions.MActionMap["NillAction"] = 1
	MActions.MActionMap["DATA_SPIKE"] = DataSpike
	/*
		MActions.MActionMap["DATA_SPIKE"] = test
		MActions.MActionMap["DATA"] = DataSpike*/
	//return &MActions
}

func checkAction(actionName string) (bool, string) {
	actionIsGood := false
	var mActionName string
	switch actionName {
	case "BRUTE_FORCE":
		actionIsGood = true
		mActionName = "BRUTE_FORCE"
		return actionIsGood, mActionName
	case "SWITCH_INTERFACE_MODE":
		actionIsGood = true
		mActionName = "SWITCH_INTERFACE_MODE"
		return actionIsGood, mActionName
	case "CRACK_FILE":
		actionIsGood = true
		mActionName = "CRACK_FILE"
		return actionIsGood, mActionName
	case "CHECK_OVERWATCH_SCORE":
		actionIsGood = true
		mActionName = "CHECK_OVERWATCH_SCORE"
		return actionIsGood, mActionName
	case "DATA_SPIKE":
		actionIsGood = true
		mActionName = "DATA_SPIKE"
		return actionIsGood, mActionName
	case "DISARM_DATABOMB":
		actionIsGood = true
		mActionName = "DISARM_DATABOMB"
		return actionIsGood, mActionName
	case "EDIT":
		actionIsGood = true
		mActionName = "EDIT"
		return actionIsGood, mActionName
	case "ENTER_HOST":
		actionIsGood = true
		mActionName = "ENTER_HOST"
		return actionIsGood, mActionName
	case "EXIT_HOST":
		actionIsGood = true
		mActionName = "EXIT_HOST"
		return actionIsGood, mActionName
	case "ERASE_MARK":
		actionIsGood = true
		mActionName = "ERASE_MARK"
		return actionIsGood, mActionName
	case "GRID_HOP":
		actionIsGood = true
		mActionName = "GRID_HOP"
		return actionIsGood, mActionName
	case "HACK_ON_THE_FLY":
		actionIsGood = true
		mActionName = "HACK_ON_THE_FLY"
		return actionIsGood, mActionName
	case "MATRIX_PERCEPTION":
		actionIsGood = true
		mActionName = "MATRIX_PERCEPTION"
		return actionIsGood, mActionName
	case "MATRIX_SEARCH":
		actionIsGood = true
		mActionName = "MATRIX_SEARCH"
		return actionIsGood, mActionName
	case "SCAN_ENVIROMENT":
		actionIsGood = true
		mActionName = "SCAN_ENVIROMENT"
		return actionIsGood, mActionName
	case "SET_DATABOMB":
		actionIsGood = true
		mActionName = "SET_DATABOMB"
		return actionIsGood, mActionName
	case "SWAP_ATTRIBUTES":
		actionIsGood = true
		mActionName = "SWAP_ATTRIBUTES"
		return actionIsGood, mActionName
	case "LOAD_PROGRAM":
		actionIsGood = true
		mActionName = "LOAD_PROGRAM"
		return actionIsGood, mActionName
	case "LOGIN":
		actionIsGood = true
		mActionName = "LOGIN"
		return actionIsGood, mActionName
	case "UNLOAD_PROGRAM":
		actionIsGood = true
		mActionName = "UNLOAD_PROGRAM"
		return actionIsGood, mActionName
	case "SWAP_PROGRAMS":
		actionIsGood = true
		mActionName = "SWAP_PROGRAMS"
		return actionIsGood, mActionName
	case "LONGACT":
		actionIsGood = true
		mActionName = "LONGACT"
		return actionIsGood, mActionName
	case "WAIT":
		actionIsGood = true
		mActionName = "WAIT"
		return actionIsGood, mActionName
	case "FULL_DEFENCE":
		actionIsGood = true
		mActionName = "FULL_DEFENCE"
		return actionIsGood, mActionName
	//////////////////////////////////////////
	case "EXECUTE_SCAN":
		actionIsGood = true
		mActionName = "EXECUTE_SCAN"
		return actionIsGood, mActionName
	case "PATROL_IC_ACTION":
		actionIsGood = true
		mActionName = "PATROL_IC_ACTION"
		return actionIsGood, mActionName
	case "ICWAIT":
		actionIsGood = true
		mActionName = "ICWAIT"
		return actionIsGood, mActionName
	default:
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: Unknown command", congo.ColorGreen)
		//panic("fi")

	}
	return actionIsGood, mActionName
}

func doAction(s string) bool {
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("doActtion: "+s, congo.ColorYellow)
	if SourceIcon == nil {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("--DEBUG-- Error: SourceIcon = nil. Try again", congo.ColorRed)
		return false
	}
	if val, ok := MActions.MActionMap[s]; ok {
		if attacker, ok := SourceIcon.(IIcon); ok {
			if defender, ok := TargetIcon.(IPersona); ok {
				if defender.GetID() == player.GetID() && attacker.GetID() != player.GetID() {
					printLog(attacker.GetName()+" attack detected...", congo.ColorYellow)
					printLog("..."+defender.GetName()+": Evaiding", congo.ColorGreen)
				}
			}
		}
		val.(func(IObj, IObj))(SourceIcon, TargetIcon)
		return true
	}
	draw()
	return false
}

//BruteForce - ++
func BruteForce(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	persona := src.(IPersona)
	isComplexAction()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")
	attMod := 0
	markRound := 1
	for i := range comm {
		if comm[i] == "-2M" {
			markRound = 2
		}
		if comm[i] == "-3M" {
			markRound = 3
		}
	}
	printLog("Initiating Brute Force sequence...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetCyberCombatSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetAttack()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure", congo.ColorRed)
	}
	printLog("..."+persona.GetName()+": "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)
	for i := range targetList {
		if grid, ok := targetList[i].(*TGrid); ok {
			dp2 := grid.GetDeviceRating() * 2
			suc2, _, _ := simpleTest(grid.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			netHits = suc1 - suc2
			if netHits > 0 {
				printLog("...Grid encryption bypassed", congo.ColorGreen)
				persona.SetGrid(grid)
			}
		} else if icon, ok := targetList[i].(IIcon); ok {
			host := icon.GetHost()
			dp2 := icon.GetDeviceRating() + icon.GetFirewall()
			suc2, rgl, rcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("...Unexpected exploit detected!", congo.ColorGreen)
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("...Target's firewall critical falure", congo.ColorGreen)
			}
			netHits = suc1 - suc2
			if netHits > 0 {
				for i := 0; i < markRound; i++ {
					placeMARK(persona, icon)
				}
				damage := netHits / 2
				if damage > 0 {
					realDamage := icon.ResistMatrixDamage(damage)
					icon.ReceiveMatrixDamage(realDamage)
				}
				if host.GetHostAlertStatus() == "No Alert" {
					host.alert = "Passive Alert"
				}
			} else {
				persona.ReceiveMatrixDamage(-netHits)
			}
		} else {
			printLog("...Error: Target "+strconv.Itoa(i+1)+" is not a valid type", congo.ColorDefault)
		}
	}
	endAction()
}

//SwitchInterfaceMode - ++
func SwitchInterfaceMode(src IObj, trg IObj) {
	src = SourceIcon
	if persona, ok := src.(*TPersona); ok {
		printLog("Initiate Interface Mode switching...", congo.ColorGreen)
		//text := TargetIcon.(string)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)
		newMode := comm[2]
		newMode = strings.Replace(newMode, "-", "_", -1)
		newMode = strings.Replace(newMode, " ", "_", -1)
		newMode = strings.Replace(newMode, "SIM", "_", -1)
		newMode = strings.Replace(newMode, "VR", "_", -1)
		newMode = strings.Replace(newMode, "_", "", -1)
		printLog(newMode, congo.ColorGreen)
		switch newMode {
		case "AR":
			newMode = "AR"
			printLog("...Switching to Augmented Reality", congo.ColorGreen)
		case "COLD":
			newMode = "COLD-SIM"
			printLog("...Switching to Virtial Reality", congo.ColorGreen)
			printLog("...Safety mode : ON", congo.ColorGreen)
		case "HOT":
			newMode = "HOT-SIM"
			printLog("...Switching to Virtial Reality", congo.ColorGreen)
			printLog("...Safety mode : OFF", congo.ColorYellow)
		default:
			printLog("...Error: User Mode is invalid...", congo.ColorGreen)
			newMode = persona.GetSimSence()

		}
		isSimpleAction()
		persona.SetSimSence(newMode)
		//printLog("--DEBUG--: Changes will be applied on next turn. Canonic Initiative System planed on later date", congo.ColorDefault)

	}
	endAction()
}

//CheckOverwatchScore - ++
func CheckOverwatchScore(src IObj, trg IObj) {
	persona := SourceIcon.(IPersona)
	grid := src.(IPersona).GetGrid()
	dp1 := persona.GetElectronicSkill() + persona.GetLogic() //+ attMod
	limit := persona.GetSleaze()
	if persona.GetSimSence() == "HOT-SIM" {
		dp1 = dp1 + 2
	}
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl {
		addOverwatchScore(8)
	}
	if cgl {
		addOverwatchScore(40)
	}
	isSimpleAction()
	printLog("Checking Overwatch Score...", congo.ColorGreen)
	dp2 := 6
	suc2, dgl, dcgl := simpleTest(-1, dp2, 1000, 0)
	if dgl {
		dgl = false
	}
	if dcgl {
		dcgl = false
	}
	netHits := suc1 - suc2
	if netHits > 0 {
		persona.GetOverwatchScore()
		persona.GetGrid().SetLastSureOS(persona.GetOverwatchScore())
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+grid.GetGridName()+": current Overwatch Score = "+strconv.Itoa(persona.GetGrid().GetLastSureOS()), congo.ColorGreen)
		hold()
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...failed", congo.ColorYellow)
		hold()
	}
	addOverwatchScore(suc2)
	endAction()
}

//CrackFile - ++
func CrackFile(src IObj, trg IObj) {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Data Spike sequence...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetHackingSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetAttack()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Attack protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Attack protocol critical failure", congo.ColorRed)
	}
	for i := range targetList {
		if file, ok := targetList[i].(IFile); ok {
			if !checkExistingMarks(persona.GetID(), file.GetID(), 1) { // проверяем марки
				printLog("..."+file.GetName()+": ACCESS DENIED", congo.ColorYellow)
				continue
			}
			if file.GetDataBombRating() > 0 { //проверяем бомбу
				persona.TriggerDataBomb(file.GetDataBombRating())
				file.SetDataBombRating(0)
				persona.ChangeFOWParametr(file.GetID(), 3, strconv.Itoa(file.GetDataBombRating())) // 3- отвечает за DataBomb
			}
			dp2 := file.GetEncryptionRating() * 2
			suc2, glt, cglt := simpleTest(file.GetID(), dp2, 1000, 0)
			if glt == true {
				addOverwatchScore(dp2 - suc2)
				printLog("...Encryption exploit found", congo.ColorGreen)
			}
			if cglt == true {
				persona.GetGrid().SetOverwatchScore(0)
				printLog("...Overwatch Score cleared", congo.ColorGreen)
			}
			netHits := suc1 - suc2
			addOverwatchScore(suc2)
			if netHits > 0 {
				file.SetEncryptionRating(0)
				printLog("...File encryption disabled", congo.ColorGreen)
				persona.ChangeFOWParametr(file.GetID(), 12, strconv.Itoa(file.GetEncryptionRating())) // 12- отвечает за Encryption
			} else {
				printLog("...Failure! File encryption is not disabled", congo.ColorYellow)
			}
		}
	}
	endAction()
}

//DataSpike - ++
func DataSpike(src IObj, trg IObj) {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Data Spike sequence...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetCyberCombatSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetAttack()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Attack protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Attack protocol critical failure", congo.ColorRed)
	}
	for i := range targetList {
		if icon, ok := targetList[i].(IIcon); ok {
			host := icon.GetHost()
			dp2 := icon.GetDeviceRating() + icon.GetFirewall()
			suc2, rgl, rcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("..."+icon.GetName()+": Firewall exploit detected", congo.ColorGreen)
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("..."+icon.GetName()+": Firewall critical failure", congo.ColorGreen)
			}
			netHits := suc1 - suc2
			if netHits > 0 {
				damage := netHits + persona.GetAttack()
				realDamage := icon.ResistMatrixDamage(damage)
				icon.ReceiveMatrixDamage(realDamage)
				if persona.CheckRunningProgram("Biofeedback") {
					if target, ok := icon.(IPersona); ok {
						bfDamage := target.ResistBiofeedbackDamage(realDamage)
						target.ReceivePhysBiofeedbackDamage(bfDamage)
					} else {
						printLog("...Error: "+icon.GetName()+" is immune to Biofeedback Damage", congo.ColorGreen)
					}
				}
				if persona.CheckRunningProgram("Blackout") {
					if target, ok := icon.(IPersona); ok {
						bfDamage := target.ResistBiofeedbackDamage(realDamage)
						target.ReceiveStunBiofeedbackDamage(bfDamage)
					} else {
						printLog("...Error: "+icon.GetName()+" is immune to Biofeedback Damage", congo.ColorGreen)
					}
				}
				if host.GetHostAlertStatus() == "No Alert" {
					host.alert = "Passive Alert"
				}
			} else {
				persona.ReceiveMatrixDamage(-netHits)
			}

		} else {
			printLog("...Error: Target "+strconv.Itoa(i+1)+" is not a valid type", congo.ColorDefault)
		}
	}
	endAction()

}

//DisarmDataBomb -
func DisarmDataBomb(src IObj, trg IObj) {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Disarm Databomb sequence...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetSoftwareSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetFirewall()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Disarming protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Disarming protocol critical failure", congo.ColorRed)
	}
	for i := range targetList {
		if file, ok := targetList[i].(IFile); ok {
			dp2 := file.GetDataBombRating() * 2
			suc2, rgl, rcgl := simpleTest(file.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("..."+file.GetName()+": Firewall exploit detected", congo.ColorGreen)
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("..."+file.GetName()+": Firewall critical failure", congo.ColorGreen)
			}
			netHits := suc1 - suc2
			if netHits > 0 {
				file.SetDataBombRating(0)
				printLog(persona.GetName()+": Databomb Disarmed", congo.ColorGreen)
			} else {
				persona.TriggerDataBomb(file.GetDataBombRating())
				file.SetDataBombRating(0)
			}
			persona.ChangeFOWParametr(file.GetID(), 3, strconv.Itoa(file.GetDataBombRating())) // 3- отвечает за DataBomb
		} else {
			printLog("...Error: Target "+strconv.Itoa(i+1)+" is not a valid type", congo.ColorDefault)
		}
	}
	endAction()
}

//SetDatabomb -
func SetDatabomb(src IObj, trg IObj) {
	src = SourceIcon.(*TPersona)
	trg = TargetIcon
	icon := SourceIcon.(IPersona)
	//attMod, defMod := getModifiers(src, trg)
	host := trg.(IFile).GetHost()
	dp1 := icon.GetSoftwareSkill() + icon.GetLogic() // + attMod
	limit := src.(IPersona).GetSleaze()

	suc1, gl, cgl := simpleTest(icon.GetID(), dp1, limit, 0)
	if icon.GetFaction() == player.GetFaction() {
		printLog("Setting up databomb on "+trg.GetName()+"...", congo.ColorGreen)
		/*congo.WindowsMap.ByTitle["Log"].WPrintLn("Setting up databomb on "+trg.GetName()+"...", congo.ColorGreen)
		hold()*/
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+icon.GetName()+" "+strconv.Itoa(suc1)+" succeses", congo.ColorGreen)
		hold()
	}

	//congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 0", congo.ColorYellow)
	isComplexAction()
	if trg, ok := trg.(*TFile); ok {
		printLog("...installing databomb", congo.ColorGreen)
		dp2 := host.GetDeviceRating() * 2
		suc2, glt, cglt := simpleTest(trg.GetID(), dp2, 1000, 0)
		if gl == true {
			addOverwatchScore(dp1 - suc1)
			printLog("...Error: Unexpected trigger initiated", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(dp1 - suc1)
			icon.TriggerDataBomb(suc1)
			printLog("...critical error erupted", congo.ColorRed)
		}
		if glt == true {
			addOverwatchScore(-suc2)
			suc2--
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...firewall exploit detected", congo.ColorGreen)
		}
		if cglt == true {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...exploit critical", congo.ColorGreen)
			suc1++
		}
		//Тут надо остановиться и спросить про перебросс

		netHits := suc1 - suc2
		addOverwatchScore(suc2)

		if netHits > 0 {
			if icon.CheckRunningProgram("Demolition") {
				netHits++
				printLog("...Databomb rating infused by Demolition program", congo.ColorGreen)
			}
			trg.SetDataBombRating(netHits)
			printLog("...Databomb rating "+strconv.Itoa(netHits)+" installed", congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Databomb installation failed", congo.ColorGreen)
		}
	}

	endAction()
}

//Edit -
func Edit(src IObj, trg IObj) {
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 4)
	if len(comm) < 4 {
		comm = append(comm, "")
		comm = append(comm, "")
		comm = append(comm, "")
		comm = append(comm, "")
	}
	//editor := src.(*TPersona)
	src = SourceIcon.(*TPersona)
	trg = TargetIcon
	//attMod, defMod := getModifiers(src, trg)
	host := src.(*TPersona).GetHost()
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 0", congo.ColorYellow)
	if persona, ok := src.(*TPersona); ok {
		dp1 := persona.GetComputerSkill() + persona.GetLogic() // + attMod
		limit := persona.GetFirewall()
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if gl == true {
			addOverwatchScore(8)
		}
		if cgl == true {
			addOverwatchScore(40)
		}
		if persona.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Enable Edit mode", congo.ColorGreen)
			if gl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...unexpected error ocured", congo.ColorYellow)
			}
			if cgl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...critical error erupted", congo.ColorRed)
			}
		}
		if file, ok := trg.(*TFile); ok {
			if checkMarks(1) == false {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...ACCESS DENIDED", congo.ColorRed)
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("Not Enough Marks on "+file.GetName(), congo.ColorYellow)
			} else {

				//	congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 1", congo.ColorYellow)

				dp2 := host.GetDeviceRating() + host.GetFirewall()
				suc2, glt, cglt := simpleTest(file.GetID(), dp2, 1000, 0)
				if glt == true {
					addOverwatchScore(-suc2)
				}
				if cglt == true {
					suc1++
				}
				if file.GetEncryptionRating() > 0 {
					suc1 = 0
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: File Encrypted", congo.ColorGreen)
				}
				if file.GetDataBombRating() > 0 {
					persona.TriggerDataBomb(file.GetDataBombRating())
					file.SetDataBombRating(0)
				}
				//Тут надо остановиться и спросить про перебросс

				netHits := suc1 - suc2
				//addOverwatchScore(suc2)
				if netHits > 0 {
					hold()
					if comm[3] == "COPY" {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...copying", congo.ColorGreen)
						hold()
						copy := host.NewFile(file.GetFileName())
						copy.SetFileName("Copy of " + file.GetFileName())
						copy.SetDataBombRating(0)
						copy.SetEncryptionRating(0)
						copy.SetSize(file.GetSize())
						copy.SetLastEditDate(STime)
						copy.markSet.MarksFrom[persona.GetID()] = 4
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...completed", congo.ColorGreen)
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("New file spotted:", congo.ColorGreen)
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Icon: "+copy.GetName(), congo.ColorGreen)
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("File Name: "+copy.GetFileName(), congo.ColorGreen)
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("File Size: "+strconv.Itoa(copy.GetSize())+" Mp", congo.ColorGreen)
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("File Owner: "+persona.GetName(), congo.ColorGreen)
						hold()
					} else if comm[3] == "DELETE" {
						printLog("...deleting file '"+file.GetFileName()+"'", congo.ColorGreen)
						host.DeleteFile(file)
						printLog("...complete", congo.ColorGreen)
					} else if comm[3] == "ENCRYPT" {
						printLog("...encrypting file", congo.ColorGreen)
						file.SetEncryptionRating(netHits)
						printLog("...complete", congo.ColorGreen)
					} else if comm[3] == "DOWNLOAD" {
						printLog("...initiate download: "+file.GetFileName(), congo.ColorGreen)
						printLog("...file size: "+strconv.Itoa(file.GetSize()), congo.ColorGreen)
						persona.SetDownloadProcess(file.GetSize(), file.GetFileName())

					}

					//file.kno
					//удалось
				} else {
					hold()
					//не удалось
				}
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn(file.GetName(), congo.ColorRed)
		}

	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" debug", congo.ColorRed)
	endAction()
}

//EnterHost - ++
func EnterHost(src IObj, trg IObj) {
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	persona := SourceIcon.(IPersona)
	if host, ok := TargetIcon.(*THost); ok {
		printLog("Entering Host...", congo.ColorGreen)
		printLog("...Target host: "+host.GetName(), congo.ColorGreen)
		if checkLinkLock(persona) == true {
			printLog("...Error: "+persona.GetName()+" is Locked", congo.ColorYellow)
		} else {
			if checkMarks(1) == false {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...ACCESS DENIED", congo.ColorRed)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...Not Enough Marks on "+host.GetName(), congo.ColorYellow)
			} else { //выполняем само действие
				persona.SetHost(host)
			}
		}
	} else {
		printLog("Entering Host...", congo.ColorGreen)
		printLog("...Error: Target is not a Host", congo.ColorGreen)
	}
	endAction()
}

//ExitHost - ++
func ExitHost(src IObj, trg IObj) {
	persona := SourceIcon.(IPersona)
	host := persona.GetHost()
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	printLog("Leaving host... ", congo.ColorGreen)
	if checkLinkLock(persona) == true && src.(IObj).GetFaction() == player.GetFaction() {
		printLog("...Error: "+src.(IPersona).GetName()+" is Locked", congo.ColorYellow)
	} else {
		persona.SetHost(host.GetHost())
		//src.(IPersona).SetHost(Matrix)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...successful", congo.ColorGreen)
	}
	endAction()
}

//EraseMark -
func EraseMark(src IObj, trg IObj) {

	isComplexAction() // есть вероятность что стрельнет механизм возврата
	if icon, ok := src.(*TPersona); ok {
		//trg = pickObjByID(2)
		totalMarks := icon.CountMarks()
		if totalMarks > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Erase Mark... ", congo.ColorGreen)
			hold()
			icon.ClearMarks()
			allMarks := icon.GetMarkSet()
			for r := range allMarks.MarksFrom {
				if allMarks.MarksFrom[r] > 0 && r != icon.GetID() {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...MARK found", congo.ColorGreen)
					//congo.WindowsMap.ByTitle["Log"].WPrintLn("ID " + strconv.Itoa(r), congo.ColorYellow)
					hold()
					trg = pickObjByID(r)
					//icon.markSet.MarksFrom[r] = icon.markSet.MarksFrom[r] - 1
					icon.ClearMarks()
					break
				}
			}
		}
		netHits := 0
		dp1 := src.(IPersona).GetComputerSkill() + src.(IPersona).GetLogic()
		limit := src.(IPersona).GetAttack()
		suc1, gl, cgl := simpleTest(icon.GetID(), dp1, limit, 0)
		if gl {
			addOverwatchScore(2)
		}
		if cgl {
			addOverwatchScore(8)
		}
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+strconv.Itoa(suc1)+" successes", congo.ColorGreen)
		hold()
		// if mark from IC
		if markOwner, ok := trg.(*TIC); ok {
			host := markOwner.GetHost()
			dp2 := host.GetDeviceRating() + host.GetFirewall()
			suc2, dgl, dcgl := simpleTest(markOwner.GetID(), dp2, 999, 0)
			if dgl {
				addOverwatchScore(-suc1)
			}
			if dcgl {
				addOverwatchScore(-suc1)
			}
			netHits = suc1 - suc2
			addOverwatchScore(suc2)
			if netHits > 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...MARK erased", congo.ColorGreen)
				hold()
				icon.markSet.MarksFrom[markOwner.GetID()] = icon.markSet.MarksFrom[markOwner.GetID()] - 1
				icon.ClearMarks()
			} else {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...failed ", congo.ColorYellow)
				hold()
			}
		}
		//if mark from Host
		if markOwner, ok := trg.(*THost); ok {
			dp2 := markOwner.GetDeviceRating() + markOwner.GetFirewall()
			suc2, dgl, dcgl := simpleTest(markOwner.GetID(), dp2, 999, 0)
			if dgl {
				addOverwatchScore(-suc1)
			}
			if dcgl {
				addOverwatchScore(-suc1)
			}
			netHits = suc1 - suc2
			addOverwatchScore(suc2)
			if netHits > 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...MARK erased", congo.ColorGreen)
				hold()
				icon.markSet.MarksFrom[markOwner.GetID()] = icon.markSet.MarksFrom[markOwner.GetID()] - 1
				icon.ClearMarks()
			} else {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...failed ", congo.ColorYellow)
				hold()
			}
		}
	}
	endAction()
}

//GridHop -
func GridHop(src IObj, trg IObj) {
	src = SourceIcon.(IPersona)
	if grid, ok := trg.(*TGrid); ok {
		src.(IPersona).SetGrid(grid)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Switching grid to...", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+grid.GetGridName(), congo.ColorGreen)
		isComplexAction()
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...completed", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Welcome to "+grid.GetGridName()+"!", congo.ColorGreen)
		//src.(IPersona).SetInitiative(src.(IPersona).GetInitiative() - 10) //Complex Action
		//присваиваем изменения
		for i := range objectList {
			if attacker, ok := objectList[i].(IPersona); ok {
				if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
					objectList[i] = attacker
				}
			}
		}
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Switching grid to...", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+trg.(IObj).GetName(), congo.ColorYellow)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error!! "+trg.(IObj).GetName()+" is not a grid.", congo.ColorYellow)
	}

	endAction()
}

//HackOnTheFly - ++
func HackOnTheFly(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	persona := src.(IPersona)
	isComplexAction()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")
	//lComm := len(comm) - 1
	attMod := 0

	markRound := 1
	for i := range comm {
		if comm[i] == "-2M" {
			markRound = 2
			attMod = attMod - 4
		}
		if comm[i] == "-3M" {
			markRound = 3
			attMod = attMod - 10
		}
	}

	printLog("Initiating Brute Force sequence...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetHackingSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetSleaze()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure", congo.ColorRed)
	}
	printLog("..."+persona.GetName()+": "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)

	for i := range targetList {
		if grid, ok := targetList[i].(*TGrid); ok {
			dp2 := grid.GetDeviceRating() * 2
			suc2, _, _ := simpleTest(grid.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			netHits = suc1 - suc2
			if netHits > 0 {
				printLog("...Grid encryption bypassed", congo.ColorGreen)
				persona.SetGrid(grid)
			}
		} else if icon, ok := targetList[i].(IIcon); ok {
			host := icon.GetHost()
			dp2 := icon.GetDeviceRating() + icon.GetFirewall()
			suc2, rgl, rcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("...Unexpected exploit detected!", congo.ColorGreen)
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("...Target's firewall critical falure", congo.ColorGreen)
			}
			netHits = suc1 - suc2
			if netHits > 0 {
				for i := 0; i < markRound; i++ {
					placeMARK(persona, icon)
				}
				needToReveal := netHits / 2
				if needToReveal > 0 {
					persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
				}

			} else {
				host.SetAlert("Active Alert")
				placeMARK(icon, persona)
			}
		} else {
			printLog("...Error: "+icon.GetName()+" is not a valid type", congo.ColorDefault)
		}
	}

	endAction()
}

//MatrixPerception - ++
func MatrixPerception(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	persona := src.(IPersona)
	isComplexAction()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")
	attMod := 0
	printLog("Initiating Matrix Perception sequence...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure", congo.ColorRed)
	}
	printLog("..."+persona.GetName()+": "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)

	for j := range targetList {
		needToReveal := suc1
		if icon, ok := targetList[j].(IIcon); ok {
			if icon.GetSilentRunningMode() {
				dp2 := icon.GetSleaze() + icon.GetDeviceRating()
				suc2, dgl, cdgl := simpleTest(icon.GetID(), dp2, 1000, 0)
				if dgl == true {
					addOverwatchScore(-suc1)
					printLog("...Encryption weakness detected", congo.ColorGreen)
				}
				if cdgl == true {
					addOverwatchScore(-suc1 * 2)
					suc1 = suc1 + 100
					printLog("...Encryption critical falure", congo.ColorGreen)

				}
				netHits = suc1 - suc2
				needToReveal = netHits
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
				if netHits < 0 {
					printLog("...Matrix Perception failed", congo.ColorYellow)
				}
			} else {
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
			}
			/*if netHits > 0 {
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
			} else {
				printLog("...Matrix Perception failed", congo.ColorYellow)
			}*/

		}
	}

	endAction()

}

//MatrixSearch -
func MatrixSearch(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	//var netHits int
	persona := src.(IPersona)
	isComplexAction()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")
	//lComm := len(comm) - 1
	attMod := 0
	var iconType string
	var iconName string
	if len(comm) > 2 {
		iconType = comm[2]
	}
	if len(comm) > 3 {
		iconName = comm[3]
	}

	/*markRound := 1
	for i := range comm {
		if comm[i] == "-2M" {
			markRound = 2
		}
		if comm[i] == "-3M" {
			markRound = 3
		}
	}*/

	printLog("Initiating Matrix Search sequense...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	if persona.CheckRunningProgram("Search") {
		attMod = attMod + 2
		printLog("...'Search' program running: "+strconv.Itoa(2)+" op/p", congo.ColorGreen)
	}
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure", congo.ColorRed)
	}
	if suc1 > 0 {

		printLog("..."+persona.GetName()+": "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)
		searchBase := 20
		if persona.CheckRunningProgram("Browse") {
			searchBase = 10
		}
		resultIn := searchBase / suc1
		if resultIn == 0 {
			resultIn = 1
		}
		//persona.SetSearchResultIn(resultIn)
		persona.SetSearchProcess(resultIn, iconType, iconName)

		printLog("..."+persona.GetName()+": Search in progress", congo.ColorGreen)
		for i := range comm {
			comm[i] = formatTargetName(comm[i])
			//	printLog("comm["+strconv.Itoa(i)+"] = "+comm[i], congo.ColorDefault)
		}
		//printLog("...Search ETA: "+strconv.Itoa(persona.GetSearchResultIn()*3)+" seconds", congo.ColorGreen)
	} else {
		printLog("..."+persona.GetName()+": Search Failed", congo.ColorGreen)
	}
	endAction()
}

//ScanEnviroment - ++
func ScanEnviroment(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	var targetList []IObj
	persona := src.(IPersona)
	isComplexAction()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")
	attMod := 0
	printLog("Initiating Matrix Perception sequence...", congo.ColorGreen)
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			canSee := persona.GetFieldOfView().KnownData[icon.GetID()]
			if icon.GetHost() == persona.GetHost() && canSee[0] != "Spotted" {
				targetList = append(targetList, icon)
			}
		}
	}

	//targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList[:1])
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: "+strconv.Itoa(limit)+" op/p", congo.ColorGreen)
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure", congo.ColorRed)
	}
	printLog("..."+persona.GetName()+": "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)

	for j := range targetList {
		needToReveal := suc1
		if icon, ok := targetList[j].(IIcon); ok {
			if icon.GetSilentRunningMode() {
				dp2 := icon.GetSleaze() + icon.GetDeviceRating()
				suc2, dgl, cdgl := simpleTest(icon.GetID(), dp2, 1000, 0)
				if dgl == true {
					addOverwatchScore(-suc1)
					printLog("...Encryption weakness detected", congo.ColorGreen)
				}
				if cdgl == true {
					addOverwatchScore(-suc1 * 2)
					suc1 = suc1 + 100
					printLog("...Encryption critical falure", congo.ColorGreen)

				}
				netHits = suc1 - suc2
				needToReveal = netHits
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
				if netHits < 0 {
					printLog("...Matrix Perception failed", congo.ColorYellow)
				}
			} else {
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
			}
			/*if netHits > 0 {
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
			} else {
				printLog("...Matrix Perception failed", congo.ColorYellow)
			}*/

		}
		printLog("..."+strconv.Itoa(len(targetList))+" icons running silent detected", congo.ColorYellow)
		break
	}

	endAction()
}

//SwapAttributes -
func SwapAttributes(src IObj, trg IObj) {
	src = SourceIcon
	if persona, ok := src.(*TPersona); ok {
		printLog("Initiate attributes swapping...", congo.ColorGreen)
		//text := TargetIcon.(string)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)

		att1 := 0
		switch comm[2] {
		case "ATTACK":
			att1 = SourceIcon.(*TPersona).GetAttackRaw()
			printLog("...Attribute 1 = Attack", congo.ColorGreen)
		case "SLEAZE":
			att1 = SourceIcon.(*TPersona).GetSleazeRaw()
			printLog("...Attribute 1 = Sleaze", congo.ColorGreen)
		case "DATA_PROCESSING":
			att1 = SourceIcon.(*TPersona).GetDataProcessingRaw()
			printLog("...Attribute 1 = Data Processing", congo.ColorGreen)
		case "FIREWALL":
			att1 = SourceIcon.(*TPersona).GetFirewallRaw()
			printLog("...Attribute 1 = Firewall", congo.ColorGreen)
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 1 is invalid...", congo.ColorYellow)

		}
		att2 := 0
		switch comm[3] {
		case "ATTACK":
			att2 = SourceIcon.(*TPersona).GetAttackRaw()
			printLog("...Attribute 2 = Attack", congo.ColorGreen)
		case "SLEAZE":
			att2 = SourceIcon.(*TPersona).GetSleazeRaw()
			printLog("...Attribute 2 = Sleaze", congo.ColorGreen)
		case "DATA_PROCESSING":
			att2 = SourceIcon.(*TPersona).GetDataProcessingRaw()
			printLog("...Attribute 2 = Data Processing", congo.ColorGreen)
		case "FIREWALL":
			att2 = SourceIcon.(*TPersona).GetFirewallRaw()
			printLog("...Attribute 2 = Firewall", congo.ColorGreen)
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 2 is invalid...", congo.ColorYellow)
		}
		if persona.device.canSwapAtt == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("This Persona cann't swap attributes!", congo.ColorRed)
			comm[2] = " "
			comm[3] = " "
		}
		swap1 := false
		if comm[2] == "ATTACK" {
			persona.SetDeviceAttackRaw(att2)
			swap1 = true
		} else if comm[2] == "SLEAZE" {
			persona.SetDeviceSleazeRaw(att2)
			swap1 = true
		} else if comm[2] == "DATA_PROCESSING" {
			persona.SetDeviceDataProcessingRaw(att2)
			swap1 = true
		} else if comm[2] == "FIREWALL" {
			persona.SetDeviceFirewallRaw(att2)
			swap1 = true
		} else {
			swap1 = false
		}
		swap2 := false
		if comm[3] == "ATTACK" {
			persona.SetDeviceAttackRaw(att1)
			swap2 = true
		} else if comm[3] == "SLEAZE" {
			persona.SetDeviceSleazeRaw(att1)
			swap2 = true
		} else if comm[3] == "DATA_PROCESSING" {
			persona.SetDeviceDataProcessingRaw(att1)
			swap2 = true
		} else if comm[3] == "FIREWALL" {
			persona.SetDeviceFirewallRaw(att1)
			swap2 = true
		} else {
			swap2 = false
		}
		if comm[2] == comm[3] {
			swap1 = false
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 1 = Attribute 2", congo.ColorYellow)
		}
		if swap1 == true && swap2 == true {
			isFreeAction()
			printLog("Attribute swapping complete", congo.ColorGreen)
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Attribute swapping failed", congo.ColorYellow)
		}
	}
	endAction()
}

//LoadProgram -
func LoadProgram(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(*TPersona); ok {
		//text := TargetIcon.(string)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)
		programINfound := false
		programINRunning := false
		for i := range persona.GetDeviceSoft().programName {
			program := persona.GetDeviceSoft().programName[i]
			prgName := persona.GetDeviceSoft().programName[i]
			prgName = formatString(prgName)
			prgName = cleanText(prgName)
			dur := time.Second / 3
			draw()
			if prgName == comm[2] {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Loading program...", congo.ColorGreen)
				time.Sleep(dur)
				draw()
				if persona.GetDeviceSoft().programStatus[i] == "Running" {
					programINRunning = true
				}
				if persona.GetDeviceSoft().programStatus[i] == "inStore" {

					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program name: "+persona.GetDeviceSoft().programName[i], congo.ColorGreen)
					time.Sleep(dur)
					draw()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program type: "+persona.GetDeviceSoft().programType[i], congo.ColorGreen)
					time.Sleep(dur)
					draw()
					if persona.LoadProgram(program) { //////Проверка устройства
						//persona.LoadProgram(program)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...program status: "+persona.GetDeviceSoft().programStatus[i], congo.ColorGreen)
						time.Sleep(dur)
						draw()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Loading Complete", congo.ColorGreen)
						programINfound = true
					}
				}
				if persona.GetDeviceSoft().programStatus[i] == "Crashed" {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program name: "+persona.GetDeviceSoft().programName[i], congo.ColorGreen)
					time.Sleep(dur)
					draw()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program type: "+persona.GetDeviceSoft().programType[i], congo.ColorGreen)
					time.Sleep(dur)
					draw()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program status: "+persona.GetDeviceSoft().programStatus[i], congo.ColorRed)
					time.Sleep(dur)
					draw()
				}
			}
		}
		if programINfound {
			isFreeAction()
		} else if programINRunning {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: program '"+comm[2]+"' already running", congo.ColorYellow)
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program '"+comm[2]+"' cannot be loaded", congo.ColorGreen)
		}
	}

	endAction()
}

//Login -
func Login(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 3)
	target := formatTargetName(comm[2])
	if persona, ok := src.(*TPersona); ok {
		printLog(">>>LOGIN: "+target, congo.ColorGreen)
		printLog(">>>PASSCODE: XXXXXXXXXXXXXXXXX", congo.ColorGreen)
		if persona.GetName() == "Unknown" {

			if target == "Unknown" {
				printLog("...Error: 'Guest' already signed in", congo.ColorGreen)
			} else {
				var valid bool
				player, valid = ImportPlayerFromDB(target)
				delete(ObjByNames, "Unknown")
				if valid {
					printLog("...Passcode accepted", congo.ColorGreen)
					printLog("...Biometric data generated", congo.ColorGreen)
					printLog("...Start session:", congo.ColorGreen)
				}
			}
		} else {
			printLog("SYSTEM ERROR: Persona already logged in.", congo.ColorDefault)
			printLog("Terminate session if you want to use another account", congo.ColorDefault)
		}
	}
	/*for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
				printLog(">>>LOGIN: "+icon.GetName(), congo.ColorGreen)
				printLog(">>>PASSCODE: XXXXXXXXXXXXXXXXX", congo.ColorGreen)
				printLog("...Passcode accepted", congo.ColorGreen)
				printLog("...Biometric data generated", congo.ColorGreen)
				printLog("Play nice chummer!", congo.ColorDefault)
				printLog("...Begin Session:", congo.ColorGreen)

		}
	}*/
	endAction()

	SourceIcon = player
}

//UnloadProgram -
func UnloadProgram(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(*TPersona); ok {
		//text := TargetIcon.(string)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)
		programFound := false
		for i := range persona.GetDeviceSoft().programName {
			program := persona.GetDeviceSoft().programName[i]
			prgName := persona.GetDeviceSoft().programName[i]
			prgName = formatString(prgName)
			prgName = cleanText(prgName)
			dur := time.Second / 3
			draw()

			if prgName == comm[2] {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Stopping program...", congo.ColorGreen)
				time.Sleep(dur)
				draw()
				if persona.GetDeviceSoft().programStatus[i] == "Running" {
					if true { //////Проверка устройства
						persona.UnloadProgram(program)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+program+" Terminated", congo.ColorGreen)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Program exit code:0", congo.ColorGreen)
						isFreeAction()

					}
				} else {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: "+program+" is "+persona.GetDeviceSoft().programStatus[i], congo.ColorGreen)
				}
				programFound = true
			}

		}
		if programFound == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: program '"+comm[2]+"' not found", congo.ColorGreen)
		}
	}

	endAction()
}

//SwapPrograms - ++
func SwapPrograms(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(*TPersona); ok {
		//text := TargetIcon.(string)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)
		programFound := false
		for i := range persona.GetDeviceSoft().programName {
			program := persona.GetDeviceSoft().programName[i]
			prgNameOut := persona.GetDeviceSoft().programName[i]
			prgNameOut = formatString(prgNameOut)
			prgNameOut = cleanText(prgNameOut)
			dur := time.Second / 3
			draw()
			if prgNameOut == comm[2] {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Stopping program...", congo.ColorGreen)
				time.Sleep(dur)
				draw()
				if persona.GetDeviceSoft().programStatus[i] == "Running" {
					if true { //////Проверка устройства
						persona.UnloadProgram(program)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+program+" Terminated", congo.ColorGreen)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Program exit code:0", congo.ColorGreen)
						isFreeAction()

					}
				} else {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: "+program+" is "+persona.GetDeviceSoft().programStatus[i], congo.ColorGreen)
				}
				programFound = true
			}
		}
		if programFound == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: program '"+comm[2]+"' not found", congo.ColorGreen)
		}
		programINfound := false
		programINRunning := false
		for i := range persona.GetDeviceSoft().programName {
			program := persona.GetDeviceSoft().programName[i]
			prgName := persona.GetDeviceSoft().programName[i]
			prgName = formatString(prgName)
			prgName = cleanText(prgName)
			dur := time.Second / 3
			draw()
			if prgName == comm[3] {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Loading program...", congo.ColorGreen)
				time.Sleep(dur)
				draw()
				if persona.GetDeviceSoft().programStatus[i] == "Running" {
					programINRunning = true
				}
				if persona.GetDeviceSoft().programStatus[i] == "inStore" {

					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program name: "+persona.GetDeviceSoft().programName[i], congo.ColorGreen)
					time.Sleep(dur)
					draw()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program type: "+persona.GetDeviceSoft().programType[i], congo.ColorGreen)
					time.Sleep(dur)
					draw()
					if persona.LoadProgram(program) { //////Проверка устройства
						//persona.LoadProgram(program)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...program status: "+persona.GetDeviceSoft().programStatus[i], congo.ColorGreen)
						time.Sleep(dur)
						draw()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Loading Complete", congo.ColorGreen)
						programINfound = true
					}
				}
			}
		}
		if programINfound {
			isFreeAction()
		} else if programINRunning {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: program '"+comm[3]+"' already running", congo.ColorYellow)
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program '"+comm[3]+"' cannot be loaded", congo.ColorGreen)
		}
	}
	endAction()
}

//Wait -
func Wait(src IObj, trg IObj) {
	src = SourceIcon.(IPersona)
	src = SourceIcon

	icon := SourceIcon.(IPersona)
	//icon.RollInitiative()
	//text := TargetIcon.(string)
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 4)
	if len(comm) > 2 {
		waitTime := comm[2]
		waitTimeInt, _ := strconv.Atoi(waitTime)
		icon.SetInitiative(icon.GetInitiative() - waitTimeInt)
	} else {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("Wait time unspecified...", congo.ColorGreen)
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("Waiting until end of turn...", congo.ColorDefault)
		src.(IPersona).SetInitiative(0)
	}
	//icon.RollInitiative()

	endAction()

}

//FullDefence -
func FullDefence(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(*TPersona); ok {
		printLog("Full defence protocol initiated", congo.ColorGreen)
		persona.SetFullDeffenceFlag(true)
		isComplexAction()
	}
	endAction()

}

func addOverwatchScore(suc2 int) {
	if icon, ok := SourceIcon.(IIcon); ok {
		icon.SetOverwatchScore(icon.GetOverwatchScore() + suc2)
	}
}

func addOverwatchScoreToTarget(suc2 int) {
	if icon, ok := TargetIcon.(IIcon); ok {
		icon.SetOverwatchScore(icon.GetOverwatchScore() + suc2)
	}
}

func endAction() {
	SourceIcon = nil
	TargetIcon = nil
	TargetIcon2 = nil
	command = "--EMPTY"
	//outIndex := 0
	for _, obj := range ObjByNames {
		if ic, ok := obj.(IIC); ok {
			if ic.GetMatrixCM() < 0 {
				host := ic.GetHost()
				host.DeleteIC(ic.(*TIC))
			}
		}
	}

	/*for _, x := range objectList {
		if objectList[outIndex].(IObj).GetType() != "File" {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Validate Icon: "+objectList[outIndex].(IObj).GetName(), congo.ColorDefault)
		}
		if icon, ok := objectList[outIndex].(IIcon); ok {
			if approveDeletion(icon) == false {
				if objectList[outIndex].(IObj).GetType() != "File" {
					//congo.WindowsMap.ByTitle["Log"].WPrintLn("Icon validated: "+objectList[outIndex].(IObj).GetName(), congo.ColorDefault)
				}
				objectList[outIndex] = x
				outIndex++
				//			objectList = objectList[:outIndex]
			} else {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("DELETE OBJECT: "+objectList[outIndex].(IObj).GetName(), congo.ColorDefault)
				if icon.GetType() == "IC" {
					host := icon.GetHost()
					host.DeleteIC(icon.(*TIC))
				}

			}
		}

	}*/
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(command, congo.ColorDefault)
	checkTurn()
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
	refreshProcessWin()
}

func getModifiers(src IObj, trg IObj) (int, int) {
	attMod := 0
	defMod := 0
	//Public Grid
	/*srcGrid := src.(*TIcon).GetGrid()
	trgGrid := src.(*TIcon).GetGrid()
	gridMod := 0
	if srcGrid.GetGridName() == "Public Grid" {
		gridMod = gridMod - 2
	}
	if srcGrid.GetName() != trgGrid.GetName() {
		gridMod = gridMod - 2
	}
	if attacker, ok := src.(IPersona); ok {
		//получаем бонусные кубы от программ
	}*/
	return attMod, defMod
}

func isComplexAction() {
	/*src := SourceIcon
	if source, ok := src.(IIcon); ok {
		source.SetInitiative(source.GetInitiative() - 10) //Complex Action
	}
	/*if source, ok := src.(IIC); ok {
		source.SetInitiative(source.GetInitiative() - 10) //Complex Action
	}

	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = attacker
			}
		}
	}*/
	if src, ok := SourceIcon.(IIcon); ok {
		src.SetInitiative(src.GetInitiative() - 10)
	}
}

func isFreeAction() {
	src := SourceIcon
	src.(IPersona).SetInitiative(src.(IPersona).GetInitiative() - 2) //Free Action
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
				objectList[i] = attacker
			}
		}
	}
}

func isSimpleAction() {
	src := SourceIcon
	src.(IPersona).SetInitiative(src.(IPersona).GetInitiative() - 5) //Simple Action
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
				objectList[i] = attacker
			}
		}
	}
}

func checkMarks(neededMarks int) bool {
	src := SourceIcon.(*TPersona)
	if trg, ok := TargetIcon.(*TIcon); ok {
		markSet := trg.GetMarkSet()
		if markSet.MarksFrom[src.GetID()] < neededMarks { //проверяем наличие марок
			return false
		}
		return true
	} else if trg, ok := TargetIcon.(*THost); ok {
		markSet := trg.GetMarkSet()
		if markSet.MarksFrom[src.GetID()] < neededMarks { //проверяем наличие марок
			return false
		}
		return true
	} else if trg, ok := TargetIcon.(*TFile); ok {
		markSet := trg.GetMarkSet()
		if markSet.MarksFrom[src.GetID()] < neededMarks { //проверяем наличие марок
			return false
		}
		return true
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Ошибка! Неизвестный тип для checkMarks()!", congo.ColorRed)
	return false
}

func checkExistingMarks(srcID, trgID, neededMarks int) bool {
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetID() != trgID {
				continue
			}
			currentMARKS := icon.GetMarkSet().MarksFrom[srcID]
			if neededMarks > currentMARKS {
				return false
			}
			return true
		}
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Ошибка! Неизвестный тип для checkMarks()!", congo.ColorRed)
	return false
}

/*
if comm[2] == "ALL" {
		for o := range objectList {
		if objectList[o].(IObj).GetType() == "Host" {
			host := objectList[o].(*THost)
			dp2 := host.GetDeviceRating() + host.GetSleaze()
			suc2, glt, cglt := simpleTest( focusIcon.GetID() ,dp2, 1000, 0)
			if glt {
				suc2 = 0
			}
			if cglt {
				addOverwatchScore(0)
			}
			netHits := suc1 - suc2
			if netHits > 0 {
				bytesToReveal := netHits
				if src.(IObj).GetID() == player.GetID(){
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Player reveals " + strconv.Itoa(netHits) + " bytes of data...", congo.ColorRed)
				}
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("Player reveals " + strconv.Itoa(netHits) + " bytes of data...", congo.ColorRed)
				allInfo := src.(*TPersona).canSee.KnownData[host.GetID()]
				allInfo[0] = "Unknown" //Spot
				allInfo[1] = "Unknown" //EditDate - File
				allInfo[2] = "Unknown" //MCM
				allInfo[3] = "Unknown" //DataBomb
				allInfo[4] = "Unknown" //ProgramList
				allInfo[5] = "Unknown" //DeviceRating
				allInfo[6] = "Unknown" //ComMode
				allInfo[7] = "Unknown" //AtT
				allInfo[8] = "Unknown" //SLZ
				allInfo[9] = "Unknown" //DTPRC
				allInfo[10] = "Unknown" //FRW
				allInfo[11] = "Unknown" //uType
				allInfo[12] = "Unknown" //Encrypt
				allInfo[13] = "Unknown" //Grid
				allInfo[14] = "Unknown" //Proxy - возможно не надо
				allInfo[15] = "Unknown" //SilentSpot - возможно не нужно
				allInfo[16] = "Unknown" //LastAction - time
				allInfo[17] = "Unknown" //Marks - howmany - ID
				allInfo[18] = "Unknown" //owner
				canSee := src.(*TPersona).canSee.KnownData[host.GetID()]
				if canSee[0] == "Unknown" {
					canSee = allInfo
					canSee[0] = "Spotted"
					bytesToReveal--
					src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
				}

				scanableForHost := make([]int, 12, 12)
				for i := 0; i< len(canSee); i++ {
					if i == 0 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 4 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 5 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 7 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 8 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 9 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 10 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 11 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 13 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 16 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 17 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
					if i == 18 && canSee[i] != "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} //12

				}
				ableToScan := scanableForHost
				for i := bytesToReveal; i>0;i-- {
					if len(ableToScan) == 0 {
						break
					}
					//reveal := 0
					setSeed()
					choosen := ableToScan[rand.Intn(len(ableToScan)-1)]
					setSeed()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Player try to reveal  parametr # " + strconv.Itoa(choosen), congo.ColorYellow)
					switch choosen {
						case 4:
						canSee[choosen] = "All IC List"
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
						case 5:
						canSee[choosen] = strconv.Itoa(host.GetDeviceRating())
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
						case 7:
						canSee[choosen] = strconv.Itoa(host.GetAttack())
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
						case 8:
						canSee[choosen] = strconv.Itoa(host.GetSleaze())
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
						case 9:
						canSee[choosen] = strconv.Itoa(host.GetDataProcessing())
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
						case 10:
						canSee[choosen] = strconv.Itoa(host.GetFirewall())
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
						default:
						src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
					}
					src.(*TPersona).canSee.KnownData[host.GetID()] = canSee
				}
				src.(*TPersona).canSee.KnownData[host.GetID()] = canSee

			} else {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("No DATA", congo.ColorRed)
			}
		}
	}*/

func checkLinkLock(icon IIcon) bool {
	allLocks := icon.GetLinkLockStatus()
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is checking Locks in "+icon.GetGridName(), congo.ColorDefault)
	var lockedBy []int
	for key, value := range allLocks.LockedByID { //check if non-slave marked by slave
		//		congo.WindowsMap.ByTitle["Log"].WPrintLn("checking: "+strconv.Itoa(key), congo.ColorDefault)
		if isLocked(allLocks.LockedByID, key) { //if true
			if value == true {
				lockedBy = append(lockedBy, key)
				//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is Locked in "+icon.GetGridName()+" by "+pickObjByID(key).(IObj).GetName(), congo.ColorRed)
			} else {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is NOT Locked in "+icon.GetGridName()+" by "+pickObjByID(key).(IObj).GetName(), congo.ColorRed)
			}
			//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is Locked in "+icon.GetGridName()+" by "+objectList[key].(IObj).GetName(), congo.ColorRed)
		}
	}
	mustReturn := false
	for i := range lockedBy {
		i++
		i--
		//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is Locked in "+icon.GetGridName()+" by "+pickObjByID(lockedBy[i]).(IObj).GetName(), congo.ColorGreen)
		mustReturn = true
	}
	if mustReturn == true {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is Locked", congo.ColorGreen)
	}
	return mustReturn
}

func isLocked(m map[int]bool, key int) bool {
	for _, a := range m {
		if a == true {
			return true
		}
	}
	return false
}

func calculateAttMods(comm []string, attacker IIcon, targetList []IObj) (attMod int) {
	/*if len(comm) > 3 {
		if comm[len(comm)-1] == "-2M" {
			attMod = attMod - 4
			printLog("...Additional operation cycles: "+strconv.Itoa(-4)+" op/p", congo.ColorGreen)
		}
		if comm[len(comm)-1] == "-3M" {
			attMod = attMod - 10
			printLog("...Additional operation cycles: "+strconv.Itoa(-10)+" op/p", congo.ColorGreen)
		}
	}*/
	var oppCyc bool
	for i := range comm {
		if comm[i] == "-2M" && oppCyc == false {
			printLog("...Additional operation cycles: "+strconv.Itoa(-4)+" op/p", congo.ColorGreen)
			attMod = attMod - 4
			oppCyc = true
		}
		if comm[i] == "-3M" && oppCyc == false {
			printLog("...Additional operation cycles: "+strconv.Itoa(-10)+" op/p", congo.ColorGreen)
			attMod = attMod - 10
			oppCyc = true
		}
	}
	if attacker.GetGrid().name == "Public Grid" {
		attMod = attMod - 2
		printLog("...Public Grid lags: "+strconv.Itoa(-2)+" op/p", congo.ColorGreen)
	}
	for i := range targetList {
		if trgt, ok := targetList[i].(IIcon); ok {
			if attacker.GetGrid() != trgt.GetGrid() {

				attMod = attMod - 2
				printLog("...Target "+strconv.Itoa(i+1)+" is in another Grid: "+strconv.Itoa(-2)+" op/p", congo.ColorGreen)
			}
		}
	}

	if attacker.GetSimSence() == "HOT-SIM" {
		attMod = attMod + 2
		printLog("...HOT-SIM connection boost: "+strconv.Itoa(2)+" op/p", congo.ColorGreen)
	}

	return attMod
}

//helper funcs:

func formatTargetName(targetName string) string {
	targetName = strings.ToLower(targetName)
	targetName = strings.Replace(targetName, "_", " ", -1)
	targetName = strings.Title(targetName)
	targetName = strings.Replace(targetName, " Ic", " IC", -1)
	return targetName
}

func pickTargets(comm []string) []IObj {
	var targetList []IObj
	if len(comm) < 3 {
		return targetList
	}
	targetName := formatTargetName(comm[2])
	if grid, ok := ObjByNames[targetName].(*TGrid); ok {
		targetList = append(targetList, grid)
		printLog("...Target 1: "+grid.GetGridName()+" has top priority", congo.ColorYellow)
		//printLog("...Target 1 replaced", congo.ColorGreen)
		return targetList
	}

	if icon1, ok := ObjByNames[targetName]; ok {
		newIcon := icon1.(IIcon) //не выводится за зону видимости((
		targetList = append(targetList, newIcon)
		printLog("...Target 1: "+newIcon.GetName(), congo.ColorGreen)
		persona := SourceIcon.(IIcon)
		if persona.CheckRunningProgram("Fork") && len(comm) > 3 {
			targetName2 := formatTargetName(comm[3])
			if targetName != targetName2 {
				if icon2, ok := ObjByNames[targetName2]; ok {
					if grid, ok := ObjByNames[targetName].(*TGrid); ok {
						targetList = nil
						targetList = append(targetList, grid)
						printLog("...Target 2: "+grid.GetGridName()+" has top priority", congo.ColorYellow)
						printLog("...Target 1 replaced", congo.ColorGreen)
						return targetList
					}
					newIcon2 := icon2.(IIcon)
					targetList = append(targetList, newIcon2)
					printLog("...Target 2: "+newIcon2.GetName(), congo.ColorGreen)
				}
			} else {
				printLog("...Error: Target 1 = Target 2", congo.ColorYellow)
			}
		}
	}
	return targetList
}

func placeMARK(source, target IIcon) {
	/*markMap := target.GetMarkSet().MarksFrom
	if markMap != nil {
		currentMARKS := target.GetMarkSet().MarksFrom[source.GetID()]
		currentMARKS++
		if currentMARKS > 3 {
			currentMARKS = 3
		}
		target.GetMarkSet().MarksFrom[source.GetID()] = currentMARKS
	} else {
		host := target.GetHost()
		currentMARKS := host.GetMarkSet().MarksFrom[source.GetID()]
		currentMARKS++
		if currentMARKS > 3 {
			currentMARKS = 3
		}
		host.GetMarkSet().MarksFrom[source.GetID()] = currentMARKS
	}*/

	currentMARKS := target.GetMarkSet().MarksFrom[source.GetID()]
	currentMARKS++
	printLog("...new MARK on "+target.GetName()+" was successfuly planted", congo.ColorGreen)
	if currentMARKS > 3 {
		currentMARKS = 3
	}
	target.GetMarkSet().MarksFrom[source.GetID()] = currentMARKS
	master := target.GetOwner()
	if master != nil && master != target {
		placeMARK(source, master)
	}
}

func revealData(persona IPersona, icon IIcon, needToReveal int) [30]string {
	canSee := persona.GetFieldOfView().KnownData[icon.GetID()]
	mem := make([]int, 0, 30)
	for i := range canSee {
		//printLog("canSee["+strconv.Itoa(i)+"] = "+canSee[i], congo.ColorGreen)
		if canSee[i] == "Unknown" { //|| canSee[0] != "Spotted" {
			//printLog("canSee["+strconv.Itoa(i)+"] = "+canSee[i], congo.ColorGreen)
			//printLog("mem: "+strconv.Itoa(i), congo.ColorGreen)
			mem = append(mem, i)
		}

	}
	/*for i := range mem {
		congo.WindowsMap.ByTitle["Log"].WPrint(" "+strconv.Itoa(mem[i]), congo.ColorDefault)
	}*/
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(" ", congo.ColorDefault)
	for i := rand.Intn(33); i > 0; i-- {
		shuffleInt(mem)
	}
	/*	for i := range mem {
			congo.WindowsMap.ByTitle["Log"].WPrint(" "+strconv.Itoa(mem[i]), congo.ColorDefault)
	}*/
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(" ", congo.ColorDefault)
	for i := needToReveal; i > 0; i-- {
		if i < len(mem) {
			//	congo.WindowsMap.ByTitle["Log"].WPrintLn("reveal data: "+strconv.Itoa(mem[i]), congo.ColorDefault)
			choosen := mem[i]
			if canSee[0] != "Spotted" {
				canSee[0] = "Spotted"
				//printLog("...Icon Spotted: "+icon.GetName(), congo.ColorGreen)
				persona.GetFieldOfView().KnownData[icon.GetID()] = canSee
				//	break
			}
			switch choosen {
			/*case 0:
			if target, ok := icon.(IIcon); ok {
				canSee[choosen] = "Spotted"
				printLog("...Icon Spotted: "+target.GetName(), congo.ColorGreen)
				//printLog("..."+target.GetName()+": Last Edit Date = "+target.GetLastEditDate(), congo.ColorGreen)
			}*/
			case 1:
				if target, ok := icon.(IFile); ok {
					canSee[choosen] = target.GetLastEditDate()
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Last Edit Date = "+target.GetLastEditDate(), congo.ColorGreen)
				}
			case 2:
				if target, ok := icon.(IICOnly); ok {
					canSee[choosen] = strconv.Itoa(target.GetMatrixCM())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.(IObj).GetName()+": Matrix Condition Monitor = "+strconv.Itoa(target.GetMatrixCM()), congo.ColorGreen)
				}
			case 3:
				if target, ok := icon.(IFile); ok {
					canSee[choosen] = strconv.Itoa(target.GetDataBombRating())
					printLog("...Data revealed: ", congo.ColorGreen)
					if target.GetDataBombRating() > 0 {
						printLog("..."+target.GetName()+": Databomb Detected", congo.ColorGreen)
						printLog("..."+target.GetName()+": Databomb Rating = "+strconv.Itoa(target.GetDataBombRating()), congo.ColorYellow)
					} else {
						printLog("..."+target.GetName()+": No Databomb Detected", congo.ColorGreen)
					}
				}
			case 4:
				if target, ok := icon.(IHost); ok {
					canSee[choosen] = "IC List Revealed"
					printLog("...Data revealed: ", congo.ColorGreen)
					icLIST := target.GetICState()
					for j := range icLIST.icName {
						congo.WindowsMap.ByTitle["Log"].WPrint("..."+target.GetName()+": "+icLIST.icName[j]+" Detected ", congo.ColorGreen)
						if icLIST.icStatus[j] {
							printLog("(Status: Active)", congo.ColorYellow)
						} else {
							printLog("(Status: Passive)", congo.ColorGreen)
						}
					}
				}
			case 5:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = strconv.Itoa(target.GetDeviceRating())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Device Rating = "+strconv.Itoa(target.GetDeviceRating()), congo.ColorGreen)
				}
			case 7:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = strconv.Itoa(target.GetAttack())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Attack = "+strconv.Itoa(target.GetAttack()), congo.ColorGreen)
				}
			case 8:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = strconv.Itoa(target.GetSleaze())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Sleaze = "+strconv.Itoa(target.GetSleaze()), congo.ColorGreen)
				}
			case 9:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = strconv.Itoa(target.GetDataProcessing())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Data Processing = "+strconv.Itoa(target.GetDataProcessing()), congo.ColorGreen)
				}
			case 10:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = strconv.Itoa(target.GetFirewall())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Firewall = "+strconv.Itoa(target.GetFirewall()), congo.ColorGreen)
				}
			case 12:
				if target, ok := icon.(IFile); ok {
					canSee[choosen] = strconv.Itoa(target.GetEncryptionRating())
					printLog("...Data revealed: ", congo.ColorGreen)
					if target.GetDataBombRating() > 0 {
						printLog("..."+target.GetName()+": File encryption detected", congo.ColorGreen)
						printLog("..."+target.GetName()+": File encryption = "+strconv.Itoa(target.GetEncryptionRating()), congo.ColorYellow)
					} else {
						printLog("..."+target.GetName()+": No file encryption Detected", congo.ColorGreen)
					}
				}
			case 13:
				if target, ok := icon.(IHost); ok {
					canSee[choosen] = target.GetGrid().GetGridName()
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Located in "+target.GetGrid().GetGridName(), congo.ColorGreen)
				}
			case 15:
				if target, ok := icon.(IFile); ok {
					canSee[choosen] = strconv.Itoa(target.GetSize())
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": File size evaluated", congo.ColorGreen)
					printLog("..."+target.GetName()+": File size = "+strconv.Itoa(target.GetSize())+" Mp", congo.ColorYellow)
				}
			case 18:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = target.GetOwner().GetName()
					printLog("...Data revealed: ", congo.ColorGreen)
					printLog("..."+target.GetName()+": Owner = "+target.GetOwner().GetName(), congo.ColorGreen)
				}

			default:
			}
			if len(mem) > i+1 {
				mem = append(mem[:i], mem[i+1:]...)
			}
			needToReveal--
			persona.GetFieldOfView().KnownData[icon.GetID()] = canSee
		}
	}
	/*allInfo := src.canSee.KnownData[trg.GetID()]
	allInfo[0] = "Spotted" //Spot
	allInfo[1] = "Unknown" //EditDate - File
	allInfo[2] = "Unknown" //MCM
	allInfo[3] = "Unknown" //DataBomb
	allInfo[4] = "Unknown" //ProgramList
	allInfo[5] = "Unknown" //DeviceRating
	allInfo[6] = "Unknown" //ComMode - Persona
	allInfo[7] = "Unknown" //AtT
	allInfo[8] = "Unknown" //SLZ
	allInfo[9] = "Unknown" //DTPRC
	allInfo[10] = "Unknown" //FRW
	allInfo[11] = "Unknown" //uType - no need
	allInfo[12] = "Unknown" //Encrypt - file
	allInfo[13] = "Unknown" //Grid - obj
	allInfo[14] = "Unknown" //Proxy - возможно не надо
	allInfo[15] = "Unknown" //Size - file
	allInfo[16] = "Unknown" //LastAction - icon/  time
	allInfo[17] = "Unknown" //Marks - howmany - ID
	allInfo[18] = "Unknown" //owner   / IIcon */
	return canSee
}

//GetComm -
func GetComm() []string {
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")
	return comm
}
