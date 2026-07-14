func isAnagram(s string, t string) bool {
    smap := make(map[rune]int)  
    tmap := make(map[rune]int) 

    for _,letter := range s {
        smap[letter] = smap[letter]+1
        }
    for _,letter := range t {
        tmap[letter] = tmap[letter]+1
        }
    if len(smap) != len(tmap) {
        return false
    }

    for kalit, qiymat := range smap {
        if tmap[kalit] != qiymat {
            return false
        }
    }

    return true


}
