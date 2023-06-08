create or replace function Reservar_turno(paciente_ID int, dni_medique int, fecha timestamp) returns boolean as $$
declare
	p paciente%rowtype;
	me medique%rowtype;
	os record;
	t turno%rowtype;
	cantidad_de_turnos_reservados int;
begin
	select * from paciente into p where paciente_ID = paciente.nro_paciente;
	
	if not found then
		raise 'el número de paciente no se encuentra registrado', paciente_ID;
	end if;
	
	select * from medique into me where dni_medique = medique.dni_medique;
	
	if not found then
		raise 'no se encontró medique', dni_medique; 
	end if;
	
	select * from paciente, cobertura, medique into os where (paciente.nro_obra_social = cobertura.nro_obra_social 
	and medique.dni_medique = cobertura.dni_medique);
	
	if not found then
		raise 'el medique seleccionado no cubre la obra social del paciente', paciente.nro_obra_social;
	end if;
	
	cantidad_de_turnos_reservados := (select count (nro_paciente) from turno group by (estado = 'reservado'));
	if cantidad_de_turnos_reservados >=5 then
		raise 'supera el limite de reserva de turnos', cantidad_de_turnos_reservados;
	end if;

	select * from turno into t where fecha = turno.fecha and dni_medique = dni_medique and estado = 'disponible';
	
	if not found then
		raise 'no hay disponibilidad de turnos para la fecha requerida', fecha;
	end if;
	
	update turno set nro_paciente = paciente_ID, nro_obra_social_consulta = paciente.nro_obra_social,
		nro_afiliade_consulta = paciente.nro_afiliade, monto_paciente = cobertura.monto_paciente,
		monto_obra_social = cobertura.monto_obra_social, f_reserva = fecha, estado = 'reservado' where (fecha = turno.fecha and dni_medique = dni_medique and estado = 'disponible');
	return true;
end;
$$ language plpgsql;
