package middlewares

import (
  "fmt"
	"context"
  "os"

  firebase "firebase.google.com/go"
//	"firebase.google.com/go/auth"

//	"google.golang.org/api/option"
)

func SetupFirebase() {

  ctx := context.Background()

  opt := os.Getenv("FIREBASE_PRIVATE_KEY")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
  	 return nil, fmt.Errorf("error initializing app: %v", err)
  }

  client, err := app.Auth(ctx)
  if err != nil {
	  log.Fatalf("errorgetting Auth client: %v\n", err)
  }
}
