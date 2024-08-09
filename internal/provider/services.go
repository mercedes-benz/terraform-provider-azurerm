// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/aadb2c"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/advisor"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/analysisservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/apimanagement"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/appconfiguration"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/applicationinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/appservice"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/arckubernetes"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/arcresourcebridge"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/attestation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/authorization"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/automanage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/automation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/azurestackhci"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/batch"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/billing"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/blueprints"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/bot"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cdn"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cognitive"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/communication"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/compute"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/confidentialledger"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/connections"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/consumption"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/containerapps"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/containers"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cosmos"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/costmanagement"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/customproviders"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/dashboard"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/databasemigration"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/databoxedge"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/databricks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/datadog"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/datafactory"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/dataprotection"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/datashare"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/desktopvirtualization"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/devtestlabs"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/digitaltwins"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/disks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/dns"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/domainservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/elastic"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/elasticsan"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/eventgrid"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/eventhub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/firewall"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/fluidrelay"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/frontdoor"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/graphservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/hdinsight"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/healthcare"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/hsm"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/hybridcompute"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iotcentral"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iothub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iottimeseriesinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/keyvault"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/kusto"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/labservice"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/legacy"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/lighthouse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/loadbalancer"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/loadtestservice"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/loganalytics"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/logic"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/logz"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/machinelearning"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/maintenance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/managedapplications"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/managedhsm"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/managedidentity"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/managementgroup"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/maps"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mariadb"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/media"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mixedreality"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mobilenetwork"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/monitor"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mssql"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mssqlmanagedinstance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mysql"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/netapp"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/network"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/networkfunction"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/newrelic"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/nginx"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/notificationhub"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/orbital"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/paloalto"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/policy"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/portal"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/postgres"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/powerbi"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/privatedns"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/privatednsresolver"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/purview"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/recoveryservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/redhatopenshift"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/redis"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/redisenterprise"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/relay"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/resource"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/search"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/securitycenter"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/sentinel"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicebus"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/serviceconnector"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabric"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicefabricmanaged"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/servicenetworking"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/signalr"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/springcloud"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/sql"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/storage"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/storagecache"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/storagemover"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/streamanalytics"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/subscription"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/synapse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/systemcentervirtualmachinemanager"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/trafficmanager"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/videoanalyzer"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/vmware"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/voiceservices"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/web"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/workloads"
)

//go:generate go run ../tools/generator-services/main.go -path=../../

func SupportedTypedServices() []sdk.TypedServiceRegistration {
	services := []sdk.TypedServiceRegistration{
		aadb2c.Registration{},
		apimanagement.Registration{},
		appconfiguration.Registration{},
		applicationinsights.Registration{},
		appservice.Registration{},
		arckubernetes.Registration{},
		arcresourcebridge.Registration{},
		authorization.Registration{},
		automanage.Registration{},
		automation.Registration{},
		azurestackhci.Registration{},
		batch.Registration{},
		bot.Registration{},
		cognitive.Registration{},
		communication.Registration{},
		compute.Registration{},
		consumption.Registration{},
		containerapps.Registration{},
		cosmos.Registration{},
		costmanagement.Registration{},
		dashboard.Registration{},
		databoxedge.Registration{},
		databricks.Registration{},
		datafactory.Registration{},
		dataprotection.Registration{},
		desktopvirtualization.Registration{},
		digitaltwins.Registration{},
		disks.Registration{},
		domainservices.Registration{},
		elasticsan.Registration{},
		eventhub.Registration{},
		fluidrelay.Registration{},
		graphservices.Registration{},
		hybridcompute.Registration{},
		iotcentral.Registration{},
		iothub.Registration{},
		keyvault.Registration{},
		kusto.Registration{},
		labservice.Registration{},
		loadbalancer.Registration{},
		loadtestservice.Registration{},
		loganalytics.Registration{},
		machinelearning.Registration{},
		maintenance.Registration{},
		managedhsm.Registration{},
		media.Registration{},
		mobilenetwork.Registration{},
		monitor.Registration{},
		mssql.Registration{},
		mssqlmanagedinstance.Registration{},
		mysql.Registration{},
		netapp.Registration{},
		network.Registration{},
		networkfunction.Registration{},
		newrelic.Registration{},
		nginx.Registration{},
		orbital.Registration{},
		paloalto.Registration{},
		policy.Registration{},
		postgres.Registration{},
		privatednsresolver.Registration{},
		recoveryservices.Registration{},
		redhatopenshift.Registration{},
		redis.Registration{},
		resource.Registration{},
		search.Registration{},
		securitycenter.Registration{},
		sentinel.Registration{},
		serviceconnector.Registration{},
		servicefabricmanaged.Registration{},
		servicenetworking.Registration{},
		signalr.Registration{},
		springcloud.Registration{},
		storage.Registration{},
		storagecache.Registration{},
		storagemover.Registration{},
		streamanalytics.Registration{},
		subscription.Registration{},
		systemcentervirtualmachinemanager.Registration{},
		vmware.Registration{},
		voiceservices.Registration{},
		web.Registration{},
		workloads.Registration{},
	}
	services = append(services, autoRegisteredTypedServices()...)
	return services
}

