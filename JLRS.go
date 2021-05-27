/*
 provides the zLog programming interface to the Go language.
 Copyright (C) 2020 JA1ZLO.
*/
package main

import (
	_ "embed"
	"strings"
	"strconv"
	"github.com/nextzlog/zylo"
)


//go:embed JLRS.dat
var JLRS string

func zcities() string {
	return JLRS
}

func zlaunch() {
	zylo.Notify("CQ!")
}

func zfinish() {
	zylo.Notify("Bye")
}

func zattach(test string, path string) {
}

func zdetach() {
}

func zinsert(qso *zylo.QSO) {
}

func zdelete(qso *zylo.QSO) {
}

func zverify(qso *zylo.QSO) {
	//multi
	call := strings.TrimSpace(qso.GetCall())
	if call != ""{
		qso.SetMul1(call[0:3])
	}

	rcvd := strings.TrimSpace(qso.GetRcvd())
	qso.SetMul2(rcvd)
	
	if rcvd != ""{
		//score
		get_serial,_ :=strconv.Atoi(rcvd)
		sent_serial ,_ := strconv.Atoi(strings.TrimSpace(qso.GetSent()))
		if qso.Dupe {
			qso.Score = 0
		} else {
			//OM
			if sent_serial <= 2000 {
				if get_serial <= 2000{
					qso.Score = 0
					qso.SetMul1("NO")
					qso.SetNote("invalid QSO, multi and score are 0")
				}
				if 2001 <= get_serial  &&   get_serial <= 5000{
					qso.Score = 1
				} 
				if 5001 <= get_serial{
					qso.Score = 5
				} 
			}
			//YL
			if 2001 <= sent_serial  &&   sent_serial <= 5000{
				if get_serial <= 2000{
					qso.Score = 1
				}
				if 2001 <= get_serial  &&   get_serial <= 5000{
					qso.Score = 5
				} 
				if 5001 <= get_serial{
					qso.Score = 5
				} 
			}
			//member
			if 5001 <= sent_serial{
				qso.Score = 1
			} 
		}	
	}	
}

func zpoints(score, mults int) int {
	return score * mults 
}

func zeditor(key int, name string) bool {
	return	false
}

func zbutton(btn int, name string) bool {
	return false
}

func main() {}
