package main

import (
	"math/rand"
	"strconv"

	"github.com/Galdoba/ConGo/congo"
)

//"github.com/Galdoba/ConGo/congo"

//Abstract Object/////////////////////////////////////////////////////
//Сюда будет входить все: Иконы, Персоны, программы, хосты, айсы

//MarkSet -
type MarkSet struct {
	MarksFrom map[int]int
}

//Locked -
type Locked struct {
	LockedByID map[int]bool
}

//FieldOfView -
type FieldOfView struct {
	KnownData map[int][30]string
}

//SearchProcess -
type SearchProcess struct {
	SearchTime     []int
	SpentTurns     []int
	SearchIconType []string
	SearchIconName []string
}

//DownloadProcess -
type DownloadProcess struct {
	FileSize         []int
	DownloadedData   []int
	DownloadIconName []string
}

//TObj -
type TObj struct {
	uDevice      string
	uType        string
	name         string
	id           int
	markSet      MarkSet
	canSee       FieldOfView
	linklocked   Locked
	faction      string
	deviceRating int
}

//IObj -
type IObj interface {
	GetID() int
	SetID()
	GetType() string
	GetFaction() string
	GetName() string
	GetUDevice() string
	GetMarkSet() MarkSet
	GetFieldOfView() FieldOfView
	GetLinkLockStatus() Locked
	ChangeFOWParametr(int, int, string)
	Scanable() bool
	CountMarks() int
	ClearMarks()
	GetDeviceRating() int
	SetDeviceRating(int)
	GetPing() string
	CheckThreadedForm(string) (bool, int)
}

//CheckThreadedForm -
func (o *TObj) CheckThreadedForm(cFormName string) (bool, int) {
	for j := range CFDBMap {
		if getComplexForm(j).madeOnID != o.GetID() {
			continue
		}
		if getComplexForm(j).cfName == cFormName {
			continue
		}
		return true, getComplexForm(j).formNum
	}
	return false, 0
}

//GetPing -
func (o *TObj) GetPing() string {
	return "Ping"
}

//SetID -
func (o *TObj) SetID() {
	o.id = id
	id = id + xd6Test(3)

}

//GetLinkLockStatus -
func (o *TObj) GetLinkLockStatus() Locked {
	//panic("Abs Func Call TOBJ")
	a := 0
	b := 1
	if a > b {
		a = a - b
	} else {
		panic("Abs Func Call TOBJ")
	}
	return o.linklocked
}

//CountMarks -
func (o *TObj) CountMarks() int {
	totalMarks := 0
	for r := range o.markSet.MarksFrom {
		if o.markSet.MarksFrom[r] > 0 && r != o.id {
			totalMarks = totalMarks + o.markSet.MarksFrom[r]
		}
	}
	return totalMarks
}

//GetType -
func (o *TObj) GetType() string {
	a := 1
	b := 2
	if b > a {
		panic("Abs Func Call")
	}
	return ""
}

//GetUDevice -
func (o *TObj) GetUDevice() string {
	return o.uDevice
}

//GetFaction -
func (o *TObj) GetFaction() string {
	a := 1
	b := 2
	if b > a {
		panic("Abs Func Call")
	}
	return ""
}

//GetName -
func (o *TObj) GetName() string {
	a := 1
	b := 2
	if b > a {
		panic("Abs Func Call")
	}
	return ""
}

//GetID -
func (o *TObj) GetID() int {
	return o.id
}

//SetDeviceRating -
func (o *TObj) SetDeviceRating(newDR int) {
	o.deviceRating = newDR
}

//GetDeviceRating -
func (o *TObj) GetDeviceRating() int {
	return o.deviceRating
}

//GetMarkSet -
func (o *TObj) GetMarkSet() MarkSet {
	//panic("Abs Func Call")
	return o.markSet
}

//GetFieldOfView -
func (o *TObj) GetFieldOfView() FieldOfView {
	//panic("Abs Func Call")
	return o.canSee
}

//SetFieldOfView -
func (o *TObj) SetFieldOfView(data FieldOfView) {
	//panic("Abs Func Call")
	o.canSee = data
}

//ChangeFOWParametr -
func (o *TObj) ChangeFOWParametr(id, key int, val string) {
	//o.canSee.KnownData[id][key] = val
	//allData := o.canSee
	idData := o.canSee.KnownData[id]
	idData[key] = val
	o.canSee.KnownData[id] = idData
	//o.canSee.KnownData[id][key] = val

}

//ClearMarks -
func (o *TObj) ClearMarks() {
	for i := range o.markSet.MarksFrom {
		for _, j := range ObjByNames {
			if obj, ok := j.(*TObj); ok {
				obj.GetFaction()
			} else {
				delete(o.markSet.MarksFrom, i)
			}
		}
	}
}

//Scanable -
func (o *TObj) Scanable() bool {
	return true
}

func pickObjByID(id int) IObj {
	//printLog("len(ObjectList) = "+strconv.Itoa(len(objectList)), congo.ColorDefault)
	for _, obj := range ObjByNames {
		if obj.GetID() == id {
			return obj
		}
	}
	return nil
}

func pickIconByName(oName string) IIcon {
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetName() == oName {
				return icon
			}
		}
	}
	return nil
}

func pickHost(hName string) IHost { //ненужная функция?
	for _, obj := range ObjByNames {
		if host, ok := obj.(IHost); ok {
			if host.GetName() == hName {
				return host
			}
		}
	}
	return nil
}

///////////////////////////////////////////////////////

//TIcon - в икону входят файлы, персоны, айсы и хосты
type TIcon struct {
	TObj
	grid         *TGrid
	lastLocation *TGrid
	host         *THost
	device       *TDevice

	simSence           string
	silentMode         bool
	initiative         int
	id                 int
	owner              IObj
	isPlayer           bool
	convergenceFlag    bool
	connected          bool
	searchLen          int
	freeActionsCount   int
	simpleActionsCount int
}

//IIcon - в икону входят файлы, персоны, айсы и хосты
type IIcon interface {
	IObj
	IIconOnly
}

var _ IIcon = (*TIcon)(nil)

//IIconOnly - в икону входят файлы, персоны, айсы и хосты
type IIconOnly interface {
	GetSilentRunningMode() bool
	SetSilentRunningMode(bool)
	GetGrid() *TGrid
	//GetGridName() string
	SetGrid(*TGrid)
	GetHost() *THost
	SetHost(*THost)
	//GetID() int
	//SetID()
	GetInitiative() int
	SetInitiative(int)
	GetSimSence() string
	SetSimSence(string)
	GetOwner() IObj
	SetOwner(IObj)
	IsPlayer() bool
	LockIcon(IIcon)
	UnlockIcon(IIcon)
	ReceiveMatrixDamage(int)
	ToggleConvergence() bool
	GetConvergenceFlag() bool
	ToggleConnection()
	GetOverwatchScore() int
	SetOverwatchScore(int)
	//GetLastSureOS() int
	//SetLastSureOS(int)
	GetLongAct() int
	GetDevice() *TDevice
	CheckRunningProgram(string) bool
	ResistMatrixDamage(int) int
	RollInitiative()
	GetFreeActionsCount() int
	GetSimpleActionsCount() int
	SpendFreeAction()
	SpendSimpleAction()
	SpendComplexAction()
	ResetActionsCount()
	SetName(string)
	GetAttackMod() int
	GetSleazeMod() int
	GetDataProcessingMod() int
	GetFirewallMod() int
	GetAttackRaw() int
	GetSleazeRaw() int
	GetDataProcessingRaw() int
	GetFirewallRaw() int
	GetAttack() int
	GetSleaze() int
	GetDataProcessing() int
	GetFirewall() int
	SetAttackMod(int)
	SetSleazeMod(int)
	SetDataProcessingMod(int)
	SetFirewallMod(int)
	SetAttackRaw(int)
	SetSleazeRaw(int)
	SetDataProcessingRaw(int)
	SetFirewallRaw(int)
	GetPrograms() []*TcyberProgram
}

//GetPrograms -
func (i *TIcon) GetPrograms() []*TcyberProgram {
	/*if i.GetDevice().cyberSoftware == nil {
		panic(0)
	}

	if i.GetDevice() == nil {
		addDevice(i.GetUDevice())
		//i.device.AddProgramtoDevice("NULL", 0)
	}*/
	/*	if i.GetDevice() != nil {
		return i.GetDevice().cyberSoftware
	}*/
	return i.GetDevice().GetCyberSoftwareList()
}

//SetName -
func (i *TIcon) SetName(name string) {
	i.name = name
}

//SpendFreeAction -
func (i *TIcon) SpendFreeAction() {
	if i.freeActionsCount > 0 {
		i.freeActionsCount--
	} else {
		i.simpleActionsCount--
	}
}

//SpendSimpleAction -
func (i *TIcon) SpendSimpleAction() {
	i.simpleActionsCount--
}

//SpendComplexAction -
func (i *TIcon) SpendComplexAction() {
	i.simpleActionsCount--
	i.simpleActionsCount--
}

//RollInitiative -
func (i *TIcon) RollInitiative() {
	i.SetInitiative(i.GetDataProcessing() + i.GetIntuition() + xd6Test(4))
}

//GetIntuition -
func (i *TIcon) GetIntuition() int {
	return i.GetDeviceRating()
}

//CheckRunningProgram -
func (i *TIcon) CheckRunningProgram(name string) bool {
	programs := i.GetPrograms()
	for i := range programs {
		if programs[i].programName != name {
			continue
		}
		if programs[i].programStatus == "Running" {
			return true
		}
	}

	/*for j := range i.GetDevice().GetSoftwareList().programName {
		if i.GetDevice().software.programName[j] == name {
			if i.GetDevice().software.programStatus[j] == "Running" {
				return true
				//test.programName[0] = test.programName[0] + "__"
			}
		}
	}*/
	return false
}

//GetOverwatchScore -
func (i *TIcon) GetOverwatchScore() int {
	return i.grid.overwatchScore
}

//GetDevice -
func (i *TIcon) GetDevice() *TDevice {
	if i.device == nil {
		return addDevice(i.GetUDevice())
	}
	return i.device
}

//GetDeviceRating -
func (i *TIcon) GetDeviceRating() int {
	return i.GetDevice().deviceRating
}

