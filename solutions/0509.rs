struct Solution {}

impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }

        let mut dp_pre = 0;
        let mut dp = 1;

        for _i in 2..=n as usize {
            let cur = dp;
            dp = dp + dp_pre;
            dp_pre = cur;
        }

        dp
    }
}

fn main() {
    println!("problem 509: ");
    println!("--------------------------");

    let input = 2;
    let output = Solution::fib(input);

    println!("output: {}", output);
}
