package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/yourusername/iletimerkezi-go"
)

func main() {
    // Create client with options
    client := iletimerkezi.NewClient(
        "your-api-key",
        "your-api-hash",
        iletimerkezi.WithDefaultSender("eMarka"),
        iletimerkezi.WithDebug(true),
    )

    // Example 1: Send SMS
    fmt.Println("Sending SMS...")
    smsService := client.SMS()
    smsService.Schedule(time.Now().Add(time.Hour))
    smsService.EnableIysConsent()

    // Example 2: Get Report
    fmt.Println("\nGetting Report...")
    reportService := client.Reports()
    report, err := reportService.Get(12345, 1, 1000)
    if err != nil {
        log.Fatal("Report error:", err)
    }
    fmt.Printf("Order Status: %s\n", report.GetOrderStatus())

    // Example 3: Get Summary
    fmt.Println("\nGetting Summary...")
    summaryService := client.Summary()
    startDate := time.Now().AddDate(0, 0, -1)
    endDate := time.Now()
    
    summary, err := summaryService.List(startDate, endDate, 1)
    if err != nil {
        log.Fatal("Summary error:", err)
    }

    for summary.HasMorePages() {
        fmt.Printf("Orders count: %d\n", len(summary.GetOrders()))
        summary, err = summaryService.Next()
        if err != nil {
            log.Fatal("Next page error:", err)
        }
    }
} 