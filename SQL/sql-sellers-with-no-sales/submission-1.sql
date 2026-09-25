-- Write your query below
select seller_name
from seller s
left join orders o on o.seller_id = s.seller_id AND EXTRACT(YEAR FROM sale_date) = '2020'
where o.order_id IS NULL
order by seller_name asc