package main

import "fmt"

//用于求数的逆元
//由于除法不能用取模运算，所以需要将式子转为乘法运算才能进行取模运算
//一个数的逆元可以将除运算转为乘法运算
//n/m==n*e（e为m的逆元）
//需要参数：i、mod
//i为需要求逆元的数
//mod为逆元的取模大小
func inv(i,mod int)int{
	if i==1{
        return 1
    }
	return (mod-mod/i)*inv(mod%i,mod)%mod
}

func main(){
	n:=inv(3,10)
	fmt.Println(n)
}