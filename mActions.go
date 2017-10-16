package main

import (
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
	MActions.MActionMap["UNLOAD_PROGRAM"] = UnloadProgram
	MActions.MActionMap["SET_DATABOMB"] = SetDatabomb
	MActions.MActionMap["SWAP_PROGRAMS"] = SwapPrograms
	MActions.MActionMap["LONGACT"] = LongAct
	MActions.MActionMap["WAIT"] = Wait
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		suc1 = 0
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Patrol succeses: "+strconv.Itoa(suc1), congo.ColorRed)
	//iconInSilentMode := false
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "No Alert" && patrolIC.actionReady == 0 {
		for o := range objectList {
			if obj, ok := objectList[o].(IPersona); ok {
				obj.GetName()
				icon := *objectList[o].(*TPersona)
				if icon.GetHost().name == src.(*TIC).GetHost().name {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Target id: "+strconv.Itoa(icon.GetID()), congo.ColorRed)
					if icon.GetSilentRunningMode() == true {
						if icon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName()+" attack detected...", congo.ColorGreen)
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...attack target: "+icon.GetName(), congo.ColorGreen)
						}
						dp2 := icon.GetLogic() + icon.GetSleaze()
						suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
		for o := range objectList {
			if obj, ok := objectList[o].(IPersona); ok {
				obj.GetName()
				icon := *objectList[o].(*TPersona)
				if icon.GetHost().name == src.(*TIC).GetHost().name {
					if icon.GetSilentRunningMode() == true {
						dp2 := icon.GetLogic() + icon.GetSleaze()
						suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
					suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
					if focusIcon.GetSimSence() == "Hot-SIM VR" || focusIcon.GetSimSence() == "Cold-SIM VR" {
						biofeedbackDamage = realDamage
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "Hot-SIM VR" {
						focusIcon.ReceivePhysBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "Cold-SIM VR" {
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
					if focusIcon.GetSimSence() == "Hot-SIM VR" || focusIcon.GetSimSence() == "Cold-SIM VR" {
						biofeedbackDamage = realDamage
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "Hot-SIM VR" {
						focusIcon.ReceivePhysBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "Cold-SIM VR" {
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
					if focusIcon.GetSimSence() == "Hot-SIM VR" || focusIcon.GetSimSence() == "Cold-SIM VR" {
						biofeedbackDamage = realDamage
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "Hot-SIM VR" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "Cold-SIM VR" {
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
					if focusIcon.GetSimSence() == "Hot-SIM VR" || focusIcon.GetSimSence() == "Cold-SIM VR" {
						biofeedbackDamage = netHits + focusIcon.markSet.MarksFrom[host.GetID()]
					} else {
						biofeedbackDamage = 0
					}
					realBiofeedbackDamage := focusIcon.ResistBiofeedbackDamage(biofeedbackDamage)
					if focusIcon.GetSimSence() == "Hot-SIM VR" {
						focusIcon.ReceiveStunBiofeedbackDamage(realBiofeedbackDamage)
					} else if focusIcon.GetSimSence() == "Cold-SIM VR" {
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
						suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
				suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
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

	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIC); ok {
			if objectList[i].(IIC).GetID() == src.(IIC).GetID() {
				objectList[i] = attacker
			}
		}
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

	}
	return actionIsGood, mActionName
}

func doAction(s string) bool {
	if SourceIcon == nil {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("SourceIcon = nil. Попробуйте вбить команду еще раз и сообщите разработчику", congo.ColorRed)
		return false
	}
	if val, ok := MActions.MActionMap[s]; ok {
		val.(func(IObj, IObj))(SourceIcon, TargetIcon)
		return true
	}
	draw()
	return false
}

//BruteForce0 -
func BruteForce0(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon

	//attMod, defMod := getModifiers(src, trg)
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	//может быть чем угодно.
	if trg, ok := trg.(*TDevice); ok {
		dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetAttack()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		if netHits > 0 {
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Source put MARK", congo.ColorYellow)
		} else {
			src.(IPersona).ReceiveMatrixDamage(-netHits)
			//src.(IPersona).SetMatrixCM(src.(IPersona).GetMatrixCM() + netHits)
		}

	}
	if trg, ok := trg.(*THost); ok {
		dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetAttack()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		gridOV := src.(IPersona).GetGrid()
		src.(IPersona).SetGrid(gridOV)

		if netHits > 0 {
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			if trg.alert == "No Alert" {
				trg.alert = "Passive Alert"
			}
			if src.(IPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("MARK was successfuly planted!", congo.ColorGreen)
			}
		} else {
			src.(IPersona).ReceiveMatrixDamage(-netHits)
			//src.(IPersona).SetMatrixCM(src.(IPersona).GetMatrixCM() + netHits)
		}

	}
	if trg, ok := trg.(*TGrid); ok {
		//panic(0)
		dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating()
		limit := src.(IPersona).GetAttack()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		if netHits > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Action succeeded.", congo.ColorGreen)
			src.(IPersona).SetGrid(*trg)
		}
		addOverwatchScore(suc2)

	}

	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
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

//BruteForce -
func BruteForce(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	persona := src.(IPersona)
	isComplexAction()

	text := command
	text = formatString(text)
	text = cleanText(text)
	//comm := strings.SplitN(text, ">", 5)
	attMod := 0
	/*if comm[3] == "2" && len(comm) > 4 {
		attMod = -4
	}
	if comm[3] == "3" && len(comm) > 4 {
		attMod = -10
	}*/

	dp1 := persona.GetCyberCombatSkill() + persona.GetLogic() + attMod
	limit := persona.GetAttack()

	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	printLog("Initiating Brute Force sequence...", congo.ColorGreen)

	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...error: Encryption protocol glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		printLog("...error: Encryption protocol critical failure", congo.ColorRed)
	}
	printLog("..."+persona.GetName()+": "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)
	if icon, ok := trg.(IIcon); ok {
		host := icon.GetHost()
		dp2 := icon.GetDeviceRating() + icon.GetFirewall()
		printLog(strconv.Itoa(icon.GetDeviceRating())+" Device Rating", congo.ColorDefault)
		printLog(strconv.Itoa(icon.GetFirewall())+" Firewall", congo.ColorDefault)
		suc2, rgl, rcgl := simpleTest(dp2, 1000, 0)
		addOverwatchScore(suc2)
		if rgl == true {
			suc1++
		}
		if rcgl == true {
			addOverwatchScore(-dp2)
		}
		netHits = suc1 - suc2
		printLog(icon.GetName(), congo.ColorDefault)
		printLog(strconv.Itoa(icon.GetDeviceRating())+" File DR", congo.ColorDefault)
		if netHits > 0 {
			placeMARK(persona, icon)
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
	}
	if host, ok := trg.(IHost); ok {
		printLog("Diong HOST"+host.GetName(), congo.ColorDefault)
		host.(IIcon).GetName()
	}

	//attMod, defMod := getModifiers(src, trg)
	//isComplexAction() // есть вероятность что стрельнет механизм возврата
	//может быть чем угодно.
	/*if trg, ok := trg.(*TDevice); ok {
		dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetAttack()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		if netHits > 0 {
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Source put MARK", congo.ColorYellow)
		} else {
			src.(IPersona).ReceiveMatrixDamage(-netHits)
			//src.(IPersona).SetMatrixCM(src.(IPersona).GetMatrixCM() + netHits)
		}

	}
	if trg, ok := trg.(*THost); ok {
		dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetAttack()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		gridOV := src.(IPersona).GetGrid()
		src.(IPersona).SetGrid(gridOV)

		if netHits > 0 {
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			if trg.alert == "No Alert" {
				trg.alert = "Passive Alert"
			}
			if src.(IPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("MARK was successfuly planted!", congo.ColorGreen)
			}
		} else {
			src.(IPersona).ReceiveMatrixDamage(-netHits)
			//src.(IPersona).SetMatrixCM(src.(IPersona).GetMatrixCM() + netHits)
		}

	}
	if trg, ok := trg.(*TGrid); ok {
		//panic(0)
		dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating()
		limit := src.(IPersona).GetAttack()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		if netHits > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Action succeeded.", congo.ColorGreen)
			src.(IPersona).SetGrid(*trg)
		}
		addOverwatchScore(suc2)

	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
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
	}*/
	endAction()
}

//CheckOverwatchScore -
func CheckOverwatchScore(src IObj, trg IObj) {
	src = SourceIcon.(IPersona)
	icon := src.(IPersona)
	grid := src.(IPersona).GetGrid()
	dp1 := src.(IPersona).GetElectronicSkill() + src.(IPersona).GetLogic() //+ attMod
	limit := icon.GetSleaze()
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	if gl {
		addOverwatchScore(8)
	}
	if cgl {
		addOverwatchScore(40)
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Checking Overwatch Score...", congo.ColorGreen)
	hold()

	congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+grid.GetGridName(), congo.ColorGreen)
	hold()
	isSimpleAction()
	congo.WindowsMap.ByTitle["Log"].WPrintLn("..."+strconv.Itoa(suc1), congo.ColorGreen)
	dp2 := 6
	suc2, dgl, dcgl := simpleTest(dp2, 1000, 0)
	if dgl {
		dgl = false
	}
	if dcgl {
		dcgl = false
	}
	netHits := suc1 - suc2
	if netHits > 0 {
		icon.GetOverwatchScore()
		src.(*TPersona).grid.SetLastSureOS(icon.GetOverwatchScore())
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...current OS = "+strconv.Itoa(src.(*TPersona).grid.GetLastSureOS()), congo.ColorGreen)
		hold()
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...failed", congo.ColorYellow)
		hold()
	}
	addOverwatchScore(suc2)

	endAction()
}

//CrackFile -
func CrackFile(src IObj, trg IObj) {
	src = SourceIcon.(*TPersona)
	trg = TargetIcon
	if checkMarks(1) == false {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("ACCESS DENIDED", congo.ColorRed)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Not Enough Marks on "+trg.(*TFile).GetName(), congo.ColorYellow)
	} else { //выполняем само действие
		attMod, defMod := getModifiers(src, trg)
		if trg, ok := trg.(IFile); ok {
			dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() + attMod
			dp2 := trg.GetEncryptionRating()*2 + defMod
			limit := src.(IPersona).GetAttack()
			suc1, gl, cgl := simpleTest(dp1, limit, 0)
			suc2, glt, cglt := simpleTest(dp2, 1000, 0)
			if gl == true {
				addOverwatchScore(2)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
			}
			if cgl == true {
				addOverwatchScore(8)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
			}
			if glt == true {
				addOverwatchScore(-suc2)
			}
			if cglt == true {
				src.(*TPersona).grid.SetOverwatchScore(0)
			}
			//Тут надо остановиться и спросить про перебросс

			netHits := suc1 - suc2
			addOverwatchScore(suc2)
			if trg.GetDataBombRating() > 0 {
				src.(*TPersona).TriggerDataBomb(trg.GetDataBombRating())
				trg.SetDataBombRating(0)
				canSee := src.(*TPersona).canSee.KnownData[trg.GetID()]
				canSee[3] = strconv.Itoa(trg.GetDataBombRating()) // 3- отвечает за рейтинг бомбы
				src.(*TPersona).canSee.KnownData[trg.GetID()] = canSee
			}

			if netHits > 0 {
				trg.SetEncryptionRating(0)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("File encryption disabled...", congo.ColorGreen)
				canSee := src.(*TPersona).canSee.KnownData[trg.GetID()]
				canSee[12] = strconv.Itoa(trg.GetDataBombRating()) // 12- отвечает за Encryption
				src.(*TPersona).canSee.KnownData[trg.GetID()] = canSee
			} else {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Failure! File encryption is not disabled...", congo.ColorGreen)
			}
		}

	}

	endAction()
}

//DataSpike -
func DataSpike(src IObj, trg IObj) {
	src = SourceIcon.(*TPersona)
	trg = TargetIcon
	attMod, defMod := getModifiers(src, trg)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("AttMOD: "+strconv.Itoa(attMod), congo.ColorRed)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("DefMOD: "+strconv.Itoa(defMod), congo.ColorRed)
	dp1 := src.(IPersona).GetCyberCombatSkill() + src.(IPersona).GetLogic() + attMod
	limit := src.(IPersona).GetAttack()
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	if gl == true {
		addOverwatchScore(2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(8)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
	}
	if src.(*TPersona).GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Initiating Data Spike protocol: "+strconv.Itoa(suc1)+" successes", congo.ColorGreen)
	}
	//suc1 = suc1 + 20
	//Тут надо остановиться и спросить про перебросс
	if trg, ok := trg.(*TIC); ok {
		host := trg.GetHost()
		dp2 := trg.GetDeviceRating() + trg.GetFirewall() + defMod
		resistPool := trg.GetDeviceRating() + trg.GetFirewall()

		suc2, glt, cglt := simpleTest(dp2, 1000, 0)
		if glt {
			addOverwatchScore(0 - suc2)
		}
		if cglt {
			resistPool = 0
		}
		netHits := suc1 - suc2
		if src.(*TPersona).GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("//DEBUG: "+strconv.Itoa(netHits)+" netHits", congo.ColorDefault)
		}
		addOverwatchScore(suc2)
		resistHits, rgl, rcgl := simpleTest(resistPool, 999, 0)
		fullDamage := src.(IPersona).GetAttack() + netHits
		realDamage := fullDamage - resistHits
		if realDamage < 0 {
			realDamage = 0
		}
		if rgl == true {
			addOverwatchScore(-suc2)
		}
		if rcgl == true {
			src.(*TPersona).grid.SetOverwatchScore(0)
		}
		if netHits > 0 {
			trg.ReceiveMatrixDamage(realDamage)
			//trg.SetMatrixCM(trg.GetMatrixCM() - realDamage)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" received "+strconv.Itoa(realDamage)+" matrix damage", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" have "+strconv.Itoa(trg.GetMatrixCM())+" matrix boxes left - THIS IS DEBUG MESSAGE", congo.ColorGreen)
			if host.alert == "No Alert" {
				host.alert = "Passive Alert"
			}
		} else {
			src.(IPersona).ReceiveMatrixDamage(-netHits)
			//src.(IPersona).SetMatrixCM(src.(IPersona).GetMatrixCM() + netHits)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(IPersona).GetName()+" received "+strconv.Itoa(netHits)+" unresisted matrix damage", congo.ColorYellow)
		}
	}

	if trg, ok := trg.(*TDevice); ok {
		dp2 := trg.GetDeviceRating() + trg.GetFirewall() + defMod

		resistPool := trg.GetDeviceRating() + trg.GetFirewall()

		suc2, glt, cglt := simpleTest(dp2, 1000, 0)
		if glt {
			addOverwatchScore(0 - suc2)
		}
		if cglt {
			resistPool = 0
		}
		netHits := suc1 - suc2
		addOverwatchScore(suc2)
		resistHits, rgl, rcgl := simpleTest(resistPool, 999, 0)
		fullDamage := src.(IPersona).GetAttack() + netHits
		realDamage := fullDamage - resistHits
		if realDamage < 0 {
			realDamage = 0
		}
		if rgl == true {
			addOverwatchScore(-suc2)
		}
		if rcgl == true {
			src.(*TPersona).grid.SetOverwatchScore(0)
		}
		if netHits > 0 {
			trg.ReceiveMatrixDamage(realDamage)
			//trg.SetMatrixCM(trg.GetMatrixCM() - realDamage)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" received "+strconv.Itoa(realDamage)+" matrix damage", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" have "+strconv.Itoa(trg.GetMatrixCM())+" matrix boxes left", congo.ColorGreen)

		} else {
			src.(IPersona).ReceiveMatrixDamage(-netHits)
			//src.(IPersona).SetMatrixCM(src.(IPersona).GetMatrixCM() + netHits)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(IPersona).GetName()+" received "+strconv.Itoa(netHits)+" unresisted matrix damage", congo.ColorYellow)
		}
	}
	src.(IPersona).SetInitiative(src.(IPersona).GetInitiative() - 10) //Complex Action
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
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

	endAction()

}

//DisarmDataBomb -
func DisarmDataBomb(src IObj, trg IObj) {
	src = SourceIcon.(*TPersona)
	trg = TargetIcon
	icon := SourceIcon.(IPersona)
	attMod, defMod := getModifiers(src, trg)
	isComplexAction()
	dp1 := icon.GetSoftwareSkill() + icon.GetIntuition() + attMod
	limit := src.(IPersona).GetFirewall()
	if icon.CheckRunningProgram("Defuse") {
		dp1 = dp1 + 4
	}
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 0", congo.ColorYellow)
	if trg, ok := trg.(*TFile); ok {
		dp2 := trg.GetDataBombRating()*2 + defMod
		suc2, glt, cglt := simpleTest(dp2, 1000, 0)
		if gl == true {
			addOverwatchScore(dp1 - suc1)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Unexpected trigger initiated....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(dp1 - suc1)
			suc2++
			trg.SetDataBombRating(trg.GetDataBombRating() + 1)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		if glt == true {
			addOverwatchScore(-suc2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb lag detected....", congo.ColorGreen)
		}
		if cglt == true {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb critical lag detected....", congo.ColorGreen)
			suc1++
		}
		//Тут надо остановиться и спросить про перебросс

		netHits := suc1 - suc2
		addOverwatchScore(suc2)

		if netHits > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 2 - success", congo.ColorYellow)
			trg.SetDataBombRating(0)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb Defused...", congo.ColorGreen)
		} else {
			src.(*TPersona).TriggerDataBomb(trg.GetDataBombRating())
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 3 - fail", congo.ColorYellow)
			//
			/*congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb triggered...", congo.ColorRed)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Host Alert triggered...", congo.ColorRed)
			trg.GetHost().SetAlert("Active Alert")
			prgBonus := 0
			if src.(*TPersona).CheckRunningProgram("Armor") {
				prgBonus = prgBonus + 2
			}
			if src.(*TPersona).CheckRunningProgram("Defuse") {
				prgBonus = prgBonus + 4
			}
			resistPool := src.(*TPersona).GetDeviceRating() + src.(*TPersona).GetFirewall() + prgBonus
			resistHits, rgl, rcgl := simpleTest(resistPool, 999, 0)
			//остановиться и перебросить при необходимости
			congo.WindowsMap.ByTitle["Log"].WPrintLn(strconv.Itoa(resistHits)+" of incomming Matrix damage has beeb resisted", congo.ColorGreen)
			fullDamage := xd6Test(trg.GetDataBombRating())
			if rgl == true {
				fullDamage = fullDamage + 2
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Warning!! Firewall error erupted...", congo.ColorYellow)
			}
			if rcgl == true {
				addOverwatchScore(xd6Test(trg.GetDataBombRating()))
				fullDamage = fullDamage + xd6Test(trg.GetDataBombRating())
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Danger!! Critical error erupted...", congo.ColorRed)
			}

			realDamage := fullDamage - resistHits
			if realDamage < 0 {
				realDamage = 0
			}
			src.(*TPersona).ReceiveMatrixDamage(realDamage)
			//src.(*TPersona).SetMatrixCM(src.(*TPersona).GetMatrixCM() - realDamage)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TPersona).GetName()+" receive "+strconv.Itoa(realDamage)+" of matrix damage", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb destroyed", congo.ColorGreen)
			//*/
			trg.SetDataBombRating(0)
			canSee := src.(*TPersona).canSee.KnownData[trg.GetID()]
			canSee[3] = strconv.Itoa(trg.GetDataBombRating()) // 3- отвечает за рейтинг бомбы
			src.(*TPersona).canSee.KnownData[trg.GetID()] = canSee

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

	suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
		suc2, glt, cglt := simpleTest(dp2, 1000, 0)
		if gl == true {
			addOverwatchScore(dp1 - suc1)
			printLog("...error: Unexpected trigger initiated", congo.ColorYellow)
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
		suc1, gl, cgl := simpleTest(dp1, limit, 0)
		if gl == true {
			addOverwatchScore(8)
		}
		if cgl == true {
			addOverwatchScore(40)
		}
		if persona.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Enable Edit mode...", congo.ColorGreen)
			if gl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...unexpected error ocured", congo.ColorYellow)
			}
			if cgl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...critical error erupted", congo.ColorRed)
			}
		}
		if file, ok := trg.(*TFile); ok {
			if checkMarks(1) == false {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("ACCESS DENIDED", congo.ColorRed)
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Not Enough Marks on "+file.GetName(), congo.ColorYellow)
			} else {

				//	congo.WindowsMap.ByTitle["Log"].WPrintLn("Step 1", congo.ColorYellow)

				dp2 := host.GetDeviceRating() + host.GetFirewall()
				suc2, glt, cglt := simpleTest(dp2, 1000, 0)
				if glt == true {
					addOverwatchScore(-suc2)
				}
				if cglt == true {
					suc1++
				}
				if file.GetEncryptionRating() > 0 {
					suc1 = 0
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...error: File Encrypted", congo.ColorGreen)
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
						host.DeleteFile(file)
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

//EnterHost -
func EnterHost(src IObj, trg IObj) {
	//icon := SourceIcon.(*TPersona)
	//checkLinkLock(icon)
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	src = SourceIcon.(*TPersona)
	trg = TargetIcon.(*THost)
	if checkLinkLock(src.(*TPersona)) == true {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("ACTION INTERRUPTED", congo.ColorRed)
		congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(IPersona).GetName()+" is Locked", congo.ColorYellow)
	} else {
		if checkMarks(1) == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("ACCESS DENIDED", congo.ColorRed)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not Enough Marks on "+trg.(*THost).GetName(), congo.ColorYellow)
		} else { //выполняем само действие
			src.(*TPersona).SetHost(trg.(*THost))
		}
	}
	/*if checkMarks(1) == false {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("ACCESS DENIDED", congo.ColorRed)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Not Enough Marks on "+trg.(*THost).GetName(), congo.ColorYellow)
	} else { //выполняем само действие
		src.(*TPersona).SetHost(trg.(*THost))
	}
	trg.(*THost).LockIcon(src.(*TPersona))*/
	endAction()
}

//ExitHost -
func ExitHost(src IObj, trg IObj) {

	isComplexAction() // есть вероятность что стрельнет механизм возврата
	src = SourceIcon.(*TPersona)
	icon := SourceIcon.(*TPersona)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Leaving host... ", congo.ColorGreen)
	if checkLinkLock(icon) == true && src.(IObj).GetFaction() == player.GetFaction() {
		src.(IPersona).SetHost(src.(IPersona).GetHost())
		congo.WindowsMap.ByTitle["Log"].WPrintLn("ACTION INTERRUPTED", congo.ColorRed)
		congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(IPersona).GetName()+" is Locked", congo.ColorYellow)
	} else {
		src.(IPersona).SetHost(Matrix)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...successful", congo.ColorGreen)
	}
	//trg = TargetIcon.(*THost)
	// нужен механизм LINK LOCK
	//src.(IPersona).SetHost(Matrix)
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
		suc1, gl, cgl := simpleTest(dp1, limit, 0)
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
			suc2, dgl, dcgl := simpleTest(dp2, 999, 0)
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
			suc2, dgl, dcgl := simpleTest(dp2, 999, 0)
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
		src.(IPersona).SetGrid(*grid)
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

//HackOnTheFly -
func HackOnTheFly(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	//attMod, defMod := getModifiers(src, trg)
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	//может быть чем угодно.
	prgBonus := 0
	if src.(*TPersona).CheckRunningProgram("Exploit") {
		prgBonus = prgBonus + 2
	}
	dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod

	if trg, ok := trg.(*TDevice); ok {
		//		dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod
		//dp1 = Hacking + Logic[Sleaze]

		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetSleaze()
		congo.WindowsMap.ByTitle["Log"].WPrintLn("LIMIT="+strconv.Itoa(src.(*TPersona).GetSleaze()), congo.ColorYellow)
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		if netHits > 0 {
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on "+trg.GetName()+" was successfuly planted!", congo.ColorGreen)
		} else {
			src.(*TPersona).markSet.MarksFrom[trg.GetID()] = src.(*TPersona).markSet.MarksFrom[trg.GetID()] + 1
			if src.(*TPersona).markSet.MarksFrom[trg.GetID()] > 3 {
				src.(*TPersona).markSet.MarksFrom[trg.GetID()] = 3
			}
		}

	}
	if trg, ok := trg.(*TDevice); ok {
		dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod
		//dp1 = Hacking + Logic[Sleaze]

		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetSleaze()
		congo.WindowsMap.ByTitle["Log"].WPrintLn("LIMIT="+strconv.Itoa(src.(*TPersona).GetSleaze()), congo.ColorYellow)
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		if netHits > 0 {
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on "+trg.GetName()+" was successfuly planted!", congo.ColorGreen)
		} else {
			src.(*TPersona).markSet.MarksFrom[trg.GetID()] = src.(*TPersona).markSet.MarksFrom[trg.GetID()] + 1
			if src.(*TPersona).markSet.MarksFrom[trg.GetID()] > 3 {
				src.(*TPersona).markSet.MarksFrom[trg.GetID()] = 3
			}
		}

	}

	if trg, ok := trg.(*THost); ok {
		//host := trg
		dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		limit := src.(IPersona).GetSleaze()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		addOverwatchScore(suc2)
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		gridOV := src.(IPersona).GetGrid()
		src.(IPersona).SetGrid(gridOV)

		if netHits > 0 {
			needToReveal := netHits / 2
			if src, ok := src.(*TPersona); ok {
				canSee := src.canSee.KnownData[trg.GetID()]
				scanableForHost := make([]int, 0, 9) // 9-количество данных которые можно собрать для этого типа
				for i := 0; i < len(canSee); i++ {
					if i == 0 && canSee[i] != "Spotted" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 4 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 5 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 7 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 8 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 9 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 10 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 11 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					} else if i == 13 && canSee[i] == "Unknown" {
						scanableForHost = append(scanableForHost, i)
					}
				}
				for i := needToReveal; i > 0; i-- {
					if len(scanableForHost) > 0 && needToReveal > 0 {
						shuffleInt(scanableForHost)
						choosen := scanableForHost[0]
						switch scanableForHost[0] {
						case 0:
							canSee[choosen] = "Spotted"
						case 4:
							if src.GetHost().name == trg.GetName() {
								canSee[choosen] = "All IC List"
							}
						case 5:
							//if src.GetHost() == host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetDeviceRating())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Host rating "+strconv.Itoa(trg.GetDeviceRating()), congo.ColorGreen)
						//}
						case 7:
							if src.GetHost().name == trg.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetAttack())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Attack rating "+strconv.Itoa(trg.GetAttack()), congo.ColorGreen)
							}
						case 8:
							if src.GetHost().name == trg.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetSleaze())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Sleaze rating "+strconv.Itoa(trg.GetSleaze()), congo.ColorGreen)
							}
						case 9:
							if src.GetHost().name == trg.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetDataProcessing())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Data Processing rating "+strconv.Itoa(trg.GetDataProcessing()), congo.ColorGreen)
							}
						case 10:
							if src.GetHost().name == trg.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetFirewall())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Firewall rating "+strconv.Itoa(trg.GetFirewall()), congo.ColorGreen)
							}
						case 11:
							if src.GetHost().name == trg.GetName() {
								canSee[choosen] = trg.GetType()
							}
						case 13:
							canSee[choosen] = trg.grid.GetGridName()
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" is located in "+trg.grid.GetGridName(), congo.ColorGreen)
						default:
						}
						src.canSee.KnownData[trg.GetID()] = canSee
						if len(scanableForHost) > 0 {
							scanableForHost = append(scanableForHost[:0], scanableForHost[1:]...)
						}
						needToReveal--
						//a = append(a[:i], a[i+1:]...)
					}
				}
			}
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			if src.(IPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on "+trg.GetName()+" was successfuly planted!", congo.ColorGreen)
			}
		} else {
			src.(*TPersona).markSet.MarksFrom[trg.GetID()] = src.(*TPersona).markSet.MarksFrom[trg.GetID()] + 1
			if src.(*TPersona).markSet.MarksFrom[trg.GetID()] > 3 {
				src.(*TPersona).markSet.MarksFrom[trg.GetID()] = 3
			}
			if src.(*TPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on "+src.(*TPersona).GetName()+" was planted!", congo.ColorRed)
				/*congo.WindowsMap.ByTitle["Log"].WPrintLn("//DEBUG: "+src.(*TPersona).GetName()+" was Locked by "+trg.GetName(), congo.ColorRed)
				trg.LockIcon(src.(*TPersona))*/
			}
			trg.alert = "Active Alert"
		}

	}
	if trg, ok := trg.(*TIC); ok {
		host := trg.GetHost()
		dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating() + trg.GetFirewall()
		//dp2 := trg.
		limit := src.(IPersona).GetSleaze()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		addOverwatchScore(suc2)
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		congo.WindowsMap.ByTitle["Log"].WPrintLn("nethits="+strconv.Itoa(netHits), congo.ColorDefault)
		gridOV := src.(IPersona).GetGrid()
		src.(IPersona).SetGrid(gridOV)

		if netHits > 0 {
			needToReveal := netHits / 2
			congo.WindowsMap.ByTitle["Log"].WPrintLn("needtoreveal="+strconv.Itoa(needToReveal), congo.ColorDefault)
			if src, ok := src.(*TPersona); ok {
				canSee := src.canSee.KnownData[trg.GetID()]

				///////////////////////////////////////////
				scanableForIC := make([]int, 0, 9) // 9-количество данных которые можно собрать для этого типа
				for i := 0; i < len(canSee); i++ {
					if i == 0 && canSee[i] != "Spotted" {
						scanableForIC = append(scanableForIC, i)
					} else if i == 2 && canSee[i] == "Unknown" { //MCM
						scanableForIC = append(scanableForIC, i)
					} else if i == 5 && canSee[i] == "Unknown" { //Rating
						scanableForIC = append(scanableForIC, i)
					} else if i == 7 && canSee[i] == "Unknown" { //att
						scanableForIC = append(scanableForIC, i)
					} else if i == 8 && canSee[i] == "Unknown" { //slz
						scanableForIC = append(scanableForIC, i)
					} else if i == 9 && canSee[i] == "Unknown" { //dtprc
						scanableForIC = append(scanableForIC, i)
					} else if i == 10 && canSee[i] == "Unknown" { //frw
						scanableForIC = append(scanableForIC, i)
					} else if i == 11 && canSee[i] == "Unknown" { //name?
						scanableForIC = append(scanableForIC, i)
					}
					////////////////////////
				}
				//////////////////////////////////////////

				for i := needToReveal; i > 0; i-- {
					if len(scanableForIC) > 0 && needToReveal > 0 {
						shuffleInt(scanableForIC)
						choosen := scanableForIC[0]
						switch scanableForIC[0] {
						case 0:
							canSee[choosen] = "Spotted"
						case 2:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetMatrixCM())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has "+strconv.Itoa(trg.GetMatrixCM())+" Matrix Boxes left", congo.ColorGreen)
							}
						case 5:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetDeviceRating())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Device rating "+strconv.Itoa(trg.GetDeviceRating()), congo.ColorGreen)
							}
						case 7:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetAttack())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Attack rating "+strconv.Itoa(trg.GetAttack()), congo.ColorGreen)
							}
						case 8:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetSleaze())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Sleaze rating "+strconv.Itoa(trg.GetSleaze()), congo.ColorGreen)
							}
						case 9:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetDataProcessing())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Data Processing rating "+strconv.Itoa(trg.GetDataProcessing()), congo.ColorGreen)
							}
						case 10:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetFirewall())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Firewall rating "+strconv.Itoa(trg.GetFirewall()), congo.ColorGreen)
							}
						case 11:
							if src.GetHost().name == host.GetName() {
								canSee[choosen] = trg.GetType()
							}

						default:
						}
						src.canSee.KnownData[trg.GetID()] = canSee
						if len(scanableForIC) > 0 {
							scanableForIC = append(scanableForIC[:0], scanableForIC[1:]...)
						}
						needToReveal--
						//a = append(a[:i], a[i+1:]...)
					}
				}
			}
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			if src.(IPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on "+trg.GetName()+" was successfuly planted!", congo.ColorGreen)
			}
		} else {
			src.(*TPersona).markSet.MarksFrom[trg.GetID()] = src.(*TPersona).markSet.MarksFrom[trg.GetID()] + 1
			if src.(*TPersona).markSet.MarksFrom[trg.GetID()] > 3 {
				src.(*TPersona).markSet.MarksFrom[trg.GetID()] = 3
			}
			if src.(*TPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on "+src.(*TPersona).GetName()+" was planted!", congo.ColorRed)
				/*congo.WindowsMap.ByTitle["Log"].WPrintLn("//DEBUG: "+src.(*TPersona).GetName()+" was Locked by "+trg.GetName(), congo.ColorRed)
				trg.LockIcon(src.(*TPersona))*/
			}
			//trg.alert = "Active Alert"
		}

	}
	if trg, ok := trg.(*TGrid); ok {
		//panic(0)
		dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod
		dp2 := trg.GetDeviceRating()
		limit := src.(IPersona).GetSleaze()
		suc1, suc2, gl, cgl := opposedTest(dp1, dp2, limit)
		netHits := suc1 - suc2
		addOverwatchScore(suc2)
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		if netHits > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Action succeeded.", congo.ColorGreen)
			src.(IPersona).SetGrid(*trg)
			addOverwatchScore(suc2)
		} else {
			addOverwatchScore(suc2)
		}

	}

	if trg, ok := trg.(*TFile); ok {
		dp1 := src.(IPersona).GetHackingSkill() + src.(IPersona).GetLogic() // + attMod
		//dp1 = Hacking + Logic[Sleaze]
		trg.GetDataBombRating()
		//dp2 := trg.host.GetDeviceRating() + trg.host.GetFirewall() +
		//dp2 := trg.GetHost()
		limit := src.(IPersona).GetSleaze()
		suc1, gl, cgl := simpleTest(dp1, limit, 0)
		var host THost
		for i := range objectList {
			if filesHost, ok := objectList[i].(*THost); ok {
				if filesHost.name == src.(IPersona).GetHost().name {
					host = *filesHost
				}
			}
		}
		dp2 := host.GetDeviceRating() + host.GetFirewall()
		suc2, _, _ := simpleTest(dp2, 1000, 0)
		netHits := suc1 - suc2
		if gl == true {
			addOverwatchScore(2)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error! Encryption failed....", congo.ColorYellow)
		}
		if cgl == true {
			addOverwatchScore(8)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Critical Error Erupted....", congo.ColorRed)
		}
		addOverwatchScore(suc2)
		if netHits > 0 {
			needToReveal := netHits / 2
			if src, ok := src.(*TPersona); ok {
				canSee := src.canSee.KnownData[trg.GetID()]
				scanableForFile := make([]int, 0, 6)
				for i := 0; i < len(canSee); i++ {
					if i == 0 && canSee[i] != "Spotted" {
						scanableForFile = append(scanableForFile, i)
					} else if i == 1 && canSee[i] == "Unknown" { //Last Edit
						scanableForFile = append(scanableForFile, i)
					} else if i == 3 && canSee[i] == "Unknown" { //Databomb Rating
						scanableForFile = append(scanableForFile, i)
					} else if i == 12 && canSee[i] == "Unknown" { //Encryption
						scanableForFile = append(scanableForFile, i)
					} else if i == 13 && canSee[i] == "Unknown" { //
						scanableForFile = append(scanableForFile, i)
					} else if i == 15 && canSee[i] == "Unknown" { //
						scanableForFile = append(scanableForFile, i)
					}
				}
				for i := needToReveal; i > 0; i-- {
					if len(scanableForFile) > 0 && needToReveal > 0 {
						shuffleInt(scanableForFile)
						choosen := scanableForFile[0]
						switch scanableForFile[0] {
						case 0:
							canSee[choosen] = "Spotted"
						case 1:
							if src.host.GetName() == trg.host.GetName() {
								//canSee[choosen] = "EditDate revealed"
								canSee[choosen] = trg.GetLastEditDate()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("Last edit date: "+canSee[choosen], congo.ColorGreen)
							}
						case 3:
							if src.host.GetName() == trg.host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetDataBombRating())
								if trg.GetDataBombRating() > 0 {
									congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb rating "+strconv.Itoa(trg.GetDataBombRating())+" detected on "+trg.GetName(), congo.ColorGreen)
								} else {
									congo.WindowsMap.ByTitle["Log"].WPrintLn("No databomb detected on "+trg.GetName(), congo.ColorGreen)
								}
							}
						case 12:
							if src.host.GetName() == trg.host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetEncryptionRating())
								if trg.GetEncryptionRating() > 0 {
									congo.WindowsMap.ByTitle["Log"].WPrintLn("Encryption rating "+strconv.Itoa(trg.GetEncryptionRating())+" detected on "+trg.GetName(), congo.ColorGreen)
								} else {
									congo.WindowsMap.ByTitle["Log"].WPrintLn("No encryption detected on "+trg.GetName(), congo.ColorGreen)
								}
							}
						case 13:
							if src.host.GetName() == trg.host.GetName() {
								canSee[choosen] = trg.GetGridName()
								//congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName() + "Reveal target's GRID", congo.ColorYellow) - не знаю как литературно сформулировать и нужно ли. Имеет смысл заменить это на Истинное имя файла
							}
						case 15:
							if src.host.GetName() == trg.host.GetName() {
								canSee[choosen] = strconv.Itoa(trg.GetSize())
								congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" contains "+strconv.Itoa(trg.GetSize())+" Mp of data", congo.ColorGreen)
							}
						default:
						}
						src.canSee.KnownData[trg.GetID()] = canSee
						if len(scanableForFile) > 0 {
							scanableForFile = append(scanableForFile[:0], scanableForFile[1:]...)
						}
						needToReveal--
						//congo.WindowsMap.ByTitle["Log"].WPrintLn("scanable for File = " + strconv.Itoa(len(scanableForFile)), congo.ColorYellow)
						//a = append(a[:i], a[i+1:]...)
					}
				}
			}
			trg.markSet.MarksFrom[src.(IPersona).GetID()] = trg.markSet.MarksFrom[src.(IPersona).GetID()] + 1
			if trg.markSet.MarksFrom[src.(IPersona).GetID()] > 3 {
				trg.markSet.MarksFrom[src.(IPersona).GetID()] = 3
			}
			if src.(IPersona).GetID() == 0 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("MARK was successfuly planted!", congo.ColorGreen)
			}
		} else {
			src.(*TPersona).markSet.MarksFrom[trg.GetID()] = src.(*TPersona).markSet.MarksFrom[trg.GetID()] + 1
			if src.(*TPersona).markSet.MarksFrom[trg.GetID()] > 3 {
				src.(*TPersona).markSet.MarksFrom[trg.GetID()] = 3
			}
		}

	}

	//присваиваем изменения
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
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
	for i := range gridList {
		if host, ok := gridList[i].(IHost); ok {
			if gridList[i].(IHost).GetID() == src.(IHost).GetID() {
				gridList[i] = host

			}
		}
	}
	endAction()
}

