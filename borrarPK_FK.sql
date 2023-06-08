alter table paciente drop constraint nro_obra_social_fk;
alter table agenda drop constraint dni_medique_fk;
alter table agenda drop constraint nro_consultorio_fk;
alter table turno drop constraint nro_consultorio_fk;
alter table turno drop constraint dni_medique_fk;
alter table turno drop constraint dni_paciente_fk;
alter table turno drop constraint nro_obra_social_consulta_fk;
alter table reprogramacion drop constraint nro_turno_fk;
alter table error drop constraint nro_consultorio_fk;
alter table error drop constraint dni_medique_fk;
alter table error drop constraint dni_paciente_fk;
alter table cobertura drop constraint dni_medique_fk;
alter table cobertura drop constraint nro_obra_social_fk;
alter table liquidacion_cabecera drop constraint nro_obra_social_fk;
alter table liquidacion_detalle drop constraint nro_liquidacion_fk;
alter table liquidacion_detalle drop constraint dni_medique_fk;

alter table paciente drop constraint paciente_pk;
alter table medique drop constraint medique_pk;
alter table consultorio drop constraint consultorio_pk;
alter table agenda drop constraint agenda_pk;
alter table turno drop constraint turno_pk;
alter table reprogramacion drop constraint reprogramacion_pk;
alter table error drop constraint error_pk;
alter table cobertura drop constraint cobertura_pk;
alter table obra_social drop constraint obra_social_pk;
alter table liquidacion_cabecera drop constraint liquidacion_cabecera_pk;
alter table liquidacion_detalle drop constraint liquidacion_detalle_pk;
alter table envio_email drop constraint envio_email_pk;


