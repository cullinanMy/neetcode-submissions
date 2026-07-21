func topKFrequent(nums []int, k int) []int {
    var answers []int
    numberMap := make(map[int]int)
   for _,number := range nums {
    numberMap[number]++
    } 
    for son := range numberMap {
    answers = append(answers, son)  
    }

    sort.Slice(answers,func(i,j int)bool{
        return numberMap[answers[i]]>numberMap[answers[j]]
    })

    return answers[:k]



}
