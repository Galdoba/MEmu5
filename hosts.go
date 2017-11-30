package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/Galdoba/ConGo/congo"
	"github.com/Galdoba/ConGo/utils"
)

//ICList -
type ICList struct {
	icName   []string
	icStatus []bool
	icID     []int //////////
}

//THost -
type THost struct {
	TGrid
	attack         int
	sleaze         int
	dataProcessing int
	firewall       int
	id             int
	matrixCM       int
	markSet        MarkSet
	grid           *TGrid
	//icOrder        []string
	icState            ICList
	alert              string
	owner              IIcon
	freeActionsCount   int
	simpleActionsCount int
}

//IHost - в икону входят файлы, персоны, айсы и хосты
type IHost interface {
	IObj
	IGridOnly
	IIconOnly
	LoadNextIC() bool
	//SetID()
	GetHostAlertStatus() string
	GatherMarks()
	//GiveMarks()
	SetAlert(string)
	DeleteIC(IIC) bool
	DeleteFile(IFile) bool
	PickPatrolIC() IIC
	GetICState() ICList
}

var _ IHost = (*THost)(nil)

//RollInitiative -
func (h *THost) RollInitiative() {
	h.SetInitiative(h.GetDataProcessing() + h.GetIntuition() + xd6Test(4))
}

//GetFreeActionsCount -
func (h *THost) GetFreeActionsCount() int {
	return h.freeActionsCount
}

//GetSimpleActionsCount -
func (h *THost) GetSimpleActionsCount() int {
	return h.simpleActionsCount
}

//ResetActionsCount -
func (h *THost) ResetActionsCount() {
	h.freeActionsCount = 1
	//
	h.simpleActionsCount = 2
}

//SpendFreeAction -
func (h *THost) SpendFreeAction() {
	if h.freeActionsCount > 0 {
		h.freeActionsCount--
	} else {
		h.simpleActionsCount--
	}
}

//SpendSimpleAction -
func (h *THost) SpendSimpleAction() {
	h.simpleActionsCount--
}

//SpendComplexAction -
func (h *THost) SpendComplexAction() {
	h.simpleActionsCount--
	h.simpleActionsCount--
}

//GetIntuition -
func (h *THost) GetIntuition() int {
	return h.GetDeviceRating()
}

//GetICState -
func (h *THost) GetICState() ICList {
	return h.icState
}

//CheckRunningProgram -
func (h *THost) CheckRunningProgram(name string) bool {
	return false
}

//GetConvergenceFlag -
func (h *THost) GetConvergenceFlag() bool {
	return false
}

//GetDevice -
func (h *THost) GetDevice() *TDevice {
	return nil
}

//GetHost -
func (h *THost) GetHost() *THost {
	return h
}

//GetInitiative -
func (h *THost) GetInitiative() int {
	return -1
}

//GetLongAct -
func (h *THost) GetLongAct() int {
	return -1
}

//GetOwner -
func (h *THost) GetOwner() IIcon {
	return h.owner
}

//GetSilentRunningMode -
func (h *THost) GetSilentRunningMode() bool {
	return false
}

//GetSimSence -
func (h *THost) GetSimSence() string {
	return "HOT-SIM"
}

//IsPlayer -
func (h *THost) IsPlayer() bool {
	return false
}

//ReceiveMatrixDamage -
func (h *THost) ReceiveMatrixDamage(damage int) {
	printLog("...error: Host is immune to Matrix Damage", congo.ColorYellow)
}

//ResistMatrixDamage -
func (h *THost) ResistMatrixDamage(damage int) int {
	return 0
}

//SetGrid -
func (h *THost) SetGrid(grid *TGrid) {
	//h.lastLocation = h.grid
	//h.grid = grid
}

//SetHost -
func (h *THost) SetHost(host *THost) {
	//h.lastLocation = h.grid
	//h.grid = grid
}

//SetInitiative -
func (h *THost) SetInitiative(int) {
}

//SetSilentRunningMode -
func (h *THost) SetSilentRunningMode(bool) {
}

//SetSimSence -
func (h *THost) SetSimSence(string) {
}

//ToggleConnection -
func (h *THost) ToggleConnection() {
}

//ToggleConvergence -
func (h *THost) ToggleConvergence() bool {
	return false
}

//PickPatrolIC -
func (h *THost) PickPatrolIC() IIC {
	var patrolFound bool
	var try int
	for !patrolFound {
		patroIC := pickIconByName("Patrol IC")
		if patroIC.GetHost() != h {
			try++
		}
		if try > 1000 {
			panic("Patrol IC was not found in 1000 cycles")
		}
		return patroIC.(IIC)
	}
	return nil
}

