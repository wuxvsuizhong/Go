package main

import (
	. "mypro/requestPocess"
)

func main() {

	//获取表单数据
	// GetFormDataByPOST()
	//上传文件
	// UploadFile()
	//使用路由组
	// UseRouteGroup()

	//按照json格式解析数据
	// ParseAsJSON()

	//按照json格式解析变单数据
	// ParseFormDataAsJson()

	//按照json格式解析Uri中的字段参数
	ParseUriParamsAsJson()
}
