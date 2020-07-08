import "strings"


func wordPattern(pattern string, str string) bool {
    
    ss:=strings.Split(str," ")
    if len(ss)!=len(pattern){
        return false
    }
    emap:=make(map[string]byte)
    emap2:=make(map[byte]bool)
    for i:=range ss{
        v,ok := emap[ss[i]]
        if !ok {
            if emap2[pattern[i]] {
                return false
            }
            emap[ss[i]]=pattern[i]
        } else{
            if v!=pattern[i]{
                return false
            }
            
        }
        emap2[pattern[i]]=true
        
        
    }
    return true
    
    
}
