//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_listcertificates.json
func ExampleCertificatesClient_ListByIotHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().ListByIotHub(ctx, "myResourceGroup", "testhub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateListDescription = armiothub.CertificateListDescription{
	// 	Value: []*armiothub.CertificateDescription{
	// 		{
	// 			Name: to.Ptr("cert"),
	// 			Type: to.Ptr("Microsoft.Devices/IotHubs/Certificates"),
	// 			Etag: to.Ptr("AAAAAAExpNs="),
	// 			ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/andbuc-hub/certificates/cert"),
	// 			Properties: &armiothub.CertificateProperties{
	// 				Certificate: to.Ptr("############################################"),
	// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 				Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 				IsVerified: to.Ptr(false),
	// 				Subject: to.Ptr("CN=testdevice1"),
	// 				Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_getcertificate.json
func ExampleCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().Get(ctx, "myResourceGroup", "testhub", "cert", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateDescription = armiothub.CertificateDescription{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpNs="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/andbuc-hub/certificates/cert"),
	// 	Properties: &armiothub.CertificateProperties{
	// 		Certificate: to.Ptr("############################################"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=testdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_certificatescreateorupdate.json
func ExampleCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().CreateOrUpdate(ctx, "myResourceGroup", "iothub", "cert", armiothub.CertificateDescription{
		Properties: &armiothub.CertificateProperties{
			Certificate: to.Ptr("############################################"),
		},
	}, &armiothub.CertificatesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateDescription = armiothub.CertificateDescription{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpNs="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServives/myFirstProvisioningService/certificates/cert"),
	// 	Properties: &armiothub.CertificateProperties{
	// 		Certificate: to.Ptr("############################################"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=testdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_certificatesdelete.json
func ExampleCertificatesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificatesClient().Delete(ctx, "myResourceGroup", "myhub", "cert", "AAAAAAAADGk=", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_generateverificationcode.json
func ExampleCertificatesClient_GenerateVerificationCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().GenerateVerificationCode(ctx, "myResourceGroup", "testHub", "cert", "AAAAAAAADGk=", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateWithNonceDescription = armiothub.CertificateWithNonceDescription{
	// 	Name: to.Ptr("cert"),
	// 	Properties: &armiothub.CertificatePropertiesWithNonce{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 		IsVerified: to.Ptr(false),
	// 		Subject: to.Ptr("CN=andbucdevice1"),
	// 		Thumbprint: to.Ptr("##############################"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:26:56 GMT"); return t}()),
	// 		VerificationCode: to.Ptr("##################################"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/32c178f2467f792a322f56174be244135d2c907f/specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_certverify.json
func ExampleCertificatesClient_Verify() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().Verify(ctx, "myResourceGroup", "myFirstProvisioningService", "cert", "AAAAAAAADGk=", armiothub.CertificateVerificationDescription{
		Certificate: to.Ptr("#####################################"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateDescription = armiothub.CertificateDescription{
	// 	Name: to.Ptr("cert"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs/Certificates"),
	// 	Etag: to.Ptr("AAAAAAExpTQ="),
	// 	ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/ProvisioningServices/myFirstProvisioningService/certificates/cert"),
	// 	Properties: &armiothub.CertificateProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:23:50 GMT"); return t}()),
	// 		Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 31 Dec 2039 23:59:59 GMT"); return t}()),
	// 		IsVerified: to.Ptr(true),
	// 		Subject: to.Ptr("CN=andbucdevice1"),
	// 		Thumbprint: to.Ptr("97388663832D0393C9246CAB4FBA2C8677185A25"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 12 Oct 2017 19:26:56 GMT"); return t}()),
	// 	},
	// }
}