//CheckThreadedForm -
func (i *TIcon) CheckThreadedForm(cFormName string) (bool, int) {
	for j := range CFDBMap {
		if getComplexForm(j).madeOnID != i.GetID() {
			continue
		}
		if getComplexForm(j).cfName != cFormName {
			continue
		}
		return true, getComplexForm(j).formNum
	}
	return false, 0
}

//GetComplexFormEffect -
func (i *TIcon) GetComplexFormEffect(cFormName string) (int, int) {
	cFormIndex := 0
	cFormEffect := 0
	for j := range CFDBMap {
		if getComplexForm(j).madeOnID != i.GetID() {
			continue
		}
		if getComplexForm(j).cfName != cFormName {
			continue
		}
		if cFormEffect < getComplexForm(j).succ {
			cFormEffect = getComplexForm(j).succ
			cFormIndex = getComplexForm(j).formNum
		}
	}

	return cFormEffect, cFormIndex
}

//GetAttack -
func (i *TIcon) GetAttack() int {
	programBoost := 0
	if i.CheckRunningProgram("Decryption") {
		programBoost = 1
	}
	cFormsBoost := 0
	infusion, _ := i.CheckThreadedForm("Infusion of Attack")
	if infusion {
		boost, _ := i.GetComplexFormEffect("Infusion of Attack")
		cFormsBoost = cFormsBoost + boost
	}
	diffusion, _ := i.CheckThreadedForm("Diffusion of Attack")
	if diffusion {
		boost, _ := i.GetComplexFormEffect("Diffusion of Attack")
		cFormsBoost = cFormsBoost - boost
	}

	att := i.GetDevice().attack + i.GetDevice().attackMod + programBoost + cFormsBoost
	if att < 0 {
		return 0
	}
	//printLog(strconv.Itoa(att), congo.ColorGreen)
	return att
}

//GetAttackMod -
func (i *TIcon) GetAttackMod() int {
	return i.GetDevice().attackMod //+ get effect of complex forms
}

//GetAttackRaw -
func (i *TIcon) GetAttackRaw() int {
	return i.GetDevice().attack
}

//SetAttack -
func (i *TIcon) SetAttack(newAttack int) {
	i.GetDevice().attack = newAttack
} //Возможно не нужен

//SetAttackMod -
func (i *TIcon) SetAttackMod(newAttack int) {
	if i.GetDevice() != nil {
		i.GetDevice().attackMod = i.GetDevice().attackMod + newAttack
	}
}

//SetAttackRaw -
func (i *TIcon) SetAttackRaw(newAttack int) {
	if i.GetDevice() != nil {
		i.GetDevice().attack = newAttack
	}
}

//GetSleaze -
func (i *TIcon) GetSleaze() int {
	programBoost := 0
	if i.CheckRunningProgram("Stealth") {
		programBoost = 1
	}
	cFormsBoost := 0
	infus, _ := i.CheckThreadedForm("Infusion of Sleaze")
	if infus {
		boost, _ := i.GetComplexFormEffect("Infusion of Sleaze")
		cFormsBoost = cFormsBoost + boost
	}
	diffus, _ := i.CheckThreadedForm("Diffusion of Sleaze")
	if diffus {
		boost, _ := i.GetComplexFormEffect("Diffusion of Sleaze")
		cFormsBoost = cFormsBoost - boost
	}

	//att := i.GetDevice().sleaze + i.GetDevice().sleazeMod + programBoost + cFormsBoost
	att := i.GetSleazeRaw() + i.GetSleazeMod() + programBoost + cFormsBoost
	if att < 0 {
		return 0
	}
	return att
}

//GetSleazeMod -
func (i *TIcon) GetSleazeMod() int {
	return i.GetDevice().sleazeMod
}

//GetSleazeRaw -
func (i *TIcon) GetSleazeRaw() int {
	return i.GetDevice().sleaze
}

//SetSleaze -
func (i *TIcon) SetSleaze(newSleaze int) {
	i.GetDevice().sleaze = newSleaze
}

//SetSleazeMod -
func (i *TIcon) SetSleazeMod(newSleaze int) {
	if i.GetDevice() != nil {
		i.GetDevice().sleaze = i.GetDevice().sleaze + newSleaze
	}
}

//SetSleazeRaw -
func (i *TIcon) SetSleazeRaw(newSleaze int) {
	if i.GetDevice() != nil {
		i.GetDevice().sleaze = newSleaze
	}
}

//GetDataProcessing -
func (i *TIcon) GetDataProcessing() int {
	programBoost := 0
	if i.CheckRunningProgram("Toolbox") {
		programBoost = 1
	}
	cFormsBoost := 0
	infus, _ := i.CheckThreadedForm("Infusion of Data Processing")
	if infus {
		boost, _ := i.GetComplexFormEffect("Infusion of Data Processing")
		cFormsBoost = cFormsBoost + boost
	}
	diffus, _ := i.CheckThreadedForm("Diffusion of Data Processing")
	if diffus {
		boost, _ := i.GetComplexFormEffect("Diffusion of Data Processing")
		cFormsBoost = cFormsBoost - boost
	}
	att := i.GetDevice().dataProcessing + i.GetDevice().dataProcessingMod + programBoost + cFormsBoost
	if att < 0 {
		return 0
	}
	return att
}

//GetDataProcessingMod -
func (i *TIcon) GetDataProcessingMod() int {
	return i.GetDevice().dataProcessingMod
}

//GetDataProcessingRaw -
func (i *TIcon) GetDataProcessingRaw() int {
	return i.GetDevice().dataProcessing
}

//SetDataProcessing -
func (i *TIcon) SetDataProcessing(newDataProcessing int) {
	i.GetDevice().dataProcessing = newDataProcessing
}

//SetDataProcessingMod -
func (i *TIcon) SetDataProcessingMod(newDataProcessing int) {
	if i.GetDevice() != nil {
		i.GetDevice().dataProcessing = i.GetDevice().dataProcessing + newDataProcessing
	}
}

//SetDataProcessingRaw -
func (i *TIcon) SetDataProcessingRaw(newDataProcessing int) {
	if i.GetDevice() != nil {
		i.GetDevice().dataProcessing = newDataProcessing
	}
}

//GetFirewall -
func (i *TIcon) GetFirewall() int {
	programBoost := 0
	if i.CheckRunningProgram("Encryption") {
		programBoost = 1
	}
	cFormsBoost := 0
	infus, _ := i.CheckThreadedForm("Infusion of Firewall")
	if infus {
		boost, _ := i.GetComplexFormEffect("Infusion of Firewall")
		cFormsBoost = cFormsBoost + boost
	}
	diffus, _ := i.CheckThreadedForm("Diffusion of Firewall")
	if diffus {
		boost, _ := i.GetComplexFormEffect("Diffusion of Firewall")
		cFormsBoost = cFormsBoost - boost
	}

	att := i.GetDevice().firewall + i.GetDevice().firewallMod + programBoost + cFormsBoost
	if att < 0 {
		return 0
	}
	return att
}

//GetFirewallMod -
func (i *TIcon) GetFirewallMod() int {
	return i.GetDevice().firewallMod
}

//GetFirewallRaw -
func (i *TIcon) GetFirewallRaw() int {
	return i.GetDevice().firewall
}

//SetFirewall -
func (i *TIcon) SetFirewall(newFirewall int) {
	i.GetDevice().firewall = newFirewall
}

//SetFirewallMod -
func (i *TIcon) SetFirewallMod(newFirewallMod int) {
	if i.GetDevice() != nil {
		i.GetDevice().firewall = i.GetDevice().firewall + newFirewallMod
	}
}

//SetFirewallRaw -
func (i *TIcon) SetFirewallRaw(newFirewall int) {
	if i.GetDevice() != nil {
		i.GetDevice().firewall = newFirewall
	}
}

