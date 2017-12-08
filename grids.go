package main

//TGrid -
type TGrid struct {
	TObj
	name string
	//deviceRating   int
	overwatchScore int
	lastSureOS     int
}

//IGrid - в икону входят файлы, персоны, айсы и хосты
type IGrid interface {
	IObj
	IGridOnly
}

var _ IGrid = (*TGrid)(nil)

//ToDo: перенести DeviceRating в TObj

//IGridOnly -
type IGridOnly interface {
	GetGridName() string
	//GetDeviceRating() int
	//GetOverwatchScore() int
	GetLastSureOS() int
	SetGridName(string)
	//SetDeviceRating(int)
	//SetOverwatchScore(int)
	SetLastSureOS(int)
}

// GetType -
func (g *TGrid) GetType() string {
	return "Grid"
}

//GetName -
func (g *TGrid) GetName() string {
	return g.name
}

//GetGridName -
func (g *TGrid) GetGridName() string {
	return g.name
}

//GetDeviceRating -
func (g *TGrid) GetDeviceRating() int {
	return g.deviceRating
}

//GetOverwatchScore -
func (g *TGrid) GetOverwatchScore() int {
	if g.overwatchScore < 0 {
		g.overwatchScore = 0
	}
	return g.overwatchScore
}

//GetLastSureOS -
func (g *TGrid) GetLastSureOS() int {
	return g.lastSureOS
}

//SetGridName -
func (g *TGrid) SetGridName(nam string) {
	g.name = nam
}

//SetDeviceRating -
func (g *TGrid) SetDeviceRating(dr int) {
	g.deviceRating = dr
}

//SetOverwatchScore -
func (g *TGrid) SetOverwatchScore(os int) {
	g.overwatchScore = os
}

//SetLastSureOS -
func (g *TGrid) SetLastSureOS(os int) {
	g.lastSureOS = os
}

//GetMarkSet -
func (g *TGrid) GetMarkSet() MarkSet {
	return g.markSet
}

//GetID -
func (g *TGrid) GetID() int {
	return g.id
}

//NewGrid -
func NewGrid(name string, dr int) *TGrid {
	g := TGrid{}
	g.name = name
	g.deviceRating = dr
	g.overwatchScore = 0
	g.id = -999
	gridList = append(gridList, &g)
	ObjByNames[g.name] = &g
	return &g
}

func createDefaultGrids() {
	NewGrid("Public Grid", 1)
	NewGrid("Local Grid", 2)
	NewGrid("Ares Global Grid", 3)
	NewGrid("AzGrid", 3)
	NewGrid("Eternal Horizon", 3)
	NewGrid("EvoGrid", 3)
	NewGrid("MCT GlobeNet", 3)
	NewGrid("NeoNetwork", 3)
	NewGrid("Renraku Okoku", 3)
	NewGrid("Saeder-Krupp Uberwelt", 3)
	NewGrid("Shiawase Central", 3)
	NewGrid("Wuxing Worldwide", 3)

}
