func isAnagram(s string, t string) bool {
smap, tmap := make(map[rune]int), make(map[rune]int)
	if len(s)!= len(t){
		return false
	}

	for i:=0;i<len(s);i++{
		smap[rune(s[i])]++
		tmap[rune(t[i])]++
		}
    for key,value := range smap{
		if value != tmap[key]{
			return false
		}
	}

	return true



}
