/*
 * 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 * 示例:
 *
 * 输入:
 * [
 *   [1,3,1],
 *   [1,5,1],
 *   [4,2,1]
 * ]
 * 输出: 7
 * 解释: 因为路径 1→3→1→1→1 的总和最小。
 */
struct Solution {}

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        println!("input: {:?}", grid);

        let m = grid.len();
        if m == 0 {
            return 0;
        }
        let n = grid[0].len();
        if n == 0 {
            let sum: i32 = grid[0].iter().sum();
            return sum;
        }

        let mut dp = vec![vec![0; n]; m];
        dp[0][0] = grid[0][0];

        for i in 0..m {
            for j in 0..n {
                if i == 0 && j == 0 {
                    continue;
                }
                //dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
                let mut up = std::i32::MAX;
                if j != 0 {
                    up = dp[i][j - 1];
                }
                let mut left = std::i32::MAX;
                if i != 0 {
                    left = dp[i - 1][j]
                }

                dp[i][j] = std::cmp::min(up, left) + grid[i][j];
            }
        }

        return dp[m - 1][n - 1];
    }
}

fn main() {
    println!("0064: 最小路径和");
    let input = vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1]];
    let output = Solution::min_path_sum(input);

    println!("output: {}", output);
}
