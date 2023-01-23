package dashboard

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2022-08-01/grafanaresource"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type DashboardGrafanaModel struct {
	Name                              string                                            `tfschema:"name"`
	ResourceGroupName                 string                                            `tfschema:"resource_group_name"`
	ApiKeyEnabled                     bool                                              `tfschema:"api_key_enabled"`
	AutoGeneratedDomainNameLabelScope grafanaresource.AutoGeneratedDomainNameLabelScope `tfschema:"auto_generated_domain_name_label_scope"`
	DeterministicOutboundIPEnabled    bool                                              `tfschema:"deterministic_outbound_ip_enabled"`
	Location                          string                                            `tfschema:"location"`
	PublicNetworkAccessEnabled        bool                                              `tfschema:"public_network_access_enabled"`
	Sku                               string                                            `tfschema:"sku"`
	Tags                              map[string]string                                 `tfschema:"tags"`
	ZoneRedundancyEnabled             bool                                              `tfschema:"zone_redundancy_enabled"`
	Endpoint                          string                                            `tfschema:"endpoint"`
	GrafanaVersion                    string                                            `tfschema:"grafana_version"`
	OutboundIPs                       []string                                          `tfschema:"outbound_ip"`
}

type DashboardGrafanaResource struct{}

var _ sdk.ResourceWithUpdate = DashboardGrafanaResource{}

func (r DashboardGrafanaResource) ResourceType() string {
	return "azurerm_dashboard_grafana"
}

func (r DashboardGrafanaResource) ModelObject() interface{} {
	return &DashboardGrafanaModel{}
}

func (r DashboardGrafanaResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return grafanaresource.ValidateGrafanaID
}

func (r DashboardGrafanaResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:     pluginsdk.TypeString,
			Required: true,
			ForceNew: true,
			ValidateFunc: validation.StringMatch(
				regexp.MustCompile(`^[a-zA-Z][-a-zA-Z\d]{0,21}[a-zA-Z\d]$`),
				`The name length must be from 2 to 23 characters. The name can only contain letters, numbers and dashes, and it must begin with a letter and end with a letter or digit.`,
			),
		},

		"resource_group_name": commonschema.ResourceGroupName(),

		"location": commonschema.Location(),

		"api_key_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
		},

		"auto_generated_domain_name_label_scope": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			Default:  string(grafanaresource.AutoGeneratedDomainNameLabelScopeTenantReuse),
			ValidateFunc: validation.StringInSlice([]string{
				string(grafanaresource.AutoGeneratedDomainNameLabelScopeTenantReuse),
			}, false),
		},

		"deterministic_outbound_ip_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  false,
		},

		"identity": commonschema.SystemAssignedIdentityOptionalForceNew(),

		"public_network_access_enabled": {
			Type:     pluginsdk.TypeBool,
			Optional: true,
			Default:  true,
		},

		"sku": {
			Type:     pluginsdk.TypeString,
			Optional: true,
			ForceNew: true,
			Default:  "Standard",
			ValidateFunc: validation.StringInSlice([]string{
				"Standard",
			}, false),
		},

		"tags": commonschema.Tags(),

		"zone_redundancy_enabled": {
			Type:     pluginsdk.TypeBool,
			ForceNew: true,
			Optional: true,
			Default:  false,
		},
	}
}

func (r DashboardGrafanaResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"endpoint": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"grafana_version": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"outbound_ip": {
			Type:     pluginsdk.TypeList,
			Computed: true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},
	}
}

func (r DashboardGrafanaResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			var model DashboardGrafanaModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			client := metadata.Client.Dashboard.GrafanaResourceClient
			subscriptionId := metadata.Client.Account.SubscriptionId
			id := grafanaresource.NewGrafanaID(subscriptionId, model.ResourceGroupName, model.Name)
			existing, err := client.GrafanaGet(ctx, id)
			if err != nil && !response.WasNotFound(existing.HttpResponse) {
				return fmt.Errorf("checking for existing %s: %+v", id, err)
			}

			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			identityValue := expandLegacySystemAndUserAssignedMap(metadata.ResourceData.Get("identity").([]interface{}))

			apiKey := grafanaresource.ApiKeyDisabled
			if model.ApiKeyEnabled {
				apiKey = grafanaresource.ApiKeyEnabled
			}

			deterministicOutboundIP := grafanaresource.DeterministicOutboundIPDisabled
			if model.DeterministicOutboundIPEnabled {
				deterministicOutboundIP = grafanaresource.DeterministicOutboundIPEnabled
			}

			publicNetworkAccess := grafanaresource.PublicNetworkAccessDisabled
			if model.PublicNetworkAccessEnabled {
				publicNetworkAccess = grafanaresource.PublicNetworkAccessEnabled
			}

			zoneRedundancy := grafanaresource.ZoneRedundancyDisabled
			if model.ZoneRedundancyEnabled {
				zoneRedundancy = grafanaresource.ZoneRedundancyEnabled
			}

			properties := &grafanaresource.ManagedGrafana{
				Identity: identityValue,
				Location: utils.String(location.Normalize(model.Location)),
				Properties: &grafanaresource.ManagedGrafanaProperties{
					ApiKey:                            &apiKey,
					AutoGeneratedDomainNameLabelScope: &model.AutoGeneratedDomainNameLabelScope,
					DeterministicOutboundIP:           &deterministicOutboundIP,
					PublicNetworkAccess:               &publicNetworkAccess,
					ZoneRedundancy:                    &zoneRedundancy,
				},
				Sku: &grafanaresource.ResourceSku{
					Name: model.Sku,
				},
				Tags: &model.Tags,
			}

			if err := client.GrafanaCreateThenPoll(ctx, id, *properties); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r DashboardGrafanaResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dashboard.GrafanaResourceClient

			id, err := grafanaresource.ParseGrafanaID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model DashboardGrafanaModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			resp, err := client.GrafanaGet(ctx, *id)
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			properties := resp.Model
			if properties == nil {
				return fmt.Errorf("retrieving %s: properties was nil", id)
			}

			if metadata.ResourceData.HasChange("api_key_enabled") {
				apiKey := grafanaresource.ApiKeyDisabled
				if model.ApiKeyEnabled {
					apiKey = grafanaresource.ApiKeyEnabled
				}

				properties.Properties.ApiKey = &apiKey
			}

			if metadata.ResourceData.HasChange("auto_generated_domain_name_label_scope") {
				properties.Properties.AutoGeneratedDomainNameLabelScope = &model.AutoGeneratedDomainNameLabelScope
			}

			if metadata.ResourceData.HasChange("deterministic_outbound_ip_enabled") {
				deterministicOutboundIP := grafanaresource.DeterministicOutboundIPDisabled
				if model.DeterministicOutboundIPEnabled {
					deterministicOutboundIP = grafanaresource.DeterministicOutboundIPEnabled
				}

				properties.Properties.DeterministicOutboundIP = &deterministicOutboundIP
			}

			if metadata.ResourceData.HasChange("public_network_access_enabled") {
				publicNetworkAccess := grafanaresource.PublicNetworkAccessDisabled
				if model.PublicNetworkAccessEnabled {
					publicNetworkAccess = grafanaresource.PublicNetworkAccessEnabled
				}

				properties.Properties.PublicNetworkAccess = &publicNetworkAccess
			}

			properties.SystemData = nil

			if metadata.ResourceData.HasChange("tags") {
				properties.Tags = &model.Tags
			}

			if err := client.GrafanaCreateThenPoll(ctx, *id, *properties); err != nil {
				return fmt.Errorf("updating %s: %+v", *id, err)
			}

			return nil
		},
	}
}