/*	for i := 0; i < len(h.icState.icName); i++ {
		if h.icState.icName[i] == "Patrol IC" && h.icState.icStatus[i] == true {
			for j := range objectList {
				if patrolIC, ok := objectList[j].(*TIC); ok {
					if patrolIC.GetName() == "Patrol IC" && patrolIC.GetHost().name == h.name {
						return patrolIC
					}
				}
			}
		}
	}
	return nil
}*/

//LoadNextIC -
func (h *THost) LoadNextIC() bool {
	for i := range h.icState.icName {
		//for i := 0; i < h.deviceRating; i++ {
		if h.icState.icStatus[i] == false {
			if player.GetHost() == h {
				printLog(h.icState.icName[i]+" was loaded...", congo.ColorYellow)
			}
			//congo.WindowsMap.ByTitle["Log"].WPrintLn(h.icState.icName[i]+" was loaded...", congo.ColorRed)
			//congo.WindowsMap.ByTitle["Process"].WPrint(".", congo.ColorGreen)
			////ObjByNames[h.icState.icName[i]] =
			h.NewIC(h.icState.icName[i])
			h.icState.icStatus[i] = true
			//ObjByNames[h.icState.icName[i]] = &i
			//h.icState.icID[i] = true

			return true
		}
	}
	return true
}

//DeleteIC -
func (h *THost) DeleteIC(ic IIC) bool {
	//ObjByNames[ic.GetName()] = nil
	//delete(ObjByNames, ic.GetName())
	congo.WindowsMap.ByTitle["Log"].WPrintLn("IC Name= "+ic.GetName(), congo.ColorDefault)
	for i := 0; i < h.deviceRating; i++ {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Check= "+h.icState.icName[i], congo.ColorDefault)
		icName := h.icState.icName[i]
		icObj := ObjByNames[h.icState.icName[i]].(IIC)

		//	for j := range objectList {
		//		if icObj, ok := objectList[j].(*TIC); ok {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Compare= "+icObj.GetName()+" with "+icName, congo.ColorDefault)
		if icObj.GetName() == icName {
			congo.WindowsMap.ByTitle["Log"].WPrintLn(icObj.GetName()+" = ok!", congo.ColorGreen)
			congo.WindowsMap.ByTitle["Log"].WPrintLn(icObj.GetName()+" has "+strconv.Itoa(icObj.GetMatrixCM())+" matrix boxes...", congo.ColorGreen)
			if icObj.GetMatrixCM() < 1 {
				congo.WindowsMap.ByTitle["Log"].WPrintLn(icObj.GetName()+" must be deleted...", congo.ColorGreen)
				h.icState.icStatus[i] = false
				h.SetAlert("Active Alert")
				hold()
				congo.WindowsMap.ByTitle["Log"].WPrintLn(h.GetName()+": Set Active Alert ON", congo.ColorGreen)
				//objectList = append(objectList[:j], objectList[j+1:]...)
				delete(ObjByNames, ic.GetName())
				return true

			}

		}
		//		}
		//	}

		/*congo.WindowsMap.ByTitle["Log"].WPrintLn("Check= "+h.icState.icName[i], congo.ColorDefault)
		if h.icState.icStatus[i] == true && ic.GetName() == h.icState.icName[i] {
			congo.WindowsMap.ByTitle["Log"].WPrintLn(h.icState.icName[i]+" = "+ic.GetName()+"   (h.icState.icStatus[i] == true && ic.GetName() == h.icState.icName[i])", congo.ColorDefault)
			h.icState.icStatus[i] = false
			return true
		}*/
		if h.icState.icStatus[i] == true {
			congo.WindowsMap.ByTitle["Log"].WPrintLn(h.icState.icName[i]+" is loaded", congo.ColorDefault)
		} else {
			congo.WindowsMap.ByTitle["Log"].WPrintLn(h.icState.icName[i]+" is NOT loaded", congo.ColorDefault)
		}
	}
	return true
}

//DeleteFile -
func (h *THost) DeleteFile(file IFile) bool {
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Delete file: "+file.GetName()+"...", congo.ColorGreen)
	hold()
	for _, obj := range ObjByNames {
		if flDel, ok := obj.(IFile); ok {
			congo.WindowsMap.ByTitle["Log"].WPrint(".", congo.ColorGreen)
			hold()
			if flDel.GetName() == file.GetName() {
				delete(ObjByNames, file.GetName())
				return true
			}
		}
	}
	congo.WindowsMap.ByTitle["Log"].WPrintLn("", congo.ColorGreen)
	congo.WindowsMap.ByTitle["Log"].WPrintLn("...error: "+file.GetName()+" was not found", congo.ColorGreen)
	hold()
	return true
}

