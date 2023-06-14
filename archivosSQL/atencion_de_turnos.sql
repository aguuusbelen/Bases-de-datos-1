create or replace function Atencion_de_turno(turno_nro int) returns boolean as $$
declare
	t turno%rowtype;
	fecha_actual timestamp := current_date + current_time ;
	
begin
	select * from turno into t where turno_nro = turno.nro_turno;
	
		if not found then
			insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'atención', fecha_actual, 'número de turno no válido');
			raise 'número de turno no válido'; 
			return false;
		else 
			if  not (t.estado = 'reservado') then
				insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'atención', fecha_actual, 'turno no reservado');
				raise 'turno no reservado'; 
				return false;
			else
				if not (t.fecha::date = current_date) then
					insert into error (f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
								values (t.fecha, t.nro_consultorio, t.dni_medique, t.nro_paciente, 'atención', fecha_actual, 'turno no corresponde a la fecha del día');
					raise 'turno no corresponde a la fecha del dia';
					return false;
				else
					update turno set estado = 'atendido' where turno.nro_turno = turno_nro;
					return true;
				end if;
			end if;
		end if;
		
	
end;
$$ language plpgsql;
