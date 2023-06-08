create or replace function Atencion_de_turno(turno_nro int) returns boolean as $$
declare
	t turno%rowtype;
begin
	select * into t from turno where turno_nro = turno.nro_turno;
	
		if not found then
			raise 'número de turno no válido'; 
		else 
			if  not (t.estado = 'reservado') then
				raise 'turno no reservado'; 
			else
				if not (t.fecha::date = current_date) then
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


--select * from turno into t where nro_turno = turno.nro_turno;
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
