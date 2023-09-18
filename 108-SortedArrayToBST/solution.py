class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# Plan
# 1. Find Center of array segment
# 2. Split into two subproblems
# 3. Recur ( find center of array segment )
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        node = self.findBranches(self, nums, 0, len(nums) - 1)
        

        
    def findBranches(self, nums: List[int], lower: int, upper: int) -> [TreeNode]:
        if lower >= upper:
            return TreeNode(nums[lower])
        center = (lower + upper) // 2
        lowerNode = self.findBranches(self, nums, lower, center - 1)
        upperNode = self.findBranches(self, nums, center + 1, upper)

        toReturn = TreeNode(nums[center], lowerNode, upperNode)
        return toReturn


    def inorder_traversal(self, nums: List[int]):