//MatrixPerception -
func MatrixPerception(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	dp1 := src.(IPersona).GetComputerSkill() + src.(IPersona).GetIntuition() // + attMod
	limit := src.(IPersona).GetDataProcessing()
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	if gl {
		addOverwatchScore(2)
	}
	if cgl {
		addOverwatchScore(8)
	}

	if src.(IObj).GetID() == 0 {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Player has "+strconv.Itoa(suc1)+" sucesses...", congo.ColorRed)
	}
	needToReveal := suc1
	if trg, ok := trg.(*TFile); ok {
		if src, ok := src.(*TPersona); ok {
			canSee := src.canSee.KnownData[trg.GetID()]
			scanableForFile := make([]int, 0, 6)
			for i := 0; i < len(canSee); i++ {
				if i == 0 && canSee[i] != "Spotted" {
					scanableForFile = append(scanableForFile, i)
				} else if i == 1 && canSee[i] == "Unknown" { //Last Edit
					scanableForFile = append(scanableForFile, i)
				} else if i == 3 && canSee[i] == "Unknown" { //Databomb Rating
					scanableForFile = append(scanableForFile, i)
				} else if i == 12 && canSee[i] == "Unknown" { //Encryption
					scanableForFile = append(scanableForFile, i)
				} else if i == 13 && canSee[i] == "Unknown" { //
					scanableForFile = append(scanableForFile, i)
				} else if i == 15 && canSee[i] == "Unknown" { //
					scanableForFile = append(scanableForFile, i)
				}
			}
			for i := needToReveal; i > 0; i-- {
				if len(scanableForFile) > 0 && needToReveal > 0 {
					src.GetHost().GetName()
					//src.host.GetName()
					shuffleInt(scanableForFile)
					choosen := scanableForFile[0]
					switch scanableForFile[0] {
					case 0:
						canSee[choosen] = "Spotted"
					case 1:
						if src.GetHost().GetName() == trg.host.GetName() {
							//canSee[choosen] = "EditDate revealed"
							canSee[choosen] = trg.GetLastEditDate()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("Last edit date: "+canSee[choosen], congo.ColorGreen)
						}
					case 3:
						if src.GetHost().GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetDataBombRating())
							if trg.GetDataBombRating() > 0 {
								congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb rating "+strconv.Itoa(trg.GetDataBombRating())+" detected on "+trg.GetName(), congo.ColorGreen)
							} else {
								congo.WindowsMap.ByTitle["Log"].WPrintLn("No databomb detected on "+trg.GetName(), congo.ColorGreen)
							}
						}
					case 12:
						if src.GetHost().GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetEncryptionRating())
							if trg.GetEncryptionRating() > 0 {
								congo.WindowsMap.ByTitle["Log"].WPrintLn("Encryption rating "+strconv.Itoa(trg.GetEncryptionRating())+" detected on "+trg.GetName(), congo.ColorGreen)
							} else {
								congo.WindowsMap.ByTitle["Log"].WPrintLn("No encryption detected on "+trg.GetName(), congo.ColorGreen)
							}
						}
					case 13:
						if src.GetHost().GetName() == trg.host.GetName() {
							canSee[choosen] = trg.GetGridName()
							//congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName() + "Reveal target's GRID", congo.ColorYellow) - не знаю как литературно сформулировать и нужно ли. Имеет смысл заменить это на Истинное имя файла
						}
					case 15:
						if src.GetHost().GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetSize())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" contains "+strconv.Itoa(trg.GetSize())+" Mp of data", congo.ColorGreen)
						}
					default:
					}
					src.canSee.KnownData[trg.GetID()] = canSee
					if len(scanableForFile) > 0 {
						scanableForFile = append(scanableForFile[:0], scanableForFile[1:]...)
					}
					needToReveal--
					//congo.WindowsMap.ByTitle["Log"].WPrintLn("scanable for File = " + strconv.Itoa(len(scanableForFile)), congo.ColorYellow)
					//a = append(a[:i], a[i+1:]...)
				}
			}
		}
	}
	//////////////////////////////////////////////////////////////////////

	if trg, ok := trg.(*THost); ok {
		dp2 := trg.GetDeviceRating() + trg.GetSleaze()
		suc2, dgl, cdgl := simpleTest(dp2, 1000, 0)
		needToReveal = suc1 - suc2
		if dgl {
			needToReveal++
		}
		if cdgl {
			needToReveal = needToReveal + 12
		}
		if needToReveal < 0 {
			needToReveal = 0
		}
		if src, ok := src.(*TPersona); ok {
			canSee := src.canSee.KnownData[trg.GetID()]
			/*allInfo := src.canSee.KnownData[trg.GetID()]
			allInfo[0] = "Spotted" //Spot
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
			allInfo[18] = "Unknown" //owner*/
			scanableForHost := make([]int, 0, 9) // 9-количество данных которые можно собрать для этого типа
			for i := 0; i < len(canSee); i++ {
				if i == 0 && canSee[i] != "Spotted" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 4 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 5 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 7 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 8 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 9 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 10 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 11 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} else if i == 13 && canSee[i] == "Unknown" {
					scanableForHost = append(scanableForHost, i)
				}
				/*if i == 16 && canSee[i] != "Unknown" {
					scanableForHost = append(scanableForHost, i)
				}
				if i == 17 && canSee[i] != "Unknown" {
					scanableForHost = append(scanableForHost, i)
				}
				if i == 18 && canSee[i] != "Unknown" {
					scanableForHost = append(scanableForHost, i)
				} */ //12
			}
			for i := needToReveal; i > 0; i-- {
				if len(scanableForHost) > 0 && needToReveal > 0 {
					shuffleInt(scanableForHost)
					choosen := scanableForHost[0]
					switch scanableForHost[0] {
					case 0:
						canSee[choosen] = "Spotted"
					case 4:
						if src.host.GetName() == trg.GetName() {
							canSee[choosen] = "ICList"
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" IC order revealed", congo.ColorGreen)
						}
					case 5:
						//if src.GetHost() == host.GetName() {f
						canSee[choosen] = strconv.Itoa(trg.GetDeviceRating())
						congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Host rating "+strconv.Itoa(trg.GetDeviceRating()), congo.ColorGreen)
					//}
					case 7:
						if src.host.GetName() == trg.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetAttack())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Attack rating "+strconv.Itoa(trg.GetAttack()), congo.ColorGreen)
						}
					case 8:
						if src.host.GetName() == trg.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetSleaze())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Sleaze rating "+strconv.Itoa(trg.GetSleaze()), congo.ColorGreen)
						}
					case 9:
						if src.host.GetName() == trg.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetDataProcessing())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Data Processing rating "+strconv.Itoa(trg.GetDataProcessing()), congo.ColorGreen)
						}
					case 10:
						if src.host.GetName() == trg.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetFirewall())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has Firewall rating "+strconv.Itoa(trg.GetFirewall()), congo.ColorGreen)
						}
					case 11:
						if src.host.GetName() == trg.GetName() {
							canSee[choosen] = trg.GetType()
						}
					case 13:
						canSee[choosen] = trg.grid.GetGridName()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" is located in "+trg.grid.GetGridName(), congo.ColorGreen)
					default:
					}
					src.canSee.KnownData[trg.GetID()] = canSee
					if len(scanableForHost) > 0 {
						scanableForHost = append(scanableForHost[:0], scanableForHost[1:]...)
					}
					needToReveal--
					//a = append(a[:i], a[i+1:]...)
				}
			}

		}
	}

	///////////////////////////////////////////////////////////////////////
	if trg, ok := trg.(*TIC); ok {
		if src, ok := src.(*TPersona); ok {
			canSee := src.canSee.KnownData[trg.GetID()]
			scanableForIC := make([]int, 0, 9) // 9-количество данных которые можно собрать для этого типа
			for i := 0; i < len(canSee); i++ {
				if i == 0 && canSee[i] != "Spotted" {
					scanableForIC = append(scanableForIC, i)
				} else if i == 2 && canSee[i] == "Unknown" { //MCM
					scanableForIC = append(scanableForIC, i)
				} else if i == 5 && canSee[i] == "Unknown" { //Rating
					scanableForIC = append(scanableForIC, i)
				} else if i == 7 && canSee[i] == "Unknown" { //att
					scanableForIC = append(scanableForIC, i)
				} else if i == 8 && canSee[i] == "Unknown" { //slz
					scanableForIC = append(scanableForIC, i)
				} else if i == 9 && canSee[i] == "Unknown" { //dtprc
					scanableForIC = append(scanableForIC, i)
				} else if i == 10 && canSee[i] == "Unknown" { //frw
					scanableForIC = append(scanableForIC, i)
				} else if i == 11 && canSee[i] == "Unknown" { //name?
					scanableForIC = append(scanableForIC, i)
				}
				////////////////////////
			}
			for i := needToReveal; i > 0; i-- {
				if len(scanableForIC) > 0 && needToReveal > 0 {
					shuffleInt(scanableForIC)
					choosen := scanableForIC[0]
					switch scanableForIC[0] {
					case 0:
						canSee[choosen] = "Spotted"
					case 2:
						if src.host.GetName() == trg.host.GetName() {
							//canSee[choosen] = "EditDate revealed"
							canSee[choosen] = strconv.Itoa(trg.GetMatrixCM())
							congo.WindowsMap.ByTitle["Log"].WPrintLn("Condition Monitor: "+canSee[choosen], congo.ColorGreen)
						}
					case 5:
						if src.host.GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetDeviceRating())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has rating "+canSee[choosen], congo.ColorGreen)

						}
					case 7:
						if src.host.GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetAttack())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has attack rating "+canSee[choosen], congo.ColorGreen)
						}
					case 8:
						if src.host.GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetSleaze())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has sleaze rating "+canSee[choosen], congo.ColorGreen)
						}
					case 9:
						if src.host.GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetDataProcessing())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has data processing rating "+canSee[choosen], congo.ColorGreen)
						}
					case 10:
						if src.host.GetName() == trg.host.GetName() {
							canSee[choosen] = strconv.Itoa(trg.GetFirewall())
							congo.WindowsMap.ByTitle["Log"].WPrintLn(trg.GetName()+" has data firewall rating "+canSee[choosen], congo.ColorGreen)
						}
					default:
					}
					src.canSee.KnownData[trg.GetID()] = canSee
					if len(scanableForIC) > 0 {
						scanableForIC = append(scanableForIC[:0], scanableForIC[1:]...)
					}
					needToReveal--
					//congo.WindowsMap.ByTitle["Log"].WPrintLn("scanable for File = " + strconv.Itoa(len(scanableForFile)), congo.ColorYellow)
					//a = append(a[:i], a[i+1:]...)
				}
			}
		}
	}
	///////////////////////////////////////////////////////////////////////

	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
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
	isComplexAction()
	endAction()

}

