class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        results = []
        list1 = []
        num = 0
        for i in range(len(s)):
            num = i
            while True:
                if num == len(s) or s[num] in list1:
                    results.append(len(list1))
                    list1.clear()
                    break
                list1.append(s[num])
                num += 1

        return max(results) if results else 0

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        results, list1 = [], []
        num, result= 0, 0
        for i in range(len(s)):
            if i != 0:
                list1.remove(s[i-1])
            while num < len(s) and s[num] not in list1:
                list1.append(s[num])
                num += 1
            result = max(result, num-i)
        return result


        