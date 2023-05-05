func sortPeople(names []string, heights []int) []string {
    var res []string
    store:=make(map[int]string,len(names))  //构建身高到名字的字典
    for i:=0;i<len(names);i++{              //输入数据
        store[heights[i]]=names[i]
    }
    
    sort.Ints(heights)                      //对身高排序
    for i:=len(heights)-1;i>-1;i--{         //按大到小取出名字
        res=append(res,store[heights[i]])
    }
    return res
}