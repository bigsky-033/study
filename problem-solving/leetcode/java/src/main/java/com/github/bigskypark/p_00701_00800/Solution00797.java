package com.github.bigskypark.p_00701_00800;

// 797. All Paths From Source to Target,
// https://leetcode.com/problems/all-paths-from-source-to-target/

import java.util.ArrayList;
import java.util.List;

public class Solution00797 {
  public List<List<Integer>> allPathsSourceTarget(int[][] adjMatrix) {
    List<List<Integer>> allPaths = new ArrayList<>();
    List<Integer> path = new ArrayList<>();
    path.add(0); // starting point
    dfs(adjMatrix, 0, allPaths, path);
    return allPaths;
  }

  private void dfs(
      final int[][] adjMatrix,
      final int currentNode,
      final List<List<Integer>> allPaths,
      final List<Integer> path) {
    if (currentNode == adjMatrix.length - 1) {
      allPaths.add(new ArrayList<>(path));
      return;
    }

    for (int node : adjMatrix[currentNode]) {
      path.add(node);
      dfs(adjMatrix, node, allPaths, path);
      path.remove(path.size() - 1);
    }
  }
}
