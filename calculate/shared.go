package calculate

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateTreeNumbers(targetNumber int) []int {
	firstNumber := rand.Intn(targetNumber-2) + 1
	secondNumber := rand.Intn(targetNumber-firstNumber-1) + 1
	thirdNumber := targetNumber - firstNumber - secondNumber

	return []int{firstNumber, secondNumber, thirdNumber}
}

func generateRemainderList(targetNumberList []int, divisor int) []int {
	remainderList := []int{}

	for _, currentNumber := range targetNumberList {
		JingJueNumber := currentNumber % divisor

		if JingJueNumber == 0 {
			JingJueNumber = 4
		}
		remainderList = append(remainderList, JingJueNumber)
	}

	return remainderList
}

func getGuaXiangPercentList(times int, guaXiangGenerator func() string) []GuaXiang {
	guaXiangList := []GuaXiang{}
	guaXiangMap := make(map[string]int)

	for i := 0; i <= times; i++ {
		guaXiang := guaXiangGenerator()
		guaXiangMap[guaXiang] = guaXiangMap[guaXiang] + 1
	}

	for k, v := range guaXiangMap {
		guaXiang := GuaXiang{k, 100 * (float32(v) / float32(times))}
		guaXiangList = append(guaXiangList, guaXiang)
	}

	sort.SliceStable(guaXiangList, func(i, j int) bool {
		return guaXiangList[i].percent > guaXiangList[j].percent
	})

	for _, guanXiang := range guaXiangList {
		println(guanXiang.name, fmt.Sprintf("%.2f", guanXiang.percent), "%")
	}

	return guaXiangList
}

func printList(title string, dataList []int) {
	// stringDataList := strings.Join(lo.Map(dataList, func(item int, _ int) string {
	// 	return string(item)
	// }), ",")

	// println(title, stringDataList)

	println()
	print(title, ": ")
	for _, item := range dataList {
		print(item)
	}
}
