#include<stdio.h>

void sort(int arr[],int left,int right,int len){
	
	if (len<2){
		return;
	}
	
	int temp;
	int x=left;
	int y=right;
	left++;
	
	while(left<right){
		while(arr[left]<arr[x]){left++;}
		while(arr[right]>arr[x]){right--;}
		if (left<right){
			temp=arr[left];
			arr[left]=arr[right];
			arr[right]=temp;
		}
	}
	temp=arr[x];
	arr[x]=arr[left-1];
	arr[left-1]=temp;
	
	sort(arr,x,left-1,left-x);
	sort(arr,right+1,y,y-right);
	
	return;
}

int main(){
	int arr[10]={4,7,2,8,9,5,7,10,-1,3};
	int len=sizeof(arr)/sizeof(arr[0]);
	
	sort(arr,0,len-1,len);
	
	for(int i=0;i<len;i++){
		printf("%d ",arr[i]);
	}
}
