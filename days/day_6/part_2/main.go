package main

import (
	"fmt"
	"strconv"
	"strings"
)

const maxvalues = 1000

// Process the instructions and update the grid
func processInstructions(instructions []string) [maxvalues][maxvalues]int {
	var grid [maxvalues][maxvalues]int

	for _, instruction := range instructions {
		parts := strings.Fields(instruction)
		var action string
		var startX, startY, endX, endY int

		if len(parts) == 7 {
			action = parts[0] + " " + parts[1]
			startX, _ = strconv.Atoi(parts[2][:strings.Index(parts[2], ",")])
			startY, _ = strconv.Atoi(parts[2][strings.Index(parts[2], ",")+1:])
			endX, _ = strconv.Atoi(parts[4][:strings.Index(parts[4], ",")])
			endY, _ = strconv.Atoi(parts[4][strings.Index(parts[4], ",")+1:])
		} else if len(parts) == 6 {
			action = parts[0]
			startX, _ = strconv.Atoi(parts[1][:strings.Index(parts[1], ",")])
			startY, _ = strconv.Atoi(parts[1][strings.Index(parts[1], ",")+1:])
			endX, _ = strconv.Atoi(parts[3][:strings.Index(parts[3], ",")])
			endY, _ = strconv.Atoi(parts[3][strings.Index(parts[3], ",")+1:])
		}

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				switch action {
				case "turn on":
					grid[x][y] = 1
				case "turn off":
					grid[x][y] = 0
				case "toggle":
					grid[x][y] ^= 1
				}
			}
		}
	}
	return grid
}

