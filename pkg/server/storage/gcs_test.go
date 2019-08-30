package storage

import (
	"os"
	"reflect"
	"testing"
)

func TestGCSK8sSecretName(t *testing.T) {
	gcs := newGCS(&Config{})

	if name := gcs.K8sSecretName(); name != "s3-storage" {
		t.Errorf("expected s3-storage, got %s", name)
	}
}

func TestGCSType(t *testing.T) {
	gcs := newGCS(&Config{})

	if tmp := gcs.Type(); tmp != "gcs" {
		t.Errorf("expected gcs, got %s", tmp)
	}
}

func TestGCSAccessData(t *testing.T) {
	conf := &Config{
		AwsBucket:  "bucket",
		GCSKeyFile: "credentialsjsoncontent",
	}
	gcs := newGCS(conf)
	ad := gcs.AccessData()
	var testCases = []struct {
		key   string
		field string
	}{
		{"builder-bucket", "AwsBucket"},
		{"gcs-key-file", "GCSKeyFile"},
	}

	for _, tc := range testCases {
		v := reflect.ValueOf(conf)
		expected := reflect.Indirect(v).FieldByName(tc.field).String()
		got := string(ad[tc.key])
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	}
}

func TestGCSUploadFile(t *testing.T) {
	gcs := newGCS(&Config{})
	//gcs.Client = &fakeGCSClient{}

	/*if err := gcs.UploadFile("/test", &fakeReadSeeker{}); err != nil {
		t.Errorf("expected no error, got %v", err)
	}*/

	// unmocked code below
	gcs.Bucket = "teresa-minio-minikube"
	fd, err := os.Open("storage.go")
	if err != nil {
		t.Errorf("error opening file %v", err)
	}
	if err := gcs.UploadFile("testupload/testfile", fd); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if err := gcs.UploadFile("testupload/testdelete", fd); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGCSDelete(t *testing.T) {
	gcs := newGCS(&Config{})
	//gcs.Client = &fakeGCSClient{}

	/*if err := gcs.Delete("/test"); err != nil {
		t.Errorf("expected no error, got %s", err)
	}*/

	// unmocked code below
	gcs.Bucket = "teresa-minio-minikube"
	if err := gcs.Delete("testupload/testdelete"); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	objects, err := gcs.List("testupload/testfile")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(objects) != 1 {
		t.Errorf("expected 1 object, got %v", objects)
	}
}

func TestGCSPodEnvVars(t *testing.T) {
	gcs := newGCS(&Config{})
	ev := gcs.PodEnvVars()
	if len(ev) != 0 {
		t.Errorf("expected 0, got %d", len(ev))
	}
}

func TestGCSList(t *testing.T) {
	gcs := newGCS(&Config{})
	//gcs.Client = &fakeGCSClient{}
	/*if _, err := gcs.List("some/path"); err != nil {
		t.Errorf("expected no error, got %v", err)
	}*/

	// unmocked code below
	gcs.Bucket = "teresa-minio-minikube"
	objects, err := gcs.List("testupload/")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(objects) != 2 {
		t.Errorf("expected 2 object, got %d", len(objects))
	}
}
