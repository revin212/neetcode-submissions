-- Write your query below
select name, COALESCE(SUM(distance), 0) AS travelled_distance
from users u
left join rides r on u.id = r.user_id
GROUP BY name
order by travelled_distance desc, name asc