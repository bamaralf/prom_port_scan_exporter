package main

import (
    "context"
    "fmt"
	"log"
	
	//"io/ioutil"
	//"github.com/ghodss/yaml"

    "github.com/ericchiang/k8s"
    corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func getPods() (map[string]string) {
	//client, err := loadClient("")
    client, err := k8s.NewInClusterClient()
    if err != nil {
     log.Fatal(err)
    }

/*
    var nodes corev1.NodeList
    if err := client.List(context.Background(), "", &nodes); err != nil {
        log.Fatal(err)
    }
    for _, node := range nodes.Items {
        fmt.Printf("name=%q schedulable=%t\n", *node.Metadata.Name, !*node.Spec.Unschedulable)
    }
*/

   // Pods in all namespaces
   var pods corev1.PodList
   
   if err = client.List(context.Background(), "", &pods); err != nil {
	   log.Fatal(err)
   }
   
   
   //var results []string
   portMap := make(map[string]string)

   for _, pod := range pods.Items {
       // Exclude pods that run on the host network
       if *pod.Status.PodIP != *pod.Status.HostIP {
        fmt.Printf("%q: %q, %q\n",*pod.Metadata.Name,*pod.Status.PodIP,*pod.Status.HostIP)
        //results = append(results, *pod.Status.PodIP)
          portMap[*pod.Metadata.Name] = *pod.Status.PodIP
       }
    }
       return portMap
 }

/*
// loadClient parses a kubeconfig from a file and returns a Kubernetes
// client. It does not support extensions or client auth providers.
func loadClient(kubeconfigPath string) (*k8s.Client, error) {
    data, err := ioutil.ReadFile(kubeconfigPath)
    if err != nil {
        return nil, fmt.Errorf("read kubeconfig: %v", err)
    }

    // Unmarshal YAML into a Kubernetes config object.
    var config k8s.Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
    }
    return k8s.NewClient(&config)
}
*/
