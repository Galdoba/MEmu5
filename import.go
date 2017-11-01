package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/Galdoba/ConGo/congo"
)

func fileDBExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//ImportHostFromDB -
func ImportHostFromDB(hostName string) *THost {
	if !fileDBExists("HostDB.txt") {
		panic("'HostDB.txt' is not exists. Please create file in the same directory as 'MEmu.exe'")
	}
	h := THost{}
	icList := new(ICList)
	h.icState = *icList
	allIC := make([]string, 0, 30)
	allIC = append(allIC, "Acid IC")
	allIC = append(allIC, "Binder IC")
	allIC = append(allIC, "Black IC")
	allIC = append(allIC, "Blaster IC")
	allIC = append(allIC, "Bloodhound IC")
	allIC = append(allIC, "Catapult IC")
	allIC = append(allIC, "Crash IC")
	allIC = append(allIC, "Jammer IC")
	allIC = append(allIC, "Killer IC")
	allIC = append(allIC, "Marker IC")
	allIC = append(allIC, "Patrol IC")
	allIC = append(allIC, "Probe IC")
	allIC = append(allIC, "Scramble IC")
	allIC = append(allIC, "Sparky IC")
	allIC = append(allIC, "Tar Baby IC")
	allIC = append(allIC, "Track IC")
	allIC = append(allIC, "Shoker IC")
	file, err := os.OpenFile("HostDB.txt", os.O_APPEND|os.O_WRONLY|os.O_RDWR, 0600) // открываем файл: Имя, ключи, что-то еще
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()                           //закрываем файл когда он уже не нужен
	dataDB, err := ioutil.ReadFile("HostDB.txt") //читаем файл
	if err != nil {
		log.Fatal(err)
	}
	allData := string(dataDB)
	hostToImport := ""
	hostData := strings.Split(allData, "####################\n")
	for i := range hostData {
		if strings.Contains(hostData[i], hostName) {
			hostToImport = hostData[i]
		}

	}

	lines := strings.Split(hostToImport, "\n")
	for i := range lines {
		//Import Name
		switch strings.Contains(lines[i], "Host: ") {
		case true:
			nameS := strings.Split(lines[i], "Host: ")
			name := nameS[1]
			h.name = name
		default:
		}
		//Import Grid
		switch strings.Contains(lines[i], "Grid: ") {
		case true:
			nameS := strings.Split(lines[i], "Grid: ")
			name := nameS[1]
			randoGo := true
			for i := range gridList {
				if gridForHost, ok := gridList[i].(*TGrid); ok {
					if gridForHost.name == name {
						h.grid = gridForHost
						randoGo = false
					}
				}
			}
			//Если сеть не известна - выбираем из известных
			for randoGo {
				r := rand.Intn(len(gridList))
				if gr, ok := gridList[r].(*TGrid); ok {
					h.grid = gr
					randoGo = false
				}
			}

		default:
		}
		//Import Rating
		switch strings.Contains(lines[i], "Rating: ") {
		case true:
			ratingS := strings.Split(lines[i], "Rating: ")
			rating, _ := strconv.Atoi(ratingS[1])
			h.deviceRating = rating
		default:
		}
		//Import Attack
		switch strings.Contains(lines[i], "Attack: ") {
		case true:
			attackS := strings.Split(lines[i], "Attack: ")
			attack, _ := strconv.Atoi(attackS[1])
			h.attack = attack
		default:
		}
		//Import Sleaze
		switch strings.Contains(lines[i], "Sleaze: ") {
		case true:
			sleazeS := strings.Split(lines[i], "Sleaze: ")
			sleaze, _ := strconv.Atoi(sleazeS[1])
			h.sleaze = sleaze
		default:
		}
		//Import Attack
		switch strings.Contains(lines[i], "Data Processing: ") {
		case true:
			dtpS := strings.Split(lines[i], "Data Processing: ")
			dtp, _ := strconv.Atoi(dtpS[1])
			h.dataProcessing = dtp
		default:
		}
		//Import Attack
		switch strings.Contains(lines[i], "Firewall: ") {
		case true:
			frwS := strings.Split(lines[i], "Firewall: ")
			frw, _ := strconv.Atoi(frwS[1])
			h.firewall = frw
		default:
		}
		//Import IC
		switch strings.Contains(lines[i], " >") {
		case true:
			icNamesS := strings.Split(lines[i], " >")
			icName := icNamesS[1]
			h.icState.icName = append(h.icState.icName, icName)
			h.icState.icStatus = append(h.icState.icStatus, false)
		default:
		}
	}
	//windowList[0].(*congo.TWindow).WPrintLn("Found: "+name+"...", congo.ColorYellow)
	//windowList[0].(*congo.TWindow).WPrintLn("Found: "+strconv.Itoa(rating)+"...", congo.ColorYellow)

	//h.name = name
	//h.faction = name
	h.alert = "No Alert"
	h.matrixCM = 999999
	h.SetID()
	h.markSet.MarksFrom = make(map[int]int)
	h.markSet.MarksFrom[h.id] = 4
	h.canSee.KnownData = make(map[int][30]string)
	//h.deviceRating = rating
	//	h.attack = attack
	//	h.sleaze = sleaze
	//	h.dataProcessing = dtp
	//	h.firewall = frw

	data := player.canSee.KnownData[h.id]
	data[0] = "Spotted"
	data[4] = "Unknown"
	data[5] = "Unknown"
	data[7] = "Unknown"
	data[8] = "Unknown"
	data[9] = "Unknown"
	data[10] = "Unknown"
	data[11] = "Unknown"
	data[13] = "Unknown"
	player.canSee.KnownData[h.id] = data
	h.FillHostWithFiles()
	h.LoadNextIC()
	printLog("Importing host: "+h.GetName(), congo.ColorDefault)
	gridList = append(gridList, &h)
	ObjByNames[h.name] = &h
	return &h
}

