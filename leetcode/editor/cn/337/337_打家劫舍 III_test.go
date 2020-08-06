package leetcode_cn_337

import (
	. "go-leetcode-practice/leetcode/editor/cn/common"
	"testing"
)

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
// Related Topics 树 深度优先搜索
// 👍 466 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	result := robInternal(root)
	return max(result[0], result[1])
}

func robInternal(root *TreeNode) []int {
	result := make([]int, 2)
	if root == nil {
		return result
	}
	left := robInternal(root.Left)
	right := robInternal(root.Right)
	//result[0]是当前节点不偷 result[1]是当前节点偷
	result[0] = max(left[0], left[1]) + max(right[0], right[1])
	//当前节点偷 那么他的左右节点都不能偷
	result[1] = root.Val + left[0] + right[0]
	return result
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func Test(t *testing.T) {

}
