package model

type AccessInfo struct {
	Id              int64
	EndpointAddress string
	Method          string
	Role            Role
}
