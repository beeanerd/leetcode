'''
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or
false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.



Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
'''


class Solution:
    def checkInclusionBruteForce(self, s1: str, s2: str) -> bool:
        s1length = len(s1)
        sortedS1 = sorted(s1)
        for i in range(len(s2) - len(s1)):
            if sortedS1 == sorted(s2[i, s1length]):
                return True
        return False

    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        s1Map = [0] * 26
        s2Map = [0] * 26

        for i in range(len(s1)):
            s1Map[ord(s1[i]) - ord('a')] += 1
            s2Map[ord(s2[i]) - ord('a')] += 1

        matching = 0
        for letter in range(26):
            matching += 1 if s1Map[letter] == s2Map[letter] else 0

        for idx in range(len(s2) - len(s1)):
            if matching == 26:
                return True

            ind = ord(s2[idx]) - ord('a')
            s2Map[ind] -= 1
            if s2Map[ind] == s1Map[ind]:
                matching += 1
            elif s2Map[ind] == s1Map[ind] - 1:
                matching -= 1

            ind = ord(s2[idx + len(s1)]) - ord('a')
            s2Map[ind] += 1
            if s2Map[ind] == s1Map[ind]:
                matching += 1
            elif s2Map[ind] == s1Map[ind] + 1:
                matching -= 1

        return matching == 26
