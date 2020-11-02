from typing import List
from collections import deque


def maxSlidingWindow(nums: "List[int]", k: "int") -> "List[int]":  # base cases
    n = len(nums)
    if n * k == 0:
        return []
    if k == 1:
        return nums

    def clean_deque(i):
        # remove indexes of elements not from sliding window
        if deq and deq[0] == i - k:
            deq.popleft()

        # remove from deq indexes of all elements
        # which are smaller than current element nums[i]
        while deq and nums[i] > nums[deq[-1]]:
            deq.pop()

    # init deque and output
    deq = deque()
    max_idx = 0
    for i in range(k):
        clean_deque(i)
        deq.append(i)
        # compute max in nums[:k]
        if nums[i] > nums[max_idx]:
            max_idx = i
    output = [nums[max_idx]]
    print(deq)

    # build output
    for i in range(k, n):
        clean_deque(i)
        deq.append(i)
        print(deq)
        output.append(nums[deq[0]])
    return output


if __name__ == "__main__":
    nums = [1, 3, -1, -3, 5, 3, 6, 7]
    maxSlidingWindow(nums, 3)
