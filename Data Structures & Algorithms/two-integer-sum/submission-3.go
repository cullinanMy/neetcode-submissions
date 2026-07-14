func twoSum(nums []int, target int) []int {
index_array := make([]int, 2)
    for i := 0;i<len(nums);i++{
        for j:= i+1;j<len(nums);j++{
            if nums[i]+nums[j]==target{
                index_array[0] = i
                index_array[1] = j
                return  index_array 
            }
        }
    }
    return nil
}
