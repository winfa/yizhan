import "math/rand"

func generateJinJianNumber() {
	threeNumbers := generateTreeNumbers(30)

	jinJianNumber := []
	// for each threeNumbers
	for i := 0; i < threeNumbers {
			
	}
}

func generateTreeNumbers(int targetNumber) {
	firstNumber := rand.Intn(targetNumber)
	secondNumber := rand.Intn(targetNumber - firstNumber)
	thirdNumber := targetNumber - firstNumber - secondNumber

	return [3]int{firstNumber, secondNumber, thirdNumber}
}