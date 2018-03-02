package main

import (
	//"strconv"

	"strings"

	"github.com/Galdoba/ConGo/congo"
	//"encoding/base64"
	"encoding/hex"
)

//UserInput -
func UserInput(input string) bool {
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("Processing: '"+input+"'")
	var mActionName string
	var actionIsGood bool
	var comm []string
	command = input
	command = formatString(command)
	comm = strings.SplitN(command, ">", 6)
	text := formatString(input)
	text = cleanText(text)
	if text != "" {
		printLog(text)
	}
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(text)
	//hold()

	if len(comm) < 2 {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("WARNING! Sintax Error!")
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [TARGET]' Format")
		congo.WindowsMap.ByTitle["Log"].WDraw()
		return false
	}
	//////Checking if action isValid
	mAction := comm[1]
	mAction = formatString(mAction)
	mAction = cleanText(mAction)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(mAction)
	//printLog(mAction)
	mActionName = mAction
	//printLog(mActionName)
	actionIsGood, mActionName = checkAction(mAction)
	if actionIsGood == false {
		//	congo.WindowsMap.ByTitle["Log"].WPrintLn("Action: "+mActionName+" is correct", congo.ColorYellow)
		return false
	}
	checkSource(comm[0])
	if mActionName == "EXIT_HOST" || mActionName == "ERASE_MARK" || mActionName == "CHECK_OVERWATCH_SCORE" || mActionName == "LONGACT" || mActionName == "SWITCH_INTERFACE_MODE" || mActionName == "SILENT_MODE" {
		TargetIcon = SourceIcon
		doAction(mActionName)
		return true
	}
	if mActionName == "WAIT" || mActionName == "FULL_DEFENCE" {
		//TargetIcon = text
		TargetIcon = SourceIcon
		doAction(mActionName)
		return true
	}
	if mActionName == "SWAP_ATTRIBUTES" || mActionName == "SWAP_PROGRAMS" {
		TargetIcon = SourceIcon
		if len(comm) < 4 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not enough data...", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [ATTRIBUTE 1] > [ATTRIBUTE 2]' Format")
			return false
		} else if len(comm) == 4 {
			TargetIcon = SourceIcon
			doAction(mActionName)
			return true
		}
		return false
	}
	if mActionName == "LOAD_PROGRAM" || mActionName == "UNLOAD_PROGRAM" || mActionName == "LOGIN" || mActionName == "COMPILE" {
		TargetIcon = SourceIcon
		if len(comm) < 3 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not enough data...", congo.ColorYellow)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [ATTRIBUTE 1] > [ATTRIBUTE 2]' Format")
			return false
		}
		TargetIcon = SourceIcon
		doAction(mActionName)
		return true

		//return false
	}
	if mActionName == "MATRIX_PERCEPTION" {
		if len(comm) > 2 {
			target := comm[2]
			target = formatString(target)
			target = cleanText(target)
			if target == "ALL" {
				mActionName = "SCAN_ENVIROMENT"
				doAction(mActionName)
				return true
			}
		}
	}
	if mActionName == "MATRIX_SEARCH" {
		//TargetIcon = text
		TargetIcon = SourceIcon
		if len(comm) < 3 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Not enough data...", congo.ColorYellow)
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [TARGET TYPE] > [TARGET NAME]' Format")
			congo.WindowsMap.ByTitle["Log"].WPrintLn("[TARGET NAME] is optional, if left blank random name will be generated")
			return false
		}
		doAction(mActionName)
		return true
	}
	if len(comm) < 3 {
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("WARNING! Sintax Error!")
		//congo.WindowsMap.ByTitle["Log"].WPrintLn("Use '[SOURCE] > [COMMAND] > [TARGET]' Format")
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: Target not designated", congo.ColorYellow)
		congo.WindowsMap.ByTitle["Log"].WDraw()
		return false
	}
	//printLog(comm[2], congo.ColorYellow)

	if checkTarget(comm[2], mActionName) == true {
		doAction(mActionName)
	} else {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: Unknown target")

	}
	congo.WindowsMap.ByTitle["Log"].WDraw()
	return true
}

func formatString(s string) string {
	s = strings.ToUpper(s)
	s = strings.Replace(s, " ", "_", -1)
	//s = strings.Replace(s, "-2M", "-2m", -1)
	//s = strings.Replace(s, "-3M", "-3m", -1)
	return s
}

func checkSource(source string) bool {
	source = formatString(source)
	source = cleanText(source)
	isGood := false
	//var alias string
	for _, obj := range ObjByNames {
		if srcObj, ok := obj.(IPersona); ok {
			//if srcObj.(IIcon).GetType() == "Persona" {
			alias := string(srcObj.(IPersona).GetName())
			alias = formatString(alias)
			s := (hex.EncodeToString([]byte(source)))
			a := (hex.EncodeToString([]byte(alias)))
			if a == s {
				//congo.WindowsMap.ByTitle["Log"].WPrintLn("SourceIcon is " + objectList[i].(*TPersona).GetName(), congo.ColorYellow)
				if obj.(IPersona).IsPlayer() == true {
					SourceIcon = srcObj
					isGood = true
					return isGood
				}

			}
		}

	}
	return isGood
}

func checkTarget(target, mActionName string) bool {
	self := formatString(target)
	self = cleanText(self)
	if self == "SELF" {
		TargetIcon = player
		return true
	}
	if pickTarget(target, mActionName) {
		return true
	}
	//printLog("Error: target not found")
	return false
}

func pickTarget(target, mActionName string) bool {
	target = formatString(target)
	target = cleanText(target)
	for _, obj := range ObjByNames {
		if grid, ok := obj.(IGrid); ok {
			var alias string
			alias = grid.GetName()
			alias = formatString(alias)
			alias = cleanText(alias)
			if alias == target {
				TargetIcon = grid
				return true
			}
		}
		if icon, ok := obj.(IIcon); ok {
			var alias string
			alias = icon.GetName()
			alias = formatString(alias)
			alias = cleanText(alias)
			if alias == target {
				TargetIcon = icon
				return true
			}
		}
	}
	return false
}

func cleanText(s string) string {
	out := ""
	plain := hex.EncodeToString([]byte(s))
	//plain = strings.Replace(plain, "10", "", -1)
	char := strings.Split(plain, "1001")
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(plain)
	for i := range char {

		if char[i] == "10" {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("remove \x10")
		} else if char[i] == "01" {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("remove \x01", congo.ColorYellow)
		} else {
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("keep x__")
			out = out + char[i]
		}
	}
	hexOut, _ := hex.DecodeString(out)

	return string(hexOut)

}
