from typing import List
'''
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and
you may not use the same element twice.

You can return the answer in any order.

[Runtime] 64ms - Beats 82.49%
Aug 29, 2023
'''


def twoSum(self, nums: List[int], target: int) -> List[int]:
    val = set(nums)
    for count, i in enumerate(nums):
        diff = target - i
        if diff in val:
            secondIndex = nums.index(diff)
            if count != secondIndex:
                return [count, secondIndex]
