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

package controller

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	crv1 "k8s.io/apiextensions-apiserver/examples/gopher-k8s/apis/cr/v1"
)

// Watcher is an astaXie of watching on resource create/update/delete events
type AstaXieController struct {
	AstaXieClient *rest.RESTClient
	AstaXieScheme *runtime.Scheme
}

// Run starts an AstaXie resource controller
func (c *AstaXieController) Run(stopCh <-chan struct{}) error {
	source := cache.NewListWatchFromClient(
		c.AstaXieClient,
		crv1.AstaXieResourcePlural,
		apiv1.NamespaceAll,
		fields.Everything())

	_, controller := cache.NewInformer(
		source,

		// The object type.
		&crv1.AstaXie{},

		// resyncPeriod
		// Every resyncPeriod, all resources in the cache will retrigger events.
		// Set to 0 to disable the resync.
		0,

		// Your custom resource event handlers.
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.onAdd,
			UpdateFunc: c.onUpdate,
			DeleteFunc: c.onDelete,
		})

	go controller.Run(stopCh)
	<-stopCh
	return nil
}

func (c *AstaXieController) onAdd(obj interface{}) {
	astaXie := obj.(*crv1.AstaXie)
	fmt.Printf("[CONTROLLER] OnAdd %s\n", astaXie.ObjectMeta.SelfLink)

	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use astaXieScheme.Copy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	copyObj, err := c.AstaXieScheme.Copy(astaXie)
	if err != nil {
		fmt.Printf("ERROR creating a deep copy of astaXie object: %v\n", err)
		return
	}

	astaXieCopy := copyObj.(*crv1.AstaXie)
	astaXieCopy.Status = crv1.AstaXieStatus{
		State:   crv1.AstaXieStateAccepted,
		Message: "Asta Xie successfully accepted invitation and joined Kubernetes community!",
	}

	err = c.AstaXieClient.Put().
		Name(astaXie.ObjectMeta.Name).
		Namespace(astaXie.ObjectMeta.Namespace).
		Resource(crv1.AstaXieResourcePlural).
		Body(astaXieCopy).
		Do().
		Error()

	if err != nil {
		fmt.Printf("ERROR updating status: %v\n", err)
	}

	// Fetch a list of our CRs
	astaXieList := crv1.AstaXieList{}
	err = c.AstaXieClient.Get().Resource(crv1.AstaXieResourcePlural).Do().Into(&astaXieList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("LIST: %#v\n", astaXieList)
}

func (c *AstaXieController) onUpdate(oldObj, newObj interface{}) {
	oldAstaXie := oldObj.(*crv1.AstaXie)
	newAstaXie := newObj.(*crv1.AstaXie)
	fmt.Printf("[CONTROLLER] OnUpdate oldObj: %s\n", oldAstaXie.ObjectMeta.SelfLink)
	fmt.Printf("[CONTROLLER] OnUpdate newObj: %s\n", newAstaXie.ObjectMeta.SelfLink)
}

func (c *AstaXieController) onDelete(obj interface{}) {
	astaXie := obj.(*crv1.AstaXie)
	fmt.Printf("[CONTROLLER] OnDelete %s\n", astaXie.ObjectMeta.SelfLink)
}
