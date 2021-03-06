package v3

import (
	projectv3 "github.com/rancher/types/apis/project.cattle.io/v3"
	"github.com/rancher/types/image"
)

var (
	m = image.Mirror

	ToolsSystemImages = struct {
		PipelineSystemImages projectv3.PipelineSystemImages
		AuthSystemImages     AuthSystemImages
	}{
		PipelineSystemImages: projectv3.PipelineSystemImages{
			Jenkins:       m("rancher/pipeline-jenkins-server:v0.1.4"),
			JenkinsJnlp:   m("jenkins/jnlp-slave:3.10-1-alpine"),
			AlpineGit:     m("rancher/pipeline-tools:v0.1.13"),
			PluginsDocker: m("plugins/docker:17.12"),
			Minio:         m("minio/minio:RELEASE.2018-05-25T19-49-13Z"),
			Registry:      m("registry:2"),
			RegistryProxy: m("rancher/pipeline-tools:v0.1.13"),
			KubeApply:     m("rancher/pipeline-tools:v0.1.13"),
		},
		AuthSystemImages: AuthSystemImages{
			KubeAPIAuth: m("rancher/kube-api-auth:v0.1.3"),
		},
	}
)
