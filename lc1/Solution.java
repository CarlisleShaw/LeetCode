import java.util.HashMap;
import java.util.*;

class Solution {
    public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < nums.length; i++) {
        map.put(nums[i], i);
    }
    System.out.println(map);
    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (map.containsKey(complement) && map.get(complement) != i) {
            return new int[] { i, map.get(complement) };
        }
    }
    throw new IllegalArgumentException("No two sum solution");
}
    public static void main(String[] args) {
        int a[] = {3, 3};
        int ans[] = {0,0};
        Solution s = new Solution();
        ans = s.twoSum(a,6);
        System.out.println(ans[0]);
        System.out.println(ans[1]);
    }
}