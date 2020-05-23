package main

import "fmt"

func longestPalindrome(s string) string {
	if l := len(s); l == 0 {
		return ""
	} else if l == 1 || (l == 2 && s[0] == s[1]) {
		return s
	}

	from := 0
	to := 0

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			f, t := getPalindrome(s, i-1, i)
			if t-f > to-from {
				from = f
				to = t
			}
		}
		if i < len(s)-1 && s[i-1] == s[i+1] {
			f, t := getPalindrome(s, i-1, i+1)
			if t-f > to-from {
				from = f
				to = t
			}
		}
	}

	return s[from : to+1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func getPalindrome(s string, left, right int) (int, int) {
	offset := 0
	for (left-offset) >= 0 && (right+offset) < len(s) {
		x := s[left-offset]
		y := s[right+offset]
		if x != y {
			break
		}
		offset++
	}
	offset = offset - 1
	return max(left-offset, 0), min(right+offset, len(s)-1)
}

func test(input, expect string) {
	actual := longestPalindrome(input)
	if actual != expect {
		fmt.Println("Input: ", input)
		fmt.Println("Output: ", actual)
		fmt.Println("Expect: ", expect)
	}
}

func main() {
	test("", "")
	test("b", "b")
	test("ccc", "ccc")
	test("babad", "bab")
	test("abcdefg", "a")
	test("rgczcpratwyqxaszbuwwcadruayhasynuxnakpmsyhxzlnxmdtsqqlmwnbxvmgvllafrpmlfuqpbhjddmhmbcgmlyeypkfpreddyencsdmgxysctpubvgeedhurvizgqxclhpfrvxggrowaynrtuwvvvwnqlowdihtrdzjffrgoeqivnprdnpvfjuhycpfydjcpfcnkpyujljiesmuxhtizzvwhvpqylvcirwqsmpptyhcqybstsfgjadicwzycswwmpluvzqdvnhkcofptqrzgjqtbvbdxylrylinspncrkxclykccbwridpqckstxdjawvziucrswpsfmisqiozworibeycuarcidbljslwbalcemgymnsxfziattdylrulwrybzztoxhevsdnvvljfzzrgcmagshucoalfiuapgzpqgjjgqsmcvtdsvehewrvtkeqwgmatqdpwlayjcxcavjmgpdyklrjcqvxjqbjucfubgmgpkfdxznkhcejscymuildfnuxwmuklntnyycdcscioimenaeohgpbcpogyifcsatfxeslstkjclauqmywacizyapxlgtcchlxkvygzeucwalhvhbwkvbceqajstxzzppcxoanhyfkgwaelsfdeeviqogjpresnoacegfeejyychabkhszcokdxpaqrprwfdahjqkfptwpeykgumyemgkccynxuvbdpjlrbgqtcqulxodurugofuwzudnhgxdrbbxtrvdnlodyhsifvyspejenpdckevzqrexplpcqtwtxlimfrsjumiygqeemhihcxyngsemcolrnlyhqlbqbcestadoxtrdvcgucntjnfavylip", "qgjjgq")
}
