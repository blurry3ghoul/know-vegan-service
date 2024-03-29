package service

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/blurry3ghoul/know-vegan-service/src/api/config"
	"github.com/blurry3ghoul/know-vegan-service/src/api/domain"
	"github.com/blurry3ghoul/know-vegan-service/src/api/persistance/repository"
)

type IProductService interface {
	CreateProduct(product *domain.Product, photo *domain.Photo) error
}

type ProductService struct {
	productRepository repository.IProductRepository
	s3Client          *s3.S3
}

func NewProductService(productRepository repository.IProductRepository, s3Client *s3.S3) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		s3Client:          s3Client,
	}
}

func (ps *ProductService) CreateProduct(product *domain.Product, photo *domain.Photo) error {

	data, err := base64.StdEncoding.DecodeString(photo.Base64)
	if err != nil {
		return err
	}
	product.Photo = ps.saveImageToS3(data, photo)
	if product.Photo == "" {
		return errors.New("there's was an error trying to save the image to S3")
	}
	return ps.productRepository.CreateProduct(product, photo)
}

func (ps *ProductService) saveImageToS3(data []byte, photo *domain.Photo) string {
	svc := ps.s3Client
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(config.PHOTOBUCKETNAME),
		Key:                aws.String(photo.Name),
		Body:               bytes.NewReader(data),
		ACL:                aws.String("public-read"),
		ContentDisposition: aws.String("inline"),
		ContentType:        aws.String(photo.Type),
	})
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	// Generate pre-signed URL for photo in S3
	url := retrieveImageAndGenerateURL(svc, photo)

	return url
}

func retrieveImageAndGenerateURL(svc *s3.S3, photo *domain.Photo) string {
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.PHOTOBUCKETNAME),
		Key:    aws.String(photo.Name),
	})

	url, err := req.Presign(24 * time.Hour)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return url
}
