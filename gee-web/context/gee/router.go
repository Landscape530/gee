package gee

type router struct {
	handlers map[string]HandlerFunc
}
