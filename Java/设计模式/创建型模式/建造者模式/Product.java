package builder_mode;

public class Product {
	private String One;
	private String Two;
	private String Three;
	
	public String getOne() {
		return One;
	}
	public String getTwo() {
		return Two;
	}
	public String getThree() {
		return Three;
	}
	
	public void setOne(String One) {
		this.One=One;
	}
	public void setTwo(String Two) {
		this.Two=Two;
	}
	public void setThree(String Three) {
		this.Three=Three;
	}
	
	public String getContent() {
		return "One:"+One+" Two:"+Two+" Three:"+Three;
	}
}
