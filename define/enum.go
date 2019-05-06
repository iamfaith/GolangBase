package define

type EnumType int

const (
	begin      EnumType = iota
	Status     EnumType = iota
	NotProcess EnumType = iota
	Processed  EnumType = iota
	end        EnumType = iota
)

var enums = [...]string{"begin", "Status", "NotProcess", "Processed", "end"}

func (a EnumType) String() string {
	if a <= begin || a >= end {
		return ""
	}
	return enums[a]
}

func ContainsEnum(modelType string) bool {
	for _, t := range enums {
		if t == begin.String() || t == end.String() {
			continue
		}
		if modelType == t {
			return true
		}
	}
	return false
}
