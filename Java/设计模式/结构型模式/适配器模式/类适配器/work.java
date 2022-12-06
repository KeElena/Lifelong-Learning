package adapter_mode;

public class work {
	public static void main(String args[]) {
		NetAdapter na=new NetAdapter();
		Computer pc =new Computer();
		pc.connect(na);
	}
}