func (r DashboardGrafanaResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dashboard.GrafanaResourceClient

			id, err := grafanaresource.ParseGrafanaID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.GrafanaGet(ctx, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}

				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			model := resp.Model
			if model == nil {
				return fmt.Errorf("retrieving %s: model was nil", id)
			}

			state := DashboardGrafanaModel{
				Name:              id.WorkspaceName,
				ResourceGroupName: id.ResourceGroupName,
				Location:          location.NormalizeNilable(model.Location),
			}

			identityValue := flattenLegacySystemAndUserAssignedMap(model.Identity)

			if err := metadata.ResourceData.Set("identity", identityValue); err != nil {
				return fmt.Errorf("setting `identity`: %+v", err)
			}

			if properties := model.Properties; properties != nil {
				if properties.ApiKey != nil {
					if *properties.ApiKey == grafanaresource.ApiKeyEnabled {
						state.ApiKeyEnabled = true
					} else {
						state.ApiKeyEnabled = false
					}
				}

				if properties.AutoGeneratedDomainNameLabelScope != nil {
					state.AutoGeneratedDomainNameLabelScope = *properties.AutoGeneratedDomainNameLabelScope
				}

				if properties.DeterministicOutboundIP != nil {
					if *properties.DeterministicOutboundIP == grafanaresource.DeterministicOutboundIPEnabled {
						state.DeterministicOutboundIPEnabled = true
					} else {
						state.DeterministicOutboundIPEnabled = false
					}
				}

				if properties.Endpoint != nil {
					state.Endpoint = *properties.Endpoint
				}

				if properties.GrafanaVersion != nil {
					state.GrafanaVersion = *properties.GrafanaVersion
				}

				if properties.OutboundIPs != nil {
					state.OutboundIPs = *properties.OutboundIPs
				}

				if properties.PublicNetworkAccess != nil {
					if *properties.PublicNetworkAccess == grafanaresource.PublicNetworkAccessEnabled {
						state.PublicNetworkAccessEnabled = true
					} else {
						state.PublicNetworkAccessEnabled = false
					}
				}

				if properties.ZoneRedundancy != nil {
					if *properties.ZoneRedundancy == grafanaresource.ZoneRedundancyEnabled {
						state.ZoneRedundancyEnabled = true
					} else {
						state.ZoneRedundancyEnabled = false
					}
				}
			}

			if model.Sku != nil {
				state.Sku = model.Sku.Name
			}

			if model.Tags != nil {
				state.Tags = *model.Tags
			}

			return metadata.Encode(&state)
		},
	}
}

func (r DashboardGrafanaResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Dashboard.GrafanaResourceClient

			id, err := grafanaresource.ParseGrafanaID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if err := client.GrafanaDeleteThenPoll(ctx, *id); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}

func expandLegacySystemAndUserAssignedMap(input []interface{}) *identity.LegacySystemAndUserAssignedMap {
	identityValue, err := identity.ExpandSystemAssigned(input)
	if err != nil {
		return nil
	}

	return &identity.LegacySystemAndUserAssignedMap{
		Type: identityValue.Type,
	}
}

func flattenLegacySystemAndUserAssignedMap(input *identity.LegacySystemAndUserAssignedMap) *[]interface{} {
	if input == nil {
		return &[]interface{}{}
	}

	identityValue := &identity.SystemAssigned{
		Type:        input.Type,
		PrincipalId: input.PrincipalId,
		TenantId:    input.TenantId,
	}

	output := identity.FlattenSystemAssigned(identityValue)
	return &output
}