//GatherMarks -
func (h *THost) GatherMarks() {
	var slaves []int
	var notSlaves []int
	slaves = append(slaves, h.id) //add host
	for _, obj := range ObjByNames {
		if icon, ok := obj.(IIcon); ok {
			if icon.GetOwner() == h {
				slaves = append(slaves, icon.GetID()) //add slave
			} else {
				notSlaves = append(notSlaves, icon.GetID()) //add non-slave
			}
		}
	}
	/*	for _, obj := range ObjByNames{
		if icon, ok := obj.(IIcon); ok {
			if icon.GetOwner() != h {
				notSlaves = append(notSlaves, icon.GetID()) //add non-slave
			}
		}
	}*/

	/*	for s := range objectList {
		if icon, ok := objectList[s].(IIcon); ok {
			if icon.GetOwner() == h {
				slaves = append(slaves, icon.GetID()) //add slave
			}
		}
	}*/
	/*for ns := range objectList {
		if icon, ok := objectList[ns].(IIcon); ok {
			if icon.GetOwner() != h {
				notSlaves = append(notSlaves, icon.GetID()) //add non-slave
			}
		}
	}*/
	for _, obj := range ObjByNames {
		if notSlaveToCheck, ok := obj.(IIcon); ok {
			if notSlaveToCheck.GetOwner() != h {
				markMap := notSlaveToCheck.GetMarkSet()
				//				canSee := notSlaveToCheck.GetFieldOfView()
				for key, value := range markMap.MarksFrom { //check if non-slave marked by slave
					if contains(slaves, key) { //if true
						if key != h.id { //and not host
							if notSlaveToCheck.GetMarkSet().MarksFrom[h.id] < value && value < 4 { //share mark with host
								notSlaveToCheck.GetMarkSet().MarksFrom[h.id] = value
								if h.alert == "No Alert" {
									h.alert = "Passive Alert"
								}
							}
						}
					}
				}
			}
		}
	}

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//NewHost -
func (g *TGrid) NewHost(name string, rating int) *THost {
	if name == "Matrix" {
		h := THost{}
		h.name = "Matrix"
		//h.
		/*for _, obj := range ObjByNames {
			if gridToPick, ok := obj.(*TGrid); ok {
				h.grid = gridToPick
				h.deviceRating = h.grid.deviceRating
				h.dataProcessing = h.grid.deviceRating
			}
		}*/
		h.grid = player.grid
		h.deviceRating = player.grid.deviceRating
		h.attack = player.grid.deviceRating
		h.sleaze = player.grid.deviceRating
		h.dataProcessing = player.grid.deviceRating
		h.firewall = player.grid.deviceRating
		return &h
	}
	h := THost{}
	setSeed()
	h.name = name
	h.faction = name
	h.matrixCM = 999999
	randoGo := true
	for randoGo {
		r := rand.Intn(len(gridList))
		if gr, ok := gridList[r].(*TGrid); ok {
			h.grid = gr
			randoGo = false
		}
	}
	//r := rand.Intn(len(gridList))
	//	h.grid = *gridList[r].(*TGrid)
	//h.overwatchScore = 0
	h.SetID()
	h.markSet.MarksFrom = make(map[int]int)
	h.markSet.MarksFrom[h.id] = 4
	h.canSee.KnownData = make(map[int][30]string)
	if rating > 0 && rating < 13 {
		h.deviceRating = rating
	} else {
		h.deviceRating = rand.Intn(12) + 1
	}
	//costMult := h.deviceRating
	atributeArray := []int{0, 1, 2, 3}
	for i := rand.Intn(100); i > 0; i-- {
		shuffleInt(atributeArray)
	}
	h.attack = atributeArray[0] + h.deviceRating
	h.sleaze = atributeArray[1] + h.deviceRating
	h.dataProcessing = atributeArray[2] + h.deviceRating
	h.firewall = atributeArray[3] + h.deviceRating
	h.alert = "No Alert"
	h.FillHostWithFiles()

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
	data[18] = "Unknown"
	player.canSee.KnownData[h.id] = data

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

	icList := new(ICList)
	h.icState = *icList
	for i := 0; i < h.deviceRating; i++ {
		if i == 0 {
			n := 10 // 10 - is patrol's position in "allIC[]"
			h.icState.icName = append(h.icState.icName, allIC[n])
			h.icState.icStatus = append(h.icState.icStatus, false)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Choose= "+allIC[n], congo.ColorRed)
			if len(allIC) > 1 {
				allIC = append(allIC[:n], allIC[n+1:]...)

			}
		} else {
			setSeed()
			n := rand.Intn(len(allIC))
			h.icState.icName = append(h.icState.icName, allIC[n])
			h.icState.icStatus = append(h.icState.icStatus, false)
			//congo.WindowsMap.ByTitle["Log"].WPrintLn("Choose= "+allIC[n], congo.ColorYellow)
			if len(allIC) > 1 {
				allIC = append(allIC[:n], allIC[n+1:]...)
			}
		}
	}

	//windowList[0].(*congo.TWindow).WPrintLn("Host located:", congo.ColorGreen)
	//windowList[0].(*congo.TWindow).WPrintLn("Grid: "+h.grid.name, congo.ColorGreen)
	h.LoadNextIC()
	//	windowList[0].(*congo.TWindow).WPrintLn(h.name+" located.", congo.ColorGreen)
	//	windowList[0].(*congo.TWindow).WPrintLn("//Debug: Atribute Array:", congo.ColorYellow)
	//	windowList[0].(*congo.TWindow).WPrintLn(strconv.Itoa(h.deviceRating)+" "+strconv.Itoa(h.attack)+" "+strconv.Itoa(h.sleaze)+" "+strconv.Itoa(h.dataProcessing)+" "+strconv.Itoa(h.firewall), congo.ColorYellow)
	//objectList = append(objectList, &h)
	h.owner = &h
	gridList = append(gridList, &h)
	ObjByNames[h.name] = &h
	//	windowList[0].(*congo.TWindow).WPrintLn(h.HostToString(), congo.ColorYellow)
	h.HostToString()
	return &h
}

// GetType -
func (h *THost) GetType() string {
	return "Host"
}

//GetGrid -
func (h *THost) GetGrid() *TGrid {
	return h.grid
}

//GetGridName -
func (h *THost) GetGridName() string {
	return h.grid.GetGridName()
}

//GetFaction -
func (h *THost) GetFaction() string {
	return h.faction
}

//SetID -
func (h *THost) SetID() {
	h.id = id
	id = id + xd6Test(3)
}

//GetID -
func (h *THost) GetID() int {
	return h.id
}

//GetDeviceRating -
func (h *THost) GetDeviceRating() int {
	return h.deviceRating
}

//GetAttack -
func (h *THost) GetAttack() int {
	return h.attack
}

//GetSleaze -
func (h *THost) GetSleaze() int {
	return h.sleaze
}

//GetDataProcessing -
func (h *THost) GetDataProcessing() int {
	return h.dataProcessing
}

//GetFirewall -
func (h *THost) GetFirewall() int {
	return h.firewall
}

//GetName -
func (h *THost) GetName() string {
	if h.name == "" {
		return "Matrix"
	}
	return h.name
}

//GetHostAlertStatus -
func (h *THost) GetHostAlertStatus() string {
	return h.alert
}

func shuffleInt(atributeArray []int) {
	for i := len(atributeArray) - 1; i > 0; i-- {
		j := rand.Intn(i)
		atributeArray[i], atributeArray[j] = atributeArray[j], atributeArray[i]
	}
}

func shuffleString(atributeArray []string) {
	for i := len(atributeArray) - 1; i > 0; i-- {
		j := rand.Intn(i)
		atributeArray[i], atributeArray[j] = atributeArray[j], atributeArray[i]
	}
}

//GetMarkSet -
func (h *THost) GetMarkSet() MarkSet {
	return h.markSet
}

//SetAlert -
func (h *THost) SetAlert(newAlert string) {
	notify := false
	if h.alert != "Active Alert" {
		notify = true
	}
	h.alert = newAlert
	if player.GetHost() == h && h.name != "Matrix" {
		if h.alert == "Passive Alert" {
			printLog("...Host now in Passive Alert mode!", congo.ColorYellow)
		}
		if h.alert == "Active Alert" && notify == true {
			printLog("...Host now in Active Alert mode!", congo.ColorRed)
			printLog("SYSTEM MESSAGE:", congo.ColorDefault)
			printLog("Attention all users!", congo.ColorDefault)
			printLog("IC Activation Protocol engaged.", congo.ColorDefault)
			printLog("Please terminate all operations and sign off.", congo.ColorDefault)
			printLog("We deeply regret any inconvenience we may have caused.", congo.ColorDefault)
		}
	}
}

//GetLinkLockStatus -
func (h *THost) GetLinkLockStatus() Locked {
	//panic("Abs Func Call")
	return h.linklocked
}

//LockIcon -
func (h *THost) LockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[h.id] = true
}

