
func waysToMakeFair(nums []int) int {
    length := len(nums)
    odd_sum := 0
    even_sum := 0
    result := 0
    var odd_sum_list []int = make([]int, length)  
    var even_sum_list []int = make([]int, length)  
    for index, num := range nums {
        if(index%2 == 1){
            odd_sum += num
        }else{
            even_sum += num
        }
        odd_sum_list[index] = odd_sum
        even_sum_list[index] = even_sum
    }
    
    for index, _ := range nums {
        current_back_odd_sum := (even_sum - even_sum_list[index])
        current_back_even_sum := (odd_sum - odd_sum_list[index])    

        if(index> 0){
            if(current_back_odd_sum + odd_sum_list[index - 1] == current_back_even_sum + even_sum_list[index - 1]){
                result +=1
            }
            
        }else{
            if(current_back_odd_sum == current_back_even_sum){
                result +=1
            }
        }
    }
    return result
}