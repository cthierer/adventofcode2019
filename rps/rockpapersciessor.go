package rps

type Action interface {
	Score() int
	Compare(interface{}) int
}
