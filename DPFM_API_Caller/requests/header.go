package requests

type Header struct {
	BillOfMaterial       int     `json:"BillOfMaterial"`
	HeaderDeliveryStatus *string `json:"HeaderDeliveryStatus"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
