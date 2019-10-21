package k8s

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/luizalabs/teresa/pkg/server/app"
	"github.com/luizalabs/teresa/pkg/server/spec"
	"k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAddVolumeMountOfSecrets(t *testing.T) {
	var testCases = []struct {
		vols       []k8sv1.VolumeMount
		volName    string
		path       string
		expectedVM int
	}{
		{
			vols:       []k8sv1.VolumeMount{{Name: "vm"}},
			volName:    "s",
			path:       "p",
			expectedVM: 2,
		},
		{
			vols:       []k8sv1.VolumeMount{{Name: "s"}},
			volName:    "s",
			path:       "p",
			expectedVM: 1,
		},
		{
			vols:       []k8sv1.VolumeMount{{Name: "vm"}, {Name: "s"}},
			volName:    "s",
			path:       "p",
			expectedVM: 2,
		},
	}

	for _, tc := range testCases {
		vols := addVolumeMountOfSecrets(tc.vols, tc.volName, tc.path)
		if actual := len(vols); actual != tc.expectedVM {
			t.Errorf("expected %d vols, got %d", tc.expectedVM, actual)
		}
		found := false
		for _, vol := range vols {
			if vol.Name == tc.volName {
				found = true
				break
			}
		}
		if !found {
			t.Error("volume ref to secret not found")
		}
	}
}

func TestAddVolumeOfSecretFile(t *testing.T) {
	var testCases = []struct {
		vols       []k8sv1.Volume
		volName    string
		secretName string
		fileName   string
	}{
		{
			vols:       []k8sv1.Volume{{Name: "vl"}},
			volName:    "volSecret",
			secretName: "secret",
			fileName:   "fileSecret",
		},
		{
			vols: []k8sv1.Volume{{
				Name: "volSecret",
				VolumeSource: k8sv1.VolumeSource{
					Secret: &k8sv1.SecretVolumeSource{
						SecretName: "secret",
						Items:      []k8sv1.KeyToPath{},
					},
				},
			}},
			volName:    "volSecret",
			secretName: "secret",
			fileName:   "fileSecret",
		},
		{
			vols: []k8sv1.Volume{{
				Name: "volSecret",
				VolumeSource: k8sv1.VolumeSource{
					Secret: &k8sv1.SecretVolumeSource{
						SecretName: "secret",
						Items:      []k8sv1.KeyToPath{{Key: "fs", Path: "fs"}},
					},
				},
			}},
			volName:    "volSecret",
			secretName: "secret",
			fileName:   "fs",
		},
	}

	for _, tc := range testCases {
		vols := addVolumeOfSecretFile(tc.vols, tc.volName, tc.secretName, tc.fileName)
		found := false
		for _, vol := range vols {
			if vol.Name == tc.volName {
				found = true
				if actual := vol.Secret.SecretName; actual != tc.secretName {
					t.Errorf("expected %s, got %s", tc.secretName, actual)
				}
				foundItem := false
				for _, item := range vol.Secret.Items {
					if item.Key == tc.fileName {
						foundItem = true
						break
					}
				}
				if !foundItem {
					t.Error("item not found in volume")
				}
				break
			}
		}
		if !found {
			t.Error("volume ref to secret not found")
		}
	}
}

func deploySpec(envs []k8sv1.EnvVar) *v1beta2.Deployment {
	evs := make([]k8sv1.EnvVar, len(envs))
	copy(evs, envs)

	return &v1beta2.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "test"},
		Spec: v1beta2.DeploymentSpec{
			Template: k8sv1.PodTemplateSpec{
				Spec: k8sv1.PodSpec{
					Containers: []k8sv1.Container{
						{
							Name: "test",
							Env:  evs,
						},
					},
				},
			},
		},
	}
}

func cronJobSpec(envs []k8sv1.EnvVar) *v1beta1.CronJob {
	evs := make([]k8sv1.EnvVar, len(envs))
	copy(evs, envs)

	return &v1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{Name: "test"},
		Spec: v1beta1.CronJobSpec{
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: k8sv1.PodTemplateSpec{
						Spec: k8sv1.PodSpec{
							Containers: []k8sv1.Container{
								{
									Name: "test",
									Env:  evs,
								},
							},
						},
					},
				},
			},
		},
	}
}

