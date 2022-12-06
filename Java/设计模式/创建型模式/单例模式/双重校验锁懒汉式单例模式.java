package Simgle_Mode_Code;

class mySimgle{
	//使用volatile保证对象实例化的原子性
	private volatile static mySimgle myObj;
	//构造方法私有化，外部不能访问
	private mySimgle() {}
	//获取唯一对象
	public static mySimgle getInstence() {
        //判断对象是否为null
		if(myObj==null){
            //是null则对对象加锁
			synchronized(mySimgle.class) {
                //加锁后再检测对象是否为空（并发情况下是必要的）
				if(myObj==null){
					myObj=new mySimgle();
				}
			}
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