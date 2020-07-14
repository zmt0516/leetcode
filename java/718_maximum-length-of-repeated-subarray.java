class Solution {
    public int findLength(int[] A, int[] B) {
        int lenB=B.length;
        int r=0;
        int[] emap=new int[lenB+1];
        


                
        for (int v:A){
            for (int i=lenB;i>0;i--){
                if (v==B[i-1]){
                    int lene=emap[i-1]+1;
                    emap[i]=lene;
                    if (lene>r){
                        r=lene;
                    }
                }else{
                    emap[i]=0;
                }
            }
        }
        
        return r;
        
        
    }
}