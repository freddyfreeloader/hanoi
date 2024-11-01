package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

/* The width of the towers on the console */
var widthOfTower = 17

/* The 3 towers as a list */
var towers [3][]int

/* Each tower gets its own list for stacking */
var slice1 []int
var slice2 []int
var slice3 []int

/* Counter for "moves" */
var count = 1

/* Tower labels */
const tower1 = "Start tower: 1"
const tower2 = "Intermediate: 2"
const tower3 = "Target: 3"

/* Number of disks for the game, default 5 */
var disks = 5

/* Maximum allowed number of disks */
const maxDisks = 15

func main() {
	userInput()
	setTowerWidth()
	initTowers()
	printTowers() // Draw towers once before the recursion
	move(1, 3, disks)
}

func userInput() {
	fmt.Println("How many disks?")
	_, err := fmt.Scan(&disks)
	if err != nil || disks < 1 || disks > maxDisks {
		fmt.Println("Error!")
		if disks < 1 {
			fmt.Println("More disks please!!!")
		}
		if disks > maxDisks {
			total := math.Pow(2, float64(disks)) - 1
			fmt.Printf("That would be %.0f recursions! Let's not do that..", total)
		}
		return
	}
}

func move(a int, b int, i int) {
	if i == 1 {
		visualize(a, b)
	} else {
		move(a, 6-a-b, i-1)
		move(a, b, 1)
		move(6-a-b, b, i-1)
	}
}

func setTowerWidth() {
	// Adjust the width of the towers to the number of disks
	widthOfTower += 2 * disks
	// Must be even, otherwise towers are misaligned
	if widthOfTower%2 != 0 {
		widthOfTower++
	}
}

func initTowers() {
	var k = 0
	// Create a list (slice) that corresponds to the number of disks...
	for k < disks {
		slice1 = append(slice1, disks-k)
		k++
	}
	// ...and then put the list in the first tower
	towers[0] = slice1
	// The other towers are initially empty
	towers[1] = slice2
	towers[2] = slice3
}

func visualize(a int, b int) {
	printLine()
	printHeadline(a, b)
	changeSlices(a, b)
	printTowers()
	printLegend()
	count++
}

func printLine() {
	// \u2500 makes a horizontal line
	line := strings.Repeat("\u2500", widthOfTower*3)
	fmt.Println(line)
}

func printHeadline(a int, b int) {
	fmt.Printf("%d. Move: Transfer a disk from %d to %d:", count, a, b)
}

func getSizeOfGreatestTower() int {
	maxSize := 0
	for _, tower := range towers {
		if len(tower) > maxSize {
			maxSize = len(tower)
		}
	}
	return maxSize
}

func printTowers() {
	maxDiskSize := getSizeOfGreatestTower()
	fmt.Println()

	for i := maxDiskSize - 1; i >= 0; i-- {
		for j := 0; j < len(towers); j++ {
			if i < len(towers[j]) {
				// Make the disk a bit larger, looks better..
				diskSize := towers[j][i] * 2
				// "Disk" is initially whitespace, gets a colored background later
				disk := strings.Repeat(" ", diskSize)
				// Fill the empty space between horizontal disks with spaces
				padding := strings.Repeat(" ", (widthOfTower-diskSize)/2)

				fmt.Printf(padding + colorDisk(disk) + padding)
			} else {
				whitespace := strings.Repeat(" ", widthOfTower)
				fmt.Printf(whitespace)
			}
		}
		fmt.Println()
	}
}

func colorDisk(disk string) string {
	colorCodes := []string{
		"\033[31;41m", // red
		"\033[32;42m", // green
		"\033[33;43m", // yellow
		"\033[34;44m", // blue
		"\033[35;45m", // purple
		"\033[36;46m", // cyan
		"\033[30;47m", // white
	}
	index := (len(disk)/2 - 1) % len(colorCodes)
	color := colorCodes[index]
	reset := "\033[0m"

	return color + disk + reset
}

func printLegend() {
	legends := []string{tower1, tower2, tower3}
	for _, legend := range legends {
		// Use a rune array to correctly count string length with Unicode characters
		runes := []rune(legend)
		// If the RuneCount of the legend string is odd, add a space before the last character
		if utf8.RuneCountInString(legend)%2 != 0 {
			legend = string(runes[:len(runes)-1]) + " " + string(runes[len(runes)-1])
		}

		// Calculate the padding based on the visible length of the string in runes
		padding := (widthOfTower - utf8.RuneCountInString(legend)) / 2

		// Print the legend with padding on both sides
		fmt.Print(strings.Repeat(" ", padding) + legend + strings.Repeat(" ", padding))
	}
	fmt.Println()

}

func changeSlices(a int, b int) {
	// Slice indices start at 0, but towers start at 1, so subtract 1
	a--
	b--
	lastIndexA := len(towers[a]) - 1

	// Add the last element from towers[a] to the slice towers[b].
	towers[b] = append(towers[b], towers[a][lastIndexA])
	// Remove the last element from the slice towers[a].
	towers[a] = towers[a][:lastIndexA]
}
