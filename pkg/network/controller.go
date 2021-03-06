/*
Copyright 2015 Juniper Networks, Inc. All rights reserved.

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

package network

import (
	"io"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/cache"
)

type NetworkController interface {
	Init(global *Config, configFile io.Reader) error

	SetNamespaceStore(store cache.Store)
	AddNamespace(obj *api.Namespace)
	UpdateNamespace(oldObj, newObj *api.Namespace)
	DeleteNamespace(obj *api.Namespace)

	SetPodStore(store cache.Store)
	AddPod(obj *api.Pod)
	UpdatePod(oldObj, newObj *api.Pod)
	DeletePod(obj *api.Pod)

	SetReplicationControllerStore(store cache.Store)
	AddReplicationController(obj *api.ReplicationController)
	UpdateReplicationController(oldObj, newObj *api.ReplicationController)
	DeleteReplicationController(obj *api.ReplicationController)

	SetServiceStore(store cache.Store)
	AddService(obj *api.Service)
	UpdateService(oldObj, newObj *api.Service)
	DeleteService(obj *api.Service)

	Run(shutdown chan struct{})
}
