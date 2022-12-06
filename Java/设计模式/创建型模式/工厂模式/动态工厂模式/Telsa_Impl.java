package dynamic_factory;

public class Telsa_Impl implements Car {
	@Override
	public void getName() {
		System.out.println("Tesla");
	}
}