//MatrixSearch -
func MatrixSearch(src IObj, trg IObj) {
	seeker := SourceIcon.(IPersona)
	host := seeker.(IPersona).GetHost()
	//text := TargetIcon.(string)
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 4)

	if comm[2] == "HOST" {
		var hostName string
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Search Host Initiated...", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Entering global registry...", congo.ColorGreen)
		if len(comm) < 4 {
			hostName = "Unknown Host " + strconv.Itoa(id)
		} else {
			hostName = comm[3]
			hostName = strings.ToLower(hostName)
			hostName = strings.Replace(hostName, "_", " ", -1)
			hostName = strings.Title(hostName)
			//примерно тут надо будет запускать поиск на наличие хоста в библиотеке
		}
		if HostExist(hostName) {

			//			congo.WindowsMap.ByTitle["Log"].WPrintLn(gridList[0].(*TGrid).GetName(), congo.ColorGreen)
			ImportHostFromDB(hostName)

			//player.grid.NewHost(name, rating)
		} else {
			player.grid.NewHost(hostName, 0) // -создаем всегда третий хост
		}

	}
	if comm[2] == "FILE" {
		//host := src.(IPersona).GetHost()
		var fileName string
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Search File Initiated...", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Entering global registry...", congo.ColorGreen)
		if len(comm) < 4 {
			fileName = "File " + strconv.Itoa(id)
		} else {
			fileName = comm[3]
			fileName = strings.ToLower(fileName)
			fileName = strings.Replace(fileName, "_", " ", -1)
			fileName = strings.Title(fileName)
			//примерно тут надо будет запускать поиск на наличие хоста в библиотеке
		}
		file := host.NewFile(fileName)
		file.SetFileName(host.GetName() + " " + fileName)
		//file.markSet.MarksFrom[seeker.GetID()] = 4
		//gridList[0].(*TGrid).NewHost(hostName, 4) // -создаем всегда третий хост
	}

	isComplexAction()
	endAction()
}

