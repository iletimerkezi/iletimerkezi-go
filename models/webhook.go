package models

type WebhookReport struct {
    ID       int    `json:"id"`
    PacketID int    `json:"packet_id"`
    Status   string `json:"status"`
    To       string `json:"to"`
    Body     string `json:"body"`
}

func NewWebhookReport(data map[string]interface{}) *WebhookReport {
    report := &WebhookReport{}
    
    if reportData, ok := data["report"].(map[string]interface{}); ok {
        // Parse ID
        if id, ok := reportData["id"].(float64); ok {
            report.ID = int(id)
        }
        
        // Parse PacketID
        if packetID, ok := reportData["packet_id"].(float64); ok {
            report.PacketID = int(packetID)
        }
        
        // Parse string fields
        if status, ok := reportData["status"].(string); ok {
            report.Status = status
        }
        
        if to, ok := reportData["to"].(string); ok {
            report.To = to
        }
        
        if body, ok := reportData["body"].(string); ok {
            report.Body = body
        }
    }

    return report
}

// Status check methods
func (r *WebhookReport) IsDelivered() bool {
    return r.Status == "delivered"
}

func (r *WebhookReport) IsAccepted() bool {
    return r.Status == "accepted"
}

func (r *WebhookReport) IsUndelivered() bool {
    return r.Status == "undelivered"
}