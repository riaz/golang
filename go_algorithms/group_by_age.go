package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
This program takes in stream of inputs of the format
age Name
and we need to map the Name to a age group (10 yrs span) on the fly.
*/

func groupByAge() {
	fmt.Println("Please make entries manually of the format: Age <space> Name")

	reader := bufio.NewReader(os.Stdin)
	ageMap := make(map[string][]string)

	granularity := 10

	for {
		// read string will block until the delimiter is entered
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occured while reading the input")
			return
		}

		// removing the delimiter from the input
		input = strings.TrimSuffix(input, "\n")

		//fmt.Println(input)

		if len(input) == 0 {
			break
		}

		res := strings.Split(input, " ")

		age, err := strconv.Atoi(res[0])
		if err != nil {
			fmt.Printf("Bad input %s , this entry will be omitted", input)
			continue
		}

		name := res[1]

		lower := int(age/granularity) * granularity
		upper := lower + granularity

		// generating the key in the format <lower>-<upper>
		key := strconv.Itoa(lower) + "-" + strconv.Itoa(upper)

		// we check if the key already exists, if it exists append else do the same
		/*if _, ok := ageMap[key]; ok {
			//do something here
			ageMap[key] = append(ageMap[key], name)
		} else {
			ageMap[key] = append(ageMap[key], name)
		}*/
		ageMap[key] = append(ageMap[key], name)
	}

	fmt.Println(ageMap)

}
