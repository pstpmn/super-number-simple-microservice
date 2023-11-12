package configs

type (
	App struct {
		Name  string
		Url   string
		Stage string
	}
)

func NewApp(name, url, stage string) App {
	return App{
		Name:  name,
		Url:   url,
		Stage: stage,
	}
}
