package main

import "fmt"

func accessMapVal() {

	//create map
	flowerColour := map[string]string{"sunflower": "yellow", "rose": "red", "lily": "pinkish"}

	fmt.Println(flowerColour)
	fmt.Println(flowerColour["rose"])

	// now changeing the data of map

	fmt.Println("")
	fmt.Println("after changeing the value")
	flowerColour["rose"] = "black"
	fmt.Println(flowerColour)
	fmt.Println(flowerColour["rose"])

	//add elements in the map

	flowerColour["lotus"] = "pink"
	flowerColour["jasmin"] = "white"

	fmt.Println(flowerColour)
	//itrate on map
	fmt.Println("\n")
	for key, pair := range flowerColour {
		fmt.Println(key, " :", pair)
	}

	fmt.Println("")
	// delete element from  the map
	fmt.Println("delete element from  the map")
	delete(flowerColour, "lily")
	for key, pair := range flowerColour {
		fmt.Println(key, " :", pair)
	}

}
