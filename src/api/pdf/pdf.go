package pdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func Main(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("PANIC: Erro Interno")
			http.Error(response, "Erro interno", http.StatusInternalServerError)
		}
	}()
	startHour := time.Now().UTC()

	if request.Method == "GET" {
		// Decodifica o corpo JSON da requisição para a struct ReceitaMedica
		var receita ReceitaMedica
		err := json.NewDecoder(request.Body).Decode(&receita)
		if err != nil {
			log.Panicln("Erro ao ler request: ", err)
			return
		}

		resp := PostPDF(receita)

		// Definir cabeçalhos de resposta para retornar o arquivo PDF
		response.Header().Set("Content-Type", "application/pdf")
		response.Header().Set("Content-Disposition", "attachment; filename=receita_medica.pdf")

		if resp.Error.Status == 0 {
			_, err = response.Write(resp.PDFbytes)
			if err != nil {
				log.Panicln("Erro ao retornar reqeust")
			}
			log.Println(request.Method, "- Terminated Request time", time.Since(startHour), " - ", http.StatusOK)
		} else {
			erros, err := json.Marshal(resp.Error)
			if err != nil {
				log.Panicln("Erro ao converter em JSON: ", err)
			}
			http.Error(response, string(erros), resp.Error.Status)
			log.Println(request.Method, "- Terminated Request time", time.Since(startHour), " - ", resp.Error.Status)
		}

		return
	} else {
		http.Error(response, "Metodo inválido", http.StatusMethodNotAllowed)
		return
	}

}

// Função para converter de UTF-8 para ISO-8859-1
func TransformUTF8ToISO8859_1(input string) string {
	encoder := charmap.ISO8859_1.NewEncoder()
	isoText, _, err := transform.String(encoder, input)
	if err != nil {
		return input // Retorna o texto original se houver erro
	}
	return isoText
}

// GerarReceitaPDF gera um PDF a partir do JSON enviado pelo cliente e retorna como resposta
func PostPDF(receita ReceitaMedica) PDF {
	// Criar uma nova instância do PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFontLocation("") // Permitir o uso de fontes embutidas

	pdf.AddPage()

	// Configurar Arial com suporte a caracteres especiais
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, TransformUTF8ToISO8859_1("RECEITUÁRIO"), "", 0, "C", false, 0, "")
	pdf.Ln(6)
	// Configurar uma fonte menor para o CRM e Especialidade
	pdf.SetFont("Arial", "", 9) // Fonte regular e menor
	pdf.CellFormat(0, 10, fmt.Sprintf("%s - CRM: %s - %s", receita.Medico.Nome, receita.Medico.CRM, receita.Medico.Especialidade), "", 0, "C", false, 0, "")
	pdf.Ln(10)

	// Adicionar espaço após as informações
	pdf.Line(10, 25, 200, 25) // Desenha uma linha horizontal da posição x=10, y=20 até x=200

	// Informações do Paciente
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, TransformUTF8ToISO8859_1("Paciente"))
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, TransformUTF8ToISO8859_1(fmt.Sprintf("Nome: %s", receita.Paciente.Nome)))
	pdf.Ln(6)
	pdf.Cell(0, 8, TransformUTF8ToISO8859_1(fmt.Sprintf("CPF: %s", receita.Paciente.Documento)))
	pdf.Ln(6)
	pdf.Cell(0, 8, TransformUTF8ToISO8859_1(fmt.Sprintf("Endereço: %s", receita.Paciente.Endereco)))
	pdf.Ln(10)

	// Medicamentos e Posologia
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, TransformUTF8ToISO8859_1("Prescrição"))
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	for _, medicamento := range receita.Medicamentos {
		pdf.Cell(0, 8, TransformUTF8ToISO8859_1(fmt.Sprintf("Medicamento: %s", medicamento.Nome)))
		pdf.Ln(6)
		pdf.Cell(0, 8, TransformUTF8ToISO8859_1(medicamento.Posologia))
		pdf.Ln(10)
	}

	// Formatar Data e Hora no formato brasileiro (DD/MM/YYYY HH:mm)
	dataFormatada := receita.DataHora.Format("02/01/2006 15:04:05")

	// Rodapé
	pdf.Line(10, 245, 200, 245) // Desenha uma linha horizontal da posição x=10, y=20 até x=200
	pdf.SetY(248)
	pdf.SetFont("Arial", "I", 8)

	pdf.CellFormat(0, 10, TransformUTF8ToISO8859_1("***Esta receita foi assinada digitalmente e não tem validade legal.***"), "", 0, "C", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(0, 10, TransformUTF8ToISO8859_1("Projeto de Desenvolvimento de Software - 2024 "), "", 0, "C", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(0, 10, TransformUTF8ToISO8859_1(fmt.Sprintf("\nEndereço: %s", receita.Medico.Endereco)), "", 0, "C", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(0, 10, TransformUTF8ToISO8859_1(dataFormatada), "", 0, "C", false, 0, "")

	// Usar um buffer para capturar os bytes do PDF
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		log.Panicln("Erro ao passar o pdf para bytes: ", err)
	}

	return PDF{PDFbytes: buf.Bytes()}
}
