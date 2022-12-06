package dynamic_factory;

public class BYD_factory implements factory {
	public Car getInstance() {
		return new BYD_Impl();
	}
}
