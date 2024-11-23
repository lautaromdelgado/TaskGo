create database taskgo;

use taskgo;

create table usuario (
	id int auto_increment not null primary key,
    nombre varchar(100) not null,
    apellido varchar(100) not null,
    correo varchar(100) unique not null,
    fecha_alta timestamp default current_timestamp,
    fecha_baja  timestamp default null
);

create table nivel_importancia (
	idnivel int not null auto_increment,
    importancia varchar(25) not null unique,
    primary key (idnivel)
);

create table task (
	idTask int not null auto_increment,
	tittle varchar(100) not null,
    descripcion text,
    nivel int not null,
    idusuario int not null unique,
    primary key (idtask),
    foreign key task(idusuario) references usuario(id),
    foreign key task(nivel) references nivel_importancia(idnivel)
);