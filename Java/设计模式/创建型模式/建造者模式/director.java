package builder_mode;

public class director {
	public static Product getProduct(Builder build) {
		build.setOne("第一步");
		build.setTwo("第二步");
		build.setThree("第三步");
		return build.getProduct();
	}
}
