package util

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
)

//保存json字符串到文件
func SaveJsonStr2File(data string,filename string)  {
	err := ioutil.WriteFile(filename, []byte(data), 766)
	if err != nil {
		logs.Error("save ",filename, "to file error.",err.Error())

		panic(err.Error())
		return
	}
	logs.Debug("save ",filename," to file success.")
}

//保存成json文件
func Save2JsonFile(vData interface{}, filename string) {
	data, err := json.Marshal(vData)
	if err != nil {
		logs.Error("json error:", err.Error())
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		logs.Error("write file error:", err.Error())
		return
	}
	logs.Debug("save ",filename," success")
}
//保存成json文件
func Save2FormatJsonFile(vData interface{}, filename string,indent string) {
	data, err := json.MarshalIndent(vData," ",indent)
	if err != nil {
		logs.Error("json error:", err.Error())
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		logs.Error("write file error:", err.Error())
		return
	}
	logs.Debug("save json success")
}

//将对象转换成不带缩进格式的json字符串
func Obj2JsonStr(vData interface{}) string {
	if data, err := json.Marshal(vData); err != nil {
		logs.Error("object to json error:", err.Error())
		panic(err.Error())
	} else {
		logs.Debug("object to json str success")
		return string(data)
	}
}

//将对象转换成带缩进格式的json字符串
func Obj2JsonStrIndex(vData interface{}, prefix, indent string) string {
	data, err := json.MarshalIndent(vData, prefix, indent)
	if err != nil {
		logs.Error("object to json error:", err.Error())
		panic(err.Error())
	}
	logs.Debug("object to json str success")
	return string(data)
}

//保存成csv文件,需要考虑反射实现
func save2CsvFile(vData interface{}, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 766)
	if err != nil {
		logs.Error("open file error:", err.Error())
		panic(err.Error())
	}
	writer := csv.NewWriter(file)
	test := []string{"1231", "234324234"}
	err = writer.Write(test)
	if err != nil {
		logs.Error("write csv file error:", err.Error())
		return
	}
	writer.Flush()
}

//保存成xml文件
func Save2XmlFile1(vData interface{}, filename string) {
	if file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0766); err != nil {
		logs.Error("open file error:", err.Error())
		panic(err.Error())
	} else {
		encoder := xml.NewEncoder(file)
		encoder.Indent(" ", "\t")
		err = encoder.Encode(vData)
		if err != nil {
			logs.Error("write to xml file error:", err.Error())
			return
		}
		err = encoder.Flush()
		if err != nil {
			logs.Error("flush to xml file error:", err.Error())
			return
		}
		logs.Debug("save to xml file success ")
	}
}

//保存成xml文件
func Save2XmlFile2(vData interface{}, filename string) {
	data := Obj2XmlStrIndent(vData, " ", "\t")
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)

	if err != nil {
		logs.Error("open file error:", err.Error())
		panic(err.Error())
	}
	//也可以直接写bytes
	if _, err = file.WriteString(data); err != nil {
		logs.Error("write data to file error:", err.Error())
		return
	}
	if err = file.Close(); err != nil {
		logs.Error("close file error:", err.Error())
		return
	}
	logs.Debug("save to xml file success")
}

//将对象转换成不带缩进的xml字符串
func Obj2XmlStr(vData interface{}) string {
	data, err := xml.Marshal(vData)
	if err != nil {
		logs.Error("xml to byte error:", err.Error())
		panic(err.Error())
	}
	logs.Debug("object to xml str success")
	return string(data)
}

//将对象转换成带缩进的xml字符串
func Obj2XmlStrIndent(vData interface{}, prefix, indent string) string {
	data, err := xml.MarshalIndent(vData, prefix, indent)
	if err != nil {
		logs.Error("xml to byte error:", err.Error())
		panic(err.Error())
	}
	logs.Debug("object to xml indent str success")
	return string(data)
}
