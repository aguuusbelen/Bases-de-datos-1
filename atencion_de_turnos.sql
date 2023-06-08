create or replace function Atencion_de_turno(nro_turno int) returns boolean as $$
declare
	t turno%rowtype;
begin
	t := (select * from turno where nro_turno = turno.nro_turno);
	
		if not t then
			raise 'número de turno no válido'; 
		else 
			if  not (t.estado = 'reservado') then
				raise 'turno no reservado'; 
			else
				if not (t.fecha = localtimestamp) then
					raise 'turno no corresponde a la fecha del dia';
					return false;
				else
					update turno set estado = 'atendido' where turno.nro_turno = nro_turno;
					return true;
			end if;
		end if;
		
	
end;
$$ language plpgsql;


--select * from turno where nro_turno = turno.nro_turno;
--	
--		if not found then
--			raise 'número de turno no válido'; 
--		else 
--			if  not (nro_turno = turno.nro_turno and estado = 'reservado') then
--				raise 'turno no reservado'; 
--			else
--				if not (nro_turno = turno.nro_turno and estado = 'reservado' and 
--			end if;
--		end if;				
