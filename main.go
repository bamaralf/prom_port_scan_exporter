
package main

import (
  "fmt"
  "net/http"
   //"time"
  log "github.com/Sirupsen/logrus"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

      var openedPorts = promauto.NewGaugeVec(prometheus.GaugeOpts{
		        Name:        "opened_ports_sum",
		        Help:        "Current number of opened TCP ports.",
		        //ConstLabels: prometheus.Labels{"pod": "podName"},
          },
          []string{"Host"},
      )

/*
     var openedPortsHist = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "opened_ports",
            Help:    "Opened ports distribution",
            Buckets: prometheus.LinearBuckets(1, 10000, 7),
        },
        []string{"TCPPortHist"},
      )
*/    
    func getMetrics(){
      go func() {
        for {
                 pods := getPods()
                 for podName, podIP := range pods {
                     ps := portScan(podIP)
                     fmt.Printf("%q:", podName)
                     psLen := float64(len(ps))
                     //for i := 0; i < len(ps); i++ {
                     //    openedPortsHist.WithLabelValues(host).Observe(float64(ps[i]))
                     openedPorts.WithLabelValues(podName).Set(psLen)
                    // }
                     
                 }
                //time.Sleep(10 * time.Second)
        }
      }()
    }

/*    
     func init() {
      // prometheus.MustRegister(openedPorts)
       prometheus.MustRegister(openedPortsHist)
     }
*/
func main() {
  getMetrics()
  
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :8081")
  http.ListenAndServe(":8081", nil)
}
