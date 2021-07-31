package requests

type RequestType int

const (
	SetPlanName RequestType = iota
)

func (requestType RequestType) String() string {
	return [...]string{"setPlanName"}[requestType]
}
