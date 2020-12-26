package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"regexp"
	"strings"
)

// 手机号正则
const (
	regular = "^(13|14|15|16|17|18|19)\\d{9}$"
)

/**
 * 正则匹配是否手机号
 */
func validate(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

/**
 * 读取excel的手机号码
 */
func readExcel(excelPath string) []string {

	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}

	rows := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))

	result := make([]string, 0)

	for _, row := range rows {
		newstr := strings.Replace(row[0], " ", "", -1)
		if !validate(newstr) {
			fmt.Println("跳过，不是手机号码:", row[0], newstr)
			continue
		}
		result = append(result, newstr)
	}

	newResult := removeDuplicateElement(result)

	fmt.Println("读取excel有效手机号码总数:", len(newResult))
	return newResult
}

/**
 * 字符串数组剔重
 */
func removeDuplicateElement(languages []string) []string {

	fmt.Println("开始剔除重复数据,数组总数长度:", len(languages))

	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		} else {
			fmt.Println("重复的数据:", item)
		}
	}

	fmt.Println("结束剔除重复数据,数组总数长度:", len(result))

	return result
}

/**
 * 逐行覆盖写txt文件
 */
func writerTXT(filename string, content []string) error {

	fmt.Println("需要写入的手机号数据长度:", len(content))

	// 写入文件
	// 判断文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}

	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	defer fd.Close()
	if err != nil {
		fmt.Println("写txt文件异常", err)
		return err
	}

	w := bufio.NewWriter(fd)
	for _, item := range content {

		//加密md5保存
		fmt.Fprintln(w, str2MD5(item))

		//不加密md5保存明文手机号
		//fmt.Fprintln(w, item)
	}

	w.Flush()
	fd.Sync()
	return nil
}

/**
 * 字符串转成md5
 */
func str2MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

/**
 * main函数
 */
func main() {
	// excel路径
	exlPath := "./test.xlsx"

	// 输出的文本路径
	desTxtPath := "./result.txt"

	// 获取待解析的excel路径，需要执行时传入参
	if len(os.Args) > 1 {
		exlPath = os.Args[1]
		fmt.Println("需要读取的手机号文件路径是:", exlPath)
	} else {
		fmt.Println("需要读取的excel文件路径和文件名称")
	}

	// 解析excel并生成数组
	result := readExcel(exlPath)

	// 生成md5文件
	err := writerTXT(desTxtPath, result)
	if err == nil {
		fmt.Println("处理完成，生成文件地址:", desTxtPath)
	}
}