//ResistMatrixDamage -
func (i *TIcon) ResistMatrixDamage(damage int) int {
	resistDicePool := 0
	resistDicePool = resistDicePool + i.GetDeviceRating()
	resistDicePool = resistDicePool + i.GetFirewall()
	if i.CheckRunningProgram("Shell") {
		resistDicePool = resistDicePool + 1
		if i.GetFaction() == player.GetFaction() {
			printLog("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if i.CheckRunningProgram("Armor") {
		resistDicePool = resistDicePool + 2
		if i.GetFaction() == player.GetFaction() {
			printLog("Program Armor: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	damageSoak, gl, cgl := simpleTest(i.GetID(), resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if i.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			printLog(i.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if i.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			printLog(i.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if i.GetFaction() == player.GetFaction() {
		printLog("..."+i.GetName()+": "+strconv.Itoa(damageSoak)+" Matrix damage soaked", congo.ColorGreen)
		hold()
	}
	return realDamage
}

//SetOverwatchScore -
func (i *TIcon) SetOverwatchScore(newScore int) {
	i.grid.overwatchScore = newScore
}

//GetLastSureOS -
func (i *TIcon) GetLastSureOS() int {
	return i.grid.lastSureOS
}

//SetLastSureOS -
func (i *TIcon) SetLastSureOS(newScore int) {
	i.grid.lastSureOS = newScore
}

//GetSimSence -
func (i *TIcon) GetSimSence() string {
	return i.simSence
}

//ToggleConnection -
func (i *TIcon) ToggleConnection() {
	if i.connected {
		i.connected = false
	} else {
		i.connected = true
	}
}

//SetSimSence -
func (i *TIcon) SetSimSence(smsence string) {
	smsence = formatString(smsence)
	smsence = cleanText(smsence)
	switch smsence {
	case "AR":
		i.simSence = "AR"
	case "COLD-SIM":
		i.simSence = "COLD-SIM"
	case "HOT-SIM":
		i.simSence = "HOT-SIM"
	default:
		i.simSence = i.GetSimSence()

	}
	i.simSence = smsence
}

//ToggleConvergence -
func (i *TIcon) ToggleConvergence() bool {
	if i.convergenceFlag == false {
		return true
	}
	return false
}

//GetConvergenceFlag -
func (i *TIcon) GetConvergenceFlag() bool {
	return i.convergenceFlag
}

//GetInitiative -
func (i *TIcon) GetInitiative() int {
	return i.initiative
}

//SetInitiative -
func (i *TIcon) SetInitiative(init int) {
	i.initiative = init
}

//GetID -
func (i *TIcon) GetID() int {
	return i.id
}

//SetID -
func (i *TIcon) SetID() {
	i.id = id
	id++
}

//GetGrid -
func (i *TIcon) GetGrid() *TGrid {
	return i.grid
}

//GetGridName -
func (i *TIcon) GetGridName() string {
	return i.grid.name
}

//SetGrid -
func (i *TIcon) SetGrid(grid *TGrid) {
	i.lastLocation = i.grid
	i.grid = grid
}

//GetHost -
func (i *TIcon) GetHost() *THost {
	return i.host
}

//SetHost -
func (i *TIcon) SetHost(host *THost) {
	i.host = host
}

//GetSilentRunningMode -
func (i *TIcon) GetSilentRunningMode() bool {
	return i.silentMode
}

//SetSilentRunningMode -
func (i *TIcon) SetSilentRunningMode(mode bool) {
	i.silentMode = mode
}

//GetType -
func (i *TIcon) GetType() string {
	return "Icon"
}

//GetFaction -
func (i *TIcon) GetFaction() string {
	return i.faction
}

//GetName -
func (i *TIcon) GetName() string {
	return i.GetType() + " " + strconv.Itoa(i.id)
}

//GetMarkSet -
func (i *TIcon) GetMarkSet() MarkSet {
	return i.markSet
}

//GetFieldOfView -
func (i *TIcon) GetFieldOfView() FieldOfView {
	//panic("Abs Func Call")
	return i.canSee
}

//GetLinkLockStatus -
func (i *TIcon) GetLinkLockStatus() Locked {
	//panic("Abs Func Call")
	return i.linklocked
}

//GetOwner -
func (i *TIcon) GetOwner() IObj {
	return i.owner
}

//SetOwner -
func (i *TIcon) SetOwner(o IObj) {
	i.owner = o
}

//IsPlayer -
func (i *TIcon) IsPlayer() bool {
	return i.isPlayer
}

//LockIcon -
func (i *TIcon) LockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[i.id] = true
}

//UnlockIcon -
func (i *TIcon) UnlockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[i.id] = false
}

//ReceiveMatrixDamage -
func (i *TIcon) ReceiveMatrixDamage(damage int) {
	printLog("...Error: "+i.GetName()+" is immune to Matrix Damage", congo.ColorYellow)
}

//GetLongAct -
func (i *TIcon) GetLongAct() int {
	return i.searchLen
}

//GetMatrixCM -
func (i *TIcon) GetMatrixCM() {
	printLog("...Error: "+i.GetName()+" is immune to Matrix Damage", congo.ColorYellow)
}

//GetFreeActionsCount -
func (i *TIcon) GetFreeActionsCount() int {
	return i.freeActionsCount
}

//GetSimpleActionsCount -
func (i *TIcon) GetSimpleActionsCount() int {
	return i.simpleActionsCount
}

//ResetActionsCount -
func (i *TIcon) ResetActionsCount() {
	i.freeActionsCount = 1
	//
	i.simpleActionsCount = 2
}

////////////////////////////////////////////////////////
//IC

//TIC -
type TIC struct {
	TIcon
	name           string
	deviceRating   int
	attack         int
	sleaze         int
	dataProcessing int
	firewall       int
	isLoaded       bool
	id             int
	markSet        MarkSet
	initiative     int
	matrixCM       int
	actionReady    int
	lastTargetName string

	//host           *THost
}

//IIC -
type IIC interface {
	IIcon
	IICOnly
}

//IICOnly -
type IICOnly interface {
	//	IIcon
	IsLoaded() bool
	SetLoadStatus(bool)
	GetMatrixCM() int
	SetMatrixCM(int)
	GetActionReady() int
	SetActionReady(int)
	GetLastTargetName() string
	SetLastTargetName(string)
	TakeFOWfromHost()

	//RollInitiative()
}

var _ IIC = (*TIC)(nil)

//NewIC -
func (h *THost) NewIC(name string) *TIC {
	id = id + xd6Test(3)
	i := TIC{}
	i.name = name
	i.uType = "IC"
	i.host = h
	i.owner = h
	i.grid = h.grid
	i.deviceRating = h.deviceRating
	i.attack = h.attack
	i.sleaze = h.sleaze
	i.dataProcessing = h.dataProcessing
	i.firewall = h.firewall
	i.id = id + xd6Test(3)
	i.isLoaded = true
	i.simSence = "HOT-SIM"
	for n := range h.icState.icName {
		if h.icState.icName[n] == i.name {
			h.icState.icStatus[n] = true
			//h.icState.icID[n] = i.id
		}
	}
	i.markSet.MarksFrom = make(map[int]int)
	//f.markSet.MarksFrom[f.id] = 4
	i.markSet.MarksFrom[i.GetID()] = 4
	i.canSee.KnownData = make(map[int][30]string)
	i.matrixCM = i.deviceRating/2 + 8
	setSeed()
	i.initiative = i.deviceRating + i.dataProcessing + xd6Test(4)
	data := player.GetFieldOfView().KnownData[i.id]
	//
	data[0] = "Spotted"
	data[2] = "Unknown"
	data[5] = "Unknown"
	data[7] = "Unknown"
	data[8] = "Unknown"
	data[9] = "Unknown"
	data[10] = "Unknown"
	data[11] = "Unknown"
	data[13] = "Unknown"
	data[18] = "Unknown"
	player.GetFieldOfView().KnownData[i.id] = data
	if i.name == "Patrol IC" {
		i.actionReady = calculatePartolScan(i.deviceRating)
	} else {
		i.actionReady = -1
	}
	i.freeActionsCount = 0
	i.simpleActionsCount = 2
	//objectList = append(objectList, &i)
	id++
	ObjByNames[i.name] = &i
	return &i
}

//RollInitiative -
func (i *TIC) RollInitiative() {
	i.SetInitiative(i.GetDataProcessing() + i.GetIntuition() + xd6Test(4))
}

//GetIntuition -
func (i *TIC) GetIntuition() int {
	return i.GetDeviceRating()
}

func calculatePartolScan(rating int) int {
	actionReady := 0
	switch rating {
	case 1:
		actionReady = 1
	case 2:
		actionReady = 1
	case 3:
		actionReady = xd6Test(1)
	case 4:
		actionReady = xd6Test(1)
	case 5:
		actionReady = xd6Test(1) + 2
	case 6:
		actionReady = xd6Test(1) + 2
	case 7:
		actionReady = xd6Test(2)
	case 8:
		actionReady = xd6Test(2)
	case 9:
		actionReady = xd6Test(2) + 2
	case 10:
		actionReady = xd6Test(2) + 2
	case 11:
		actionReady = xd6Test(3)
	case 12:
		actionReady = xd6Test(3)
	}
	return actionReady
}

//TakeFOWfromHost -
func (i *TIC) TakeFOWfromHost() {
	host := i.host
	i.canSee = host.canSee
}

//IsLoaded -
func (i *TIC) IsLoaded() bool {
	return i.isLoaded
}

//SetLoadStatus -
func (i *TIC) SetLoadStatus(newState bool) {
	i.isLoaded = newState
}

//GetType -
func (i *TIC) GetType() string {
	return i.uType
}

//GetFaction -
func (i *TIC) GetFaction() string {
	return i.faction
}

//GetName -
func (i *TIC) GetName() string {
	return i.name
}

/*/GetHost -
func (i *TIC) GetHost() THost {
	return *i.host
}*/

//GetHost -
func (i *TIC) GetHost() *THost {
	return i.host
}

//GetID -
func (i *TIC) GetID() int {
	return i.id
}

//GetInitiative -
func (i *TIC) GetInitiative() int {
	return i.initiative
}

//SetInitiative -
func (i *TIC) SetInitiative(init int) {
	i.initiative = init
}

//GetMatrixCM -
func (i *TIC) GetMatrixCM() int {
	return i.matrixCM
}

//SetMatrixCM -
func (i *TIC) SetMatrixCM(newCM int) {
	i.matrixCM = newCM
}

//GetDeviceRating -
func (i *TIC) GetDeviceRating() int {
	return i.deviceRating
}

//GetAttack -
func (i *TIC) GetAttack() int {
	return i.attack
}

//GetSleaze -
func (i *TIC) GetSleaze() int {
	return i.sleaze
}

//GetDataProcessing -
func (i *TIC) GetDataProcessing() int {
	return i.dataProcessing
}

//GetFirewall -
func (i *TIC) GetFirewall() int {
	return i.firewall
}

//GetActionReady -
func (i *TIC) GetActionReady() int {
	return i.actionReady
}

//SetActionReady -
func (i *TIC) SetActionReady(newAR int) {
	i.actionReady = newAR
}

//GetLinkLockStatus -
func (i *TIC) GetLinkLockStatus() Locked {
	//panic("Abs Func Call")
	return i.linklocked
}

//LockIcon -
func (i *TIC) LockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[i.id] = true
}

//UnlockIcon -
func (i *TIC) UnlockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[i.id] = false
}

//GetLastTargetName -
func (i *TIC) GetLastTargetName() string {
	return i.lastTargetName
}

//SetLastTargetName -
func (i *TIC) SetLastTargetName(name string) {
	i.lastTargetName = name
}

//ResistMatrixDamage -
func (i *TIC) ResistMatrixDamage(damage int) int {
	host := i.GetHost()
	resistDicePool := host.GetDeviceRating() + host.GetFirewall()
	damageSoak, gl, cgl := simpleTest(i.GetID(), resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if i.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			printLog(i.GetName()+": Firewall glitch detected!", congo.ColorYellow)
		}
	}
	if cgl {
		if i.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			printLog(i.GetName()+": Firewall critical failure!", congo.ColorRed)
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}
	if i.GetFaction() == player.GetFaction() {
		printLog("..."+i.GetName()+": "+strconv.Itoa(damageSoak)+" Matrix damage soaked", congo.ColorGreen)
	}
	return realDamage
}

//GetMarkSet -
func (i *TIC) GetMarkSet() MarkSet {
	return i.markSet
}

//GetFieldOfView -
func (i *TIC) GetFieldOfView() FieldOfView {
	//panic("Abs Func Call")
	host := i.GetHost()
	return host.canSee
}

//ReceiveMatrixDamage -
func (i *TIC) ReceiveMatrixDamage(damage int) {

	i.SetMatrixCM(i.GetMatrixCM() - damage)
}

////////////////////////////////////////////////////////

//TProgram -
type TProgram struct {
	programName   []string
	programType   []string
	programStatus []string
	programRating []int
}

