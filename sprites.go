package main

import (
	"strconv"
)

//SPNumber - index of Sprite Power
var SPNumber int

//TSprite -
type TSprite struct {
	TPersona
	actionProtocol string
	level          int
	resonance      int
}

//IAgent -
type ISprite interface {
	IPersona
	ISpriteOnly
}

//IAgentOnly -
type ISpriteOnly interface {
	GetSprLevel() int
	ActivatePower(s string) bool
}

var _ ISprite = (*TSprite)(nil)

//NewTechnom -
func (t *TTechnom) NewSprite(spriteType string, level int) *TSprite {
	s := TSprite{}
	id = id + xd6Test(3)
	s.id = id
	s.level = level
	s.device = addDevice(spriteType)
	s.GetDevice().AddProgramtoDevice("", 0)
	s.uDevice = spriteType
	if s.GetDevice().model == spriteType {
		s.uDevice = spriteType
	} else {
		s.uDevice = s.GetDevice().model
	}
	s.deviceRating = level
	s.uType = "Sprite"
	s.isPlayer = false
	s.name = s.GetType() + " " + strconv.Itoa(s.id)
	s.alias = "Sprite ALIAS"
	s.faction = t.GetFaction()
	s.grid = t.GetGrid()
	s.willpower = level
	s.logic = level
	s.intuition = level
	s.charisma = level
	switch spriteType {
	case "Courier Sprite":
		s.SetAttackRaw(level)
		s.SetSleazeRaw(level + 3)
		s.SetDataProcessingRaw(level + 1)
		s.SetFirewallRaw(level + 2)
		s.SetDeviceRating(level)
		s.SetMaxMatrixCM((s.level+1)/2 + 8)
		s.SetMatrixCM(s.maxMatrixCM)
		s.computerSkill = level
		s.hackingSkill = level
	case "Crack Sprite":
		s.SetAttackRaw(level)
		s.SetSleazeRaw(level + 3)
		s.SetDataProcessingRaw(level + 2)
		s.SetFirewallRaw(level + 1)
		s.SetDeviceRating(level)
		s.SetMaxMatrixCM((s.level+1)/2 + 8)
		s.SetMatrixCM(s.maxMatrixCM)
		s.computerSkill = level
		s.electronicSkill = level
		s.hackingSkill = level
	case "Data Sprite":
		s.SetAttackRaw(level - 1)
		if s.GetAttackRaw() < 1 {
			s.SetAttackRaw(1)
		}
		s.SetSleazeRaw(level)
		s.SetDataProcessingRaw(level + 4)
		s.SetFirewallRaw(level + 1)
		s.SetDeviceRating(level)
		s.SetMaxMatrixCM((s.level+1)/2 + 8)
		s.SetMatrixCM(s.maxMatrixCM)
		//s.SetMatrixCM((level+1)/2 + 8)
		s.computerSkill = level
		s.electronicSkill = level
	case "Fault Sprite":
		s.SetAttackRaw(level + 3)
		s.SetSleazeRaw(level)
		s.SetDataProcessingRaw(level + 1)
		s.SetFirewallRaw(level + 2)
		s.SetDeviceRating(level)
		s.computerSkill = level
		s.cybercombatSkill = level
		s.hackingSkill = level
	case "Machine Sprite":
		s.SetAttackRaw(level + 1)
		s.SetSleazeRaw(level)
		s.SetDataProcessingRaw(level + 3)
		s.SetFirewallRaw(level + 2)
		s.SetDeviceRating(level)
		s.SetMaxMatrixCM((s.level+1)/2 + 8)
		s.SetMatrixCM(s.maxMatrixCM)
		s.computerSkill = level
		s.electronicSkill = level
		s.hardwareSkill = level
	default:
		panic("Error: Sprite unknown type")
	}
	s.SetMaxMatrixCM((s.level+1)/2 + 8)
	s.SetMatrixCM(s.GetMaxMatrixCM())
	//s.SetMatrixCM((level+1)/2 + 8)
	//SKILLS
	/*a.cybercombatSkill = agent.programRating
	a.computerSkill = agent.programRating
	a.hackingSkill = agent.programRating
	a.softwareSkill = -1
	a.electronicSkill = -1
	a.hardwareSkill = -1*/
	//MATRIX ATTRIBUTES
	/*	a.body = agent.programRating
		a.reaction = agent.programRating
		a.willpower = agent.programRating
		a.logic = agent.programRating
		a.intuition = agent.programRating
		a.charisma = agent.programRating
		a.edge = 0
		a.maxEdge = 0*/
	//a.id = id
	s.silentMode = t.GetSilentRunningMode()
	s.simSence = "HOT-SIM"
	s.maxStunCM = 999999
	s.stunCM = 999999
	s.maxPhysCM = 999999
	s.physCM = 999999
	s.actionProtocol = "Follow"
	//s.SetMaxMatrixCM((s.level+1)/2 + 8)
	//s.SetMatrixCM(s.maxMatrixCM)
	//s.matrixCM = s.maxMatrixCM
	s.host = t.GetHost()
	s.markSet.MarksFrom = make(map[int]int)
	s.markSet.MarksFrom[s.id] = 4
	s.linklocked.LockedByID = make(map[int]bool)
	s.canSee.KnownData = make(map[int][30]string)
	s.owner = t
	s.connected = true
	s.physLocation = false
	s.freeActionsCount = 1
	s.simpleActionsCount = 2
	t.ChangeFOWParametr(s.id, 0, "Spotted")

	data := t.GetFieldOfView().KnownData[s.id]
	data[0] = "Spotted"
	data[1] = "Unknown"
	data[2] = "MCM"
	data[3] = "Unknown"
	data[4] = "Unknown"
	data[5] = "Unknown"
	data[6] = s.GetSimSence()
	data[7] = strconv.Itoa(s.GetAttack())
	data[8] = strconv.Itoa(s.GetSleaze())
	data[9] = strconv.Itoa(s.GetDataProcessing())
	data[10] = strconv.Itoa(s.GetFirewall())
	data[11] = s.GetUDevice()
	data[12] = "Unknown"
	data[13] = s.GetGridName()
	data[14] = "Unknown"
	data[15] = "Unknown"
	data[16] = "Unknown"
	data[17] = "Unknown"
	data[18] = s.GetOwner().GetName()
	t.GetFieldOfView().KnownData[s.id] = data
	ObjByNames[s.name] = &s
	//printLog(fmt.Sprintf("INFO: %v", a)+" --is an Agent", congo.ColorDefault)
	//printLog(fmt.Sprintf("INFO: %v", a.GetType())+" --is an Agent's type", congo.ColorDefault)
	s.initiative = 1
	id++
	return &s
}

