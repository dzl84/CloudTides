package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eci"
	"os"
)

func main() {
	region := os.Args[1]
	accessKeyId := os.Args[2]
	accessKeySecret := os.Args[3]
	ContainerId := os.Args[4]

	client, err := eci.NewClientWithAccessKey(region, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}

	request := eci.CreateRestartContainerGroupRequest()
	request.ContainerGroupId = ContainerId
	request.RegionId = region

	_, err = client.RestartContainerGroup(request)
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}
}

