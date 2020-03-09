struct Solution {}

impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut ret: Vec<Vec<i32>> = Vec::new();

        let get_idx_value = |datas: &Vec<Vec<i32>>, i: usize, j: usize| -> i32 {
            if i == 0 {
                return 1;
            }

            let sum: i32;
            if j == 0 {
                sum = datas[i - 1][j];
            } else if j == i {
                sum = datas[i - 1][j - 1];
            } else {
                sum = datas[i - 1][j - 1] + datas[i - 1][j];
            }

            sum
        };

        for i in 0..num_rows {
            let mut line_vec: Vec<i32> = Vec::new();

            for j in 0..=i {
                line_vec.push(get_idx_value(&ret, i as usize, j as usize));
            }

            ret.push(line_vec);
        }

        return ret;
    }
}

fn main() {
    println!("problem 118: ");
    println!("-----------------------------------");

    let input = 5;
    let output = Solution::generate(input);

    for v in &output {
        println!("{:?}", v);
    }
}
