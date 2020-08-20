package hash

type Hash interface {
	Make(plainText string) (string, error)
	Check(plainText string, hashedText string) bool
}
