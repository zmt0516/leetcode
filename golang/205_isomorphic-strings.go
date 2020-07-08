func isIsomorphic(s string, t string) bool {
        

    if len(s)!=len(t){
        return false
    }
    emap:=make(map[byte]byte)
    emap2:=make(map[byte]bool)
    for i:=range s{
        v,ok := emap[s[i]]
        if !ok {
            if emap2[t[i]] {
                return false
            }
            emap[s[i]]=t[i]
        } else{
            if v!=t[i]{
                return false
            }
            
        }
        emap2[t[i]]=true
        
        
    }
    return true
    
}
