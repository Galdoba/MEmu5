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
	hostData := strings.Split(allData, "####################\r\n")
	for i := range hostData {
		if strings.Contains(hostData[i], hostName) {
			hostToImport = hostData[i]
		}

	}

	lines := strings.Split(hostToImport, "\r\n")
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

	data := player.GetFieldOfView().KnownData[h.id]
	data[0] = "Spotted"
	data[4] = "Unknown"
	data[5] = "Unknown"
	data[7] = "Unknown"
	data[8] = "Unknown"
	data[9] = "Unknown"
	data[10] = "Unknown"
	data[11] = "Unknown"
	data[13] = "Unknown"
	player.GetFieldOfView().KnownData[h.id] = data
	h.FillHostWithFiles()
	h.LoadNextIC()
	//printLog("Importing host: "+h.GetName(), congo.ColorDefault)
	gridList = append(gridList, &h)
	ObjByNames[h.name] = &h
	return &h
}

//ImportPlayerFromDB0 -
func ImportPlayerFromDB0(alias string) (IPersona, bool) {
	if !fileDBExists("PlayerDB.txt") {
		//panic("'PlayerDB.txt' is not exists. Please create file in the same directory as 'MEmu.exe'")
		os.Create("PlayerDB.txt")

	}
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
	found := false
	for i := range playerData {
		if strings.Contains(playerData[i], alias) {
			playerToImport = playerData[i]
			found = true
		}

	}
	if !found {
		printLog("...Incorrect username or passcode", congo.ColorRed)
		return player, false
	}
	////////////////////////////////////////////
	p := TPersona{}
	p.isPlayer = true
	lines := strings.Split(playerToImport, "\r\n")
	class := "Decker"
	for i := range lines {
		//Import Device
		switch strings.Contains(lines[i], "Class: ") {
		case true:
			deviceNameS := strings.Split(lines[i], "Class: ")
			devName := deviceNameS[1]
			if devName == "Technomancer" {
				class = "Technomancer"
			}
		default:
		}
	}

	if class == "Decker" {
		printLog("Create Decker", congo.ColorDefault)
	}
	var name string
	var devName string
	var specs []string
	var cyberCombRating int

	for i := range lines {
		//Import Name
		switch strings.Contains(lines[i], "Alias: ") {
		case true:
			nameS := strings.Split(lines[i], "Alias: ")
			name = strings.Trim(nameS[1], SPACES)
			p.name = name
			p.faction = name
		default:
		}
		//Import Device
		switch strings.Contains(lines[i], "Device: ") {
		case true:
			deviceNameS := strings.Split(lines[i], "Device: ")
			devName = deviceNameS[1]
			devName = strings.Trim(devName, SPACES)
			p.device = addDevice(devName)
			p.maxMatrixCM = p.GetDevice().GetMatrixCM()
		default:
		}
		//Import Computer Skill
		switch strings.Contains(lines[i], "Computer: ") {
		case true:
			ratingS := strings.Split(lines[i], "Computer: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.computerSkill = rating
		default:
		}
		//Import Computer Specialization
		switch strings.Contains(lines[i], "ComputerSpec: ") {
		case true:
			specListS := strings.Split(lines[i], "ComputerSpec: ")
			specList := specListS[1]
			specList = strings.Trim(specList, SPACES)
			specs = strings.Split(specList, ";")
			for i := range specs {
				specs[i] = "CompSpec_" + specs[i]
				p.specialization = append(p.specialization, specs[i])
			}
			//p.specialization = specs
		default:
		}
		//Import Cybercombat Skill
		switch strings.Contains(lines[i], "Cybercombat: ") {
		case true:
			ratingS := strings.Split(lines[i], "Cybercombat: ")
			cyberCombRating, _ = strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.cybercombatSkill = cyberCombRating
		default:
		}
		//Import CyberCombat Specialization
		switch strings.Contains(lines[i], "CybercombatSpec: ") {
		case true:
			specListS := strings.Split(lines[i], "CybercombatSpec: ")
			specList := specListS[1]
			specList = strings.Trim(specList, SPACES)
			specs := strings.Split(specList, ";")
			for i := range specs {
				specs[i] = "CyberSpec_" + specs[i]
				p.specialization = append(p.specialization, specs[i])
			}
		default:
		}
		//Import Hacking Skill
		switch strings.Contains(lines[i], "Hacking: ") {
		case true:
			ratingS := strings.Split(lines[i], "Hacking: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.hackingSkill = rating
		default:
		}
		//Import Hacking Specialization
		switch strings.Contains(lines[i], "HackingSpec: ") {
		case true:
			specListS := strings.Split(lines[i], "HackingSpec: ")
			specList := specListS[1]
			specList = strings.Trim(specList, SPACES)
			specs := strings.Split(specList, ";")
			for i := range specs {
				specs[i] = "HackingSpec_" + specs[i]
				p.specialization = append(p.specialization, specs[i])
			}
		default:
		}
		//Import HardWare Skill
		switch strings.Contains(lines[i], "Hardware: ") {
		case true:
			ratingS := strings.Split(lines[i], "Hardware: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.hardwareSkill = rating
		default:
		}
		//Import Hardware Specialization
		switch strings.Contains(lines[i], "HardwareSpec: ") {
		case true:
			specListS := strings.Split(lines[i], "HardwareSpec: ")
			specList := specListS[1]
			specList = strings.Trim(specList, SPACES)
			specs := strings.Split(specList, ";")
			for i := range specs {
				specs[i] = "HardwareSpec_" + specs[i]
				p.specialization = append(p.specialization, specs[i])
			}
		default:
		}
		//Import Software Skill
		switch strings.Contains(lines[i], "Software: ") {
		case true:
			ratingS := strings.Split(lines[i], "Software: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
		}
		//Import Software Specialization
		switch strings.Contains(lines[i], "SoftwareSpec: ") {
		case true:
			specListS := strings.Split(lines[i], "SoftwareSpec: ")
			specList := specListS[1]
			specList = strings.Trim(specList, SPACES)
			specs := strings.Split(specList, ";")
			for i := range specs {
				specs[i] = "SoftwareSpec_" + specs[i]
				p.specialization = append(p.specialization, specs[i])
			}
		default:
		}
		//Import Electronic Warfare Skill
		switch strings.Contains(lines[i], "Electronic Warfare: ") {
		case true:
			ratingS := strings.Split(lines[i], "Electronic Warfare: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.electronicSkill = rating
		default:
		}
		//Import Electronic Warfare Specialization
		switch strings.Contains(lines[i], "Electronic WarfareSpec: ") {
		case true:
			specListS := strings.Split(lines[i], "Electronic WarfareSpec: ")
			specList := specListS[1]
			specList = strings.Trim(specList, SPACES)
			specs := strings.Split(specList, ";")
			for i := range specs {
				specs[i] = "Electronic WarfareSpec_" + specs[i]
				p.specialization = append(p.specialization, specs[i])
			}
		default:
		}
		//Import BODY Attribute
		switch strings.Contains(lines[i], "BODY: ") {
		case true:
			ratingS := strings.Split(lines[i], "BODY: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.body = rating
		default:
		}
		//Import AGILITY Attribute
		switch strings.Contains(lines[i], "AGILITY: ") {
		case true:
			ratingS := strings.Split(lines[i], "AGILITY: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.AGILITY = 0
		}
		//Import REACTION Attribute
		switch strings.Contains(lines[i], "REACTION: ") {
		case true:
			ratingS := strings.Split(lines[i], "REACTION: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.reaction = rating
		default:
			//p.REACTION = 0
		}
		//Import STRENGTH Attribute
		switch strings.Contains(lines[i], "STRENGTH: ") {
		case true:
			ratingS := strings.Split(lines[i], "STRENGTH: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//	p.STRENGTH = 0
		}
		//Import WILLPOWER Attribute
		switch strings.Contains(lines[i], "WILLPOWER: ") {
		case true:
			ratingS := strings.Split(lines[i], "WILLPOWER: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.willpower = rating
		default:
		}
		//Import LOGIC Attribute
		switch strings.Contains(lines[i], "LOGIC: ") {
		case true:
			ratingS := strings.Split(lines[i], "LOGIC: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.logic = rating
		default:
		}
		//Import INTUITION Attribute
		switch strings.Contains(lines[i], "INTUITION: ") {
		case true:
			ratingS := strings.Split(lines[i], "INTUITION: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.intuition = rating
		default:
		}
		//Import CHARISMA Attribute
		switch strings.Contains(lines[i], "CHARISMA: ") {
		case true:
			ratingS := strings.Split(lines[i], "CHARISMA: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.charisma = rating
		default:
		}
		//Import EDGE Attribute
		switch strings.Contains(lines[i], "EDGE: ") {
		case true:
			ratingS := strings.Split(lines[i], "EDGE: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.edge = rating
			p.maxEdge = rating
		default:
			//p.EDGE = 0
		}
		//Import RESONANCE Attribute
		switch strings.Contains(lines[i], "RESONANCE: ") {
		case true:
			ratingS := strings.Split(lines[i], "RESONANCE: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.RESONANCE = 0
		}
		//Import DEPTH Attribute
		switch strings.Contains(lines[i], "DEPTH: ") {
		case true:
			ratingS := strings.Split(lines[i], "DEPTH: ")
			rating, _ := strconv.Atoi(strings.Trim(ratingS[1], SPACES))
			p.softwareSkill = rating
		default:
			//p.DEPTH = 0
		}
		/*	//Import Specialization
			switch strings.Contains(lines[i], "Specialisations: ") {
			case true:
				specListS := strings.Split(lines[i], "Specialisations: ")
				specList := specListS[1]
				specList = strings.Trim(specList, SPACES)
				specs := strings.Split(specList, ";")
				p.specialization = specs
			default:
			}
		*/
	}

	p.grid = gridList[0].(*TGrid) //временно - должен стартовать из публичной сети
	//p.maxMatrixCM = p.device.GetMatrixCM()
	p.matrixCM = p.maxMatrixCM
	//p.id = id
	p.simSence = "HOT-SIM"
	p.maxStunCM = (p.GetWillpower()+1)/2 + 8
	p.stunCM = p.maxStunCM
	p.maxPhysCM = (p.body+1)/2 + 8
	p.physCM = p.maxPhysCM
	//p.SetID()
	p.host = Matrix
	p.markSet.MarksFrom = make(map[int]int)
	p.markSet.MarksFrom[p.id] = 4
	p.linklocked.LockedByID = make(map[int]bool)
	//p.linklocked.LockedByID[p.id] = true
	p.canSee.KnownData = make(map[int][30]string)
	//p.(IIcon)name = alias
	p.connected = true
	p.physLocation = false
	//p.freeActionsCount = 1
	//p.simpleActionsCount = 2
	ObjByNames[p.name] = &p
	//ObjByNames[p.name] = ObjByNames["Unknown"]
	//objectList = append(objectList, &p)
	//id++
	return &p, true
}

//ImportPlayerFromDB -
func ImportPlayerFromDB(alias string) (IPersona, bool) {
	if alias == "Pupa" {
		return player, true
	}
	//////////////////////////////////////////////////
	if !fileDBExists("PlayerDB.txt") {
		os.Create("PlayerDB.txt")

	}
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
	found := false
	for i := range playerData {
		if strings.Contains(playerData[i], alias) {
			playerToImport = playerData[i]
			found = true
		}

	}
	if !found {
		printLog("...Incorrect username or passcode", congo.ColorRed)
		return player, false
	}
	lines := strings.Split(playerToImport, "\r\n")
	charMAP := make(map[string]string)
	for i := range lines {
		parts := strings.Split(lines[i], ":")
		if len(parts) < 2 {
			continue
		}
		parts[0] = strings.Trim(parts[0], SPACES)
		key := parts[0]
		parts[1] = strings.Join(parts[1:], ":")
		parts[1] = strings.Trim(parts[1], SPACES)
		val := parts[1]
		charMAP[key] = val
		congo.WindowsMap.ByTitle["Log"].WPrintLn(key+":"+val, congo.ColorDefault)
	}
	var p IPersona
	switch charMAP["Class"] {
	case "Decker":
		printLog("is Decker", congo.ColorDefault)
		p = NewPersona(charMAP["Alias"], charMAP["Device"])
	case "Technomancer":
		printLog("is TECH", congo.ColorDefault)
		t := NewTechnom(charMAP["Alias"], charMAP["Device"])
		p = t

	default:
		return player, false
	}
	body, _ := strconv.Atoi(charMAP["BODY"])
	agi, _ := strconv.Atoi(charMAP["AGILITY"])
	rea, _ := strconv.Atoi(charMAP["REACTION"])
	str, _ := strconv.Atoi(charMAP["STRENGHT"])
	will, _ := strconv.Atoi(charMAP["WILLPOWER"])
	log, _ := strconv.Atoi(charMAP["LOGIC"])
	intu, _ := strconv.Atoi(charMAP["INTUITION"])
	cha, _ := strconv.Atoi(charMAP["CHARISMA"])
	edge, _ := strconv.Atoi(charMAP["EDGE"])
	computSkill, _ := strconv.Atoi(charMAP["Computer"])
	cyberSkill, _ := strconv.Atoi(charMAP["Cybercombat"])
	electronicSkill, _ := strconv.Atoi(charMAP["Electronic Warfare"])
	hackSkill, _ := strconv.Atoi(charMAP["Hacking"])
	hardwareSkill, _ := strconv.Atoi(charMAP["Hardware"])
	softwareSkill, _ := strconv.Atoi(charMAP["Software"])
	p.SetAttribute("B", body)
	p.SetAttribute("A", agi)
	p.SetAttribute("R", rea)
	p.SetAttribute("S", str)
	p.SetAttribute("W", will)
	p.SetAttribute("L", log)
	p.SetAttribute("I", intu)
	p.SetAttribute("C", cha)
	p.SetAttribute("E", edge)
	p.SetMaxEdge(edge)
	p.SetSkill("Computer", computSkill)
	p.SetSkill("Cybercombat", cyberSkill)
	p.SetSkill("Electronic", electronicSkill)
	p.SetSkill("Hacking", hackSkill)
	p.SetSkill("Hardware", hardwareSkill)
	p.SetSkill("Software", softwareSkill)
	/////////////SPECS
	compSpecSTR := charMAP["ComputerSpec"]
	compSpec := strings.Split(compSpecSTR, ",")
	for i := range compSpec {
		compSpec[i] = strings.Trim(compSpec[i], SPACES)
		p.AddSpecialization("CompSpec", compSpec[i])
	}
	cyberSpecSTR := charMAP["CybercombatSpec"]
	cyberSpec := strings.Split(cyberSpecSTR, ",")
	for i := range cyberSpec {
		cyberSpec[i] = strings.Trim(cyberSpec[i], SPACES)
		p.AddSpecialization("CyberSpec", cyberSpec[i])
	}
	electronicSpecSTR := charMAP["Electronic WarfareSpec"]
	electronicSpec := strings.Split(electronicSpecSTR, ",")
	for i := range electronicSpec {
		electronicSpec[i] = strings.Trim(electronicSpec[i], SPACES)
		p.AddSpecialization("electronicSpec", electronicSpec[i])
	}
	hackingSpecSTR := charMAP["HackingSpec"]
	hackingSpec := strings.Split(hackingSpecSTR, ",")
	for i := range hackingSpec {
		hackingSpec[i] = strings.Trim(hackingSpec[i], SPACES)
		p.AddSpecialization("HackingSpec", hackingSpec[i])
	}
	hardwareSpecSTR := charMAP["HardwareSpec"]
	hardwareSpec := strings.Split(hardwareSpecSTR, ",")
	for i := range hardwareSpec {
		hardwareSpec[i] = strings.Trim(hardwareSpec[i], SPACES)
		p.AddSpecialization("HardwareSpec", hardwareSpec[i])
	}
	softwareSpecSTR := charMAP["SoftwareSpec"]
	softwareSpec := strings.Split(softwareSpecSTR, ",")
	for i := range softwareSpec {
		softwareSpec[i] = strings.Trim(softwareSpec[i], SPACES)
		p.AddSpecialization("SoftwareSpec", softwareSpec[i])
	}

	if lp, ok := p.(ITechnom); ok {
		res, _ := strconv.Atoi(charMAP["RESONANCE"])
		lp.SetResonance(res)
		lp.SetMatrixCM(lp.GetStunCM())
		compile, _ := strconv.Atoi(charMAP["Compiling"])
		decompile, _ := strconv.Atoi(charMAP["Decompiling"])
		register, _ := strconv.Atoi(charMAP["Registering"])
		lp.SetSkill("Compiling", compile)
		lp.SetSkill("Decompiling", decompile)
		lp.SetSkill("Registering", register)

		compilingSpecSTR := charMAP["CompilingSpec"]
		compilingSpec := strings.Split(compilingSpecSTR, ",")
		for i := range compilingSpec {
			compilingSpec[i] = strings.Trim(compilingSpec[i], SPACES)
			p.AddSpecialization("CompilingSpec", compilingSpec[i])
		}
		decompilingSpecSTR := charMAP["DecompilingSpec"]
		decompilingSpec := strings.Split(decompilingSpecSTR, ",")
		for i := range decompilingSpec {
			decompilingSpec[i] = strings.Trim(decompilingSpec[i], SPACES)
			p.AddSpecialization("DecompilingSpec", decompilingSpec[i])
		}
		registeringSpecSTR := charMAP["RegisteringSpec"]
		registeringSpec := strings.Split(registeringSpecSTR, ",")
		for i := range registeringSpec {
			registeringSpec[i] = strings.Trim(registeringSpec[i], SPACES)
			p.AddSpecialization("RegisteringSpec", registeringSpec[i])
		}

	}

	ObjByNames[p.GetName()] = p

	return p, true
}
