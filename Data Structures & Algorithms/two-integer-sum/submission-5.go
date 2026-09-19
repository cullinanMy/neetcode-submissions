func twoSum(nums []int, target int) []int {
	answers := make(map[int]int)

	for i,v := range nums{
		answers[v]=i
		}
    for i,n := range nums{
		diff := target - n 
		if j,found := answers[diff];found && j != i{
			return []int{i,j}
		}

	}
	return []int{}
}
