package utils

import (
	"fmt"
	"strings"
)

// core function
func coreCompute(input []string) {
	var dirSize map[string]int = make(map[string]int)
	var currentDirectory string = ""

	for i := 0; i < len(input); i++ {
		var line string = input[i]
		// if nextLine > 0 {
		// 	nextLine--
		// 	continue
		// }
		/*
			Je veux savoir si la ligne est une commande et si c'est une commande on l'execute
		*/
		cmd := strings.Split(line, " ")[1]
		if cmd == "cd" {
			currentDirectory = getNameDirectory(input, &i)
			fmt.Println("currentDir = ", currentDirectory)
		} else if cmd == "ls" {
			dirSize[currentDirectory] = getDirectoryFileSizeSum(input, &i)
		}

		// si c'est une commande on l'execute
		// if isCommand(line) {
		// 	cmdLine := strings.Split(line, " ")
		// 	if cmdLine[1] == "cd" {
		// 		// nextLine = execCd(line, input, &i)
		// 		execCd(line, input, &i)
		// 	} else if cmdLine[1] == "ls" {
		// 		// nextLine = execLs(line, input, &i)
		// 		execLs(line, input, &i)
		// 	}
		// }
		// fmt.Println(line)
		// addOneToIdx(&i)
	}

}

func getDirectoryFileSizeSum(input []string, idx *int) int {
	var firstIter bool = true

	for _, data := range input[*idx:] {
		if firstIter {
			fmt.Println(data)
			*idx++
			firstIter = false
			continue
		}
		fmt.Println("data = ", data)
		if isCommand(data) {
			break
		}
		dataLine := strings.Split(data, " ")
		if dataLine[0] == "dir" {
			return 0
		}
		fmt.Println("", data)
		*idx++
		firstIter = false
	}
	*idx--
	return 0
}

func getNameDirectory(input []string, idx *int) string {
	cmdLine := strings.Split(input[*idx], " ")

	fmt.Println("cmdLine= ", cmdLine)
	if len(cmdLine) < 2 {
		panic("Error in file, cd must have an argument")
	}
	if cmdLine[1] == ".." {
		*idx++
		getNameDirectory(input, idx)
	}
	// *idx++
	return cmdLine[2]
}

// func checkType() {
// if reflect.TypeOf(xxx) == int => FILE
// else => DIR
// }

func isCommand(line string) bool {
	fmt.Println("line = ", line)
	rawLine := strings.Split(line, " ")
	fmt.Println("rawLine[0] = ", rawLine[0])
	if rawLine[0] == "$" {
		return true
	}
	return false
}

// func execCd(line string, input []string, currentInputIdx *int) {
// 	fmt.Println("=>", line)
// }

// func execLs(line string, input []string, currentInputIdx *int) {
// 	// fmt.Println("We pass in execLs")
// 	// for _, line := range input[*currentInputIdx:] {
// 	// 	*currentInputIdx++
// 	// 	if isCommand(line) {
// 	// 		return
// 	// 	}
// 	// 	fmt.Println("=", line)
// 	// }
// 	// foreach
// }

// func computeCommand(line string) {
// 	cmdLine := strings.Split(line, " ")
// 	cmd := cmdLine[1]

// 	if cmd == CD {
// 		arg := cmdLine[2]
// 	}
// }

/*
	Fonction recursive pour aller voir tous les dir
*/
