package enum

type AttrShort string
type RoleAttr string

const (
	StrengthAttr    RoleAttr = "strength"
	CultureAttr     RoleAttr = "culture"
	EnvironmentAttr RoleAttr = "environment"
	CharismaAttr    RoleAttr = "charisma"
	TalentAttr      RoleAttr = "talent"
	IntellectAttr   RoleAttr = "intellect"

	StrengthShort    AttrShort = "str"
	CultureShort     AttrShort = "cul"
	EnvironmentShort AttrShort = "env"
	CharismaShort    AttrShort = "cha"
	TalentShort      AttrShort = "tal"
	IntellectShort   AttrShort = "int"
)

const (
	ObjectiveCreatedExp = 10

	TodoCreatedExp   = 5
	TodoCompletedExp = 10

	MessageCreateExp = 1
)
