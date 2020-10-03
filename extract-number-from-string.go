import "math"

// extract number from a given string, e.g. " s-123.67ddds" will return -123.67
func extractNum(s string) float64 {
	if len(s) == 0 {
		return 0
	}

	numMap := initNumMap()
	numFound := false
	result := float64(1)
	decimalBit := 0
	for i := 0; i < len(s); i++ {
		current := s[i : i+1]
		if current == " " && !numFound {
			continue
		} else if current == "-" && !numFound {
			if i+1 < len(s) {
				if _, ok := numMap[s[i+1:i+2]]; ok {
					result = float64(-1)
				}
			}
		} else if v, ok := numMap[current]; ok {
			if numFound {
				if s[i-1:i] == "." || decimalBit != 0 {
					decimalBit++
				}

				if decimalBit == 0 {
					if result < 0 {
						result = result*float64(10) - float64(v)
					} else {
						result = result*float64(10) + float64(v)
					}
				} else {
					if result < 0 {
						result -= float64(v) / math.Pow(float64(10), float64(decimalBit))
					} else {
						result += float64(v) / math.Pow(float64(10), float64(decimalBit))
					}
				}
			} else {
				numFound = true
				result = result * float64(v)
			}
		}
	}

	if !numFound {
		return 0
	}

	return result
}

func initNumMap() map[string]int {
	numMap := make(map[string]int)
	numMap["0"] = 0
	numMap["1"] = 1
	numMap["2"] = 2
	numMap["3"] = 3
	numMap["4"] = 4
	numMap["5"] = 5
	numMap["6"] = 6
	numMap["7"] = 7
	numMap["8"] = 8
	numMap["9"] = 9
	return numMap
}