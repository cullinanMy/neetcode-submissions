func hasDuplicate(nums []int) bool {
    filteredNums:=make(map[int]bool) 
    for i := 0;i<len(nums);i++{
        if filteredNums[nums[i]]{
            return true
        }
        filteredNums[nums[i]] = true
    }

    return false
    
}
