import (
    "fmt"
    "strings"
)

func romanToInt(s string) int {
    symbolMap := initSymbolMap()
    subtractionSymbolMap := initSubtractionSymbolMap()

    result:=0
    ss := strings.ToUpper(s)
    for i := 0; i <len(ss); i++ {
        if(i+1<len(ss)) {
            subtractionSymbol := ss[i:i+2]
            v, ok := subtractionSymbolMap[subtractionSymbol]
            if ok {
                result += v
                i++
                continue
            }
        }
        
        symbol := ss[i:i+1]
        vv, ok := symbolMap[symbol]
        if !ok {
            fmt.Errorf("invalid character in Roman numeral: %s", symbol)
        }
        
        result += vv
    }

    return result
}

func initSymbolMap() map[string]int {
    symbolMap := make(map[string]int)
    symbolMap["I"] = 1
    symbolMap["V"] = 5
    symbolMap["X"] = 10
    symbolMap["L"] = 50
    symbolMap["C"] = 100
    symbolMap["D"] = 500
    symbolMap["M"] = 1000
    return symbolMap
}

func initSubtractionSymbolMap() map[string]int {
    symbolMap := make(map[string]int)
    symbolMap["IV"] = 4
    symbolMap["IX"] = 9
    symbolMap["XL"] = 40
    symbolMap["XC"] = 90
    symbolMap["CD"] = 400
    symbolMap["CM"] = 900
    return symbolMap
}