//ScanEnviroment -
func ScanEnviroment(src IObj, trg IObj) {
	src = SourceIcon
	trg = TargetIcon
	isComplexAction()
	dp1 := src.(IPersona).GetComputerSkill() + src.(IPersona).GetIntuition() // + attMod
	limit := src.(IPersona).GetDataProcessing()
	suc1, gl, cgl := simpleTest(dp1, limit, 0)
	if gl {
		addOverwatchScore(2)
	}
	if cgl {
		addOverwatchScore(8)
	}
	if src.(IObj).GetID() == 0 {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Scanning completed...", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(IPersona).GetName()+" has "+strconv.Itoa(suc1)+" sucesses", congo.ColorGreen)
	}
	needToReveal := suc1
	needToReveal++
	needToReveal--
	silentIconsFound := 0

	for o := range objectList {
		if obj, ok := objectList[o].(IFile); ok {
			obj.GetName()
			file := *objectList[o].(*TFile)
			var host THost
			for i := range objectList {
				if filesHost, ok := objectList[i].(*THost); ok {
					if filesHost.GetName() == src.(IPersona).GetHost().name {
						host = *filesHost
					}
				}
			}
			if file.host.GetName() == src.(IPersona).GetHost().name {
				if file.GetSilentRunningMode() {
					silentIconsFound++
				}
				dp2 := host.GetDeviceRating() + host.GetSleaze()
				suc2, dgl, cdgl := simpleTest(dp2, 1000, 0)
				needToReveal = suc1 - suc2
				if dgl {
					needToReveal++
				}
				if cdgl {
					needToReveal = needToReveal + 12
				}
				if needToReveal < 0 {
					needToReveal = 0
				}
				canSee := src.(*TPersona).canSee.KnownData[file.GetID()]

				scanableForFile := make([]int, 0, 6)
				for i := 0; i < len(canSee); i++ {
					if i == 0 && canSee[i] != "Spotted" {
						scanableForFile = append(scanableForFile, i)
					} /* else if i == 1 && canSee[i] == "Unknown" { //Last Edit
						scanableForFile = append(scanableForFile, i)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Data " + strconv.Itoa(i) + " in file " + file.GetName() + " Is SCANABLE", congo.ColorYellow)
					} else if i == 3 && canSee[i] == "Unknown" { //Databomb Rating
						scanableForFile = append(scanableForFile, i)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Data " + strconv.Itoa(i) + " in file " + file.GetName() + " Is SCANABLE", congo.ColorYellow)
					} else if i == 12 && canSee[i] == "Unknown" { //Encryption
						scanableForFile = append(scanableForFile, i)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Data " + strconv.Itoa(i) + " in file " + file.GetName() + " Is SCANABLE", congo.ColorYellow)
					} else if i == 13 && canSee[i] == "Unknown" { //
						scanableForFile = append(scanableForFile, i)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Data " + strconv.Itoa(i) + " in file " + file.GetName() + " Is SCANABLE", congo.ColorYellow)
					} else if i == 15 && canSee[i] == "Unknown" { //
						scanableForFile = append(scanableForFile, i)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Data " + strconv.Itoa(i) + " in file " + file.GetName() + " Is SCANABLE", congo.ColorYellow)
					}*/
				}
				for i := needToReveal; i > 0; i-- {
					if len(scanableForFile) > 0 && needToReveal > 0 {
						/*shuffleInt(scanableForHost)
						shuffleInt(scanableForHost)
						shuffleInt(scanableForHost)
						shuffleInt(scanableForHost)*/

						shuffleInt(scanableForFile)
						choosen := scanableForFile[0]
						switch scanableForFile[0] {
						case 0:
							if src.(*TPersona).GetHost().name == file.GetHost().name {
								canSee[choosen] = "Spotted"
								congo.WindowsMap.ByTitle["Log"].WPrintLn(file.GetName()+" was spotted", congo.ColorGreen)
							}
						/*case 1:
						if src.(*TPersona).GetHost() == file.GetHost() {
							canSee[choosen] = "EditDate revealed"
						}
						case 3:
						if src.(*TPersona).GetHost() == file.GetName() {
							canSee[choosen] = strconv.Itoa(file.GetDataBombRating())
						}
						case 12:
						if src.(*TPersona).GetHost() == file.GetHost() {
							canSee[choosen] = strconv.Itoa(file.GetEncryptionRating())
						}
						case 13:
						if src.(*TPersona).GetHost() == file.GetName() {
							canSee[choosen] = file.GetGridName()
						}
						case 15:
						if src.(*TPersona).GetHost() == file.GetName() {
							canSee[choosen] = strconv.Itoa(file.GetSize())
						}*/
						default:
						}
						src.(*TPersona).canSee.KnownData[file.GetID()] = canSee
						if len(scanableForFile) > 0 {
							scanableForFile = append(scanableForFile[:0], scanableForFile[1:]...)
						}
						needToReveal--
						//a = append(a[:i], a[i+1:]...)
					}
				}
			}
		}

	}
	if silentIconsFound > 0 {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Warning!!! There are tolal of "+strconv.Itoa(silentIconsFound)+" Icons running in Silent Mode!!", congo.ColorYellow)
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("No icons running in Silent Mode found", congo.ColorGreen)
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
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
	endAction()
}

//SwapAttributes -
func SwapAttributes(src IObj, trg IObj) {
	src = SourceIcon
	if persona, ok := src.(*TPersona); ok {
		//text := TargetIcon.(string)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)

		att1 := 0
		switch comm[2] {
		case "ATTACK":
			att1 = SourceIcon.(*TPersona).GetAttackRaw()
		case "SLEAZE":
			att1 = SourceIcon.(*TPersona).GetSleazeRaw()
		case "DATA_PROCESSING":
			att1 = SourceIcon.(*TPersona).GetDataProcessingRaw()
		case "FIREWALL":
			att1 = SourceIcon.(*TPersona).GetFirewallRaw()
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Attribute1 name is invalid...", congo.ColorYellow)

		}
		att2 := 0
		switch comm[3] {
		case "ATTACK":
			att2 = SourceIcon.(*TPersona).GetAttackRaw()
		case "SLEAZE":
			att2 = SourceIcon.(*TPersona).GetSleazeRaw()
		case "DATA_PROCESSING":
			att2 = SourceIcon.(*TPersona).GetDataProcessingRaw()
		case "FIREWALL":
			att2 = SourceIcon.(*TPersona).GetFirewallRaw()
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Attribute2 name is invalid...", congo.ColorYellow)
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
		}
		if swap1 == true && swap2 == true {
			isFreeAction()
			//присваиваем изменения
			for i := range objectList {
				if attacker, ok := objectList[i].(IIcon); ok {
					if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
						objectList[i] = attacker
					}
				}
			}
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Attribute configuration failed...", congo.ColorYellow)
		}
	}
	endAction()
}

