func reformatNumber(number string) string {
	var res string
	number = strings.Replace(number, "-", "", -1)		//使用strings.Replcae()去除字符
	number = strings.Replace(number, " ", "", -1)
	for len(number) > 0 {								//使用for循环重新格式化
		if len(number) == 4 {							//长度剩余4时的格式化
			res+=number[0:2]+"-"+number[2:]
			number=""									//赋予空串
		} else if len(number)<4 {						//长度小于4时的格式化
			res+=number	
			number=""									//赋予空串
		}else{
			res+=number[0:3]+"-"						//其他情况以3个为单位进行格式化
			number=number[3:]							//裁剪前面3个字符
		}
	}
	return res
}