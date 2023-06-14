create or replace function envio_mail_update() returns trigger as $$
declare
	body text; 
	subject text;
	med_aux medique%rowtype; 	
	pac_aux paciente%rowtype;
	
begin 
		if (old.estado='disponible' and new.estado='reservado') then --chequeo que haya sido una nueva reserva
		
			select * from medique into med_aux where new.dni_medique = medique.dni_medique; --en med_aux ingreso los datos del medique a cargo de este turno
			select * from paciente into pac_aux where new.nro_paciente = paciente.nro_paciente; --en pac_aux ingreso los datos del paciente de este turno
			
			body:= 'Usted reservo un turno con fecha y hora: ' || to_char(old.fecha,'DD Mon YYYY HH12:MI:SS') || ' con el medique: ' || med_aux.nombre || ' ' || med_aux.apellido;
			subject:= 'Reserva de turno';
			
			insert into envio_email(f_generacion, email_paciente, asunto, cuerpo, f_envio, estado) 
						values (current_date + current_time, pac_aux.email, subject, body, null, 'pendiente');
	
		end if;
		
		if (old.estado='reservado' and new.estado='cancelado') then
			
			select * from medique into med_aux where new.dni_medique = medique.dni_medique;
			select * from paciente into pac_aux where new.nro_paciente = paciente.nro_paciente;
			
			body:= ' Lamentamos informarle que el medique ' || med_aux.nombre || ' tuvo que cancelar su turno el día ' || old.fecha;
			subject:= 'Cancelación de turno';
		
			insert into envio_email(f_generacion, email_paciente, asunto, cuerpo, f_envio, estado) 
						values (current_date + current_time, pac_aux.email, subject, body, null, 'pendiente');
						
		end if;
	return new;
end;
$$ language plpgsql;

create or replace function envio_mail_diario() returns void as $$ --tiene que retornar trigger (no se como triggerear cada x tiempo)
declare
	body text;
	subject text;
	turno_aux turno%rowtype;
	med_aux medique%rowtype;
	pac_aux paciente%rowtype;
	
begin
	
	
	for turno_aux in select * from turno where estado='reservado' and (current_date + interval '2 days')= date_trunc('day',turno.fecha) loop --esta query me da los turnos reservados a 2 días de la fecha actual
		
		select * from medique into med_aux where turno_aux.dni_medique = medique.dni_medique;
		select * from paciente into pac_aux where turno_aux.nro_paciente = paciente.nro_paciente;
		
		body:= 'Le recordamos que su turno con el medique ' || med_aux.nombre || ' ' || med_aux.apellido || ' es el día: ' || turno_aux.fecha;
		subject:= 'Recordatorio de turno';
		insert into envio_email(f_generacion, email_paciente, asunto, cuerpo, f_envio, estado) 
						values (current_date + current_time, pac_aux.email, subject, body, null, 'pendiente');
	end loop;
	
	for turno_aux in select turno.fecha from turno where estado='reservado' and current_date = date_trunc('day',turno.fecha) loop --esta query me da los turnos que pasaron el día de hoy sin pasarse a atendidos
		
		select * from medique into med_aux where turno_aux.dni_medique = medique.dni_medique;
		select * from paciente into pac_aux where turno_aux.nro_paciente = paciente.nro_paciente;
		
		body:= 'Hoy perdió su turno con el medique ' || med_aux.nombre || ' ' || med_aux.apellido || 'del día y hora: ' || turno_aux.fecha;
		subject:= 'Pérdida de turno reservado';
		insert into envio_email(f_generacion, email_paciente, asunto, cuerpo, f_envio, estado) 
						values (current_date + current_time, pac_aux.email, subject, body, null, 'pendiente');
	end loop;	
end;
$$
language plpgsql;

create or replace trigger envio_mail_reserva_trg
after update on turno
for each row
execute function envio_mail_update();