func SupportedUntypedServices() []sdk.UntypedServiceRegistration {
	return func() []sdk.UntypedServiceRegistration {
		out := []sdk.UntypedServiceRegistration{
			advisor.Registration{},
			analysisservices.Registration{},
			apimanagement.Registration{},
			appconfiguration.Registration{},
			applicationinsights.Registration{},
			arckubernetes.Registration{},
			attestation.Registration{},
			authorization.Registration{},
			automation.Registration{},
			azurestackhci.Registration{},
			batch.Registration{},
			billing.Registration{},
			blueprints.Registration{},
			bot.Registration{},
			cdn.Registration{},
			cognitive.Registration{},
			compute.Registration{},
			confidentialledger.Registration{},
			connections.Registration{},
			consumption.Registration{},
			containers.Registration{},
			cosmos.Registration{},
			customproviders.Registration{},
			dashboard.Registration{},
			databasemigration.Registration{},
			databoxedge.Registration{},
			databricks.Registration{},
			datadog.Registration{},
			datafactory.Registration{},
			dataprotection.Registration{},
			datashare.Registration{},
			desktopvirtualization.Registration{},
			devtestlabs.Registration{},
			digitaltwins.Registration{},
			dns.Registration{},
			domainservices.Registration{},
			elastic.Registration{},
			eventgrid.Registration{},
			eventhub.Registration{},
			firewall.Registration{},
			frontdoor.Registration{},
			hdinsight.Registration{},
			healthcare.Registration{},
			hsm.Registration{},
			iotcentral.Registration{},
			iothub.Registration{},
			iottimeseriesinsights.Registration{},
			keyvault.Registration{},
			kusto.Registration{},
			legacy.Registration{},
			lighthouse.Registration{},
			loadbalancer.Registration{},
			loganalytics.Registration{},
			logic.Registration{},
			logz.Registration{},
			machinelearning.Registration{},
			maintenance.Registration{},
			managedapplications.Registration{},
			managedhsm.Registration{},
			managedidentity.Registration{},
			managementgroup.Registration{},
			maps.Registration{},
			mariadb.Registration{},
			media.Registration{},
			mixedreality.Registration{},
			monitor.Registration{},
			mssql.Registration{},
			mssqlmanagedinstance.Registration{},
			mysql.Registration{},
			netapp.Registration{},
			network.Registration{},
			notificationhub.Registration{},
			policy.Registration{},
			portal.Registration{},
			postgres.Registration{},
			powerbi.Registration{},
			privatedns.Registration{},
			purview.Registration{},
			recoveryservices.Registration{},
			redis.Registration{},
			redisenterprise.Registration{},
			relay.Registration{},
			resource.Registration{},
			search.Registration{},
			securitycenter.Registration{},
			sentinel.Registration{},
			servicebus.Registration{},
			servicefabric.Registration{},
			signalr.Registration{},
			springcloud.Registration{},
			sql.Registration{},
			storage.Registration{},
			storagecache.Registration{},
			streamanalytics.Registration{},
			subscription.Registration{},
			synapse.Registration{},
			trafficmanager.Registration{},
			videoanalyzer.Registration{},
			vmware.Registration{},
			web.Registration{},
		}
		return out
	}()
}
