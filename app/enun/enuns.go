package enun

type TipoSangue struct {
	sangue string
}

type DiaSemana struct {
	diaSemana string
}

type Sexo struct {
	sexo string
}

type TipoEvento struct {
	tipoEvento string
}

func (ts TipoSangue) String() string {
	return ts.sangue
}

func (dia DiaSemana) String() string {
	return dia.diaSemana
}
func (sexo Sexo) String() string {
	return sexo.sexo
}

func (te TipoEvento) String() string {
	return te.tipoEvento
}

var (
	A_POSITIVO  = TipoSangue{"A+"}
	B_POSITIVO  = TipoSangue{"B+"}
	AB_POSTIVO  = TipoSangue{"AB+"}
	O_POSITIVO  = TipoSangue{"O+"}
	A_NEGATIVO  = TipoSangue{"A-"}
	B_NEGATIVO  = TipoSangue{"B-"}
	AB_NEGATIVO = TipoSangue{"AB-"}
	O_NEGATIVO  = TipoSangue{"O-"}
)

var (
	SEGUNDA = DiaSemana{"Segunda"}
	TERCA   = DiaSemana{"Terca"}
	QUARTA  = DiaSemana{"Quarta"}
	QUINTA  = DiaSemana{"Quinta"}
	SEXTA   = DiaSemana{"Sexta"}
	SABADO  = DiaSemana{"Sabado"}
	DOMINGO = DiaSemana{"Domingo"}
)

var (
	MASCULINO = Sexo{"Masculino"}
	FEMININO  = Sexo{"Feminino"}
)

var (
	PEDAL                = TipoEvento{"Pedal"}
	PASSEIO              = TipoEvento{"Passeio"}
	VIAGEM               = TipoEvento{"Viagem"}
	VIAGEM_INTERNACIONAL = TipoEvento{"Viagem Internacional"}
)
