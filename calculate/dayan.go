package calculate

import (
	"strconv"
)

func generateYao() int {
	_, firtRoundEndNumber := generateRoundNumber(49)
	_, secondRoundEndNumber := generateRoundNumber(firtRoundEndNumber)
	_, thirdRoundEndNumber := generateRoundNumber(secondRoundEndNumber)

	return thirdRoundEndNumber / 4
}

func dayanYaoGenerator() string {
	return strconv.Itoa(generateYao())
}

func ProcessDayanTest() {
	// hexagrams := generateHexagrams()

	// for _, yao := range hexagrams {
	// 	fmt.Println("爻：", yao)
	// }

	getGuaXiangPercentList(1000000, dayanYaoGenerator)
}

func generateRoundNumber(startNumber int) ([]int, int) {
	threeNumbers := generateTreeNumbers(startNumber)
	roundNumberList := generateRemainderList(threeNumbers, 4)

	return roundNumberList, startNumber - roundNumberList[0] - roundNumberList[1] - roundNumberList[2]
}

func generateHexagrams() []int {
	hexagrams := []int{}
	for i := 0; i < 6; i++ {
		generateYao := generateYao()
		hexagrams = append(hexagrams, generateYao)
	}

	return hexagrams
}
