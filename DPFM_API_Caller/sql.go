package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-organization-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-organization-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var organization *[]dpfm_api_output_formatter.Organization
	for _, fn := range accepter {
		switch fn {
		case "Organization":
			func() {
				organization = c.Organization(mtx, input, output, errs, log)
			}()
		case "Organizations":
			func() {
				organization = c.Organizations(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Organization:     organization,
	}

	return data
}

func (c *DPFMAPICaller) Organization(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Organization {
	where = fmt.Sprintf("WHERE BusinessPartner = %d", input.Organization.BusinessPartner)
	where := fmt.Sprintf("%s\nAND Organization = '%s'", where, *input.Organization.Organization)

	if input.Organization.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Organization.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_organization_organization_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, BusinessPartner DESC Organization DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToOrganization(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Organizations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Organization {

	if input.Organization.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.Organization.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_organization_organization_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, Organization DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToOrganization(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
