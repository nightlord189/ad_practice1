-- сколько единиц каждого товара и на какую сумму продано от каждой фарм. компании? 
select 
    firm_name, 
    sum(quantity) as quantity, 
    sum(quantity*price) as income
from fullSalesView
group by firm_name

--сколько единиц лекарств от кашля и на какую сумму было продано в Аптеке 1?
select 
    sum(quantity) as quantity, 
    sum(quantity*price) as income
from fullSalesView
where address = 'Адрес 1'
and type_name = 'от кашля'

--сколько единиц лекарств от кашля и на какую сумму было продано в Аптеке 1 в разрезе по годам?
select 
    substring(bill_date from 0 for 5) as year, 
    sum(quantity) as quantity, 
    sum(quantity*price) as income
from fullSalesView
where address = 'Адрес 1'
and type_name = 'от кашля'
group by substring(bill_date from 0 for 5)

--на какую сумму и каких типов было куплено лекарств в Аптеке 2 зимой 2011?
select 
    type_name,
    sum (quantity*price) as income
from fullSalesView
where address = 'Адрес 2'
and substring (bill_date from 0 for 8) in ('2011-01', '2011-02', '2011-12')
group by type_name

--сколько и на какую сумму было продано лекарств от каждой фарм. компании в Аптеке 3 весной 2012?
select 
    firm_name,
    sum (quantity) as quantity,
    sum (quantity*price) as income
from fullSalesView
where address = 'Адрес 3'
and substring (bill_date from 0 for 8) in ('2012-03', '2012-04', '2012-05')
group by firm_name

--кросс-таблица, где столбцы - месяцы, а в строках - типы товаров и суммы продаж по месяцам
select * from crosstab ('select 
	type_name, 
	substring (bill_date from 6 for 2) as month, 
	sum (quantity*price)::text as income
from fullSalesView
group by type_name, substring (bill_date from 6 for 2)
order by type_name, month
') AS ct ("type_name" text, "month" text, "income" text);
