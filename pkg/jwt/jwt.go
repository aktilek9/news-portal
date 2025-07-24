package jwt

type JWTService interface {
	GenerateToken(userID int, role string) (string, error)
	ParseToken(token string) (int, string, error)
}

type jwtService struct {
	secretKey []byte
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{secretKey: []byte(secretKey)}
}

func (j *jwtService) GenerateToken(userID int, role string) (string, error) {
	return "", nil
}

func (j *jwtService) ParseToken(token string) (int, string, error) {
	return 0, "", nil
}