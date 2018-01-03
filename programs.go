package main

import "fmt"
import "github.com/Galdoba/ConGo/congo"
import "strconv"

//TcyberProgram -
type TcyberProgram struct {
	programName   string
	programType   string
	programStatus string
	programRating int
}

func preapareCyberProgram(prgName string, rating int) *TcyberProgram {
	program := new(TcyberProgram)
	program.programName = prgName
	//program2 := make(map[string]ICyberProgramData)
	//program2[prgName].SetRating(rating)
	//program2[prgName].SetStatus("inStorage")
	//program2[prgName].SetRating(rating)
	program.programRating = rating
	program.programStatus = "Stored"
	program.programType = "--PLACEHOLDER--"
	return program
}

//TAgent -
type TAgent struct {
	TPersona
}

//IAgent -
type IAgent interface {
	IPersona
	IAgentOnly
}

//IAgentOnly -
type IAgentOnly interface {
}

//NewTechnom -
func (d *TDevice) NewAgent() IAgent {
	programs := d.GetPrograms()

	ownerPersona := d.GetOwner()

	if ownerPersona == nil { //not sure how to handle "owner==nil" problem
		ownerPersona = SourceIcon
	}
	printLog(fmt.Sprintf("INFO: %v", ownerPersona)+" --is the owner for an Agent", congo.ColorYellow)
	//printLog(ownerPersona.GetName(), congo.ColorDefault)
	var agent *TcyberProgram
	for i := range programs {
		if programs[i].programName != "Agent" {
			continue
		}
		agent = programs[i]
	}
	a := TAgent{}
	a.SetID()
	a.device = d
	a.uDevice = d.model
	a.uType = "Agent"
	a.isPlayer = false
	a.name = a.GetType() + " " + strconv.Itoa(a.id)
	a.alias = "Agent ALIAS"
	a.faction = ownerPersona.(IIcon).GetFaction()
	a.grid = ownerPersona.(IIcon).GetGrid()
	//a.grid = player.GetGrid()
	a.cybercombatSkill = agent.programRating
	a.computerSkill = agent.programRating
	a.hackingSkill = agent.programRating
	a.softwareSkill = -1
	a.electronicSkill = -1
	a.hardwareSkill = -1
	a.body = agent.programRating
	a.reaction = agent.programRating
	a.willpower = agent.programRating
	a.logic = agent.programRating
	a.intuition = agent.programRating
	a.charisma = agent.programRating
	a.edge = 0
	a.maxEdge = 0
	//a.id = id
	a.silentMode = false
	a.simSence = "HOT-SIM"
	a.maxStunCM = 10
	a.stunCM = 10
	a.maxPhysCM = 10
	a.physCM = 10

	a.maxMatrixCM = a.GetDeviceRating()/2 + 8
	a.matrixCM = 10
	a.host = Matrix
	a.markSet.MarksFrom = make(map[int]int)
	a.markSet.MarksFrom[a.id] = 4
	a.linklocked.LockedByID = make(map[int]bool)
	a.canSee.KnownData = make(map[int][30]string)
	a.owner = ownerPersona
	a.connected = true
	a.physLocation = false
	a.freeActionsCount = 1
	a.simpleActionsCount = 2
	//ownerPersona.ChangeFOWParametr(a.id, 0, "Spotted")
	data := player.GetFieldOfView().KnownData[a.id]
	data[0] = "Spotted"
	data[1] = "Unknown"
	data[2] = "Unknown"
	/*data[3] = "Unknown"
	data[4] = "Unknown"
	data[5] = "Unknown"
	data[6] = "Unknown"
	data[7] = "Unknown"
	data[8] = "Unknown"
	data[9] = "Unknown"
	data[10] = "Unknown"
	data[11] = "Unknown"
	data[12] = "Unknown"
	data[13] = "Unknown"
	data[14] = "Unknown"
	data[15] = "Unknown"
	data[16] = "Unknown"
	data[17] = "Unknown"
	data[18] = "Unknown"*/
	player.GetFieldOfView().KnownData[a.id] = data
	player.ChangeFOWParametr(a.id, 0, "Spotted")

	ObjByNames[a.name] = &a
	printLog(fmt.Sprintf("INFO: %v", a)+" --is an Agent", congo.ColorDefault)
	printLog(fmt.Sprintf("INFO: %v", a.GetType())+" --is an Agent's type", congo.ColorDefault)
	id++
	return &a
}

//RollInitiative -
func (a *TAgent) RollInitiative() {
	a.SetInitiative(a.GetDataProcessing() + 4 + xd6Test(4)) //taken from Hero Lab
}

//GetType -
func (a *TAgent) GetType() string {
	return "Agent"
}
