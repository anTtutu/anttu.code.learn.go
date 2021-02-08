package tool

import "testing"

/**
 * 测试手机号
 */
func TestValidateMobile(t *testing.T) {
	mobileDemo1 := "13512345678"
	mobileDemo2 := "14712345678"
	mobileDemo3 := "15712345678"
	mobileDemo4 := "16612345678"
	mobileDemo5 := "17712345678"
	mobileDemo6 := "18912345678"
	mobileDemo7 := "19912345678"
	mobileDemo8 := "12345678901"

	boolDemo1 := validateMobile(mobileDemo1)
	if !boolDemo1 {
		t.Errorf("1- validate:%t mobile:%s", boolDemo1, mobileDemo1)
	}
	boolDemo2 := validateMobile(mobileDemo2)
	if !boolDemo2 {
		t.Errorf("2- validate:%t mobile:%s", boolDemo2, mobileDemo2)
	}

	boolDemo3 := validateMobile(mobileDemo3)
	if !boolDemo3 {
		t.Errorf("3- validate:%t mobile:%s", boolDemo3, mobileDemo3)
	}

	boolDemo4 := validateMobile(mobileDemo4)
	if !boolDemo4 {
		t.Errorf("4- validate:%t mobile:%s", boolDemo4, mobileDemo4)
	}

	boolDemo5 := validateMobile(mobileDemo5)
	if !boolDemo5 {
		t.Errorf("5- validate:%t mobile:%s", boolDemo5, mobileDemo5)
	}

	boolDemo6 := validateMobile(mobileDemo6)
	if !boolDemo6 {
		t.Errorf("6- validate:%t mobile:%s", boolDemo6, mobileDemo6)
	}

	boolDemo7 := validateMobile(mobileDemo7)
	if !boolDemo7 {
		t.Errorf("7- validate:%t mobile:%s", boolDemo7, mobileDemo7)
	}

	boolDemo8 := validateMobile(mobileDemo8)
	if !boolDemo8 {
		t.Errorf("8- validate:%t mobile:%s", boolDemo8, mobileDemo8)
	}
}

/**
 * 测试中国身份证号
 */
func TestValidateIdCard(t *testing.T) {

	/**
	 * 网上找了一批通过工具生成的测试用的公开身份证号，不涉及泄露隐私 https://www.cnblogs.com/linus-tan/p/7111797.html
	 */
	/**
	姓名和身份证号码	性别	年龄	所属地区
	刁以松 230227198302151067	女	34	黑龙江省 齐齐哈尔市 富裕县
	舒佳乐 371522199402189127	女	23	山东省 聊城市 莘县
	史梦玉 120111199207178301	女	25	天津市 市辖区 西青区
	蓝秋柳 230713198012022856	男	37	黑龙江省 伊春市 带岭区
	柯惜寒 45128119860426100X	女	31	广西壮族自治区
	华蔚蓝 141081198101053723	女	36	山西省
	崔书竹 652701198205147107	女	35	新疆维吾尔自治区 博尔塔拉蒙古自治州 博乐市
	程嘉昌 410922198502159418	男	32	河南省 濮阳市 清丰县
	厍问风 500111199412160849	女	23	重庆市 市辖区 双桥区
	蓝淼   33010619930228635X	男	24	浙江省 杭州市 西湖区
	戈雯艳 220102198405100831	男	33	吉林省 长春市 南关区
	柳晨程 370724198508236874	男	32	山东省 潍坊市 临朐县
	*/

	// 18位
	idCardNoDemo1 := "123456789012345678"
	idCardNoDemo2 := "230227198302151067"
	idCardNoDemo3 := "371522199402189127"
	idCardNoDemo4 := "120111199207178301"
	idCardNoDemo5 := "230713198012022856"
	idCardNoDemo6 := "45128119860426100X"
	idCardNoDemo7 := "141081198101053723"
	idCardNoDemo8 := "652701198205147107"
	idCardNoDemo9 := "410922198502159418"
	idCardNoDemo10 := "500111199412160849"
	idCardNoDemo11 := "33010619930228635X"
	idCardNoDemo12 := "220102198405100831"
	idCardNoDemo13 := "370724198508236874"

	boolDemo1 := validateIdCard(idCardNoDemo1)
	if !boolDemo1 {
		t.Errorf("1- validate:%t idCardNo:%s", boolDemo1, idCardNoDemo1)
	}

	boolDemo2 := validateIdCard(idCardNoDemo2)
	if !boolDemo2 {
		t.Errorf("2- validate:%t idCardNo:%s", boolDemo2, idCardNoDemo2)
	}

	boolDemo3 := validateIdCard(idCardNoDemo3)
	if !boolDemo3 {
		t.Errorf("3- validate:%t idCardNo:%s", boolDemo3, idCardNoDemo3)
	}

	boolDemo4 := validateIdCard(idCardNoDemo4)
	if !boolDemo4 {
		t.Errorf("4- validate:%t idCardNo:%s", boolDemo4, idCardNoDemo4)
	}

	boolDemo5 := validateIdCard(idCardNoDemo5)
	if !boolDemo5 {
		t.Errorf("5- validate:%t idCardNo:%s", boolDemo5, idCardNoDemo5)
	}

	boolDemo6 := validateIdCard(idCardNoDemo6)
	if !boolDemo6 {
		t.Errorf("6- validate:%t idCardNo:%s", boolDemo6, idCardNoDemo6)
	}

	boolDemo7 := validateIdCard(idCardNoDemo7)
	if !boolDemo7 {
		t.Errorf("7- validate:%t idCardNo:%s", boolDemo7, idCardNoDemo7)
	}

	boolDemo8 := validateIdCard(idCardNoDemo8)
	if !boolDemo8 {
		t.Errorf("8- validate:%t idCardNo:%s", boolDemo8, idCardNoDemo8)
	}

	boolDemo9 := validateIdCard(idCardNoDemo9)
	if !boolDemo9 {
		t.Errorf("9- validate:%t idCardNo:%s", boolDemo9, idCardNoDemo9)
	}

	boolDemo10 := validateIdCard(idCardNoDemo10)
	if !boolDemo10 {
		t.Errorf("10- validate:%t idCardNo:%s", boolDemo10, idCardNoDemo10)
	}

	boolDemo11 := validateIdCard(idCardNoDemo11)
	if !boolDemo11 {
		t.Errorf("11- validate:%t idCardNo:%s", boolDemo11, idCardNoDemo11)
	}

	boolDemo12 := validateIdCard(idCardNoDemo12)
	if !boolDemo12 {
		t.Errorf("12- validate:%t idCardNo:%s", boolDemo12, idCardNoDemo12)
	}

	boolDemo13 := validateIdCard(idCardNoDemo13)
	if !boolDemo13 {
		t.Errorf("13- validate:%t idCardNo:%s", boolDemo13, idCardNoDemo13)
	}
}

