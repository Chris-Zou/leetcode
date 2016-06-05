# Write your MySQL query statement below
SELECT S1.Score, COUNT(*) as Rank
    FROM 
            (SELECT DISTINCT Score from Scores) as S1,
                    (SELECT DISTINCT Score from Scores) as S2 
                            Where S1.Score<=S2.Score
                                    Group By S1.Score
                                    )
