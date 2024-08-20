package cmd

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

var costPanelCmd = &cobra.Command{
	Use:   "costpanel",
	Short: "Generate cost panel for EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize a session in the us-east-1 region
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
		}))

		// Create EC2 service client
		ec2Svc := ec2.New(sess)

		// Create Cost Explorer service client
		ceSvc := costexplorer.New(sess)

		// Call the necessary functions to get data and generate panel
		generateCostPanel(ec2Svc, ceSvc)
	},
}

func init() {
	rootCmd.AddCommand(costPanelCmd)
}

func generateCostPanel(ec2Svc *ec2.EC2, ceSvc *costexplorer.CostExplorer) {
	// Describe EC2 instances
	describeInstancesInput := &ec2.DescribeInstancesInput{}
	result, err := ec2Svc.DescribeInstances(describeInstancesInput)
	if err != nil {
		log.Fatalf("Failed to describe instances: %v", err)
	}

	// Iterate over instances and get their costs
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s, Type: %s\n", *instance.InstanceId, *instance.InstanceType)
			// Fetch the cost details using Cost Explorer
			// (You can add a detailed implementation here)
		}
	}

	// Example of a Cost Explorer API call (customize as needed)
	costInput := &costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("MONTHLY"),
		Metrics:     aws.StringSlice([]string{"AmortizedCost"}),
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2023-07-01"),
			End:   aws.String("2023-07-31"),
		},
	}

	costOutput, err := ceSvc.GetCostAndUsage(costInput)
	if err != nil {
		log.Fatalf("Failed to get cost and usage: %v", err)
	}

	for _, result := range costOutput.ResultsByTime {
		for _, group := range result.Groups {
			fmt.Printf("Cost: %s\n", *group.Metrics["AmortizedCost"].Amount)
		}
	}
}
