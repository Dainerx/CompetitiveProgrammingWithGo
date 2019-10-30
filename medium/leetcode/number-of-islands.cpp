class Solution {
public:
vector<vector<bool>> visited;

bool inBoundaries(pair<int, int> u, int r, int c)
{
    if ((u.first>=0 && u.first < r) && ( u.second >=0 && u.second < c))
        return true;
    return false;
}
vector<pair<int, int>> possibleMoves(pair<int, int> u)
{
    vector<pair<int, int>> moves(4);
    moves[0] = make_pair(u.first - 1, u.second);
    moves[1] = make_pair(u.first + 1, u.second);
    moves[2] = make_pair(u.first, u.second - 1);
    moves[3] = make_pair(u.first, u.second + 1);
    return moves;
}

void dfs(vector<vector<char>> &grid, pair<int, int> u, int r, int c)
{
    visited[u.first][u.second] = true;
    vector<pair<int, int>> moves = possibleMoves(u);
    for (int i = 0; i < moves.size(); i++)
    {
        pair<int, int> v = moves[i];
        if (inBoundaries(v, r, c) == true)
        {
            if (grid[v.first][v.second] == '1' && visited[v.first][v.second] == false )
                dfs(grid, v, r, c);
        }
    }
    return;
}

int numIslands(vector<vector<char>> &grid)
{
    if (grid.size() == 0)
        return 0;
    int islands = 0;
    int r = grid.size();
    int c = grid[0].size();
    cout << r << " " << c << "\n";
    visited.resize(r);
    for (int i = (0); i < r; i++)
    {
        for(int j=0; j<c; j++)
            visited[i].push_back(false);
    }

    for (int i = (0); i < r; i++)
    {
        for (int j = 0; j < c; j++)
        {
            if (!visited[i][j] && grid[i][j] == '1')
            {
                auto u = make_pair(i, j);
                dfs(grid, u, r, c);
                islands++;
            }
        }
    }

    return islands;
}
};
