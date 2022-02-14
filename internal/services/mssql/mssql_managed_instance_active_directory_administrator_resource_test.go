package mssql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mssql/parse"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type MsSqlManagedInstanceActiveDirectoryAdministratorResource struct{}

func TestAccMsSqlManagedInstanceActiveDirectoryAdministrator_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_mssql_managed_instance_active_directory_administrator", "test")
	r := MsSqlManagedInstanceActiveDirectoryAdministratorResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.template(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("administrator_login_password"),
		{
			Config: r.basic(data, false),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("administrator_login_password"),
		{
			Config: r.basic(data, true),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("administrator_login_password"),
	})
}

func (r MsSqlManagedInstanceActiveDirectoryAdministratorResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.ManagedInstanceAzureActiveDirectoryAdministratorID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.MSSQL.ManagedInstanceAdministratorsClient.Get(ctx, id.ResourceGroup, id.ManagedInstanceName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	return utils.Bool(true), nil
}

func (r MsSqlManagedInstanceActiveDirectoryAdministratorResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_client_config" "test" {}

data "azuread_service_principal" "test" {
  object_id = data.azuread_client_config.test.object_id
}

resource "azuread_directory_role" "reader" {
  display_name = "Directory Readers"
}

resource "azuread_directory_role_member" "test" {
  role_object_id   = azuread_directory_role.reader.object_id
  member_object_id = azurerm_mssql_managed_instance.test.identity.0.principal_id
}

%[1]s
`, MsSqlManagedInstanceResource{}.identity(data))
}

func (r MsSqlManagedInstanceActiveDirectoryAdministratorResource) basic(data acceptance.TestData, aadOnly bool) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_mssql_managed_instance_active_directory_administrator" "test" {
  managed_instance_id = azurerm_mssql_managed_instance.test.id
  login               = data.azuread_service_principal.test.display_name
  object_id           = data.azuread_service_principal.test.object_id
  tenant_id           = data.azuread_client_config.test.tenant_id

  azuread_authentication_only = %[2]t

  depends_on = [azuread_directory_role_member.test]
}
`, r.template(data), aadOnly)
}
