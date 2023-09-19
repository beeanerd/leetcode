'''
Given two strings s and t of lengths m and n respectively, return the minimum
window substring of s such that every character in t (including duplicates) is
included in the window. If there is no such substring, return the empty string
"".

The testcases will be generated such that the answer is unique.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C'
from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
'''

'''
Brute Force Approach:
    First, add each element in main string to HashMap, then same for
    substring. Check if elements in substring are contained in it. If
    yes try every string size smaller. Almost like fibonnacci but with
    n-1 size of string. Kinda like mergesort partitioning.

    Procedural solution:
        tMap = dict containing all the letters & frequencies in t
        Starting from left
        Left Pointer:
            1. If letter is in tMap mark as such and expand window to next
            letter
                1.1 If letter is not in tMap move left side
                1.2 If right pointer encounters letter in tMap that is the same
                    as the letter left is on & now exceeds the frequency of the
                    letter in tMap move left until it reaches the next letter
                    in tMap
        Right Pointer:
            1. Continue until reaching the end of the string or minsubstring
            length is equal to length of the t-string (smallest possible)
            2. Once all letters are contained in the substring, check if length
            is
                less than the previously found substring's length

        How to store?
            1. Check each frequency of each letter O(26n) operation for add &
            subtract
                1.1 Optimizations:
                    1.1.1
'''


class Solution:
    def minWindow(self, s: str, t: str) -> str:
