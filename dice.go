/*
How to Call Tests outside of "dice.go"
opposedTest()  =>   opposedTest(dicePoolSrc int, dicePoolTrgt int , limit int) (return netHits int, glitch bool, critGlitch bool)

simpleTest()   =>   simpleTest(dicePoolSrc int, limit int, threshold int) (return netHits int, glitch bool, critGlitch bool)

xd6Test()      =>   xd6Test(dicePoolSrc) (return summ int)

extendedTest() => extendedTest(dicePoolSrc int, limit int, threshold int) (return netHits int, glitch bool, critGlitch bool)
*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/ConGo/congo"
)

//DicePool -
type DicePool struct {
	pool     []int
	isOk     bool
	isRolled bool
}

func makeDicePool(size int) *DicePool {
	var dp DicePool
	dp.pool = make([]int, size, size*2)
	dp.isOk = true
	return &dp
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())  //получаем рандомный сид от текущего времени с точностью до наносекунд
	time.Sleep(time.Millisecond * 30) //ждем 3 милисекунды чтобы сид гарантированно сменился к следующему заходу
}

func (dp *DicePool) roll() {
	assert(dp.isOk, "DicePool not initialized")
	setSeed()
	//windowList[5].(*congo.TWindow).WPrint("& ", congo.ColorGreen)
	//windowList[0].(*congo.TWindow).WPrint("...", congo.ColorGreen)
	for i := range dp.pool {
		dp.pool[i] = rand.Intn(6) + 1
		//	windowList[5].(*congo.TWindow).WPrint(".", congo.ColorGreen)
		draw()
		if src, ok := SourceIcon.(IPersona); ok {
			//				windowList[0].(*congo.TWindow).WPrint(strconv.Itoa(dp.pool[i])+" ", congo.ColorGreen)
			draw()
			//				hold()
			src.(IObj).GetID()
		}
	}
	//windowList[0].(*congo.TWindow).WPrintLn("...", congo.ColorGreen)
	//windowList[5].(*congo.TWindow).WPrint(".", congo.ColorGreen)
	dp.isRolled = true
}

func (dp *DicePool) successes() int {
	assert(dp.isRolled, "DicePool not Rolled")
	successes := 0
	for i := range dp.pool {
		if dp.pool[i] == 5 || dp.pool[i] == 6 {
			successes++
		}
	}
	//fmt.Println(successes)
	return successes
}

func (dp *DicePool) sixes() int {
	assert(dp.isRolled, "DicePool not Rolled")
	sixes := 0
	for i := range dp.pool {
		if dp.pool[i] == 6 {
			sixes++
		}
	}
	//fmt.Println(successes)
	return sixes
}

func (dp *DicePool) glitch() bool {
	assert(dp.isRolled, "DicePool not Rolled")
	glitch := false
	ones := 0
	for i := range dp.pool {
		if dp.pool[i] == 1 {
			ones++
		}
	}
	//fmt.Println(ones)
	if ones > len(dp.pool)/2 {
		glitch = true
	}
	return glitch
}

func (dp *DicePool) critGlitch() bool {
	assert(dp.isRolled, "DicePool not Rolled")
	critGlitch := false
	if dp.successes() == 0 && dp.glitch() {
		critGlitch = true
	}
	return critGlitch
}

func (dp *DicePool) summ() int {
	assert(dp.isRolled, "DicePool not Rolled")
	xd6 := 0
	total := 0
	for i := range dp.pool {
		xd6 = dp.pool[i]
		total = total + xd6
		//	windowList[0].(*congo.TWindow).WPrint(strconv.Itoa(dp.pool[i])+" ", congo.ColorGreen)
	}
	//windowList[0].(*congo.TWindow).WPrintLn(" ", congo.ColorGreen)
	//fmt.Println(xd6)
	//windowList[0].(*congo.TWindow).WPrint(strconv.Itoa(total)+" ", congo.ColorGreen)
	return total
}

func reRoll(rerollDp int) (int, bool, bool) {
	if rerollDp < 0 {
		return 0, false, false
	}
	sourceIcon := makeDicePool(rerollDp)
	sourceIcon.roll()
	suc := sourceIcon.successes()
	glitch := sourceIcon.glitch()
	critGlitch := sourceIcon.critGlitch()
	return suc, glitch, critGlitch
}

func interruptProcess(i int) {
	congo.WindowsMap.ByTitle["Process"].WClear()
	congo.WindowsMap.ByTitle["Process"].WPrintLn("Interrupt protocol Keys: ('S' - Skip timer)", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Process"].WPrintLn("'R' - Reroll", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Process"].WPrintLn("'N' - Negate Glitch", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Process"].WPrintLn("'P' - Push the Limit", congo.ColorDefault)
	congo.WindowsMap.ByTitle["Process"].WPrintLn("Deactivation in : "+strconv.Itoa(i/4), congo.ColorDefault)

}

func rule6Test(rollerID int, dicePool1 int, limit int, threshold int) (int, bool, bool) {

	if dicePool1 < 1 {
		return 0, false, false
	}
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	sixes := sourceIcon.sixes()
	suc := sourceIcon.successes()
	glitch := sourceIcon.glitch()
	critGlitch := sourceIcon.critGlitch()
	if rollerID == player.GetID() {
		congo.WindowsMap.ByTitle["Log"].WPrint("......Performance Array: ", congo.ColorGreen)
		for i := range sourceIcon.pool {
			congo.WindowsMap.ByTitle["Log"].WPrint(strconv.Itoa(sourceIcon.pool[i])+" ", congo.ColorGreen)
			draw()
		}
		congo.WindowsMap.ByTitle["Log"].WPrintLn("", congo.ColorGreen)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("......Result: "+strconv.Itoa(suc)+" successes", congo.ColorGreen)
		if sixes > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("......Reevaluating resources: "+strconv.Itoa(sixes)+" Mp/p available", congo.ColorGreen)
		}
	}
	for sixes > 0 {
		//printLog("create DP = sixes ("+strconv.Itoa(sixes), congo.ColorDefault)
		//	hold()
		addDp := makeDicePool(sixes)
		//printLog("roll DP = addDp.pool ("+strconv.Itoa(len(addDp.pool)), congo.ColorDefault)
		addDp.roll()
		addSuc := addDp.successes()
		suc = suc + addSuc
		sixes = addDp.sixes()
		if rollerID == player.GetID() && len(addDp.pool) > 0 {
			congo.WindowsMap.ByTitle["Log"].WPrint("......Performance Array: ", congo.ColorGreen)
			for i := range addDp.pool {
				congo.WindowsMap.ByTitle["Log"].WPrint(strconv.Itoa(addDp.pool[i])+" ", congo.ColorGreen)
				draw()
			}
			congo.WindowsMap.ByTitle["Log"].WPrintLn("", congo.ColorGreen)

			printLog("......Result: "+strconv.Itoa(addSuc)+" successes", congo.ColorGreen)
			if sixes > 0 {
				printLog("......Reevaluating resources: "+strconv.Itoa(sixes)+" Mp/p available", congo.ColorGreen)
			}

		}
	}
	printLog("...Final Result: "+strconv.Itoa(suc), congo.ColorGreen)
	return suc, glitch, critGlitch
	//return 0, false, false
}

func simpleTest(rollerID int, dicePool1 int, limit int, threshold int) (int, bool, bool) {
	if rollerID == player.GetID() {
		text := command
		text = formatString(text)
		text = cleanText(text)
		comm := strings.Split(text, ">")
		for i := range comm {
			if comm[i] == "-PL" {
				printLog("...Additional resources evaluated", congo.ColorGreen)
				if player.GetEdge() > 0 {
					dicePool1 = dicePool1 + player.GetMaxEdge()
					printLog("...Expected efficiency "+strconv.Itoa(dicePool1)+" Mp/p", congo.ColorGreen)
					limit = 999
					printLog("...Warning: Hardware limit deactivated", congo.ColorYellow)
					suc, gl, cgl := rule6Test(rollerID, dicePool1, limit, threshold)
					netHits := (suc - threshold)
					player.SetEdge(player.GetEdge() - 1)
					return netHits, gl, cgl
				}
				//dicePool1 = dicePool1 + player.GetMaxEdge()

			}

		}
	}
	if dicePool1 < 0 {
		return 0, false, false
	}
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	if rollerID == player.GetID() {
		congo.WindowsMap.ByTitle["Log"].WPrint("......Performance Array: ", congo.ColorGreen)
		for i := range sourceIcon.pool {
			congo.WindowsMap.ByTitle["Log"].WPrint(strconv.Itoa(sourceIcon.pool[i])+" ", congo.ColorGreen)
		}
		congo.WindowsMap.ByTitle["Log"].WPrintLn("", congo.ColorGreen)

	}
	suc := sourceIcon.successes()
	//sixes := sourceIcon.sixes()
	glitch := sourceIcon.glitch()
	critGlitch := sourceIcon.critGlitch()
	if rollerID == player.GetID() && player.GetEdge() > 0 {
		printLog("......Roll result: "+strconv.Itoa(suc)+" successes", congo.ColorGreen)
		if glitch {
			printLog("......Error: Glitch detected!", congo.ColorYellow)
		}
		if critGlitch {
			printLog("......Warning: Error critical!", congo.ColorRed)
		}
		printLog("......Interrupt protocol ready", congo.ColorGreen)
		for i := 20; i > 0; i-- {
			congo.WindowsMap.ByTitle["Process"].WClear()
			interruptProcess(i)
			hold()
			if congo.GetKeyboard().KeyPressed() {
				ev := congo.GetKeyboard().ReadEvent()
				if ev.GetEventType() == "Keyboard" {
					key := ev.(*congo.KeyboardEvent).GetRune()
					if key != 0 {
						char := string(key)
						if char == "r" {
							printLog("......Reroll protocol Initiated", congo.ColorYellow)
							player.SetEdge(player.GetEdge() - 1)
							rerollDp := dicePool1 - suc
							printLog("......Rerolling "+strconv.Itoa(rerollDp)+" dices...", congo.ColorGreen)
							suc2, sgl, scgl := reRoll(rerollDp)
							if sgl {
								glitch = sgl
							}
							if scgl {
								critGlitch = scgl
							}
							suc = suc + suc2
							printLog("......Reroll result: "+strconv.Itoa(suc2)+" successes", congo.ColorGreen)
							printLog("...Final result: "+strconv.Itoa(suc2+suc), congo.ColorGreen)
							break
						}
						if char == "n" {
							printLog("......Emergency data steam rerouting", congo.ColorYellow)
							player.SetEdge(player.GetEdge() - 1)
							if critGlitch {
								critGlitch = false
								break
							}
							if glitch {
								glitch = false
								break
							}
							break
						}
						if char == "p" {
							printLog("......Allocating additional resourses", congo.ColorYellow)
							player.SetEdge(player.GetEdge() - 1)
							limit = 999
							addSucc, _, _ := rule6Test(player.GetID(), player.GetMaxEdge(), limit, 0)
							suc = suc + addSucc
							break
						}
						if char == "s" {
							printLog("......Interrupt protocol deactivated", congo.ColorGreen)
							break
						}
					}
				}
			}
		}
		congo.WindowsMap.ByTitle["Process"].WClear()
	}
	if suc > limit {
		suc = limit
	}
	netHits := (suc - threshold)
	if glitch == true {
		if critGlitch == true {
			//windowList[0].(*congo.TWindow).WPrintLn("CRITICAL GLITCH!!!", congo.ColorGreen)
		} else {
			//windowList[0].(*congo.TWindow).WPrintLn("GLITCH!", congo.ColorGreen)
		}
	}
	/*if rollerID == player.GetID() {
		printLog("...Roll Result: ", congo.ColorGreen)
	}*/
	return netHits, glitch, critGlitch
}