//UnlockIcon -
func (h *THost) UnlockIcon(icon IIcon) {
	icon.GetLinkLockStatus().LockedByID[h.id] = false
}

//HostToString -
func (h *THost) HostToString() string {
	hString := "          \r\n"
	hString = hString + "Host: " + h.name + "\r\n"
	hString = hString + "Grid: " + h.grid.name + "\r\n"
	hString = hString + "Host_Atributes:" + "\r\n"
	hString = hString + "  Rating: " + strconv.Itoa(h.deviceRating) + "\r\n"
	hString = hString + "  Attack: " + strconv.Itoa(h.attack) + "\r\n"
	hString = hString + "  Sleaze: " + strconv.Itoa(h.sleaze) + "\r\n"
	hString = hString + "  Data Processing: " + strconv.Itoa(h.dataProcessing) + "\r\n"
	hString = hString + "  Firewall: " + strconv.Itoa(h.firewall) + "\r\n"
	hString = hString + "Host_IC:" + "\r\n"

	for i := 0; i < h.deviceRating; i++ {
		hString = hString + " >" + h.icState.icName[i] + "\r\n"
	}
	hString = hString + "\r\n"
	hString = hString + "#########################\r\n" // Конец хоста (сепаратор)

	file, err := os.OpenFile("HostDB.txt", os.O_APPEND|os.O_WRONLY|os.O_RDWR, 0600) // открываем файл: Имя, ключи, что-то еще
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //закрываем файл когда он уже не нужен
	file.WriteString(hString)

	return hString
}

