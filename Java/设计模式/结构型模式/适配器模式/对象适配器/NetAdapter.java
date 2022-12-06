package adapter_mode;
//构建适配器，使用构造方法获取网线对象
public class NetAdapter implements Adapter {
	private NetWork net;
	public NetAdapter(NetWork net) {
		this.net=net;
	}
	@Override
	public void handleFunc() {
		net.working();
	}
}