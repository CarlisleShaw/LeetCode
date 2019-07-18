import java.util.ArrayDeque;
import java.util.Deque;

public class MyQueue {
    Deque<Integer> stack = new ArrayDeque<>();
    /** Initialize your data structure here. */
    public MyQueue() {
        
    }
    
    /** Push element x to the back of queue. */
    public void push(int x) {
        stack.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    public int pop() {
        Deque<Integer> stack2 = new ArrayDeque<>();
        while (!stack.isEmpty()) stack2.push(stack.pop());
        try{
            return stack2.pop();
        } finally {
            while (!stack2.isEmpty()) {
                stack.push(stack2.pop());
            }
        }
        
    }
    
    /** Get the front element. */
    public int peek() {
        Deque<Integer> stack2 = new ArrayDeque<>();
        while (!stack.isEmpty()) stack2.push(stack.pop());
        try{
            return stack2.peekFirst();
        } finally {
            while (!stack2.isEmpty()) {
                stack.push(stack2.pop());
            }
        }
    }
    
    /** Returns whether the queue is empty. */
    public boolean empty() {
        return stack.isEmpty();
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = new MyQueue();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.peek();
 * boolean param_4 = obj.empty();
 */