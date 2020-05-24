def groupAnagrams(strs):
    res = []
    memo = {}
    for str in strs:
        key = toKey(str)
        if key in memo:
            res[memo[key]].append(str)
        else:
            memo[key] = len(res)
            res.append([str])
    return res


def toKey(str):
    key = [0]*26
    for ch in str:
        idx = ord(ch)-ord('a')
        key[idx] = key[idx] + 1
    return tuple(key)


print(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))
