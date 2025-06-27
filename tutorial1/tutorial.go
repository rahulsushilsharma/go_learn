package tutorial1_test

import "errors"

func sum(val int) (int, error) {
	sum := 0
	var errorVal error
	for i := range val {
		sum += i
	}
	if sum < 1 {
		errorVal = errors.New("sum low")
	}

	return sum, errorVal

}
func twoSum(nums []int, target int) []int {
	arr := []int{0, 0}
	m := make(map[int]int)
	for i := range nums {
		m[target-nums[i]] = i
	}

	for i := range nums {
		val, found := m[nums[i]]
		if found && val != i {
			arr[0] = val
			arr[1] = i
			return arr
		}

	}
	return arr
}

func isPalindrome(x int) bool {
	result := false

	println(x)
	for {

	}

	return result

}

func removeDuplicates(nums []int) int {
	// hash := make(map[int]int)
	j := 0
	curr := -101

	for i := range nums {
		// _, found := hash[nums[i]]

		if curr != nums[i] {
			nums[j] = nums[i]
			curr = nums[i]
			j++

		}
	}

	return j + 1
}

func longestCommonPrefix(strs []string) string {
	result := ""
	curr_i := 0

	for {
		if len(strs[0]) == 0 {
			return ""
		}
		char_c := strs[0][curr_i]

		for i := range len(strs) {
			if len(strs[i]) == 0 {
				return ""
			}
			if curr_i < len(strs[i]) {
				val := strs[i][curr_i]
				if char_c != val {
					return result
				}
				println(string(val))
			} else {
				return result

			}
		}
		result += string(char_c)
		curr_i++
		println("result", curr_i, result)
		if curr_i >= len(strs[0]) {
			break
		}

	}

	return result
}

func isValid(s string) bool {
	// hash := make(map[rune]int)

	// hash['{'] = 0
	// hash['['] = 0
	// hash['('] = 0
	// hash['}'] = 0
	// hash[']'] = 0
	// hash[')'] = 0

	// for i := range s {
	// 	val := rune(s[i])
	// 	_, found := hash[val]
	// 	if found {
	// 		hash[val] += 1
	// 		println(val, hash[val])
	// 	}
	// }
	// if hash['{'] == hash['}'] && hash['['] == hash[']'] && hash['('] == hash[')'] {
	// 	return true
	// }

	// reimpliment with stack

	stack := make([]rune, len(s))
	point := 0

	for i, v := range s {
		r_val := rune(v)

		if r_val == '(' || r_val == '[' || r_val == '{' {
			stack[point] = r_val
			point++
			println("in")
		} else if r_val == ')' || r_val == ']' || r_val == '}' {
			if point != 0 {
				top := stack[point-1]
				switch top {
				case '(':
					if r_val == ')' {
						println("out ()")

						point--
					} else {
						return false
					}
				case '[':

					if r_val == ']' {
						println("out []")

						point--
					} else {
						return false
					}
				case '{':
					if r_val == '}' {
						println("out {}")

						point--
					} else {
						return false
					}
				}
			} else {
				return false
			}

		}

		if point != 0 {
			println(string(r_val), string(stack[point-1]), i, point)
		}
	}
	return point == 0
}

func maxDifference(s string) int {
	max_e := 101
	max_o := 0

	hash := make(map[rune]int)

	for i := range s {
		hash[rune(s[i])] += 1
	}
	for i := range hash {
		if hash[i]%2 == 0 && max_e > hash[i] {
			max_e = hash[i]
		} else if hash[i]%2 != 0 && max_o < hash[i] {
			max_o = hash[i]
		}
		println(hash[i], max_e, max_o)
	}

	return max_o - max_e
}

func plusOne(digits []int) []int {
	n := len(digits)
	carry := 0
	sum := 1
	cpy := false
	for i := n - 1; i >= 0; i -= 1 {
		val := digits[i] + carry + sum
		if i == 0 && val > 9 {
			cpy = true
		}
		carry = val / 10
		val = val % 10
		println(val, carry)
		digits[i] = val
		sum = 0
	}

	if cpy {
		n := len(digits) + 1
		var slice = make([]int, n)
		for i := range digits {
			if i == 0 {
				slice[i] = carry
			} else {
				slice[i] = digits[i-1]
			}
		}
		return slice
	}

	return digits
}

