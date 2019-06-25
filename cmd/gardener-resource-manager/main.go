// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"

	"github.com/gardener/gardener-resource-manager/cmd/gardener-resource-manager/app"

	"sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

func main() {
	log.SetLogger(log.ZapLogger(false))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		<-signals.SetupSignalHandler()
	}()

	cmd := app.NewControllerManagerCommand(ctx)

	if err := cmd.Execute(); err != nil {
		log.Log.Error(err, "error executing the main controller command")
		os.Exit(1)
	}
}