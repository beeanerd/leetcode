from collections import defaultdict
'''    class Solution:
        def isAnagram(self, s: str, t: str) -> bool:
            return sorted(s) == (sorted(t)
'''


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        freq = defaultdict(int)
        if len(s) != len(t):
            return False

        for letter in s:
            freq[letter] += 1

        for letter in t:
            if letter in freq:
                freq[letter] -= 1
                if freq[letter] <= 0:
                    freq.pop(letter)
            else:
                return False
        return not freq
