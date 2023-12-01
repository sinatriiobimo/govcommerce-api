package product

type (
	Product struct {
		SKU         string  `json:"sku"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		ImgURL      string  `json:"imgURL"`
		Weight      float64 `json:"weight"`
		Price       float64 `json:"price"`
	}

	UpdateProductRequest struct {
		SKU         string `json:"sku"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
	}

	SearchProductData struct {
		SKU         string  `json:"sku"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		ImgURL      string  `json:"imgURL"`
		Weight      float64 `json:"weight"`
		Price       float64 `json:"price"`
		AvgRating   float64 `json:"avgRating"`
	}

	ParamSearch struct {
		SKU       string
		Title     string
		Category  string
		SortBy    string
		Direction string
		PageSize  int
		PageIndex int
	}

	SearchProductResponse struct {
		Data      []SearchProductData `json:"data"`
		TotalPage int                 `json:"total_page"`
		TotalData int                 `json:"total_data"`
	}
)
