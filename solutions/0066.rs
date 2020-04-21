/*
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 */

impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let size = digits.len();

        let mut ret = vec![0; size];

        let mut step = 1;
        for i in (0..size).rev() {
            let sum = step + digits[i];
            ret[size - 1 - i] = sum % 10;
            step = sum / 10;
        }
        // NOTE: 这里有个隐藏特性
        // 如果step > 0，则必然是10000000
        if step > 0 {
            ret.push(step);
        }

        ret.reverse();

        ret
    }
}
