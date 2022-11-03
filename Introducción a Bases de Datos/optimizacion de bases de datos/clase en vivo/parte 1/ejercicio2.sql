-- En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. 
alter table series
add index title_index (title);

show index from series;

-- Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.
/*
	Elegi el titulo de las series, es muy comun que las personas busquen series con este atributo de la entidad.
    Es importante resaltar que el index no es unico, pues el titulo puede ser el mismo para varias series.
    Esto mejora los tiempos de busqueda del usuario final.
*/