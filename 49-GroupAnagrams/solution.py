class List:
    def __init__(self):
        print("hi")


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for word in strs:
            freqList = [0] * 26
            for s in word:
                freqList[ord(s.lower()) - ord('a')] += 1
            freqList = tuple(freqList)
            if freqList not in groups:
                groups[freqList] = []
            groups[freqList].append(word)
        toReturn = []
        for key in groups:
            newList = []
            for word in groups[key]:
                newList.append(word)
            toReturn.append(newList)
        return toReturn
