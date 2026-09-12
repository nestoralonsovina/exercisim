package raindrops

import "fmt"

func Convert(number int) string {
    var divisibleBy3 bool = number % 3 == 0
	var divisibleBy5 bool = number % 5 == 0
    var divisibleBy7 bool = number % 7 == 0

    result := ""

    if (divisibleBy3) { result += "Pling" }
    if (divisibleBy5) { result += "Plang" }
    if (divisibleBy7) { result += "Plong" }

    if (result != "") {
        return result
    }
    
    return fmt.Sprintf("%d", number)
}
