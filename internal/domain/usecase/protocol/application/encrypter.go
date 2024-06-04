package protocol_application

type Encrypter interface {
	Hash(string) (string, error)
	CheckPasswordHash(string, string) bool
}
