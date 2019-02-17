package main
import "fmt"

func isAnagram(s string, t string) bool {
    var x, y [26]int
    for i := 0; i < len(s); i++ {
        x[s[i] - 'a']++
    }
    for i := 0; i < len(t); i++ {
        y[t[i] - 'a']++
    }
    for i := 0; i < len(x); i++ {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}

func main() {
    a := isAnagram("abcdeff", "gfedcba")
    b := isAnagram("abcdefg", "gfedcba")
    fmt.Println(a)
    fmt.Println(b)
}