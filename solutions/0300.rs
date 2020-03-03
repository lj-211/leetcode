//use std::cmp;
/*
#[derive(Debug)]
pub struct SolStruct {
    pub name: String,
}

pub trait Solution {
    fn length_of_lis(nums: Vec<i32>) -> i32;
}

impl Solution for SolStruct {
    fn length_of_lis(nums: Vec<i32>) -> i32 {
        for v in nums {
            println!("{} - {}", 0, v)
        }

        0
    }
}
*/

struct Solution {}

impl Solution {
    #[allow(dead_code)]
    fn length_of_list_binary(nums: &Vec<i32>) -> i32 {
        let size = nums.len();
        if size < 2 {
            return size as i32;
        }

        let mut tails = vec![0; nums.len()];
        let mut max_len = 0;
        tails[0] = nums[0];

        for i in 0..nums.len() {
            let v = nums[i];

            // 找到第一个比他大的值替换
            let mut l = 0;
            let mut r = max_len;
            while l <= r {
                let m = (l + r) / 2;
                if v > tails[m] {
                    l = m + 1
                } else {
                    if m < 1 {
                        break;
                    }
                    r = m - 1
                }
            }

            tails[l] = v;
            if l > max_len {
                max_len = l;
            }

            println!("l: {} {}", l, max_len);

            println!("{:?}", tails);
        }

        (max_len + 1) as i32
    }

    #[allow(dead_code)]
    fn length_of_lis(nums: &Vec<i32>) -> i32 {
        let size = nums.len();
        if size < 2 {
            return size as i32;
        }

        let mut dp = vec![0; nums.len()];
        dp[0] = 1;
        let mut max_len: i32 = 0;

        for i in 0..nums.len() {
            let mut cur_dp: i32;
            let mut max_dp: i32 = 1;

            for j in 0..i {
                if nums[j] < nums[i] {
                    cur_dp = dp[j] + 1
                } else {
                    cur_dp = 1
                }

                if cur_dp > max_dp {
                    max_dp = cur_dp
                }
            }

            dp[i] = max_dp;

            if dp[i] > max_len {
                max_len = dp[i];
            }

            println!("dp: {:?}", dp);
        }

        max_len
    }
}

fn main() {
    println!("problem {}:", 300);

    let input = vec![1, 2, 3, 1, 1, 4, 5];
    let out = Solution::length_of_lis(&input);

    println!("input: {:?}", input);
    println!("output: {}", out);
}
