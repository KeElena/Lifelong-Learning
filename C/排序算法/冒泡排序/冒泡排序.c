#include<stdio.h>

//冒泡排序每遍历一次会将最大值移动到最后一个位置
//冒泡排序每遍历一次，遍历的元素减少1个

void sort(int arr[],int len){		//输入数组和数组长度
	int temp=0;						//定义中间值
	while(len!=0){					//长度不等于0时循环
		for(int i=1;i<len;i++){		//根据len进行遍历
			if(arr[i-1]>arr[i]){	//将最大值移到arr[len-1]的位置
				temp=arr[i];
				arr[i]=arr[i-1];
				arr[i-1]=temp;
			}
		}
		len--;						//遍历一次后len-1，最后减到0完成排序
	}
}

int main(){
	int arr[10]={4,7,2,8,9,4,7,10,-1,3};
	int len=sizeof(arr)/sizeof(arr[0]);
	sort(arr,len);
	for(int i=0;i<len;i++){
		printf("%d ",arr[i]);
	}
}
