package simple_factory;

public class consumer {
	public static void main(String args[]) {
		Car myBYD=factory.getBYD();
		Car myTesla=factory.getTesla();
		
		myBYD.getName();
		myTesla.getName();
	}
}
