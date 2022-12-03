package pkg

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	v13 "k8s.io/client-go/informers/core/v1"
	v14 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	v12 "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
)

type controller struct {
	client        kubernetes.Interface
	ingressLister v1.IngressLister
	serviceLister v12.ServiceLister
	queue         workqueue.RateLimitingInterface
}

func NewController(client kubernetes.Interface, serviceInformer v13.ServiceInformer, ingessInformrer v14.IngressInformer) controller {
	c := controller{
		client:        client,
		ingressLister: ingessInformrer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManger"),
	}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})
	ingessInformrer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})

	return c
}

func (c *controller) addService(obj interface{}) {
	c.enqueue(obj)
}

func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	//todo 比较annotation
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enqueue(newObj)
}

func (c *controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c *controller) deleteIngress(obj interface{}) {

}

func (c *controller) Run(stopCh chan struct{}) {
	<-stopCh
}