func TestRemoveEnvVarsWithSecretsFromDeployAndCronJob(t *testing.T) {
	var testCases = []struct {
		envs     []k8sv1.EnvVar
		toRemove []string
		expected []k8sv1.EnvVar
	}{
		{
			envs:     []k8sv1.EnvVar{{Name: "FOO"}},
			toRemove: []string{"FOO"},
			expected: []k8sv1.EnvVar{},
		},
		{
			envs:     []k8sv1.EnvVar{{Name: "FOO"}, {Name: "BAR"}},
			toRemove: []string{"FOO"},
			expected: []k8sv1.EnvVar{{Name: "BAR"}},
		},
		{
			envs:     []k8sv1.EnvVar{{Name: "FOO"}, {Name: "BAR"}},
			toRemove: []string{"BAR"},
			expected: []k8sv1.EnvVar{{Name: "FOO"}},
		},
		{
			envs:     []k8sv1.EnvVar{{Name: "FOO"}, {Name: "BAR"}},
			toRemove: []string{"FOO", "BAR"},
			expected: []k8sv1.EnvVar{},
		},
	}

	for _, tc := range testCases {
		d := removeEnvVarsWithSecretsFromDeploy(deploySpec(tc.envs), tc.toRemove)
		cj := removeEnvVarsWithSecretsFromCronJob(cronJobSpec(tc.envs), tc.toRemove)
		cns := []k8sv1.Container{
			d.Spec.Template.Spec.Containers[0],
			cj.Spec.JobTemplate.Spec.Template.Spec.Containers[0],
		}
		for _, cn := range cns {
			if len(tc.expected) == 0 && len(cn.Env) != 0 {
				t.Errorf("expected no envs, there are some: %v - %v - %v", cn.Env, tc.toRemove, tc.envs)
			}
			for i := range tc.expected {
				actual := cn.Env[i]
				if actual != tc.expected[i] {
					t.Errorf("expected %s, got %s", tc.expected[i], actual)
				}
			}
		}
	}
}

func deploySpecWithVolumes(vols []k8sv1.Volume, volMounts []k8sv1.VolumeMount) *v1beta2.Deployment {
	d := deploySpec([]k8sv1.EnvVar{})
	d.Spec.Template.Spec.Volumes = vols
	d.Spec.Template.Spec.Containers[0].VolumeMounts = volMounts
	return d
}

func cronJobSpecWithVolumes(vols []k8sv1.Volume, volMounts []k8sv1.VolumeMount) *v1beta1.CronJob {
	cj := cronJobSpec([]k8sv1.EnvVar{})
	cj.Spec.JobTemplate.Spec.Template.Spec.Volumes = vols
	cj.Spec.JobTemplate.Spec.Template.Spec.Containers[0].VolumeMounts = volMounts
	return cj
}

