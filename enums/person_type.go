package enums

type PersonType string

const (
	PessoaFisica   PersonType = "Pessoa Física"
	PessoaJuridica PersonType = "Pessoa Jurídica"
)

func (pt PersonType) String() string {
	return string(pt)
}
