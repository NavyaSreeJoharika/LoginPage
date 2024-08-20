package main

import (
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-cost-cli",
	Short: "AWS Cost CLI",
	Long:  `A CLI tool to retrieve and display AWS EC2 cost data using AWS SDK for Go.`,
}

var getCostDataCmd = &cobra.Command{
	Use:   "getEC2CostData",
	Short: "Get cost of EC2 services",
	Long:  `Get the cost of EC2 services for a specified period`,
	Run: func(cmd *cobra.Command, args []string) {
		region, _ := cmd.Flags().GetString("region")
		startDate, _ := cmd.Flags().GetString("start-date")
		endDate, _ := cmd.Flags().GetString("end-date")

		totalCost, err := getTotalEC2Cost(region, startDate, endDate)
		if err != nil {
			log.Fatalln("Error getting total EC2 cost:", err)
			return
		}
		log.Printf("Total EC2 Cost: $%f", totalCost)
	},
}

func main() {
	rootCmd.AddCommand(getCostDataCmd)

	getCostDataCmd.Flags().StringP("region", "r", "us-east-1", "AWS region to query")
	getCostDataCmd.Flags().StringP("start-date", "s", "2024-02-01", "Start date for the cost retrieval in YYYY-MM-DD format")
	getCostDataCmd.Flags().StringP("end-date", "e", "2024-03-01", "End date for the cost retrieval in YYYY-MM-DD format")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func getTotalEC2Cost(region, startDate, endDate string) (float64, error) {
	log.Println("Getting EC2 cost data")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return 0, err
	}

	costClient := costexplorer.New(sess)

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(startDate),
			End:   aws.String(endDate),
		},
		Metrics: []*string{
			aws.String("UnblendedCost"),
		},
		Granularity: aws.String("DAILY"),
		Filter: &costexplorer.Expression{
			Dimensions: &costexplorer.DimensionValues{
				Key: aws.String("SERVICE"),
				Values: []*string{
					aws.String("Amazon Elastic Compute Cloud - Compute"),
				},
			},
		},
	}

	result, err := costClient.GetCostAndUsage(input)
	if err != nil {
		return 0, err
	}

	totalCost := 0.0
	for _, resultByTime := range result.ResultsByTime {
		for _, group := range resultByTime.Groups {
			for _, metricValue := range group.Metrics {
				amountStr := *metricValue.Amount
				cost, err := strconv.ParseFloat(amountStr, 64)
				if err != nil {
					return 0, err
				}

				totalCost += cost
			}
		}
	}

	return totalCost, nil
}