//TDevice -
type TDevice struct {
	TIcon
	name               string
	deviceRating       int
	attack             int
	attackMod          int
	sleaze             int
	sleazeMod          int
	dataProcessing     int
	dataProcessingMod  int
	firewall           int
	firewallMod        int
	maxRunningPrograms int
	curRunningPrograms int
	software           *TProgram
	cyberSoftware      []*TcyberProgram
	//storedPrograms
	modifications []string
	matrixCM      int
	maxMatrixCM   int
	deviceType    string
	model         string
	owner         IObj
	id            int
	markSet       MarkSet
	initiative    int
	canSwapAtt    bool
}

//IDevice -
type IDevice interface {
	IIcon
	//GetDeviceRating() int
	//GetDataProcessing() int
	//GetFirewall() int
	GetMatrixCM() int
	SetMatrixCM(int)
	//GetAttack() int
	//GetSleaze() int
	GetModel() string
	GetSoftwareList() *TProgram
	LoadProgramToDevice(string) bool
	UnloadProgramFromDevice(string) bool
}

//NewDevice -
func NewDevice(model string, rating int) *TDevice {
	d := TDevice{}
	if _, ok := DB.DeviceDB[model]; ok {
		d = *DB.DeviceDB[model]
		//d.id = id

	} else {
		d.model = model
		d.deviceRating = rating
		d.dataProcessing = rating
		d.firewall = rating
		d.maxMatrixCM = (rating+1)/2 + 8
		d.matrixCM = d.maxMatrixCM
		d.simSence = "HOT-SIM"
		//d.grid = "Public Grid"
		//d.id = id
	}
	d.uDevice = model
	d.SetID()
	d.maxRunningPrograms = d.deviceRating
	//d.cyberSoftware = nil
	d.cyberSoftware = make([]*TcyberProgram, 0, 99)

	//d.software = make([]TProgram, 20)
	//add all soft:
	//d.software = preaparePrograms()
	//d.software.programName = append(d.software.programName, "brows")
	d.markSet.MarksFrom = make(map[int]int)
	d.markSet.MarksFrom[d.GetID()] = 4
	d.canSee.KnownData = make(map[int][30]string)
	d.linklocked.LockedByID = make(map[int]bool)
	d.name = d.GetType() + " " + strconv.Itoa(d.id)
	//id++
	return &d
}

//GetPrograms -
func (d *TDevice) GetPrograms() []*TcyberProgram {
	if d.cyberSoftware == nil {
		panic(0)
	}
	return d.cyberSoftware
}

//LoadProgramToDevice -
func (d *TDevice) LoadProgramToDevice(name string) bool {
	programs := d.GetPrograms()
	for i := range programs {
		if name != programs[i].programName {
			continue
		}
		if d.GetRunningProgramsQty() < d.GetMaxRunningPrograms() { //тест проверки на то может ли загрузиться данная программа
			programs[i].programStatus = "Running"
		} else {
			return false
		}
	}
	if name == "Agent" {
		d.NewAgent()
	}
	return true
}

//UnloadProgramFromDevice -
func (d *TDevice) UnloadProgramFromDevice(name string) bool {
	programs := d.GetPrograms()
	for i := range programs {
		if name != programs[i].programName {
			continue
		}
		if programs[i].programStatus != "Running" {
			return false
		}
		programs[i].programStatus = "Stored"
	}
	if name == "Agent" {
		printLog("вот тут мы его отключаем", congo.ColorDefault)
	}
	return true

	/*	for i := 0; i < len(d.software.programName); i++ {
			if d.software.programName[i] == name {
				if d.software.programStatus[i] == "Running" {
					d.software.programStatus[i] = "inStore"
				} else {
					return false
				}

			}
		}
		return true*/
}

//GetRunningProgramsQty -
func (d *TDevice) GetRunningProgramsQty() int {
	d.curRunningPrograms = 0
	programs := d.GetPrograms()
	for i := range programs {
		if programs[i].programStatus == "Running" {
			d.curRunningPrograms++
		}
	}
	return d.curRunningPrograms
}

//GetMaxRunningPrograms -
func (d *TDevice) GetMaxRunningPrograms() int {
	prgBoost := 0
	if d.CheckRunningProgram("Virtual Machine") {
		prgBoost = 2
	}
	d.maxRunningPrograms = d.deviceRating + prgBoost
	return d.maxRunningPrograms
}

//GetSoftwareList -
func (d *TDevice) GetSoftwareList() *TProgram {
	return d.software
}

//GetCyberSoftwareList -
func (d *TDevice) GetCyberSoftwareList() []*TcyberProgram {
	return d.cyberSoftware
}

//AddProgramtoDevice -
func (d *TDevice) AddProgramtoDevice(prgName string, prgRat int) {
	d.cyberSoftware = append(d.cyberSoftware, preapareCyberProgram(prgName, prgRat))
}

//RemoveProgramFromDevice -
func (d *TDevice) RemoveProgramFromDevice(prgName string) {
	for i := range d.cyberSoftware {
		if d.cyberSoftware[i].programName == prgName {
			d.cyberSoftware = append(d.cyberSoftware[:i], d.cyberSoftware[i+1:]...)
			break
		}
	}
}

func addDevice(model string) *TDevice {
	//objectList = append(objectList, DB.DeviceDB[model])
	return DB.DeviceDB[model]
}

//GetInitiative -
func (d *TDevice) GetInitiative() int {
	return d.initiative
}

//SetInitiative -
func (d *TDevice) SetInitiative(init int) {
	d.initiative = init
}

//SetDeviceRating -
func (d *TDevice) SetDeviceRating(init int) {
	d.initiative = init
}

//GetDeviceRating -
func (d *TDevice) GetDeviceRating() int {
	return d.deviceRating
}

//GetAttack -
func (d *TDevice) GetAttack() int {
	return d.attack + d.attackMod
}

//GetSleaze -
func (d *TDevice) GetSleaze() int {
	return d.attack + d.sleazeMod
}

//GetDataProcessing -
func (d *TDevice) GetDataProcessing() int {
	return d.dataProcessing + d.dataProcessingMod
}

//SetDataProcessing -
func (d *TDevice) SetDataProcessing(dp int) {
	d.deviceRating = dp
}

//GetFirewall -
func (d *TDevice) GetFirewall() int {
	return d.firewall + d.firewallMod
}

//SetFirewall -
func (d *TDevice) SetFirewall(fw int) {
	d.firewall = fw
}

//GetMatrixCM -
func (d *TDevice) GetMatrixCM() int {
	return d.matrixCM
}

//SetMatrixCM -
func (d *TDevice) SetMatrixCM(mcm int) {
	d.matrixCM = mcm
	if d.matrixCM > d.maxMatrixCM {
		d.matrixCM = d.maxMatrixCM
	}
	if d.matrixCM < 1 {
		d.matrixCM = 0
	}
}

//GetModel -
func (d *TDevice) GetModel() string {
	return d.model
}

//GetFaction -
func (d *TDevice) GetFaction() string {
	return d.faction
}

//GetGrid -
func (d *TDevice) GetGrid() *TGrid {
	return d.grid
}

//SetGrid -
func (d *TDevice) SetGrid(grid *TGrid) {
	d.grid = grid
}

//GetHost -
func (d *TDevice) GetHost() *THost {
	return d.host
}

//SetHost -
func (d *TDevice) SetHost(host *THost) {
	d.host = host
}

//GetMarkSet -
func (d *TDevice) GetMarkSet() MarkSet {
	return d.markSet
}

//GetFieldOfView -
func (d *TDevice) GetFieldOfView() FieldOfView {
	//panic("Abs Func Call")
	return d.canSee
}

//GetLinkLockStatus -
func (d *TDevice) GetLinkLockStatus() Locked {
	//panic("Abs Func Call")
	return d.linklocked
}

//LockIcon -
func (d *TDevice) LockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[d.id] = true
}

//UnlockIcon -
func (d *TDevice) UnlockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[d.id] = false
}

//CheckRunningProgram -
func (d *TDevice) CheckRunningProgram(name string) bool {
	return false
	programs := d.GetPrograms()
	for i := range programs {
		if programs[i].programName == name {
			if programs[i].programStatus == "Running" {
				return true
			}
		}
	}
	return false
}

//ToggleConvergence -
func (d *TDevice) ToggleConvergence() bool {
	if d.convergenceFlag == false {
		return true
	}
	return false
}

//GetConvergenceFlag -
func (d *TDevice) GetConvergenceFlag() bool {
	return d.convergenceFlag
}

//ToggleConnection -
func (d *TDevice) ToggleConnection() {
	if d.connected {
		d.connected = false
	} else {
		d.connected = true
	}
}

//ReceiveMatrixDamage -
func (d *TDevice) ReceiveMatrixDamage(damage int) {
	if d.CheckRunningProgram("Virtual Machine") && damage > 0 {
		damage++

	}
	d.SetMatrixCM(d.GetMatrixCM() - damage)
}

////////////////////////////////////////////////////

//TPersona -
type TPersona struct {
	//TObj
	TIcon
	name                  string
	alias                 string
	userMode              string
	device                *TDevice
	computerSkill         int
	hackingSkill          int
	softwareSkill         int
	electronicSkill       int
	hardwareSkill         int
	cybercombatSkill      int
	initiative            int
	body                  int
	agility               int
	reaction              int
	strenght              int
	logic                 int
	intuition             int
	willpower             int
	charisma              int
	edge                  int
	maxEdge               int
	maxPhysCM             int
	physCM                int
	maxStunCM             int
	stunCM                int
	maxMatrixCM           int
	matrixCM              int
	id                    int
	physLocation          bool
	fullDefFlag           bool
	waitFlag              bool
	markSet               MarkSet
	searchProcessStatus   SearchProcess
	downloadProcessStatus DownloadProcess
	specialization        []string
}

//IPersona -
type IPersona interface {
	IIcon
	IPersonaOnly
}

