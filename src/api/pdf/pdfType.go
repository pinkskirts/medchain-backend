package pdf

import "time"

type Medicamento struct {
	Nome      string `json:"nome"`
	Posologia string `json:"posologia"`
}

type ReceitaMedica struct {
	Medico       Medico        `json:"medico"`
	Paciente     Paciente      `json:"paciente"`
	DataHora     time.Time     `json:"data_hora"`
	Medicamentos []Medicamento `json:"medicamentos"`
}

type Medico struct {
	Nome          string `json:"medico_nome"`
	CRM           string `json:"medico_crm"`
	Especialidade string `json:"especialidade"`
	Endereco      string `json:"medico_endereco"`
}

type Paciente struct {
	Nome      string `json:"paciente_nome"`
	Documento string `json:"paciente_documento"`
	Endereco  string `json:"paciente_endereco"`
}
type PDF struct {
	PDFbytes []byte    `json:"PDFbytes"`
	Error    typeError `json:"errror"`
}

type typeError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