func testRemoveVolumesWithSecrets(fromDeploy bool) func(*testing.T) {
	return func(t *testing.T) {
		var testCases = []struct {
			Volumes           []k8sv1.Volume
			VolumeMounts      []k8sv1.VolumeMount
			toRemove          []string
			expectedVols      []k8sv1.Volume
			expectedVolMounts []k8sv1.VolumeMount
		}{
			{
				Volumes: []k8sv1.Volume{{
					Name: spec.AppSecretName,
					VolumeSource: k8sv1.VolumeSource{
						Secret: &k8sv1.SecretVolumeSource{
							Items: []k8sv1.KeyToPath{{Key: "FOO"}},
						},
					},
				}},
				VolumeMounts: []k8sv1.VolumeMount{{
					Name: spec.AppSecretName,
				}},
				toRemove:          []string{"FOO"},
				expectedVols:      []k8sv1.Volume{},
				expectedVolMounts: []k8sv1.VolumeMount{},
			},
			{
				Volumes: []k8sv1.Volume{{
					Name: spec.AppSecretName,
					VolumeSource: k8sv1.VolumeSource{
						Secret: &k8sv1.SecretVolumeSource{
							Items: []k8sv1.KeyToPath{{Key: "FOO"}, {Key: "BAR"}},
						},
					},
				}},
				VolumeMounts: []k8sv1.VolumeMount{{
					Name: spec.AppSecretName,
				}},
				toRemove: []string{"FOO"},
				expectedVols: []k8sv1.Volume{{
					Name: spec.AppSecretName,
					VolumeSource: k8sv1.VolumeSource{
						Secret: &k8sv1.SecretVolumeSource{
							Items: []k8sv1.KeyToPath{{Key: "BAR"}},
						},
					},
				}},
				expectedVolMounts: []k8sv1.VolumeMount{{
					Name: spec.AppSecretName,
				}},
			},
			{
				Volumes: []k8sv1.Volume{{
					Name: spec.AppSecretName,
					VolumeSource: k8sv1.VolumeSource{
						Secret: &k8sv1.SecretVolumeSource{
							Items: []k8sv1.KeyToPath{{Key: "FOO"}, {Key: "BAR"}},
						},
					},
				}},
				VolumeMounts: []k8sv1.VolumeMount{{
					Name: spec.AppSecretName,
				}},
				toRemove:          []string{"FOO", "BAR"},
				expectedVols:      []k8sv1.Volume{},
				expectedVolMounts: []k8sv1.VolumeMount{},
			},
		}

		for _, tc := range testCases {
			var spec k8sv1.PodSpec
			if fromDeploy {
				d := deploySpecWithVolumes(tc.Volumes, tc.VolumeMounts)
				d = removeVolumesWithSecretsFromDeploy(d, tc.toRemove)
				spec = d.Spec.Template.Spec
			} else {
				cj := cronJobSpecWithVolumes(tc.Volumes, tc.VolumeMounts)
				cj = removeVolumesWithSecretsFromCronJob(cj, tc.toRemove)
				spec = cj.Spec.JobTemplate.Spec.Template.Spec
			}

			if actual := len(spec.Volumes); actual != len(tc.expectedVols) {
				t.Fatalf("expected %d, got %d vols", len(tc.expectedVols), actual)
			}
			if len(tc.expectedVols) > 0 {
				for i, item := range tc.expectedVols[0].Secret.Items {
					actual := spec.Volumes[0].Secret.Items[i].Key
					if actual != item.Key {
						t.Errorf("expected %s, got %s", item.Key, actual)
					}
				}
			}

			if actual := len(spec.Containers[0].VolumeMounts); actual != len(tc.expectedVolMounts) {
				t.Errorf("expected %d, got %d vol mounts", len(tc.expectedVolMounts), actual)
			}
		}
	}
}

func TestRemoveSecretVols(t *testing.T) {
	t.Run("TestRemoveVolumesWithSecretsFromDeploy", testRemoveVolumesWithSecrets(true))
	t.Run("TestRemoveVolumesWithSecretsFromCronJob", testRemoveVolumesWithSecrets(false))
}

func TestClientCreateNamespace(t *testing.T) {
	a := &app.App{Name: "test"}
	cli := &Client{testing: true}

	if err := cli.CreateNamespace(a, "test"); err != nil {
		t.Fatal("got unexpected error:", err)
	}

	ns, err := cli.fake.CoreV1().Namespaces().Get("test", metav1.GetOptions{})
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	an := ns.Annotations[app.TeresaLastUser]
	if an != "test" {
		t.Errorf("got %s; want test", an)
	}
}

