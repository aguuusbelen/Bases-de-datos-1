--Stored procedure para reservar los turnos cargados en la tabla solicitud_reservas
create or replace function test_reservas() returns void as $$

declare
	
	res_aux solicitud_reservas%rowtype;
	tur_aux turno%rowtype;
	time_aux timestamp;
	b boolean;

begin 
	for res_aux in (select * from solicitud_reservas order by nro_orden) loop --se reservan por nro_orden
		time_aux :=	res_aux.fecha + res_aux.hora;
		perform reservar_turno(res_aux.nro_paciente, res_aux.dni_medique,time_aux);
	end loop;
end;
$$ language plpgsql;

--Carga de datos de prueba en la tabla solicitud_reservas
insert into solicitud_reservas values(2, 1, 31759846, '2023-06-15','12:00'); --Mismo turno que el de nro_orden = 1
insert into solicitud_reservas values(1, 2, 31759846, '2023-06-15','12:00'); --Mismo turno que el de nro_orden = 2
insert into solicitud_reservas values(3, 6, 31759846, '2023-06-15','17:20');
insert into solicitud_reservas values(5, 5, 20147852, '2023-06-15','10:00');
insert into solicitud_reservas values(4, 14, 20147852, '2023-06-22','10:20');
insert into solicitud_reservas values(8, 8, 29541019, '2023-06-16','13:00');
insert into solicitud_reservas values(6, 13,31759846, '2023-06-19','10:00');
insert into solicitud_reservas values(9, 4, 30668951, '2023-06-19','15:15');
insert into solicitud_reservas values(7, 15, 30668951, '2023-06-21','12:00');
insert into solicitud_reservas values(10, 16, 27401511, '2023-06-27','10:40');
insert into solicitud_reservas values(11, 6,18945123, '2023-06-20', '15:30'); --Turno invalido:Medico no atiende obra social del paciente
insert into solicitud_reservas values(12, 20,20147852, '2023-06-21', '10:40'); --Turno invalido: Dia fuera de atencion
insert into solicitud_reservas values(13, 8,21036258, '2023-06-16', '10:30'); --Turno invalido: Horario fuera de atencion
insert into solicitud_reservas values(14, 8,21036258, '2023-06-16', '12:10'); --Turno invalido: Horario fuera de intervalos de atencion
insert into solicitud_reservas values(15, 8,21036258, '2023-06-16', '18:00'); --Turno invalido: Horario fuera de atencion/limite 
insert into solicitud_reservas values (15, 2, 30668951, '2023-06-14', '09:00');



--Se ejecuta la funcion
select test_reservas();