func ImportPlayerFromDB(alias string) *TPersona {
	if !fileDBExists("PlayerDB.txt") {
		//panic("'PlayerDB.txt' is not exists. Please create file in the same directory as 'MEmu.exe'")
		os.Create("PlayerDB.txt")

	}
	p := TPersona{}
	p.isPlayer = true
	//deviceName := "d"
	file, err := os.OpenFile("PlayerDB.txt", os.O_APPEND|os.O_WRONLY|os.O_RDWR, 0600) // открываем файл: Имя, ключи, что-то еще
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()                             //закрываем файл когда он уже не нужен
	dataDB, err := ioutil.ReadFile("PlayerDB.txt") //читаем файл
	if err != nil {
		log.Fatal(err)
	}
	allData := string(dataDB)
	playerToImport := ""
	playerData := strings.Split(allData, "####################")
	for i := range playerData {
		if strings.Contains(playerData[i], alias) {
			playerToImport = playerData[i]
		}

	}
	////////////////////////////////////////////

	lines := strings.Split(playerToImport, "\n")

	for i := range lines {
		//Import Name
		switch strings.Contains(lines[i], "Alias: ") {
		case true:
			nameS := strings.Split(lines[i], "Alias: ")
			name := strings.Trim(nameS[1], SPACES)
			p.name = name
			p.faction = name
		default:
		}
		//Import Device
		switch strings.Contains(lines[i], "Device: ") {
		case true:
			deviceNameS := strings.Split(lines[i], "Device: ")
			devName := deviceNameS[1]
			devName = strings.Trim(devName, SPACES)
			p.device = addDevice(devName)
			p.maxMatrixCM = p.device.GetMatrixCM()
		default:
		}
		//Import Computer Skill
		switch strings.Contains(lines[i], "Computer : ") {
		case true:
			ratingS := strings.Split(lines[i], "Computer : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.computerSkill = rating
		default:
		}
		//Import Cybercombat Skill
		switch strings.Contains(lines[i], "Cybercombat : ") {
		case true:
			ratingS := strings.Split(lines[i], "Cybercombat : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.cybercombatSkill = rating
		default:
		}
		//Import Hacking Skill
		switch strings.Contains(lines[i], "Hacking : ") {
		case true:
			ratingS := strings.Split(lines[i], "Hacking : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.hackingSkill = rating
		default:
		}
		//Import HardWare Skill
		switch strings.Contains(lines[i], "HardWare : ") {
		case true:
			ratingS := strings.Split(lines[i], "HardWare : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.hardwareSkill = rating
		default:
		}
		//Import Software Skill
		switch strings.Contains(lines[i], "Software : ") {
		case true:
			ratingS := strings.Split(lines[i], "Software : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
		}
		//Import Electronic Warfare Skill
		switch strings.Contains(lines[i], "Electronic Warfare : ") {
		case true:
			ratingS := strings.Split(lines[i], "Electronic Warfare : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.electronicSkill = rating
		default:
		}
		//Import BODY Attribute
		switch strings.Contains(lines[i], "BODY : ") {
		case true:
			ratingS := strings.Split(lines[i], "BODY : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.body = rating
		default:
		}
		//Import AGILITY Attribute
		switch strings.Contains(lines[i], "AGILITY : ") {
		case true:
			ratingS := strings.Split(lines[i], "AGILITY : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.AGILITY = 0
		}
		//Import REACTION Attribute
		switch strings.Contains(lines[i], "REACTION : ") {
		case true:
			ratingS := strings.Split(lines[i], "REACTION : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.REACTION = 0
		}
		//Import STRENGTH Attribute
		switch strings.Contains(lines[i], "STRENGTH : ") {
		case true:
			ratingS := strings.Split(lines[i], "STRENGTH : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//	p.STRENGTH = 0
		}
		//Import WILLPOWER Attribute
		switch strings.Contains(lines[i], "WILLPOWER : ") {
		case true:
			ratingS := strings.Split(lines[i], "WILLPOWER : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.willpower = rating
		default:
		}
		//Import LOGIC Attribute
		switch strings.Contains(lines[i], "LOGIC : ") {
		case true:
			ratingS := strings.Split(lines[i], "LOGIC : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.logic = rating
		default:
		}
		//Import INTUITION Attribute
		switch strings.Contains(lines[i], "INTUITION : ") {
		case true:
			ratingS := strings.Split(lines[i], "INTUITION : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.intuition = rating
		default:
		}
		//Import CHARISMA Attribute
		switch strings.Contains(lines[i], "CHARISMA : ") {
		case true:
			ratingS := strings.Split(lines[i], "CHARISMA : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.charisma = rating
		default:
		}
		//Import EDGE Attribute
		switch strings.Contains(lines[i], "EDGE : ") {
		case true:
			ratingS := strings.Split(lines[i], "EDGE : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.EDGE = 0
		}
		//Import RESONANCE Attribute
		switch strings.Contains(lines[i], "RESONANCE : ") {
		case true:
			ratingS := strings.Split(lines[i], "RESONANCE : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.RESONANCE = 0
		}
		//Import DEPTH Attribute
		switch strings.Contains(lines[i], "DEPTH : ") {
		case true:
			ratingS := strings.Split(lines[i], "DEPTH : ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.DEPTH = 0
		}

	}

	p.grid = gridList[0].(*TGrid) //временно - должен стартовать из публичной сети
	//p.maxMatrixCM = p.device.GetMatrixCM()
	p.matrixCM = p.maxMatrixCM
	p.id = id
	p.silentMode = true
	p.simSence = "Hot-SIM VR"
	p.maxStunCM = (p.GetWillpower()+1)/2 + 8
	p.stunCM = p.maxStunCM
	p.maxPhysCM = (p.body+1)/2 + 8
	p.physCM = p.maxPhysCM
	p.SetID()
	p.host = Matrix
	p.markSet.MarksFrom = make(map[int]int)
	p.markSet.MarksFrom[p.id] = 4
	p.linklocked.LockedByID = make(map[int]bool)
	//p.linklocked.LockedByID[p.id] = true
	p.canSee.KnownData = make(map[int][30]string)
	//p.(IIcon)name = alias
	p.connected = true
	p.physLocation = false

	objectList = append(objectList, &p)
	id++
	return &p
}
