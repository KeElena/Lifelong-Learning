package abstract_factory;

public class consumer {
	public static void main(String args[]) {
		company huawei=new Huawei();
		huawei.getChip().Produce();
		huawei.getMobilePhone().Produce();
		huawei.getVM().IaaS();
	}
}
