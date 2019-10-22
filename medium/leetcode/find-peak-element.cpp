void recSearch(vector<int> &arr, int l, int r, int &result)
{
    if (r >= l)
    {
        int mid = l + (r - l) / 2;
        long l_neighbor = (mid-1<0) ? LONG_MIN : arr[mid-1];
        long r_neighbor = (mid+1>arr.size()-1) ? LONG_MIN : arr[mid+1];
        if (arr[mid] > l_neighbor && arr[mid] > r_neighbor) {
            result =  mid;
            return;
        }
        recSearch(arr, l, mid - 1,result);
        recSearch(arr, mid + 1, r,result);
    }
    else
        return;
}

int findPeakElement(vector<int> &nums)
{
    int result = -1;
    recSearch(nums,0,nums.size()-1,result);
    return result;
}
