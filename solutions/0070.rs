struct Solution {}

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        // dp = dp[i-1] + dp[i-2]
        // 3 ..=n
        // dp[1] = 1 dp[2] = 2
        if n <= 1 {
            return n;
        }
        let mut dp_pre = 1;
        let mut dp = 2;

        for i in 3..=n {
            let pre = dp;
            dp = dp + dp_pre;
            dp_pre = pre;
        }

        dp
    }
}
