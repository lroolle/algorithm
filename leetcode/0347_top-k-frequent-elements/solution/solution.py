import collections
import heapq
from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        table = collections.Counter(nums)
        heap, ret = list(), [0] * k
        for i, n in table.items():
            heapq.heappush(heap, (n, i))
            if len(heap) > k:
                heapq.heappop(heap)

        for i in range(k):
            ret[k - i - 1] = heapq.heappop(heap)[1]
        return ret


if __name__ == "__main__":
    example1 = ([1, 1, 1, 2, 2, 3], 2)
    ans1 = Solution().topKFrequent(example1[0], example1[1])
    print(ans1)
