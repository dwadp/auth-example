# Auth Example
I'm trying to make a very simple authentication application and trying to implement `Clean Architecture` into the application. The goal is to make the application as flexible as possible. Flexible in the sense that the application can be replaced using any database such as `Mongo DB` or `MySQL` or others by simply making the implementation of the database without changing the logical business code and so on. And also I try to make as many abstractions from the library as I use to make it more flexible.

And at the same time I did testing for the library that I just created [https://github.com/dwadp/mantau](https://github.com/dwadp/mantau) 😆.

# The Stack
- **Golang**
- **Mongo DB**
- **Redis** (*for storing user session*)

## Framework & Libraries
- **Gin** (*web framework*) [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Viper** (*for managing configuration*) [https://github.com/spf13/viper](https://github.com/spf13/viper)
- **Bcrypt** (*for password hashing*) [https://github.com/golang/crypto](https://github.com/golang/crypto)
- **JWT** (*a json web token implementation for golang*) [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- **Redis** (*a redis client implementation for golang*) [https://github.com/go-redis/redis](https://github.com/go-redis/redis)
- **Mongo DB Driver** (*a mongo db driver implementation for golang*) [https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
- **Mantau** (*a schema based golang data transformer*) [https://github.com/dwadp/mantau](https://github.com/dwadp/mantau)

