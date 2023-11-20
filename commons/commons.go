package commons

type Product struct {
	ProductURL string `json:"product_url" bson:"product_url"`
	ProductID  string `json:"product_id" bson:"product_id"`
	Brand      string `json:"brand" bson:"brand"`
	Model      string `json:"model" bson:"model"`
	ImageURL   string `json:"image_url" bson:"image_url"`   // w razie potrzeby zmienić na []string
	Price      string `json:"sale_price" bson:"sale_price"` // w razie potrzeby zmienić na int żeby było łatwiej porównywać
	Display    string `json:"display" bson:"display"`
	Processor  string `json:"processor" bson:"processor"`
	RAM        string `json:"ram" bson:"ram"`
	Storage    string `json:"storage" bson:"storage"`
	Battery    string `json:"battery" bson:"battery"`
}
