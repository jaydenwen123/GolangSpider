# LeetCode每日一题:数组拆分 I（No.561） #

### 题目：数组拆分 I ###

` 给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。 复制代码`

### 示例： ###

` 输入: [1,4,3,2] 输出: 4 解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4). 复制代码`

### 思考： ###

` 这道题先将数组排序，再从下标0开始，间隔相加，即为结果。 复制代码`

### 实现： ###

` class Solution { public int arrayPairSum(int[] nums) { int sum = 0; Arrays.sort(nums); for(int i=0;i<nums.length;i+=2) sum += nums[i]; return sum; } } 复制代码`