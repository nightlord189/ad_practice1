--update discount for bills
with bill_sum as (
	select sum(price*quantity) as amount, bill_id from sale group by bill_id
)

update bill set discount = 
case 
	when (select amount from bill_sum where bill_id=bill.id) > 5000 then 5
	when (select amount from bill_sum where bill_id=bill.id) > 1000 then 2
	else 0
end;

update sale set price=price*(100.0-(select discount from bill where id = sale.bill_id))/100.0;