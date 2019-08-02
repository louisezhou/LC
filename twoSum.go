func twoSum(nums []int, target int) []int {
    for i, v1 := range nums {
        if i+1 != len(nums) {
            for j, v2 := range nums[i+1:] {
                if target == (v1 + v2) {
                    return []int{i, i + j + 1}
                }
            }
        }
    }
    return []int{}
}
