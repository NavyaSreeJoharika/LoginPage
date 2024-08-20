package cmd

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

var CostCmd = &cobra.Command{
	Use:   "cost",
	Short: "Get the cost breakdown by instance type",
	Run: func(cmd *cobra.Command, args []string) {
		region, _ := cmd.Flags().GetString("region")
		getCostByInstanceType(region)
	},
}

func init() {
	CostCmd.Flags().StringP("region", "r", "us-east-1", "AWS region to query")
}

func getCostByInstanceType(region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	svc := ec2.New(sess)

	result, err := svc.DescribeInstances(nil)
	if err != nil {
		log.Fatalf("Failed to describe instances: %v", err)
	}

	instanceCostMap := map[string]float64{
		"t2.micro":   0.0116,
		"m5.large":   0.096,
		"c5.xlarge":  0.17,
		"r5.2xlarge": 0.504,
		"t3.medium":  0.0416,
		"m4.xlarge":  0.2,
	}

	costBreakdown := map[string]float64{}

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceType := *instance.InstanceType
			costBreakdown[instanceType] += instanceCostMap[instanceType]
		}
	}

	fmt.Println("Cost by Instance Type:")
	for instanceType, cost := range costBreakdown {
		fmt.Printf("%s: $%.2f\n", instanceType, cost)
	}
}
