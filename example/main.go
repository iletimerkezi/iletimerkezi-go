package main

import (
    "fmt"
    "log"
    // "time"
    "github.com/iletimerkezi/iletimerkezi-go"
)

func main() {
    // Client oluşturma
    client := iletimerkezi.NewClient(
        "API_KEY",
        "API_HASH",
    )

    client.SetDefaultSender("SENDER")

    // SMS Gönderme örneği
    smsResp, err := client.Sms().Send("5551234567", "Test mesajı", "")
    if err != nil {
        log.Fatal("SMS gönderme hatası:", err)
    }

    fmt.Println(smsResp)
    fmt.Println("\nDebug Bilgileri:")
    fmt.Println(client.Debug())

    /*
    if !smsResp.Ok() {
        log.Printf("API Hatası: %s (Kod: %d)\n", smsResp.GetMessage(), smsResp.GetStatusCode())
        return
    }

    fmt.Printf("SMS başarıyla gönderildi. Order ID: %s\n", smsResp.OrderID)

    // Rapor sorgulama örneği
    startDate := time.Now().AddDate(0, 0, -7) // Son 7 gün
    endDate := time.Now()
    
    summaryResp, err := client.Summary().Get(&startDate, &endDate, 1)
    if err != nil {
        log.Fatal("Rapor sorgulama hatası:", err)
    }

    if !summaryResp.Ok() {
        log.Printf("API Hatası: %s (Kod: %d)\n", summaryResp.GetMessage(), summaryResp.GetStatusCode())
        return
    }

    fmt.Printf("Toplam SMS sayısı: %d\n", summaryResp.GetCount())

    // Bakiye sorgulama örneği
    accountResp, err := client.Account().Balance()
    if err != nil {
        log.Fatal("Bakiye sorgulama hatası:", err)
    }

    if !accountResp.Ok() {
        log.Printf("API Hatası: %s (Kod: %d)\n", accountResp.GetMessage(), accountResp.GetStatusCode())
        return
    }

    fmt.Printf("Bakiye: %.2f TL\n", accountResp.Amount)
    fmt.Printf("Kalan SMS: %d\n", accountResp.Credits)

    // Debug bilgilerini görüntüleme
    fmt.Println("\nDebug Bilgileri:")
    fmt.Println(client.Debug())
    */
}