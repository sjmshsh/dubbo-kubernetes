/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package services

import (
	"context"

	"github.com/apache/dubbo-kubernetes/pkg/admin/model"
)

type TestingService interface {
	GetMethodsNames(ctx context.Context, target, serviceName string) ([]string, error)
	GetMethodDescribe(ctx context.Context, target, methodName string) (*model.MethodDescribe, error)
	GetMessageTemplateString(ctx context.Context, target, messageName string) (string, error)
	GetMessageDescribeString(ctx context.Context, target, messageName string) (string, error)
	Invoke(ctx context.Context, target, methodName, input string, headers map[string]string) (string, bool, error)
}
