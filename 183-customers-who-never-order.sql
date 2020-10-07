# Write your MySQL query statement below
select c.Name as Customers from Customers c where not exists (select * from Orders o where o.CustomerId = c.Id)