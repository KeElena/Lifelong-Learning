package abstract_factory;

public class chip_Impl implements chip {
	@Override
	public void Produce() {
		System.out.println("获取麒麟芯片");
	}
}
