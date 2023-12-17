package commons

type Product struct {
	ProductURL string  `json:"product_url" bson:"product_url"`
	ProductID  string  `json:"product_id" bson:"product_id"`
	Name       string  `json:"name" bson:"name"`
	Brand      string  `json:"brand" bson:"brand"`
	Model      string  `json:"model" bson:"model"`
	ImageURL   string  `json:"image" bson:"image"`
	Price      float64 `json:"price" bson:"price"`
	Display    string  `json:"display" bson:"display"`
	Processor  string  `json:"processor" bson:"processor"`
	RAM        int     `json:"ram" bson:"ram"`
	Storage    int     `json:"storage" bson:"storage"`
	Battery    int     `json:"battery" bson:"battery"`
}
