package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"md5Tools/config"
	"os"
	"path"
	"strconv"
	"strings"
)

/**
 * 读取excel的数据
 */
func readExcel(excelPath string, md5Config *config.Md5ToolsConfig) [][]string {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}

	// 全部数据加载
	//rows, _ := xlsx.GetRows(md5Config.InSheetName)
	//return rows

	// 申明新的二维数组
	rows, _ := xlsx.GetRows(md5Config.InSheetName)
	var newRows = make([][]string, len(rows))
	for i := 0; i < len(rows); i++ {
		newRows[i] = make([]string, len(md5Config.Title))
	}

	// 空的excel无法脱敏
	if len(rows) == 0 {
		fmt.Println("待脱敏的excel是空的，请检查")
		os.Exit(1)
	}

	// [][]第二个下标
	var idx int
	// 列迭代模式
	cols, err := xlsx.Cols(md5Config.InSheetName)
	if err != nil {
		fmt.Println(err)
	}
	for cols.Next() {
		col, err := cols.Rows()
		if err != nil {
			fmt.Println(err)
		}

		// 判断title是否需要脱敏的
		var breakFlag bool
		for _, element := range md5Config.Title {
			if element == col[0] {
				breakFlag = true
			}
		}

		// 重新给二维数组赋值列数据
		if breakFlag {
			for i, v := range col {
				newRows[i][idx] = v
			}
			idx++
		}
	}

	// 校验没读取到脱敏数据，可能空的 excel 或者 title 匹配不上
	//if newRows[0][0] != md5Config.Title[0] || newRows[0][1] != md5Config.Title[1] {
	//	fmt.Println("没有匹配到需要脱敏的列标题，请检查excel")
	//	os.Exit(1)
	//}
	if len(newRows) == 0 || newRows[0][0] == "" || newRows[0][1] == "" {
		fmt.Println("没有匹配到需要脱敏的列标题，请检查excel")
		os.Exit(1)
	}

	return newRows
}

/**
 * 写excel到新的 sheet 页
 */
func writeExcelResult2Sheet (filePath string, content [][]string, md5Config *config.Md5ToolsConfig) {
	var xlsx *excelize.File
	var err error
	// 比较输入和输出 sheet 名称一样就新创建一个 xlsx
	if md5Config.InSheetName != md5Config.OutSheetName {
		// 继续在原来的 excel 操作
		xlsx, err = excelize.OpenFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// 新创建
		xlsx = excelize.NewFile()
	}

	//创建新表单
	index := xlsx.NewSheet(md5Config.OutSheetName)

	// 写入读取的 excel 二维数组的md5值
	for i, row := range content {
		for j, colCell := range row {
			cellIndex :=  changIndexToAxis(i, j)
			modifyExcelCellByAxis(xlsx, md5Config, cellIndex, colCell)
		}
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	if md5Config.InSheetName != md5Config.OutSheetName {
		err = xlsx.SaveAs(filePath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//获取文件名称带后缀
		fileNameWithSuffix := path.Base(filePath)
		//获取文件的后缀(文件类型)
		fileType := path.Ext(fileNameWithSuffix)
		//获取文件名称(不带后缀)
		fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
		// 存放的新文件名
		newFile := path.Base(fileNameOnly + "_result" + fileType)
		err = xlsx.SaveAs(newFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}

/**
 * 数据坐标转换成 excel 坐标
 */
func changIndexToAxis(intIndexX int, intIndexY int ) string{
	var arr = [...]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	intIndexY = intIndexY + 1
	resultY := ""
	for true {
		if intIndexY <= 26 {
			resultY = resultY + arr[intIndexY - 1]
			break
		}
		mo := intIndexY % 26
		resultY = arr[mo - 1] + resultY
		shang := intIndexY / 26
		if shang <= 26 {
			resultY = arr[shang - 1] + resultY
			break
		}
		intIndexY = shang
	}
	return resultY + strconv.Itoa(intIndexX + 1)
}

/**
 * 二维数组转成excel的坐标后往excel数据
 */
func modifyExcelCellByAxis( xlsx *excelize.File, md5Config *config.Md5ToolsConfig, axis string, value interface{}) int{
	var cellValue string

	str := value.(string)
	trimStr := strings.Trim(str, " ")
	halfStr := full2Half(trimStr)

	for _, element := range md5Config.Title {
		// 比较标题，标题不脱敏
		if element == halfStr {
			cellValue = halfStr
			break
		} else {
			if md5Config.SaltModule.Activate == "on" {
				cellValue = str2md5Salt(halfStr, md5Config.SaltModule.SaltKey)
			} else if md5Config.SaltModule.Activate == "off" {
				cellValue = str2md5Normal(halfStr)
			}
		}
	}

	xlsx.SetCellValue(md5Config.OutSheetName, axis, cellValue)
	return 0
}

func main() {
	// excel路径
	excelPath := "./股东测试数据明文.xlsx"

	// config配置文件路径
	configPath := "./md5Tools.json"

	// 获取待解析的excel路径，需要执行时传入参
	if len(os.Args) > 1 {
		excelPath = os.Args[1]
		fmt.Println("需要读取的手机号文件路径是:", excelPath)
	} else {
		fmt.Println("需要读取的excel文件路径和文件名称")
	}

	fmt.Println("---------------开始脱敏---------------")

	md5Config := config.LoadConfig(configPath)

	// 调测时查看下配置参数
	//json, _ := json.Marshal(md5Config)
	//fmt.Println(string(json))
	fmt.Println("加载配置参数如下：（可以手动在配置文件修改）")
	fmt.Println("是否开启脱敏加盐模式：", md5Config.SaltModule.Activate)
	fmt.Println("脱敏的盐值：", md5Config.SaltModule.SaltKey)
	fmt.Println("匹配的标题：", md5Config.Title)
	fmt.Println("存放脱敏数据的sheet页：", md5Config.OutSheetName)

	// 解析excel并生成数组
	result := readExcel(excelPath, md5Config)
	fmt.Println("解析excel完毕...")

	// 转成 excel 坐标写入新的 sheet 页中
	writeExcelResult2Sheet(excelPath, result, md5Config)
	fmt.Println("数据脱敏完毕，请打开原excel的", md5Config.OutSheetName, ",注意标题复制过去了")
}

// 带固定盐的 md5，非标准 md5值
func str2md5Salt (plantStr string, salt string) string {
	saltData := []byte(salt)
	strData := []byte(plantStr)

	has := md5.New()
	has.Write(strData)
	has.Write(saltData)

	value := has.Sum(nil)
	return hex.EncodeToString(value)
}

// 不带盐，标准的 md5 值
func str2md5Normal (plantStr string) string {
	data := []byte(plantStr)
	has := md5.Sum(data)
	// 将[]byte转成16进制
	value := fmt.Sprintf("%x", has)
	return value
}

// 全角字符串转换半角字符串
// 参数：str：字符串
// 返回：转换半角后的字符串
//注意：并不是所有的全角字符都能被转换为半角字符，例如汉字是全角字符，占2个字符的位置，但它无法被转换；只有英文字母、数字键、符号键等才能可以做全角和半角之间的转换
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