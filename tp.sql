create table paciente(
nro_paciente int, --número de historia clínica
nombre text,
apellido text,
dni_paciente int,
f_nac date,
nro_obra_social int,
nro_afiliade int,
domicilio text,
telefono char(12),
email text -- válido
);

create table medique(
dni_medique int,
nombre text,
apellido text,
especialidad varchar(64),
monto_consulta_privada decimal(12,2), --para pacientes sin obra social
telefono char(12)
);

create table consultorio(
nro_consultorio int,
nombre text,
domicilio text,
codigo_postal char(8),
telefono char(12)
);

create table agenda(
dni_medique int,
dia int, --1:lunes, 2:martes...
nro_consultorio int,
hora_desde time, --08:45, 11:30...
hora_hasta time,
duracion_turno interval
);

create table turno(
nro_turno int,
fecha timestamp,
nro_consultorio int,
dni_medique int,
nro_paciente int,
nro_obra_social_consulta int, --para las liquidaciones
nro_afiliade_consulta int,
monto_paciente decimal(12,2),
monto_obra_social decimal(12,2), --para las liquidaciones
f_reserva timestamp,
estado char(10) --`disponible',`reservado',`cancelado',`atendido',`liquidado'
);

create table reprogramacion(
nro_turno int,
nombre_paciente text,
apellido_paciente text,
telefono_paciente char(12),
email_paciente text,
nombre_medique text,
apellido_medique text,
estado char(12) --`pendiente', `reprogramado', `desistido'
);

create table error(
nro_error int,
f_turno timestamp,
nro_consultorio int,
dni_medique int,
nro_paciente int,
operacion char(12), --`reserva', `cancelación', `atención', `liquidación'
f_error timestamp,
motivo varchar(64)
);

create table cobertura(
dni_medique int,
nro_obra_social int,
monto_paciente decimal(12,2), --monto a abonar por el paciente
monto_obra_social decimal(12,2) --monto a liquidar a la obra social
);

create table obra_social (
nro_obra_social int,
nombre text,
contacto_nombre text,
contacto_apellido text,
contacto_telefono char(12),
contacto_email text
);

create table liquidacion_cabecera(
nro_liquidacion int,
nro_obra_social int,
desde date,
hasta date,
total decimal(15,2)
);

create table liquidacion_detalle(
nro_liquidacion int,
nro_linea int,
f_atencion date,
nro_afiliade int,
dni_paciente int,
nombre_paciente text,
apellido_paciente text,
dni_medique int,
nombre_medique text,
apellido_medique text,
especialidad varchar(64),
monto decimal(12,2)
);

create table envio_email(
nro_email int,
f_generacion timestamp,
email_paciente text,
asunto text,
cuerpo text,
f_envio timestamp,
estado char(10) --`pendiente', `enviado'
);

create table solicitud_reservas(
nro_orden int,
nro_paciente int,
dni_medique int,
fecha date,
hora time
);

alter table paciente add constraint paciente_pk primary key (nro_paciente);
alter table medique add constraint medique_pk primary key (dni_medique);
alter table consultorio add constraint consultorio_pk primary key (nro_consultorio);
alter table agenda add constraint agenda_pk primary key (dni_medique, dia);
alter table turno add constraint turno_pk primary key (nro_turno);
alter table reprogramacion add constraint reprogramacion_pk primary key (nro_turno);
alter table error add constraint error_pk primary key (nro_error);
alter table cobertura add constraint cobertura_pk primary key (dni_medique, nro_obra_social);
alter table obra_social add constraint obra_social_pk primary key (nro_obra_social);
alter table liquidacion_cabecera add constraint liquidacion_cabecera_pk primary key (nro_liquidacion);
alter table liquidacion_detalle add constraint liquidacion_detalle_pk primary key (nro_liquidacion, nro_linea);
alter table envio_email add constraint envio_email_pk primary key (nro_email);

