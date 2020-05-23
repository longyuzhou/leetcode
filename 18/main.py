from collections import Counter


def fourSum(nums, target):
    return nSum(nums, target, 4)


def nSum(nums, target, n):
    res = {}

    def collector(path, total):
        if len(path) > 0 and total + path[len(path) - 1] * (n - len(path)) > target:
            return False
        if len(path) < n:
            return True
        if len(path) == n and total == target:
            res[tuple(path)] = True
        return False

    def dfs(start, path, total):
        if not collector(path, total) or start >= len(nums):
            return

        for i in range(start, len(nums)):
            path.append(nums[i])
            dfs(i + 1, path, total + nums[i])
            path.pop()

    nums.sort()
    dfs(0, [], 0)
    return res.keys()


print(fourSum([], 0))
print(fourSum([0, 0, 0, 0, 0], 0))
print(fourSum([1, 0, -1, 0, -2, 2], 0))
print(fourSum([-1, 0, 1, 2, -1, -4], -1))
print(fourSum([-479, -472, -469, -461, -456, -420, -412, -403, -391, -377, -362, -361, -340, -336, -336, -323, -315, -301, -288, -272, -271, -246, -244, -227, -226, -225, -210, -194, -190, -187, -183, -176, -167, -143, -140, -123, -120, -74, -60, -40, -39, -37, -34, -33, -29, -26, -23, -18, -17, -11, -9, 6, 8, 20, 29, 35, 45, 48, 58, 65, 122, 124, 127, 129, 145, 164, 182, 198, 199, 206, 207, 217, 218, 226, 267, 274, 278, 278, 309, 322, 323, 327, 350, 361, 372, 376, 387, 391, 434, 449, 457, 465, 488], 1979))
