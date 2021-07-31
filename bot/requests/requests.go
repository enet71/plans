package requests

var requests = map[int]*Request{}

func addRequest(request *Request) {
	requests[request.UserId] = request
}

func getRequest(userId int) *Request {
	return requests[userId]
}

func userHasRequest(userId int) bool {
	return requests[userId] != nil
}
