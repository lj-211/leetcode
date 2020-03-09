struct Solution {}

impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let index = row_index as usize;
        let mut ret = vec![0; index + 1];
        ret[index] = 1;

        let mut start: usize;
        for i in 1..=index {
            start = index - i;
            for j in start..=index {
                if j == index {
                    ret[j] = ret[j]
                } else if j == 0 {
                    ret[j] = ret[j + 1]
                } else {
                    ret[j] = ret[j] + ret[j + 1]
                }
            }
        }

        ret
    }
}

fn main() {
    println!("problem 119: ");
    println!("-----------------------------------");

    let input = 4;
    let output = Solution::get_row(input);

    println!("vec is: {:?}", output);
}
