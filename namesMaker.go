package main

import (
	"math/rand"
	"strconv"
	"time"
)

//import "github.com/Galdoba/ConGo/congo"

func generateFileName() string {
	setSeed()
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789")
	name := make([]rune, 8)
	extention := make([]rune, 3)
	for i := range name {
		name[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	for i := range extention {
		extention[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	fileName := string(name) + "." + string(extention)
	return fileName
}

func generateLastEditTime() string {
	setSeed()
	date := "00-00-00 00-00-2078"
	hours := rand.Intn(23)
	sHours := ""
	if hours < 10 {
		sHours = "0" + strconv.Itoa(hours)
	} else {
		sHours = strconv.Itoa(hours)
	}
	setSeed()
	minutes := rand.Intn(59)
	sMinutes := ""
	if minutes < 10 {
		sMinutes = "0" + strconv.Itoa(minutes)
	} else {
		sMinutes = strconv.Itoa(minutes)
	}
	setSeed()
	seconds := rand.Intn(59)
	sSeconds := ""
	if seconds < 10 {
		sSeconds = "0" + strconv.Itoa(seconds)
	} else {
		sSeconds = strconv.Itoa(seconds)
	}
	setSeed()
	days := rand.Intn(28) + 1
	sDays := ""
	if days < 10 {
		sDays = "0" + strconv.Itoa(days)
	} else {
		sDays = strconv.Itoa(days)
	}
	setSeed()
	months := rand.Intn(12) + 1
	sMonth := ""
	if months < 10 {
		sMonth = "0" + strconv.Itoa(months)
	} else {
		sMonth = strconv.Itoa(months)
	}
	setSeed()
	sYear := strconv.Itoa(2079 - rand.Intn(5))

	//date = strconv.Itoa(hours) + "-" + strconv.Itoa(minutes) + "-" + strconv.Itoa(seconds) + " " + strconv.Itoa(days) + "-" + strconv.Itoa(months) + "-" + year
	date = sYear + "-" + sMonth + "-" + sDays + " " + sHours + ":" + sMinutes + ":" + sSeconds
	return date
}

func generateCurrentTime() string {
	setSeed()
	cTime := time.Now()
	SrTime = cTime.AddDate(62, -1, -14)
	//diference := time.Now().Sub(cTime)
	date := (SrTime.String())
	date = (SrTime.Format("2006-01-02 15:04:05"))
	/*	sDate := strings.Split(date, ":")
		year := sDate[0]
		month := sDate[1]
		day := sDate[2]
		hour := sDate[3]
		minute := sDate[4]
		second := sDate[5]
		intYear, _ := strconv.Atoi(year) //+ 61
		intYear = intYear + 61
		intMonth, _ := strconv.Atoi(month) // - 1
		intMonth = intMonth - 1

		year = strconv.Itoa(intYear)
		if intMonth < 10 {
			month = "0" + strconv.Itoa(intMonth)
		} else {
			month = strconv.Itoa(intMonth)
		}
		if
		//month = strconv.Itoa(intMonth)
		test := year + month + day + "date" + hour + minute + second*/

	return date
}

func forwardShadowrunTime() string {
	//cTime := generateCurrentTime()
	SrTime = SrTime.Add(3 * time.Second)
	date := (SrTime.String())
	date = (SrTime.Format("2006-01-02 15:04:05"))
	//SrTime = SrTime.Add(3 * time.Second)
	return date
}

/*func test(trg IObj) {
	//var testObj interface{}
	if trg, ok := trg.(IPersona); ok {

	}
	switch trg.GetType() {
	case "Persona":
		testObj := trg.(IPersona)
	case "File":
		testObj := trg.(IFile)
	default:
		endAction()
	}

	testObj.GetName()
}*/