//IPersonaOnly -
type IPersonaOnly interface {
	GetMatrixCM() int
	GetHackingSkill() int
	GetCyberCombatSkill() int
	GetComputerSkill() int
	GetElectronicSkill() int
	GetSoftwareSkill() int
	SetAttribute(string, int)
	GetAttribute(string) int
	GetBody() int
	GetWillpower() int
	GetLogic() int
	GetIntuition() int
	GetCharisma() int
	GetEdge() int
	GetMaxEdge() int
	SetBody(int)
	//SetAgility(int)
	SetReaction(int)
	//SetBody(int)
	//SetWillpower(int)
	//SetLogic(int)
	//SetIntuition(int)
	//SetCharisma(int)
	SetEdge(int)
	SetMaxEdge(int)
	SetMatrixCM(int)
	GetAlias() string
	GetStunCM() int
	GetPhysCM() int
	GetMaxStunCM() int
	GetMaxPhysCM() int
	SetStunCM(int)
	SetPhysCM(int)
	Dumpshock()
	IsConnected() bool
	SetConnection(bool)
	GetPhysicalLocation() bool
	SetPhysicalLocation(bool)
	TriggerDataBomb(int)
	ReceiveBiofeedbackDamage(int)
	ReceiveStunBiofeedbackDamage(int)
	ReceivePhysBiofeedbackDamage(int)
	ResistBiofeedbackDamage(int) int
	GetSearchResultIn() int
	SetSearchResultIn(int)
	GetSearchProcess() SearchProcess
	SetSearchProcess(int, string, string)
	UpdateSearchProcess()
	GetDownloadProcess() DownloadProcess
	SetDownloadProcess(int, string)
	UpdateDownloadProcess()
	CheckConvergence()
	GetFullDeffenceFlag() bool
	SetFullDeffenceFlag(bool)
	GetWaitFlag() bool
	SetWaitFlag(bool)
	HaveValidSpec([]string) (bool, string)
	GetDeviceSoft() *TProgram
	CrashRandomProgram() bool
	isOnline() bool

	SetSkill(string, int)
	AddSpecialization(string, string)
	GetSpecializationList() []string

	/*SetDeviceAttackMod(int)
	SetDeviceSleazeMod(int)
	SetDeviceDataProcessingMod(int)
	SetDeviceFirewallMod(int)*/
}

var _ IPersona = (*TPersona)(nil)

//NewPersona -
func NewPersona(alias string, d string) IPersona {
	p := TPersona{}
	p.isPlayer = true
	p.uType = "Persona"
	p.name = alias
	p.faction = alias
	p.alias = alias
	p.device = addDevice(d)
	p.uDevice = p.device.model
	//p.device.owner = &p
	p.grid = gridList[0].(*TGrid) //временно - должен стартовать из публичной сети
	p.maxMatrixCM = p.device.GetMatrixCM()
	p.matrixCM = p.maxMatrixCM
	p.cybercombatSkill = 1
	p.computerSkill = 1
	p.hackingSkill = 1
	p.softwareSkill = 1
	p.body = 1
	p.reaction = 1
	p.willpower = 1
	p.logic = 1
	p.intuition = 1
	p.charisma = 1
	p.edge = 100
	p.maxEdge = 100
	p.id = id
	//p.silentMode = false
	p.simSence = "HOT-SIM"
	p.maxStunCM = (p.willpower+1)/2 + 8
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
	p.freeActionsCount = 1
	p.simpleActionsCount = 2
	//p.specializations

	id++
	return &p
}

//GetAttribute -
func (p *TPersona) GetAttribute(name string) int {
	switch name {
	default:
	case "B":
		return p.body
	case "A":
		return p.agility
	case "R":
		return p.reaction
	case "S":
		return p.strenght
	case "W":
		return p.willpower
	case "L":
		return p.logic
	case "I":
		return p.intuition
	case "C":
		return p.charisma
	case "E":
		return p.edge
	}
	return 0
}

//SetAttribute -
func (p *TPersona) SetAttribute(name string, rating int) {
	switch name {
	default:
	case "B":
		p.body = rating
	case "A":
		p.agility = rating
	case "R":
		p.reaction = rating
	case "S":
		p.strenght = rating
	case "W":
		p.willpower = rating
	case "L":
		p.logic = rating
	case "I":
		p.intuition = rating
	case "C":
		p.charisma = rating
	case "E":
		p.edge = rating
	}
}

//SetSkill -
func (p *TPersona) SetSkill(name string, rating int) {
	if rating == 0 {
		rating = -1
	}
	switch name {
	default:
	case "Cybercombat":
		p.cybercombatSkill = rating
	case "Electronic":
		p.electronicSkill = rating
	case "Hacking":
		p.hackingSkill = rating
	case "Computer":
		p.computerSkill = rating
	case "Hardware":
		p.hardwareSkill = rating
	case "Software":
		p.softwareSkill = rating
	}
}

//SetName -
func (p *TPersona) SetName(name string) {
	p.name = name
}

//GetSpecializationList -
func (p *TPersona) GetSpecializationList() []string {
	return p.specialization
}

//AddSpecialization -
func (p *TPersona) AddSpecialization(skill string, spec string) {
	p.specialization = append(p.specialization, skill+"_"+spec)
}

//HaveValidSpec -
func (p *TPersona) HaveValidSpec(spec []string) (bool, string) {
	for i := range p.specialization {
		for j := range spec {
			if p.specialization[i] == spec[j] {
				return true, spec[j]
			}
		}
	}
	return false, "--NO_SPEC--"
}

//RollInitiative -
func (p *TPersona) RollInitiative() {
	mode := p.GetSimSence()

	switch mode {
	case "AR":
		p.SetInitiative(p.GetReaction() + p.GetIntuition() + xd6Test(1))
	case "COLD-SIM":
		p.SetInitiative(p.GetDataProcessing() + p.GetIntuition() + xd6Test(3))
	case "HOT-SIM":
		p.SetInitiative(p.GetDataProcessing() + p.GetIntuition() + xd6Test(4))
	default:
	}
	/*if p.waitFlag {
		p.SetInitiative(0)
	}*/
	//p.SetInitiative(p.GetDataProcessing() + p.GetIntuition() + xd6Test(4))
}

//GetFullDeffenceFlag -
func (p *TPersona) GetFullDeffenceFlag() bool {
	return p.fullDefFlag
}

//SetFullDeffenceFlag -
func (p *TPersona) SetFullDeffenceFlag(newFDF bool) {
	p.fullDefFlag = newFDF
}

//GetWaitFlag -
func (p *TPersona) GetWaitFlag() bool {
	return p.waitFlag
}

//SetWaitFlag -
func (p *TPersona) SetWaitFlag(newWF bool) {
	p.waitFlag = newWF
}

//IsPlayer -
func (p *TPersona) IsPlayer() bool {
	return p.isPlayer
}

//GetType -
func (p *TPersona) GetType() string {
	return "Persona"
}

//GetFaction -
func (p *TPersona) GetFaction() string {
	return p.faction
}

//GetDevice -
func (p *TPersona) GetDevice() *TDevice {
	return p.device
}

//GetAlias -
func (p *TPersona) GetAlias() string {
	return p.alias
}

//GetName -
func (p *TPersona) GetName() string {
	return p.name
}

//GetInitiative -
func (p *TPersona) GetInitiative() int {
	return p.initiative
}

//SetInitiative -
func (p *TPersona) SetInitiative(init int) {
	p.initiative = init
}

//SetDeviceRating -
func (p *TPersona) SetDeviceRating(init int) {
	p.initiative = init
}

//GetCyberCombatSkill -
func (p *TPersona) GetCyberCombatSkill() int {
	return p.cybercombatSkill
}

//GetElectronicSkill -
func (p *TPersona) GetElectronicSkill() int {
	return p.electronicSkill
}

//GetComputerSkill -
func (p *TPersona) GetComputerSkill() int {
	return p.computerSkill
}

//GetHackingSkill -
func (p *TPersona) GetHackingSkill() int {
	return p.cybercombatSkill
}

//GetSoftwareSkill -
func (p *TPersona) GetSoftwareSkill() int {
	return p.softwareSkill
}

//GetBody -
func (p *TPersona) GetBody() int {
	return p.body
}

//SetBody -
func (p *TPersona) SetBody(newRating int) {
	p.body = newRating
}

//GetReaction -
func (p *TPersona) GetReaction() int {
	return p.reaction
}

//SetReaction -
func (p *TPersona) SetReaction(newRating int) {
	p.reaction = newRating
}

//GetWillpower -
func (p *TPersona) GetWillpower() int {
	return p.willpower
}

//GetLogic -
func (p *TPersona) GetLogic() int {
	return p.logic
}

//GetIntuition -
func (p *TPersona) GetIntuition() int {
	return p.intuition
}

//GetCharisma -
func (p *TPersona) GetCharisma() int {
	return p.charisma
}

//GetEdge -
func (p *TPersona) GetEdge() int {
	return p.edge
}

//GetMaxEdge -
func (p *TPersona) GetMaxEdge() int {
	return p.maxEdge
}

//SetEdge -
func (p *TPersona) SetEdge(newEdge int) {
	p.edge = newEdge
}

//SetMaxEdge -
func (p *TPersona) SetMaxEdge(newEdge int) {
	p.maxEdge = newEdge
}

//GetDeviceRating -
func (p *TPersona) GetDeviceRating() int {
	return p.device.deviceRating
}

//GetDeviceSoft -
func (p *TPersona) GetDeviceSoft() *TProgram {
	return p.device.software
}

///////////////

//GetMatrixCM -
func (p *TPersona) GetMatrixCM() int {
	return p.device.matrixCM
}

//SetMatrixCM -
func (p *TPersona) SetMatrixCM(cmValue int) {
	p.device.matrixCM = cmValue
}

//GetStunCM -
func (p *TPersona) GetStunCM() int {
	return p.stunCM
}

//GetMaxStunCM -
func (p *TPersona) GetMaxStunCM() int {
	return p.maxStunCM
}

//SetStunCM -
func (p *TPersona) SetStunCM(cmValue int) {
	p.stunCM = cmValue
}

/*//ReceiveStunDamage -
func (p *TPersona) ReceiveStunDamage(cmValue int) {
	sDamage := p.stunCM - cmValue
	if p.GetFaction() == player.GetFaction() {
		printLog(strconv.Itoa(sDamage)+" Stun Damage inflicted to "+p.name, congo.ColorYellow)
	}
	if cmValue < 0 {
		pDamage := cmValue / 2
		printLog(strconv.Itoa(sDamage), congo.ColorDefault)
		p.SetPhysCM(p.GetPhysCM() + pDamage)
		if p.GetFaction() == player.GetFaction() {
			printLog(strconv.Itoa(sDamage)+" Physical Damage converted from Stun Damage", congo.ColorRed)
		}
	}
	p.stunCM = cmValue
}*/

//GetPhysCM -
func (p *TPersona) GetPhysCM() int {
	return p.physCM
}

