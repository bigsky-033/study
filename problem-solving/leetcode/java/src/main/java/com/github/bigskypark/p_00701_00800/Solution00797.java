package com.github.bigskypark.p_00701_00800;

// 797. All Paths From Source to Target,
// https://leetcode.com/problems/all-paths-from-source-to-target/

import java.util.ArrayList;
import java.util.List;

public class Solution00797 {
  public List<List<Integer>> allPathsSourceTarget(int[][] adjMatrix) {
    List<List<Integer>> allPaths = new ArrayList<>();
    dfs(adjMatrix, 0, allPaths, new ArrayList<>());
    return allPaths;
  }

  private void dfs(
      final int[][] adjMatrix,
      final int currentNode,
      final List<List<Integer>> allPaths,
      final List<Integer> path) {
    path.add(currentNode);

    if (currentNode == adjMatrix.length - 1) {
      allPaths.add(new ArrayList<>(path));
    } else {
      for (int node : adjMatrix[currentNode]) {
        dfs(adjMatrix, node, allPaths, path);
      }
    }

    path.remove(path.size() - 1);
  }
}
