func removeElement(nums []int, val int) int {
    var result int
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[result] = nums[i]
            result++
        }
    }
    return result
}
