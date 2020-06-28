package sqlxx

type DBType uint8

const (
	DB_TYPE_POSTGRES DBType = iota // postgres
)

func (dt DBType) String() string {
	switch dt {
	case DB_TYPE_POSTGRES:
		return "postgres"
	default:
		return ""
	}
}
