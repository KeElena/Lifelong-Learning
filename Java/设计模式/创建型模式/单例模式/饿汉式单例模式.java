package Simgle_Mode_Code;

class mySimgle{
	//使用static关键字在不依赖类实例化的情况下完成实例化唯一的对象
	private static mySimgle myObj =new mySimgle();
	//构造方法私有化，外部不能调用构造方法实例化类
	private mySimgle() {}
	//构建一个获取唯一对象的函数
	public static mySimgle getInstence() {
		return myObj;
	}
	//构建对象的相关方法
	public void Info() {
		System.out.println("OK!");
	}
}

public class Simgle_Mode_Code {
	public static void main(String args[]) {
		mySimgle myObj=mySimgle.getInstence();
		myObj.Info();
	}
}