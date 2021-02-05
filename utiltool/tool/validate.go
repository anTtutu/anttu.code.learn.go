package tool

import (
	"fmt"
	"regexp"
	module "utiltool/module"
)

// 手机号正则   支持不带86、+86，也支持带，
const (
	mobileRegular = "^1[3456789]\\d{9}$"
)

/**
 * 正则匹配是否手机号
 */
func validateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(mobileRegular)
	return reg.MatchString(mobileNum)
}

/**
 * 校验中国大陆身份证号
 */
func validateIdCard(idCardNo string) bool {
	idCardMsgRelation, err := module.AnalysisIDCard(idCardNo)
	if err != nil {
		fmt.Println("check idCard error,", err.Error())
		return false
	} else {
		fmt.Printf("IdCard:%v\n", idCardMsgRelation)
		return true
	}
}

/**
 * 校验税号，统一信用代码
 */
func validateTaxNo(taxNo string) bool {
	module.New(taxNo)
	boolResult := module.CheckTaxNo(module.New(taxNo))
	if boolResult {
		return true
	} else {
		fmt.Printf("tax no :%s", taxNo)
		return false
	}
}
