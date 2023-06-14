alter table paciente add constraint paciente_pk primary key (nro_paciente);
alter table medique add constraint medique_pk primary key (dni_medique);
alter table consultorio add constraint consultorio_pk primary key (nro_consultorio);
alter table agenda add constraint agenda_pk primary key (dni_medique, dia);
alter table turno add constraint turno_pk primary key (nro_turno);
alter table reprogramacion add constraint reprogramacion_pk primary key (nro_turno);
alter table error add constraint error_pk primary key (nro_error);
alter table cobertura add constraint cobertura_pk primary key (dni_medique, nro_obra_social);
alter table obra_social add constraint obra_social_pk primary key (nro_obra_social);
alter table liquidacion_cabecera add constraint liquidacion_cabecera_pk primary key (nro_liquidacion);
alter table liquidacion_detalle add constraint liquidacion_detalle_pk primary key (nro_liquidacion, nro_linea);
alter table envio_email add constraint envio_email_pk primary key (nro_email);

alter table paciente add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social(nro_obra_social);

alter table agenda add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table agenda add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);

alter table turno add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table turno add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table turno add constraint dni_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);
alter table turno add constraint nro_obra_social_consulta_fk foreign key (nro_obra_social_consulta) references obra_social (nro_obra_social); --duda si es necesario

alter table reprogramacion add constraint nro_turno_fk foreign key (nro_turno) references turno (nro_turno);

alter table error add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table error add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table error add constraint dni_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);

alter table cobertura add constraint dni_medique_fk  foreign key (dni_medique) references medique (dni_medique);
alter table cobertura add constraint nro_obra_social_fk  foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table liquidacion_cabecera add constraint nro_obra_social_fk  foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table liquidacion_detalle add constraint nro_liquidacion_fk  foreign key (nro_liquidacion) references liquidacion_cabecera (nro_liquidacion);
alter table liquidacion_detalle add constraint dni_medique_fk  foreign key (dni_medique) references medique (dni_medique);
