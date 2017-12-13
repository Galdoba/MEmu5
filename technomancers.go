package main

import (
	"strconv"

	"github.com/Galdoba/ConGo/congo"
)

//TTechnom -
type TTechnom struct {
	TPersona
	resonance        int
	submersion       int
	compilingSkill   int
	decompilingSkill int
	registeringSkill int
}

//ITechnom -
type ITechnom interface {
	IPersona
	ITechnomOnly
}

//ITechnomOnly -
type ITechnomOnly interface {
	GetResonance() int
	SetResonance(int)
	GetSubmersion() int
	GetSustainedForms() []ComplexForm
	ResistFade(int, string)
}

//NewTechnom -
func NewTechnom(alias string, d string) ITechnom {
	t := TTechnom{}
	t.isPlayer = true
	t.name = alias
	t.faction = alias
	t.alias = alias
	t.device = addDevice(d)

	//t.SetDeviceAttackRaw(t.GetCharisma())
	//r := rand.Intn(len(gridList))
	t.grid = gridList[0].(*TGrid) //временно - должен стартовать из публичной сети
	t.cybercombatSkill = -1
	t.computerSkill = -1
	t.hackingSkill = -1
	t.softwareSkill = -1
	t.electronicSkill = -1
	t.hardwareSkill = -1
	t.compilingSkill = -1
	t.decompilingSkill = -1
	t.registeringSkill = -1
	t.body = 1
	t.reaction = 1
	t.willpower = 1
	t.logic = 1
	t.intuition = 1
	t.charisma = 1
	t.edge = 0
	t.maxEdge = 0
	t.id = id
	if t.GetDevice().model == "Living Persona" {
		t.uDevice = "Living Persona"
	} else {
		t.uDevice = t.GetDevice().model
	}
	//t.silentMode = false
	t.simSence = "HOT-SIM"
	t.maxStunCM = (t.willpower+1)/2 + 8
	t.stunCM = t.maxStunCM
	t.maxPhysCM = (t.body+1)/2 + 8
	t.physCM = t.maxPhysCM
	t.maxMatrixCM = t.maxStunCM
	t.matrixCM = t.maxStunCM
	t.SetID()
	t.host = Matrix
	t.markSet.MarksFrom = make(map[int]int)
	t.markSet.MarksFrom[t.id] = 4
	t.linklocked.LockedByID = make(map[int]bool)
	t.canSee.KnownData = make(map[int][30]string)
	t.connected = true
	t.physLocation = false
	t.freeActionsCount = 1
	t.simpleActionsCount = 2
	id++
	return &t
}

//GetPing -
func (t *TTechnom) GetPing() string {
	return "Pong"
}

//GetDeviceRating -
func (t *TTechnom) GetDeviceRating() int {
	if t.device.model == "Living Persona" {
		return t.resonance
	}
	return t.device.deviceRating
}

//GetAttack -
/*func (t *TTechnom) GetAttack() int {
	//	if t.device.model == "Living Persona" {
	//	return t.charisma + t.device.attackMod
	///
	return t.device.attack + t.device.attackMod
}*/

//GetSleaze -
func (t *TTechnom) GetSleaze() int {
	/*if t.device.model == "Living Persona" {
		return t.intuition + t.device.sleazeMod
	}*/
	return t.device.sleaze + t.device.sleazeMod
}

//GetDataProcessing -
func (t *TTechnom) GetDataProcessing() int {
	/*if t.device.model == "Living Persona" {
		return t.logic + t.device.dataProcessingMod
	}*/
	return t.device.dataProcessing + t.device.dataProcessingMod
}

//GetFirewall -
func (t *TTechnom) GetFirewall() int {
	/*if t.device.model == "Living Persona" {
		return t.willpower + t.device.firewallMod
	}*/
	return t.device.firewall + t.device.firewallMod
}

//GetResonance -
func (t *TTechnom) GetResonance() int {
	return t.resonance
}

//SetResonance -
func (t *TTechnom) SetResonance(newRes int) {
	t.resonance = newRes
}

//GetSubmersion -
func (t *TTechnom) GetSubmersion() int {
	return t.submersion
}

