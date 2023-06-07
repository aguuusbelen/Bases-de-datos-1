create function Generar_turnos_disponibles(anio int, mes int) returns boolean as $$
declare
	turnos_en_fecha record;
	nro_turno int;
	t agenda%rowtype;
	
begin
	-- Si hay turnos creados para ese año y mes, devuelve false.
	if exists (
		select * into turnos_en_fecha from turno where 
		((select extract (year from timestamp turno.fecha))= anio and (select extract (month from timestamp turno.fecha)) = mes) -- Comparo el año y el mes con la fecha del turno
		) then return false;
	end if;
	for t in agenda loop -- para recorrer la tabla agenda? 
		nro_turno:= 1;
		insert into turno values (nro_turno, t.fecha , t.nro_consultorio, t.dni_medique, null, 
		null, null, null, null, "disponible") -- Insertar datos en la tabla turno
		nro_turno:= nro_turno +1;
	end loop;
	return true;
	
end;
$$ language plpgsql;
