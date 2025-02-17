package main

import (
	"fmt"
	"strconv"
)

// import pkg

func TimeViewer(from int64, to int64) {

	var listOfData []Result
	// get a group Json
	groupJson := LoadGroupJson()
	for _, v := range groupJson {
		// result time of each group
		var _indexedResult Result
		_indexedResult.GroupName = v.Group

		for _, p := range v.WorkSpace {
			// result time of each workSpace
			var _current spaceNdCalc
			_current.WorkSpace = p
			_indexedResult.SumGroup = append(_indexedResult.SumGroup, _current)
		}

		listOfData = append(listOfData, _indexedResult)
	}

	listOfData = TimeCalc(from, to, listOfData)

	PrintSumResult(listOfData)

	// _currentDir, _ := os.Getwd()
	// directoryName := filepath.Base(_currentDir)
	// logJson := LoadLogJson(directoryName)

}

// print calculated datas
func PrintSumResult(listOfData []Result) {

	// all time calclation
	var _allTime int64
	for _, v := range listOfData {
		_allTime += v.Sum
	}
	_, min, sec := UnixToHMS(_allTime)
	hour := int((_allTime - int64(sec) - int64(min)*60) / 3600)

	fmt.Printf(string(colorGreen))
	fmt.Print("\n")
	fmt.Print("[*]", string(colorReset))
	fmt.Println(" " + strconv.Itoa(hour) + "h " + strconv.Itoa(min) + "m " + strconv.Itoa(sec) + "s\n")

	for _, v := range listOfData {

		fmt.Printf(string(colorGreen))
		fmt.Print("["+v.GroupName+"]", string(colorReset))
		_, min_g, sec_g := UnixToHMS(v.Sum)
		hour_g := int((v.Sum - int64(sec) - int64(min)*60) / 3600)
		fmt.Println(" " + strconv.Itoa(hour_g) + "h " + strconv.Itoa(min_g) + "m " + strconv.Itoa(sec_g) + "s")

		for _, r := range v.SumGroup {
			_, min_sps, sec_sps := UnixToHMS(r.Sum)
			hour_sps := int((r.Sum - int64(sec) - int64(min)*60) / 3600)
			fmt.Println("> " + r.WorkSpace + " : " + strconv.Itoa(hour_sps) + "h " + strconv.Itoa(min_sps) + "m " + strconv.Itoa(sec_sps) + "s")
		}
		fmt.Print("\n")
	}

}

// UnixToHMS
