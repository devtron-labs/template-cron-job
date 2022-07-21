/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package pubsub

import (
	"github.com/devtron-labs/template-cron-job/pkg/event"
	"go.uber.org/zap"
)

type CronBasedEventReceiver interface {
	Subscribe() error
}

type CronBasedEventReceiverImpl struct {
	logger       *zap.SugaredLogger
	eventService event.EventService
}

func NewCronBasedEventReceiverImpl(logger *zap.SugaredLogger, eventService event.EventService) *CronBasedEventReceiverImpl {
	cronBasedEventReceiverImpl := &CronBasedEventReceiverImpl{
		logger:       logger,
		eventService: eventService,
	}
	return cronBasedEventReceiverImpl
}

func (impl *CronBasedEventReceiverImpl) Subscribe() error {
	return nil
}
