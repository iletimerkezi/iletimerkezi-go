package models

type WebhookReport struct {
    id       string
    packetID string
    status   string
    to       string
    body     string
}

func NewWebhookReport(data map[string]interface{}) *WebhookReport {
    report := &WebhookReport{}
    if r, ok := data["report"].(map[string]interface{}); ok {
        report.id = getString(r, "id")
        report.packetID = getString(r, "packet_id")
        report.status = getString(r, "status")
        report.to = getString(r, "to")
        report.body = getString(r, "body")
    }
    return report
}

func getString(m map[string]interface{}, key string) string {
    if val, ok := m[key].(string); ok {
        return val
    }
    return ""
}

func (r *WebhookReport) GetID() string {
    return r.id
}

func (r *WebhookReport) GetPacketID() string {
    return r.packetID
}

func (r *WebhookReport) GetStatus() string {
    return r.status
}

func (r *WebhookReport) GetTo() string {
    return r.to
}

func (r *WebhookReport) GetBody() string {
    return r.body
}

func (r *WebhookReport) IsDelivered() bool {
    return r.status == "delivered"
}

func (r *WebhookReport) IsAccepted() bool {
    return r.status == "accepted"
}

func (r *WebhookReport) IsUndelivered() bool {
    return r.status == "undelivered"
} 