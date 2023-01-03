package connectivity

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

func NewClient(accessKeyId *string, accessKeySecret *string, region *string) (client *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		//Region
		RegionId: region,
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs20140526.cn-shanghai.aliyuncs.com")
	client = &ecs20140526.Client{}
	client, _err = ecs20140526.NewClient(config)
	return client, _err
}

type AliCloudClient struct {
	Region       string
	AccessKey    string
	AccessSecret string
	ecsConn      *ecs20140526.Client
}

// EcsClient 客户端
func (c *AliCloudClient) EcsClient() (*ecs20140526.Client, error) {
	client, err := ecs20140526.NewClient(&(c.AccessKey), &(c.AccessSecret), &(c.Region))
	if err != nil {
		return nil, err
	}
	return client, nil
}
