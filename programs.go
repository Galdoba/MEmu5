package main

//import "fmt"
//import "github.com/Galdoba/ConGo/congo"
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
	program.programRating = rating
	program.programStatus = "Stored"
	program.programType = "--PLACEHOLDER--" // Common/Hacking/Rigger
	return program
}

//TAgent -
type TAgent struct {
	TPersona
	actionProtocol string
	rating         int
}

//IAgent -
type IAgent interface {
	IPersona
	IAgentOnly
}

//IAgentOnly -
type IAgentOnly interface {
	GetActionProtocol() string
	SetActionProtocol(string)
	GetRating() int
	RunActionProtocol() (string, string)
}

//NewTechnom -
func (d *TDevice) NewAgent() IAgent {
	programs := d.GetPrograms()

	ownerPersona := d.GetOwner()

	if ownerPersona == nil { //not sure how to handle "owner==nil" problem
		ownerPersona = SourceIcon
	}
	//printLog(fmt.Sprintf("INFO: %v", ownerPersona)+" --is the owner for an Agent", congo.ColorYellow)
	//printLog(ownerPersona.GetName(), congo.ColorDefault)
	var agent *TcyberProgram
	for i := range programs {
		if programs[i].programName != "Agent" {
			continue
		}
		agent = programs[i]
	}
	a := TAgent{}
	id = id + xd6Test(3)
	a.id = id
	a.rating = agent.programRating
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
	a.silentMode = ownerPersona.(IPersona).GetSilentRunningMode()
	a.simSence = "HOT-SIM"
	a.maxStunCM = 999999
	a.stunCM = 999999
	a.maxPhysCM = 999999
	a.physCM = 999999
	a.actionProtocol = "Overwatch"
	a.maxMatrixCM = a.GetDeviceRating()/2 + 8
	a.matrixCM = 10
	a.host = ownerPersona.(IIcon).GetHost()
	a.markSet.MarksFrom = make(map[int]int)
	a.markSet.MarksFrom[a.id] = 4
	a.linklocked.LockedByID = make(map[int]bool)
	a.canSee.KnownData = make(map[int][30]string)
	a.owner = ownerPersona
	a.connected = true
	a.physLocation = false
	a.freeActionsCount = 1
	a.simpleActionsCount = 2
	ownerPersona.ChangeFOWParametr(a.id, 0, "Spotted")

	data := ownerPersona.GetFieldOfView().KnownData[a.id]
	data[0] = "Spotted"
	data[1] = "Unknown"
	data[2] = "MCM"
	data[3] = "Unknown"
	data[4] = "Unknown"
	data[5] = "Unknown"
	data[6] = a.GetSimSence()
	data[7] = strconv.Itoa(a.GetAttack())
	data[8] = strconv.Itoa(a.GetSleaze())
	data[9] = strconv.Itoa(a.GetDataProcessing())
	data[10] = strconv.Itoa(a.GetFirewall())
	data[11] = "Unknown"
	data[12] = "Unknown"
	data[13] = a.GetGridName()
	data[14] = "Unknown"
	data[15] = "Unknown"
	data[16] = "Unknown"
	data[17] = "Unknown"
	data[18] = a.GetOwner().GetName()
	ownerPersona.GetFieldOfView().KnownData[a.id] = data
	ObjByNames[a.name] = &a
	//printLog(fmt.Sprintf("INFO: %v", a)+" --is an Agent", congo.ColorDefault)
	//printLog(fmt.Sprintf("INFO: %v", a.GetType())+" --is an Agent's type", congo.ColorDefault)
	a.initiative = 1
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

//GetActionProtocol -
func (a *TAgent) GetActionProtocol() string {
	return a.actionProtocol
}

//SetActionProtocol -
func (a *TAgent) SetActionProtocol(newProtocol string) {
	a.actionProtocol = newProtocol
}

//GetRating -
func (a *TAgent) GetRating() int {
	return a.rating
}

func (a *TAgent) RunActionProtocol() (string, string) {
	if owner, ok := a.GetOwner().(IPersona); ok {
		switch a.actionProtocol {
		default:
			return "", ""
			//////////////////////////////////////
		case "Idle":
			return "WAIT", ""
			//////////////////////////////////////
		case "Follow":
			if owner.GetGrid() != a.GetGrid() { //if Agent is not on the same Grid as Owner
				return "GRID_HOP", owner.GetGrid().GetGridName() //            follow to the same Grid
			}
			host := owner.GetHost()
			if host != a.GetHost() { //if Agent is not on the same Host as Owner
				if host.GetName() == "Matrix" {
					return "EXIT_HOST", host.GetName()
				}
				if !checkExistingMarks(a.id, host.GetID(), 1) { //if not have MARKS on Host - get them
					if a.GetSleaze() < a.GetAttack() {
						return "BRUTE_FORCE", host.GetName()
					} else {
						return "HACK_ON_THE_FLY", host.GetName()
					}
				}
				return "ENTER_HOST", host.GetName() //            Enter host where owner is
			}
			////////////////////////////////////////
		case "Overwatch":
			return "SCAN_ENVIROMENT", "ALL" //do 'matrix perception>all' repeteadly
		}
	}

	return "WAIT", a.GetName()
}
