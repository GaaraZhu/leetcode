# Write your MySQL query statement below
select w.id from Weather w left outer join Weather w2 on DATE_ADD(w.recordDate, INTERVAL -1 DAY)=w2.recordDate where w.Temperature>w2.Temperature