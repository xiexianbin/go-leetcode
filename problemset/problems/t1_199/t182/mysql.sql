# Write your MySQL query statement below
select Email from (select Email, count(1) as c from person group by Email) as t where t.c > 1;
