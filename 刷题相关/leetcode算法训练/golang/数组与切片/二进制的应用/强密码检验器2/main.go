func strongPasswordCheckerII(password string) bool {
    //长度小于8则返回
    if len(password)<8{
        return false
    }
    //定义掩码，在二进制下，使用或运算记录1，并防止1变为0
    mask:=0
    for i,char:=range password{
        //出现连续子串时返回false
        if i>0 && password[i]==password[i-1]{
            return false
        }
        //处理小写
        if unicode.IsLower(char){
            mask|=1
        //处理大写
        }else if unicode.IsUpper(char){
            mask|=2
        //处理数字
        }else if unicode.IsDigit(char){
            mask|=4
        //处理是否包含特殊字符
        }else if strings.Contains("!@#$%^&*()-+",password[i:i+1]){
            mask|=8
        }
    }
    //掩码为15时返回true（二进制下为1111）
    return mask==15
}