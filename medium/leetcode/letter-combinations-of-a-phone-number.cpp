class Solution {
public:
unordered_map<long long, vector<string>> ht;

vector<vector<string>> groupAnagrams(vector<string> &strs)
{
    for (int i = 0; i < strs.size(); i++)
    {
        string s = strs[i];
        long long hv = 0;
        for (int j = 0; j < s.length(); j++)
        {
            string ss = "";
            ss+= s[j];
            hash<string> mystdhash; 
            hv += mystdhash(ss);
        }
        ht[hv].push_back(s);
    }

    vector<vector<string>> r(0);
    for(auto it = ht.begin(); it!=ht.end(); it++)
    {
        r.push_back(it->second);
    }
    return r;
}


};
