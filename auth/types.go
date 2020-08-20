package auth

type Type string

var (
	Uint    Type = "uint"
	Integer Type = "integer"
	Float   Type = "float"
	Boolean Type = "boolean"
	String  Type = "string"
	JSON    Type = "json"
)