func lengthOfLastWord(s string) int {
	n := len(s)
	canskip := true
	count := 0
	for i := n - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if canskip {
				continue
			} else {
				return count
			}
		} else {
			count++
			canskip = false

		}
	}
	return count
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	prev := &ListNode{}
	res := &ListNode{}
	// curr := & ListNode{}

	if list1.Val >= list2.Val {
		res = list2
	} else {
		res = list1
	}

	// curr = res

	for list1 != nil || list2 != nil {
		if list1.Val >= list2.Val {

			prev.Next = list2
			list2 = list2.Next

		} else {
			prev = list1
			list1 = list1.Next
		}
	}

	// for list1.Next != nil || list2.Next != nil {
	// 	if list1.Val <= list2.Val {
	// 		head.Next = list1
	// 		head = head.Next
	// 		list1 = list1.Next
	// 	} else {
	// 		head.Next = list2
	// 		head = head.Next
	// 		list2 = list2.Next
	// 	}
	// }
	// for list1.Next != nil {
	// 	head.Next = list1
	// 	head = head.Next
	// 	list1 = list1.Next
	// }
	// for list2.Next != nil {

	// 	head.Next = list2
	// 	head = head.Next
	// 	list2 = list2.Next
	// }

	//  in memmory

	return res.Next

}

func maximumDifference(nums []int) int {
	max_diff := -1
	min := nums[0]

	for _, val := range nums {
		if min > val {
			min = val
		}
		if min < val {
			diff := val - min
			if diff > max_diff {
				max_diff = diff
			}
		}
	}
	return max_diff
}
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := ((high - low) / 2) + low
	for low <= high {
		println(mid, low, high)
		val := nums[mid]
		if val == target {
			return mid
		} else if val > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

		mid = ((high - low) / 2) + low

	}
	return mid
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0

	for i < m {
		if i >= m || j >= n {
			break
		}
		if nums1[i] == 0 {
			break
		} else if nums1[i] > nums2[j] {
			temp := nums1[i]
			nums1[i] = nums2[j]
			nums2[j] = temp
			if !(j+1 >= n) && nums2[j] > nums2[j+1] {
				j++
			}
		}
		i++
	}

	for j < n {
		nums1[i] = nums2[j]
		j++
		i++
	}
	i, j = 0, 0

	for i < m {
		if i >= m || j >= n {
			break
		}
		if nums1[i] == 0 {
			break
		} else if nums1[i] > nums2[j] {
			temp := nums1[i]
			nums1[i] = nums2[j]
			nums2[j] = temp
			if !(j+1 >= n) && nums2[j] > nums2[j+1] {
				j++
			}
		}
		i++
	}

	for j < n {
		nums1[i] = nums2[j]
		j++
		i++
	}

}

// func main() {
// 	// data := "world"
// 	// num := 1

// 	// if num > 1 {

// 	// 	println("hello " + data)
// 	// } else {
// 	// 	println("hello " + fmt.Sprint(num))
// 	// }

// 	// n := 3

// 	// result, err := sum(n)

// 	// switch {
// 	// case err != nil:
// 	// 	println(err.Error())
// 	// case result > 10:
// 	// 	println("sum data baka", result)
// 	// default:
// 	// 	println(result)

// 	// // }
// 	// val := twoSum([]int{3, 2, 4}, 6)
// 	// for v := range val {
// 	// 	println(val[v])
// 	// }

// 	// isPalindrome(121)
// 	// test := []string{"flower", "flow", "flight"}
// 	// result := longestCommonPrefix(test)
// 	// println(result)

// 	// result := isValid("(()))")
// 	// println(result)
// 	// result := maxDifference("mmsmsym")
// 	// println(result)
// 	// input := []int{9}
// 	// val := plusOne(input)
// 	// for i := range val {
// 	// 	println(val[i])
// 	// }
// 	// result := lengthOfLastWord("   fly me   to   the moon  ")
// 	// println(result)
// 	searchInsert([]int{1, 3, 5, 6, 9, 11, 13, 15, 16, 17}, 5)
// }
