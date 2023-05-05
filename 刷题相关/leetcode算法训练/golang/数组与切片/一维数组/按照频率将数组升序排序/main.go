func frequencySort(nums []int) []int {
    store:=make(map[int]int,len(nums)/2)        //构造散列表
    
    for _,val:=range nums{                      //遍历nums，记录元素出现频率
        store[val]++
    }

    var a,b int
    //i和j表示的是索引，不是值
    sort.Slice(nums,func(i,j int)bool{          //排序

        a=nums[i]                               //取得下一个索引对应的元素
        b=nums[j]                               //取得当前索引对应的元素
        if store[a]<store[b]{                   //下一个元素出现频率小于当前元素的出现频率
            return true                         //交换
        }
        if store[a]==store[b] && a>b{           //出现频率相同，但是下一个元素的值大于当前元素的值
            return true                         //交换
        }
        return false
    })

    return nums
}