//GetMaxPhysCM -
func (p *TPersona) GetMaxPhysCM() int {
	return p.maxPhysCM
}

//SetPhysCM -
func (p *TPersona) SetPhysCM(cmValue int) {
	p.physCM = cmValue
}

//GetGrid -
func (p *TPersona) GetGrid() *TGrid {
	return p.grid
}

//SetGrid -
func (p *TPersona) SetGrid(grid *TGrid) {
	p.grid = grid
	printLog("WELCOME TO: "+p.grid.name, congo.ColorDefault)
}

//GetHost -
func (p *TPersona) GetHost() *THost {
	if p.host != nil {
		return p.host
	}
	return Matrix
}

//SetHost -
func (p *TPersona) SetHost(host *THost) {
	p.host = host
}

//GetFieldOfView -
func (p *TPersona) GetFieldOfView() FieldOfView {
	//panic("Abs Func Call")
	return p.canSee
}

//GetMarkSet -
func (p *TPersona) GetMarkSet() MarkSet {
	p.ClearMarks()
	/*for i := range p.markSet.MarksFrom {
		if pickObjByID(i) == nil {
			printLog("Mark on obj"+strconv.Itoa(i)+" = "+strconv.Itoa(p.markSet.MarksFrom[i]), congo.ColorYellow)
			delete(p.markSet.MarksFrom, i)
		}
	}*/

	return p.markSet
}

//ClearMarks -
func (p *TPersona) ClearMarks() {
	//hostID := 999999
	for i := range p.markSet.MarksFrom {
		valid := false
		for _, obj := range ObjByNames {
			if obj.GetID() == i {
				valid = true
				//break
			}
		}
		/*	for j := range objectList {
				if objectList[j].(IObj).GetID() == i {

					valid = true
					break
				}

			}
			for k := range gridList {
				if host, ok := gridList[k].(*THost); ok {
					if i == host.GetID() {
						valid = true
						break
					}
				}
			}

			if i == 2 {
				valid = true
			}*/

		if valid == false {
			//p.markSet.MarksFrom[i] = 0
			delete(p.markSet.MarksFrom, i)
		}
	}
}

//ClearLocks -
func (p *TPersona) ClearLocks() {
	//hostID := 999999
	for i := range p.linklocked.LockedByID {
		valid := false
		for _, obj := range ObjByNames {
			if obj.GetID() == i {

				valid = true
				break
			}

		}
		for k := range gridList {
			if host, ok := gridList[k].(*THost); ok {
				if i == host.GetID() {
					valid = true
					break
				}
			}
		}

		/*if i == 2 {
			valid = true
		}*/

		if valid == false {
			//p.markSet.MarksFrom[i] = 0
			delete(p.linklocked.LockedByID, i)
		}
	}
}

//GetLinkLockStatus -
func (p *TPersona) GetLinkLockStatus() Locked {
	p.ClearLocks()
	//panic("Abs Func Call")
	return p.linklocked
}

//LockIcon -
func (p *TPersona) LockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[p.id] = true
}

//UnlockIcon -
func (p *TPersona) UnlockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[p.id] = false
}

//Dumpshock -
func (p *TPersona) Dumpshock() {
	if p.GetMatrixCM() < 1 {
		p.SetAttackRaw(0)
		p.SetSleazeRaw(0)
		p.SetDataProcessingRaw(0)
		p.SetFirewall(0)
		prgs := p.GetRunningPrograms()
		for i := 0; i < len(prgs); i++ {
			p.CrashProgram(prgs[i])
		}
	}
	printLog("Warning!! Dumpshock imminent!!", congo.ColorRed)
	dp1 := p.GetWillpower() + p.GetFirewall()
	suc1, gl, cgl := simpleTest(p.GetID(), dp1, 1000, 0)
	biofeedbackDamage := 6 - suc1
	if gl {
		biofeedbackDamage = biofeedbackDamage + 2
	}
	if cgl {
		biofeedbackDamage = biofeedbackDamage + 2
	}
	if biofeedbackDamage < 0 {
		biofeedbackDamage = 0
	}
	if p.GetSimSence() == "COLD-SIM" {
		p.SetStunCM(p.GetStunCM() - biofeedbackDamage)
		printLog(strconv.Itoa(biofeedbackDamage)+" Stun Damage inflicted by Dumpshock...", congo.ColorRed)
		if p.GetStunCM() < 0 {
			physDamage := p.GetStunCM() / -2
			p.SetPhysCM(p.GetPhysCM() - physDamage)
			printLog(strconv.Itoa(physDamage)+" Physical Damage inflicted by Dumpshock...", congo.ColorRed)
		}
	} else if p.GetSimSence() == "HOT-SIM" {
		p.SetPhysCM(p.GetPhysCM() - biofeedbackDamage)
		printLog(strconv.Itoa(biofeedbackDamage)+" Physical Damage inflicted by Dumpshock...", congo.ColorRed)

	}
	//p.SetSimSence("OFFLINE")
	p.SetInitiative(999999)
	if p.id == player.GetID() {
		printLog("Session terminated...", congo.ColorDefault)
	}
}

//IsConnected -
func (p *TPersona) IsConnected() bool {
	return p.connected
}

//SetConnection -
func (p *TPersona) SetConnection(b bool) {
	p.connected = b
}

//ToggleConnection -
func (p *TPersona) ToggleConnection() {
	if p.connected {
		p.connected = false
	} else {
		p.connected = true
	}
}

//ToggleConvergence -
func (p *TPersona) ToggleConvergence() bool {
	if p.convergenceFlag == false {
		return true
	}
	return false
}

//GetConvergenceFlag -
func (p *TPersona) GetConvergenceFlag() bool {
	return p.convergenceFlag
}

//isOnline -
func (p *TPersona) isOnline() bool {
	mOK := false
	if p.GetMatrixCM() > 0 {
		mOK = true
	}
	sOK := false
	if p.GetStunCM() > 0 {
		sOK = true
	}
	pOK := false
	if p.GetPhysCM() > 0 {
		pOK = true
	}
	tOK := false
	if mOK == true && sOK == true && pOK == true {
		tOK = true
	}
	return tOK
}

//GetPhysicalLocation -
func (p *TPersona) GetPhysicalLocation() bool {
	return p.physLocation
}

//SetPhysicalLocation -
func (p *TPersona) SetPhysicalLocation(location bool) {
	p.physLocation = location
}

/*//LoadProgram -
func (p *TPersona) LoadProgram(name string) bool {
	for i := 0; i < len(p.device.software.programName); i++ {
		if p.device.software.programName[i] == name {
			if p.device.GetRunningProgramsQty() < p.device.GetMaxRunningPrograms() { //тест проверки на то может ли загрузиться данная программа
				if p.device.software.programStatus[i] == "inStore" {
					p.device.software.programStatus[i] = "Running"
				} else {
					printLog("Error: Program '"+name+"' is "+p.device.software.programStatus[i], congo.ColorYellow)
				}
				p.device.software.programStatus[i] = "Running"
			} else {
				printLog("Error: No free program slots available", congo.ColorYellow)
				return false
			}
		}
	}
	return true
}*/

//UnloadProgram -
func (p *TPersona) UnloadProgram(name string) bool {
	for i := 0; i < len(p.device.software.programName); i++ {
		if p.device.software.programName[i] == name {
			if p.device.software.programStatus[i] == "Running" {
				p.device.software.programStatus[i] = "Stored"
			} else {
				return false
			}
		}
	}
	return true
}

//CrashProgram -
func (p *TPersona) CrashProgram(name string) bool {
	for i := 0; i < len(p.device.software.programName); i++ {
		if p.device.software.programName[i] == name {
			if p.device.software.programStatus[i] == "Running" {
				p.device.software.programStatus[i] = "Crashed"
				if p.GetFaction() == player.GetFaction() {
					printLog("WARNING! "+p.device.software.programName[i]+" program crased!", congo.ColorYellow)
				}
			}
		}
	}
	return true
}

//CheckRunningProgram -
func (p *TPersona) CheckRunningProgram(name string) bool {
	programs := p.GetPrograms()
	for i := range programs {
		if programs[i].programName != name {
			continue
		}
		if programs[i].programStatus == "Running" {
			return true
		}
	}
	/*for i := range p.device.software.programName {
		if p.device.software.programName[i] == name {
			if p.device.software.programStatus[i] == "Running" {
				return true
			}
		}
	}*/
	return false
}

//GetRunningPrograms -
func (p *TPersona) GetRunningPrograms() []string {
	var prgs []string
	for i := range p.device.software.programName {
		if p.device.software.programStatus[i] == "Running" {
			prgs = append(prgs, p.device.software.programName[i])
		}
	}
	return prgs
}

//CrashRandomProgram -
func (p *TPersona) CrashRandomProgram() bool {
	prgs := p.GetRunningPrograms()
	r := rand.Intn(32)
	for i := 0; i < r; i++ {
		shuffleString(prgs)
	}
	if len(prgs) > 0 {
		p.CrashProgram(prgs[0])
	}
	return true
}

//GetMaxRunningPrograms -
func (p *TPersona) GetMaxRunningPrograms() int {
	prgBoost := 0
	if p.device.CheckRunningProgram("Virtual Machine") {
		prgBoost = 2
	}
	p.device.maxRunningPrograms = p.device.deviceRating + prgBoost
	return p.device.maxRunningPrograms
}

//ReceiveMatrixDamage -
func (p *TPersona) ReceiveMatrixDamage(damage int) {
	if p.CheckRunningProgram("Virtual Machine") && damage > 0 {
		damage++
		if p.GetFaction() == player.GetFaction() {
			printLog("...WARNING! 1 additional Matrix Damage caused by Virtual Machine program", congo.ColorYellow)
			hold()
		}
	}
	p.SetMatrixCM(p.GetMatrixCM() - damage)
	if p.GetFaction() == player.GetFaction() {
		printLog("..."+p.GetName()+" takes "+strconv.Itoa(damage)+" Matrix damage", congo.ColorYellow)
		hold()
	}
	if p.GetMatrixCM() < 1 {
		//p.Dumpshock()
		p.SetSimSence("OFFLINE")
	}
}

