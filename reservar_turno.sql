create or replace function Reservar_turno(paciente_ID int, medique_dni int, fecha_f timestamp) returns boolean as $$
declare
	p paciente%rowtype;
	me medique%rowtype;
	os record;
	t turno%rowtype;
	cantidad_de_turnos_reservados int;
	
	fecha_actual timestamp := current_date + current_time ;
	
begin
	select * from medique into me where medique_dni = medique.dni_medique;
	
	if not found then
		insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'reserva', fecha_actual, 'dni de médique no válido');
		--raise 'no se encontró medique'; 
		return false;
	end if;
	
	select * from paciente into p where paciente_ID = paciente.nro_paciente;
	
	if not found then
		insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'reserva', fecha_actual, 'número de historia clinica no válido');
		--raise 'el número de paciente no se encuentra registrado';
		return false;
	end if;
	
	select * from paciente, cobertura, medique into os where (paciente.nro_obra_social = cobertura.nro_obra_social 
	and medique.dni_medique = cobertura.dni_medique);
	
	if not found then
		insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'reserva', fecha_actual, 'obra social de paciente no atendida por le médique');
		--raise 'el medique seleccionado no cubre la obra social del paciente';
		return false;
	end if;
	
	select * from turno into t where fecha_f = turno.fecha and medique_dni = turno.dni_medique and estado = 'disponible';
	
	if not found then
		insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'reserva', fecha_actual, 'turno inexistente ó no disponible');
		--raise 'no hay disponibilidad de turnos para la fecha requerida';
		return false;
	end if;
	
	cantidad_de_turnos_reservados := (select count(*) from turno where estado = 'reservado' group by nro_paciente having nro_paciente= paciente_ID) ;
	if cantidad_de_turnos_reservados >=5 then
		insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'reserva', fecha_actual, 'supera el límite de reserva de turnos');
		--raise 'supera el limite de reserva de turnos';
		return false;
	end if;

	update turno set nro_paciente = paciente_ID, nro_obra_social_consulta = p.nro_obra_social,
		nro_afiliade_consulta = p.nro_afiliade, monto_paciente = os.monto_paciente,
		monto_obra_social = os.monto_obra_social, f_reserva = fecha_actual, estado = 'reservado' where (fecha_f = turno.fecha and medique_dni = turno.dni_medique and estado = 'disponible');
	return true;
end;
$$ language plpgsql;
