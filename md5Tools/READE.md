## 配置文件说明 

```json
{
  "saltModule": {
    "activate": "on",
    "saltKey": "saltKey"
  },
  "title": [
    "证券账户名称",
    "证券账户号码"
  ],
  "inSheetName": "Sheet1",
  "outSheetName": "Sheet1"
}
```


key|value|说明
-|-|-
saltModule|-|加盐模式对象，下面2个参数  
-|activate|on: 开启加盐模式<br>off：关闭加盐模式
-|saltKey|盐值参数的值，可以自己随意填充，默认值：saltKey
title|证券账户名称，证券账户号码|字符串数组结构，待脱敏的excel标题
inSheetName|Sheet1|待脱敏的excel的sheet页签名称
outSheetName|Sheet1|脱敏后的数据存放的sheet页签名称<br>1、outSheetName跟inSheetName一样的话会自动创建一个新的excel存放脱敏数据，新excel文件名格式是"原文件名_result.xlsx<br>2、outSheetName跟inSheetName不一样的话，会在原文件增加一个sheet页签保存脱敏数据

## 执行方式
mac:
```bash
chmod 700 md5Tools
./md5Tools ***.xlsx
```
win:
```bat
cmd
cd md5Tools工具所在的目录
md5Tools ***.xlsx
```

输出日志：
```
---------------开始脱敏---------------
加载配置参数如下：（可以手动在配置文件修改）
是否开启脱敏加盐模式： on
脱敏的盐值： saltKey
匹配的标题： [证券账户名称 证券账户号码]
存放脱敏数据的sheet页： Sheet1
解析excel完毕...
数据脱敏完毕，请打开原excel的 Sheet1 ,注意标题复制过去了
```

## FAQ
1、待脱敏的excel是空的，请检查  
>A: 待脱敏的目标文件的目标sheet页签数据是空的

2、没有匹配到需要脱敏的列标题，请检查excel  
>A：待脱敏的目标文件的目标 sheet 中没有匹配到需要脱敏的列标题，请检查配置文件的title和目标excel的待脱敏sheet页签的数据标题

3、其他异常信息  
>A: 请联系开发者

4、出现0开头的数字或者字符串
>A: 因为没有约束excel的模板格式，需要使用者保障这种0开头的数据不要用"数字"或者"自定义"格式，建议调整成文本格式，这样脱敏后数据才完整，否着excel自动会去掉前面的0
<br>比如：003A自动成了3A
