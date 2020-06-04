Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Solution:
输入：-1, 2, 3, -5, -2, 0, -1

(1) 分割成不含0的子列表
  -1, 2, 3, -5, -2
  1

(2) -1, 2, 3, -5, -2
  计算全部元素的乘积 => -60
  计算0至左侧第一个负数的乘积 => -1
  计算size-1至右侧第一个负数的乘积 => -2
  此子列表的对应解为max(-60, -60/-1, -60/2)，即60

(3) 对所有子列表执行(2)的操作

(4) 最终解为
  max(
    0,   // 因为列表中存在0
    60,  // 第一个子列表的解
    1    // 第二个子列表的解
  )
