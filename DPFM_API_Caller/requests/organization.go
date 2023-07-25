package requests

type Organization struct {
	BusinessPartner			int		`json:"BusinessPartner"`
	Organization			string	`json:"Organization"`
	OrganizationName		string	`json:"OrganizationName"`
	ValidityStartDate		string	`json:"ValidityStartDate"`
	ValidityEndDate			string	`json:"ValidityEndDate"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