//RollInitiative -
func (s *TSprite) GetSprLevel() int {
	return s.level
}

//RollInitiative -
func (s *TSprite) RollInitiative() {
	s.SetInitiative(s.GetDataProcessing() + 4 + xd6Test(4)) //taken from Hero Lab
}

//GetType -
func (s *TSprite) GetType() string {
	return "Sprite"
}

//GetActionProtocol -
func (s *TSprite) GetActionProtocol() string {
	return s.actionProtocol
}

//SetActionProtocol -
func (s *TSprite) SetActionProtocol(newProtocol string) {
	s.actionProtocol = newProtocol
}

//GetRating -
func (s *TSprite) GetRating() int {
	return s.level
}

func (s *TSprite) RunActionProtocol() (string, string) {
	if owner, ok := s.GetOwner().(IPersona); ok {
		switch s.actionProtocol {
		default:
			return "", ""
			//////////////////////////////////////
		case "Idle":
			return "WAIT", ""
			//////////////////////////////////////
		case "Follow":
			if owner.GetGrid() != s.GetGrid() { //if Agent is not on the same Grid as Owner
				return "GRID_HOP", owner.GetGrid().GetGridName() //            follow to the same Grid
			}
			host := owner.GetHost()
			if host != s.GetHost() { //if Agent is not on the same Host as Owner
				if host.GetName() == "Matrix" {
					return "EXIT_HOST", host.GetName()
				}
				if !checkExistingMarks(s.id, host.GetID(), 1) { //if not have MARKS on Host - get them
					if s.GetSleaze() < s.GetAttack() {
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

	return "WAIT", s.GetName()
}

//
func (s *TSprite) SetMaxMatrixCM(mxMCM int) {
	s.GetDevice().maxMatrixCM = mxMCM
}

type SpritePower struct {
	powerName        string
	activeUser       ISprite
	activeEnviroment IHost
	activeTarget     IIcon
}

//type AffectedTarget interface {
//	Get
//}

func (s *TSprite) Supression() bool {
	if s.uDevice != "Crack Sprite" {
		return false
	}
	SPNumber++
	SPowerMap[SPNumber] = SpritePower{
		"Supression",
		s,
		s.GetHost(),
		nil,
	}
	return true
}

func checkActivePower(icon IIcon, powerName string) bool {
	for _, val := range SPowerMap {
		if val.powerName == powerName {
			if val.activeEnviroment == icon.GetHost() {
				return true
			}
			if val.activeTarget == icon {
				return true
			}
		}
	}
	/*_, ok := SPowerMap[id]
	if ok {
		return true
	}*/
	return false
}

func (s *TSprite) ActivatePower(powerName string) bool {
	var validPowers []string
	switch s.uDevice {
	case "Crack Sprite":
		validPowers = append(validPowers, "Supression")
	default:
		return false
	}
	SPNumber++
	SPowerMap[SPNumber] = SpritePower{
		powerName,
		s,
		s.GetHost(),
		nil,
	}
	return true
}
