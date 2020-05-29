def maxArea(height):
    i, j, capacity = 0, len(height) - 1, 0
    while i < j:
        capacity = max(capacity, min(height[i], height[j]) * (j - i))
        if height[i] < height[j]:
            i += 1
        else:
            j -= 1
    return capacity


def test(expect, input):
    actual = maxArea(input)
    if expect != actual:
        print("fail! expect %d, got %d" % (expect, actual))
    else:
        print("pass")


test(49, [1, 8, 6, 2, 5, 4, 8, 3, 7])
