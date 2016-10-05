//
//  numIslands.cpp
//  leetcode_algorithm
//
//  Created by Chriszou on 10/5/16.
//  Copyright © 2016 Chriszou. All rights reserved.
//

#include <iostream>
#include <vector>
#include <queue>

using namespace std;

class Index{
public:
    Index(int _i, int _j){
        i = _i;
        j = _j;
    }
    int i;
    int j;
};

bool isIndexValid(int m, int n, int i, int j){
    if(i < m && j < n && i >= 0 && j >= 0)
        return true;
    return false;
}

void bfs(const vector<vector<char>>& grid, vector<vector<bool>>& mark, int m, int n, int i, int j){
    queue<Index> qu;
    Index tmpIndex(i,j);
    qu.push(tmpIndex);
    mark[i][j] = true;
    
    while(!qu.empty()){
        Index tmp = qu.front();
        qu.pop();
        //mark[tmp.i][tmp.j] = true;
        if(isIndexValid(m, n, tmp.i+1, tmp.j) && mark[tmp.i+1][tmp.j] == false && grid[tmp.i+1][tmp.j] == '1'){
            qu.push(Index(tmp.i+1,tmp.j));
            mark[tmp.i+1][tmp.j] = true;
        }
        if(isIndexValid(m, n, tmp.i-1, tmp.j) && mark[tmp.i-1][tmp.j] == false && grid[tmp.i-1][tmp.j] == '1'){
            qu.push(Index(tmp.i-1,tmp.j));
            mark[tmp.i-1][tmp.j] = true;
        }
        if(isIndexValid(m, n, tmp.i, tmp.j+1) && mark[tmp.i][tmp.j+1] == false && grid[tmp.i][tmp.j+1] == '1'){
            qu.push(Index(tmp.i,tmp.j+1));
            mark[tmp.i][tmp.j+1] = true;
        }
        if(isIndexValid(m, n, tmp.i, tmp.j-1) && mark[tmp.i][tmp.j-1] == false && grid[tmp.i][tmp.j-1] == '1'){
            qu.push(Index(tmp.i,tmp.j-1));
            mark[tmp.i][tmp.j-1] = true;
        }
    }
}

int numIslands(vector<vector<char>>& grid) {
    int ret = 0;
    int m = (int)grid.size();
    if (m == 0)
        return 0;
    int n = (int)grid[0].size();
    
    vector<vector<bool>> mark(m, vector<bool>(n, false));
    for(int i = 0; i < m; i ++){
        for(int j = 0; j < n; j++){
            if(grid[i][j] == '0' || mark[i][j])
                continue;
            bfs(grid, mark, m, n, i, j);
            ret++;
        }
    }
    
    return ret;
}

int main(int argc, char *argv[]){
    vector<vector<char>> grid(4, vector<char>(5));
    grid[0] = {'1','1', '1', '1', '0'};
    grid[1] = {'1','1', '0', '1', '0'};
    grid[2] = {'1','1', '0', '0', '0'};
    grid[3] = {'0','0', '0', '0', '0'};
    
    
    vector<vector<char>> grid2(4, vector<char>(5));
    grid2[0] = {'1','1', '0', '0', '0'};
    grid2[1] = {'1','1', '0', '0', '0'};
    grid2[2] = {'0','0', '1', '0', '0'};
    grid2[3] = {'0','0', '0', '1', '1'};
    
    int ret = numIslands(grid);
    std::cout << ret << std::endl;
    
    ret = numIslands(grid2);
    std::cout << ret << std::endl;
}

