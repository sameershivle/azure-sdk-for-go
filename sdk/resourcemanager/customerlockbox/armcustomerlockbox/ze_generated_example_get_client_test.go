//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/TenantOptedIn.json
func ExampleGetClient_TenantOptedIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerlockbox.NewGetClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.TenantOptedIn(ctx,
		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
