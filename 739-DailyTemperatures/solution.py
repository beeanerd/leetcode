from collections import deque
from typing import List


'''
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to
wait after the ith day to get a warmer temperature. If there is no future day
for which this is possible, keep answer[i] == 0 instead.


Example 1:

Input: temperatures = [73, 74, 75, 71, 69, 72, 76, 73]

Stack

73, 1
74, 1
75, 4
76, 0
_____
Output: [1, 1, 4, 2, 1, 1, 0, 0]

Example 2:

Input: temperatures = [30, 40, 50, 60]
Output: [1, 1, 1, 0]

Example 3:

Input: temperatures = [30, 60, 90]
Output: [1, 1, 0]
'''


class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        tempStack = deque()
        differenceList = [0] * len(temperatures)
        for count, temperature in enumerate(temperatures):
            while tempStack and tempStack[-1][0] < temperature:
                temp, index = tempStack.pop()
                differenceList[index] = count - index
            tempStack.append((temperature, count))
        return differenceList
