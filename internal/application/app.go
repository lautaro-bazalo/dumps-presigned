package application

import (
	"context"
	"dumps-presigned/internal/presigner"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

type Application struct {
	Presigner *presigner.Presigner
}

func NewApplication() Application {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(s3Client)

	presigner := presigner.NewPresigner(presignClient)

	return Application{
		Presigner: presigner,
	}
}
