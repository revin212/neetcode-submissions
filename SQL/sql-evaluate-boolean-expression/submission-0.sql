-- Write your query below
select left_operand, operator, right_operand,
CASE
    WHEN operator = '='
        THEN CASE WHEN lo.value = ro.value THEN 'true' ELSE 'false' END
    WHEN operator = '>'
        THEN CASE WHEN lo.value > ro.value THEN 'true' ELSE 'false' END
    WHEN operator = '<'
        THEN CASE WHEN lo.value < ro.value THEN 'true' ELSE 'false' END
END AS value
from expressions e
join variables lo on e.left_operand = lo.name
join variables ro on e.right_operand = ro.name
