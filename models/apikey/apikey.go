package apikey

import "github.com/dylantic/SSLurper/helpers"

type apiKey struct {
	Id          int
	Owner       string
	Application string
	Key         string
}

func Create(owner string, app string) apiKey {
	id := 1
	key := helpers.RandomBase64String(20)
	n := apiKey{
		Id:          id,
		Owner:       owner,
		Application: app,
		Key:         key,
	}
	return n
}
