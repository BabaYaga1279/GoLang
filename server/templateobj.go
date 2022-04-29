package server

type TempObj struct {
	HostUrl       string
	TemplatesPath string
	Username      string
	Password      string
}

type TempErrorObj struct {
	ErrorLogin string
}

const HostServerAddress = "192.168.1.8:6969"
const TemplateFilesPath = "./server/templates"

func NewTempObj() TempObj {
	return TempObj{
		HostServerAddress,
		TemplateFilesPath,
		"",
		"",
	}
}

func NewTempUserObj(username string, password string) TempObj {
	tmp := NewTempObj()
	tmp.Username = username
	tmp.Password = password
	return tmp
}

func NewTempErrorObj(err string) TempErrorObj {
	const ps = "<p	id=\"warning\" style=\"font-style: italic;font-size: 15px;color:Tomato;\">"
	const pe = "</p>"

	return TempErrorObj{
		err,
	}
}
