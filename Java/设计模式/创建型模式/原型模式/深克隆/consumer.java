package clone_mode;

import java.util.Date;

public class consumer {
	public static void main(String args[]) throws CloneNotSupportedException {
		Date time=new Date();
		video v1=new video("myVideo",time);
		System.out.println(v1.toString());
		
		video v2=(video)v1.clone();
		time.setTime(1111111111);
		System.out.println(v2.toString());
		System.out.println(v1.toString());
	}
}
