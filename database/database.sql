
create table `usuarios` (
  `id` integer not null auto_increment primary key,
  `nombre` varchar(30) not null,
  `apellidos` varchar(30) not null,
  `correo` varchar(120),
  `username` varchar(30) not null unique,
  `password` varchar(30) not null
);

create table `actividades` (
  `id` integer not null auto_increment primary key,
  `nombre` varchar(30) not null
);

create table `tareas` (
  `id` integer not null auto_increment primary key,
  `nombre` varchar(30) not null,
  `actividad_id` integer not null,
  foreign key(actividad_id) references actividades(id)
);