//LoadProgram -
func LoadProgram(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["z"].WClear()
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

//UnloadProgram -
func UnloadProgram(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["z"].WClear()
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

//SwapPrograms -
func SwapPrograms(src IObj, trg IObj) {
	src = SourceIcon
	congo.WindowsMap.ByTitle["z"].WClear()
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

	//text := TargetIcon.(string)
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 4)
	if len(comm) > 2 {
		wt := comm[2]
		wi, _ := strconv.Atoi(wt)
		icon.SetInitiative(icon.GetInitiative() - wi)

	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Wait time unspecified...", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Waiting until end of turn...", congo.ColorDefault)
		src.(IPersona).SetInitiative(0)
	}
	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IPersona); ok {
			if objectList[i].(IIcon).GetID() == src.(IPersona).GetID() {
				objectList[i] = attacker
			}
		}
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

	outIndex := 0
	for _, x := range objectList {
		if objectList[outIndex].(IObj).GetType() != "File" {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Validate Icon: "+objectList[outIndex].(IObj).GetName(), congo.ColorDefault)
		}
		if icon, ok := objectList[outIndex].(IIcon); ok {
			if approveDetetion(icon) == false {
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

	}
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
	src := SourceIcon
	if source, ok := src.(IIcon); ok {
		source.SetInitiative(source.GetInitiative() - 10) //Complex Action
	}
	/*if source, ok := src.(IIC); ok {
		source.SetInitiative(source.GetInitiative() - 10) //Complex Action
	}*/

	//присваиваем изменения
	for i := range objectList {
		if attacker, ok := objectList[i].(IIcon); ok {
			if objectList[i].(IIcon).GetID() == src.(IIcon).GetID() {
				objectList[i] = attacker
			}
		}
	}
}

func isFreeAction() {
	src := SourceIcon
	src.(IPersona).SetInitiative(src.(IPersona).GetInitiative() - 3) //Free Action
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

/*
if comm[2] == "ALL" {
		for o := range objectList {
		if objectList[o].(IObj).GetType() == "Host" {
			host := objectList[o].(*THost)
			dp2 := host.GetDeviceRating() + host.GetSleaze()
			suc2, glt, cglt := simpleTest(dp2, 1000, 0)
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