//ReceiveBiofeedbackDamage -
func (p *TPersona) ReceiveBiofeedbackDamage(damage int) {
	if p.GetSimSence() == "HOT-SIM" {
		p.SetPhysCM(p.GetPhysCM() - damage)
		printLog(p.GetName()+" takes "+strconv.Itoa(damage)+" Physical damage", congo.ColorYellow)
	} else if p.GetSimSence() == "COLD-SIM" {
		p.SetStunCM(p.GetStunCM() - damage)
		//p.ReceiveStunDamage(damage)
		printLog(p.GetName()+" takes "+strconv.Itoa(damage)+" Stun damage", congo.ColorYellow)
	} else if p.GetSimSence() == "AR" {
		printLog("Biofeedback code detected", congo.ColorYellow)
	} else {
		printLog("--DEBUG--: Simsence Mode Error: func (p *TPersona) ReceiveBiofeedbackDamage(damage int)", congo.ColorDefault)
	}

}

//ReceivePhysBiofeedbackDamage -
func (p *TPersona) ReceivePhysBiofeedbackDamage(damage int) {
	p.SetPhysCM(p.GetPhysCM() - damage)
	if p.GetFaction() == player.GetFaction() {
		printLog(p.GetName()+" takes "+strconv.Itoa(damage)+" Physical damage", congo.ColorYellow)
		hold()
	}
}

//ReceiveStunBiofeedbackDamage -
func (p *TPersona) ReceiveStunBiofeedbackDamage(damage int) {
	p.SetStunCM(p.GetStunCM() - damage)
	if p.GetFaction() == player.GetFaction() {
		printLog(p.GetName()+" takes "+strconv.Itoa(damage)+" Stun damage", congo.ColorYellow)
		hold()
	}
}

