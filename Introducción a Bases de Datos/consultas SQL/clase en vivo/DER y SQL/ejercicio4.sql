#Consultar todos los clientes con todos sus datos.
select * from client;

#Consultar todos los clientes que nacieran en los 90s.
select first_name, last_name from client where year(birth_date) between 1990 and 1999;

#consultar la velocidad de los planes con precio menor a 150.
select speed from plan where price < 150;

#Consultar el precio promedio de los planes.
select avg(price) precio_promedio from plan;

#Consultar la cantidad de planes que estan activos.
select count(*) from service where status = "activo";

#Consultar nombre apellido y la velocidad de todos los clientes que tiene un plan activo.
select c.first_name, c.last_name, p.speed
from client c 
join service s on c.idclient = s.idclient
join plan p on p.idplan = s.idplan
where s.status = "activo";

#Consultar los clientes que tiene un plan de precio mayor 150.
select c.first_name, c.last_name
from client c 
join service s on c.idclient = s.idclient
join plan p on p.idplan = s.idplan
where p.price > 150;

#Consultar los clientes y el valor total que pagan por todos los servicios activos.
select c.first_name, c.last_name, sum(s.discount*p.price) total_pago
from client c 
join service s on c.idclient = s.idclient
join plan p on p.idplan = s.idplan
where s.status = "activo"
group by c.first_name, c.last_name;

#Consultar los clientes y la cantidad servicios activos contratados que san mayor a 1.
select c.first_name, c.last_name, count(s.idservice) servicios
from client c 
join service s on c.idclient = s.idclient
join plan p on p.idplan = s.idplan
where s.status = "activo"
group by c.first_name, c.last_name
having servicios > 1;