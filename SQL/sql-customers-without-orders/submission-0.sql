-- Write your query below
select c.name
from customers c
left join orders o ON c.id = o.customer_id
where o.id IS NULL