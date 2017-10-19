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
	grid           TGrid
	//icOrder        []string
	icState ICList
	alert   string
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
	DeleteIC(*TIC) bool
	DeleteFile(*TFile) bool
	PickPatrolIC() *TIC
}

var _ IHost = (*THost)(nil)

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
func (h *THost) GetOwner() string {
	return h.name
}

//GetSilentRunningMode -
func (h *THost) GetSilentRunningMode() bool {
	return false
}

//GetSimSence -
func (h *THost) GetSimSence() string {
	return "Hot-SIM VR"
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
func (h *THost) SetGrid(grid TGrid) {
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
func (h *THost) PickPatrolIC() *TIC {
	for i := 0; i < len(h.icState.icName); i++ {
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
}

//LoadNextIC -
func (h *THost) LoadNextIC() bool {
	for i := range h.icState.icName {
		//for i := 0; i < h.deviceRating; i++ {
		if h.icState.icStatus[i] == false {
			congo.WindowsMap.ByTitle["Log"].WPrintLn(h.icState.icName[i]+" was loaded...", congo.ColorRed)
			h.NewIC(h.icState.icName[i])
			h.icState.icStatus[i] = true
			//h.icState.icID[i] = true

			return true
		}
	}
	return true
}

//DeleteIC -
func (h *THost) DeleteIC(ic *TIC) bool {
	congo.WindowsMap.ByTitle["Log"].WPrintLn("IC Name= "+ic.GetName(), congo.ColorDefault)
	for i := 0; i < h.deviceRating; i++ {
		congo.WindowsMap.ByTitle["Log"].WPrintLn("Check= "+h.icState.icName[i], congo.ColorDefault)
		icName := h.icState.icName[i]
		for j := range objectList {
			if icObj, ok := objectList[j].(*TIC); ok {
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
						objectList = append(objectList[:j], objectList[j+1:]...)
						return true

					}

				}
			}
		}

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
func (h *THost) DeleteFile(file *TFile) bool {
	congo.WindowsMap.ByTitle["Log"].WPrintLn("Delete file: "+file.GetName()+"...", congo.ColorGreen)
	hold()
	for i := range objectList {
		if ftDel, ok := objectList[i].(*TFile); ok {
			congo.WindowsMap.ByTitle["Log"].WPrint(".", congo.ColorGreen)
			hold()
			if ftDel.GetName() == file.GetName() {
				congo.WindowsMap.ByTitle["Log"].WPrintLn(".."+ftDel.GetName()+" deleted", congo.ColorGreen)
				hold()
				objectList = append(objectList[:i], objectList[i+1:]...)
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
	for s := range objectList {
		if icon, ok := objectList[s].(IIcon); ok {
			if icon.GetOwner() == h.GetName() {
				slaves = append(slaves, icon.GetID()) //add slave
			}
		}
	}
	for ns := range objectList {
		if icon, ok := objectList[ns].(IIcon); ok {
			if icon.GetOwner() != h.GetName() {
				notSlaves = append(notSlaves, icon.GetID()) //add non-slave
			}
		}
	}
	for i := range objectList {
		if notSlaveToCheck, ok := objectList[i].(IIcon); ok {
			if notSlaveToCheck.GetOwner() != h.GetName() {
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
		//h.grid =
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
			h.grid = *gr
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
			congo.WindowsMap.ByTitle["Log"].WPrintLn("Choose= "+allIC[n], congo.ColorYellow)
			if len(allIC) > 1 {
				allIC = append(allIC[:n], allIC[n+1:]...)
			}
		}
	}

	//windowList[0].(*congo.TWindow).WPrintLn("Host located:", congo.ColorGreen)
	//windowList[0].(*congo.TWindow).WPrintLn("Grid: "+h.grid.name, congo.ColorGreen)
	h.LoadNextIC()
	windowList[0].(*congo.TWindow).WPrintLn(h.name+" located.", congo.ColorGreen)
	windowList[0].(*congo.TWindow).WPrintLn("//Debug: Atribute Array:", congo.ColorYellow)
	windowList[0].(*congo.TWindow).WPrintLn(strconv.Itoa(h.deviceRating)+" "+strconv.Itoa(h.attack)+" "+strconv.Itoa(h.sleaze)+" "+strconv.Itoa(h.dataProcessing)+" "+strconv.Itoa(h.firewall), congo.ColorYellow)
	//objectList = append(objectList, &h)
	gridList = append(gridList, &h)
	ObjByNames[h.name] = &h
	windowList[0].(*congo.TWindow).WPrintLn(h.HostToString(), congo.ColorYellow)
	return &h
}

// GetType -
func (h *THost) GetType() string {
	return "Host"
}

//GetGrid -
func (h *THost) GetGrid() TGrid {
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
	h.alert = newAlert
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
	hString := "          \n"
	hString = hString + "Host: " + h.name + "\n"
	hString = hString + "Grid: " + h.grid.name + "\n"
	hString = hString + "Host_Atributes:" + "\n"
	hString = hString + "  Rating: " + strconv.Itoa(h.deviceRating) + "\n"
	hString = hString + "  Attack: " + strconv.Itoa(h.attack) + "\n"
	hString = hString + "  Sleaze: " + strconv.Itoa(h.sleaze) + "\n"
	hString = hString + "  Data Processing: " + strconv.Itoa(h.dataProcessing) + "\n"
	hString = hString + "  Firewall: " + strconv.Itoa(h.firewall) + "\n"
	hString = hString + "Host_IC:" + "\n"

	for i := 0; i < h.deviceRating; i++ {
		hString = hString + " >" + h.icState.icName[i] + "\n"
	}
	hString = hString + "\n"
	hString = hString + "#########################\n" // Конец хоста (сепаратор)

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
	cont := strings.Split(content, "\n")
	windowList[0].(*congo.TWindow).WPrintLn("Checking "+hostName+"...", congo.ColorYellow)
	subStr := "Host: " + hostName

	for i := 0; i < len(cont); i++ {
		if strings.Contains(cont[i], subStr) {
			windowList[0].(*congo.TWindow).WPrintLn(hostName+" exist !", congo.ColorYellow)
			return true
		}
	}
	windowList[0].(*congo.TWindow).WPrintLn(hostName+" is not exist", congo.ColorYellow)
	return false
}

//ImportHostFromDB -
func ImportHostFromDB(hostName string) *THost {
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
						h.grid = *gridForHost
						randoGo = false
					}
				}
			}
			//Если сеть не известна - выбираем из известных
			for randoGo {
				r := rand.Intn(len(gridList))
				if gr, ok := gridList[r].(*TGrid); ok {
					h.grid = *gr
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

		r, _, _ := simpleTest(h.deviceRating+h.deviceRating, h.dataProcessing, 0)
		file.SetEncryptionRating(r)
		if r > 0 {
			bR, _, _ := simpleTest(h.deviceRating+h.deviceRating, h.sleaze, 0)
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

func applyGrid(name string) TGrid {
	for i := range gridList {
		if grid, ok := gridList[i].(*TGrid); ok {
			windowList[0].(*congo.TWindow).WPrintLn(grid.GetGridName(), congo.ColorGreen)
			if grid.GetGridName() == name {
				return *grid
			}
		}
	}
	return player.grid
}
