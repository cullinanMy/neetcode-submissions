import (
    "maps"
)

func groupAnagrams(strs []string) [][]string {
    var answers [][]string

    for _, str := range strs {
        if len(answers) == 0 {
            answers = append(answers, []string{str})
            continue
        }

        lettersMap := make(map[rune]int)
        for _, letter := range str {
            lettersMap[letter] = lettersMap[letter] + 1
        }

        found := false

        for indexAns, answer := range answers {
            answersMap := make(map[rune]int)
            for _, letter := range answer[0] {
                answersMap[letter] = answersMap[letter] + 1
            }

            if maps.Equal(lettersMap, answersMap) {
                answers[indexAns] = append(answers[indexAns], str) // guruhga qo'sh
                found = true
                break
            }
        }

        if !found {
            answers = append(answers, []string{str})
        }
    }

    return answers
}