//ResistMatrixDamage -
func (p *TPersona) ResistMatrixDamage(damage int) int {
	resistDicePool := 0
	resistDicePool = resistDicePool + p.GetDeviceRating()
	resistDicePool = resistDicePool + p.GetFirewall()
	printLog("...Incoming matrix damage detected", congo.ColorGreen)

	if p.CheckRunningProgram("Shell") {
		resistDicePool = resistDicePool + 1
		if p.GetFaction() == player.GetFaction() {
			printLog("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if p.CheckRunningProgram("Armor") {
		resistDicePool = resistDicePool + 2
		if p.GetFaction() == player.GetFaction() {
			printLog("Program Armor: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	printLog("...Evaluated Firewall resources: "+strconv.Itoa(resistDicePool)+" mp/p", congo.ColorGreen)
	damageSoak, gl, cgl := simpleTest(p.GetID(), resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			printLog(p.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			printLog(p.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if p.GetFaction() == player.GetFaction() {
		printLog("..."+p.GetName()+": "+strconv.Itoa(damageSoak)+" Matrix damage soaked", congo.ColorGreen)
		hold()
	}
	return realDamage
}

//ResistBiofeedbackDamage -
func (p *TPersona) ResistBiofeedbackDamage(damage int) int {
	resistDicePool := 0
	resistDicePool = resistDicePool + p.GetWillpower()
	resistDicePool = resistDicePool + p.GetFirewall()
	if p.CheckRunningProgram("Shell") {
		resistDicePool = resistDicePool + 1
		if p.GetFaction() == player.GetFaction() {
			printLog("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if p.CheckRunningProgram("Biofeedback Filter") {
		resistDicePool = resistDicePool + 2
		if p.GetFaction() == player.GetFaction() {
			printLog("Program Biofeedback Filter: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	damageSoak, gl, cgl := simpleTest(p.GetID(), resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			printLog(p.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			printLog(p.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if p.GetFaction() == player.GetFaction() {
		printLog("..."+p.GetName()+": "+strconv.Itoa(damageSoak)+" Biofeedback damage soaked", congo.ColorGreen)
		hold()
	}
	return realDamage
}

//CheckConvergence -
func (p *TPersona) CheckConvergence() {
	if p.grid.GetOverwatchScore() > 39 {
		if p.convergenceFlag == false {
			p.convergenceFlag = true
			if p.GetFaction() == player.GetFaction() {
				p.ToggleConvergence()
				printLog("Warning!!! Convergence protocol engaged...", congo.ColorRed)
				//printLog("Warning!!! Convergence protocol engaged...", congo.ColorRed)
			}
			convergenceDamage := p.ResistMatrixDamage(12)
			p.ReceiveMatrixDamage(convergenceDamage)
			if p.GetHost() == Matrix {
				p.ToggleConnection()
				p.ClearMarks()

				if p.GetMatrixCM() > 0 {
					printLog("1", congo.ColorDefault)
					p.Dumpshock()
				}
				if p.GetFaction() == player.GetFaction() {
					printLog("All MARKs have been lost...", congo.ColorRed)
					hold()
				}
			} else {
				host := p.GetHost()
				host.markSet.MarksFrom[p.GetID()] = 3
				if p.GetFaction() == player.GetFaction() {
					printLog(host.GetName()+" "+host.GetType()+" now have 3 MARKs on "+p.GetName()+"...", congo.ColorRed)
					hold()
				}
			}
			p.SetPhysicalLocation(true)
			if p.GetFaction() == player.GetFaction() {
				printLog("Physical location discovered...", congo.ColorYellow)
				hold()
				printLog("...reporting physical location to nearest police force", congo.ColorGreen)
				hold()
				printLog("...calling nearest GOD agent", congo.ColorGreen)
				hold()
				if p.GetName() == player.GetName() {
					printLog("Grid Overwatch Division wishes you to have a nice day!", congo.ColorDefault)
				}
				hold()
			}
		} else {
			if p.GetHost() == Matrix && p.GetMatrixCM() > 0 && p.GetInitiative() < 9999 {
				p.ToggleConnection()
				p.ClearMarks()
				p.Dumpshock()
				if p.GetFaction() == player.GetFaction() {
					printLog("All MARKs have been lost...", congo.ColorRed)
					hold()
				}
			}
		}

	}
}

//TriggerDataBomb -
func (p *TPersona) TriggerDataBomb(bombRating int) {
	host := p.GetHost()
	printLog("...Warning! Databomb triggered", congo.ColorRed)
	prgBonus := 0
	if p.CheckRunningProgram("Armor") {
		prgBonus = prgBonus + 2
	}
	if p.CheckRunningProgram("Shell") {
		prgBonus = prgBonus + 1
	}
	if p.CheckRunningProgram("Defuse") {
		prgBonus = prgBonus + 4
	}
	resistPool := p.GetDeviceRating() + p.GetFirewall() + prgBonus
	resistHits, rgl, rcgl := simpleTest(p.GetID(), resistPool, 999, 0)
	//остановиться и перебросить при необходимости
	fullDamage := xd6Test(bombRating)
	if rgl == true {
		fullDamage = fullDamage + bombRating
		printLog("...Warning!! Firewall error erupted", congo.ColorYellow)
	}
	if rcgl == true {
		//addOverwatchScore(xd6Test(trg.GetDataBombRating()))
		fullDamage = fullDamage + xd6Test(bombRating)
		printLog("...Danger!! Critical error erupted", congo.ColorRed)
	}
	printLog("..."+strconv.Itoa(resistHits)+" of incomming Matrix damage has been resisted", congo.ColorGreen)
	realDamage := fullDamage - resistHits
	if realDamage < 0 {
		realDamage = 0
	}
	p.ReceiveMatrixDamage(realDamage)
	printLog("...Databomb destroyed", congo.ColorGreen)
	if host.GetHostAlertStatus() != "Active Alert" {
		host.SetAlert("Active Alert")
	}
}

//CountMarks -
func (p *TPersona) CountMarks() int {
	totalMarks := 0
	for r := range p.markSet.MarksFrom {
		if p.markSet.MarksFrom[r] > 0 && r != p.id {
			totalMarks = totalMarks + p.markSet.MarksFrom[r]
		}
	}
	return totalMarks
}

//GetOverwatchScore -
func (p *TPersona) GetOverwatchScore() int {
	return p.grid.overwatchScore
}

//SetOverwatchScore -
func (p *TPersona) SetOverwatchScore(newScore int) {
	p.grid.overwatchScore = newScore
}

//GetLastSureOS -
func (p *TPersona) GetLastSureOS() int {
	return p.grid.lastSureOS
}

//SetLastSureOS -
func (p *TPersona) SetLastSureOS(newScore int) {
	p.grid.lastSureOS = newScore
}

//GetSearchResultIn -
func (p *TPersona) GetSearchResultIn() int {
	return p.searchLen
}

//SetSearchResultIn -
func (p *TPersona) SetSearchResultIn(val int) {
	p.searchLen = val
}

//GetSearchProcess -
func (p *TPersona) GetSearchProcess() SearchProcess {
	return p.searchProcessStatus
}

//SetSearchProcess -
func (p *TPersona) SetSearchProcess(turns int, sIconType, sIconName string) {
	p.searchProcessStatus.SearchTime = append(p.searchProcessStatus.SearchTime, turns)
	p.searchProcessStatus.SearchIconName = append(p.searchProcessStatus.SearchIconName, sIconName)
	p.searchProcessStatus.SearchIconType = append(p.searchProcessStatus.SearchIconType, sIconType)
	p.searchProcessStatus.SpentTurns = append(p.searchProcessStatus.SpentTurns, 0)
}

//UpdateSearchProcess -
func (p *TPersona) UpdateSearchProcess() {
	host := player.GetHost()
	for i := range p.searchProcessStatus.SpentTurns {
		if i < (len(p.searchProcessStatus.SearchIconName)) {
			p.searchProcessStatus.SpentTurns[i] = p.searchProcessStatus.SpentTurns[i] + 1
			if p.searchProcessStatus.SpentTurns[i] == p.searchProcessStatus.SearchTime[i] {
				p.waitFlag = false
				switch formatTargetName(p.searchProcessStatus.SearchIconType[i]) {
				case "Host":
					hostName := formatTargetName(p.searchProcessStatus.SearchIconName[i])
					if HostExist(hostName) {
						ImportHostFromDB(hostName)
					} else {
						player.GetGrid().NewHost(hostName, 0) // -DEBUG: тут можно указать какой хост создавать (1-12: рейтинг/0: рандом)
					}
				case "File":
					if host != Matrix {
						host.NewFile(p.searchProcessStatus.SearchIconName[i])
					} else {
						printLog("Matrix is just to vast. No File can be found outside the host...", congo.ColorGreen)
						printLog("DEBUG: сложности с использованием прото-хоста в качестве носителя объектов", congo.ColorGreen)
					}

				default:
				}
				printLog("Connection with "+p.searchProcessStatus.SearchIconType[i]+" '"+p.searchProcessStatus.SearchIconName[i]+"' established " /* + host.GetName()*/, congo.ColorGreen)
				p.RollInitiative()
				p.searchProcessStatus.SearchIconName = append(p.searchProcessStatus.SearchIconName[:i], p.searchProcessStatus.SearchIconName[i+1:]...)
				p.searchProcessStatus.SearchIconType = append(p.searchProcessStatus.SearchIconType[:i], p.searchProcessStatus.SearchIconType[i+1:]...)
				p.searchProcessStatus.SpentTurns = append(p.searchProcessStatus.SpentTurns[:i], p.searchProcessStatus.SpentTurns[i+1:]...)
				p.searchProcessStatus.SearchTime = append(p.searchProcessStatus.SearchTime[:i], p.searchProcessStatus.SearchTime[i+1:]...)
			}
		}
		//kill search process
	}
}

//SetDownloadProcess -
func (p *TPersona) SetDownloadProcess(size int, fileName string) {
	p.downloadProcessStatus.FileSize = append(p.downloadProcessStatus.FileSize, size)
	p.downloadProcessStatus.DownloadIconName = append(p.downloadProcessStatus.DownloadIconName, fileName)
	p.downloadProcessStatus.DownloadedData = append(p.downloadProcessStatus.DownloadedData, 0)
}

//UpdateDownloadProcess -
func (p *TPersona) UpdateDownloadProcess() {
	for i := range p.downloadProcessStatus.DownloadIconName {
		p.downloadProcessStatus.DownloadedData[i] = p.downloadProcessStatus.DownloadedData[i] + p.GetDataProcessing()*5
		if p.downloadProcessStatus.DownloadedData[i] >= p.downloadProcessStatus.FileSize[i] {
			p.waitFlag = false
			printLog("Downloading of "+p.downloadProcessStatus.DownloadIconName[i]+" complete", congo.ColorGreen)
			p.RollInitiative()
			p.downloadProcessStatus.DownloadIconName = append(p.downloadProcessStatus.DownloadIconName[:i], p.downloadProcessStatus.DownloadIconName[i+1:]...)
			p.downloadProcessStatus.FileSize = append(p.downloadProcessStatus.FileSize[:i], p.downloadProcessStatus.FileSize[i+1:]...)
			p.downloadProcessStatus.DownloadedData = append(p.downloadProcessStatus.DownloadedData[:i], p.downloadProcessStatus.DownloadedData[i+1:]...)
		}
	}
}

//GetDownloadProcess -
func (p *TPersona) GetDownloadProcess() DownloadProcess {
	return p.downloadProcessStatus
}

//SetSimSence -
func (p *TPersona) SetSimSence(smsence string) {
	smsence = formatString(smsence)
	smsence = cleanText(smsence)
	switch smsence {
	case "AR":
		p.simSence = "AR"
	case "COLD-SIM":
		p.simSence = "COLD-SIM"
	case "HOT-SIM":
		p.simSence = "HOT-SIM"
	case "OFFLINE":
		if p.GetSimSence() != "AR" {
			p.Dumpshock()
		}
		p.simSence = "OFFLINE"
	default:
	}
}

///////////////////////////////////////////////////////
//File

//TFile -
type TFile struct {
	owner    IObj //*TPersona
	name     string
	fileName string
	//encryptionFlag bool - ////////////Flags aren't nesessary
	//dataBombFlag bool
	encryptionRating int
	dataBombRating   int
	size             int
	value            int
	lastEditTime     string
	TIcon
	//TObj
}

//IFile -
type IFile interface {
	IIcon //сделать IFileOnly
	IFileOnly
}

//IFileOnly -
type IFileOnly interface {
	//SetOwner(IObj)
	//SetName(string)
	SetFileName(string)
	GetFileName() string
	GetEncryptionRating() int
	SetEncryptionRating(int)
	GetDataBombRating() int
	SetDataBombRating(int)
	GetSize() int
	SetSize(int)
	GetLastEditDate() string
	SetLastEditDate(string)
	GetValue() int
}

var _ IFile = (*TFile)(nil)

//NewFile -
func (h *THost) NewFile(name string) *TFile {
	f := TFile{}
	data := player.GetFieldOfView().KnownData[f.id]
	data[0] = "Unknown"
	data[1] = "Unknown"
	data[3] = "Unknown"
	data[12] = "Unknown"
	data[13] = "Unknown"
	data[15] = "Unknown"
	data[18] = "Unknown"
	f.owner = h
	f.host = h
	hRatMod := 1
	congo.WindowsMap.ByTitle["Process"].WPrint(".", congo.ColorGreen)
	if f.host == Matrix {
		f.grid = player.GetGrid()
		//f.host = Matrix
		f.device = addDevice("noDevice")
		f.silentMode = false
		ObjByNames[f.name] = &f
		return &f
		//} else {
		//h.GetGrid()
		//f.grid = h.GetGrid()

		/*for _, obj := range ObjByNames {
			if gridToPick, ok := obj.(*TGrid); ok {
				f.grid = gridToPick
			}
		}*/
	}
	congo.WindowsMap.ByTitle["Process"].WPrint(".", congo.ColorGreen)
	enRat, _, _ := simpleTest(h.GetID(), h.deviceRating+h.deviceRating, h.dataProcessing, 0)
	f.encryptionRating = enRat
	bombRat, _, _ := simpleTest(h.GetID(), h.deviceRating+h.deviceRating, h.sleaze, 0)
	f.dataBombRating = bombRat
	f.lastEditTime = generateLastEditTime()

	hRatMod = (h.deviceRating-1)/3 + 1

	f.device = addDevice("noDevice")
	f.name = f.GetType() + " " + strconv.Itoa(id)

	//file := NewFile("File")
	dataDensity := xd6Test(2)
	f.size = hRatMod * 5 * dataDensity

	f.id = id
	congo.WindowsMap.ByTitle["Process"].WPrint(".", congo.ColorGreen)
	f.faction = h.faction
	if name == "random" {
		f.fileName = generateFileName()
		f.silentMode = true
		r := rand.Intn(9)
		if r > 4 {
			f.silentMode = false
			data[0] = "Spotted"
			//f.nameName = generateFileName()
		}
	} else {
		f.fileName = name
		f.silentMode = false
		data[0] = "Spotted"
	}
	//f.silentMode = true
	f.markSet.MarksFrom = make(map[int]int)
	f.markSet.MarksFrom[f.id] = 4

	//r := rand.Intn(9)
	/*if r > 4 {
		f.silentMode = false
		data[0] = "Spotted"
		f.name = generateFileName()
	}*/
	player.GetFieldOfView().KnownData[f.id] = data
	//f.grid = "ARES GRID"
	//f.name = f.GetType() + " " + strconv.Itoa(id)
	ObjByNames[f.name] = &f
	id = id + xd6Test(3)
	return &f

}

//GetType -
func (f *TFile) GetType() string {
	return "File"
}

//GetFaction -
func (f *TFile) GetFaction() string {
	return f.faction
}

//GetOwner -
func (f *TFile) GetOwner() IObj {
	return f.owner
}

//GetName -
func (f *TFile) GetName() string {
	return f.name
}

//GetFileName -
func (f *TFile) GetFileName() string {
	return f.fileName
}

//GetEncryptionRating -
func (f *TFile) GetEncryptionRating() int {
	return f.encryptionRating
}

//GetDataBombRating -
func (f *TFile) GetDataBombRating() int {
	return f.dataBombRating
}

//SetOwner -
func (f *TFile) SetOwner(owner IObj) {
	f.owner = owner
}

//SetName -
func (f *TFile) SetName(name string) {
	f.name = name
}

//SetFileName -
func (f *TFile) SetFileName(name string) {
	f.fileName = name
}

//SetEncryptionRating -
func (f *TFile) SetEncryptionRating(rating int) {
	f.encryptionRating = rating
}

//SetDataBombRating -
func (f *TFile) SetDataBombRating(rating int) {
	f.dataBombRating = rating
}

//GetGrid -
func (f *TFile) GetGrid() *TGrid {
	return f.grid
}

//SetGrid -
func (f *TFile) SetGrid(grid *TGrid) {
	f.grid = grid
}

//GetSize -
func (f *TFile) GetSize() int {
	return f.size
}

//SetSize -
func (f *TFile) SetSize(newSize int) {
	f.size = newSize
}

//GetLastEditDate -
func (f *TFile) GetLastEditDate() string {
	return f.lastEditTime
}

//SetLastEditDate -
func (f *TFile) SetLastEditDate(newTime string) {
	f.lastEditTime = STime
}

//GetValue -
func (f *TFile) GetValue() int {
	return f.value
}

//GetHost -
func (f *TFile) GetHost() *THost {
	return f.host
}

//GetLongAct -
func (p *TPersona) GetLongAct() int {
	return p.searchLen
}

//GetDeviceRating -
func (f *TFile) GetDeviceRating() int {
	if f.device.deviceType == "noDevice" {
		return f.host.deviceRating
	}
	return f.device.deviceRating
}

//GetAttack -
func (f *TFile) GetAttack() int {
	if f.device.deviceType == "noDevice" {
		return f.host.attack
	}
	return f.device.attack
}

//GetSleaze -
func (f *TFile) GetSleaze() int {
	if f.device.deviceType == "noDevice" {
		return f.host.sleaze
	}
	return f.device.sleaze
}

//GetDataProcessing -
func (f *TFile) GetDataProcessing() int {
	//if f.device.deviceType == "noDevice" {
	if f.GetDevice().deviceType == "noDevice" {
		return f.GetHost().GetDataProcessing()
	}
	return f.device.dataProcessing
}

//GetFirewall -
func (f *TFile) GetFirewall() int {
	if f.device.deviceType == "noDevice" {
		return f.host.firewall
	}
	return f.device.firewall
}

//RollInitiative -
func (f *TFile) RollInitiative() {
	f.SetInitiative(-1)
}

//GetIntuition -
func (f *TFile) GetIntuition() int {
	owner := f.GetOwner()
	if persona, ok := owner.(IPersona); ok {
		return persona.GetIntuition()
	}
	if host, ok := owner.(IHost); ok {
		return host.GetDeviceRating()
	}
	if device, ok := owner.(IDevice); ok {
		return device.GetDeviceRating()
	}
	return 0
}
