func intToRoman(num int) string {
    
    var result string
    result = ""
    
    if (num >= 1000) {
        for i:=0; i<num/1000; i++ {
            result += "M"
        }
    }
    
    num = num%1000
    if (num>=100) {
        if(num/100<4){
            for i:=0; i<num/100; i++ {
                result += "C"
            }
        } else if (num/100==4) {
            result += "CD"
        } else if (num/100<9){
            result += "D"
            for i:=0; i<(num-500)/100; i++ {
                result += "C"
            }
        } else {
            result += "CM"
        }
    }
    
    num = num%100
    if (num>=10) {
        if(num/10<4){
            for i:=0; i<num/10; i++ {
                result += "X"
            }
        } else if (num/10==4) {
            result += "XL"
        } else if (num/10<9){
            result += "L"
            for i:=0; i<(num-50)/10; i++ {
                result += "X"
            }
        } else {
            result += "XC"
        }
    }
    
    num = num%10
    if(num<4){
        for i:=0; i<num; i++ {
            result += "I"
        }
    } else if (num==4) {
        result += "IV"
    } else if (num<9){
        result += "V"
        for i:=0; i<num-5; i++ {
            result += "I"
        }
    } else {
        result += "IX"
    }
    
    return result
}
