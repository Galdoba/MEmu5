package main

//DB -
var DB TDeviceDB

//TDeviceDB -
type TDeviceDB struct {
	DeviceDB map[string]*TDevice
}

//InitDeviceDatabase -
func InitDeviceDatabase() {
	DB = TDeviceDB{}
	DB.DeviceDB = map[string]*TDevice{}
}

//AddDevice -
func AddDevice() TDeviceDB {
	DB.DeviceDB["Camera3"] = &TDevice{}
	DB.DeviceDB["Camera3"].SetDataProcessing(3)
	////////////////////////
	DB.DeviceDB["Camera5"] = &TDevice{}
	DB.DeviceDB["Camera5"].deviceRating = 5
	//////////////
	DB.DeviceDB["Erika MCD-1"] = &TDevice{}
	DB.DeviceDB["Erika MCD-1"].deviceType = "Cyberdeck"
	DB.DeviceDB["Erika MCD-1"].model = "Erika MCD-1"
	DB.DeviceDB["Erika MCD-1"].deviceRating = 1
	DB.DeviceDB["Erika MCD-1"].attack = 4
	DB.DeviceDB["Erika MCD-1"].sleaze = 3
	DB.DeviceDB["Erika MCD-1"].dataProcessing = 2
	DB.DeviceDB["Erika MCD-1"].firewall = 1
	DB.DeviceDB["Erika MCD-1"].maxMatrixCM = 9
	DB.DeviceDB["Erika MCD-1"].matrixCM = 9
	DB.DeviceDB["Erika MCD-1"].canSwapAtt = true
	DB.DeviceDB["Erika MCD-1"].software = preaparePrograms()
	//////////////
	DB.DeviceDB["Sony CIY-720"] = &TDevice{}
	DB.DeviceDB["Sony CIY-720"].deviceType = "Cyberdeck"
	DB.DeviceDB["Sony CIY-720"].model = "Sony CIY-720"
	DB.DeviceDB["Sony CIY-720"].deviceRating = 4
	DB.DeviceDB["Sony CIY-720"].attack = 7
	DB.DeviceDB["Sony CIY-720"].sleaze = 6
	DB.DeviceDB["Sony CIY-720"].dataProcessing = 5
	DB.DeviceDB["Sony CIY-720"].firewall = 4
	DB.DeviceDB["Sony CIY-720"].maxMatrixCM = 10
	DB.DeviceDB["Sony CIY-720"].matrixCM = 10
	DB.DeviceDB["Sony CIY-720"].canSwapAtt = true
	DB.DeviceDB["Sony CIY-720"].software = preaparePrograms()
	DB.DeviceDB["Sony CIY-720"].deviceType = "Cyberdeck"
	DB.DeviceDB["Sony CIY-720"].deviceType = "Cyberdeck"
	DB.DeviceDB["Sony CIY-720"].deviceType = "Cyberdeck"
	//////////////
	DB.DeviceDB["noDevice"] = &TDevice{}
	DB.DeviceDB["noDevice"].deviceType = "noDevice"
	DB.DeviceDB["noDevice"].software = preaparePrograms()

	/*	DB.DeviceDB["noDevice"].software.programName = append(DB.DeviceDB["noDevice"].software.programName, "Browse")
		DB.DeviceDB["noDevice"].software.programStatus = append(DB.DeviceDB["noDevice"].software.programStatus, "inStore")
		DB.DeviceDB["noDevice"].software.programType = append(DB.DeviceDB["noDevice"].software.programType, "COMMON")*/
	//DB.DeviceDB["noDevice"].software.programName[0] = "c"
	//DB.DeviceDB["noDevice"].programs = 1
	return DB
}

func preaparePrograms() *TProgram {
	prgList := new(TProgram)
	//Browse
	prgList.programName = append(prgList.programName, "Browse")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Confgurator
	prgList.programName = append(prgList.programName, "Confgurator")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Edit
	prgList.programName = append(prgList.programName, "Edit")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Encryption
	prgList.programName = append(prgList.programName, "Encryption")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Signal Scrub
	prgList.programName = append(prgList.programName, "Signal Scrub")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Toolbox
	prgList.programName = append(prgList.programName, "Toolbox")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Virtual Machine
	prgList.programName = append(prgList.programName, "Virtual Machine")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Bootstrap
	prgList.programName = append(prgList.programName, "Bootstrap")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Search
	prgList.programName = append(prgList.programName, "Search")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Shredder
	prgList.programName = append(prgList.programName, "Shredder")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "COMMON")
	//Armor
	prgList.programName = append(prgList.programName, "Armor")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Baby Monitor
	prgList.programName = append(prgList.programName, "Baby Monitor")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Biofeedback
	prgList.programName = append(prgList.programName, "Biofeedback")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Biofeedback Filter
	prgList.programName = append(prgList.programName, "Biofeedback Filter")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Blackout
	prgList.programName = append(prgList.programName, "Blackout")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Decryption
	prgList.programName = append(prgList.programName, "Decryption")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Defuse
	prgList.programName = append(prgList.programName, "Defuse")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Demolition
	prgList.programName = append(prgList.programName, "Demolition")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Exploit
	prgList.programName = append(prgList.programName, "Exploit")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Fork
	prgList.programName = append(prgList.programName, "Fork")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Guard
	prgList.programName = append(prgList.programName, "Guard")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Hammer
	prgList.programName = append(prgList.programName, "Hammer")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Lockdown
	prgList.programName = append(prgList.programName, "Lockdown")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Mugger
	prgList.programName = append(prgList.programName, "Mugger")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Shell
	prgList.programName = append(prgList.programName, "Shell")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Sneak
	prgList.programName = append(prgList.programName, "Sneak")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Stealth
	prgList.programName = append(prgList.programName, "Stealth")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Track
	prgList.programName = append(prgList.programName, "Track")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Wrapper
	prgList.programName = append(prgList.programName, "Wrapper")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Cat’s Paw
	prgList.programName = append(prgList.programName, "Cat’s Paw")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Cloudless
	prgList.programName = append(prgList.programName, "Cloudless")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Detonator
	prgList.programName = append(prgList.programName, "Detonator")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Evaluate
	prgList.programName = append(prgList.programName, "Evaluate")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Fly on a Wall
	prgList.programName = append(prgList.programName, "Fly on a Wall")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Hitchhiker
	prgList.programName = append(prgList.programName, "Hitchhiker")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Nuke-from-Orbit
	prgList.programName = append(prgList.programName, "Nuke-from-Orbit")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Paintjob
	prgList.programName = append(prgList.programName, "Paintjob")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Smoke-and-Mirrors
	prgList.programName = append(prgList.programName, "Smoke-and-Mirrors")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Swerve
	prgList.programName = append(prgList.programName, "Swerve")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Tantrum
	prgList.programName = append(prgList.programName, "Tantrum")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	//Tarball
	prgList.programName = append(prgList.programName, "Tarball")
	prgList.programStatus = append(prgList.programStatus, "inStore")
	prgList.programType = append(prgList.programType, "HACKING")
	///////////////////////////////////////////////////////////

	return prgList

}
