package main

import (
	"fmt"
	"sort"
)

type strArray []string

func (a strArray) Len() int      { return len(a) }
func (a strArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a strArray) Less(i, j int) bool {
	if len(a[i]) == len(a[j]) {
		for x := 0; x < len(a[i]); x++ {
			if a[i][x] == a[j][x] {
				continue
			} else {
				return a[i][x] < a[j][x]
			}
		}
	}
	return len(a[i]) > len(a[j])
}

func isVaild(s, str string) bool {
	i, j := 0, 0
	// 注意这种写法，优于两层循环
	for i < len(str) && j < len(s) {
		if str[i] == s[j] {
			i++
		}
		j++
	}
	return i == len(str)
}

func findLongestWord(s string, d []string) string {
	// 按长度和字母ASCII排序
	sort.Sort(strArray(d))
	for _, str := range d {
		if isVaild(s, str) {
			return str
		}
	}
	return ""
}

func main() {
	// s := "abpcplea"
	// s := "bab"
	// s := "aewfafwafjlwajflwajflwafj"
	s := "wsmzffsupzgauxwokahurhhikapmqitytvcgrfpavbxbmmzdhnrazartkzrnsmoovmiofmilihynvqlmwcihkfskwozyjlnpkgdkayioieztjswgwckmuqnhbvsfoevdynyejihombjppgdgjbqtlauoapqldkuhfbynopisrjsdelsfspzcknfwuwdcgnaxpevwodoegzeisyrlrfqqavfziermslnlclbaejzqglzjzmuprpksjpqgnohjjrpdlofruutojzfmianxsbzfeuabhgeflyhjnyugcnhteicsvjajludwizklkkosrpvhhrgkzctzwcghpxnbsmkxfydkvfevyewqnzniofhsriadsoxjmsswgpiabcbpdjjuffnbvoiwotrbvylmnryckpnyemzkiofwdnpnbhkapsktrkkkakxetvdpfkdlwqhfjyhvlxgywavtmezbgpobhikrnebmevthlzgajyrmnbougmrirsxi"
	// d := []string{"ale", "apple", "monkey", "plea"}
	// d := []string{"a", "b", "c"}
	// d := []string{"ba", "ab", "a", "b"}
	// d := []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}
	d := []string{"nbmxgkduynigvzuyblwjepn", "jpthiudqzzeugzwwsngebdeai"}
	fmt.Println(s)
	fmt.Println(d)
	ans := findLongestWord(s, d)
	fmt.Println(d)
	fmt.Println(ans)
}
