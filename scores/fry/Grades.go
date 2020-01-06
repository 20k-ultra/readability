package fry

import (
	"fmt"
)

// Coordinate represents the x,y position on a graph
type Coordinate struct {
	x float32
	y float32
}

// Bounds represents graph coordinates for an area
type Bounds struct {
	topLeft Coordinate
	topRight Coordinate
	bottomLeft Coordinate
	bottomRight Coordinate
}

// Grade represents a grade on the Fry Readability Graph
// Provides the grade value (grade 1, 2, 3, etc) and the Bounds of that Grade
type Grade struct {
	value int
	boundaries Bounds
}

// GRADES stores all the possible grades in the Fry Readability Graph
var GRADES = []Grade{
	//Grade 1
	{1, Bounds{
		Coordinate{108, 25},   //topLeft
		Coordinate{120, 25},   //topRight
		Coordinate{108, 9.7},  //bottomLeft
		Coordinate{120, 12.9}, //bottomRight
	}},
	// Grade 2
	{2, Bounds{
		Coordinate{108, 9.7},  //topLeft
		Coordinate{120, 12.9}, //topRight
		Coordinate{108, 7.8},  //bottomLeft
		Coordinate{121, 11},   //bottomRight
	}},
	// Grade 3
	{3, Bounds{
		Coordinate{108, 7.8},  //topLeft
		Coordinate{121, 11},   //topRight
		Coordinate{108, 6.5},  //bottomLeft
		Coordinate{122, 8.8},  //bottomRight
	}},
	//Grade 4
	{4, Bounds{
		Coordinate{108, 6.5},  //topLeft
		Coordinate{122, 8.8},  //topRight
		Coordinate{108, 5.7},  //bottomLeft
		Coordinate{124, 7.7},  //bottomRight
	}},
	//Grade 5
	{5, Bounds{
		Coordinate{108, 5.7},  //topLeft
		Coordinate{134, 7.7},  //topRight
		Coordinate{108, 5},    //bottomLeft
		Coordinate{126, 6.9},  //bottomRight
	}},
	//Grade 6
	{6, Bounds{
		Coordinate{108, 5},    //topLeft
		Coordinate{126, 6.9},  //topRight
		Coordinate{108, 3.9},  //bottomLeft
		Coordinate{129, 5.9},  //bottomRight
	}},
	//Grade 7
	{7, Bounds{
		Coordinate{108, 3.9},  //topLeft
		Coordinate{129, 5.9},  //topRight
		Coordinate{109, 2},    //bottomLeft
		Coordinate{136, 5.1},  //bottomRight
	}},
	//Grade 8
	{8, Bounds{
		Coordinate{109, 2},    //topLeft
		Coordinate{136, 5.1},  //topRight
		Coordinate{122, 2},    //bottomLeft
		Coordinate{141, 4.8},  //bottomRight
	}},
	//Grade 9
	{9, Bounds{
		Coordinate{122, 2},    //topLeft
		Coordinate{141, 4.8},  //topRight
		Coordinate{136, 2},    //bottomLeft
		Coordinate{148, 4.4},  //bottomRight
	}},
	//Grade 10
	{10, Bounds{
		Coordinate{136, 2},    //topLeft
		Coordinate{148, 4.4},  //topRight
		Coordinate{143, 2},    //bottomLeft
		Coordinate{153, 4.3},  //bottomRight
	}},
	//Grade 11
	{11, Bounds{
		Coordinate{143, 2},    //topLeft
		Coordinate{153, 4.3},  //topRight
		Coordinate{148, 2},    //bottomLeft
		Coordinate{158, 4.2},  //bottomRight
	}},
	//Grade 12
	{12, Bounds{
		Coordinate{148, 2},    //topLeft
		Coordinate{158, 4.2},  //topRight
		Coordinate{155, 2},    //bottomLeft
		Coordinate{162, 4.1},  //bottomRight
	}},
	//Grade 13
	{13, Bounds{
		Coordinate{155, 2},    //topLeft
		Coordinate{162, 4.1},  //topRight
		Coordinate{161, 2},    //bottomLeft
		Coordinate{166, 4},    //bottomRight
	}},
	//Grade 14
	{14, Bounds{
		Coordinate{161, 2},    //topLeft
		Coordinate{166, 4},    //topRight
		Coordinate{168, 2},    //bottomLeft
		Coordinate{170, 3.9},  //bottomRight
	}},
	//Grade 15
	{15, Bounds{
		Coordinate{168, 2},     //topLeft
		Coordinate{170, 3.9},   //topRight
		Coordinate{168, 2},     //bottomLeft
		Coordinate{174, 3.8},   //bottomRight
	}},
}

func getGrade(x float32, y float32) int {
	bool1, bool2 := false, false
	var selGrade Grade
	for i, grade := range GRADES {
		if (i == 7){
			fmt.Printf("\n X: %f", x)
			fmt.Printf("\n Y: %f", y)
			fmt.Printf("\n\n")
		}


		if (y > grade.boundaries.bottomLeft.y && y > grade.boundaries.bottomRight.y &&
			y < grade.boundaries.topLeft.y && y > grade.boundaries.topRight.y){
				fmt.Printf("%f is between %v", y, grade.boundaries)
				bool1 = true
		}
		if (x > grade.boundaries.bottomLeft.x && x > grade.boundaries.bottomRight.x &&
			x < grade.boundaries.topLeft.x && x > grade.boundaries.topRight.x){
				bool2 = true
		}

		if (bool1 && bool2){
			selGrade = grade
		}
	}

	return selGrade.value
}

func betweenCoordinates(coord1 Coordinate, coord2 Coordinate, location Coordinate) bool {
	return greaterThenCoordinate(coord1, location) && lesserThenCoordinate(coord2, location)
}

func lesserThenCoordinate(coord Coordinate, location Coordinate) bool {
	return coord.x <= location.x && coord.y <= location.y
}

func greaterThenCoordinate(coord Coordinate, location Coordinate) bool {
	return coord.x >= location.x && coord.y >= location.y
}