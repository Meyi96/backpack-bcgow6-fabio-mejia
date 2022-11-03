-- Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.
create temporary table `TWD` 
select e.*, se.title season_title from series s
join seasons se on s.id = se.serie_id
join episodes e  on se.id = e.season_id
where s.title = "The Walking Dead";

-- Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.
select * from TWD where season_title = "Primer Temporada";
