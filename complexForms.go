package main

//DB -
//var CFDB TCompFormsDB
var CFNumber int

type ComplexForm struct {
	formNum      int
	madeByID     int
	madeOnID     int
	cfName       string
	level        int
	succ         int
	turnsActive  int
	durationCode string
}

func getComplexForm(formNum int) ComplexForm {
	return CFDBMap[formNum]
}

func deleteComplexForm(srcID int, formName string) {
	//
}

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
	}
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
