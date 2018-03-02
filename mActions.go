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
	MActions.MActionMap["COMMAND"] = Command
	MActions.MActionMap["LONGACT"] = LongAct
	MActions.MActionMap["WAIT"] = Wait
	MActions.MActionMap["FULL_DEFENCE"] = FullDefence
	/////////////////////////////////////
	MActions.MActionMap["INFUSE_ATTACK"] = InfuseAttack
	MActions.MActionMap["INFUSE_SLEAZE"] = InfuseSleaze
	MActions.MActionMap["INFUSE_DATA_PROCESSING"] = InfuseDataProcessing
	MActions.MActionMap["INFUSE_FIREWALL"] = InfuseFirewall
	MActions.MActionMap["DIFFUSE_ATTACK"] = DiffuseAttack
	MActions.MActionMap["DIFFUSE_SLEAZE"] = DiffuseSleaze
	MActions.MActionMap["DIFFUSE_DATA_PROCESSING"] = DiffuseDataProcessing
	MActions.MActionMap["DIFFUSE_FIREWALL"] = DiffuseFirewall
	MActions.MActionMap["KILL_COMPLEX_FORM"] = KillComplexForm
	MActions.MActionMap["COMPILE"] = Compile
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
func LongAct(src IObj, trg IObj) bool {
	icon := src.(IPersona)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Start LongAct by " + icon.GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName() + "`s searchLen = " + strconv.Itoa(icon.GetLongAct()))

	return true
}

//PatrolICActionArea - patrol scans EVERY icon in the host acording to rules in Data Trails
func PatrolICActionArea(src IObj, trg IObj) bool {
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
				printLog("..." + icon.GetName() + " is being scanned")
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
					printLog("..." + icon.GetName() + " evaiding")
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
					printLog("..." + icon.GetName() + " detected")
				}
				if icon.GetFaction() != host.GetFaction() {
					host.alert = "Passive Alert"
				}
				host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
			} else {
				if icon.GetFaction() == player.GetFaction() {
					printLog("..." + icon.GetName() + " evaded")
				}
			}
		}
	}

	return true
}

//PatrolICActionTarget - Patrol scanning 1 target
func PatrolICActionTarget(src IObj, trg IObj) bool {
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
						printLog("..." + icon.GetName() + " was affected")
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
						printLog("..." + icon.GetName() + " was affected")
					}
				}
			}
		}
	}

	return true
}

//AcidICActionTarget -
func AcidICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Acid IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetFirewall() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetFirewallMod(focusIcon.GetFirewallMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall Rating reduced to " + strconv.Itoa(focusIcon.GetFirewall()))
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//BinderICActionTarget -
func BinderICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Binder IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetDataProcessing() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetDataProcessingMod(focusIcon.GetDataProcessingMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Data Processing Rating reduced to " + strconv.Itoa(focusIcon.GetDataProcessing()))
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}

	return true
}

//JammerICActionTarget -
func JammerICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Jammer IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetAttack() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetAttackMod(focusIcon.GetAttackMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Attack Rating reduced to " + strconv.Itoa(focusIcon.GetAttack()))
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}

	return true
}

//MarkerICActionTarget -
func MarkerICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Marker IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetSleaze() < 1 {
						fullDamage := netHits
						realDamage := focusIcon.ResistMatrixDamage(fullDamage)
						focusIcon.ReceiveMatrixDamage(realDamage)

					} else {
						focusIcon.SetSleazeMod(focusIcon.GetSleazeMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Sleaze Rating reduced to " + strconv.Itoa(focusIcon.GetSleaze()))
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}

		}

	}

	return true
}

//KillerICActionTarget -
func KillerICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Killer IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - " + src.(*TIC).GetName() + " have " + strconv.Itoa(m) + " marks on " + focusIcon.GetName())
					fullDamage := host.GetAttack() + netHits + 2*m
					realDamage := focusIcon.ResistMatrixDamage(fullDamage)
					focusIcon.ReceiveMatrixDamage(realDamage)

				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//SparkyICActionTarget -
func SparkyICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Sparky IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - " + src.(*TIC).GetName() + " have " + strconv.Itoa(m) + " marks on " + focusIcon.GetName())
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
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//TarBabyICActionTarget -
func TarBabyICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Tar Baby IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
			if focusIcon.GetHost().name == src.(*TIC).GetHost().name {
				dp2 := focusIcon.GetLogic() + focusIcon.GetFirewall()
				suc2, dgl, dcgl := simpleTest(focusIcon.GetID(), dp2, 1000, 0)
				if dgl {
					congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected!")
					addOverwatchScoreToTarget(suc1)
				}
				if dcgl {
					suc1++
					congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure!!!")
					addOverwatchScoreToTarget(10)
				}
				netHits := suc1 - suc2

				if focusIcon.GetFaction() == player.GetFaction() {
					hold()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetLinkLockStatus().LockedByID[src.(*TIC).GetID()] == true {
						focusIcon.GetMarkSet().MarksFrom[src.(*TIC).GetID()] = focusIcon.GetMarkSet().MarksFrom[src.(*TIC).GetID()] + 1
						if focusIcon.GetMarkSet().MarksFrom[src.(*TIC).GetID()] > 3 {
							focusIcon.GetMarkSet().MarksFrom[src.(*TIC).GetID()] = 3
						}
					} else {
						src.(*TIC).LockIcon(focusIcon)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...connection locked")
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//BlackICActionTarget -
func BlackICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(IIC)
	trg = TargetIcon.(IPersona)
	host := src.(IIC).GetHost()

	dp1 := host.GetDeviceRating() * 2
	limit := host.GetAttack()
	suc1, gl, cgl := simpleTest(src.GetID(), dp1, limit, 0)
	if gl == true {
		suc1--
	}
	if cgl == true {
		host.alert = "No Alert"
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Black IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - " + src.(*TIC).GetName() + " have " + strconv.Itoa(m) + " marks on " + focusIcon.GetName())
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...connection locked")
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//BlasterICActionTarget -
func BlasterICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Black IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					marks := focusIcon.GetMarkSet()
					m := marks.MarksFrom[src.(*TIC).GetID()]
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - " + src.(*TIC).GetName() + " have " + strconv.Itoa(m) + " marks on " + focusIcon.GetName())
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...connection locked")
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//ProbeICActionTarget -
func ProbeICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Probe IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					focusIcon.GetMarkSet().MarksFrom[host.GetID()] = focusIcon.GetMarkSet().MarksFrom[host.GetID()] + 1
					if focusIcon.GetMarkSet().MarksFrom[host.GetID()] > 3 {
						focusIcon.GetMarkSet().MarksFrom[host.GetID()] = 3
					}
					congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG -" + focusIcon.GetName() + " marked by " + host.GetName())
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//ScrambleICActionTarget -
func ScrambleICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Scramble IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetMarkSet().MarksFrom[host.GetID()] > 3 {
						focusIcon.GetMarkSet().MarksFrom[host.GetID()] = 3
					}
					if focusIcon.GetMarkSet().MarksFrom[host.GetID()] > 2 {
						//focusIcon.Dumpshock()
						focusIcon.SetSimSence("OFFLINE")
						focusIcon.SetInitiative(999999)
						//	congo.WindowsMap.ByTitle["Log"].WPrintLn("Connection terminated")
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//CatapultICActionTarget -
func CatapultICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Catapult IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetFirewall() < 1 {
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn("DEBUG - Firewall reduced to 0 ")
						}
					} else {
						focusIcon.SetFirewallMod(focusIcon.GetFirewallMod() - 1)
						if focusIcon.GetFaction() == player.GetFaction() {
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...firewall rating reduced to " + strconv.Itoa(focusIcon.GetFirewall()))
						}
					}
					//////////BIOFEEDBACK DAMAGE//////////////
					biofeedbackDamage := 0
					if focusIcon.GetSimSence() == "HOT-SIM" || focusIcon.GetSimSence() == "COLD-SIM" {
						biofeedbackDamage = netHits + focusIcon.GetMarkSet().MarksFrom[host.GetID()]
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
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//ShokerICActionTarget -
func ShokerICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Acid IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					focusIcon.SetInitiative(focusIcon.GetInitiative() - 5)
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...initiative reduced by 5")
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//TrackICActionTarget -
func TrackICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Probe IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetMarkSet().MarksFrom[host.GetID()] > 1 {
						focusIcon.SetPhysicalLocation(true)
						if focusIcon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...physical location tracked")
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...local authorities reported by " + host.GetName() + " administration")
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//BloodhoundICActionTarget -
func BloodhoundICActionTarget(src IObj, trg IObj) bool {
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Bloodhound IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetMarkSet().MarksFrom[host.GetID()] > 1 {
						focusIcon.SetPhysicalLocation(true)
						if focusIcon.GetFaction() == player.GetFaction() {
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...physical location tracked")
							hold()
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...local authorities reported by " + host.GetName() + " administration")
						}
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//BloodhoundICActionArea -
func BloodhoundICActionArea(src IObj, trg IObj) bool {
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Bloodhound succeses: " + strconv.Itoa(suc1))
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
							congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + icon.GetName() + " evading: " + strconv.Itoa(suc2) + " successes")
							if dgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName() + ": Encryption glitch detected")
							}
							if dcgl {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName() + ": Encryption critical failure")
							}
						}
						if netHits > 0 {
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + icon.GetName() + " affected")
							}
							bloodhoundIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + icon.GetName() + " detected")
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
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + icon.GetName() + " affected")
							}
							bloodhoundIC.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							host.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
							if icon.GetFaction() == player.GetFaction() {
								hold()
								congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + icon.GetName() + " detected")
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

	return true
}

//CrashICActionTarget -
func CrashICActionTarget(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon.(IPersona)
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn(src.(*TIC).GetName() + " attacking " + trg.(IObj).GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Crash IC succeses: " + strconv.Itoa(suc1))
	/////////////////////////////////////////////////////////////////////////
	if host.alert == "Passive Alert" || host.alert == "Active Alert" {
		isComplexAction()
		if focusIcon, ok := trg.(IPersona); ok {
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...evading: " + strconv.Itoa(suc2) + " successes")
					if dgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall glitch detected")
					}
					if dcgl {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn(focusIcon.GetName() + ": Firewall critical failure")
					}
				}
				if netHits > 0 {
					if focusIcon.GetFaction() == player.GetFaction() {
						hold()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...affected")
					}
					if focusIcon.GetMarkSet().MarksFrom[host.GetID()] > 0 {
						focusIcon.CrashRandomProgram()
					}
				} else {
					if focusIcon.GetFaction() == player.GetFaction() {
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...evaded")
					}
					src.(*TIC).ReceiveMatrixDamage(-netHits)
				}
			}
		}
	}

	return true
}

//ICWait - action for all ICes: do nothing
func ICWait(src IObj, trg IObj) bool {
	src = SourceIcon.(*TIC)
	trg = TargetIcon
	if ic, ok := SourceIcon.(*TIC); ok {
		ic.SetInitiative(ic.GetInitiative() - 0)
		ic.SpendComplexAction()
		//printLog(ic.GetName()+" waiting")
	}
	//src.(IIC).SpendComplexAction()

	return true
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
	case "COMMAND":
		actionIsGood = true
		mActionName = "COMMAND"
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
	case "INFUSE_ATTACK":
		actionIsGood = true
		mActionName = "INFUSE_ATTACK"
		return actionIsGood, mActionName
	case "INFUSE_SLEAZE":
		actionIsGood = true
		mActionName = "INFUSE_SLEAZE"
		return actionIsGood, mActionName
	case "INFUSE_DATA_PROCESSING":
		actionIsGood = true
		mActionName = "INFUSE_DATA_PROCESSING"
		return actionIsGood, mActionName
	case "INFUSE_FIREWALL":
		actionIsGood = true
		mActionName = "INFUSE_FIREWALL"
		return actionIsGood, mActionName
	case "DIFFUSE_ATTACK":
		actionIsGood = true
		mActionName = "DIFFUSE_ATTACK"
		return actionIsGood, mActionName
	case "DIFFUSE_SLEAZE":
		actionIsGood = true
		mActionName = "DIFFUSE_SLEAZE"
		return actionIsGood, mActionName
	case "DIFFUSE_DATA_PROCESSING":
		actionIsGood = true
		mActionName = "DIFFUSE_DATA_PROCESSING"
		return actionIsGood, mActionName
	case "DIFFUSE_FIREWALL":
		actionIsGood = true
		mActionName = "DIFFUSE_FIREWALL"
		return actionIsGood, mActionName
	case "KILL_COMPLEX_FORM":
		actionIsGood = true
		mActionName = "KILL_COMPLEX_FORM"
		return actionIsGood, mActionName
	case "COMPILE":
		actionIsGood = true
		mActionName = "COMPILE"
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
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: Unknown command")
		//panic("fi")
		return false, mActionName

	}
	return actionIsGood, mActionName
}

//doAction - handler for matrix action for both player AND enemies
func doAction(mActionName string) bool {
	if SourceIcon == nil {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("{RED}--DEBUG-- {YELLOW}Error: SourceIcon = nil. Try again")
		return false
	}
	//	player.RevealDataTo(player.GetID(), 7)
	if val, ok := MActions.MActionMap[mActionName]; ok { // if s string match any Matrix Action than execute action
		if mActionName != "ICWAIT" {
			if attacker, ok := SourceIcon.(IIcon); ok {
				if defender, ok := TargetIcon.(IPersona); ok {
					if defender.GetID() == player.GetID() && attacker.GetID() != player.GetID() {
						printLog(attacker.GetName() + " attack detected...")
						printLog("..." + defender.GetName() + ": Evaiding")
						player.SetWaitFlag(false)
					}
				}
			}
		}
		if canIconCanDoAction(mActionName, SourceIcon.(IIcon)) {
			actionSuccsesful := val.(func(IObj, IObj) (actionSuccsesful bool))(SourceIcon, TargetIcon)
			if actionSuccsesful {
				endAction()
			}
			//draw()
			return true
		}
		//printLog(SourceIcon.GetName()+" cant do: "+mActionName)
		//checkTurn()

	}
	//draw()
	return false
}

func canIconCanDoAction(mActionName string, icon IIcon) bool {
	canDo := false
	//Оцениваем действие
	actionType := ""
	switch mActionName {
	case "LOAD_PROGRAM":
		actionType = "free"
	case "UNLOAD_PROGRAM":
		actionType = "free"
	case "SWAP_PROGRAMS":
		actionType = "free"
	case "SWAP_ATTRIBUTES":
		actionType = "free"
		//
	case "WAIT":
		actionType = "simple"
	case "SWITCH_INTERFACE_MODE":
		actionType = "simple"
	case "SILENT_MODE":
		actionType = "simple"
	case "SEND_MESSAGE":
		actionType = "simple"
	case "CHECK_OVERWATCH_SCORE":
		actionType = "simple"
		//////////////////////////////////
	default:
		actionType = "complex"
	}
	//Оцениваем свободные действия у иконы
	avFree, avSimple := getActions(icon)
	//Сравниваем вес действия и свободные действия
	if actionType == "free" {
		if avFree > 0 {
			canDo = true
		} else if avSimple > 0 {
			canDo = true
		}
	}
	if actionType == "simple" {
		if avSimple > 0 {
			canDo = true
		}
	}
	if actionType == "complex" {
		if avSimple > 1 {
			canDo = true
		}
	}
	if !canDo {
		if icon == player {
			printLog("Error: Impossible to comply on this Action Phase")
		}
	}
	return canDo
}

func getActions(icon IIcon) (int, int) {
	availFree := icon.GetFreeActionsCount()
	availSimple := icon.GetSimpleActionsCount()
	return availFree, availSimple
}

//BruteForce - ++
func BruteForce(src IObj, trg IObj) bool {
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
		printLog("ValidSpec is " + spec)
	}
	printLog("Initiating Brute Force sequence...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetCyberCombatSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	/*	if haveSpec || targetSpec {
		dp1 = dp1 + 2
		printLog("...Specialization: +"+strconv.Itoa(2)+" op/p")
	}*/
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetAttack()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure")
	}
	printLog("..." + persona.GetName() + ": " + strconv.Itoa(suc1) + " successes")
	for i := range targetList {
		if grid, ok := targetList[i].(*TGrid); ok {
			dp2 := grid.GetDeviceRating() * 2
			suc2, _, _ := simpleTest(grid.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			netHits := suc1 - suc2
			if netHits > 0 {
				printLog("...Grid encryption bypassed")
				persona.SetGrid(grid)
			}
		} else if icon, ok := targetList[i].(IIcon); ok {
			host := icon.GetHost()
			dp2 := icon.GetDeviceRating() + icon.GetFirewall()
			suc2, rgl, rcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("...Unexpected exploit detected!")
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("...Target's firewall critical falure")
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
			printLog("...Error: Target " + strconv.Itoa(i+1) + " is not a valid type")
		}
	}

	return true
}

//SwitchInterfaceMode - ++
func SwitchInterfaceMode(src IObj, trg IObj) bool {
	src = SourceIcon
	if persona, ok := src.(IPersona); ok {
		printLog("Initiate Interface Mode switching...")
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
		//printLog(newMode)
		switch newMode {
		case "AR":
			newMode = "AR"
			printLog("...Switching to Augmented Reality")
		case "COLD":
			newMode = "COLD-SIM"
			printLog("...Switching to Virtial Reality")
			printLog("...Safety mode : ON")
			if persona.GetDevice().GetModel() == "Living Persona" {
				printLog("...Error: @$#%**$$@^!@")
				printLog("...Interface Mode switching failed")
				return false
			}
		case "HOT":
			newMode = "HOT-SIM"
			printLog("...Switching to Virtial Reality")
			printLog("...Safety mode : OFF")
		default:
			printLog("...Error: User Mode is invalid...")
			newMode = persona.GetSimSence()

		}
		isSimpleAction()
		persona.SetSimSence(newMode)
		//printLog("--DEBUG--: Changes will be applied on next turn. Canonic Initiative System planed on later date")

	}

	return true
}

//CheckOverwatchScore - ++
func CheckOverwatchScore(src IObj, trg IObj) bool {
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
	printLog("Checking Overwatch Score...")
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
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + grid.GetGridName() + ": current Overwatch Score = " + strconv.Itoa(persona.GetGrid().GetLastSureOS()))
		hold()
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...failed")
		hold()
	}
	addOverwatchScore(suc2)

	return true
}

//CrackFile - ++
func CrackFile(src IObj, trg IObj) bool {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Data Spike sequence...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetHackingSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetAttack()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Attack protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Attack protocol critical failure")
	}
	for i := range targetList {
		if file, ok := targetList[i].(IFile); ok {
			if !checkExistingMarks(persona.GetID(), file.GetID(), 1) { // проверяем марки
				printLog("..." + file.GetName() + ": ACCESS DENIED")
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
				printLog("...Encryption exploit found")
			}
			if cglt == true {
				persona.GetGrid().SetOverwatchScore(0)
				printLog("...Overwatch Score cleared")
			}
			netHits := suc1 - suc2
			addOverwatchScore(suc2)
			if netHits > 0 {
				file.SetEncryptionRating(0)
				printLog("..." + file.GetName() + " decrypted")
				persona.ChangeFOWParametr(file.GetID(), 12, strconv.Itoa(file.GetEncryptionRating())) // 12- отвечает за Encryption
			} else {
				printLog("...Failure! File encryption is not disabled")
			}
		}
	}

	return true
}

//DataSpike - ++
func DataSpike(src IObj, trg IObj) bool {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Data Spike sequence...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetCyberCombatSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetAttack()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Attack protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Attack protocol critical failure")
	}
	for i := range targetList {
		if icon, ok := targetList[i].(IIcon); ok {
			host := icon.GetHost()
			dp2 := icon.GetDeviceRating() + icon.GetFirewall()
			suc2, rgl, rcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("..." + icon.GetName() + ": Firewall exploit detected")
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("..." + icon.GetName() + ": Firewall critical failure")
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
						printLog("...Error: " + icon.GetName() + " is immune to Biofeedback Damage")
					}
				}
				if persona.CheckRunningProgram("Blackout") {
					if target, ok := icon.(IPersona); ok {
						bfDamage := target.ResistBiofeedbackDamage(realDamage)
						target.ReceiveStunBiofeedbackDamage(bfDamage)
					} else {
						printLog("...Error: " + icon.GetName() + " is immune to Biofeedback Damage")
					}
				}
				if host.GetHostAlertStatus() == "No Alert" {
					host.alert = "Passive Alert"
				}
			} else {
				persona.ReceiveMatrixDamage(-netHits)
			}

		} else {
			printLog("...Error: Target " + strconv.Itoa(i+1) + " is not a valid type")
		}
	}

	return true
}

//DisarmDataBomb -
func DisarmDataBomb(src IObj, trg IObj) bool {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Disarm Databomb sequence...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetSoftwareSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetFirewall()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Disarming protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Disarming protocol critical failure")
	}
	for i := range targetList {
		if file, ok := targetList[i].(IFile); ok {
			dp2 := file.GetDataBombRating() * 2
			suc2, rgl, rcgl := simpleTest(file.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("..." + file.GetName() + ": Firewall exploit detected")
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("..." + file.GetName() + ": Firewall critical failure")
			}
			netHits := suc1 - suc2
			if netHits > 0 {
				file.SetDataBombRating(0)
				printLog("...Databomb Disarmed")
			} else {
				persona.TriggerDataBomb(file.GetDataBombRating())
				file.SetDataBombRating(0)
			}
			persona.ChangeFOWParametr(file.GetID(), 3, strconv.Itoa(file.GetDataBombRating())) // 3- отвечает за DataBomb
		} else {
			printLog("...Error: Target " + strconv.Itoa(i+1) + " is not a valid type")
		}
	}

	return true
}

//SetDatabomb -
func SetDatabomb(src IObj, trg IObj) bool {
	src = SourceIcon.(IPersona)
	trg = TargetIcon
	icon := SourceIcon.(IPersona)
	host := trg.(IFile).GetHost()
	dp1 := icon.GetSoftwareSkill() + icon.GetLogic() // + attMod
	limit := src.(IPersona).GetSleaze()
	suc1, gl, cgl := simpleTest(icon.GetID(), dp1, limit, 0)
	if icon.GetFaction() == player.GetFaction() {
		printLog("Setting up databomb on " + trg.GetName() + "...")
	}
	isComplexAction()
	if trg, ok := trg.(*TFile); ok {
		printLog("...installing databomb")
		dp2 := host.GetDeviceRating() * 2
		suc2, glt, cglt := simpleTest(trg.GetID(), dp2, 1000, 0)
		if gl == true {
			addOverwatchScore(dp1 - suc1)
			printLog("...Error: Unexpected trigger initiated")
		}
		if cgl == true {
			addOverwatchScore(dp1 - suc1)
			icon.TriggerDataBomb(suc1)
			printLog("...critical error erupted")
		}
		if glt == true {
			addOverwatchScore(-suc2)
			suc2--
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...firewall exploit detected")
		}
		if cglt == true {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...exploit critical")
			suc1++
		}
		netHits := suc1 - suc2
		addOverwatchScore(suc2)
		if netHits > 0 {
			if icon.CheckRunningProgram("Demolition") {
				netHits++
				printLog("...Databomb rating infused by Demolition program")
			}
			trg.SetDataBombRating(netHits)
			printLog("...Databomb rating " + strconv.Itoa(netHits) + " installed")
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Databomb installation failed")
		}
	}

	return true
}

/*//Edit0 - old function. Delete if Edit will receive no bugs
func Edit0(src IObj, trg IObj) bool  {
	comm := GetComm()
	trg = TargetIcon
	host := SourceIcon.(IIcon).GetHost()
	if persona, ok := SourceIcon.(IPersona); ok {
		if persona.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Enable Edit mode")
		}
		dp1 := persona.GetComputerSkill() + persona.GetLogic() // + attMod
		limit := persona.GetDataProcessing()
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...Hardware limit: "+strconv.Itoa(limit))
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if gl == true {
			addOverwatchScore(8)
		}
		if cgl == true {
			addOverwatchScore(40)
		}
		if persona.GetFaction() == player.GetFaction() {
			if gl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...unexpected error ocured")
			}
			if cgl {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...critical error erupted")
			}
		}
		if file, ok := trg.(IFile); ok {
			if !checkExistingMarks(persona.GetID(), file.GetID(), 1) {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...ACCESS DENIDED")
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: File Encrypted")
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
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...copying")
							hold()
							copy := host.NewFile(file.GetFileName())
							copy.SetFileName("Copy of " + file.GetFileName())
							copy.SetDataBombRating(0)
							copy.SetEncryptionRating(0)
							copy.SetSize(file.GetSize())
							copy.SetLastEditDate(STime)
							copy.markSet.MarksFrom[persona.GetID()] = 4
							printLog("...completed")
							printLog("New file spotted:")
							printLog("Icon: "+copy.GetName())
							printLog("File Name: "+copy.GetFileName())
							printLog("File Size: "+strconv.Itoa(copy.GetSize())+" Mp")
							printLog("File Owner: "+persona.GetName())
							break
						} else if comm[i] == "DELETE" {
							printLog("...deleting file '"+file.GetFileName()+"'")
							host.DeleteFile(file)
							printLog("...complete")
							break
						} else if comm[i] == "ENCRYPT" {
							printLog("...encrypting file")
							file.SetEncryptionRating(netHits)
							printLog("...complete")
							break
						} else if comm[i] == "DOWNLOAD" {
							if !checkExistingMarks(persona.GetID(), file.GetID(), 4) {
								printLog("ACCESS DENIDED")
								printLog("This operation reserved only for owners")
							} else {
								printLog("...initiate download: "+file.GetFileName())
								printLog("...file size: "+strconv.Itoa(file.GetSize()))
								persona.SetDownloadProcess(file.GetSize(), file.GetFileName())
							}
							break
						}
						//удалось
					}
				} else {
					printLog("...Failed")
				}
			}
		}

	}
}*/

//Edit -
func Edit(src IObj, trg IObj) bool {
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiate Edit Mode...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetComputerSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Edit Mode glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Edit Mode critical failure")
	}
	for i := range targetList {
		if file, ok := targetList[i].(IFile); ok {
			host := file.GetHost()
			dp2 := file.GetDeviceRating() + file.GetFirewall()
			suc2, rgl, rcgl := simpleTest(file.GetID(), dp2, 1000, 0)
			if rgl == true {
				printLog("..." + file.GetName() + ": Encryption exploit detected")
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("..." + file.GetName() + ": Encryption critical failure")
			}
			netHits := suc1 - suc2
			if !checkExistingMarks(persona.GetID(), file.GetID(), 1) {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...ACCESS DENIDED")
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: File Encrypted")
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
							congo.WindowsMap.ByTitle["Log"].WPrintLn("...copying")
							hold()
							copy := host.NewFile(file.GetFileName())
							copy.SetFileName("Copy of " + file.GetFileName())
							copy.SetDataBombRating(0)
							copy.SetEncryptionRating(0)
							copy.SetSize(file.GetSize())
							copy.SetLastEditDate(STime)
							copy.markSet.MarksFrom[persona.GetID()] = 4
							printLog("...completed")
							printLog("New file spotted:")
							printLog("Icon: " + copy.GetName())
							printLog("File Name: " + copy.GetFileName())
							printLog("File Size: " + strconv.Itoa(copy.GetSize()) + " Mp")
							printLog("File Owner: " + persona.GetName())
							break
						} else if comm[i] == "DELETE" {
							printLog("...deleting file '" + file.GetFileName() + "'")
							host.DeleteFile(file)
							printLog("...complete")
							break
						} else if comm[i] == "ENCRYPT" {
							printLog("...encrypting file")
							file.SetEncryptionRating(netHits)
							printLog("...complete")
							break
						} else if comm[i] == "DOWNLOAD" {
							if !checkExistingMarks(persona.GetID(), file.GetID(), 4) {
								printLog("ACCESS DENIDED")
								printLog("This operation reserved only for owners")
							} else {
								printLog("...initiate download: " + file.GetFileName())
								printLog("...file size: " + strconv.Itoa(file.GetSize()))
								persona.SetDownloadProcess(file.GetSize(), file.GetFileName())
							}
							break
						}
					}
					printLog("...File editing completed")
				} else {
					printLog("...Failed")
				}
			}
		}
	}

	return true
}

//EnterHost - ++
func EnterHost(src IObj, trg IObj) bool {
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	persona := SourceIcon.(IPersona)

	if host, ok := TargetIcon.(*THost); ok {
		printLog("Entering Host...")
		printLog("..." + host.GetName())
		if checkLinkLock(persona) == true {
			printLog("...Error: " + persona.GetName() + " is Locked")
		} else {
			if !checkExistingMarks(persona.GetID(), host.GetID(), 1) {
				printLog("...ACCESS DENIED")
			} else { //выполняем само действие
				persona.SetHost(host)
				if persona == player {
					printLog("SYSTEM MESSAGE: You are now connected to " + host.GetName())
				}
				//printLog("SYSTEM MESSAGE: You are now connected to "+host.GetName())
			}
		}
		for _, val := range ObjByNames {
			if icon, ok := val.(IIcon); ok && icon.GetHost() == host {
				if icon.GetSilentRunningMode() {
					continue
				}
				persona.ChangeFOWParametr(icon.GetID(), 0, "Spotted")
			}
		}
	} else {
		printLog("Entering Host...")
		printLog("...Error: Target is not a Host")
	}

	return true
}

//ExitHost - ++
func ExitHost(src IObj, trg IObj) bool {
	persona := SourceIcon.(IPersona)
	//host := persona.GetHost()
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	printLog("Leaving host... ")
	if checkLinkLock(persona) == true && src.(IObj).GetFaction() == player.GetFaction() {
		printLog("...Error: " + src.(IPersona).GetName() + " is Locked")
	} else {
		//persona.SetHost(host.GetHost())   -  логика для сложных хостов
		src.(IPersona).SetHost(Matrix)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("...successful")
	}

	return true
}

//SwitchSilentMode - ++
func SwitchSilentMode(src IObj, trg IObj) bool {
	if persona, ok := src.(IPersona); ok {
		printLog("Switching silent running mode...")
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
			printLog("...Silent running mode is now: ON")
			persona.SetSilentRunningMode(true)
			isSimpleAction()
		case "OFF":
			printLog("...Silent running mode is now: OFF")
			persona.SetSilentRunningMode(false)
			isSimpleAction()
		default:
			printLog("...Error: Argument is invalid...")
		}
	}

	return true
}

//EraseMark -
func EraseMark(src IObj, trg IObj) bool {
	isComplexAction() // есть вероятность что стрельнет механизм возврата
	if icon, ok := src.(IPersona); ok {
		//trg = pickObjByID(2)
		totalMarks := icon.CountMarks()
		if totalMarks > 0 {
			printLog("Erase Mark... ")
			hold()
			icon.ClearMarks()
			allMarks := icon.GetMarkSet()
			for r := range allMarks.MarksFrom {
				if allMarks.MarksFrom[r] > 0 && r != icon.GetID() {
					printLog("...MARK found")
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
		printLog("..." + strconv.Itoa(suc1) + " successes")
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
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...MARK erased")
				hold()
				icon.GetMarkSet().MarksFrom[markOwner.GetID()] = icon.GetMarkSet().MarksFrom[markOwner.GetID()] - 1
				icon.ClearMarks()
			} else {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...failed ")
				hold()
			}
		}
	}

	return true
}

//GridHop -
func GridHop(src IObj, trg IObj) bool {
	persona := SourceIcon.(IPersona)
	if grid, ok := trg.(*TGrid); ok {
		isComplexAction()
		goAct := true
		for goAct {
			goAct = false
			printLog("Switching grid to...")
			printLog("..." + grid.GetGridName())
			if persona.GetHost() != Matrix {
				printLog("...Error: Impossible to switch grids while in Host")
				break
			}
			if checkLinkLock(persona) {
				printLog("...Error: Impossible to switch grids being Link-locked")
				break
			}
			persona.(IPersona).SetGrid(grid)
			printLog("...completed")
		}
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Switching grid to...")
		congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + trg.(IObj).GetName())
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error!! " + trg.(IObj).GetName() + " is not a grid.")
	}

	return true
}

//HackOnTheFly - ++
func HackOnTheFly(src IObj, trg IObj) bool {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	persona := src.(IPersona)
	isComplexAction()
	/*text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.Split(text, ">")*/
	comm := GetComm()
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
	printLog("Initiating Hack on the Fly sequence...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetHackingSkill() + persona.GetLogic()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetSleaze()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure")
	}
	printLog("..." + persona.GetName() + ": " + strconv.Itoa(suc1) + " successes")

	for i := range targetList {
		if grid, ok := targetList[i].(*TGrid); ok {
			dp2 := grid.GetDeviceRating() * 2
			suc2, _, _ := simpleTest(grid.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			netHits = suc1 - suc2
			if netHits > 0 {
				printLog("...Grid encryption bypassed")
				persona.SetGrid(grid)
			}
		} else if icon, ok := targetList[i].(IIcon); ok {
			host := icon.GetHost()
			dp2 := icon.GetDeviceRating() + icon.GetFirewall()
			suc2, rgl, rcgl := simpleTest(icon.GetID(), dp2, 1000, 0)
			addOverwatchScore(suc2)
			if rgl == true {
				printLog("...Unexpected exploit detected!")
				suc1++
			}
			if rcgl == true {
				addOverwatchScore(-dp2)
				printLog("...Target's firewall critical falure")
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
				host.SetAlert("Passive Alert")
				placeMARK(icon, persona)
			}
		} else {
			printLog("...Error: " + icon.GetName() + " is not a valid type")
		}
	}

	return true
}

//MatrixPerception - ++
func MatrixPerception(src IObj, trg IObj) bool {
	src = SourceIcon
	trg = TargetIcon
	var netHits int
	persona := src.(IPersona)
	isComplexAction()
	comm := GetComm()
	attMod := 0
	printLog("Initiating Matrix Perception sequence...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure")
	}
	printLog("..." + persona.GetName() + ": " + strconv.Itoa(suc1) + " successes")

	for j := range targetList {
		needToReveal := suc1
		if icon, ok := targetList[j].(IIcon); ok {
			if icon.GetHost() != persona.GetHost() && icon.GetType() != "Host" {
				printLog("...Warning: " + icon.GetName() + " unreachable")
				continue
			}
			if icon.GetSilentRunningMode() {
				dp2 := icon.GetSleaze() + icon.GetDeviceRating()
				suc2, dgl, cdgl := simpleTest(icon.GetID(), dp2, 1000, 0)
				if dgl == true {
					addOverwatchScore(-suc1)
					printLog("...Encryption weakness detected")
				}
				if cdgl == true {
					addOverwatchScore(-suc1 * 2)
					suc1 = suc1 + 100
					printLog("...Encryption critical falure")

				}
				netHits = suc1 - suc2
				needToReveal = netHits
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
				if netHits < 0 {
					printLog("...Matrix Perception failed")
				}
			} else {
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
			}
		}
	}

	return true
}

//MatrixSearch -
func MatrixSearch(src IObj, trg IObj) bool {
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

	printLog("Initiating Matrix Search sequense...")
	targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList)
	if persona.CheckRunningProgram("Search") {
		attMod = attMod + 2
		printLog("...'Search' program running: " + strconv.Itoa(2) + " op/p")
	}
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure")
	}
	if suc1 > 0 {
		//	multiPrintTo("Log", mPrnt{"..." + persona.GetName() + ": ", GREEN}, mPrnt{suc1}, mPrnt{" successes", GREEN})
		printLog("..." + persona.GetName() + ": " + strconv.Itoa(suc1) + " successes")
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

		printLog("..." + persona.GetName() + ": Search in progress")
		for i := range comm {
			comm[i] = formatTargetName(comm[i])
			//	printLog("comm["+strconv.Itoa(i)+"] = "+comm[i])
		}
		//printLog("...Search ETA: "+strconv.Itoa(persona.GetSearchResultIn()*3)+" seconds")
	} else {
		printLog("..." + persona.GetName() + ": Search Failed")
	}

	return true
}

//ScanEnviroment - ++
func ScanEnviroment(src IObj, trg IObj) bool {
	//printLog("Start")
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
	printLog("Initiating Matrix Perception sequence...")
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetSilentRunningMode() == true && icon.GetHost() == persona.GetHost() {
				silentIcons++
			}
			canSee := persona.GetFieldOfView().KnownData[icon.GetID()]
			if icon.GetHost() == persona.GetHost() && canSee[0] != "Spotted" {
				if icon.GetFaction() != persona.GetFaction() {
					targetList = append(targetList, icon)
				}
				//printLog("append "+icon.GetName()+" to TargetList") // - debug
			}
		}
	}

	//targetList := pickTargets(comm)
	printLog("...Allocating resources:")
	dp1 := persona.GetComputerSkill() + persona.GetIntuition()
	printLog("...Base MPCP resources: " + strconv.Itoa(dp1) + " op/p")
	attMod = calculateAttMods(comm, persona, targetList) //targetList[:1]
	dp1 = dp1 + attMod
	if dp1 < 0 {
		dp1 = 0
	}
	printLog("...Evaluated Software resources: " + strconv.Itoa(dp1) + " op/p")
	limit := persona.GetDataProcessing()
	printLog("...Hardware limit: " + strconv.Itoa(limit) + " op/p")
	suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
	if gl == true {
		addOverwatchScore(dp1 - suc1)
		printLog("...Error: Encryption protocol glitch detected")
	}
	if cgl == true {
		addOverwatchScore(dp1)
		persona.GetHost().SetAlert("Active Alert")
		printLog("...Error: Encryption protocol critical failure")
	}
	printLog("..." + persona.GetName() + ": " + strconv.Itoa(suc1) + " successes")

	for j := range targetList {
		//printLog("Debug: Evaluating "+targetList[j].GetName())
		needToReveal := suc1
		if icon, ok := targetList[j].(IIcon); ok {
			if icon.GetSilentRunningMode() {
				dp2 := icon.GetSleaze() + icon.GetDeviceRating()
				suc2, dgl, cdgl := simpleTest(icon.GetID(), dp2, 1000, 0)
				if dgl == true {
					addOverwatchScore(-suc1)
					printLog("...Encryption weakness detected")
				}
				if cdgl == true {
					addOverwatchScore(-suc1 * 2)
					suc1 = suc1 + 100
					printLog("...Encryption critical falure")

				}
				netHits = suc1 - suc2
				//printLog("Debug: "+icon.GetName()+": netHits "+strconv.Itoa(netHits))
				needToReveal = netHits

				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)

				if netHits < 0 {
					printLog("...Matrix Perception failed")
				}
			} else {
				//printLog("Debug: ELSE:: "+icon.GetName()+": netHits "+strconv.Itoa(netHits))
				persona.GetFieldOfView().KnownData[icon.GetID()] = revealData(persona, icon, needToReveal)
			}
		}
		if persona.GetSilentRunningMode() {
			silentIcons--
		}
		printLog("..." + strconv.Itoa(silentIcons) + " icons running silent detected")
		break
	}
	//printLog("End")

	return true
}

//SwapAttributes -
func SwapAttributes(src IObj, trg IObj) bool { //need to rewrite to use IPersona
	src = SourceIcon.(IPersona)
	if persona, ok := src.(IPersona); ok {
		printLog("Initiate attributes swapping...")
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
			//att1 = SourceIcon.(IPersona).GetAttackRaw()
			att1 = persona.GetAttackRaw()
			printLog("...Attribute 1 = Attack")
			swap1 = true
		case "SLEAZE":
			att1 = persona.GetSleazeRaw()
			printLog("...Attribute 1 = Sleaze")
			swap1 = true
		case "DATA_PROCESSING":
			att1 = persona.GetDataProcessingRaw()
			printLog("...Attribute 1 = Data Processing")
			swap1 = true
		case "FIREWALL":
			att1 = persona.GetFirewallRaw()
			printLog("...Attribute 1 = Firewall")
			swap1 = true
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 1 is invalid...")
			swap1 = false

		}
		att2 := 0
		switch comm[3] {
		case "ATTACK":
			att2 = persona.GetAttackRaw()
			printLog("...Attribute 2 = Attack")
			swap2 = true
		case "SLEAZE":
			att2 = persona.GetSleazeRaw()
			printLog("...Attribute 2 = Sleaze")
			swap2 = true
		case "DATA_PROCESSING":
			att2 = persona.GetDataProcessingRaw()
			printLog("...Attribute 2 = Data Processing")
			swap2 = true
		case "FIREWALL":
			att2 = persona.GetFirewallRaw()
			printLog("...Attribute 2 = Firewall")
			swap2 = true
		default:
			congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 2 is invalid...")
			swap2 = false
		}
		if persona.GetDevice().canSwapAtt == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("This Persona can not swap attributes!")
			return false
		}
		if swap1 == true && swap2 == true {
			if comm[2] == "ATTACK" {
				persona.SetAttackRaw(att2)
				//persona.SetDeviceAttackRaw(att2)
				swap1 = true
			} else if comm[2] == "SLEAZE" {
				persona.SetSleazeRaw(att2)
				//persona.SetDeviceSleazeRaw(att2)
				swap1 = true
			} else if comm[2] == "DATA_PROCESSING" {
				persona.SetDataProcessingRaw(att2)
				//persona.SetDeviceDataProcessingRaw(att2)
				swap1 = true
			} else if comm[2] == "FIREWALL" {
				persona.SetFirewallRaw(att2)
				//persona.SetDeviceFirewallRaw(att2)
				swap1 = true
			} else {
				swap1 = false
			}
			if comm[3] == "ATTACK" {
				//persona.SetDeviceAttackRaw(att1)
				persona.SetAttackRaw(att1)
				swap2 = true
			} else if comm[3] == "SLEAZE" {
				persona.SetSleazeRaw(att1)
				//persona.SetDeviceSleazeRaw(att1)
				swap2 = true
			} else if comm[3] == "DATA_PROCESSING" {
				persona.SetDataProcessingRaw(att1)
				//persona.SetDeviceDataProcessingRaw(att1)
				swap2 = true
			} else if comm[3] == "FIREWALL" {
				persona.SetFirewallRaw(att1)
				//persona.SetDeviceFirewallRaw(att1)
				swap2 = true
			} else {
				swap2 = false
			}
			if comm[2] == comm[3] {
				swap1 = false
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...Error: Attribute 1 = Attribute 2")
			}
		}
		if swap1 == true && swap2 == true {
			printLog("Attribute swapping complete")
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Attribute swapping failed")
		}
	}

	return true
}

//LoadProgram -
func LoadProgram(src IObj, trg IObj) bool {
	src = SourceIcon
	comm := GetComm()
	targetProgram := formatTargetName(comm[2])
	printLog("Loading Program...")
	if persona, ok := src.(IPersona); ok {
		programs := persona.GetPrograms()
		printLog("...Program: " + targetProgram)
		prgFound := false
		for i := range programs {
			if programs[i].programName != targetProgram {
				continue
			}
			prgFound = true
			if programs[i].programStatus != "Stored" {
				printLog("...Error: Program '" + programs[i].programName + "' is " + programs[i].programStatus)
				return false
			}
			printLog("...Status: " + programs[i].programStatus)
			persona.GetDevice().LoadProgramToDevice(targetProgram)
		}
		if !prgFound {
			printLog("...Error: Program '" + targetProgram + "' is not available")
			return false
		}
		persona.SpendFreeAction()
	}
	printLog("Loading successful")
	return true
}

//Login -
func Login(src IObj, trg IObj) bool {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	text := command
	text = formatString(text)
	text = cleanText(text)
	comm := strings.SplitN(text, ">", 3)
	target := formatTargetName(comm[2])
	if persona, ok := src.(IPersona); ok {
		printLog(">>>LOGIN: " + target)
		printLog(">>>PASSCODE: XXXXXXXXXXXXXXXXX")
		if persona.GetName() == "Unknown" {
			if target == "Unknown" {
				printLog("...Error: 'Guest' already signed in")
			} else {
				var valid bool
				player, valid = ImportPlayerFromDB(target)
				if valid {
					printLog("...Passcode accepted")
					printLog("...Biometric data generated")
					printLog("...Start session:")
					delete(ObjByNames, "Unknown")
				}
			}
		} else {
			printLog("SYSTEM ERROR: Persona already logged in.")
			printLog("Terminate session if you want to use another account")
		}
	}
	SourceIcon = player

	//SourceIcon = player
	return true
}

//UnloadProgram -
func UnloadProgram(src IObj, trg IObj) bool {
	src = SourceIcon
	comm := GetComm()
	targetProgram := formatTargetName(comm[2])
	printLog("Unloading Program...")
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(IPersona); ok {
		programs := persona.GetPrograms()
		printLog("...Program: " + targetProgram)
		prgFound := false
		for i := range programs {
			if programs[i].programName != targetProgram {
				continue
			}
			prgFound = true
			if programs[i].programStatus != "Running" {
				printLog("...Error: Program '" + programs[i].programName + "' is not running")
				return false
			}
			persona.GetDevice().UnloadProgramFromDevice(targetProgram) //unload here
		}
		if !prgFound {
			printLog("...Error: Program '" + targetProgram + "' is not available")
			return false
		}
		persona.SpendFreeAction()

	}
	printLog("...Exit code: 0")
	printLog("Program termination completed")
	return true
}

//SwapPrograms - ++
func SwapPrograms0(src IObj, trg IObj) bool {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(IPersona); ok {
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
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Stopping program...")
				time.Sleep(dur)
				draw()
				if persona.GetDeviceSoft().programStatus[i] == "Running" {
					if true { //////Проверка устройства
						persona.GetDevice().UnloadProgramFromDevice(program)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("..." + program + " Terminated")
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Program exit code:0")
						isFreeAction()

					}
				} else {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: " + program + " is " + persona.GetDeviceSoft().programStatus[i])
				}
				programFound = true
			}
		}
		if programFound == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: program '" + comm[2] + "' not found")
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
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Loading program...")
				time.Sleep(dur)
				draw()
				if persona.GetDeviceSoft().programStatus[i] == "Running" {
					programINRunning = true
				}
				if persona.GetDeviceSoft().programStatus[i] == "inStore" {

					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program name: " + persona.GetDeviceSoft().programName[i])
					time.Sleep(dur)
					draw()
					congo.WindowsMap.ByTitle["Log"].WPrintLn("...program type: " + persona.GetDeviceSoft().programType[i])
					time.Sleep(dur)
					draw()
					if persona.GetDevice().LoadProgramToDevice(program) { //////Проверка устройства
						//persona.LoadProgram(program)
						congo.WindowsMap.ByTitle["Log"].WPrintLn("...program status: " + persona.GetDeviceSoft().programStatus[i])
						time.Sleep(dur)
						draw()
						congo.WindowsMap.ByTitle["Log"].WPrintLn("Loading Complete")
						programINfound = true
					}
				}
			}
		}
		if programINfound {
			isFreeAction()
		} else if programINRunning {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: program '" + comm[3] + "' already running")
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program '" + comm[3] + "' cannot be loaded")
		}
	}

	return true
}

//SwapPrograms - ++
func SwapPrograms(src IObj, trg IObj) bool {
	printLog("Program Swap protocol initiated... ")
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	comm := GetComm()
	if len(comm) < 4 {
		printLog("...Error: not enough data")
		printLog("Program Swap failed")
		return false
	}
	programOUT := formatTargetName(comm[2])
	printLog("...Unloading '" + programOUT + "' program")
	programIN := formatTargetName(comm[3])
	loadedOUT := false
	loadedIN := false
	if persona, ok := src.(IPersona); ok {
		programs := persona.GetPrograms()
		//clearing space
		for i := range programs {
			if programs[i].programName != programOUT {
				continue
			}
			if programs[i].programStatus == "Running" {
				persona.GetDevice().UnloadProgramFromDevice(programOUT)
				loadedOUT = true
			}
		}
		if !loadedOUT {
			printLog("...Error: Program '" + programOUT + "' is not running")
			printLog("Program Swap failed")
			return false
		}
		printLog("...completed")
		//loading new program
		printLog("...Loading '" + programIN + "' program")
		for i := range programs {
			if programs[i].programName != programIN {
				continue
			}
			if programs[i].programStatus == "Stored" {
				persona.GetDevice().LoadProgramToDevice(programIN)
				loadedIN = true
			} else {
				printLog("...Error: Program '" + programs[i].programName + "' is " + programs[i].programStatus)
			}
		}
		if !loadedIN {
			printLog("Program Swap failed")
			return false
		}
		printLog("...completed")
		printLog("Program Swap succsessful")
	}
	return true
}

func Command(src IObj, trg IObj) bool {
	comm := GetComm()

	targetList := pickTargets(comm)
	if len(comm) < 4 {
		printLog("...Error: Message has no data")
		printLog("Use '[SEND MESSAGE]>[TARGET]>[MESSAGE TEXT]' format")
		return false
	}
	if persona, ok := src.(IPersona); ok {
		persona.SpendSimpleAction()
		for i := range targetList {
			if agent, ok := targetList[i].(IAgent); ok {
				if comm[3] == "REPORT" {
					printLog("Status report:")
					printLog("Persona: " + agent.GetName())
					printLog("Owner :" + agent.GetOwner().GetName())
					printLog("Action Protocol: " + agent.GetActionProtocol())
					printLog("Current Enviroment: " + agent.GetHost().GetName() + " " + agent.GetGrid().GetGridName())
					printLog("Silent Running mode: " + strconv.FormatBool(agent.GetSilentRunningMode()))
					printLog("End report:")
				}
				if comm[3] == "WAIT" {
					agent.SetActionProtocol("Idle")
				}
				if comm[3] == "FOLLOW" {
					agent.SetActionProtocol("Follow")
				}
				if comm[3] == "SCAN_ENVIROMENT" {
					agent.SetActionProtocol("Overwatch")
				}
			}
		}
	}

	return true
}

//Wait - ++
func Wait(src IObj, trg IObj) bool {
	persona := src.(IPersona)
	comm := GetComm()
	//date := (SrTime.String())
	//date = (SrTime.Format("2006-01-02 15:04:05"))
	timePeriod := "init"
	timeInt := 0
	for i := range comm {
		//printLog(comm[i])
		if comm[i] == "-EV" {
			loop := false
			search := persona.GetSearchProcess()
			download := persona.GetDownloadProcess()
			if len(search.SearchIconName) > 0 || len(download.DownloadIconName) > 0 {
				loop = true
			}
			if !loop {
				printLog("No active process. Waiting until the end of Combat Turn")
			} else {
				persona.SetWaitFlag(true)
			}
			break
			//persona.SetWaitFlag(true)
			//persona.SetInitiative(0)

		}
		if comm[i] == "TURNS" {
			timePeriod = "turns"
			break
		}
		if comm[i] == "MINUTES" {
			timePeriod = "minutes"
			break
		}
		if comm[i] == "HOURS" {
			timePeriod = "hours"
			break
		}
	}
	if len(comm) > 2 {
		waitTime := comm[2]
		waitTimeInt, _ := strconv.Atoi(waitTime)
		timeInt = waitTimeInt
		//persona.SetInitiative(persona.GetInitiative() - timeInt)
	} else {
		//src.(IPersona).SetInitiative(0)
	}
	expectDate := SrTime
	switch timePeriod {
	default:
	case "turns":
		for t := 0; t < timeInt; t++ {
			expectDate = expectDate.Add(3 * time.Second)
			persona.SetWaitFlag(true)
			eXdate := (expectDate.Format("2006-01-02 15:04:05"))
			TimeMarker = eXdate
		}
		persona.SpendComplexAction()
		return true
	case "minutes":
		for t := 0; t < timeInt*20; t++ {
			expectDate = expectDate.Add(3 * time.Second)
			persona.SetWaitFlag(true)
			eXdate := (expectDate.Format("2006-01-02 15:04:05"))
			TimeMarker = eXdate
		}
		persona.SpendComplexAction()
		return true
	case "hours":
		for t := 0; t < timeInt*20*60; t++ {
			expectDate = expectDate.Add(3 * time.Second)
			persona.SetWaitFlag(true)
			eXdate := (expectDate.Format("2006-01-02 15:04:05"))
			TimeMarker = eXdate
		}
		persona.SpendComplexAction()
		return true
	}
	persona.SpendComplexAction()
	/*testPrint("Status name: " + Status(player.GetID()).GetStatusName())
	for _, obj := range ObjByNames {
		printLog("Obj = "+obj.GetName()+" Status: "+Status(obj.GetID()).GetStatusName())
	}*/
	return true
}

//FullDefence -
func FullDefence(src IObj, trg IObj) bool {
	src = SourceIcon
	congo.WindowsMap.ByTitle["User Input"].WClear()
	if persona, ok := src.(IPersona); ok {
		printLog("Full defence protocol initiated")
		persona.SetFullDeffenceFlag(true)
		isComplexAction()
		//Status(persona.GetID()).SetStatusName("Supperssed")
	}

	return true
}

//InfuseAttack -
func InfuseAttack(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Infusion of Attack: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetAttack() {
			if suc1 > target.GetAttack() {
				suc1 = target.GetAttack()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Infusion of Attack", level, suc1)
			printLog("...Threadeng successful")

		} else {
			printLog("...Threadeng failed")
			printLog("Target's Attack higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//InfuseSleaze -
func InfuseSleaze(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Infusion of Sleaze: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetSleaze() {
			if suc1 > target.GetSleaze() {
				suc1 = target.GetSleaze()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Infusion of Sleaze", level, suc1)
			printLog("...Threading successful")

		} else {
			printLog("...Threading failed")
			printLog("Target's Sleaze higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//InfuseDataProcessing -
func InfuseDataProcessing(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Infusion of Data Processing: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetDataProcessing() {
			if suc1 > target.GetDataProcessing() {
				suc1 = target.GetDataProcessing()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Infusion of Data Processing", level, suc1)
			printLog("...Threading successful")

		} else {
			printLog("...Threading failed")
			printLog("Target's Data Processing higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//InfuseFirewall -
func InfuseFirewall(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Infusion of Firewall: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetFirewall() {
			if suc1 > target.GetFirewall() {
				suc1 = target.GetFirewall()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Infusion of Firewall", level, suc1)
			printLog("...Threading successful")

		} else {
			printLog("...Threading failed")
			printLog("Target's Firewall higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//DiffuseAttack -
func DiffuseAttack(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Diffusion of Attack: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetAttack() {
			if suc1 > target.GetAttack() {
				suc1 = target.GetAttack()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Diffusion of Attack", level, suc1)
			printLog("...Threadeng successful")

		} else {
			printLog("...Threadeng failed")
			printLog("Target's Attack higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//DiffuseSleaze -
func DiffuseSleaze(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Diffusion of Sleaze: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetSleaze() {
			if suc1 > target.GetSleaze() {
				suc1 = target.GetSleaze()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Diffusion of Sleaze", level, suc1)
			printLog("...Threading successful")

		} else {
			printLog("...Threading failed")
			printLog("Target's Sleaze higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//DiffuseDataProcessing -
func DiffuseDataProcessing(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Diffusion of Data Processing: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetDataProcessing() {
			if suc1 > target.GetDataProcessing() {
				suc1 = target.GetDataProcessing()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Diffusion of Data Processing", level, suc1)
			printLog("...Threading successful")

		} else {
			printLog("...Threading failed")
			printLog("Target's Data Processing higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//DiffuseFirewall -
func DiffuseFirewall(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("[COMPLEX_FORM]>[TARGET]>[LEVEL]")
		return false
	}
	printLog(persona.GetName())
	level := retrieveLevel()
	if level < 1 {
		printLog("Error: Level not designated correctly")
		printLog("Use '-L3' for level = 3, '-L10' for level = 10, ect...")
		return false
	}
	if level > (persona.GetResonance() * 2) {
		printLog("Error: Level can not be higher than Resonance x 2")
		return false
	}
	targetList := pickTargets(comm)
	printLog("Begin threadeng:")
	persona.SpendComplexAction()
	printLog("...Diffusion of Firewall: " + strconv.Itoa(level) + " level")
	if target, ok := TargetIcon.(IIcon); ok {
		attMod := calculateAttMods(comm, persona, targetList)
		dp1 := persona.GetSoftwareSkill() + persona.GetResonance() + attMod
		limit := level
		fade := level - 2
		fadeType := "stun"
		if fade < 2 {
			fade = 2
		}
		suc1, gl, cgl := simpleTest(persona.GetID(), dp1, limit, 0)
		if suc1 > persona.GetResonance() {
			fadeType = "phys"
		}
		if gl == true {
			fade = fade + (xd6Test(1) / 2)
		}
		if cgl == true {
			fade = fade + (xd6Test(1) / 2)
			fadeType = "phys"
		}
		if level >= target.GetFirewall() {
			if suc1 > target.GetFirewall() {
				suc1 = target.GetFirewall()
			}
			TreadComplexForm(persona.GetID(), target.GetID(), "Diffusion of Firewall", level, suc1)
			printLog("...Threading successful")

		} else {
			printLog("...Threading failed")
			printLog("Target's Firewall higher than Complex Form level")
		}

		persona.ResistFade(fade, fadeType)
	} else {
		printLog("Error: This Complex Form is not usable for this target type")
		return false
	}

	return true
}

//KillComplexForm -
func KillComplexForm(src IObj, trg IObj) bool {
	if livPersona, ok := src.(ITechnom); ok {
		printLog(livPersona.GetDevice().GetModel())
		if livPersona.GetDevice().GetModel() != "Living Persona" {
			printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
			return false
		}
	} else {
		printLog("not a Technomancer")
		return false
	}
	////////////////////////
	persona := src.(ITechnom)
	comm := GetComm()
	targetList := pickTargets(comm)
	formName := "Random Form"
	printLog("Comm:")
	for i := range comm {
		printLog(comm[i])
	}
	printLog("Targets:")
	for i := range targetList {
		printLog(targetList[i].GetName())
	}
	printLog("Complex Form:")
	if len(comm) < 4 {
		printLog("Random Form")
	} else {
		formName = formatTargetName(comm[3])
		printLog(formName + " form")
	}

	if target, ok := TargetIcon.(IIcon); ok {
		printLog("execute target: " + target.GetName())
		if formName == "Random Form" {
			for i := range CFDBMap {
				if getComplexForm(i).madeOnID == target.GetID() {
					formName = getComplexForm(i).cfName
					break
				}
			}
			printLog("Random Form picked: " + formName)
		}
		formTreaded, formID := target.CheckThreadedForm(formName)
		if formTreaded {
			threaderID := getComplexForm(formID).madeByID
			threaderName := pickObjByID(threaderID).GetName()
			printLog(threaderName)
			if threaderID == persona.GetID() {
				printLog("This is own Form, no test Needed")
				delete(CFDBMap, formID)
				persona.SpendFreeAction()
			} else {
				printLog("--DEBUG--: This is NOT own Form. TODO: Write a test")
			}
		} else {
			printLog("Form: " + formName + " is not threded on " + target.GetName())
		}

	}
	return true
}

//Compile -
func Compile(src IObj, trg IObj) bool {
	livPersona, ok := src.(ITechnom)
	if !ok {
		printLog("--DEBUG--:Can't use Resonance abilities with mundane electronics (CRB p.251)")
		return false
	}
	comm := GetComm()
	for i := range comm { //DEBUG
		printLog("comm[" + strconv.Itoa(i) + "] = " + comm[i])
	}
	printLog("0")
	if len(comm) < 4 {
		printLog("Error: Level not designated")
		printLog("COMPILE>[SPRITE_TYPE]>[LEVEL]")
		return false
	}
	printLog("1")
	level := retrieveLevel()
	if level > livPersona.GetResonance()*2 {
		printLog("Error: Can't compile Sprite with level higher than Resonance*2")
		return false
	}
	if level < 1 {
		printLog("Error: Can't compile Sprite with level lower than 1")
		return false
	}
	printLog("2")
	var desc string
	if len(comm) > 2 {
		desc = comm[2]
	}
	printLog("3")
	var sprite *TSprite
	switch desc {
	case "COURIER_SPRITE":
		sprite = livPersona.NewSprite("Courier Sprite", level)
	case "CRACK_SPRITE":
		sprite = livPersona.NewSprite("Crack Sprite", level)
	case "DATA_SPRITE":
		sprite = livPersona.NewSprite("Data Sprite", level)
	case "FAULT_SPRITE":
		sprite = livPersona.NewSprite("Fault Sprite", level)
	case "MACHINE_SPRITE":
		sprite = livPersona.NewSprite("Machine Sprite", level)
	default:
		printLog("Error: Unknown type of sprite")
		return false
	}
	targetList := pickTargets(comm)
	fadeType := "stun"
	tasks := 0
	if target, ok := TargetIcon.(IIcon); ok {
		target.GetName()
		attMod := calculateAttMods(comm, livPersona, targetList)
		spriteHits, _, _ := simpleTest(sprite.GetID(), level, level, 0)
		printLog("Sprite has " + strconv.Itoa(spriteHits) + " hits")
		fade := spriteHits * 2
		if fade < 2 {
			fade = 2
		}
		dp1 := livPersona.GetCompilingSkill() + livPersona.GetResonance() + attMod
		suc1, gl, cgl := simpleTest(livPersona.GetID(), dp1, level, 0)
		if gl {
			fade = fade + (xd6Test(1)+1)/2
		}
		if cgl {
			fade = fade + (xd6Test(1))
		}
		if suc1 > livPersona.GetResonance() {
			fadeType = "phys"
		}
		livPersona.ResistFade(fade, fadeType)
		tasks = suc1 - spriteHits
		if tasks <= 0 {
			tasks = 0
			delete(ObjByNames, sprite.GetName())
		}
	}
	printLog(sprite.GetName() + " compiled " + desc + " type " + fadeType + " tasks = " + strconv.Itoa(tasks))

	return true
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
	if SourceIcon.(IIcon).GetSimpleActionsCount() < 1 { // if sourceIcon used all Simple Actions => calculate next sourceIcon
		endActionPhase(SourceIcon.(IIcon))
		SourceIcon = nil
		TargetIcon = nil
		TargetIcon2 = nil
		command = ""
		for _, obj := range ObjByNames {
			if ic, ok := obj.(IIC); ok {
				if ic.GetMatrixCM() < 0 {
					host := ic.GetHost()
					host.DeleteIC(ic)
				}
			}
		}
	}
	refreshEnviromentWin()
	refreshPersonaWin()
	refreshGridWin()
	refreshProcessWin()
}

func isComplexAction() { //Evaluate IconType and spend 2 Simple actions if possible
	if src, ok := SourceIcon.(IIcon); ok {
		src.SpendComplexAction()
	}
}

func isFreeAction() { //Evaluate IconType and spend 1 Free action if possible
	if src, ok := SourceIcon.(IIcon); ok {
		src.SpendFreeAction()
	}
}

func isSimpleAction() { //Evaluate IconType and spend 1 Simple action if possible
	if src, ok := SourceIcon.(IIcon); ok {
		src.SpendSimpleAction()
	}
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
	congo.WindowsMap.ByTitle["Log"].WPrintLn("--DEBUG--Error: Неизвестный тип для checkMarks()!")
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
		mustReturn = true
	}
	if mustReturn == true {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn(icon.GetName()+" is Locked")
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

func calculateAttMods(comm []string, attacker IPersona, targetList []IObj) (attMod int) {
	//Action Specializations (Computer && Software)
	haveSpec, _ := attacker.HaveValidSpec(getActionSpecs())
	if haveSpec {
		attMod = 2
	}
	//Target Specializations (Hacking && CyberCombat)
	mustApplytargetSpec := true
	for i := range targetList {
		var targetValidSpecs []string
		if actionIs("Brute Force") || actionIs("Data Spike") || actionIs("Crash Program") {
			targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+targetList[i].GetType())
		}
		if actionIs("Crack File") || actionIs("Hack On The Fly") || actionIs("Spoof Command") {
			targetValidSpecs = append(targetValidSpecs, "HackingSpec_vs."+targetList[i].GetType())
		}
		haveTargetSpec, _ := attacker.HaveValidSpec(targetValidSpecs)
		if haveTargetSpec {
			mustApplytargetSpec = mustApplytargetSpec && haveTargetSpec
		} else {
			mustApplytargetSpec = false
		}
	}
	if mustApplytargetSpec {
		attMod = 2
	}
	if attMod == 2 {
		printLog("...Active Specialization: +" + strconv.Itoa(2) + " op/p")
	}
	var woundMod int
	sWoundMod := (attacker.GetMaxStunCM() - attacker.GetStunCM()) / 3
	pWoundMod := (attacker.GetMaxPhysCM() - attacker.GetPhysCM()) / 3
	woundMod = sWoundMod + pWoundMod
	if woundMod > 0 {
		printLog("...Wound modificator: -" + strconv.Itoa(woundMod) + " op/p")
	}
	attMod = attMod + woundMod

	var oppCyc bool
	for i := range comm {
		if actionIs("Brute Force") || actionIs("Hack On The Fly") {
			if comm[i] == "-2M" && oppCyc == false {
				printLog("...Additional operation cycles: " + strconv.Itoa(-4) + " op/p")
				attMod = attMod - 4
				oppCyc = true
			}
			if comm[i] == "-3M" && oppCyc == false {
				printLog("...Additional operation cycles: " + strconv.Itoa(-10) + " op/p")
				attMod = attMod - 10
				oppCyc = true
			}
		}
	}
	cfMod := countSustainedForms(attacker.GetID())
	if cfMod > 0 {
		attMod = attMod - (cfMod * 2)
		printLog("...Sustained Complex Forms: -" + strconv.Itoa(cfMod*2) + " op/p")
	}
	if attacker.GetGrid().name == "Public Grid" {
		attMod = attMod - 2
		printLog("...Public Grid lags: " + strconv.Itoa(-2) + " op/p")
	}
	for i := range targetList {
		if trgt, ok := targetList[i].(IIcon); ok {
			if attacker.GetGrid() != trgt.GetGrid() {
				attMod = attMod - 2
				if attacker.GetHost() == trgt.GetHost() && attacker.GetHost() != Matrix {
					attMod = attMod + 2
				} else {
					printLog("...Target " + strconv.Itoa(i+1) + " is in another Grid: " + strconv.Itoa(-2) + " op/p")
				}
			}
		}
	}
	if attacker.GetSimSence() == "HOT-SIM" {
		attMod = attMod + 2
		printLog("...HOT-SIM connection boost: +" + strconv.Itoa(2) + " op/p")
	}
	return attMod
}

//helper funcs:

func formatTargetName(targetName string) string {
	targetName = strings.ToLower(targetName)
	targetName = strings.Replace(targetName, "_", " ", -1)
	targetName = strings.Title(targetName)
	targetName = strings.Replace(targetName, " Ic", " IC", -1)
	targetName = strings.Replace(targetName, " Of ", " of ", -1)
	return targetName
}

func pickTargets(comm []string) []IObj {
	var targetList []IObj
	if len(comm) < 3 {
		return targetList
	}
	targetName := formatTargetName(comm[2])
	if targetName == "SELF" {
		targetName = SourceIcon.GetName()
	}
	if grid, ok := ObjByNames[targetName].(*TGrid); ok {
		targetList = append(targetList, grid)
		printLog("...Target 1: " + grid.GetGridName() + " has top priority")
		return targetList
	}
	if icon1, ok := ObjByNames[targetName]; ok {
		newIcon := icon1.(IIcon) //не выводится за зону видимости((
		targetList = append(targetList, newIcon)
		printLog("...Target 1: " + newIcon.GetName())
		persona := SourceIcon.(IIcon)
		if persona.CheckRunningProgram("Fork") && len(comm) > 3 {
			targetName2 := formatTargetName(comm[3])
			if targetName != targetName2 {
				if icon2, ok := ObjByNames[targetName2]; ok {
					if grid, ok := ObjByNames[targetName].(*TGrid); ok {
						targetList = nil
						targetList = append(targetList, grid)
						printLog("...Target 2: " + grid.GetGridName() + " has top priority")
						printLog("...Target 1 replaced")
						return targetList
					}
					newIcon2 := icon2.(IIcon)
					targetList = append(targetList, newIcon2)
					printLog("...Target 2: " + newIcon2.GetName())
				}
			} else {
				printLog("...Error: Target 1 = Target 2")
			}
		}
	}
	return targetList
}

func getActionSpecs() []string {
	var validSpecs []string
	action := getCurrentActionName()
	switch action {
	//ComputeSkillSpecs:
	case "Edit File":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Erase Mark":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Erase Matrix Signature":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Format Device":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Matrix Perception":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Matrix Search":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Reboot Device":
		validSpecs = append(validSpecs, "CompSpec_"+action)
	case "Trace Icon":
		validSpecs = append(validSpecs, "CompSpec_"+action)
		//SoftwareSkill Specs:
	case "Disarm Databomb":
		validSpecs = append(validSpecs, "SoftwareSpec_"+action)
	case "Set Databomb":
		validSpecs = append(validSpecs, "SoftwareSpec_"+action)
	default:
	}
	return validSpecs
}

func getCurrentActionName() string {
	comm := GetComm()
	if len(comm) < 2 {
		return "--UNKNOWN--"
	}
	actionName := formatTargetName(comm[1])
	return actionName
}

func actionIs(actionName string) bool {
	comm := GetComm()
	if len(comm) < 2 {
		return false
	}
	actionName = formatTargetName(actionName)
	formatedComm := formatTargetName(comm[1])
	if formatedComm != actionName {
		return false
	}
	return true
}

func pickTargets2(comm []string) ([]IObj, bool) {
	var targetList []IObj
	persona := SourceIcon.(IPersona)
	////////////////////////
	totalSpec := true
	if len(comm) < 3 {
		return targetList, false
	}
	printLog(comm[1])
	targetName := formatTargetName(comm[2])
	printLog("...Target Type: " + ObjByNames[targetName].GetType())
	printLog("...Target Name: " + ObjByNames[targetName].GetName())
	if grid, ok := ObjByNames[targetName].(*TGrid); ok {
		targetList = append(targetList, grid)
		printLog("...Target 1: " + grid.GetGridName() + " has top priority")
		var targetValidSpecs []string
		if actionIs("Brute Force") {
			targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+grid.GetType())
		}
		haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
		totalSpec = totalSpec && haveSpec
		return targetList, totalSpec
	}
	if icon1, ok := ObjByNames[targetName]; ok {
		newIcon := icon1.(IIcon) //не выводится за зону видимости((
		targetList = append(targetList, newIcon)
		printLog("...Target 1: " + newIcon.GetName())

		var targetValidSpecs []string
		if actionIs("Brute Force") || actionIs("Data Spike") || actionIs("Crash Program") {
			targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+icon1.GetType())
		}
		if actionIs("Crack File") || actionIs("Hack On The Fly") || actionIs("Spoof Command") {
			targetValidSpecs = append(targetValidSpecs, "HackSpec_vs."+icon1.GetType())
		}
		haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
		totalSpec = totalSpec && haveSpec
		if persona.CheckRunningProgram("Fork") && len(comm) > 3 {
			targetName2 := formatTargetName(comm[3])
			if targetName != targetName2 {
				if icon2, ok := ObjByNames[targetName2]; ok {
					if grid, ok := ObjByNames[targetName].(*TGrid); ok {
						targetList = nil
						targetList = append(targetList, grid)
						printLog("...Target 2: " + grid.GetGridName() + " has top priority")
						printLog("...Target 1 replaced")
						var targetValidSpecs []string
						targetValidSpecs = append(targetValidSpecs, "Hvs."+grid.GetType())
						targetValidSpecs = append(targetValidSpecs, "CyberSpec_vs."+grid.GetType())
						haveSpec, _ := persona.HaveValidSpec(targetValidSpecs)
						totalSpec = totalSpec && haveSpec
						printLog("--DEBUG--totalTargetSpec: " + strconv.FormatBool(totalSpec))
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
					printLog("--DEBUG--totalTargetSpec: " + strconv.FormatBool(totalSpec))

					printLog("...Target 2: " + newIcon2.GetName())

				}
			} else {
				printLog("...Error: Target 1 = Target 2")
			}
		}
	}
	printLog("--DEBUG--totalTargetSpec: " + strconv.FormatBool(totalSpec)) //TODO: доразделить проверку vs.Host на типы действий
	return targetList, totalSpec
}

func placeMARK(source, target IIcon) {
	currentMARKS := target.GetMarkSet().MarksFrom[source.GetID()]
	currentMARKS++
	if target.GetName() != player.GetName() {
		printLog("...new MARK on " + target.GetName() + " was successfuly planted")
	} else {
		printLog("...Warning: MARK on " + target.GetName() + " was confirmed")
	}
	if currentMARKS > 3 {
		currentMARKS = 3
	}
	target.GetMarkSet().MarksFrom[source.GetID()] = currentMARKS
	master := target.GetOwner()
	if masterIcon, ok := master.(IIcon); ok {
		if masterIcon != nil && masterIcon != target {
			placeMARK(source, masterIcon)
		}
	}
}

func revealData(persona IPersona, icon IIcon, needToReveal int) [30]string {
	canSee := persona.GetFieldOfView().KnownData[icon.GetID()]
	mem := make([]int, 0, 30)

	for i := range canSee {
		if canSee[i] == "Unknown" || canSee[i] == "" { //|| canSee[0] != "Spotted" {
			canSee[i] = "Unknown"
			if i != 29 {
				mem = append(mem, i)
			}
			//mem = append(mem, i)
		}
	}
	for i := rand.Intn(30); i > 0; i-- { //perception has a chance to reveal some data. TOPIC TO DISSCUSS: on what chance exactly is. Note: 30 stands for chance to reaveal 30 positions from X (33 here) picks
		shuffleInt(mem)
	}
	k := -1
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(fmt.Sprintf("Mem[]: %v", mem))
	for i := needToReveal; i > 0; i-- {
		mem = append(mem, 29)
		//mem = append(mem[:0], mem[1:]...)
		//congo.WindowsMap.ByTitle["Log"].WPrintLn(fmt.Sprintf("Mem[]: %v", mem))
		needToReveal--
		k++
		if k < len(mem) {
			choosen := mem[k]
			//printLog("k = "+strconv.Itoa(k)+"i = "+strconv.Itoa(i)+" | "+"choosen = "+strconv.Itoa(choosen)+" | "+"needToReveal = "+strconv.Itoa(needToReveal))
			if canSee[0] != "Spotted" {
				canSee[0] = "Spotted"
				persona.GetFieldOfView().KnownData[icon.GetID()] = canSee
				i++
				needToReveal++
				continue
			}
			if canSee[11] == "Unknown" {
				canSee[11] = icon.GetType()
				persona.GetFieldOfView().KnownData[icon.GetID()] = canSee
				i++
				needToReveal++
				continue
			}
			if canSee[choosen] != "Unknown" {
				mem = append(mem[:0], mem[1:]...)
				continue
			}
			switch choosen {
			case 1:
				if target, ok := icon.(IFile); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = target.GetLastEditDate()
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Last Edit Date = " + target.GetLastEditDate())
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 2:
				if target, ok := icon.(IPersona); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetMatrixCM())
					printLog("...Data revealed: ")
					printLog("..." + target.(IObj).GetName() + ": Matrix Condition Monitor = " + strconv.Itoa(target.GetMatrixCM()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
				if target, ok := icon.(IIC); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetMatrixCM())
					printLog("...Data revealed: ")
					printLog("..." + target.(IObj).GetName() + ": Matrix Condition Monitor = " + strconv.Itoa(target.GetMatrixCM()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 3:
				if target, ok := icon.(IFile); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetDataBombRating())
					printLog("...Data revealed: ")
					if target.GetDataBombRating() > 0 {
						printLog("..." + target.GetName() + ": Databomb Detected")
						printLog("..." + target.GetName() + ": Databomb Rating = " + strconv.Itoa(target.GetDataBombRating()))
					} else {
						printLog("..." + target.GetName() + ": No Databomb Detected")
					}
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 4:
				if target, ok := icon.(IHost); ok && target.GetHost() == persona.GetHost() {
					if target == persona.GetHost() {
						canSee[choosen] = "IC List Revealed"
						printLog("...Data revealed: ")
						icLIST := target.GetICState()
						for j := range icLIST.icName {
							congo.WindowsMap.ByTitle["Log"].WPrint("..." + target.GetName() + ": " + icLIST.icName[j] + " Detected ")
							if icLIST.icStatus[j] {
								printLog("(Status: Active)")
							} else {
								printLog("(Status: Passive)")
							}
						}
					}
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 5:
				if target, ok := icon.(IIcon); ok && target.GetType() != "File" && target.GetType() != "IC" {
					canSee[choosen] = strconv.Itoa(target.GetDeviceRating())
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Device Rating = " + strconv.Itoa(target.GetDeviceRating()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 7:
				if target, ok := icon.(IIcon); ok && target.GetType() != "File" && target.GetType() != "IC" && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetAttack())
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Attack = " + strconv.Itoa(target.GetAttack()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 8:
				if target, ok := icon.(IIcon); ok && target.GetType() != "File" && target.GetType() != "IC" && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetSleaze())
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Sleaze = " + strconv.Itoa(target.GetSleaze()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 9:
				if target, ok := icon.(IIcon); ok && target.GetType() != "File" && target.GetType() != "IC" && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetDataProcessing())
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Data Processing = " + strconv.Itoa(target.GetDataProcessing()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 10:
				if target, ok := icon.(IIcon); ok && target.GetType() != "File" && target.GetType() != "IC" && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetFirewall())
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Firewall = " + strconv.Itoa(target.GetFirewall()))
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 11:
				if target, ok := icon.(IIcon); ok {
					canSee[choosen] = target.GetUDevice()
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 12:
				if target, ok := icon.(IFile); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetEncryptionRating())
					printLog("...Data revealed: ")
					if target.GetDataBombRating() > 0 {
						printLog("..." + target.GetName() + ": Encryption rating detected")
						printLog("..." + target.GetName() + ": Encryption rating = " + strconv.Itoa(target.GetEncryptionRating()))
					} else {
						printLog("..." + target.GetName() + ": No file encryption detected")
					}
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 13:
				if target, ok := icon.(IHost); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = target.GetGrid().GetGridName()
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": Located in " + target.GetGrid().GetGridName())
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 15:
				if target, ok := icon.(IFile); ok && target.GetHost() == persona.GetHost() {
					canSee[choosen] = strconv.Itoa(target.GetSize())
					printLog("...Data revealed: ")
					printLog("..." + target.GetName() + ": File size evaluated")
					printLog("..." + target.GetName() + ": File size = " + strconv.Itoa(target.GetSize()) + " Mp")
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 18:
				if target, ok := icon.(IIcon); ok && target.GetType() != "IC" && target.GetHost() == persona.GetHost() {
					if target.GetOwner() == nil {
						canSee[choosen] = "No Owner"
						printLog("...Data revealed: ")
						printLog("..." + target.GetName() + " have no Owner")
					} else {
						canSee[choosen] = target.GetOwner().GetName()
						printLog("...Data revealed: ")
						printLog("..." + target.GetName() + ": Owner = " + target.GetOwner().GetName())
					}
				} else {
					canSee[choosen] = "Unknown"
					i++
					needToReveal++
				}
			case 29:
				//printLog("...Breaking: ")
				//i = 0
				//needToReveal = 0
				break

			default:
				canSee[choosen] = "Unknown"
				i++
				needToReveal++
				continue
			}
			persona.GetFieldOfView().KnownData[icon.GetID()] = canSee
			owner := persona.GetOwner()
			if owner != persona.(IObj) && owner != nil {
				//owner.GetFieldOfView().KnownData[icon.GetID()] = canSee
				for q := range owner.GetFieldOfView().KnownData[icon.GetID()] {
					if owner.GetFieldOfView().KnownData[icon.GetID()][q] == "Unknown" {
						owner.ChangeFOWParametr(icon.GetID(), q, canSee[q])
					}
					if owner.GetFieldOfView().KnownData[icon.GetID()][q] == "" {
						owner.ChangeFOWParametr(icon.GetID(), q, canSee[q])
					}
				}
			}
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
	allInfo[11] = "Unknown" //uDevice
	allInfo[12] = "Unknown" //Encrypt - file
	allInfo[13] = "Unknown" //Grid - obj
	allInfo[14] = "Unknown" //Special info
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

func retrieveLevel() int {
	comm := GetComm()
	level := 0
	for i := range comm {
		if strings.Contains(comm[i], "-L") {
			levelSTR := strings.Split(comm[i], "-L")
			levelINT, _ := strconv.Atoi(levelSTR[1])
			level = levelINT
		}
	}
	return level
}
