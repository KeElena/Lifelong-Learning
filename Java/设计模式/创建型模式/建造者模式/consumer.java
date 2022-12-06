package builder_mode;

public class consumer {
	public static void main(String args[]) {
		Product product=director.getProduct(new worker());
		System.out.println(product.getContent());
	}
}
