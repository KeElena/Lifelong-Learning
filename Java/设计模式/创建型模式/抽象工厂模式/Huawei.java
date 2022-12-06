package abstract_factory;

public class Huawei implements company{
	@Override
	public chip getChip() {
		return new chip_Impl();
	}
	@Override
	public mobilePhone getMobilePhone() {
		return new mobilePhone_Impl();
	}
	@Override
	public cloudComputing getVM() {
		return new cloudComputing_Impl();
	}
}