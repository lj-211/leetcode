use std::collections::HashMap;

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

        for _i in 3..=n {
            let pre = dp;
            dp = dp + dp_pre;
            dp_pre = pre;
        }

        dp
    }

    pub fn climb_stairs_v1(n: i32) -> i32 {
        let mut cache = HashMap::new();

        return Solution::internal_climb_staires(n, &mut cache);
    }

    fn internal_climb_staires(n: i32, cache: &mut HashMap<i32, i32>) -> i32 {
        if n <= 2 {
            cache.insert(n, n);
            return n;
        }

        let one: i32;
        let two: i32;

        if let Some(v) = cache.get(&(n - 1)) {
            one = *v;
        } else {
            one = Solution::internal_climb_staires(n - 1, cache);
        }

        if let Some(v) = cache.get(&(n - 2)) {
            two = *v;
        } else {
            two = Solution::internal_climb_staires(n - 2, cache);
        }

        return one + two;
    }
}

fn main() {
    println!("problem 70:");
    println!("-------------------------------");

    let input = 4;
    let output = Solution::climb_stairs_v1(input);
    let output_2 = Solution::climb_stairs(input);
    println!("v1: {}", output);
    println!("{}", output_2);
}
