package usecase

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Dalot/goddd/internal/app/domain"
	"github.com/Dalot/goddd/internal/app/domain/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	UserErrWrongPassword     = errors.New("The password does not match.")
	UserErrCouldNotCreateJWT = errors.New("Cannot create the token.")
	UserErrAlreadyExists     = errors.New("An account with such email already exists.")
)

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginArgs struct {
	Email          string
	Password       string
	UserRepository repository.IUser
}
type RegisterArgs struct {
	Email          string
	Password       string
	Username       string
	UserRepository repository.IUser
}

func Users(userRepository repository.IUser) []domain.User {
	return userRepository.Index()
}

func Login(args LoginArgs) (*http.Cookie, error) {
	user, err := args.UserRepository.GetByEmail(args.Email)
	if err != nil {
		return &http.Cookie{}, err
	}

	_, err = user.CompareHashAndPassword([]byte(user.Password), []byte(args.Password))
	if err != nil {

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Println(err.Error())
			return &http.Cookie{}, UserErrWrongPassword
		} else {
			panic(err)
		}
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims
	claims := &Claims{
		Email: args.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Fatal(err)
		return &http.Cookie{}, UserErrCouldNotCreateJWT
	}

	cookie := &http.Cookie{
		Name:    "code_challenge_token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	return cookie, nil
}

func Register(args RegisterArgs) (*domain.User, error) {
	//TODO: Check what happens here.
	user, err := args.UserRepository.GetByEmail(args.Email)
	if user.ID > 0 {
		return &domain.User{}, UserErrAlreadyExists
	}

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return &domain.User{}, err
		} 
	}

	pw, err := user.HashAndSalt([]byte(args.Password))
	if err != nil {
		log.Printf("could no hash the password")
		panic(err)
	}

	user = domain.User{
		Username: args.Username,
		Email:    args.Email,
		Password: pw,
	}

	user, err = args.UserRepository.Create(user)
	if err != nil {
		log.Printf("could not create the user")
		return &domain.User{}, err
	}

	return &user, nil
}
