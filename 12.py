def divide(x, y):
	a = int(x / y)
	b = int(x % y)
	return (a, b)

class Solution:
	def intToRoman(self, num:int) -> str:
		symbols:List = ['M', 'D', 'C', 'L', 'X', 'V', 'I']
		values:List = [1000, 500, 100, 50, 10, 5, 1]
		powers:List = [0] * len(values)
		
		for i in range(len(values)):
			n, num = divide(num, values[i])
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

	def test(self, input, expect):
		actual = self.intToRoman(input)
		if actual != expect:
			print("Input: ", input)
			print("Output: ", actual)
			print("Expect: ", expect)
			print("==========")

def main():
	solution = Solution()
	solution.test(3, "III")
	solution.test(4, "IV")
	solution.test(9, "IX")
	solution.test(40, "XL")
	solution.test(90, "XC")
	solution.test(400, "CD")
	solution.test(900, "CM")

if __name__ == '__main__':
    main()