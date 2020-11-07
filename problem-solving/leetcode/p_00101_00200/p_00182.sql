# 182, Duplicate Emails https://leetcode.com/problems/duplicate-emails/

SELECT Email
FROM Person
GROUP BY Email
HAVING count(Email) > 1