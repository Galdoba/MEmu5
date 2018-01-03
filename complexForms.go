package main

//CFNumber -
var CFNumber int

//ComplexForm -
type ComplexForm struct {
	formNum           int
	madeByID          int
	madeOnID          int
	cfName            string
	level             int
	succ              int
	turnsActive       int
	durationCode      string
	threaderResonance int
}

func getComplexForm(formNum int) ComplexForm {
	return CFDBMap[formNum]
}

func deleteComplexForm(srcID int, formName string) {
	//
}

//TreadComplexForm -
func TreadComplexForm(srcID, trgID int, formName string, formLevel, formSucc int) bool {
	for i := range CFDBMap {
		if getComplexForm(i).madeOnID == trgID && getComplexForm(i).cfName == formName && getComplexForm(i).madeByID == srcID { //if this Form from this src on this trg exist
			if getComplexForm(i).level < formLevel { //compare effect and kill weakest
				delete(CFDBMap, getComplexForm(i).formNum)
				break
			} else {
				return false
			}

		}
	}
	CFNumber++
	CFDBMap[CFNumber] = ComplexForm{
		CFNumber,
		srcID,
		trgID,
		formName,
		formLevel,
		formSucc,
		0,
		designFormDuration(formName),
		pickObjByID(srcID).(ITechnom).GetResonance(),
	}
	//pickObjByID(srcID).(ITechnom).GetResonance()
	return true
}

func designFormDuration(name string) string {
	switch name {
	default:
		return "S"
	}
	return "--Unknown--"
}

func countSustainedForms(srcID int) int {
	sustainedFormsCount := 0
	for i := range CFDBMap {
		if getComplexForm(i).madeByID == srcID {
			sustainedFormsCount++
		}
	}
	return sustainedFormsCount
}

func calculateFade(formName string, lvl int) int {
	fade := 0
	switch formName {
	default:
	case "Diffusion of Attack":
		fade = lvl - 2
	case "Diffusion of Sleaze":
		fade = lvl - 2
	case "Diffusion of Data Processing":
		fade = lvl - 2
	case "Diffusion of Firewall":
		fade = lvl - 2
	case "Editor":
		fade = lvl - 1
	case "Infusion of Attack":
		fade = lvl - 2
	case "Infusion of Sleaze":
		fade = lvl - 2
	case "Infusion of Data Processing":
		fade = lvl - 2
	case "Infusion of Firewall":
		fade = lvl - 2
	case "Static Veil":
		fade = lvl - 3
	case "Pulse Storm":
		fade = lvl - 3
	case "Puppeteer":
		fade = lvl + 1
	case "Resonanse Channel":
		fade = lvl - 3
	case "Resonance Spike":
		fade = lvl - 3
	case "Resonance Veil":
		fade = lvl - 3
	case "Static Bomb":
		fade = lvl - 1
	case "Stiches":
		fade = lvl - 3
	case "Tattletale":
		fade = lvl - 3
	}
	if fade < 2 {
		fade = 2
	}
	return fade
}

/*
Status: Official

Fading Value Changes for Complex Forms (P. 252-3, Resonance Library)
The Fading Values for complex forms should be updated as follows. Note that the minimum Fading Value for a complex form is 2 (Threading, p. 251).

Cleaner: L–2
Diffusion of [Matrix Attribute]: L–2
Editor: L–1
Infusion of [Matrix Attribute]: L–2
Static Veil: L–3
Pulse Storm: L–3
Puppeteer: L+1
Resonance Channel: L–3
Resonance Spike: L–3
Resonance Veil: L–3
Static Bomb: L–1
Stitches: L–3
Tattletale: L–3
*/
