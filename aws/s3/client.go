package client

import (
	"crypto/tls"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

func NewS3Client() (*s3.S3, error) {
	creds := credentials.NewStaticCredentials(
		viper.GetString("vos.access.key"),
		viper.GetString("vos.secret.key"),
		"",
	)
	if _, err := creds.Get(); err != nil {
		return nil, err
	}
	conf := aws.NewConfig().
		WithRegion("US").
		WithCredentials(creds).
		WithEndpoint(viper.GetString("vos.api.url")).
		WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		})
	return s3.New(session.New(), conf), nil
}
