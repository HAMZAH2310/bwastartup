package auth

import(
	"github.com/golang-jwt/jwt/v4"
)

type Service interface{
		GenerateToken(userID int) (string,error)
}

type jwtService struct{

}

func NewService() *jwtService{
		return &jwtService{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func(s* jwtService) GenerateToken(userID int)(string,error){
		claim := jwt.MapClaims{}
		claim["user_id"] = userID

		token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

		signedToken,err := token.SignedString(SECRET_KEY)
		if err != nil{
			return signedToken,err
		}
		return signedToken,nil
}