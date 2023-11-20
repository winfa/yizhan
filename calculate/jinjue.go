package calculate

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type GuaXiang struct {
	name    string
	percent float32
}

func processJinJueTest(targetNumber int, times int) {
	guaXiangList := []GuaXiang{}
	guaXiangMap := make(map[string]int)

	for i := 0; i <= times; i++ {
		JingJueNumbers := generateJingJueNumbers(targetNumber)
		guaXiang := convertJingJueNumberToGuaXiang(JingJueNumbers)
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

	jiGuaList := lo.Filter(guaXiangList, func(guaXiang GuaXiang, _ int) bool {
		return strings.Contains(guaXiang.name, "吉")
	})
	jiGuaPercent := lo.SumBy(jiGuaList, func(item GuaXiang) float32 {
		return item.percent
	})
	println("吉卦：", fmt.Sprintf("%.2f", jiGuaPercent), "%")

	xiongGuaList := lo.Filter(guaXiangList, func(guaXiang GuaXiang, _ int) bool {
		return strings.Contains(guaXiang.name, "凶")
	})
	xiongGuaPercent := lo.SumBy(xiongGuaList, func(item GuaXiang) float32 {
		return item.percent
	})
	println("凶卦：", fmt.Sprintf("%.2f", xiongGuaPercent), "%")

	println("总共存在", len(guaXiangList), "种不同的卦象")
}

func generateJingJueNumbers(targetNumber int) []int {
	threeNumbers := generateTreeNumbers(targetNumber)

	JingJueNumbers := []int{}

	for _, currentNumber := range threeNumbers {
		JingJueNumber := currentNumber % 4

		if JingJueNumber == 0 {
			JingJueNumber = 4
		}
		JingJueNumbers = append(JingJueNumbers, JingJueNumber)
	}

	return JingJueNumbers
}

func convertJingJueNumberToGuaXiang(JingJueNumbers []int) string {
	guaXiang := ""

	for _, itemNumber := range JingJueNumbers {
		guaXiang = guaXiang + strconv.Itoa(itemNumber)
	}

	if guaXiang == "433" {
		return "甲（凶）"
	}
	if guaXiang == "411" {
		return "乙（吉）"
	}
	if guaXiang == "343" {
		return "丙（吉）"
	}
	if guaXiang == "424" {
		return "丁（吉）"
	}
	if guaXiang == "312" {
		return "戊（吉）"
	}
	if guaXiang == "334" {
		return "己（凶）"
	}
	if guaXiang == "231" {
		return "任（凶）"
	}
	if guaXiang == "222" {
		return "癸（吉）"
	}
	if guaXiang == "213" {
		return "子（吉）"
	}
	if guaXiang == "141" {
		return "丑（未）"
	}
	if guaXiang == "132" {
		return "寅（凶）"
	}
	if guaXiang == "321" {
		return "卯（凶）"
	}
	if guaXiang == "123" {
		return "辰（吉）"
	}
	if guaXiang == "114" {
		return "巳（吉）"
	}
	if guaXiang == "442" {
		return "午（吉）"
	}
	if guaXiang == "244" {
		return "未（凶）"
	}

	return guaXiang
}
