package adapter_mode;

public class work {
	public static void main(String args[]) {
		//获取网线
		NetWork net=new NetWork();
		//网线接入适配器
		NetAdapter na=new NetAdapter(net);
		//获取电脑
		Computer pc=new Computer();
		//电脑调用连接方法并插入适配器
		pc.connect(na);
	}
}
