insert into solicitud_reservas values(2, 1, 31759846, '2023-06-15','12:00');
insert into solicitud_reservas values(1, 2, 31759846, '2023-06-15','12:00');
insert into solicitud_reservas values(3, 6, 31759846, '2023-06-15','17:20');
insert into solicitud_reservas values(5, 5, 20147852, '2023-06-15','10:00');
insert into solicitud_reservas values(4, 14, 20147852, '2023-06-22','10:20');
insert into solicitud_reservas values(8, 8, 29541019, '2023-06-16','13:00');
insert into solicitud_reservas values(6, 13,31759846, '2023-06-19','10:00');
insert into solicitud_reservas values(9, 4, 30668951, '2023-06-19','15:15');
insert into solicitud_reservas values(7, 15, 30668951, '2023-06-21','12:00');
insert into solicitud_reservas values(10, 16, 27401511, '2023-06-27','10:40');

create or replace function test_reservas() returns void as $$
declare
	res_aux solicitud_reservas%rowtype;
	tur_aux turno%rowtype;
	time_aux timestamp;
	b boolean;

begin
	
	for res_aux in select * from solicitud_reservas loop
		time_aux :=	res_aux.fecha + res_aux.hora;
		perform reservar_turno(res_aux.nro_paciente, res_aux.dni_medique,time_aux);
	end loop;
end;
$$ language plpgsql;

select test_reservas();
