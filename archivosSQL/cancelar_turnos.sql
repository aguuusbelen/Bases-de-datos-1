create or replace function Cancelar_turnos(medique_dni int, desde_fecha date, hasta_fecha date) returns int as $$
declare
	
	t turno%rowtype;
	turnos_cancelados int=0;
	aux_medique medique%rowtype;
	aux_paciente paciente%rowtype;
	
begin
	--set transaction read only;
	for t in select * from turno where turno.dni_medique = medique_dni and (turno.estado='disponible' or turno.estado='reservado') loop
														
		if (t.fecha::date >= desde_fecha and t.fecha::date <= hasta_fecha) then
	
			if (t.estado='reservado')then
			
				select * into aux_medique from medique where medique.dni_medique=t.dni_medique;

				select * into aux_paciente from paciente where paciente.nro_paciente=t.nro_paciente;  
				
				insert into reprogramacion values (t.nro_turno, aux_paciente.nombre, aux_paciente.apellido, aux_paciente.telefono, 
													aux_paciente.email, aux_medique.nombre,aux_medique.apellido, 'pendiente');
				--aca se enviarìa mail
				
			end if;
			
			update turno set estado='cancelado' where t.nro_turno = turno.nro_turno;
			
			turnos_cancelados = turnos_cancelados + 1;
		end if;	
		
	end loop;
				
				
	return turnos_cancelados;		
	
end;
$$ language plpgsql;



