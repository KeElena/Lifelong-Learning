package adapter_mode;
//需要连接两端
public class NetAdapter extends NetWork implements Adapter {
	@Override
	public void handleFunc() {
		super.working();
	}
}
