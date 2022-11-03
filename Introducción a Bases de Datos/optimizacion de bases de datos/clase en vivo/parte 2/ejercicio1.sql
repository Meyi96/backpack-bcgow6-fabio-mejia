-- Agregar un género a la tabla genres.
insert into genres (created_at, name, ranking, active) values ("2022-11-03","Tecnologico", 13,true);

-- Agregar una película a la tabla movies.
-- Asociar a la película del punto 1. con el género creado en el punto 2.
insert into movies (title,rating, awards,release_date,length,genre_id) values ("Avatar 2",9.4,12,"2022-12-15",180,13);

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
update actors 
set favorite_movie_id = 13
where id in (select * from (select id from actors where first_name like "A%") as subquery);

-- Crear una tabla temporal copia de la tabla movies.
create temporary table movies_copy select * from movies; 

-- Obtener la lista de todos los géneros que tengan al menos una película.
select distinct g.name from genres g join movies m on g.id = m.genre_id;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
select a.first_name, a.last_name from actors a 
join movies m on a.favorite_movie_id = m.id 
where m.awards > 3;

-- Crear un índice sobre el nombre en la tabla movies.
alter table movies
add index title_index (title);

-- Chequee que el índice fue creado correctamente.
show index from movies;

-- En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
/*
	efectivamente hay una mejora en las busquedas por titulo de pelicula, la cual es una consulta muy recurrente 
    en la natura del problema pues no tendra que iterar sobre toda las peliculas que pueda tener la  DB
*/

-- ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
/*
	El indice seria una buena idea para los titulos de las series, peliculas, tal vez de los episodios y 
    por ultimo el nombre de los actores ya que stas seran las busquedas mas frecuentes por el usuario final, 
    seran mucho mas rapoidas aplicado indices. 
*/

