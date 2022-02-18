/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
    const n = head.length;

    if (n === 0) {
        return [];
    }

    if (n === 1) {
        return head;
    }

    for (let i = 0; i < n; i += 2) {
        if (i - 1 === n / 2) {
            res[i] = head[i];
            continue;
        }

        res[i] = head[i + 1];
        res[i + 1] = head[i];
    }

    return res;
};
