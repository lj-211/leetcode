// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn swap_pairs(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        if let Some(mut one) = head {
            if let Some(mut two) = one.next {
                one.next = Solution::swap_pairs(two.next);
                two.next = Some(Box::new(*one));

                return Some(Box::new(*two));
            } else {
                return Some(Box::new(*one));
            }
        }

        return None;
    }
}
