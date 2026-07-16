func groupAnagrams(strs []string) [][]string {
    var answers [][]string

    for _, str := range strs {
        if len(answers) == 0 {
            answers = append(answers, []string{str})
            continue
        }
        smap := make(map[rune]int)
        for _, letter := range str {
            smap[letter] = smap[letter] + 1
        }

        topildi := false

        for i, innerSlice := range answers {
            innermap := make(map[rune]int)
            for _, letter := range innerSlice[0] {
                innermap[letter] = innermap[letter] + 1
            }

            teng := true

            if len(smap) != len(innermap) {
                teng = false
            } else {
                for kalit, qiymat := range smap {
                    if innermap[kalit] != qiymat {
                        teng = false
                        break
                    }
                }
            }

            if teng {
                answers[i] = append(answers[i], str)
                topildi = true
                break
            }
        }

        if !topildi {
            answers = append(answers, []string{str})
        }
    }

    return answers
}

