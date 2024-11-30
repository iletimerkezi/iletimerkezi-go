package responses

type BlacklistResponse struct {
    *BaseResponse
    Total      int
    TotalPage  int
    Numbers    []string
    CurrentPage  int
    NextPage     int
    rowCount     int
    HasMorePages bool
}

type BlacklistCrudResponse struct {
    *BaseResponse
}

func NewBlacklistCrudResponse(resp *Response) *BlacklistCrudResponse {
    base := NewBaseResponse(resp)
    r := &BlacklistCrudResponse{
        BaseResponse: base,
    }
    return r
}

func NewBlacklistResponse(resp *Response, page int, rowCount int) *BlacklistResponse {
    base := NewBaseResponse(resp)
    r := &BlacklistResponse{
        BaseResponse: base,
        CurrentPage:  page,
        rowCount:     rowCount,
    }
    r.parseBlacklistData()
    r.GetNextPage()
    r.checkMorePages()
    r.GetTotalPageCount()
    return r
}

func (r *BlacklistResponse) parseBlacklistData() {
    if blacklist, ok := r.Data["blacklist"].(map[string]interface{}); ok {
        if numbers, ok := blacklist["number"].([]interface{}); ok {
            r.Numbers = make([]string, 0, len(numbers))
            for _, num := range numbers {
                if str, ok := num.(string); ok {
                    r.Numbers = append(r.Numbers, str)
                }
            }
        }
        
        if count, ok := blacklist["count"].(int); ok {
            r.Total = int(count)
        }
    }
}

// Sonraki sayfa numarasını döndürür
func (r *BlacklistResponse) GetNextPage() {
    r.NextPage = r.CurrentPage
    if r.HasMorePages {
        r.NextPage = r.CurrentPage + 1
    }
}

// Daha fazla sayfa olup olmadığını kontrol eder
func (r *BlacklistResponse) checkMorePages() {
    currentCount := r.CurrentPage * r.rowCount
    r.HasMorePages = currentCount < r.Total 
}

// Toplam sayfa sayısını döndürür
func (r *BlacklistResponse) GetTotalPageCount() {
    
    if r.Total == 0 {
        r.TotalPage = 0
    }
    
    // Tam bölünme durumunda yukarı yuvarlama yapmamak için
    if r.Total%r.rowCount == 0 {
        r.TotalPage = r.Total / r.rowCount
    } else {
        r.TotalPage = (r.Total + r.rowCount - 1) / r.rowCount
    }
} 