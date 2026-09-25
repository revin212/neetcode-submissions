-- Write your query below
select employee_id, 
CASE
    WHEN employee_id%2 = 0 OR SUBSTRING(name, 1, 1) = 'M'
        THEN 0
    ELSE
        salary
END AS bonus
from employees
ORDER BY employee_id