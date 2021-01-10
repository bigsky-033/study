# 178. Rank Scores, https://leetcode.com/problems/rank-scores/

SELECT score,
       dense_rank() over (order by Score desc) AS `Rank`
FROM Scores