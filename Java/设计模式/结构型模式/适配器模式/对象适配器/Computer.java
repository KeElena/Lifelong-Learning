package adapter_mode;

public class Computer {
   	 //调用适配器的方法
	public void connect(NetAdapter na) {
		na.handleFunc();
	}
}
