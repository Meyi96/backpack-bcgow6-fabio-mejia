-- Listar los datos de los autores.
select * from autor;

-- Listar nombre y edad de los estudiantes
select nombre, edad from estudiante;

-- ¿Qué estudiantes pertenecen a la carrera informática?
select * from estudiante where carrera = "informática";

-- ¿Qué autores son de nacionalidad francesa o italiana?
select * from autor where nacionalidad = "francesa" or nacionalidad = "italiana";

-- ¿Qué libros no son del área de internet?
select * from libro where area <> "internet";

-- Listar los libros de la editorial Salamandra.
select * from libro where editorial =  "Salamandra";

-- Listar los datos de los estudiantes cuya edad es mayor al promedio.
select * from estudiante where edad > (select avg(edad)from estudiante); 

-- Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
select nombre from estudiante where apellido like "G%";

-- Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
select a.nombre from autor a 
join libro_autor la on a.id_autor = la.id_autor
join libro l on l.id_libro = la.id_libro
where l.titulo = "El Universo: Guía de viaje";

-- ¿Qué libros se prestaron al lector “Filippo Galli”?
select l.titulo from libro l
join prestamo p on l.id_libro = p.id_libro
join estudiante e on e.id_lector = p.id_lector
where concat(e.nombre, " ", e.apellido) = "Filippo Galli";

-- Listar el nombre del estudiante de menor edad.
select nombre, apellido from estudiante where edad <18;

-- Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
select distinct e.nombre, e.apellido from libro l
join prestamo p on l.id_libro = p.id_libro
join estudiante e on e.id_lector = p.id_lector
where area = "Base de Datos";

-- Listar los libros que pertenecen a la autora J.K. Rowling.
select l.* from autor a 
join libro_autor la on a.id_autor = la.id_autor
join libro l on l.id_libro = la.id_libro
where a.nombre = "J.K. Rowling";

-- Listar títulos de los libros que debían devolverse el 16/07/2021.
select distinct l.titulo from libro l
join prestamo p on l.id_libro = p.id_libro
where fecha_devolucion = "2021-07-16";