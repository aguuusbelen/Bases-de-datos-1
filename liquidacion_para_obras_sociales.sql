create or replace function Liquidacion_para_obras_sociales(anio int, mes int, numero_obra_social int) returns decimal as $$
declare

	t_aux turno%rowtype;
	medique_aux medique%rowtype;
	paciente_aux paciente%rowtype;
	fecha_aux_inicio date := date_trunc('month', make_date(anio,mes,1));
	fecha_aux_final date := fecha_aux_inicio + interval '1 month' - interval '1 day';
	monto_liquidacion decimal(15,2);
	
begin
	-- si ya esta liquidado
	if exists (
		select * from turno where 
			(turno.fecha ::date >= fecha_aux_inicio  and turno.fecha ::date <= fecha_aux_final and turno.estado='liquidado')
		) then return 0;
	end if;
	-- en caso que no este liquidado
	for t_aux in select * from turno where turno.nro_obra_social_consulta = numero_obra_social and turno.estado='atendido' loop
														
		if (t_aux.fecha ::date >= fecha_aux_inicio  and t_aux.fecha ::date <= fecha_aux_final) then
			
			select * into medique_aux from medique where t_aux.dni_medique=medique.dni_medique;

			select * into paciente_aux from paciente where t_aux.nro_paciente=paciente.nro_paciente;
			
			update turno set estado='liquidado' where turno.nro_turno=t_aux.nro_turno;
			
			insert into liquidacion_detalle (f_atencion,nro_afiliade, dni_paciente, nombre_paciente, apellido_paciente, dni_medique,
				nombre_medique, apellido_medique,especialidad, monto)
				values(t_aux.fecha :: date, t_aux.nro_afiliade_consulta, paciente_aux.dni_paciente, paciente_aux.nombre, 
				paciente_aux.apellido,t_aux.dni_medique, medique_aux.nombre, medique_aux.apellido,medique_aux.especialidad, 
				t_aux.monto_obra_social);
				
			monto_liquidacion=monto_liquidacion + t_aux.monto_obra_social; 
			
		end if;	
	
	end loop;
	
	insert into liquidacion_cabecera (nro_obra_social, desde, hasta, total) values
			(numero_obra_social, fecha_aux_inicio, fecha_aux_final, monto_liquidacion);

	return monto_liquidacion;
	
end;
$$ language plpgsql;
