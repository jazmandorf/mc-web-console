package mcir

import ( 
	"/tumblebug/common
) 
type TbImageInfo struct {
	Status	string	`json:"status"`
	AssociatedObjectList	[]string	`json:"associatedObjectList"`
	ConnectionName	string	`json:"connectionName"`
	GuestOS	string	`json:"guestOS"`
	Namespace	string	`json:"namespace"`
	Id	string	`json:"id"`
	IsAutoGenerated		`json:"isAutoGenerated"`
	KeyValueList	[]KeyValue	`json:"keyValueList"`
	Name	string	`json:"name"`
	CreationDate	string	`json:"creationDate"`
	CspImageId	string	`json:"cspImageId"`
	CspImageName	string	`json:"cspImageName"`
	Description	string	`json:"description"`
}