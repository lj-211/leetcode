package solutions

import (
	"reflect"
)

/* c++ code
class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (head == NULL) {return NULL;}
        Node* start = head;
        while (start != NULL) {
            Node* copy = new Node(start->val, NULL, NULL);
            Node* cur = start;
            start = start->next;
            cur->next = copy;
            copy->next = start;
        }

        start = head;
        while (start != NULL) {
            if (start->random != NULL) {
                start->next->random = start->random->next;
            }

            start = start->next->next;
        }

        Node* l1 = head;
        Node* l2 = head->next;
        Node* ret = l2;
        for (l1=head; l1 != NULL; l1 = l1->next) {
            l2 = l1->next;
            l1->next = l2->next;
            if (l2->next != NULL) {
                l2->next = l2->next->next;
            }
        }

        return ret;
    }
};
*/

func noGoImplement() {
}

func init() {
	desc := `
A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a deep copy of the list.

Example 1:
Input:
{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}

Explanation:
Node 1's value is 1, both of its next and random pointer points to Node 2.
Node 2's value is 2, its next pointer points to null and its random pointer points to itself.
	`
	sol := Solution{
		Title:  "Copy List with Random Pointer",
		Desc:   desc,
		Method: reflect.ValueOf(noGoImplement),
		Tests:  make([]TestCase, 0),
	}

	SolutionMap["0138"] = sol
}
