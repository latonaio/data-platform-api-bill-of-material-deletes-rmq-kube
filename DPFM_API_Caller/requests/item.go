package requests

type Item struct {
	BillOfMaterial      int   `json:"BillOfMaterial"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
