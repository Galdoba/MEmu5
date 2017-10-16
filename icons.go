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

//TObj -
type TObj struct {
	uType        string
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
	GetMarkSet() MarkSet
	GetFieldOfView() FieldOfView
	GetLinkLockStatus() Locked
	ChangeFOWParametr(int, int, string)
	Scanable() bool
	CountMarks() int
	GetDeviceRating() int
	SetDeviceRating(int)
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
		for j := range objectList {
			if obj, ok := objectList[j].(*TObj); ok {
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
	//congo.WindowsMap.ByTitle["Log"].WPrintLn("len(ObjectList) = "+strconv.Itoa(len(objectList)), congo.ColorDefault)

	for i := 0; i < len(objectList); i++ {
		if id == objectList[i].(IObj).GetID() {
			return objectList[i] //.(*TObj)
		}
	}
	//check if obj is *THOST
	for i := 0; i < len(gridList); i++ {
		if host, ok := gridList[i].(*THost); ok {
			if host.GetID() == id {
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
	grid         TGrid
	lastLocation TGrid
	host         *THost
	device       *TDevice

	simSence        string
	silentMode      bool
	initiative      int
	id              int
	owner           string
	isPlayer        bool
	convergenceFlag bool
	connected       bool
	searchLen       int
}

//IIcon - в икону входят файлы, персоны, айсы и хосты
type IIcon interface {
	IObj
	IIconOnly
}

var _ IIcon = (*TIcon)(nil)

//IIcon - в икону входят файлы, персоны, айсы и хосты
type IIconOnly interface {
	GetSilentRunningMode() bool
	SetSilentRunningMode(bool)
	GetGrid() TGrid
	//GetGridName() string
	SetGrid(TGrid)
	GetHost() *THost
	SetHost(*THost)
	//GetID() int
	//SetID()
	GetInitiative() int
	SetInitiative(int)
	GetSimSence() string
	SetSimSence(string)
	GetOwner() string
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
	//GetDeviceRating() int
	GetAttack() int
	GetSleaze() int
	GetDataProcessing() int
	GetFirewall() int
	CheckRunningProgram(string) bool
	ResistMatrixDamage(int) int
}

//CheckRunningProgram -
func (i *TIcon) CheckRunningProgram(name string) bool {
	//return false
	noDevice := i.GetDevice() //.GetSoftwareList()
	if noDevice == nil {
		i.device = addDevice("noDevice")
	}
	for j := range i.GetDevice().GetSoftwareList().programName {
		if i.GetDevice().software.programName[j] == name {
			if i.device.software.programStatus[j] == "Running" {
				return true
				//test.programName[0] = test.programName[0] + "__"
			}
		}
	}
	return false
}

//GetOverwatchScore -
func (i *TIcon) GetOverwatchScore() int {
	return i.grid.overwatchScore
}

//GetDevice -
func (i *TIcon) GetDevice() *TDevice {
	return i.device
}

//GetDeviceRating -
func (i *TIcon) GetDeviceRating() int {
	return i.device.deviceRating
}

//GetAttack -
func (i *TIcon) GetAttack() int {
	boost := 0
	if i.CheckRunningProgram("Decryption") {
		boost = 1
	}
	att := i.device.attack + i.device.attackMod + boost
	if att < 0 {
		return 0
	}
	return att
}

//GetAttackMod -
func (i *TIcon) GetAttackMod() int {
	return i.device.attackMod
}

//GetAttackRaw -
func (i *TIcon) GetAttackRaw() int {
	return i.device.attack
}

//SetAttack -
func (i *TIcon) SetAttack(newAttack int) {
	i.device.attack = newAttack
} //Возможно не нужен

//SetAttackMod -
func (i *TIcon) SetAttackMod(newAttack int) {
	i.device.attackMod = newAttack
}

//SetAttackRaw -
func (i *TIcon) SetAttackRaw(newAttack int) {
	i.device.attack = newAttack
}

//GetSleaze -
func (i *TIcon) GetSleaze() int {
	boost := 0
	if i.CheckRunningProgram("Stealth") {
		boost = 1
	}
	slz := i.device.sleaze + i.device.sleazeMod + boost
	if slz < 0 {
		return 0
	}
	return slz
}

//GetSleazeMod -
func (i *TIcon) GetSleazeMod() int {
	return i.device.sleazeMod
}

//GetSleazeRaw -
func (i *TIcon) GetSleazeRaw() int {
	return i.device.sleaze
}

//SetSleaze -
func (i *TIcon) SetSleaze(newSleaze int) {
	i.device.sleaze = newSleaze
}

//SetSleazeMod -
func (i *TIcon) SetSleazeMod(newSleaze int) {
	i.device.sleazeMod = newSleaze
}

//SetSleazeRaw -
func (i *TIcon) SetSleazeRaw(newSleaze int) {
	i.device.sleaze = newSleaze
}

//GetDataProcessing -
func (i *TIcon) GetDataProcessing() int {
	boost := 0
	if i.CheckRunningProgram("Toolbox") {
		boost = 1
	}
	dtp := i.device.dataProcessing + i.device.dataProcessingMod + boost
	if dtp < 0 {
		return 0
	}
	return dtp
}

//GetDataProcessingMod -
func (i *TIcon) GetDataProcessingMod() int {
	return i.device.dataProcessingMod
}

//GetDataProcessingRaw -
func (i *TIcon) GetDataProcessingRaw() int {
	return i.device.dataProcessing
}

//SetDataProcessing -
func (i *TIcon) SetDataProcessing(newDataProcessing int) {
	i.device.dataProcessing = newDataProcessing
}

//SetDataProcessingMod -
func (i *TIcon) SetDataProcessingMod(newDataProcessing int) {
	i.device.dataProcessingMod = newDataProcessing
}

//SetDataProcessingRaw -
func (i *TIcon) SetDataProcessingRaw(newDataProcessing int) {
	i.device.dataProcessing = newDataProcessing
}

//GetFirewall -
func (i *TIcon) GetFirewall() int {
	boost := 0
	if i.CheckRunningProgram("Encryption") {
		boost = 1
	}
	fwl := i.device.firewall + i.device.firewallMod + boost
	if fwl < 0 {
		return 0
	}
	return fwl
}

//GetFirewallMod -
func (i *TIcon) GetFirewallMod() int {
	return i.device.firewallMod
}

//GetFirewallRaw -
func (i *TIcon) GetFirewallRaw() int {
	return i.device.firewall
}

//SetFirewall -
func (i *TIcon) SetFirewall(newFirewall int) {
	i.device.firewall = newFirewall
}

//SetFirewallMod -
func (i *TIcon) SetFirewallMod(newFirewallMod int) {
	i.device.firewallMod = i.device.firewallMod + newFirewallMod
}

//SetFirewallRaw -
func (i *TIcon) SetFirewallRaw(newFirewall int) {
	i.device.firewall = newFirewall
}

//ResistMatrixDamage -
func (i *TIcon) ResistMatrixDamage(damage int) int {
	resistDicePool := 0
	resistDicePool = resistDicePool + i.GetDeviceRating()
	resistDicePool = resistDicePool + i.GetFirewall()
	if i.CheckRunningProgram("Shell") {
		resistDicePool = resistDicePool + 1
		if i.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if i.CheckRunningProgram("Armor") {
		resistDicePool = resistDicePool + 2
		if i.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program Armor: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	damageSoak, gl, cgl := simpleTest(resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if i.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			congo.WindowsMap.ByTitle["Log"].WPrintLn(i.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if i.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(i.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if i.GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn(i.GetName()+": "+strconv.Itoa(damageSoak)+" Matrix damage soaked...", congo.ColorGreen)
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
		i.simSence = "Cold-SIM VR"
	case "HOT-SIM":
		i.simSence = "Hot-SIM VR"
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
func (i *TIcon) GetGrid() TGrid {
	return i.grid
}

//GetGridName -
func (i *TIcon) GetGridName() string {
	return i.grid.name
}

//SetGrid -
func (i *TIcon) SetGrid(grid TGrid) {
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
func (i *TIcon) GetOwner() string {
	return i.owner
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
	printLog("...error: this icon type is immune to Matrix Damage", congo.ColorYellow)
}

//GetLongAct -
func (i *TIcon) GetLongAct() int {
	return i.searchLen
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
	IsLoaded() bool
	SetLoadStatus(bool)
	GetMatrixCM() int
	SetMatrixCM(int)
	GetActionReady() int
	SetActionReady(int)
	GetLastTargetName() string
	SetLastTargetName(string)
	TakeFOWfromHost()
}

//NewIC -
func (h *THost) NewIC(name string) *TIC {
	id = id + xd6Test(3)
	i := TIC{}
	i.name = name
	i.uType = "IC"
	i.host = h
	i.owner = h.GetName()
	i.deviceRating = h.deviceRating
	i.attack = h.attack
	i.sleaze = h.sleaze
	i.dataProcessing = h.dataProcessing
	i.firewall = h.firewall
	i.id = id + xd6Test(3)
	i.isLoaded = true
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
	data := player.canSee.KnownData[i.id]
	data[0] = "Spotted"
	data[2] = "Unknown"
	data[5] = "Unknown"
	data[7] = "Unknown"
	data[8] = "Unknown"
	data[9] = "Unknown"
	data[10] = "Unknown"
	data[11] = "Unknown"
	//data[13] = "Unknown"
	player.canSee.KnownData[i.id] = data
	if i.name == "Patrol IC" {
		i.actionReady = calculatePartolScan(i.deviceRating)
	} else {
		i.actionReady = -1
	}
	objectList = append(objectList, &i)
	id++

	return &i
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
	//storedPrograms
	modifications []string
	matrixCM      int
	maxMatrixCM   int
	deviceType    string
	model         string
	owner         string //*TPersona
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
		d.simSence = "Hot-SIM VR"
		//d.grid = "Public Grid"
		//d.id = id
	}
	d.SetID()
	d.maxRunningPrograms = d.deviceRating
	//d.software = make([]TProgram, 20)
	//add all soft:
	//d.software = preaparePrograms()
	//d.software.programName = append(d.software.programName, "brows")
	d.markSet.MarksFrom = make(map[int]int)
	d.markSet.MarksFrom[d.GetID()] = 4
	d.canSee.KnownData = make(map[int][30]string)
	d.linklocked.LockedByID = make(map[int]bool)
	d.name = d.GetType() + " " + strconv.Itoa(d.id)
	objectList = append(objectList, &d)
	//id++
	return &d
}

//LoadProgram -
func (d *TDevice) LoadProgram(name string) bool {
	for i := 0; i < len(d.software.programName); i++ {
		if d.software.programName[i] == name {
			if d.GetRunningProgramsQty() < d.GetMaxRunningPrograms() { //тест проверки на то может ли загрузиться данная программа
				d.software.programStatus[i] = "Running"
			} else {
				return false
			}
		}
	}
	return true
}

//UnloadProgram -
func (d *TDevice) UnloadProgram(name string) bool {
	for i := 0; i < len(d.software.programName); i++ {
		if d.software.programName[i] == name {
			if d.software.programStatus[i] == "Running" {
				d.software.programStatus[i] = "inStore"
			} else {
				return false
			}

		}
	}
	return true
}

//GetRunningProgramsQty -
func (d *TDevice) GetRunningProgramsQty() int {
	d.curRunningPrograms = 0
	for i := range d.software.programStatus {
		if d.software.programStatus[i] == "Running" {
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
func (d *TDevice) GetGrid() TGrid {
	return d.grid
}

//SetGrid -
func (d *TDevice) SetGrid(grid TGrid) {
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
	for i := range d.software.programName {
		if d.software.programName[i] == name {
			if d.software.programStatus[i] == "Running" {
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
	name             string
	alias            string
	userMode         string
	device           *TDevice
	computerSkill    int
	hackingSkill     int
	softwareSkill    int
	electronicSkill  int
	hardwareSkill    int
	cybercombatSkill int
	initiative       int
	body             int
	logic            int
	intuition        int
	willpower        int
	charisma         int
	maxPhysCM        int
	physCM           int
	maxStunCM        int
	stunCM           int
	maxMatrixCM      int
	matrixCM         int
	id               int
	physLocation     bool
	markSet          MarkSet
}

//IPersona -
type IPersona interface {
	IIcon
	//IsPlayer() bool
	//GetDeviceRating() int
	GetMatrixCM() int
	GetHackingSkill() int
	GetCyberCombatSkill() int
	GetComputerSkill() int
	GetElectronicSkill() int
	GetSoftwareSkill() int
	GetBody() int
	GetWillpower() int
	GetLogic() int
	GetIntuition() int

	/*GetAttack() int
	GetAttackMod() int
	GetAttackRaw() int
	GetSleaze() int
	GetSleazeMod() int
	GetSleazeRaw() int
	GetDataProcessing() int
	GetDataProcessingMod() int
	GetDataProcessingRaw() int
	GetFirewall() int
	GetFirewallMod() int
	GetFirewallRaw() int
	SetDeviceAttack(int)
	SetDeviceAttackMod(int)
	SetDeviceAttackRaw(int)
	SetDeviceSleaze(int)
	SetDeviceSleazeMod(int)
	SetDeviceSleazeRaw(int)
	SetDeviceDataProcessing(int)
	SetDeviceDataProcessingMod(int)
	SetDeviceDataProcessingRaw(int)
	SetDeviceFirewall(int)
	SetDeviceFirewallMod(int)
	SetDeviceFirewallRaw(int)*/
	SetMatrixCM(int)
	GetAlias() string
	GetStunCM() int
	GetPhysCM() int
	SetStunCM(int)
	SetPhysCM(int)
	Dumpshock()
	IsConnected() bool
	SetConnection(bool)
	//CheckRunningProgram(string) bool
	GetPhysicalLocation() bool
	SetPhysicalLocation(bool)
	TriggerDataBomb(int)
	//GetInitiative() int
	//SetInitiative(int)

	//DataSpike(TIcon)
}

//NewPlayer -
func NewPlayer(alias string, d string) *TPersona {
	p := TPersona{}
	p.isPlayer = true
	p.name = alias
	p.faction = alias
	p.alias = alias
	p.device = addDevice(d)
	//r := rand.Intn(len(gridList))
	p.grid = *gridList[0].(*TGrid) //временно - должен стартовать из публичной сети
	p.maxMatrixCM = p.device.GetMatrixCM()
	p.matrixCM = p.maxMatrixCM
	p.cybercombatSkill = 4
	p.computerSkill = 6
	p.hackingSkill = 5
	p.softwareSkill = -1
	p.body = 3
	p.willpower = 4
	p.logic = 8
	p.intuition = 6
	p.charisma = 3
	p.id = id
	p.silentMode = true
	p.simSence = "Hot-SIM VR"
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

	objectList = append(objectList, &p)
	id++
	return &p
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

//GetDeviceRating -
func (p *TPersona) GetDeviceRating() int {
	return p.device.deviceRating
}

//GetDeviceSoft -
func (p *TPersona) GetDeviceSoft() *TProgram {
	return p.device.software
}

//GetAttack -
func (p *TPersona) GetAttack() int {
	boost := 0
	if p.CheckRunningProgram("Decryption") {
		boost = 1
	}
	att := p.device.attack + p.device.attackMod + boost
	if att < 0 {
		return 0
	}
	return att
}

//GetAttackMod -
func (p *TPersona) GetAttackMod() int {
	return p.device.attackMod
}

//GetAttackRaw -
func (p *TPersona) GetAttackRaw() int {
	return p.device.attack
}

//SetDeviceAttack -
func (p *TPersona) SetDeviceAttack(newAttack int) {
	p.device.attack = newAttack
} //Возможно не нужен

//SetDeviceAttackMod -
func (p *TPersona) SetDeviceAttackMod(newAttack int) {
	p.device.attackMod = newAttack
}

//SetDeviceAttackRaw -
func (p *TPersona) SetDeviceAttackRaw(newAttack int) {
	p.device.attack = newAttack
}

//GetSleaze -
func (p *TPersona) GetSleaze() int {
	boost := 0
	if p.CheckRunningProgram("Stealth") {
		boost = 1
	}
	slz := p.device.sleaze + p.device.sleazeMod + boost
	if slz < 0 {
		return 0
	}
	return slz
}

//GetSleazeMod -
func (p *TPersona) GetSleazeMod() int {
	return p.device.sleazeMod
}

//GetSleazeRaw -
func (p *TPersona) GetSleazeRaw() int {
	return p.device.sleaze
}

//SetDeviceSleaze -
func (p *TPersona) SetDeviceSleaze(newSleaze int) {
	p.device.sleaze = newSleaze
}

//SetDeviceSleazeMod -
func (p *TPersona) SetDeviceSleazeMod(newSleaze int) {
	p.device.sleazeMod = newSleaze
}

//SetDeviceSleazeRaw -
func (p *TPersona) SetDeviceSleazeRaw(newSleaze int) {
	p.device.sleaze = newSleaze
}

//GetDataProcessing -
func (p *TPersona) GetDataProcessing() int {
	boost := 0
	if p.CheckRunningProgram("Toolbox") {
		boost = 1
	}
	dtp := p.device.dataProcessing + p.GetDataProcessingMod() + boost
	if dtp < 0 {
		return 0
	}
	return dtp
}

//GetDataProcessingMod -
func (p *TPersona) GetDataProcessingMod() int {
	return p.device.dataProcessingMod
}

//GetDataProcessingRaw -
func (p *TPersona) GetDataProcessingRaw() int {
	return p.device.dataProcessing
}

//SetDeviceDataProcessing -
func (p *TPersona) SetDeviceDataProcessing(newDataProcessing int) {
	p.device.dataProcessing = newDataProcessing
}

//SetDeviceDataProcessingMod -
func (p *TPersona) SetDeviceDataProcessingMod(newDataProcessing int) {
	p.device.dataProcessingMod = newDataProcessing
}

//SetDeviceDataProcessingRaw -
func (p *TPersona) SetDeviceDataProcessingRaw(newDataProcessing int) {
	p.device.dataProcessing = newDataProcessing
}

//GetFirewall -
func (p *TPersona) GetFirewall() int {
	boost := 0
	if p.CheckRunningProgram("Encryption") {
		boost = 1
	}
	fwl := p.device.firewall + p.device.firewallMod + boost
	if fwl < 0 {
		return 0
	}
	return fwl
}

//GetFirewallMod -
func (p *TPersona) GetFirewallMod() int {
	return p.device.firewallMod
}

//GetFirewallRaw -
func (p *TPersona) GetFirewallRaw() int {
	return p.device.firewall
}

//SetDeviceFirewall -
func (p *TPersona) SetDeviceFirewall(newFirewall int) {
	p.device.firewall = newFirewall
}

//SetDeviceFirewallMod -
func (p *TPersona) SetDeviceFirewallMod(newFirewallMod int) {
	p.device.firewallMod = p.device.firewallMod + newFirewallMod
}

//SetDeviceFirewallRaw -
func (p *TPersona) SetDeviceFirewallRaw(newFirewall int) {
	p.device.firewall = newFirewall
}

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

//SetStunCM -
func (p *TPersona) SetStunCM(cmValue int) {
	p.stunCM = cmValue
}

//GetPhysCM -
func (p *TPersona) GetPhysCM() int {
	return p.physCM
}

//SetPhysCM -
func (p *TPersona) SetPhysCM(cmValue int) {
	p.physCM = cmValue
}

//GetGrid -
func (p *TPersona) GetGrid() TGrid {
	return p.grid
}

//SetGrid -
func (p *TPersona) SetGrid(grid TGrid) {
	p.grid = grid
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
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Mark on obj"+strconv.Itoa(i)+" = "+strconv.Itoa(p.markSet.MarksFrom[i]), congo.ColorYellow)
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
		for j := range objectList {
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

		/*if i == 2 {
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
		for j := range objectList {
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
		p.SetDeviceAttackRaw(0)
		p.SetDeviceSleazeRaw(0)
		p.SetDeviceDataProcessingRaw(0)
		p.SetDeviceFirewall(0)
		prgs := p.GetRunningPrograms()
		for i := 0; i < len(prgs); i++ {
			p.CrashProgram(prgs[i])
		}
	}
	dp1 := p.GetWillpower() + p.GetFirewall()
	suc1, gl, cgl := simpleTest(dp1, 1000, 0)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Warning!! Dumpshock imminent!!", congo.ColorRed)
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
	if p.GetSimSence() == "Cold-SIM VR" {
		p.SetStunCM(p.GetStunCM() - biofeedbackDamage)
		congo.WindowsMap.ByTitle["Log"].WPrintLn(strconv.Itoa(biofeedbackDamage)+" Stun Damage inflicted by Dumpshock...", congo.ColorRed)
		if p.GetStunCM() < 0 {
			physDamage := p.GetStunCM() / -2
			p.SetPhysCM(p.GetPhysCM() - physDamage)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(strconv.Itoa(physDamage)+" Physical Damage inflicted by Dumpshock...", congo.ColorRed)
		}
	} else if p.GetSimSence() == "Hot-SIM VR" {
		p.SetPhysCM(p.GetPhysCM() - biofeedbackDamage)
		congo.WindowsMap.ByTitle["Log"].WPrintLn(strconv.Itoa(biofeedbackDamage)+" Physical Damage inflicted by Dumpshock...", congo.ColorRed)

	}
	p.SetSimSence("Offline")
	p.SetInitiative(999999)
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

//LoadProgram -
func (p *TPersona) LoadProgram(name string) bool {
	for i := 0; i < len(p.device.software.programName); i++ {
		if p.device.software.programName[i] == name {
			if p.device.GetRunningProgramsQty() < p.device.GetMaxRunningPrograms() { //тест проверки на то может ли загрузиться данная программа
				if p.device.software.programStatus[i] == "inStore" {
					p.device.software.programStatus[i] = "Running"
				} else {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: Program '"+name+"' is "+p.device.software.programStatus[i], congo.ColorYellow)
				}
				p.device.software.programStatus[i] = "Running"
			} else {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Error: No free program slots available", congo.ColorYellow)
				return false
			}
		}
	}
	return true
}

//UnloadProgram -
func (p *TPersona) UnloadProgram(name string) bool {
	for i := 0; i < len(p.device.software.programName); i++ {
		if p.device.software.programName[i] == name {
			if p.device.software.programStatus[i] == "Running" {
				p.device.software.programStatus[i] = "inStore"
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
					congo.WindowsMap.ByTitle["Log"].WPrintLn("WARNING! "+p.device.software.programName[i]+" program crased!", congo.ColorYellow)
				}
			}
		}
	}
	return true
}

//CheckRunningProgram -
func (p *TPersona) CheckRunningProgram(name string) bool {
	for i := range p.device.software.programName {
		if p.device.software.programName[i] == name {
			if p.device.software.programStatus[i] == "Running" {
				return true
			}
		}
	}
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
			congo.WindowsMap.ByTitle["Log"].WPrintLn("WARNING! 1 additional Matrix Damage caused by Virtual Machine program", congo.ColorYellow)
			hold()
		}
	}
	p.SetMatrixCM(p.GetMatrixCM() - damage)
	if p.GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+" takes "+strconv.Itoa(damage)+" Matrix damage", congo.ColorYellow)
		hold()
	}
	if p.GetMatrixCM() < 1 {
		p.Dumpshock()
	}
}

//ReceivePhysBiofeedbackDamage -
func (p *TPersona) ReceivePhysBiofeedbackDamage(damage int) {
	p.SetPhysCM(p.GetPhysCM() - damage)
	if p.GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+" takes "+strconv.Itoa(damage)+" Physical damage", congo.ColorYellow)
		hold()
	}
}

//ReceiveStunBiofeedbackDamage -
func (p *TPersona) ReceiveStunBiofeedbackDamage(damage int) {
	p.SetStunCM(p.GetStunCM() - damage)
	if p.GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+" takes "+strconv.Itoa(damage)+" Stun damage", congo.ColorYellow)
		hold()
	}
}

//ResistMatrixDamage -
func (p *TPersona) ResistMatrixDamage(damage int) int {
	resistDicePool := 0
	resistDicePool = resistDicePool + p.GetDeviceRating()
	resistDicePool = resistDicePool + p.GetFirewall()
	if p.CheckRunningProgram("Shell") {
		resistDicePool = resistDicePool + 1
		if p.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if p.CheckRunningProgram("Armor") {
		resistDicePool = resistDicePool + 2
		if p.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program Armor: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	damageSoak, gl, cgl := simpleTest(resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if p.GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+": "+strconv.Itoa(damageSoak)+" Matrix damage soaked...", congo.ColorGreen)
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
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program Shell: add 1 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	if p.CheckRunningProgram("Biofeedback Filter") {
		resistDicePool = resistDicePool + 2
		if p.GetFaction() == player.GetFaction() {
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Program Biofeedback Filter: add 2 to resistPool - DEBUG", congo.ColorDefault)
		}
	}
	damageSoak, gl, cgl := simpleTest(resistDicePool, 1000, 0)
	realDamage := damage - damageSoak
	if gl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+": Firewall glitch detected!", congo.ColorYellow)
			hold()
		}
	}
	if cgl {
		if p.GetFaction() == player.GetFaction() {
			realDamage = realDamage + 2
			addOverwatchScoreToTarget(40)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+": Firewall critical failure!", congo.ColorRed)
			hold()
		}
	}
	if realDamage < 0 {
		realDamage = 0
	}

	if p.GetFaction() == player.GetFaction() {
		congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+": "+strconv.Itoa(damageSoak)+" Biofeedback damage soaked...", congo.ColorGreen)
		hold()
	}
	return realDamage
}

func (p *TPersona) checkConvergence() {
	if p.grid.GetOverwatchScore() > 39 {
		if p.convergenceFlag == false {
			p.convergenceFlag = true
			if p.GetFaction() == player.GetFaction() {
				p.ToggleConvergence()
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Warning!!! Convergence protocol engaged...", congo.ColorRed)
				hold()
			}
			convergenceDamage := p.ResistMatrixDamage(12)
			p.ReceiveMatrixDamage(convergenceDamage)
			if p.GetHost() == Matrix {
				p.ToggleConnection()
				p.ClearMarks()

				if p.GetMatrixCM() > 0 {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("1", congo.ColorDefault)
					p.Dumpshock()
				}
				if p.GetFaction() == player.GetFaction() {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("All MARKs have been lost...", congo.ColorRed)
					hold()
				}
			} else {
				host := p.GetHost()
				host.markSet.MarksFrom[p.GetID()] = 3
				if p.GetFaction() == player.GetFaction() {
					congo.WindowsMap.ByTitle["Log"].WPrintLn(host.GetName()+" "+host.GetType()+" now have 3 MARKs on "+p.GetName()+"...", congo.ColorRed)
					hold()
				}
			}
			p.SetPhysicalLocation(true)
			if p.GetFaction() == player.GetFaction() {
				congo.WindowsMap.ByTitle["Log"].WPrintLn("Physical location discovered...", congo.ColorYellow)
				hold()
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...reporting physical location to nearest police force", congo.ColorGreen)
				hold()
				congo.WindowsMap.ByTitle["Log"].WPrintLn("...calling nearest GOD agent", congo.ColorGreen)
				hold()
				if p.GetName() == player.GetName() {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("Grid Overwatch Division wishes you to have a nice day!", congo.ColorDefault)
				}
				hold()
			}
		} else {
			if p.GetHost() == Matrix && p.GetMatrixCM() > 0 && p.GetInitiative() < 9999 {
				p.ToggleConnection()
				p.ClearMarks()
				p.Dumpshock()
				if p.GetFaction() == player.GetFaction() {
					congo.WindowsMap.ByTitle["Log"].WPrintLn("All MARKs have been lost...", congo.ColorRed)
					hold()
				}
			}
		}

	}
}

//TriggerDataBomb -
func (p *TPersona) TriggerDataBomb(bombRating int) {
	host := p.GetHost()
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb triggered...", congo.ColorRed)
	hold()
	if host.GetHostAlertStatus() != "Active Alert" {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Host Active Alert triggered...", congo.ColorRed)
		hold()
		host.SetAlert("Active Alert")
	}
	host.SetAlert("Active Alert")
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
	resistHits, rgl, rcgl := simpleTest(resistPool, 999, 0)
	//остановиться и перебросить при необходимости

	fullDamage := xd6Test(bombRating)
	if rgl == true {
		fullDamage = fullDamage + bombRating
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Warning!! Firewall error erupted...", congo.ColorYellow)
	}
	if rcgl == true {
		//addOverwatchScore(xd6Test(trg.GetDataBombRating()))
		fullDamage = fullDamage + xd6Test(bombRating)
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Danger!! Critical error erupted...", congo.ColorRed)
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn(strconv.Itoa(resistHits)+" of incomming Matrix damage has been resisted", congo.ColorGreen)
	realDamage := fullDamage - resistHits
	if realDamage < 0 {
		realDamage = 0
	}
	p.ReceiveMatrixDamage(realDamage)
	//src.(*TPersona).SetMatrixCM(src.(*TPersona).GetMatrixCM() - realDamage)
	//congo.WindowsMap.ByTitle["Log"].WPrintLn(p.GetName()+" receive "+strconv.Itoa(realDamage)+" of matrix damage", congo.ColorYellow)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Databomb destroyed", congo.ColorGreen)
	hold()
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

///////////////////////////////////////////////////////
//File

//TFile -
type TFile struct {
	owner    string //*TPersona
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
	IIcon
	//IObj
	//GetOwner() string
	SetOwner(string)
	//GetName() string
	SetName(string)
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

//NewFile -
func (h *THost) NewFile(name string) *TFile {
	f := TFile{}
	data := player.canSee.KnownData[f.id]
	data[0] = "Unknown"
	data[1] = "Unknown"
	data[3] = "Unknown"
	data[12] = "Unknown"
	data[13] = "Unknown"
	data[15] = "Unknown"
	f.owner = h.GetName()
	f.host = h
	f.grid = h.grid
	f.device = addDevice("noDevice")
	f.name = f.GetType() + " " + strconv.Itoa(id)
	enRat, _, _ := simpleTest(h.deviceRating+h.deviceRating, h.dataProcessing, 0)
	f.encryptionRating = enRat
	bombRat, _, _ := simpleTest(h.deviceRating+h.deviceRating, h.sleaze, 0)
	f.dataBombRating = bombRat
	f.lastEditTime = generateLastEditTime()

	hRatMod := (h.deviceRating-1)/3 + 1

	//file := NewFile("File")
	dataDensity := xd6Test(2)
	f.size = hRatMod * 5 * dataDensity

	f.id = id
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
	player.canSee.KnownData[f.id] = data
	//f.grid = "ARES GRID"
	//f.name = f.GetType() + " " + strconv.Itoa(id)
	objectList = append(objectList, &f)
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
func (f *TFile) GetOwner() string {
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
func (f *TFile) SetOwner(owner string) {
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
func (f *TFile) GetGrid() TGrid {
	return f.grid
}

//SetGrid -
func (f *TFile) SetGrid(grid TGrid) {
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
	if f.device.deviceType == "noDevice" {
		return f.host.dataProcessing
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
