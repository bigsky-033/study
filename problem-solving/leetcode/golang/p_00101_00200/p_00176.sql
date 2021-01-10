# 176. Second Highest Salary, https://leetcode.com/problems/second-highest-salary/

SELECT max(salary) SecondHighestSalary
FROM employee
WHERE salary < (SELECT MAX(salary) FROM employee)