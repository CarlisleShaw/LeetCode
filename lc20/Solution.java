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
        Character[] charArray = s.toCharArray();
        Stack<Character> stack = new Stack<>();
        for (Character c : charArray) {
            if (map.containsKey(c)) {
                stack.add(c);
            } else {
                // CAUTION: If stack is empty, peek() will throw EmptyStackException.
                if (stack.empty()) return false; 
                else if (map.get(stack.peek()).equals(c)) stack.pop();
                else return false;
            }
        }
        // CAUTION: for example, "({[]" will come here and need to be handled.
        if (!stack.empty()) return false;
        return true;
    }
}