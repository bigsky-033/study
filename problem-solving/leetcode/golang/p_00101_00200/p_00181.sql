# 181. Employees Earning More Than Their Managers, https://leetcode.com/problems/employees-earning-more-than-their-managers/

SELECT e1.Name as Employee
FROM Employee e1
         LEFT JOIN Employee e2 ON e1.ManagerId = e2.Id
WHERE e1.Salary > e2.Salary