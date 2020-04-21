/*
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 *  [ 1, 2, 3 ],
 *  [ 4, 5, 6 ],
 *  [ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 * 示例 2:
 *
 * 输入:
 * [
 *   [1, 2, 3, 4],
 *   [5, 6, 7, 8],
 *   [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 */

struct Solution {}

impl Solution {
    fn traverse_circle(
        matrix: &Vec<Vec<i32>>,
        i: usize,
        j: usize,
        m: usize,
        n: usize,
        output: &mut Vec<i32>,
    ) {
        if m == 1 {
            for v in 0..n {
                output.push(matrix[i][j + v]);
            }
            return;
        } else if n == 1 {
            for v in (0..m) {
                output.push(matrix[i + v][j]);
            }
            return;
        }

        let path_len = (m + n) * 2 - 4;
        let dest_one = n;
        let dest_two = m + n - 1;
        let dest_thr = m + n * 2 - 2;

        for d in 0..path_len {
            let val: i32;
            if d < dest_one {
                val = matrix[i][j + d]
            } else if d >= dest_one && d < dest_two {
                val = matrix[i + d - dest_one + 1][j + n - 1];
            } else if d >= dest_two && d < dest_thr {
                val = matrix[i + m - 1][j + n - 1 - (d - dest_two + 1)];
            } else if d >= dest_thr && d < path_len {
                val = matrix[i + m - 1 - (d - dest_thr + 1)][j];
            } else {
                continue;
            }

            output.push(val);
        }

        if m > 2 && n > 2 {
            Solution::traverse_circle(matrix, i + 1, j + 1, m - 2, n - 2, output);
        }
    }

    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let m = matrix.len();
        if m == 0 {
            return Vec::new();
        }
        let n = matrix[0].len();

        let mut spiral_order_vec = Vec::with_capacity(m * n);

        Solution::traverse_circle(&matrix, 0, 0, m, n, &mut spiral_order_vec);

        spiral_order_vec
    }
}

fn main() {
    println!("0054: ");

    //let input = vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]];
    let input = vec![vec![1, 2], vec![4, 5], vec![7, 8]];
    //let input = vec![vec![1], vec![4], vec![7]];
    //let input = vec![vec![1, 2], vec![3, 4]];
    let output = Solution::spiral_order(input);
    println!("spriral-order: {:?}", output);
}