alter table paciente add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social(nro_obra_social);

alter table agenda add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table agenda add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);

alter table turno add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table turno add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table turno add constraint dni_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);
alter table turno add constraint nro_obra_social_consulta_fk foreign key (nro_obra_social_consulta) references obra_social (nro_obra_social); --duda si es necesario

alter table reprogramacion add constraint nro_turno_fk foreign key (nro_turno) references turno (nro_turno);

alter table error add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table error add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table error add constraint dni_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);

alter table cobertura add constraint dni_medique_fk  foreign key (dni_medique) references medique (dni_medique);
alter table cobertura add constraint nro_obra_social_fk  foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table liquidacion_cabecera add constraint nro_obra_social_fk  foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table liquidacion_detalle add constraint nro_liquidacion_fk  foreign key (nro_liquidacion) references liquidacion_cabecera (nro_liquidacion);
alter table liquidacion_detalle add constraint dni_medique_fk  foreign key (dni_medique) references medique (dni_medique);

insert into obra_social values( 100, 'OSDE', 'Jorge', 'Osdelaga', 1164358292, 'jorgeosde@yahoo.com');
insert into obra_social values( 200, 'Medifé', 'Ricardo', 'Lopez', 1141425870, 'rickylopez@hotmail.com');
insert into obra_social values( 300, 'Avalian', 'Arnaldo', 'Gomez Castro', 1171824577, 'arnaldo_gs@gmail.com');
insert into obra_social values( 400, 'Swiss Medical', 'Pedro', 'Rossi', 1164358292, 'pedrorossi@hotmail.com');

insert into paciente values ( 1, 'Bryan', 'Cranston', 16774869, '1956-03-07', 100, 1001, 'Pte. Peron 123, Grand Bourg', '1143527303', 'bryancranston@gmail.com');
insert into paciente values ( 2, 'Aaron', 'Paul', 33651237, '1979-08-27', null, null,'Rivadavia 1573, Polvorines', '1124756891', 'aaronp@yahoo.com'); 
insert into paciente values ( 3, 'Bob', 'Odenkirk', 21444139, '1962-10-22', 200, 2001, 'Paunero 374, San Miguel', '1147658473', 'bobodenkirk@hotmail.com.com');
insert into paciente values ( 4, 'Jennifer', 'Aniston', 24541769, '1969-02-11', 400, 4001, 'Segurola 1734, CABA', '1161859262', 'jenniferaniston@hotmail.com');
insert into paciente values ( 5, 'Jonathan', 'Banks', 11134762, '1947-01-31', 200, 2002, 'Miranda 1449, Jose C. Paz', '1172689412', 'jonathanbanks@gmail.com');
insert into paciente values ( 6, 'Giancarlo', 'Esposito', 20912476, '1958-04-26', 300, 3001, 'Aristobulo del valle 632, San Miguel', '1172345182', 'gcesposito@gmail.com');
insert into paciente values ( 7, 'Antonio', 'Dalton', 27731465, '1975-02-13', null, null, 'Maipu 732, Vicente Lopez', '1191763242', 'tonydalton@gmail.com');
insert into paciente values ( 8, 'Uma', 'Thurman', 25305811, '1970-04-29', 400, 4002, 'Darragueira 2362, Polvorines', '1174836792', 'umathurman@hotmail.com');
insert into paciente values ( 9, 'Tim', 'Roth', 20168762, '1961-05-14', 300, 3002, 'Saavedra 473, San Miguel', '1184679516', 'timroth@yahoo.com');
insert into paciente values ( 10, 'Anya', 'Taylor-Joy', 38576183, '1996-04-16', 200, 2003, 'Jose Ingenieros 4734, CABA', '1184596235', 'anyatj@gmail.com');
insert into paciente values ( 11, 'Scarlett', 'Johansson', 33473891, '1984-11-22', null, null, 'Compostela 523, Jose C. Paz', '1124759681', 'bwillis@gmail.com');
insert into paciente values ( 12, 'Salma', 'Hayek', 23142673, '1966-09-02', 100, 1002, 'Azcuenaga 1348, San Miguel', '1143527303', 'bryancranston@gmail.com');
insert into paciente values ( 13, 'Antonio', 'Banderas', 19846735, '1960-08-10', 100, 1003, 'Misiones 374, San Isidro', '1174859623', 'toniobanderas@hotmail.com');
insert into paciente values ( 14, 'Penélope', 'Cruz', 26843951, '1974-04-28', 200, 2004, 'Maipu 2479, Polvorines', '1184597486', 'pencruz@gmail.com');
insert into paciente values ( 15, 'Channing', 'Tatum', 34946851, '1980-04-26', 100, 1004, 'Las Heras 539, San Miguel', '1141546874', 'channingtatum@yahoo.com');
insert into paciente values ( 16, 'Adam', 'Sandler', 22718253, '1966-09-09', null, null, 'Buenos Aires 1656, Don Torcuato', '1144526398', 'adamsandler@gmail.com');
insert into paciente values ( 17, 'Ben', 'Stiller', 19745852, '1960-05-11', 200, 2005, 'Eva Peron 1345, Grand Bourg', '1158759623', 'benstiller@hotmail.com');
insert into paciente values ( 18, 'Angelina', 'Jolie', 27673951, '1975-06-04', 300, 3003, 'Corrientes 1256, San Miguel', '1174859613', 'angjolie@gmail.com');
insert into paciente values ( 19, 'John', 'Travolta', 14723441, '1954-02-18', 300, 3004, 'Santa Fe 2523, CABA', '1144526344', 'johntravolta@gmail.com');
insert into paciente values ( 20, 'Brad', 'Pitt', 21739852, '1963-12-18', 200, 2006, 'Pte. Peron 433, San Miguel', '1122536892', 'bradp@yahoo.com');

