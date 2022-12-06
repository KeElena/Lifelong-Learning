package builder_mode;

public class worker extends Builder {
	public Product product=new Product();
	@Override
	public void setOne(String One) {
		product.setOne(One);
	}
	@Override
	public void setTwo(String Two) {
		product.setTwo(Two);
	}
	@Override
	public void setThree(String Three) {
		product.setThree(Three);
	}
	@Override
	public Product getProduct() {
		return product;
	}
}
