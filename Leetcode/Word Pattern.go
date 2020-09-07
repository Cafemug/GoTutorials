func wordPattern(pattern string, str string) bool {
    str_slice := strings.Split(str, " ")
    m := make(map[rune]string)
    k := make(map[string]rune)
    pattern_len := len(pattern)
    str_len := len(str_slice)
    if pattern_len != str_len{
        return false
    }
    for i,item := range pattern{
        if val, ok := m[item]; ok {
            if val != str_slice[i] {
                return false
            }
        }else{
            m[item] = str_slice[i]
        }
        if val, ok := k[str_slice[i]]; ok {
            if val != item {
                return false
            }
        }else{
            k[str_slice[i]] = item
        }
        
        
    }
    return true
}