# 183. Customers Who Never Order, https://leetcode.com/problems/customers-who-never-order/

select Name as Customers
from Customers
where Id not in (select distinct CustomerId from Orders)