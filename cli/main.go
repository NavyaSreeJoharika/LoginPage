/*package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an EC2 client
	svc := ec2.NewFromConfig(cfg)

	// Describe EC2 instance types
	describeInstanceTypes(svc)
}

func describeInstanceTypes(svc *ec2.Client) {
	input := &ec2.DescribeInstanceTypesInput{}

	// Paginate through the results
	paginator := ec2.NewDescribeInstanceTypesPaginator(svc, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page, %v", err)
		}

		for _, instanceType := range page.InstanceTypes {
			fmt.Printf("Instance Type: %s\n", aws.ToString(instanceType.InstanceType))
			fmt.Printf("vCPUs: %d\n", instanceType.VCpuInfo.DefaultVCpus)
			fmt.Printf("Memory: %d MiB\n\n", instanceType.MemoryInfo.SizeInMiB)
		}
	}
}
*/

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an EC2 client
	svc := ec2.NewFromConfig(cfg)

	// Describe EC2 instance types
	describeInstanceTypes(svc)
}

func describeInstanceTypes(svc *ec2.Client) {
	input := &ec2.DescribeInstanceTypesInput{}

	// Paginate through the results
	paginator := ec2.NewDescribeInstanceTypesPaginator(svc, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page, %v", err)
		}

		for _, instanceType := range page.InstanceTypes {
			fmt.Printf("Instance Type: %s\n", instanceType.InstanceType)
			fmt.Printf("vCPUs: %d\n", instanceType.VCpuInfo.DefaultVCpus)
			fmt.Printf("Memory: %d MiB\n\n", instanceType.MemoryInfo.SizeInMiB)
		}
	}
}
