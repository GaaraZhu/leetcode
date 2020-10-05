# Write your MySQL query statement below
select e.Name as Employee from Employee e left outer join Employee e2 on e.ManagerId = e2.Id where e.Salary>e2.Salary