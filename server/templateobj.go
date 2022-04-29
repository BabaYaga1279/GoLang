package server

type TempObj struct {
	HostUrl       string
	TemplatesPath string
}

func NewTempObj() TempObj {
	return TempObj{
		"192.168.1.8:6969",
		"./server/templates",
	}
}
