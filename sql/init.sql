--вставляем лекарства
INSERT INTO type (id,name) VALUES 
(1,'антибиотик'),
(2,'витамины'),
(3,'жаропонижающее'),
(4,'обезболивающее'),
(5,'противодиарейное'),
(6,'от кашля'),
(7,'против гриппа');

--генерируем фирмы
insert into firm (
    id, name
)
select
    i,
    concat('Фирма ', i::text)
from generate_series(1, 20) s(i);

--генерируем аптеки
insert into drugstore (
    id, address, phone
)
select
    i,
    concat('Адрес ', i::text),
    ceiling(random()*999999+100)
from generate_series(1, 20) s(i);