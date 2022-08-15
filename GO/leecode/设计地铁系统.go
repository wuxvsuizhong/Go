package main

import (
	"fmt"
)

type CheckInfo struct {
	sName       string
	checkinTime int
}

type UndergroundSystem struct {
	cnt    map[string]int
	total  map[string]int
	record map[int]CheckInfo
}

// func (u *UndergroundSystem) checkIn(id int, stationName string, clock int) {
func (u *UndergroundSystem) checkIn(args ...interface{}) {
	id := args[0].(int)
	stationName := args[1].(string)
	clock := args[2].(int)
	// u.record[id].sName = stationName
	// u.record[id].checkinTime = clock
	u.record[id] = CheckInfo{sName: stationName, checkinTime: clock}
}

// func (u *UndergroundSystem) checkOut(id int, outStationName string, clock int) {
func (u *UndergroundSystem) checkOut(args ...interface{}) {
	id := args[0].(int)
	outStationName := args[1].(string)
	clock := args[2].(int)

	checkinStation := u.record[id].sName
	checkinTime := u.record[id].checkinTime

	u.total[checkinStation+"#"+outStationName] += clock - checkinTime
	u.cnt[checkinStation+"#"+outStationName] += 1
}

// func (u *UndergroundSystem) getAverageTime(startStation string, endStation string) (clock int) {
func (u *UndergroundSystem) getAverageTime(args ...interface{}) (clock int) {
	startStation := args[0].(string)
	endStation := args[1].(string)
	return u.total[startStation+"#"+endStation] * 1.0 / u.cnt[startStation+"#"+endStation]
}

func main() {
	var actions = []string{"UndergroundSystem", "checkIn", "checkIn", "checkIn", "checkOut", "checkOut", "checkOut", "getAverageTime", "getAverageTime", "checkIn", "getAverageTime", "checkOut", "getAverageTime"}

	var trace = [][]interface{}{{}, {45, "Leyton", 3}, {32, "Paradise", 8}, {27, "Leyton", 10}, {45, "Waterloo", 15}, {27, "Waterloo", 20}, {32, "Cambridge", 22}, {"Paradise", "Cambridge"}, {"Leyton", "Waterloo"}, {10, "Leyton", 24}, {"Leyton", "Waterloo"}, {10, "Waterloo", 38}, {"Leyton", "Waterloo"}}

	var undersys *UndergroundSystem
	results := make([]int, 0)
	for i, val := range actions {
		switch val {
		case "UndergroundSystem":
			undersys = &UndergroundSystem{
				cnt:    make(map[string]int, 0),
				total:  make(map[string]int, 0),
				record: make(map[int]CheckInfo),
			}
		case "checkIn":
			undersys.checkIn(trace[i]...)
		case "checkOut":
			undersys.checkOut(trace[i]...)
		case "getAverageTime":
			ret := undersys.getAverageTime(trace[i]...)
			results = append(results, ret)
		}
	}

	fmt.Println(results)
}
