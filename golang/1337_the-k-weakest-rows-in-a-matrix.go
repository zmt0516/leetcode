import "sort"

func sum(vec []int) int{
    res := -0
    for _,v := range vec{
        res+=v
    }
    return res
}

type astruct [][2]int

func (p astruct) Len() int {
	return len(p)
}

func (p astruct) Less(i, j int) bool {
    if p[i][0]==p[j][0]{
        return p[i][1]<p[j][1]
    } 
    return p[i][0]<p[j][0]
}

func (p astruct) Swap(i,j int){
    p[i],p[j] = p[j],p[i]
}

func kWeakestRows(mat [][]int, k int) []int {
    
    var list astruct
    
    for i,v := range mat{
        list = append(list,[2]int{sum(v),i})
    }
    sort.Sort(list)
    
    var r []int
    
    for i:=0;i<k;i++{
        r = append(r,list[i][1])
    }
    
    
    
    return r    
}