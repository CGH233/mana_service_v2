package apartment

type ApartmentInfoItem struct {
	Apartment string   `json:"apartment"`
	Phone     []string `json:"phone"`
	Place     string   `json:"place"`
}
