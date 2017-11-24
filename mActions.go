package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/ConGo/congo"
)

//MActions - Holder for all actions for "doAction(string)" function
var MActions TMActionMap

//TMActionMap - key is string, value is the function than must be executed
type TMActionMap struct {
	MActionMap map[string]interface{}
}

//InitMatrixActionMap - creating MActions and adding it's values
func InitMatrixActionMap() {
	MActions = TMActionMap{}
	MActions.MActionMap = map[string]interface{}{}
	MActions.MActionMap["BRUTE_FORCE"] = BruteForce
	MActions.MActionMap["SWITCH_INTERFACE_MODE"] = SwitchInterfaceMode
	MActions.MActionMap["SILENT_MODE"] = SwitchSilentMode
	MActions.MActionMap["CHECK_OVERWATCH_SCORE"] = CheckOverwatchScore
	MActions.MActionMap["CRACK_FILE"] = CrackFile
	MActions.MActionMap["DATA_SPIKE"] = DataSpike
	MActions.MActionMap["DISARM_DATABOMB"] = DisarmDataBomb
	MActions.MActionMap["EDIT_FILE"] = Edit
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
}

//LongAct - test function, no meaning
func LongAct(src IObj, trg IObj) {
	icon := src.(IPersona)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Start LongAct by "+icon.GetName(), congo.ColorDefault)
	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+"`s searchLen = "+strconv.Itoa(icon.GetLongAct()), congo.ColorDefault)
	endAction()
}

//PatrolICActionArea - patrol scans EVERY icon in the host acording to rules in Data Trails
func PatrolICActionArea(src IObj, trg IObj) {
	patrolIC := src.(IIC)
	host := patrolIC.GetHost()
	dp1 := host.GetDeviceRating() * 2
	limit := host.GetDataProcessing()
	isComplexAction()
	suc1, gl, cgl := simpleTest(patrolIC.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		suc1 = 0
	}
	/////////////////////////////////////////////////////////////////////////
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetHost() != patrolIC.GetHost() {
				continue
			}
			if icon.GetFaction() == player.GetFaction() {
				printLog("..."+icon.GetName()+" is being scanned", congo.ColorGreen)
			}
			dp2 := 0
			netHits := 0
			suc2 := 0
			dgl := false
			dcgl := false
			if icon.GetSilentRunningMode() {
				dp2 = icon.GetSleaze() + icon.GetDeviceRating()
				if persona, ok := icon.(IPersona); ok {
					dp2 = persona.GetSleaze() + persona.GetLogic()
				}
				if icon.GetFaction() == player.GetFaction() {
					printLog("..."+icon.GetName()+" evaiding", congo.ColorGreen)
				}
				suc2, dgl, dcgl = simpleTest(icon.GetID(), dp2, 1000, 0)
			} else {
				dp2 = 0
			}
			//suc2, dgl, dcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			if dgl {
				suc1++
			}
			if dcgl {
				suc1++
			}
			netHits = suc1 - suc2
			if netHits > 0 {
				patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
				if icon.GetFaction() == player.GetFaction() {
					printLog("..."+icon.GetName()+" detected", congo.ColorYellow)
				}
				if icon.GetFaction() != host.GetFaction() {
					host.alert = "Passive Alert"
				}
				host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
			} else {
				if icon.GetFaction() == player.GetFaction() {
					printLog("..."+icon.GetName()+" evaded", congo.ColorGreen)
				}
			}
		}
	}
	endAction()
}

