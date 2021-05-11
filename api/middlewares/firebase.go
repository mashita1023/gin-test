package middlewares

import (
  "fmt"
	"context"
  "log"
  "strings"


	"github.com/gin-gonic/gin"
  firebase "firebase.google.com/go"
  //  "firebase.google.com/go/auth"

  "google.golang.org/api/option"
)

func SetupFirebase() *firebase.App{

  opt := option.WithCredentialsFile("secretkey.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    fmt.Println(err)
    log.Fatalf("error initializing app: %v", err)
  }
  return app
}

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    opt := option.WithCredentialsFile("secretkey.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
      fmt.Println(err)
    }

    auth, err := app.Auth(context.Background())
    if err != nil {
      fmt.Printf("err: %v\n", err)
    }

    authHeader := c.Request.Header.Get("Authorization")
    idToken := strings.Replace(authHeader, "Bearer ", "", 1)

    token, err := auth.VerifyIDToken(context.Background(), idToken)

    if err != nil {
      fmt.Printf("error verifying ID token: %v\n", err)
      c.JSON(401, gin.H{
        "message": "error verifying ID token",
      })
      c.Abort()
    }
    log.Printf("Verified ID token: %v\n", token)
    c.Next()
  }
}
