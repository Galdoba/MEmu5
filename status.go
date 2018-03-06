package main

import (
	"sort"
)

type TStatus struct {
	source  IObj
	target  IObj
	name    string
	counter int
}

type IStatus interface {
	RunTest() string
	Source() IObj
	SetSource(IObj)
	Target() IObj
	SetTarget(IObj)
	Counter() int
	SetCounter(int)
	Name() string
	SetName(string)
	Update()
}

var _ IStatus = (*TStatus)(nil)

/*//CFDBMap - Complex Forms Storing Map (global)
var StatusMap = map[string]IStatus{}*/

func (st *TStatus) RunTest() string {
	if st.source != nil {
		return "Source ID = " + iStr(st.source.GetID()) + ";{red} Name = " + st.name
	}
	return "Source ID = ХЗ кто;{red} Name = " + st.name
}

func (st *TStatus) Source() IObj {
	return st.source
}

func (st *TStatus) SetSource(src IObj) {
	st.source = src
}

func (st *TStatus) Target() IObj {
	return st.target
}

func (st *TStatus) SetTarget(targ IObj) {
	st.target = targ
}

func (st *TStatus) Counter() int {
	return st.counter
}

func (st *TStatus) SetCounter(c int) {
	st.counter = c
}

func (st *TStatus) Name() string {
	return st.name
}

func (st *TStatus) SetName(n string) {
	st.name = n
}

func (st *TStatus) Update() {
	if st.counter > 0 && st.name != "Nominal" {
		st.counter--
	}
	if st.name == "Nominal" {
		st.counter++
	}

}

func NewStatus(s string, c int, o IObj) IStatus {
	st := TStatus{}
	st.target = o //.(IObj)
	st.name = s
	st.counter = c
	return &st
}

func (status StatusMap) UpdateStatuses() {
	//status := o.Status()
	for key, st := range status.ByName {
		st.Update()
		if st.Counter() == 0 && st.Name() != "Nominal" {
			delete(status.ByName, key)
		}
	}
}

func UpdateStatusesForAll() {
	for _, obj := range ObjByNames {
		stat := obj.Status()
		stat.UpdateStatuses()
	}
}

func AllStatuses(o IObj) []string {
	status := o.Status()
	var activeSt []string
	for key, st := range status.ByName {
		if st.Counter() > 0 {
			activeSt = append(activeSt, key)
		}
	}
	sort.Strings(activeSt)
	return activeSt
}

func (sm StatusMap) HaveName(s string) bool {
	for key, _ := range sm.ByName {
		//printLog("Key is " + key + " but need " + s)
		if key == s {
			return true
		}
	}
	return false
}
