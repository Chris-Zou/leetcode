# Write your MySQL query statement below
SELECT DISTINCT Num as ConsecutiveNums FROM (
  SELECT Num, COUNT(Rank) AS Cnt FROM (
    SELECT    Num,
              @curRank := @curRank + IF(@prevNum = Num, 0, 1) AS rank, @prevNum := Num
    FROM      Logs s, (SELECT @curRank := 0) r, (SELECT @prevNum := NULL) p
    ORDER BY  ID ASC
) t GROUP BY Rank HAVING Cnt >= 3 
) n;
