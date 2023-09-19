from collections import defaultdict


'''
You are given a string s and an integer k. You can choose any character of the
string and change it to any other uppercase English character. You can perform
this operation at most k times.

Return the length of the longest substring containing the same letter you can
get after performing the above operations.



Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achive this answer too.


Constraints:

1 <= s.length <= 10^5
s consists of only uppercase English letters.
0 <= k <= s.length
'''


class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        def findFreqLetter(freqChart):
            max_count = 0
            for key in freqChart:
                max_count = max(max_count, freqChart[key])
        count = defaultdict(int)
        left, right = 0, 0
        res = 0
        while right < len(s):
            count[s[right]] += 1
            if 1 + right - left - findFreqLetter(count) > k:
                count[s[left]] -= 1
                left += 1
            res = max(res, (right - left + 1))
            right += 1
        return res

    def characterReplacementOptimized(self, s: str, k: int) -> int:
        def findFreqLetter(freqChart):
            max_count = 0
            for key in freqChart:
                max_count = max(max_count, freqChart[key])
        count = defaultdict(int)
        left, right = 0, 0
        res = 0
        curr_max = 0
        while right < len(s):
            count[s[right]] += 1
            if count[s[right]] > curr_max:
                curr_max = count[s[right]]
            if 1 + right - left - curr_max > k:
                count[s[left]] -= 1
                left += 1
            res = max(res, (right - left + 1))
            right += 1
        return res
