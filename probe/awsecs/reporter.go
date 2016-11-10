
package awsecs


struct labelInfo {
	cluster string
	taskArn string
	family string
}

func getLabelInfos() []labelInfo {
	// TODO get labels for all containers and look for com.amazonaws.ecs.{cluster,task-arn,task-definition-family}
}

func getRegion() string {
	// TODO GET http://169.254.169.254/<whatever version is latest atm>/dynamic/instance-identity/document,
	// parse as json, extract .region
}

// returns a list of services in the cluster
func getServices(cluster string) []string {
	// TODO aws list-services with given cluster, extract .serviceArns
}

// returns a map from deployment ids to service names
func getDeploymentMap(cluster string) map[string]string {
	// TODO describe-services with given cluster and getServices()
	// for each service, for each deployment in service["deployments"],
	// results[deployment["id"]] = service["serviceName"]
}

// returns a map from task ARNs to deployment ids
func getTaskDeployments(cluster string, taskArns []string) map[string]string {
	// TODO describe-tasks with given cluster and task ARNs
	// for each task, results[task["taskArn"]] = task["startedBy"]
}

// returns a map from task ARNs to service names
func getTaskServices(cluster string, taskArns []string) map[string]string {
	deploymentMapChan := make(chan map[string] string)
	go func() {
		deploymentMapChan <- getDeploymentMap(cluster)
	}()

	// do these two fetches in parallel
	taskDeployments := getTaskDeployments(cluster, taskArns)
	deploymentMap := <-deploymentMapChan

	results := make(map[string] string)
	for taskArn, depID := range taskDeployments {
		if service, ok := deploymentMap[depID]; ok {
			results[taskArn] = service
		} else {
			// TODO log warning - couldn't find service for task
		}
	}

	return results
}
