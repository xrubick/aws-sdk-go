package getinfo

import (
	// "fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetRegion(client *ec2.EC2) (regionlist []*string , err error) {
	input := &ec2.DescribeRegionsInput{}
	var i *ec2.Region

	result,err := client.DescribeRegions(input)
	if err != nil {
		log.Fatal(err)
	}
	for _,i = range(result.Regions) {
		regionlist = append(regionlist, i.RegionName)
	}
	
	return regionlist,err
}