func opposedTest(dicePool1 int, dicePool2 int, limit int) (int, int, bool, bool) {
	suc1 := 0
	suc2 := 0
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	targetIcon := makeDicePool(dicePool2)
	targetIcon.roll()
	suc1 = sourceIcon.successes()
	suc2 = targetIcon.successes()
	if suc1 > limit {
		windowList[0].(*congo.TWindow).WPrintLn("Succeses by Limit: "+strconv.Itoa(limit), congo.ColorYellow)
		//fmt.Println("Succeses by Limit:", limit)
		suc1 = limit
	}
	windowList[0].(*congo.TWindow).WPrintLn("Source sucesesses: "+strconv.Itoa(suc1)+"; Target successes: "+strconv.Itoa(suc2), congo.ColorYellow)
	//fmt.Println("Source sucesesses =", suc1, "Target successes =", suc2)
	//netHits := suc1 - suc2
	glitch := sourceIcon.glitch()
	critGlitch := sourceIcon.critGlitch()
	//fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return suc1, suc2, glitch, critGlitch
}

func xd6Test(dicePool1 int) int {
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	summ := 0
	summ = sourceIcon.summ()
	return summ
}

func extendedTest(dicePool1 int, limit int, threshold int) (int, bool, bool) {
	netHits := 0
	glitch := false
	critGlitch := false
	step := 1
	i := 0
	for i = dicePool1; i > 0; i-- {
		fmt.Println("Step =", step)
		sourceIcon := makeDicePool(i)
		sourceIcon.roll()
		netHits = netHits + sourceIcon.successes()
		if sourceIcon.successes() == 0 {
			i++
		}
		if sourceIcon.glitch() == true {
			glitch = true
			fmt.Println("glitch")
			threshold = threshold + 2
		}
		if sourceIcon.critGlitch() == true {
			critGlitch = true
			fmt.Println("critGlitch")
			netHits = 0 // for making sure that test fail
			break
		}
		fmt.Println("Sucesesses =", sourceIcon.successes())
		if netHits >= threshold {
			break
		}
		step++
		fmt.Println("Nethits =", netHits, "Threshold =", threshold)
	}
	fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return netHits, glitch, critGlitch
}

func assert(ok bool, s string) {
	if !ok {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("\033[31m\n%s:%d\n[error] %s\033[0m", file, line, s))
	}
}
