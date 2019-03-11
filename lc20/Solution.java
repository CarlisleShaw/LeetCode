package LeetCode.lc20;

import java.util.*;
import java.lang.*;
public class Solution {
    Map<Character,Character> m;
    Solution() {
        m = new HashMap<Character,Character>();
        m.put('(', ')');
        m.put('{', '}');
        m.put('[', ']');
    }
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<Character>();
        for(int i = 0; i < s.length(); i++) {
            if(m.get(s.charAt(i)) == null) {
                if(stack.empty()) {
                    return false;
                } else {
                    if(m.get(stack.peek()) == s.charAt(i)) {
                        stack.pop();
                    } else {
                        return false;
                    }
                }
            } else {
                stack.push(s.charAt(i));
            }
        }
        if(!stack.empty()) {
            return false;
        }
        return true;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.isValid("{()}"));
    }
}