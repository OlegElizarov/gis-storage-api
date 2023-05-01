create table if not exists tree_data
(
    id               bigserial primary key,
    gis_id           integer unique,
    pcd_name         varchar unique,
    x_coordinate     numeric,
    y_coordinate     numeric,
    gis_height_mitro numeric,
    gis_height_il    numeric,
    order_number     integer unique,
    tree_type        varchar,
    circle           numeric,
    diameter_mitro   numeric,
    diameter_il      numeric
);

create table if not exists tree_growth
(
    id       bigserial primary key,
    gis_id   integer references tree_data (gis_id),
    ts       timestamp,
    age      integer,
    diameter numeric,
    height   numeric,
    is_alive boolean
);

select id
from tree_data
limit 1;
