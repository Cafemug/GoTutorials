import "sort"

func makeList(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := append([]int{}, root.Val)
	res1 := makeList(root.Left)
	res2 := makeList(root.Right)
	result = append(result, res1...)
	result = append(result, res2...)
	return result
}
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result_left := makeList(root1)
	result_right := makeList(root2)
	result := append(result_left, result_right...)
	sort.Ints(result)
	return result
}