def divide(x, y):
    a = int(x / y)
    b = int(x % y)
    return (a, b)


def intToRoman(num):
    symbols = ['M', 'D', 'C', 'L', 'X', 'V', 'I']
    values = [1000, 500, 100, 50, 10, 5, 1]
    powers = [0] * len(values)

    for i, value in enumerate(values):
        n, num = divide(num, value)
        powers[i] = n

    s = ''
    for i in range(len(powers)):
        if i % 2 != 0:
            if powers[i+1] == 4:
                s += symbols[i+1]
                powers[i+1] = 0
                if powers[i] == 1:
                    s += symbols[i-1]
                else:
                    s += symbols[i]
            else:
                s += symbols[i] * powers[i]
        else:
            s += symbols[i] * powers[i]
    return s


def test(input, expect):
    actual = intToRoman(input)
    if actual != expect:
        print("fail! expect %s, got %s" % (expect, actual))
    else:
        print("pass")


test(3, "III")
test(4, "IV")
test(9, "IX")
test(40, "XL")
test(90, "XC")
test(400, "CD")
test(900, "CM")
