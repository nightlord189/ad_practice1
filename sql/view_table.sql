--создаем представление как OLAP-куб
create or replace view fullSalesView as
select 
	a."name" as article_name, 
	t."name" as type_name, 
	s.price, 
	f."name" as firm_name, 
	d.address, 
	b."date" as bill_date, 
	b.id as bill_id, 
	s.quantity
from sale s 
	join article a on a.id = s.article_id
	join bill b on b.id = s.bill_id
	join drugstore d on d.id = s.drugstore_id
	join firm f on f.id = a.firm_id
	join type t on t.id = a.type_id;

--создаем вспомогательную таблицу из представления
create table if not exists fullSalesTable
(
    article_name varchar (50),
    type_name varchar (50),
    price numeric (6, 2),
    firm_name varchar (50),
    address varchar (50),
    bill_date varchar(50),
    bill_id int,
    quantity int
);
insert into fullSalesTable select * from fullSalesView