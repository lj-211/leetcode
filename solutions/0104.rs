// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if let Some(node) = root {
            let node_raw = node.borrow();
            return max_i32(
                Solution::max_depth(node_raw.left.clone()),
                Solution::max_depth(node_raw.right.clone()),
            ) + 1;
        }

        0
    }
}

fn max_i32(a: i32, b: i32) -> i32 {
    if a > b {
        return a;
    }

    b
}
