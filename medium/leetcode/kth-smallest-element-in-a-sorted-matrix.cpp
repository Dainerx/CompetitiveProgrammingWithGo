class Solution {
public:
int bs(vector<vector<int>> &matrix, int k)
{
    int n = matrix.size();
    int left = matrix[0][0], right = matrix[n - 1][n - 1];
    while (left < right)
    {
        int mid = (right + left) / 2;
        int r = countLessOrEqual(matrix, n, mid);
        if (r<k)
            left = mid + 1;
        else 
            right = mid;
    }
    return left;
}

int countLessOrEqual(vector<vector<int>> &matrix, int n, int target)
{
    int count = 0;
    int i, j;
    for (i = 0; i < n; i++)
    {
        j = 0;
        while (j < n && target >= matrix[i][j])
        {
            j++;
            count++;
        }
    }
    return count;
}

int kthSmallest(vector<vector<int>> &matrix, int k)
{
    int r = bs(matrix, k);
    return r;
}
    
};
