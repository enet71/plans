package requests

type RequestType int

const (
	SetPlanName RequestType = iota
	SetPlanStartDate
	SetPlanEndDate
	SetPlanDescription
	FinishPlanCreating
)

var RequestTypes = [...]string{
	"setPlanName",
	"setPlanStartDate",
	"setPlanEndDate",
	"setPlanDescription",
	"finishPlanCreating",
}

func (requestType RequestType) String() string {
	return RequestTypes[requestType]
}

func IsRequest(value string) bool {
	for _, requestType := range RequestTypes {
		if requestType == value {
			return true
		}
	}

	return false
}
