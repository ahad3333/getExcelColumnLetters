package excel

func getExcelColumnLetters(n int) []string {
    letters := []string{}
    for n > 0 {
        n--
        letters = append([]string{string(n%26 + 65)}, letters...)
        n /= 26
    }
    return letters
}