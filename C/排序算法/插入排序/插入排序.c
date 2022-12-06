#include<stdio.h>

void sort(int arr[],int len){
	int j,temp;
	for(int i=1;i<len;i++){
		j=i;
		while(arr[j]<arr[j-1]){
			temp=arr[j];
			arr[j]=arr[j-1];
			arr[j-1]=temp;
			j--;
			if(j==0){
				break;
			}
		}
	}
	
}

int main(){
	int arr[10]={4,7,2,8,9,5,7,10,-1,3};
	int len=sizeof(arr)/sizeof(arr[0]);
	sort(arr,len);
	
	for(int i=0;i<len;i++){
		printf("%d ",arr[i]);
	}
}
