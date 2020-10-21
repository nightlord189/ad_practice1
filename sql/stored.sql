create or replace function get_cube_node(bill integer, article integer, drugstore integer)
returns table (
    article_name varchar (50),
    type_name varchar (50),
    price numeric (6, 2),
    firm_name varchar (50),
    address varchar (50),
    bill_date varchar(50),
    bill_id int,
    quantity int
)
language plpgsql
AS $$
begin
return query 
		select
			*
		from
			fullSalesView v
		where
			v.bill_id=bill
			and v.article_id=article
			and v.drugstore_id=drugstore;
end;
$$;