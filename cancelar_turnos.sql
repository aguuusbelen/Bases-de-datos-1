create or replace function Cancelar_turnos(dni_medique int, desde_fecha timestamp, hasta_fecha timestamp) returns int as $$
declare
begin
	--recorre la tabla turno con el rango de fechas dado
	--update de estado = 'cancelado'
	-- insert en tabla reprogramacion
end;
$$ language plpgsql;