//PatrolICActionTarget - Patrol scanning 1 target
func PatrolICActionTarget(src IObj, trg IObj) {
	patrolIC := src.(IIC)
	host := patrolIC.GetHost()
	dp1 := host.GetDeviceRating() * 2
	limit := host.GetDataProcessing()
	suc1, gl, cgl := simpleTest(patrolIC.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		suc1 = 0
	}
	/////////////////////////////////////////////////////////////////////////
	scan := true
	for scan {
		scan = false
		isComplexAction()
		if icon, ok := trg.(IIcon); ok {
			if !icon.GetSilentRunningMode() {
				if suc1 > 0 {
					if icon.GetFaction() == player.GetFaction() {
						printLog("..."+icon.GetName()+" was affected", congo.ColorYellow)
					}
					patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
					host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
					host.SetAlert("Active Alert")
				}
			} else {
				dp2 := icon.GetSleaze() + icon.GetDeviceRating()
				if persona, ok := icon.(IPersona); ok {
					dp2 = persona.GetSleaze() + persona.GetLogic()
				}
				limit = 1000
				suc2, dgl, dcgl := simpleTest(icon.GetID(), dp2, limit, 0)
				if dgl {
					addOverwatchScoreToTarget(dp2 - suc2)
				}
				if dcgl {
					suc1++
					addOverwatchScoreToTarget(dp2 - suc2)
				}
				netHits := suc1 - suc2
				if netHits > 0 {
					patrolIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
					host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
					host.SetAlert("Active Alert")
					if icon.GetFaction() == player.GetFaction() {
						printLog("..."+icon.GetName()+" was affected", congo.ColorYellow)
					}
				}
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
				dp2 := focusIcon.GetWillpower() + focusIcon.GetDataProcessing()
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
	endAction()
}

//BloodhoundICActionTarget -
func BloodhoundICActionTarget(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = pickIconByName(src.(*TIC).GetLastTargetName())
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
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		for _, obj := range ObjByNames {
			if icon, ok := obj.(IIcon); ok {
				if icon.GetHost().name == src.(*TIC).GetHost().name {
					if icon.GetSilentRunningMode() == true {
						dp2 := icon.GetDeviceRating() + icon.GetSleaze()
						if persona, ok := icon.(IPersona); ok {
							dp2 = persona.GetSleaze() + persona.GetLogic()
						}
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
	endAction()
}

//ICWait - action for all ICes: do nothing
func ICWait(src IObj, trg IObj) {
	src = SourceIcon.(*TIC)
	trg = TargetIcon
	if ic, ok := SourceIcon.(*TIC); ok {
		ic.SetInitiative(ic.GetInitiative() - 10)
	}
	endAction()
}

/////////////////////////////////////////////////////////////

//checkAction - will check user input and if matrix action was spelled correctly
func checkAction(actionName string) (bool, string) {
	actionIsGood := false
	var mActionName string
	switch actionName {
	case "BRUTE_FORCE":
		actionIsGood = true
		mActionName = "BRUTE_FORCE"
		return actionIsGood, mActionName
	case "SILENT_MODE":
		actionIsGood = true
		mActionName = "SILENT_MODE"
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
	case "EDIT_FILE":
		actionIsGood = true
		mActionName = "EDIT_FILE"
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

//doAction - handler for matrix action for both player AND enemies
func doAction(s string) bool {
	if SourceIcon == nil {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("--DEBUG-- Error: SourceIcon = nil. Try again", congo.ColorRed)
		return false
	}
	if val, ok := MActions.MActionMap[s]; ok { // if s string match any Matrix Action than execute action
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
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
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

	var validSpecs []string
	//validSpecs = append(validSpecs, "Brute Force")
	haveSpec, spec := persona.HaveValidSpec(validSpecs)
	if haveSpec {
		printLog("ValidSpec is "+spec, congo.ColorDefault)
	}
	printLog("Initiating Brute Force sequence...", congo.ColorGreen)
	targetList, targetSpec := pickTargets2(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetCyberCombatSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	if haveSpec || targetSpec {
		dp1 = dp1 + 2
		printLog("...Specialization: +"+strconv.Itoa(2)+" op/p", congo.ColorGreen)
	}
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
			netHits := suc1 - suc2
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
			netHits := suc1 - suc2
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
				printLog("..."+file.GetName()+" decrypted", congo.ColorGreen)
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
				printLog("...Databomb Disarmed", congo.ColorGreen)
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
	host := trg.(IFile).GetHost()
	dp1 := icon.GetSoftwareSkill() + icon.GetLogic() // + attMod
	limit := src.(IPersona).GetSleaze()
	suc1, gl, cgl := simpleTest(icon.GetID(), dp1, limit, 0)
	if icon.GetFaction() == player.GetFaction() {
		printLog("Setting up databomb on "+trg.GetName()+"...", congo.ColorGreen)
	}
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

/*//Edit0 - old function. Delete if Edit will receive no bugs
func Edit0(src IObj, trg IObj) {
	comm := GetComm()
	trg = TargetIcon
	host := SourceIcon.(IIcon).GetHost()
	if persona, ok := SourceIcon.(IPersona); ok {
		if persona.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Enable Edit mode", congo.ColorGreen)
		}
		dp1 := persona.GetComputerSkill() + persona.GetLogic() // + attMod
		limit := persona.GetDataProcessing()
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...Hardware limit: "+strconv.Itoa(limit), congo.ColorGreen)
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if gl == true {
			addOverwatchScore(8)
		}
		if cgl == true {
			addOverwatchScore(40)
		}
		if persona.GetFaction() == player.GetFaction() {
			if gl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...unexpected error ocured", congo.ColorYellow)
			}
			if cgl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...critical error erupted", congo.ColorRed)
			}
		}
		if file, ok := trg.(IFile); ok {
			if !checkExistingMarks(persona.GetID(), file.GetID(), 1) {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...ACCESS DENIDED", congo.ColorRed)
			} else {
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
					for i := range comm {
						if comm[i] == "COPY" {
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...copying", congo.ColorGreen)
							hold()
							copy := host.NewFile(file.GetFileName())
							copy.SetFileName("Copy of " + file.GetFileName())
							copy.SetDataBombRating(0)
							copy.SetEncryptionRating(0)
							copy.SetSize(file.GetSize())
							copy.SetLastEditDate(STime)
							copy.markSet.MarksFrom[persona.GetID()] = 4
							printLog("...completed", congo.ColorGreen)
							printLog("New file spotted:", congo.ColorGreen)
							printLog("Icon: "+copy.GetName(), congo.ColorGreen)
							printLog("File Name: "+copy.GetFileName(), congo.ColorGreen)
							printLog("File Size: "+strconv.Itoa(copy.GetSize())+" Mp", congo.ColorGreen)
							printLog("File Owner: "+persona.GetName(), congo.ColorGreen)
							break
						} else if comm[i] == "DELETE" {
							printLog("...deleting file '"+file.GetFileName()+"'", congo.ColorGreen)
							host.DeleteFile(file)
							printLog("...complete", congo.ColorGreen)
							break
						} else if comm[i] == "ENCRYPT" {
							printLog("...encrypting file", congo.ColorGreen)
							file.SetEncryptionRating(netHits)
							printLog("...complete", congo.ColorGreen)
							break
						} else if comm[i] == "DOWNLOAD" {
							if !checkExistingMarks(persona.GetID(), file.GetID(), 4) {
								printLog("ACCESS DENIDED", congo.ColorYellow)
								printLog("This operation reserved only for owners", congo.ColorGreen)
							} else {
								printLog("...initiate download: "+file.GetFileName(), congo.ColorGreen)
								printLog("...file size: "+strconv.Itoa(file.GetSize()), congo.ColorGreen)
								persona.SetDownloadProcess(file.GetSize(), file.GetFileName())
							}
							break
						}
						//удалось
					}
				} else {
					printLog("...Failed", congo.ColorYellow)
				}
			}
		}
		endAction()
	}
}*/

//Edit -
func Edit(src IObj, trg IObj) {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiate Edit Mode...", congo.ColorGreen)
	targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetComputerSkill() + persona.GetLogic()
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
		printLog("...Error: Edit Mode glitch detected", congo.ColorYellow)
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Edit Mode critical failure", congo.ColorRed)
	}
	for i := range targetList {
		if file, ok := targetList[i].(IFile); ok {
			host := file.GetHost()
			dp2 := file.GetDeviceRating() + file.GetFirewall()
			suc2, rgl, rcgl := simpleTest(file.GetID(), dp2, 1000, 0)
			if rgl == true {
				printLog("..."+file.GetName()+": Encryption exploit detected", congo.ColorGreen)
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("..."+file.GetName()+": Encryption critical failure", congo.ColorGreen)
			}
			netHits := suc1 - suc2
			if !checkExistingMarks(persona.GetID(), file.GetID(), 1) {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...ACCESS DENIDED", congo.ColorRed)
			} else {
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
				netHits = suc1 - suc2
				if netHits > 0 {
					file.SetLastEditDate(STime)
					for i := range comm {
						if comm[i] == "COPY" {
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...copying", congo.ColorGreen)
							hold()
							copy := host.NewFile(file.GetFileName())
							copy.SetFileName("Copy of " + file.GetFileName())
							copy.SetDataBombRating(0)
							copy.SetEncryptionRating(0)
							copy.SetSize(file.GetSize())
							copy.SetLastEditDate(STime)
							copy.markSet.MarksFrom[persona.GetID()] = 4
							printLog("...completed", congo.ColorGreen)
							printLog("New file spotted:", congo.ColorGreen)
							printLog("Icon: "+copy.GetName(), congo.ColorGreen)
							printLog("File Name: "+copy.GetFileName(), congo.ColorGreen)
							printLog("File Size: "+strconv.Itoa(copy.GetSize())+" Mp", congo.ColorGreen)
							printLog("File Owner: "+persona.GetName(), congo.ColorGreen)
							break
						} else if comm[i] == "DELETE" {
							printLog("...deleting file '"+file.GetFileName()+"'", congo.ColorGreen)
							host.DeleteFile(file)
							printLog("...complete", congo.ColorGreen)
							break
						} else if comm[i] == "ENCRYPT" {
							printLog("...encrypting file", congo.ColorGreen)
							file.SetEncryptionRating(netHits)
							printLog("...complete", congo.ColorGreen)
							break
						} else if comm[i] == "DOWNLOAD" {
							if !checkExistingMarks(persona.GetID(), file.GetID(), 4) {
								printLog("ACCESS DENIDED", congo.ColorYellow)
								printLog("This operation reserved only for owners", congo.ColorGreen)
							} else {
								printLog("...initiate download: "+file.GetFileName(), congo.ColorGreen)
								printLog("...file size: "+strconv.Itoa(file.GetSize()), congo.ColorGreen)
								persona.SetDownloadProcess(file.GetSize(), file.GetFileName())
							}
							break
						}
						printLog("...File editing completed", congo.ColorGreen)
					}
				} else {
					printLog("...Failed", congo.ColorYellow)
				}
			}
		}
	}
	endAction()
}

//EnterHost - ++
func EnterHost(src IObj, trg IObj) {
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	persona := SourceIcon.(IPersona)

	if host, ok := TargetIcon.(*THost); ok {
		printLog("Entering Host...", congo.ColorGreen)
		printLog("..."+host.GetName(), congo.ColorGreen)
		if checkLinkLock(persona) == true {
			printLog("...Error: "+persona.GetName()+" is Locked", congo.ColorYellow)
		} else {
			if !checkExistingMarks(persona.GetID(), host.GetID(), 1) {
				printLog("...ACCESS DENIED", congo.ColorRed)
			} else { //выполняем само действие
				persona.SetHost(host)
				printLog("SYSTEM MESSAGE: You are now connected to "+host.GetName(), congo.ColorDefault)
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
	//host := persona.GetHost()
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	printLog("Switching silent running ... ", congo.ColorGreen)
	if checkLinkLock(persona) == true && src.(IObj).GetFaction() == player.GetFaction() {
		printLog("...Error: "+src.(IPersona).GetName()+" is Locked", congo.ColorYellow)
	} else {
		//persona.SetHost(host.GetHost())   -  логика для сложных хостов
		src.(IPersona).SetHost(Matrix)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...successful", congo.ColorGreen)
	}
	endAction()
}

//SwitchSilentMode - ++
func SwitchSilentMode(src IObj, trg IObj) {
	if persona, ok := src.(IPersona); ok {
		printLog("Switching silent running mode...", congo.ColorGreen)
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)
		newMode := comm[2]
		newMode = strings.Replace(newMode, "-", "_", -1)
		newMode = strings.Replace(newMode, " ", "_", -1)
		newMode = strings.Replace(newMode, "_", "", -1)
		switch newMode {
		case "ON":
			printLog("...Silent running mode is now: ON", congo.ColorGreen)
			persona.SetSilentRunningMode(true)
			isSimpleAction()
		case "OFF":
			printLog("...Silent running mode is now: OFF", congo.ColorYellow)
			persona.SetSilentRunningMode(false)
			isSimpleAction()
		default:
			printLog("...Error: Argument is invalid...", congo.ColorYellow)
		}
	}
	endAction()
}

//EraseMark -
func EraseMark(src IObj, trg IObj) {
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	if icon, ok := src.(IPersona); ok {
		//trg = pickObjByID(2)
		totalMarks := icon.CountMarks()
		if totalMarks > 0 {
			printLog("Erase Mark... ", congo.ColorGreen)
			hold()
			icon.ClearMarks()
			allMarks := icon.GetMarkSet()
			for r := range allMarks.MarksFrom {
				if allMarks.MarksFrom[r] > 0 && r != icon.GetID() {
					printLog("...MARK found", congo.ColorGreen)
					trg = pickObjByID(r)
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
		printLog("..."+strconv.Itoa(suc1)+" successes", congo.ColorGreen)
		// if mark from Icon
		if markOwner, ok := trg.(IIcon); ok {
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
				icon.GetMarkSet().MarksFrom[markOwner.GetID()] = icon.GetMarkSet().MarksFrom[markOwner.GetID()] - 1
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
	persona := SourceIcon.(IPersona)
	if grid, ok := trg.(*TGrid); ok {
		isComplexAction()
		goAct := true
		for goAct {
			goAct = false
			printLog("Switching grid to...", congo.ColorGreen)
			printLog("..."+grid.GetGridName(), congo.ColorGreen)
			if persona.GetHost() != Matrix {
				printLog("...Error: Impossible to switch grids while in Host", congo.ColorYellow)
				break
			}
			if checkLinkLock(persona) {
				printLog("...Error: Impossible to switch grids being Link-locked", congo.ColorYellow)
				break
			}
			persona.(IPersona).SetGrid(grid)
			printLog("...completed", congo.ColorGreen)
			printLog("Welcome to "+grid.GetGridName()+"!", congo.ColorGreen)
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

	var validSpecs []string
	validSpecs = append(validSpecs, "Hack on the Fly")
	haveSpec, spec := persona.HaveValidSpec(validSpecs)
	if haveSpec {
		printLog("ValidSpec is "+spec, congo.ColorDefault)
	}
	printLog("Initiating Hack on the Fly sequence...", congo.ColorGreen)
	targetList, targetSpec := pickTargets2(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetHackingSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	if haveSpec || targetSpec {
		dp1 = dp1 + 2
		printLog("...Specialization: +"+strconv.Itoa(2)+" op/p", congo.ColorGreen)
	}
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
	comm := GetComm()
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
	//printLog("Start", congo.ColorYellow)
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
	silentIcons := 0
	printLog("Initiating Matrix Perception sequence...", congo.ColorGreen)
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetSilentRunningMode() == true && icon.GetHost() == persona.GetHost() {
				silentIcons++
			}
			canSee := persona.GetFieldOfView().KnownData[icon.GetID()]
			if icon.GetHost() == persona.GetHost() && canSee[0] != "Spotted" {
				if icon.GetName() != persona.GetName() {
					targetList = append(targetList, icon)
				}
				//printLog("append "+icon.GetName()+" to TargetList", congo.ColorYellow) - debug
			}
		}
	}

	//targetList := pickTargets(comm)
	printLog("...Allocating resources:", congo.ColorGreen)
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: "+strconv.Itoa(dp1)+" op/p", congo.ColorGreen)
	attMod = calculateAttMods(comm, persona, targetList) //targetList[:1]
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
		}
		if persona.GetSilentRunningMode() {
			silentIcons--
		}
		printLog("..."+strconv.Itoa(silentIcons)+" icons running silent detected", congo.ColorYellow)
		break
	}
	//printLog("End", congo.ColorYellow)
	endAction()
}

//SwapAttributes - ++
func SwapAttributes(src IObj, trg IObj) {
	src = SourceIcon
	if persona, ok := src.(*TPersona); ok {
		printLog("Initiate attributes swapping...", congo.ColorGreen)
		//text := TargetIcon.(string)
		isFreeAction()
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.SplitN(text, ">", 4)
		swap1 := false
		swap2 := false
		att1 := 0
		switch comm[2] {
		case "ATTACK":
			att1 = SourceIcon.(*TPersona).GetAttackRaw()
			printLog("...Attribute 1 = Attack", congo.ColorGreen)
			swap1 = true
		case "SLEAZE":
			att1 = SourceIcon.(*TPersona).GetSleazeRaw()
			printLog("...Attribute 1 = Sleaze", congo.ColorGreen)
			swap1 = true
		case "DATA_PROCESSING":
			att1 = SourceIcon.(*TPersona).GetDataProcessingRaw()
			printLog("...Attribute 1 = Data Processing", congo.ColorGreen)
			swap1 = true
		case "FIREWALL":
			att1 = SourceIcon.(*TPersona).GetFirewallRaw()
			printLog("...Attribute 1 = Firewall", congo.ColorGreen)
			swap1 = true
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 1 is invalid...", congo.ColorYellow)
			swap1 = false

		}
		att2 := 0
		switch comm[3] {
		case "ATTACK":
			att2 = SourceIcon.(*TPersona).GetAttackRaw()
			printLog("...Attribute 2 = Attack", congo.ColorGreen)
			swap2 = true
		case "SLEAZE":
			att2 = SourceIcon.(*TPersona).GetSleazeRaw()
			printLog("...Attribute 2 = Sleaze", congo.ColorGreen)
			swap2 = true
		case "DATA_PROCESSING":
			att2 = SourceIcon.(*TPersona).GetDataProcessingRaw()
			printLog("...Attribute 2 = Data Processing", congo.ColorGreen)
			swap2 = true
		case "FIREWALL":
			att2 = SourceIcon.(*TPersona).GetFirewallRaw()
			printLog("...Attribute 2 = Firewall", congo.ColorGreen)
			swap2 = true
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 2 is invalid...", congo.ColorYellow)
			swap2 = false
		}
		if persona.device.canSwapAtt == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("This Persona cann't swap attributes!", congo.ColorRed)
			comm[2] = " "
			comm[3] = " "
		}
		if swap1 == true && swap2 == true {
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
		}
		if swap1 == true && swap2 == true {
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

//Wait - ++
func Wait(src IObj, trg IObj) {
	persona := src.(IPersona)
	comm := GetComm()
	for i := range comm {
		//printLog(comm[i], congo.ColorGreen)
		if comm[i] == "-EV" {
			loop := false
			search := persona.GetSearchProcess()
			download := persona.GetDownloadProcess()
			if len(search.SearchIconName) > 0 || len(download.DownloadIconName) > 0 {
				loop = true
			}
			if !loop {
				printLog("No active process. Waiting until the end of Combat Turn", congo.ColorGreen)
			} else {
				persona.SetWaitFlag(true)
			}
			//persona.SetWaitFlag(true)
			persona.SetInitiative(0)

		}
	}
	if len(comm) > 2 {
		waitTime := comm[2]
		waitTimeInt, _ := strconv.Atoi(waitTime)
		persona.SetInitiative(persona.GetInitiative() - waitTimeInt)
	} else {
		src.(IPersona).SetInitiative(0)
	}
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
	//comm := GetComm()
	//if SourceIcon.(IIcon).GetID() == player.GetID() && comm[1] != "WAIT" {
	//	hold()
	//	drawLineInWindow("Log")
	//}
	SourceIcon = nil
	TargetIcon = nil
	TargetIcon2 = nil
	command = ""
	//outIndex := 0
	for _, obj := range ObjByNames {
		if ic, ok := obj.(IIC); ok {
			if ic.GetMatrixCM() < 0 {
				host := ic.GetHost()
				host.DeleteIC(ic)
			}
		}
	}

	checkTurn()
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
	refreshProcessWin()

}

func isComplexAction() {
	if src, ok := SourceIcon.(IIcon); ok {
		src.SetInitiative(src.GetInitiative() - 10)
	}
}

func isFreeAction() {
	if src, ok := SourceIcon.(IIcon); ok {
		src.SetInitiative(src.GetInitiative() - 2) //Free Action
	}
}

func isSimpleAction() {
	if src, ok := SourceIcon.(IIcon); ok {
		src.SetInitiative(src.GetInitiative() - 5) //Simple Action
	}
}

/*func checkMarks(neededMarks int) bool {
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
}*/

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

func checkLinkLock(icon IIcon) bool {
	allLocks := icon.GetLinkLockStatus()
	var lockedBy []int
	for key, value := range allLocks.LockedByID { //check if non-slave marked by slave
		if isLocked(allLocks.LockedByID, key) { //if true
			if value == true {
				lockedBy = append(lockedBy, key)
				return true
			}
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
				if attacker.GetHost() == trgt.GetHost() && attacker.GetHost() != Matrix {
					attMod = attMod + 2
				} else {
					printLog("...Target "+strconv.Itoa(i+1)+" is in another Grid: "+strconv.Itoa(-2)+" op/p", congo.ColorGreen)
				}
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

func pickTargets2(comm []string) ([]IObj, bool) {

	var targetList []IObj
	persona := SourceIcon.(IPersona)
	////////////////////////

	totalSpec := true
	if len(comm) < 3 {
		return targetList, false
	}
	targetName := formatTargetName(comm[2])
	if grid, ok := ObjByNames[targetName].(*TGrid); ok {
		targetList = append(targetList, grid)
		printLog("...Target 1: "+grid.GetGridName()+" has top priority", congo.ColorYellow)
		var targetValidSpecs []string
		targetValidSpecs = append(targetValidSpecs, "Hvs."+grid.GetType())
		targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+grid.GetType())
		haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
		totalSpec = totalSpec && haveSpec
		return targetList, totalSpec
	}
	if icon1, ok := ObjByNames[targetName]; ok {
		newIcon := icon1.(IIcon) //не выводится за зону видимости((
		targetList = append(targetList, newIcon)
		printLog("...Target 1: "+newIcon.GetName(), congo.ColorGreen)
		var targetValidSpecs []string
		targetValidSpecs = append(targetValidSpecs, "Hvs."+icon1.GetType())
		targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+icon1.GetType())
		haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
		totalSpec = totalSpec && haveSpec
		//return targetList, totalSpec
		//	persona := SourceIcon.(IIcon)
		if persona.CheckRunningProgram("Fork") && len(comm) > 3 {
			targetName2 := formatTargetName(comm[3])
			if targetName != targetName2 {
				if icon2, ok := ObjByNames[targetName2]; ok {
					if grid, ok := ObjByNames[targetName].(*TGrid); ok {
						targetList = nil
						targetList = append(targetList, grid)
						printLog("...Target 2: "+grid.GetGridName()+" has top priority", congo.ColorYellow)
						printLog("...Target 1 replaced", congo.ColorGreen)
						var targetValidSpecs []string
						targetValidSpecs = append(targetValidSpecs, "Hvs."+grid.GetType())
						targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+grid.GetType())
						haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
						totalSpec = totalSpec && haveSpec
						printLog("--DEBUG--totalTargetSpec: "+strconv.FormatBool(totalSpec), congo.ColorYellow)
						return targetList, totalSpec
						//return targetList, false
					}
					newIcon2 := icon2.(IIcon)
					targetList = append(targetList, newIcon2)
					var targetValidSpecs []string
					targetValidSpecs = append(targetValidSpecs, "Hvs."+icon2.GetType())
					targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+icon2.GetType())
					haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
					totalSpec = totalSpec && haveSpec
					printLog("--DEBUG--totalTargetSpec: "+strconv.FormatBool(totalSpec), congo.ColorYellow)

					printLog("...Target 2: "+newIcon2.GetName(), congo.ColorGreen)

				}
			} else {
				printLog("...Error: Target 1 = Target 2", congo.ColorYellow)
			}
		}
	}
	printLog("--DEBUG--totalTargetSpec: "+strconv.FormatBool(totalSpec), congo.ColorYellow) //TODO: доразделить проверку vs.Host на типы действий
	return targetList, totalSpec
}

func placeMARK(source, target IIcon) {
	currentMARKS := target.GetMarkSet().MarksFrom[source.GetID()]
	currentMARKS++
	if target.GetName() != player.GetName() {
		printLog("...new MARK on "+target.GetName()+" was successfuly planted", congo.ColorGreen)
	} else {
		printLog("...Warning: MARK on "+target.GetName()+" was confirmed", congo.ColorYellow)
	}
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
		if canSee[i] == "Unknown" { //|| canSee[0] != "Spotted" {
			mem = append(mem, i)
		}
	}
	for i := rand.Intn(33); i > 0; i-- { //perception has a chance to reveal some data. TOPIC TO DISSCUSS: on what chance exactly is
		shuffleInt(mem)
	}
	for i := needToReveal; i > 0; i-- {
		if i < len(mem) {
			choosen := mem[i]
			if canSee[0] != "Spotted" {
				canSee[0] = "Spotted"
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
						printLog("..."+target.GetName()+": Encryption rating detected", congo.ColorGreen)
						printLog("..."+target.GetName()+": Encryption rating = "+strconv.Itoa(target.GetEncryptionRating()), congo.ColorYellow)
					} else {
						printLog("..."+target.GetName()+": No file encryption detected", congo.ColorGreen)
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
