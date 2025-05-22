class StringProblems:
    def __init__(self):
        pass


    # GET THE LONGEST COMMON PREFIX

    # PLAN
    # ["flower","flow","flight"]

    def getLongestPrefix(self, strs: list[str]) -> str:
        if len(strs) == 0:
            return ""
        elif len(strs) == 1:
            return strs[0]
        else:
            prefix = strs[0]
            for i in range(1, len(strs)):
                while length(prefix) !=  0:
                    if prefix == strs[i]:
                        break
                    else:
                        prefix = prefix[::-1]
            return prefix


    # LENGTH OF LAST WORD
    def lengthOfLastWord(self, s: str) -> int:
        if len(s) == 0:
            return 0
        elif len(s) == 1:
            return len(s)
        else:
            print(" ".join(s.split()), len("  ".join(s.split())))
            lastWord = " ".join(s.split())[-1]
            return len(lastWord)
        return 0



    # GET THE LARGEST NUMBER
    def getLargestNumber(self, nums: list[int]) -> int:
        strs = ''
        for num in range(len(nums)):
            strs += str(nums[num])

        print(strs)
        newNums = strs.split()

        for num in range(newNums):
            pass
        print(newNums)
        return -1