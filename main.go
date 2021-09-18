package main

import (
	// "context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// sess := session.NewSessionWithOptions()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("af-south-1")},
		Profile: "onion",
	}))
	client := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String("app-onion-v2ray"),
				},
			},
		},
	}

	ec2output, err := client.DescribeInstances(input)

	if err != nil {
		log.Fatal(err)
	}

	ec2list := &ec2output.Reservations[0].Instances
	// fmt.Println(ec2list)
	for _,instance := range *ec2list {
		fmt.Println(*instance.InstanceId)
	}
}
