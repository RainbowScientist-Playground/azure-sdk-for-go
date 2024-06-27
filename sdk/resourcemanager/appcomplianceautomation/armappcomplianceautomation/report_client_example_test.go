//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_List.json
func ExampleReportClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReportClient().NewListPager(&armappcomplianceautomation.ReportClientListOptions{SkipToken: to.Ptr("1"),
		Top:                   to.Ptr[int32](100),
		Select:                nil,
		Filter:                nil,
		Orderby:               nil,
		OfferGUID:             to.Ptr("00000000-0000-0000-0000-000000000000"),
		ReportCreatorTenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ReportResourceListResult = armappcomplianceautomation.ReportResourceListResult{
		// 	Value: []*armappcomplianceautomation.ReportResource{
		// 		{
		// 			Name: to.Ptr("testReportName"),
		// 			Type: to.Ptr("Microsfot.AppComplianceAutomation/reports"),
		// 			ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName"),
		// 			SystemData: &armappcomplianceautomation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 				CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 			},
		// 			Properties: &armappcomplianceautomation.ReportProperties{
		// 				CertRecords: []*armappcomplianceautomation.CertSyncRecord{
		// 					{
		// 						CertificationStatus: to.Ptr("CertIngestion"),
		// 						Controls: []*armappcomplianceautomation.ControlSyncRecord{
		// 							{
		// 								ControlID: to.Ptr("Operational_Security_10"),
		// 								ControlStatus: to.Ptr("Approved"),
		// 						}},
		// 						IngestionStatus: to.Ptr("EvidenceResubmitted"),
		// 						OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 				}},
		// 				ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
		// 					M365: &armappcomplianceautomation.OverviewStatus{
		// 						FailedCount: to.Ptr[int32](0),
		// 						ManualCount: to.Ptr[int32](0),
		// 						NotApplicableCount: to.Ptr[int32](0),
		// 						PassedCount: to.Ptr[int32](0),
		// 						PendingCount: to.Ptr[int32](0),
		// 					},
		// 				},
		// 				Errors: []*string{
		// 				},
		// 				LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
		// 				NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
		// 				OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"),
		// 				ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
		// 				Resources: []*armappcomplianceautomation.ResourceMetadata{
		// 					{
		// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
		// 						ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
		// 						ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
		// 					},
		// 					{
		// 						AccountID: to.Ptr("000000000000"),
		// 						ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
		// 						ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
		// 						ResourceType: to.Ptr("iam.user"),
		// 				}},
		// 				Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
		// 				StorageInfo: &armappcomplianceautomation.StorageInfo{
		// 					AccountName: to.Ptr("testStorageAccount"),
		// 					Location: to.Ptr("East US"),
		// 					ResourceGroup: to.Ptr("testResourceGroup"),
		// 					SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 				Subscriptions: []*string{
		// 					to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TimeZone: to.Ptr("GMT Standard Time"),
		// 					TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_Get.json
func ExampleReportClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().Get(ctx, "testReport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportResource = armappcomplianceautomation.ReportResource{
	// 	Name: to.Ptr("testReportName"),
	// 	Type: to.Ptr("Microsfot.AppComplianceAutomation/reports"),
	// 	ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName"),
	// 	SystemData: &armappcomplianceautomation.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcomplianceautomation.ReportProperties{
	// 		CertRecords: []*armappcomplianceautomation.CertSyncRecord{
	// 			{
	// 				CertificationStatus: to.Ptr("CertIngestion"),
	// 				Controls: []*armappcomplianceautomation.ControlSyncRecord{
	// 					{
	// 						ControlID: to.Ptr("Operational_Security_10"),
	// 						ControlStatus: to.Ptr("Approved"),
	// 				}},
	// 				IngestionStatus: to.Ptr("EvidenceResubmitted"),
	// 				OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 		}},
	// 		ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
	// 			M365: &armappcomplianceautomation.OverviewStatus{
	// 				FailedCount: to.Ptr[int32](0),
	// 				ManualCount: to.Ptr[int32](0),
	// 				PassedCount: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		Errors: []*string{
	// 			to.Ptr("resource-inaccessible")},
	// 			LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 			NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 			OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"),
	// 			ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
	// 			Resources: []*armappcomplianceautomation.ResourceMetadata{
	// 				{
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 					ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
	// 					ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
	// 				},
	// 				{
	// 					AccountID: to.Ptr("000000000000"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
	// 					ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
	// 					ResourceType: to.Ptr("iam.user"),
	// 			}},
	// 			Status: to.Ptr(armappcomplianceautomation.ReportStatusFailed),
	// 			StorageInfo: &armappcomplianceautomation.StorageInfo{
	// 				AccountName: to.Ptr("testStorageAccount"),
	// 				Location: to.Ptr("East US"),
	// 				ResourceGroup: to.Ptr("testResourceGroup"),
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			Subscriptions: []*string{
	// 				to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				TimeZone: to.Ptr("GMT Standard Time"),
	// 				TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_CreateOrUpdate.json
func ExampleReportClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginCreateOrUpdate(ctx, "testReportName", armappcomplianceautomation.ReportResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportResource = armappcomplianceautomation.ReportResource{
	// 	Name: to.Ptr("testReportName"),
	// 	Type: to.Ptr("Microsfot.AppComplianceAutomation/reports"),
	// 	ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName"),
	// 	SystemData: &armappcomplianceautomation.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcomplianceautomation.ReportProperties{
	// 		CertRecords: []*armappcomplianceautomation.CertSyncRecord{
	// 			{
	// 				CertificationStatus: to.Ptr("CertIngestion"),
	// 				Controls: []*armappcomplianceautomation.ControlSyncRecord{
	// 					{
	// 						ControlID: to.Ptr("Operational_Security_10"),
	// 						ControlStatus: to.Ptr("Approved"),
	// 				}},
	// 				IngestionStatus: to.Ptr("EvidenceResubmitted"),
	// 				OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 		}},
	// 		ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
	// 			M365: &armappcomplianceautomation.OverviewStatus{
	// 				FailedCount: to.Ptr[int32](0),
	// 				ManualCount: to.Ptr[int32](0),
	// 				PassedCount: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"),
	// 		ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
	// 		Resources: []*armappcomplianceautomation.ResourceMetadata{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 				ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
	// 				ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
	// 			},
	// 			{
	// 				AccountID: to.Ptr("000000000000"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
	// 				ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
	// 				ResourceType: to.Ptr("iam.user"),
	// 		}},
	// 		Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
	// 		StorageInfo: &armappcomplianceautomation.StorageInfo{
	// 			AccountName: to.Ptr("testStorageAccount"),
	// 			Location: to.Ptr("East US"),
	// 			ResourceGroup: to.Ptr("testResourceGroup"),
	// 			SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Subscriptions: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TimeZone: to.Ptr("GMT Standard Time"),
	// 			TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_Update.json
func ExampleReportClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginUpdate(ctx, "testReportName", armappcomplianceautomation.ReportResourcePatch{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportResource = armappcomplianceautomation.ReportResource{
	// 	Name: to.Ptr("testReportName"),
	// 	Type: to.Ptr("Microsfot.AppComplianceAutomation/reports"),
	// 	ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName"),
	// 	SystemData: &armappcomplianceautomation.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcomplianceautomation.ReportProperties{
	// 		CertRecords: []*armappcomplianceautomation.CertSyncRecord{
	// 			{
	// 				CertificationStatus: to.Ptr("CertIngestion"),
	// 				Controls: []*armappcomplianceautomation.ControlSyncRecord{
	// 					{
	// 						ControlID: to.Ptr("Operational_Security_10"),
	// 						ControlStatus: to.Ptr("Approved"),
	// 				}},
	// 				IngestionStatus: to.Ptr("EvidenceResubmitted"),
	// 				OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 		}},
	// 		ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
	// 			M365: &armappcomplianceautomation.OverviewStatus{
	// 				FailedCount: to.Ptr[int32](0),
	// 				ManualCount: to.Ptr[int32](0),
	// 				PassedCount: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"),
	// 		ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
	// 		Resources: []*armappcomplianceautomation.ResourceMetadata{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 				ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
	// 				ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
	// 			},
	// 			{
	// 				AccountID: to.Ptr("000000000000"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
	// 				ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
	// 				ResourceType: to.Ptr("iam.user"),
	// 		}},
	// 		Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
	// 		StorageInfo: &armappcomplianceautomation.StorageInfo{
	// 			AccountName: to.Ptr("testStorageAccount"),
	// 			Location: to.Ptr("East US"),
	// 			ResourceGroup: to.Ptr("testResourceGroup"),
	// 			SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Subscriptions: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TimeZone: to.Ptr("GMT Standard Time"),
	// 			TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_Delete.json
func ExampleReportClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginDelete(ctx, "testReportName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_NestedResourceCheckNameAvailability_Report_Evidence_Check_Name_Availability.json
func ExampleReportClient_NestedResourceCheckNameAvailability_reportEvidenceCheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().NestedResourceCheckNameAvailability(ctx, "reportABC", armappcomplianceautomation.CheckNameAvailabilityRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armappcomplianceautomation.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr("An evidence named 'evidenceABC' is already in use."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armappcomplianceautomation.CheckNameAvailabilityReasonAlreadyExists),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_NestedResourceCheckNameAvailability_Report_Snapshot_Check_Name_Availability.json
func ExampleReportClient_NestedResourceCheckNameAvailability_reportSnapshotCheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().NestedResourceCheckNameAvailability(ctx, "reportABC", armappcomplianceautomation.CheckNameAvailabilityRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armappcomplianceautomation.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr("An snapshot named 'snapshotABC' is already in use."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armappcomplianceautomation.CheckNameAvailabilityReasonAlreadyExists),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_NestedResourceCheckNameAvailability_Report_Webhook_Check_Name_Availability.json
func ExampleReportClient_NestedResourceCheckNameAvailability_reportWebhookCheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().NestedResourceCheckNameAvailability(ctx, "reportABC", armappcomplianceautomation.CheckNameAvailabilityRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResponse = armappcomplianceautomation.CheckNameAvailabilityResponse{
	// 	Message: to.Ptr("An webhook named 'webhookABC' is already in use."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armappcomplianceautomation.CheckNameAvailabilityReasonAlreadyExists),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_Fix.json
func ExampleReportClient_BeginFix() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginFix(ctx, "testReport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportFixResult = armappcomplianceautomation.ReportFixResult{
	// 	Reason: to.Ptr(""),
	// 	Result: to.Ptr(armappcomplianceautomation.ResultSucceeded),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_GetScopingQuestions.json
func ExampleReportClient_GetScopingQuestions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().GetScopingQuestions(ctx, "testReportName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopingQuestions = armappcomplianceautomation.ScopingQuestions{
	// 	Questions: []*armappcomplianceautomation.ScopingQuestion{
	// 		{
	// 			InputType: to.Ptr(armappcomplianceautomation.InputTypeBoolean),
	// 			OptionIDs: []*string{
	// 			},
	// 			QuestionID: to.Ptr("DHP_G07_customerDataProcess"),
	// 			Rules: []*armappcomplianceautomation.Rule{
	// 				to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 				ShowSubQuestionsValue: to.Ptr("true"),
	// 			},
	// 			{
	// 				InputType: to.Ptr(armappcomplianceautomation.InputTypeText),
	// 				OptionIDs: []*string{
	// 				},
	// 				QuestionID: to.Ptr("DHP_G04_graphPermissionData"),
	// 				Rules: []*armappcomplianceautomation.Rule{
	// 					to.Ptr(armappcomplianceautomation.RuleRequired),
	// 					to.Ptr(armappcomplianceautomation.RuleCharLength)},
	// 					SuperiorQuestionID: to.Ptr("DHP_G07_customerDataProcess"),
	// 				},
	// 				{
	// 					InputType: to.Ptr(armappcomplianceautomation.InputTypeBoolean),
	// 					OptionIDs: []*string{
	// 					},
	// 					QuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 					Rules: []*armappcomplianceautomation.Rule{
	// 						to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 						ShowSubQuestionsValue: to.Ptr("true"),
	// 					},
	// 					{
	// 						InputType: to.Ptr(armappcomplianceautomation.InputTypeText),
	// 						OptionIDs: []*string{
	// 						},
	// 						QuestionID: to.Ptr("DHP_G05_graphPermissionInfo"),
	// 						Rules: []*armappcomplianceautomation.Rule{
	// 							to.Ptr(armappcomplianceautomation.RuleRequired),
	// 							to.Ptr(armappcomplianceautomation.RuleCharLength),
	// 							to.Ptr(armappcomplianceautomation.RulePreventNonEnglishChar)},
	// 							SuperiorQuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 						},
	// 						{
	// 							InputType: to.Ptr(armappcomplianceautomation.InputTypeMultiSelectDropdown),
	// 							OptionIDs: []*string{
	// 								to.Ptr("Croatia"),
	// 								to.Ptr("Cuba"),
	// 								to.Ptr("Curaçao"),
	// 								to.Ptr("Cyprus"),
	// 								to.Ptr("Czechia"),
	// 								to.Ptr("Côte d'Ivoire"),
	// 								to.Ptr("Denmark"),
	// 								to.Ptr("Djibouti"),
	// 								to.Ptr("Dominica"),
	// 								to.Ptr("Dominican Republic (the)"),
	// 								to.Ptr("Ecuador"),
	// 								to.Ptr("Egypt")},
	// 								QuestionID: to.Ptr("DHP_G08_storageLocation"),
	// 								Rules: []*armappcomplianceautomation.Rule{
	// 									to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 									SuperiorQuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 								},
	// 								{
	// 									InputType: to.Ptr(armappcomplianceautomation.InputType("SingleSelectEnum")),
	// 									OptionIDs: []*string{
	// 									},
	// 									QuestionID: to.Ptr("LEG03_complianceDataTermination"),
	// 									Rules: []*armappcomplianceautomation.Rule{
	// 										to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 										SuperiorQuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 								}},
	// 							}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_SyncCertRecord.json
func ExampleReportClient_BeginSyncCertRecord() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginSyncCertRecord(ctx, "testReportName", armappcomplianceautomation.SyncCertRecordRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SyncCertRecordResponse = armappcomplianceautomation.SyncCertRecordResponse{
	// 	CertRecord: &armappcomplianceautomation.CertSyncRecord{
	// 		CertificationStatus: to.Ptr("CertIngestion"),
	// 		Controls: []*armappcomplianceautomation.ControlSyncRecord{
	// 		},
	// 		IngestionStatus: to.Ptr("InitialDocumentResubmitted"),
	// 		OfferGUID: to.Ptr("addb13fc-64bf-4005-b693-4c2f094e2187"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_Verify.json
func ExampleReportClient_BeginVerify() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginVerify(ctx, "testReport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportVerificationResult = armappcomplianceautomation.ReportVerificationResult{
	// 	Reason: to.Ptr(""),
	// 	Result: to.Ptr(armappcomplianceautomation.ResultSucceeded),
	// }
}
