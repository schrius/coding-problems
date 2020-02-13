/**
 * Given a singly linked list and an integer k, 
 * remove the kth last element from the list. 
 * k is guaranteed to be smaller than the length of the list.
 * The list is very long, so making more than one pass is prohibitively expensive.
 * Do this in constant space and in one pass.
 */

 public class ListNode{
     int value;
     ListNode next;
     ListNode(int value){
        this.value = value;
     }
     public ListNode removeNthFromEnd(ListNode head, int k){
         ListNode pre1 = head;
         LIstNode pre2 = head;
            
            for(int i = 0; i <= k; i++){
                pre1 = pre1.next;
            }

            while(pre1 != null){
                pre1 = pre1.next;
                pre2 = pre2.next;
            }
            pre2.next = pre2.next.next;
            return head;
     }
 }