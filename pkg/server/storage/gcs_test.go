package storage

import (
	"context"
	"reflect"
	"testing"

	"gocloud.dev/blob"
)

type fakeBlob struct{}

func (f *fakeBlob) OpenBucket(context.Context, string) (*blob.Bucket, error) {
	return nil, nil
}

type fakeWriteCloser struct{}

func (f *fakeWriteCloser) Write([]byte) (int, error) {
	return 0, nil
}

func (f *fakeWriteCloser) Close() error {
	return nil
}

type fakeGCSClient struct{}

func (f *fakeGCSClient) NewWriter(context.Context, string, *blob.WriterOptions) (*blob.Writer, error) {
	return nil, nil
}

func (f *fakeGCSClient) List(*blob.ListOptions) *blob.ListIterator {
	return nil
}

func (f *fakeGCSClient) Delete(context.Context, string) error {
	return nil
}

func TestGCSK8sSecretName(t *testing.T) {
	gcs, err := newGCS(&Config{})
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}

	if name := gcs.K8sSecretName(); name != "s3-storage" {
		t.Errorf("expected s3-storage, got %s", name)
	}
}

func TestGCSType(t *testing.T) {
	gcs, err := newGCS(&Config{})
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}

	if tmp := gcs.Type(); tmp != "gcs" {
		t.Errorf("expected gcs, got %s", tmp)
	}
}

func TestGCSAccessData(t *testing.T) {
	conf := &Config{
		AwsBucket:  "bucket",
		GCSKeyFile: "credentialsjsoncontent",
	}
	gcs, err := newGCS(conf)
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}
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
	gcs, err := newGCS(&Config{})
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}
	gcs.client = &fakeGCSClient{}
	gcs.blob = &fakeBlob{}

	if err := gcs.UploadFile("/test", &fakeReadSeeker{}); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGCSDelete(t *testing.T) {
	gcs, err := newGCS(&Config{})
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}
	gcs.client = &fakeGCSClient{}
	gcs.blob = &fakeBlob{}

	if err := gcs.Delete("/test"); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
}

func TestGCSPodEnvVars(t *testing.T) {
	gcs, err := newGCS(&Config{})
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}
	ev := gcs.PodEnvVars()
	if len(ev) != 0 {
		t.Errorf("expected 0, got %d", len(ev))
	}
}

func TestGCSList(t *testing.T) {
	gcs, err := newGCS(&Config{})
	if err != nil {
		t.Errorf("error creating new gcs, %v", err)
	}
	gcs.client = &fakeGCSClient{}
	gcs.blob = &fakeBlob{}
	if _, err := gcs.List("some/path"); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
