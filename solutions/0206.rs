// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    #[allow(dead_code)]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution {}

impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }

        match head.as_ref() {
            Some(node) => {
                if node.next.is_none() {
                    return head;
                }
            }
            None => {}
        }

        let mut pre = None;
        let mut cur = head;

        while let Some(mut node) = cur {
            let next = node.next;
            node.next = pre;
            pre = Some(node);
            cur = next;
        }

        pre
    }
}

fn main() {
    println!("problem 206:");
    println!("-------------------------------");

    let mut input = Some(Box::new(ListNode::new(1)));
    match input.as_mut() {
        Some(mut node) => {
            node.next = Some(Box::new(ListNode::new(2)));
        }
        None => {}
    }

    println!("input: {:?}", input);

    let mut output = Solution::reverse_list(input);
    println!("output: {:?}", output);
    while let Some(node) = output {
        println!("node: {}", node.val);
        output = node.next;
    }
}
