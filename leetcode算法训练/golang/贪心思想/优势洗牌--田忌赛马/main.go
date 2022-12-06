func advantageCount(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    idx:=make([]int,len(nums2))
    res:=make([]int,len(nums1))
    for i:=range idx{
        idx[i]=i
    }

    sort.Slice(idx,func(i,j int)bool{
        return nums2[idx[i]]<nums2[idx[j]]
    })
    
    left:=0
    right:=len(nums1)-1
    for _,val:=range nums1{
        if val>nums2[idx[left]]{
            res[idx[left]]=val
            left++
        }else{
            res[idx[right]]=val
            right--
        }
    }
    return res
}