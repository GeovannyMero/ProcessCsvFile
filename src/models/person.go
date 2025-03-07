package models

type Person struct {
	ID                        string `bson:"_id,omitempty"`
	Ruc                       string `bson:"ruc"`
	RazonSocial               string `bson:"razon_social"`
	CodigoJuridiccion         string `bson:"codigo_juridiccion"`
	EstadoContribuyente       string `bson:"estado_contribuyente"`
	ClaseContribuyente        string `bson:"clase_contribuyente"`
	FechaInicioActividades    string `bson:"fecha_inicio_actividades"`
	FechaActualización        string `bson:"fecha_actualizacion"`
	FechaSuspencionDefinitiva string `bson:"fecha_suspencion_definitiva"`
	FechaReinicioActividades  string `bson:"fecha_reinicio_actividades"`
	Obligado                  string `bson:"obligado"`
	TipoContribuyente         string `bson:"tipo_contribuyente"`
	NumeroEstablecimiento     int    `bson:"numero_establecimiento"`
	NombreFantasiaComercial   string `bson:"nombre_fantasia_comercial"`
	EstadoEstablecimiento     string `bson:"estado_establecimiento"`
	DescripcionProvinciaEst   string `bson:"descripcion_provincia_est"`
	DescripcionCantonEst      string `bson:"descripcion_canton_est"`
	DescripcionParroquiaEst   string `bson:"descripcion_parroquia_est"`
	CodigoCIIU                string `bson:"codigo_ciiu"`
	ActividadEconomica        string `bson:"actividad_economica"`
	AgenteRetencion           string `bson:"agente_retencion"`
	Especial                  string `bson:"especial"`
}