// Count the number of lit bulbs
func countLitBulbs(grid [maxvalues][maxvalues]int) int {
	count := 0
	for i := 0; i < maxvalues; i++ {
		for j := 0; j < maxvalues; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	// Example instructions
	instructions := []string{
		"turn off 660,55 through 986,197",
		"turn off 341,304 through 638,850",
		"turn off 199,133 through 461,193",
		"toggle 322,558 through 977,958",
		"toggle 537,781 through 687,941",
		"turn on 226,196 through 599,390",
		"turn on 240,129 through 703,297",
		"turn on 317,329 through 451,798",
		"turn on 957,736 through 977,890",
		"turn on 263,530 through 559,664",
		"turn on 158,270 through 243,802",
		"toggle 223,39 through 454,511",
		"toggle 544,218 through 979,872",
		"turn on 313,306 through 363,621",
		"toggle 173,401 through 496,407",
		"toggle 333,60 through 748,159",
		"turn off 87,577 through 484,608",
		"turn on 809,648 through 826,999",
		"toggle 352,432 through 628,550",
		"turn off 197,408 through 579,569",
		"turn off 1,629 through 802,633",
		"turn off 61,44 through 567,111",
		"toggle 880,25 through 903,973",
		"turn on 347,123 through 864,746",
		"toggle 728,877 through 996,975",
		"turn on 121,895 through 349,906",
		"turn on 888,547 through 931,628",
		"toggle 398,782 through 834,882",
		"turn on 966,850 through 989,953",
		"turn off 891,543 through 914,991",
		"toggle 908,77 through 916,117",
		"turn on 576,900 through 943,934",
		"turn off 580,170 through 963,206",
		"turn on 184,638 through 192,944",
		"toggle 940,147 through 978,730",
		"turn off 854,56 through 965,591",
		"toggle 717,172 through 947,995",
		"toggle 426,987 through 705,998",
		"turn on 987,157 through 992,278",
		"toggle 995,774 through 997,784",
		"turn off 796,96 through 845,182",
		"turn off 451,87 through 711,655",
		"turn off 380,93 through 968,676",
		"turn on 263,468 through 343,534",
		"turn on 917,936 through 928,959",
		"toggle 478,7 through 573,148",
		"turn off 428,339 through 603,624",
		"turn off 400,880 through 914,953",
		"toggle 679,428 through 752,779",
		"turn off 697,981 through 709,986",
		"toggle 482,566 through 505,725",
		"turn off 956,368 through 993,516",
		"toggle 735,823 through 783,883",
		"turn off 48,487 through 892,496",
		"turn off 116,680 through 564,819",
		"turn on 633,865 through 729,930",
		"turn off 314,618 through 571,922",
		"toggle 138,166 through 936,266",
		"turn on 444,732 through 664,960",
		"turn off 109,337 through 972,497",
		"turn off 51,432 through 77,996",
		"turn off 259,297 through 366,744",
		"toggle 801,130 through 917,544",
		"toggle 767,982 through 847,996",
		"turn on 216,507 through 863,885",
		"turn off 61,441 through 465,731",
		"turn on 849,970 through 944,987",
		"toggle 845,76 through 852,951",
		"toggle 732,615 through 851,936",
		"toggle 251,128 through 454,778",
		"turn on 324,429 through 352,539",
		"toggle 52,450 through 932,863",
		"turn off 449,379 through 789,490",
		"turn on 317,319 through 936,449",
		"toggle 887,670 through 957,838",
		"toggle 671,613 through 856,664",
		"turn off 186,648 through 985,991",
		"turn off 471,689 through 731,717",
		"toggle 91,331 through 750,758",
		"toggle 201,73 through 956,524",
		"toggle 82,614 through 520,686",
		"toggle 84,287 through 467,734",
		"turn off 132,367 through 208,838",
		"toggle 558,684 through 663,920",
		"turn on 237,952 through 265,997",
		"turn on 694,713 through 714,754",
		"turn on 632,523 through 862,827",
		"turn on 918,780 through 948,916",
		"turn on 349,586 through 663,976",
		"toggle 231,29 through 257,589",
		"toggle 886,428 through 902,993",
		"turn on 106,353 through 236,374",
		"turn on 734,577 through 759,684",
		"turn off 347,843 through 696,912",
		"turn on 286,699 through 964,883",
		"turn on 605,875 through 960,987",
		"turn off 328,286 through 869,461",
		"turn off 472,569 through 964,695",
		"turn on 587,297 through 735,973",
		"toggle 223,22 through 608,663",
		"toggle 194,319 through 623,768",
		"turn on 207,546 through 926,876",
		"toggle 457,183 through 488,394",
		"turn off 473,363 through 686,832",
		"turn off 597,547 through 844,631",
		"turn off 547,879 through 974,925",
		"turn on 130,264 through 212,501",
		"toggle 601,304 through 779,388",
		"toggle 182,394 through 781,609",
		"turn off 95,146 through 965,174",
		"turn off 190,402 through 563,828",
		"toggle 163,373 through 306,752",
		"turn off 338,484 through 636,651",
		"toggle 136,78 through 853,991",
		"toggle 241,794 through 413,896",
		"turn on 892,873 through 904,918",
		"turn off 577,137 through 980,911",
		"toggle 25,892 through 510,903",
		"toggle 248,158 through 396,768",
		"toggle 0,377 through 993,980",
		"turn off 24,24 through 977,975",
		"turn on 728,382 through 971,396",
		"turn off 21,789 through 990,823",
		"toggle 670,292 through 754,958",
		"turn off 446,98 through 623,810",
		"toggle 109,61 through 872,790",
		"toggle 786,527 through 915,897",
		"toggle 241,870 through 350,917",
		"turn off 371,447 through 763,859",
		"turn on 182,262 through 315,295",
		"toggle 491,317 through 874,589",
		"toggle 287,686 through 800,911",
		"turn on 105,783 through 682,832",
		"turn off 535,692 through 822,740",
		"toggle 71,107 through 417,357",
		"turn on 234,377 through 765,817",
		"toggle 359,369 through 406,687",
		"toggle 333,665 through 354,714",
		"turn off 263,382 through 277,859",
		"turn off 682,825 through 953,871",
		"turn off 56,96 through 956,557",
		"toggle 849,645 through 904,868",
		"turn on 737,689 through 785,958",
		"turn on 1,549 through 998,715",
		"toggle 351,440 through 651,741",
		"turn on 570,139 through 830,548",
		"toggle 93,650 through 730,953",
		"turn off 171,72 through 558,811",
		"toggle 640,951 through 912,981",
		"toggle 426,179 through 967,920",
		"turn off 504,117 through 758,976",
		"toggle 101,124 through 853,685",
		"toggle 248,106 through 895,261",
		"turn off 58,220 through 703,779",
		"turn on 270,342 through 829,686",
		"turn off 0,0 through 999,999",
	}

	grid := processInstructions(instructions)
	litBulbs := countLitBulbs(grid)

	fmt.Printf("Total lit bulbs: %d\n", litBulbs)
}
