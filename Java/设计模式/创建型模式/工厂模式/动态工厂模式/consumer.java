package dynamic_factory;

public class consumer {
	public static void main(String args[]) {
		Car myBYD=new BYD_factory().getInstance();
		Car myTesla=new Tesla_factory().getInstance();
		
		myBYD.getName();
		myTesla.getName();
	}
}
