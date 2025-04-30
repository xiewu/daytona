// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"context"
	"fmt"
	"time"

	"github.com/daytonaio/daytona-ai-saas/cli/apiclient"
	"github.com/daytonaio/daytona-ai-saas/cli/config"
	"github.com/daytonaio/daytona-ai-saas/cli/internal/util"
	"github.com/daytonaio/daytona-ai-saas/cli/views/logs"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type ReadLogParams struct {
	Id        string
	ServerUrl string
	ServerApi config.ServerApi
	Follow    *bool
	From      *time.Time
}

func ReadBuildLogs(ctx context.Context, params ReadLogParams) {
	logs.SetupLongestPrefixLength([]string{params.Id})

	for {
		query := ""
		if params.Follow != nil && *params.Follow {
			query = "follow=true"
		}

		ws, res, err := util.GetWebsocketConn(ctx, fmt.Sprintf("/%s/build-logs", params.Id), params.ServerUrl, params.ServerApi, &query)
		// We want to retry getting the logs if it fails
		if err != nil {
			log.Trace(apiclient.HandleErrorResponse(res, err))
			time.Sleep(250 * time.Millisecond)
			continue
		}

		readJSONLog(ctx, ws, logs.FIRST_WORKSPACE_INDEX, nil)
		ws.Close()
		break
	}
}

func readJSONLog(ctx context.Context, ws *websocket.Conn, index int, from *time.Time) {
	logEntriesChan := make(chan logs.LogEntry)
	readErr := make(chan error)
	go func() {
		for {
			var logEntry logs.LogEntry

			err := ws.ReadJSON(&logEntry)

			// An empty entry will be sent from the server on close/EOF
			// We don't want to print that
			if logEntry != (logs.LogEntry{}) {
				logEntriesChan <- logEntry
			}

			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
					log.Error(err)
				}
				readErr <- err
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case logEntry := <-logEntriesChan:
			if from != nil {
				parsedTime, err := time.Parse(time.RFC3339Nano, logEntry.Time)
				if err != nil {
					log.Trace(err)
				}

				if parsedTime.After(*from) || parsedTime.Equal(*from) {
					logs.DisplayLogEntry(logEntry, index)
				}
			} else {
				logs.DisplayLogEntry(logEntry, index)
			}

		case err := <-readErr:
			if err != nil {
				err := ws.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second))
				if err != nil {
					log.Trace(err)
				}
				ws.Close()
				return
			}
		}
	}
}
