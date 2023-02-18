package main

import (
	"stepik/anonymous"
	"stepik/conversions"
	"stepik/datetimes"
	"stepik/files"
	"stepik/interfaces"
	"stepik/jsons"
	"stepik/mappings"
	"stepik/parallelism"
)

func main() {
	//mappingsTesting()
	//conversionsTesting()
	//anonymousTesting()
	//interfacesTesting()
	//filesTesting()
	//JsonsTesting()
	//DateTimesTesting()
	ParallelismTesting()
}

func mappingsTesting() {
	//mappings.StepFive()
	mappings.StepSix()
}

func conversionsTesting() {
	//conversions.StepThree()
	//conversions.StepThirteen()
	conversions.StepFourteen()
}

func anonymousTesting() {
	anonymous.StepSeven()
}

func interfacesTesting() {
	//interfaces.StepTen()
	interfaces.StepThirteen()
}

func filesTesting() {
	//files.StepNine()
	//files.StepThirteen()
	files.StepFourteen()
}

func JsonsTesting() {
	//jsons.StepSix()
	jsons.StepNine()
}

func DateTimesTesting() {
	//datetimes.StepThree()
	//datetimes.StepFour()
	//datetimes.StepSeven()
	datetimes.StepEight()
}

func ParallelismTesting() {
	//parallelism.StepEight()
	//parallelism.StepNine()
	parallelism.StepTen()
}
