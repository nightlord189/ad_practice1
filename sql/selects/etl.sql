--удаляем все записи, кроме данных за 2011 и 2012
delete from fullsalestable where substring(bill_date from 0 for 5) not in ('2012', '2011')

--обновляем данные по фирмам без названия
update firm set name = concat('Фирма ', id::text) 
where name is null or name =''

--обновляем данные по праздничным скидкам в определенной аптеке за 30 декабря 2012
update bill set discount=15
where id in (select bill_id from sale where drugstore_id=5)
and date='2012-12-29';

--удаляем записи о продажах лекарств от производителя - фирмы 3
alter table bill disable trigger all;
alter table sale disable trigger all;
delete from bill where id in 
(select distinct bill_id from sale where article_id in (select id from article where firm_id=3));
delete from sale where article_id in (select id from article where firm_id=3);
alter table bill enable trigger all;
alter table sale enable trigger all;