package main

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