//HostExist -
func HostExist(hostName string) bool {
	file, err := os.OpenFile("HostDB.txt", os.O_APPEND|os.O_WRONLY|os.O_RDWR, 0600) // открываем файл: Имя, ключи, что-то еще
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()                         //закрываем файл когда он уже не нужен
	data, err := ioutil.ReadFile("HostDB.txt") //читаем файл
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)
	cont := strings.Split(content, "\r\n")
	subStr := "Host: " + hostName

	for i := 0; i < len(cont); i++ {
		if strings.Contains(cont[i], subStr) {
			congo.WindowsMap.ByTitle["Process"].WPrintLn("Connection with host established:", congo.ColorDefault)
			return true
		}
	}
	congo.WindowsMap.ByTitle["Process"].WPrintLn("Lost connection with host", congo.ColorDefault)
	return false
}

//GenerateFileData -
func (h *THost) GenerateFileData() {

}

//FillHostWithFiles -
func (h *THost) FillHostWithFiles() {
	totalFiles := xd6Test(h.deviceRating)
	for i := 0; i < totalFiles; i++ {
		costMult := 0
		//file := h.NewFile("File" + " " + strconv.Itoa(id))
		file := h.NewFile("random")

		r, _, _ := simpleTest(h.GetID(), h.deviceRating+h.deviceRating, h.dataProcessing, 0)
		file.SetEncryptionRating(r)
		if r > 0 {
			bR, _, _ := simpleTest(h.GetID(), h.deviceRating+h.deviceRating, h.sleaze, 0)
			file.SetDataBombRating(bR)
			costMult = r * bR
		}
		hRat := h.deviceRating/3 + 1
		//file := NewFile("File")
		file.host = h
		dataDensity := xd6Test(2)
		file.size = hRat * 5 * dataDensity
		//value
		//value := 0
		switch hRat {
		case 1:
			value := utils.Max(xd6Test(1)-1, 0)
			file.value = value * 5 * file.size * costMult
		case 2:
			value := utils.Max(xd6Test(2)-2, 0)
			file.value = value * 5 * file.size * costMult
		case 3:
			value := utils.Max(xd6Test(2), 0)
			file.value = value * 5 * file.size * costMult
		case 4:
			value := utils.Max(xd6Test(2)+2, 0)
			file.value = value * 5 * file.size * costMult
		default:
		}
	}
}

func applyGrid(name string) *TGrid {
	for i := range gridList {
		if grid, ok := gridList[i].(*TGrid); ok {
			windowList[0].(*congo.TWindow).WPrintLn(grid.GetGridName(), congo.ColorGreen)
			if grid.GetGridName() == name {
				return grid
			}
		}
	}
	return player.grid
}
