package clone_mode;

import java.util.Date;

public class consumer {
	public static void main(String args[]) throws CloneNotSupportedException {
		video v1=new video("myVideo",new Date());
		System.out.println(v1.toString());
		
		video v2=(video)v1.clone();
		System.out.println(v2.toString());
	}
}
