import java.util.*;
import java.lang.*;
public class Solution {
    Map<Character,Character> map = new HashMap<>();
    Solution() {
        map.put('(', ')');
        map.put('[', ']');
        map.put('{', '}');
    }
    public boolean isValid(String s) {
        char[] charArray = s.toCharArray();
        ArrayDeque<Character> stack = new ArrayDeque<>();
        for (Character c : charArray) {
            if (map.containsKey(c)) {
                stack.push(c);
            } else {
                // CAUTION: If stack is empty, peek() will throw EmptyStackException.
                if (stack.isEmpty()) return false; 
                else if (map.get(stack.peekFirst()).equals(c)) stack.pop();
                else return false;
            }
        }
        // CAUTION: For example, "({[]" will come here and need to be handled.
        if (!stack.isEmpty()) return false;
        return true;
    }
}