//SetName -
func (t *TTechnom) SetName(name string) {
	t.name = name
}

//GetMatrixCM -
func (t *TTechnom) GetMatrixCM() int {
	return t.stunCM
}

//ReceiveMatrixDamage -
func (t *TTechnom) ReceiveMatrixDamage(damage int) {
	if t.CheckRunningProgram("Virtual Machine") && damage > 0 {
		damage++
		if t.GetFaction() == player.GetFaction() {
			printLog("...WARNING! 1 additional Matrix Damage caused by Virtual Machine program", congo.ColorYellow)
			hold()
		}
	}
	t.SetStunCM(t.GetStunCM() - damage)
	if t.GetFaction() == player.GetFaction() {
		printLog("..."+t.GetName()+" takes "+strconv.Itoa(damage)+" Matrix damage", congo.ColorYellow)
		printLog("...Matrix damage converted to Stun", congo.ColorYellow)
		hold()
	}
	if t.GetMatrixCM() < 1 {
		t.Dumpshock()
	}
}

//SetSkill -
func (t *TTechnom) SetSkill(name string, rating int) {
	if rating == 0 {
		rating = -1
	}
	switch name {
	default:
	case "Cybercombat":
		t.cybercombatSkill = rating
	case "Electronic":
		t.electronicSkill = rating
	case "Hacking":
		t.hackingSkill = rating
	case "Computer":
		t.computerSkill = rating
	case "Hardware":
		t.hardwareSkill = rating
	case "Software":
		t.softwareSkill = rating
	case "Compiling":
		t.compilingSkill = rating
	case "Decompiling":
		t.decompilingSkill = rating
	case "Registering":
		t.registeringSkill = rating
	}
}

//ResistMatrixDamage -
func (t *TTechnom) ResistMatrixDamage(damage int) int {
	resistDicePool := 0
	resistDicePool = resistDicePool + t.GetDeviceRating()
	resistDicePool = resistDicePool + t.GetFirewall()
	printLog("...Incoming matrix damage detected", congo.ColorGreen)

	if t.CheckRunningProgram("Shell") {
		resistDicePool = resistDicePool + 1
		if t.GetFaction() == player.GetFaction() {
			printLog("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if t.CheckRunningProgram("Armor") {
		resistDicePool = resistDicePool + 2
		if t.GetFaction() == player.GetFaction() {
			printLog("Program Armor: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	printLog("...Evaluated Firewall resources: "+strconv.Itoa(resistDicePool)+" mp/p", congo.ColorGreen)
	damageSoak, gl, cgl := simpleTest(t.GetID(), resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if t.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			printLog(t.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if t.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			printLog(t.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if t.GetFaction() == player.GetFaction() {
		printLog("..."+t.GetName()+": "+strconv.Itoa(damageSoak)+" Matrix damage soaked", congo.ColorGreen)
		hold()
	}
	return realDamage
}

//GetDevice -
func (t *TTechnom) GetDevice() *TDevice {
	return t.device
}

//GetDevice -
func (t *TTechnom) GetSustainedForms() []ComplexForm {
	var formsList []ComplexForm
	for i := range CFDBMap {
		if getComplexForm(i).madeByID == t.id {
			formsList = append(formsList, getComplexForm(i))
		}
	}
	return formsList
}

func (t *TTechnom) ResistFade(fade int, fadeType string) {
	printLog("...Fading: "+strconv.Itoa(fade)+" "+fadeType, congo.ColorGreen)
	resDP := t.GetWillpower() + t.GetResonance()
	printLog("...Fade resiting: "+strconv.Itoa(resDP)+" dice", congo.ColorGreen)
	resistedFadeV, gl, cgl := simpleTest(t.id, resDP, 1000, 0)
	if gl == true {
		fade = fade + (xd6Test(1) / 2)
	}
	if cgl == true {
		fade = fade + (xd6Test(1) / 2)
		fadeType = "phys"
	}
	fade = fade - resistedFadeV
	if fade < 0 {
		fade = 0
	}
	if fadeType == "stun" {
		t.ReceiveStunBiofeedbackDamage(fade)
	}
	if fadeType == "phys" {
		t.ReceivePhysBiofeedbackDamage(fade)
	}
}
