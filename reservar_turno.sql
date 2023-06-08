create or replace function Reservar_turno(paciente_ID int, medique_dni int, fecha_f timestamp) returns boolean as $$
declare
	p paciente%rowtype;
	me medique%rowtype;
	os record;
	t turno%rowtype;
	cantidad_de_turnos_reservados int;
	
	fecha_actual date := current_date;
	
begin
	select * from paciente into p where paciente_ID = paciente.nro_paciente;
	
	if not found then
		raise 'el número de paciente no se encuentra registrado';
	end if;
	
	select * from medique into me where medique_dni = medique.dni_medique;
	
	if not found then
		raise 'no se encontró medique'; 
	end if;
	
	select * from paciente, cobertura, medique into os where (paciente.nro_obra_social = cobertura.nro_obra_social 
	and medique.dni_medique = cobertura.dni_medique);
	
	if not found then
		raise 'el medique seleccionado no cubre la obra social del paciente';
	end if;
	
	cantidad_de_turnos_reservados := (select count (nro_paciente = paciente_ID) from turno group by (estado = 'reservado'));
	if cantidad_de_turnos_reservados >=5 then
		raise 'supera el limite de reserva de turnos';
	end if;

	select * from turno into t where fecha_f = turno.fecha and medique_dni = turno.dni_medique and estado = 'disponible';
	
	if not found then
		raise 'no hay disponibilidad de turnos para la fecha requerida';
	end if;
	
	update turno set nro_paciente = paciente_ID, nro_obra_social_consulta = p.nro_obra_social,
		nro_afiliade_consulta = p.nro_afiliade, monto_paciente = os.monto_paciente,
		monto_obra_social = os.monto_obra_social, f_reserva = fecha_actual, estado = 'reservado' where (fecha_f = turno.fecha and medique_dni = turno.dni_medique and estado = 'disponible');
	return true;
end;
$$ language plpgsql;
