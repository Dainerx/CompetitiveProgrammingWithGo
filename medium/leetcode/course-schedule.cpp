class Solution {
public:
vector<bool> visited;
vector<int> toposort;
vector<vector<int>> graph;
void dfs(int curr)
{
    visited[curr] = true;
    for (int i = 0; i < graph[curr].size(); i++)
    {
        int v = graph[curr][i];
        if (!visited[v])
            dfs(v);
    }
    toposort.push_back(curr);
}

bool canFinish(int numCourses, vector<vector<int>> &prerequisites)
{
    graph.resize(numCourses);
    for (int i = 0; i < prerequisites.size(); i++)
    {
        auto vv = prerequisites[i];
        graph[vv[1]].push_back(vv[0]);
    }
    visited.resize(0);
    for (int i = 0; i < numCourses; i++)
    {
        visited.push_back(false);
    }
    for (int i = 0; i < numCourses; i++)
    {
        if (!visited[i])
            dfs(i);
    }

    for (int i = 0; i < toposort.size(); i++)
    {
        int course = toposort[i];
        auto courses_after = graph[course];
        for (int c = 0; c < courses_after.size(); c++)
        {
            auto cc = courses_after[c];
            auto it = find(toposort.begin() + i, toposort.end(), cc);
            if (it != toposort.end())
                return false;
        }
    }

    return true;
}
};
