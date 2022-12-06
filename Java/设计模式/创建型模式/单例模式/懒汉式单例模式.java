package Simgle_Mode_Code;

class mySimgle{
	//使用static在未实例化的情况下声明对象
	private static mySimgle myObj;
	//构造方法私有化，外部不能访问
	private mySimgle() {}
	//获取唯一对象
	public static mySimgle getInstence() {
		//判断对象是否为null
		if(myObj==null){
			//是则实例化对象
			myObj=new mySimgle();
		}
		return myObj;
	}
	//对象相关方法
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