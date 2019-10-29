class Solution {
public:

unordered_map<string, bool> dp;
unordered_set<string> ht;

bool found(string s)
{
    auto it = ht.find(s);
    if (it != ht.end())
        return true;
    else
        return false;
}
bool recSearchWord(string s)
{
    if (s == "")
        return true;
    if (dp.find(s) != dp.end())
        return dp[s];
    for (int i = 1; i <= s.length(); i++)
    {
        string left = s.substr(0, i);
        string right = s.substr(i, s.length() - i);
        if (found(left) && recSearchWord(right))
            return dp[right] = true;
    }
    return dp[s] = false;
}

bool wordBreak(string s, vector<string> &wordDict)
{
    copy(wordDict.begin(), wordDict.end(), inserter(ht, ht.begin()));
    return recSearchWord(s);
}
};
