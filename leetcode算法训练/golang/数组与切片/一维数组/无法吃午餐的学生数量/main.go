func countStudents(students []int, sandwiches []int) int {
    squ:=0
    for _,val:=range students{
        squ+=val
    }
    cir:=len(students)-squ
    for _,val:=range sandwiches{
        if val==0 && cir>0{
            cir-=1
        }else if val==1 && squ>0{
            squ-=1
        }else{
            break
        }
    }
    return squ+cir
}