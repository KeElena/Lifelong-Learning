package simple_factory;

public class factory {
	public static Car getBYD() {
		return new BYD_Impl();
	}
	public static Car getTesla() {
		return new Tesla_Impl();
	}
}
