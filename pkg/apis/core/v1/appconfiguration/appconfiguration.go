package appconfiguration

import (
	"kusionstack.io/kusion/pkg/modules/inputs/workload"
)

type Accessory map[string]interface{}

// AppConfiguration is a developer-centric definition that describes how to run an Application. The application model is built on a decade
// of experience from AntGroup in operating a large-scale internal developer platform and combines the best ideas and practices from the
// community.
//
// Note: AppConfiguration per se is not a Kusion Module
//
// Example:
// import models.schema.v1 as ac
// import models.schema.v1.workload as wl
// import models.schema.v1.workload.container as c
// import models.schema.v1.workload.container.probe as p
// import models.schema.v1.monitoring as m
// import models.schema.v1.database as d
//
//	helloWorld: ac.AppConfiguration {
//	   # Built-in module
//	   workload: wl.Service {
//	       containers: {
//	           "main": c.Container {
//	               image: "ghcr.io/kusion-stack/samples/helloworld:latest"
//	               # Configure a HTTP readiness probe
//	               readinessProbe: p.Probe {
//	                   probeHandler: p.Http {
//	                       url: "http://localhost:80"
//	                   }
//	               }
//	           }
//	       }
//	   }
//
//	   # extend accessories module base
//	   accessories: {
//	       # Built-in module
//	       "mysql" : d.MySQL {
//	           type: "cloud"
//	           version: "8.0"
//	       }
//	       # Built-in module
//	       "pro" : m.Prometheus {
//	           path: "/metrics"
//	       }
//	       # Customized module
//	       "customize": customizedModule {
//	               ...
//	       }
//	   }
//
//	   # extend pipeline module base
//	   pipeline: {
//	       # Step is a module
//	       "step" : Step {
//	           use: "exec"
//	           args: ["--test-all"]
//	       }
//	   }
//
//	   # Dependent app list
//	   dependency: {
//	       dependentApps: ["init-kusion"]
//	   }
//	}
type AppConfiguration struct {
	// Name of the target Application.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Workload defines how to run your application code.
	Workload *workload.Workload `json:"workload" yaml:"workload"`
	// Accessories defines a collection of accessories that will be attached to the workload.
	Accessories map[string]*Accessory `json:"accessories,omitempty" yaml:"accessories,omitempty"`
	// Labels and Annotations can be used to attach arbitrary metadata as key-value pairs to resources.
	Labels      map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}
