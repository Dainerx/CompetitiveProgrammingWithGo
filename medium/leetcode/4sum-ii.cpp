class Solution {
public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
        map<int,int> AB;
        for(int i = 0; i<A.size(); i++)
        {
            for (int j = 0; j<B.size(); j++)
            {
                int r = A[i] + B[j];
                if (AB.find(r) == AB.end())
                    AB.insert(make_pair(r,1));
                else
                    AB[r] = AB[r] + 1;
            }
        }
        
        map<int,int> CD;
        for(int i = 0; i<C.size(); i++)
        {
            for (int j = 0; j<D.size(); j++)
            {
                int r = C[i] + D[j];
                if (CD.find(r) == CD.end())
                    CD.insert(make_pair(r,1));
                else
                    CD[r] = CD[r] + 1;
            }
        }

        int count = 0;
        for(auto it=AB.begin(); it!=AB.end(); it++)
        {
            int r = it->first;
            auto tt = CD.find(-r);
            if (tt!=CD.end())
            {
                count+=it->second*tt->second; 
            }
        }
        return count;
    }
};
