package common

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// func GetBucket() (infos []Info, err error) {

// }

func Test_GetBucket(t *testing.T) {
	c, e := oss.New("http://oss-cn-hangzhou.aliyuncs.com", config.GVC.Accounts[0].SecretID, config.GVC.Accounts[0].SecretKEY)
	if e != nil {
		log.Fatal(e)
	}

	marker := ""
	for {
		lsRes, e := c.ListBuckets(oss.Marker(marker))
		if e != nil {
			t.Fatal(e)
		}
		for _, bucket := range lsRes.Buckets {
			fmt.Println("bucket: ", bucket.Name)
			GetBucketInfomation(c, bucket.Name)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
}

func GetBucketInfomation(client *oss.Client, bucketName string) {
	// 获取存储空间的信息，包括地域（Region或Location）、创建日期（CreationDate）、访问权限（ACL）、拥有者（Owner）、存储类型（StorageClass）、容灾类型（RedundancyType）等。
	res, err := client.GetBucketInfo(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("BucketInfo.Location: ", res.BucketInfo.Location)
	fmt.Println("BucketInfo.CreationDate: ", res.BucketInfo.CreationDate)
	fmt.Println("BucketInfo.ACL: ", res.BucketInfo.ACL)
	fmt.Println("BucketInfo.Owner: ", res.BucketInfo.Owner)
	fmt.Println("BucketInfo.StorageClass: ", res.BucketInfo.StorageClass)
	fmt.Println("BucketInfo.RedundancyType: ", res.BucketInfo.RedundancyType)
	fmt.Println("BucketInfo.ExtranetEndpoint: ", res.BucketInfo.ExtranetEndpoint)
	fmt.Println("BucketInfo.IntranetEndpoint: ", res.BucketInfo.IntranetEndpoint)
}
