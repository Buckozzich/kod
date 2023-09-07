package client

import (
	"fmt"
	"log"
	
	
	"ctx"
	"time"
	
)

type SingInClient struct {
	service pb.AuthServiceClient
}

func NewSingInUserClient(conn *grp.clientConnect) *SingUpClient {
service := pd.AutoServiceClient(conn)

return &SingInClient{service: }
}


func (singInUserClient *SingInUserClient) SingInsUser(*pdclSingUpInput) {

ctx, cancel := 	context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := signInUserClient.service.SignInUser(ctx, credentials)

	if err != nil {
		log.Fatalf("SignInUser: %v", err)
	}

	fmt.Println(res)
	
}

