class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

public class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        if (head == null || k == 1) {
            return head;
        }

        ListNode dummy = new ListNode(0);
        dummy.next = head;

        ListNode pre = dummy;
        ListNode cur = head;
        int count = 0;

        while (cur != null) {
            count++;
            if (count % k == 0) {
                pre = reverse(pre, cur.next);
                cur = pre.next;
            } else {
                cur = cur.next;
            }
        }

        return dummy.next;
    }

    private ListNode reverse(ListNode pre, ListNode next) {
        ListNode last = pre.next;
        ListNode cur = last.next;

        while (cur != next) {
            last.next = cur.next;
            cur.next = pre.next;
            pre.next = cur;
            cur = last.next;
        }

        return last;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        // 创建链表 1->2->3->4->5->6->7->8
        ListNode head = new ListNode(1);
        ListNode node = head;
        for (int i = 2; i <= 8; i++) {
            node.next = new ListNode(i);
            node = node.next;
        }

        // 应用翻转函数，k = 3
        head = solution.reverseKGroup(head, 3);

        // 打印翻转后的链表
        while (head != null) {
            System.out.print(head.val + "->");
            head = head.next;
        }
    }
}
