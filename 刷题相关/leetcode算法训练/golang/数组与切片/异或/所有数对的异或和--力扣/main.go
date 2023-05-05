func xorAllNums(nums1 []int, nums2 []int) int {
    res:=0
    count:=make(map[int]int)
    for _,val:=range nums1{
        count[val]=count[val]+len(nums2)%2
    }
    for _,val:=range nums2{
        count[val]=count[val]+len(nums1)%2
    }
    for key,ok:=range count{
        if ok==1{
            if res==0{
                res=key
            }else{
                res=res^key
            }
        }
    }
    return res
}