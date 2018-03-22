/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package controller

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	"github.com/marun/fnord/pkg/controller/federatedconfigmap"
	"github.com/marun/fnord/pkg/controller/federatedconfigmapoverride"
	"github.com/marun/fnord/pkg/controller/federatedconfigmapplacement"
	"github.com/marun/fnord/pkg/controller/federateddeployment"
	"github.com/marun/fnord/pkg/controller/federateddeploymentoverride"
	"github.com/marun/fnord/pkg/controller/federateddeploymentplacement"
	"github.com/marun/fnord/pkg/controller/federatedreplicaset"
	"github.com/marun/fnord/pkg/controller/federatedreplicasetoverride"
	"github.com/marun/fnord/pkg/controller/federatedreplicasetplacement"
	"github.com/marun/fnord/pkg/controller/federatedsecret"
	"github.com/marun/fnord/pkg/controller/federatedsecretoverride"
	"github.com/marun/fnord/pkg/controller/federatedsecretplacement"
	"github.com/marun/fnord/pkg/controller/propagatedversion"
	"github.com/marun/fnord/pkg/controller/sharedinformers"
	"k8s.io/client-go/rest"
)

func GetAllControllers(config *rest.Config) ([]controller.Controller, chan struct{}) {
	shutdown := make(chan struct{})
	si := sharedinformers.NewSharedInformers(config, shutdown)
	return []controller.Controller{
		federatedconfigmap.NewFederatedConfigMapController(config, si),
		federatedconfigmapoverride.NewFederatedConfigMapOverrideController(config, si),
		federatedconfigmapplacement.NewFederatedConfigMapPlacementController(config, si),
		federateddeployment.NewFederatedDeploymentController(config, si),
		federateddeploymentoverride.NewFederatedDeploymentOverrideController(config, si),
		federateddeploymentplacement.NewFederatedDeploymentPlacementController(config, si),
		federatedreplicaset.NewFederatedReplicaSetController(config, si),
		federatedreplicasetoverride.NewFederatedReplicaSetOverrideController(config, si),
		federatedreplicasetplacement.NewFederatedReplicaSetPlacementController(config, si),
		federatedsecret.NewFederatedSecretController(config, si),
		federatedsecretoverride.NewFederatedSecretOverrideController(config, si),
		federatedsecretplacement.NewFederatedSecretPlacementController(config, si),
		propagatedversion.NewPropagatedVersionController(config, si),
	}, shutdown
}
