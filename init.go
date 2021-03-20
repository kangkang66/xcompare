package xcompare

var (
	IN          *inOperator
	Equal       *equalOperator
	NotEqual    *notEqualOperator
	Great       *greatOperator
	GreatEqual  *greatEqualOperator
	Litter      *litterOperator
	LitterEqual *litterEqualOperator
)

func init() {
	IN = new(inOperator)
	Equal = new(equalOperator)
	NotEqual = new(notEqualOperator)
	Great = new(greatOperator)
	GreatEqual = new(greatEqualOperator)
	Litter = new(litterOperator)
	LitterEqual = new(litterEqualOperator)
}