insert into medique values(31759846, 'Juan','Perez','clínico', 514.5, 1174859623);
insert into medique values(28455749, 'Pedro','Corvalan','clínico', 423.95, 1152869341);
insert into medique values(30668951, 'Alfredo','Mendez','cardiología', 1023.5, 1147852639);
insert into medique values(19512639, 'Claudia','Pacheco','clínico', 723.31, 1142758693);
insert into medique values(25748596, 'Esteban','Roig','clínico', 416.96, 1152647382);
insert into medique values(30205816, 'Beatriz','Mendieta','clínico', 1569.33, 1194758623);
insert into medique values(33850951, 'Juan Carlos','Caputo','clínico', 963.70, 1134526282);
insert into medique values(20147852, 'Alberto','Semit','ginecología', 4347, 1144552368);
insert into medique values(18945123, 'Javier','Greco','clínico', 698.83, 1155237896);
insert into medique values(27401511, 'Ana Maria','Lopez','clínico', 752.13, 11012586);
insert into medique values(29564812, 'Augusto','Gutierrez','clínico', 968.8, 11602539);
insert into medique values(30514789, 'Franco','Campanella','clínico', 421, 11547082);
insert into medique values(18308511, 'Lucia','Ricardini','osteopatía', 372.7, 1122526230);
insert into medique values(22174850, 'Juan Jose','Di Marco','clínico', 469.75, 1171815020);
insert into medique values(25013965, 'Walter','Gomez','clínico', 961.23, 1195648023);
insert into medique values(21980541, 'Carlos','Scasso','clínico', 847.4, 1175958072);
insert into medique values(26845623, 'Melisa','Lissi','clínico', 375, 1140536292);
insert into medique values(21036258, 'Juan Martin','Delgado','clínico', 400, 1177839254);
insert into medique values(31540210, 'Luciana','Bentancur','clínico', 500, 1137759062);
insert into medique values(29541019, 'Romina','Ricci','clínico', 600, 1141489806);
