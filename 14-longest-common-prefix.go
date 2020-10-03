func longestCommonPrefix(strs []string) string {
    prefixMap := make(map[string]int)
    for _, str := range strs {
        for i:=0; i<len(str); i++ {
            prefix := str[0:i+1]
            c, ok := prefixMap[prefix]
            if ok {
                prefixMap[prefix] = c+1
            } else {
                prefixMap[prefix] = 1
            }
        }
    }
    
    result := ""
    for p, c := range prefixMap {
        if c<len(strs) {
            continue
        }
        if len(p)>len(result) {
            result = p
        }
    }
    
    return result
}