package requests

var requests = map[int]*Request{}

func AddRequest(request *Request) {
	requests[request.UserId] = request
}

func GetRequest(userId int) *Request {
	return requests[userId]
}

func UserHasRequest(userId int) bool {
	return requests[userId] != nil
}
