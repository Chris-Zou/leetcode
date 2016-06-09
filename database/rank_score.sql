T Scores.Score, COUNT(Ranking.Score) AS Rank
  FROM Scores
       , (
        SELECT DISTINCT Score
        FROM Scores
        ) Ranking
WHERE Scores.Score <= Ranking.Score
GROUP BY Scores.Id, Scores.Score
ORDER BY Scores.Score DESC;
