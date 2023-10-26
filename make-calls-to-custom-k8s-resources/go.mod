module make-calls-to-custom-k8s-resources

go 1.21.0

require (
	git.heb.com/kub/composition-functions/resources/utils v0.0.0
	git.heb.com/provider-simplejsonapp v0.0.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/crossplane-contrib/provider-kubernetes v0.10.0-rc.0.0.20231002102226-f298faeba59f // indirect
	github.com/crossplane/crossplane-runtime v1.14.0-rc.0.0.20231003095054-b53745680067 // indirect
	github.com/crossplane/function-sdk-go v0.0.0-20230930011419-ec31b88ab696 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.10.2 // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rancher/fleet/pkg/apis v0.0.0-20231003153507-1353dee96d14 // indirect
	github.com/rancher/wrangler v1.1.1 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.28.3 // indirect
	k8s.io/apimachinery v0.28.3 // indirect
	k8s.io/client-go v0.28.3 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
	sigs.k8s.io/cli-utils v0.27.0 // indirect
	sigs.k8s.io/controller-runtime v0.16.2 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	git.heb.com/kub/composition-functions/resources/utils v0.0.0 => gitlab.com/heb-engineering/teams/platform-engineering/gke-hybrid-cloud/kon/crossplane/composition-functions/resources/utils.git v0.3.1
	git.heb.com/provider-simplejsonapp v0.0.0 => gitlab.com/heb-engineering/teams/platform-engineering/gke-hybrid-cloud/kon/crossplane/learning/simplejsonapp/provider-simplejsonapp.git v0.1.0-temp
	github.com/crossplane/function-sdk-go v0.0.0-20230930011419-ec31b88ab696 => github.com/dalton-hill-0/function-sdk-go v0.0.0-20231003210841-c5c85a7dc78f
)