func TestExposeDeploy(t *testing.T) {
	cli := &Client{
		testing: true,
		ingress: true,
	}
	var testCases = []struct {
		appName         string
		svcType         string
		vHosts          []string
		reserveStaticIp bool
		expectedIngress bool
	}{
		{
			appName:         "teresa-lb-nostaticip",
			svcType:         "LoadBalancer",
			vHosts:          []string{"a"},
			reserveStaticIp: false,
			expectedIngress: false,
		}, {
			appName:         "teresa-lb-staticip",
			svcType:         "LoadBalancer",
			vHosts:          []string{"a"},
			reserveStaticIp: true,
			expectedIngress: false,
		}, {
			appName:         "teresa-lb-vhost-staticip",
			svcType:         "LoadBalancer",
			vHosts:          []string{"foobar.luizalabs.com"},
			reserveStaticIp: true,
			expectedIngress: false,
		}, {
			appName:         "teresa-nodeport-vhost-staticip",
			svcType:         "NodePort",
			vHosts:          []string{"foobar.luizalabs.com"},
			reserveStaticIp: true,
			expectedIngress: true,
		}, {
			appName:         "teresa-nodeport-vhost-nostaticip",
			svcType:         "NodePort",
			vHosts:          []string{"foobar.luizalabs.com"},
			reserveStaticIp: false,
			expectedIngress: true,
		}, {
			appName:         "teresa-nodeport-novhost-nostaticip",
			svcType:         "NodePort",
			vHosts:          []string{},
			reserveStaticIp: false,
			expectedIngress: false,
		},
	}
	for _, tc := range testCases {
		a := &app.App{
			Name:            tc.appName,
			ReserveStaticIp: tc.reserveStaticIp,
		}
		if len(tc.vHosts) > 0 {
			a.VirtualHost = tc.vHosts[0]
		}
		if err := cli.CreateNamespace(a, tc.appName); err != nil {
			t.Fatal("got unexpected error:", err)
		}
		if err := cli.ExposeDeploy(
			tc.appName, tc.appName, tc.svcType, tc.appName, tc.vHosts,
			tc.reserveStaticIp, "nginx", ioutil.Discard,
		); err != nil {
			t.Fatal("got unexpected error:", err)
		}

		ingIface, err := cli.fake.ExtensionsV1beta1().
			Ingresses(tc.appName).Get(tc.appName, metav1.GetOptions{})
		if err != nil {
			if cli.IsNotFound(err) && tc.expectedIngress {
				t.Fatal("got unexpected error:", err)
			}
		}
		if ingIface == nil && tc.svcType == "NodePort" && len(tc.vHosts) > 0 {
			t.Errorf("got %v; want ingress", ingIface)
		}

		if ingIface != nil && tc.svcType != "NodePort" {
			t.Errorf("got ingress %v; want %v", ingIface, nil)
		}

		svcIface, err := cli.fake.CoreV1().
			Services(tc.appName).Get(tc.appName, metav1.GetOptions{})
		if err != nil {
			t.Fatal("got unexpected error:", err)
		}
		if svcIface == nil {
			t.Errorf("got nil, expected service")
		}
		if string(svcIface.Spec.Type) != tc.svcType {
			t.Errorf("got svcType %v, expected %v", string(svcIface.Spec.Type), tc.svcType)
		}

		if tc.expectedIngress {
			if err != nil {
				t.Fatal("got unexpected error:", err)
			}
			if tc.reserveStaticIp {
				if ingIface.Spec.Rules != nil {
					t.Errorf("got %v; want %v", ingIface.Spec.Rules, nil)
				}
				if ingIface.Spec.Backend == nil {
					t.Errorf("got %v; want backend", ingIface.Spec.Backend)
				}
			} else if len(tc.vHosts) > 0 {
				if tc.vHosts[0] == "" {
					t.Errorf("got %v; want %v", ingIface, nil)
				}
				if ingIface.Spec.Rules == nil {
					t.Errorf("got %v; want rules", ingIface.Spec.Rules)
				}
				if ingIface.Spec.Backend != nil {
					t.Errorf("got %v; want %v", ingIface.Spec.Backend, nil)
				}
			} else {
				t.Errorf("got %v; want %v", ingIface, nil)
			}
		}
	}
}

func TestHasIngressShouldNotDuplicate(t *testing.T) {
	appName := "teresa"
	cli := &Client{testing: true}
	a := &app.App{
		Name:            appName,
		ReserveStaticIp: true,
	}
	if err := cli.CreateNamespace(a, appName); err != nil {
		t.Fatal("got unexpected error:", err)
	}
	// This is not the same standard teresa uses to create an ingress
	// Teresa does not append -ingress to the end
	if err := cli.createIngress(
		appName,
		fmt.Sprintf("%s-ingress", appName),
		[]string{}, false, "nginx",
	); err != nil {
		t.Fatal("got unexpected error:", err)
	}
	hasIngress, err := cli.HasIngress(appName, appName)
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if hasIngress != true {
		t.Fatal("expected hasIngress to be true, got false")
	}
}
