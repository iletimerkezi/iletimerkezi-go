# İleti Merkezi Go SDK

İleti Merkezi SMS API'sini Go programlama dili ile kullanmanızı sağlayan resmi SDK.

## Gereksinimler

- Go 1.18 veya üzeri

## Kurulum

```bash
go get github.com/iletimerkezi/iletimerkezi-go
```

## Kullanım

### Client Oluşturma

```go
import "github.com/iletimerkezi/iletimerkezi-go"

// Temel kullanım
client := iletimerkezi.NewClient("API_KEY", "API_HASH")

// Varsayılan gönderici adı ile
client := iletimerkezi.NewClient("API_KEY", "API_HASH", iletimerkezi.WithDefaultSender("SENDER"))

// Debug modu ile
client := iletimerkezi.NewClient("API_KEY", "API_HASH", iletimerkezi.WithDebug(true))
```

### SMS Gönderme

```go
// Tek numaraya SMS gönderme
resp, err := client.Sms().Send("555123xxxx", "Test mesajı")
if err != nil {
    log.Fatal(err)
}

// Birden fazla numaraya SMS gönderme
numbers := []string{"555123xxxx", "555765xxxx"}
resp, err := client.Sms().Send(numbers, "Toplu test mesajı")

// İleri tarihli SMS gönderme
sendTime := time.Now().Add(24 * time.Hour)
resp, err := client.Sms().SendScheduled("5551234567", "İleri tarihli mesaj", &sendTime)
```

### Rapor Sorgulama

```go
// Özet rapor alma
startDate := time.Now().AddDate(0, 0, -7) // Son 7 gün
endDate := time.Now()
resp, err := client.Summary().Get(&startDate, &endDate, 1)
```

### Gönderici Adları

```go
// Gönderici adlarını listeleme
resp, err := client.Senders().List()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Gönderici adları:", resp.GetSenders())
```

### Bakiye Sorgulama

```go
// Hesap bakiyesi sorgulama
resp, err := client.Account().Balance()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Bakiye: %.2f TL\n", resp.Amount)
fmt.Printf("Kalan SMS: %d\n", resp.Credits)
```

### Webhook İşleme

```go
// HTTP handler içinde webhook işleme
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    report, err := client.Webhook().Handle(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if report.IsDelivered() {
        fmt.Printf("SMS teslim edildi: ID=%s\n", report.GetID())
    }
}
```

### Kara Liste İşlemleri

```go
// Numarayı kara listeye ekleme
resp, err := client.Blacklist().Add("555123xxxx")

// Kara listeden numara silme
resp, err := client.Blacklist().Delete("555123xxxx")

// Kara liste sorgulama
resp, err := client.Blacklist().Check("555123xxxx")
```

### Debug Bilgileri

```go
// Son request/response bilgilerini görüntüleme
client.EnableDebug()
resp, _ := client.Sms().Send("555123xxxx", "Test mesajı")
fmt.Println(client.Debug())
```

## Özellikler

- SMS gönderimi (tekli/çoklu)
- İleri tarihli SMS gönderimi
- Raporlama ve sorgulama
- Gönderici adı yönetimi
- Bakiye sorgulama
- Webhook desteği
- Kara liste yönetimi
- Debug modu
- Özelleştirilebilir HTTP client
- Thread-safe yapı

## Hata Yönetimi

SDK, API yanıtlarını ve hataları Go'nun standart hata yönetimi yaklaşımı ile ele alır:

```go
resp, err := client.Sms().Send("555123xxxx", "Test mesajı")
if err != nil {
    log.Fatal("SMS gönderilemedi:", err)
}

if !resp.Ok() {
    log.Printf("API Hatası: %s (Kod: %d)\n", resp.GetMessage(), resp.GetStatusCode())
    return
}
```

## Thread Safety

SDK thread-safe olarak tasarlanmıştır ve concurrent kullanıma uygundur.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakınız.