package tool

// 全角字符串转换半角字符串
// 参数：str：字符串
// 返回：转换半角后的字符串
// 注意：并不是所有的全角字符都能被转换为半角字符，例如汉字是全角字符，占2个字符的位置，但它无法被转换；只有英文字母、数字键、符号键等才能可以做全角和半角之间的转换
func full2Half(str string) (result string) {
	for _, charCode := range str {
		insideCode := charCode
		if insideCode == 12288 {
			insideCode = 32
		} else {
			insideCode -= 65248
		}

		if insideCode < 32 || insideCode > 126 {
			result += string(charCode)
		} else {
			result += string(insideCode)
		}
	}
	return result
}
