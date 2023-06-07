create function Generar_turnos_disponibles(anio int, mes int) returns boolean as $$
declare
	turnos_en_fecha record;
	nro_turno int;
	aux_horario time;
	a agenda%rowtype;
	
	fecha_aux_inicio timestamp := date_trunc('month', make_date(anio,mes,1));
	fecha_aux_final timestamp := fecha_aux_inicio + interval '1 month' - interval '1 day';
	fecha_aux_actual timestamp;
begin	
	-- Si hay turnos creados para ese año y mes, devuelve false.
	if exists (
		select * from turno where 
		((select extract (year from turno.fecha))= anio and (select extract (month from turno.fecha)) = mes) -- Comparo el año y el mes con la fecha del turno
	) then return false;
	end if;
	nro_turno:=1;
	for a in select * from agenda loop -- para recorrer la tabla agenda
		
		fecha_aux_actual:= fecha_aux_inicio;
		
		while fecha_aux_actual <= fecha_aux_final loop
			
			if (select extract (isodow from fecha_aux_actual)) = a.dia then
				aux_horario:= a.hora_desde;
				while aux_horario <= a.hora_hasta - a.duracion_turno loop
			
						insert into turno values(nro_turno, fecha_aux_actual + aux_horario, a.nro_consultorio, a.dni_medique, null, null, null, null, null, null, 'disponible');
						nro_turno:= nro_turno +1; 
						aux_horario := aux_horario + a.duracion_turno;
				
				end loop;
			end if;
			fecha_aux_actual=fecha_aux_actual + interval '1 day';
			
		end loop;
		
	end loop;
	return true;
	
end;
$$ language plpgsql;
