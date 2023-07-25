package dpfm_api_output_formatter

import (
	"data-platform-api-organization-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToOrganization(rows *sql.Rows) (*[]Organization, error) {
	defer rows.Close()
	organization := make([]Organization, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Organization{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Organization,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &organization, nil
		}

		data := pm
		organization = append(organization, Organization{
			BusinessPartner:		data.BusinessPartner,
			Organization:			data.Organization,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &organization, nil
}
