//go always for one more try
#include <bits/stdc++.h>
#define FOR(i, a, b) for (int i = (a); i < b; i++)
#define mp std::make_pair

using namespace std;

int main() {

    int t;
    cin >> t;

    while (t--) {
        string s;
        cin >> s;
        vector<int> a(0);
        int n = s.length();

        FOR(i, 0, n) {
            a.push_back(s[i] - 'a');
        }
        bool no_answer = true;
        for (int i = n - 1; i >= 1; i--) {
            if (a[i] > a[i - 1]) {
                no_answer = false;
                int posToSwap = i - 1;
                int posToSwapWith = -1;
                int diff = INT_MAX;
                for (int j = posToSwap + 1; j < n; j++) {
                    if (a[j] - a[posToSwap] > 0 && a[j] - a[posToSwap] < diff) {
                        posToSwapWith = j;
                        diff = a[j] - a[posToSwap];
                    }
                }
                int temp = a[posToSwap];
                a[posToSwap] = a[posToSwapWith];
                a[posToSwapWith] = temp;
                std::sort(a.begin() + posToSwap + 1, a.end());
                for (int i = 0; i < n; i++) {
                    char c = a[i] + 'a';
                    cout << c;
                }
                break;
            }
        }

        if (no_answer == true) {
            cout << "no answer";
        }
        cout << endl;
    }
    return 0;
}
