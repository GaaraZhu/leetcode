func convert(s string, numRows int) string {
    if(numRows==1) {
        return s
    }
    length := len(s)
    if (length <= numRows) {
        return s
    }
    
    var width int
    width = (numRows-1) * (length/(2*numRows-2)+1)
    array:=make([][]string, numRows)
    for i:=0; i<numRows; i++ {
        array[i] = make([]string, width)
    }
    
    for i:=0; i<length; i++ {
        r := (i+1)%(2*numRows-2)
        if (r==0) {
            r = 1
        } else if (r<=numRows) {
            r -= 1
        } else {
            r = 2*numRows-r-1
        }
        
        c := ((i+1)/(2*numRows-2))*(numRows-1)
        if ((i+1)%(2*numRows-2)!=0) {
            if ((i+1)%(2*numRows-2)>numRows) {
                c = c + (i+1)%(2*numRows-2)-numRows
            }
        } else {
	        c -= 1
	    }
        
        array[r][c] = s[i:i+1]
    }
    
    var result string
    result = ""
    for i:=0; i<numRows; i++ {
        for j:=0; j<width; j++ {
            if (array[i][j]!="") {
                result+=array[i][j]
            }
        }
    }
    
    return result
}