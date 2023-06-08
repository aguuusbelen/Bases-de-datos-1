create function Reservar_turno(nro_paciente int, dni_medique int, fecha timestamp) returns boolean as $$
declare
	
begin
	select * from paciente where nro_paciente = paciente.nro_paciente;
	
	if not found then
		raise 'el número de paciente no se encuentra registrado' , nro_paciente;
	end if;
	
	select * from medique where dni_medique = medique.dni_medique;
	
	if not found then
		raise 'no se encontró medique' , dni_medique;
	end if;
	
	select * from paciente, cobertura, medique where (paciente.nro_obra_social = cobertura.nro_obra_social 
	and medique.dni_medique = cobertura.dni_medique);
	
	if not found then
		raise 'el medique seleccionado no cubre la obra social del paciente';
	end if;
	
	--select * from turno 
	--falta verificar turnos pendientes de pacientes
	select * from turno where fecha = turno.fecha and dni_medique = dni_medique and estado = "disponible";
	
	if not found then
		raise 'no hay disponibilidad de turnos para la fecha requerida' , fecha;
	else
		update turno set nro_paciente = nro_paciente, nro_obra_social_consulta = paciente.nro_obra_social,
		nro_afiliade_consulta = paciente.nro_afiliade, monto_paciente = cobertura.monto_paciente,
		monto_obra_social = cobertura.monto_obra_social, f_reserva = time.Now(), estado = "reservado" where (fecha = turno.fecha and dni_medique = dni_medique and estado = "disponible");
	end if;

	
