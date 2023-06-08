create or replace function Cancelar_turnos(medique_dni int, desde_fecha date, hasta_fecha date) returns int as $$
declare
	
	t turno%rowtype;
	turnos_cancelados int=0;
	
begin
	
	for t in select * from turno, paciente, medique where turno.dni_medique = medique.dni_medique and turno.nro_paciente = paciente.nro_paciente and 
											turno.dni_medique = medique_dni and (turno.estado='disponible' or turno.estado='reservado') loop
		if (t.fecha::date >= desde_fecha and t.fecha::date <= hasta_fecha) then
			
			if (t.estado='reservado')then
				insert into reprogramacion values (t.nro_turno, paciente.nombre, paciente.apellido,paciente.telefono, paciente.email, medique.nombre,medique.apellido, 'pendiente');
				--aca se enviarìa mail
			end if;
			
			update turno set estado='cancelado' where t.nro_turno = turno.nro_turno;
			
			turnos_cancelados = turnos_cancelados + 1;
		end if;	
		
	end loop;
				
				
	return turnos_cancelados;		
	
	--recorre la tabla turno con el rango de fechas dado
	--update de estado = 'cancelado'
	-- insert en tabla reprogramacion
end;
$$ language plpgsql;