func TestValidateTaxNo(t *testing.T) {
	// 税号来源于腾讯、百度、阿里、大疆 公开信息用于测试
	taxNoDemo1 := "9144030071526726XG"
	taxNoDemo2 := "91330100799655058B"
	taxNoDemo3 := "914403007954257495"
	taxNoDemo4 := "911103026605015136"
	taxNoDemo5 := "123456789012345678"

	boolDemo1 := validateTaxNo(taxNoDemo1)
	if !boolDemo1 {
		t.Errorf("1- validate:%t taxNo:%s", boolDemo1, taxNoDemo1)
	}

	boolDemo2 := validateTaxNo(taxNoDemo2)
	if !boolDemo2 {
		t.Errorf("2- validate:%t taxNo:%s", boolDemo2, taxNoDemo2)
	}

	boolDemo3 := validateTaxNo(taxNoDemo3)
	if !boolDemo3 {
		t.Errorf("3- validate:%t taxNo:%s", boolDemo3, taxNoDemo3)
	}

	boolDemo4 := validateTaxNo(taxNoDemo4)
	if !boolDemo4 {
		t.Errorf("4- validate:%t taxNo:%s", boolDemo4, taxNoDemo4)
	}

	boolDemo5 := validateTaxNo(taxNoDemo5)
	if !boolDemo5 {
		t.Errorf("5- validate:%t taxNo:%s", boolDemo5, taxNoDemo5)
	}
}

func TestValidateDate(t *testing.T) {
	dateDemo1 := "2019-02-28"
	dateDemo2 := "2019-02-29"
	dateDemo3 := "2019-13-01"
	dateDemo4 := "2019-09-31"
	dateDemo5 := "2000-02-29"
	dateDemo6 := "2000-02-28"
	dateDemo7 := "2020-01-33"

	boolDemo1 := validateDate(dateDemo1)
	if !boolDemo1 {
		t.Errorf("1- validate:%t dateStr:%s", boolDemo1, dateDemo1)
	}

	boolDemo2 := validateDate(dateDemo2)
	if !boolDemo2 {
		t.Errorf("2- validate:%t dateStr:%s", boolDemo2, dateDemo2)
	}

	boolDemo3 := validateDate(dateDemo3)
	if !boolDemo3 {
		t.Errorf("3- validate:%t dateStr:%s", boolDemo3, dateDemo3)
	}

	boolDemo4 := validateDate(dateDemo4)
	if !boolDemo4 {
		t.Errorf("4- validate:%t dateStr:%s", boolDemo4, dateDemo4)
	}

	boolDemo5 := validateDate(dateDemo5)
	if !boolDemo1 {
		t.Errorf("5- validate:%t dateStr:%s", boolDemo5, dateDemo5)
	}

	boolDemo6 := validateDate(dateDemo6)
	if !boolDemo6 {
		t.Errorf("6- validate:%t dateStr:%s", boolDemo6, dateDemo6)
	}

	boolDemo7 := validateDate(dateDemo7)
	if !boolDemo7 {
		t.Errorf("7- validate:%t dateStr:%s", boolDemo7, dateDemo7)
	}
}
