-- Write your query below
select r.customer_id, r.customer_name
FROM (select c.customer_id, c.customer_name, 
Count(case product_name when 'A' then 1 else null end) AS ProductACount, 
Count(case product_name when 'B' then 1 else null end) AS ProductBCount,
Count(case product_name when 'C' then 1 else null end) AS ProductCCount
from customers c
left join orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id) r
where r.ProductACount > 0 AND r.ProductBCount > 0 AND r.ProductCCount = 0
ORDER BY r.customer_name