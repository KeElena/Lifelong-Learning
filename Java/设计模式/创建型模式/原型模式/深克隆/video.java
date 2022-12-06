package clone_mode;

import java.util.Date;

public class video implements Cloneable{
	
	private String name;
	private Date time;
	
	public video(String name,Date time){
		this.setName(name);
		this.setDate(time);
	}
	@Override
	protected Object clone() throws CloneNotSupportedException{
		Object obj=super.clone();
		video v=(video)obj;
		v.time=(Date)this.time.clone();
		return obj;
	}
	
	public void setName(String name) {
		this.name=name;
	}
	public void setDate(Date time) {
		this.time=time;
	}
	public String toString() {
		return "name:"+name+" date:"+time.toString();
	}
}
