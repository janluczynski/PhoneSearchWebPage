package commons

//package od struktur

type Product struct {
	ProductURL  string   `json:"product_url" bson:"product_url"`
	ProductID   string   `json:"product_id" bson:"product_id"`
	ProductName string   `json:"product_name" bson:"product_name"`
	Brand       string   `json:"brand" bson:"brand"`
	ImageURL    string   `json:"image_url" bson:"image_url"` //TODO: change to []string
	Imagetab    []string `json:"image_tab" bson:"image_tab"`
	SalePrice   float64  `json:"sale_price" bson:"sale_price"`
	Colour      string   `json:"colour" bson:"colour"`
	Description string   `json:"description" bson:"description"`
	Category    string   `json:"category" bson:"category"`
	Material    string   `json:"material" bson:"material"`
}
type ProductStock struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Count     int    `json:"stock" bson:"stock"`
}
