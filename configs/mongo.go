package configs

type (
	Mongo struct {
		Url string
	}
)

func NewMongo(url string) Mongo {
	return Mongo{
		Url: url,
	}
}
