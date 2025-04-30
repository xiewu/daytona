// Copyright 2025 Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package common

import (
	"context"
	"fmt"
	"time"

	"github.com/daytonaio/daytona-ai-saas/cli/apiclient"
	daytonaapiclient "github.com/daytonaio/daytona-ai-saas/daytonaapiclient"
)

func AwaitImageActive(ctx context.Context, apiClient *daytonaapiclient.APIClient, targetImage string) error {
	for {
		images, res, err := apiClient.ImagesAPI.GetAllImages(ctx).Execute()
		if err != nil {
			return apiclient.HandleErrorResponse(res, err)
		}

		for _, image := range images.Items {
			if image.Name == targetImage {
				if image.State == daytonaapiclient.IMAGESTATE_ACTIVE {
					return nil
				} else if image.State == daytonaapiclient.IMAGESTATE_ERROR {
					if !image.ErrorReason.IsSet() {
						return fmt.Errorf("image processing failed")
					}
					return fmt.Errorf("image processing failed: %s", *image.ErrorReason.Get())
				}
			}
		}

		time.Sleep(time.Second)
	}
}

func AwaitSandboxStarted(ctx context.Context, apiClient *daytonaapiclient.APIClient, targetSandbox string) error {
	for {
		sandboxes, res, err := apiClient.WorkspaceAPI.ListWorkspaces(ctx).Execute()
		if err != nil {
			return apiclient.HandleErrorResponse(res, err)
		}

		for _, sandbox := range sandboxes {
			if sandbox.Id == targetSandbox {
				if sandbox.State != nil && *sandbox.State == daytonaapiclient.WORKSPACESTATE_STARTED {
					return nil
				} else if sandbox.State != nil && *sandbox.State == daytonaapiclient.WORKSPACESTATE_ERROR {
					if sandbox.ErrorReason == nil {
						return fmt.Errorf("sandbox processing failed")
					}
					return fmt.Errorf("sandbox processing failed: %s", *sandbox.ErrorReason)
				}
			}
		}

		time.Sleep(time.Second)
	}
}
