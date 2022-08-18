package factories

import (
	"gohub/app/models/link"

	"github.com/bxcodec/faker/v3"
)

func MakeLinks(times int) []link.Link {

	var objs []link.Link

	for i := 0; i < times; i++ {
		model := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, model)
	}

	return objs
}
