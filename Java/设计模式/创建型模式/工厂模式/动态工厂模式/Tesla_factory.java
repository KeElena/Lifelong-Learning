package dynamic_factory;

public class Tesla_factory implements factory {
	@Override
	public Car getInstance() {
		return new Telsa_Impl();
	}
}
