package storage

import "io"

type GCS struct {
	Bucket     string
	GCSKeyFile string
}

func (g *GCS) K8sSecretName() string {
	return "s3-storage"
}

func (g *GCS) AccessData() map[string][]byte {
	return map[string][]byte{
		"builder-bucket": []byte(g.Bucket),
		"gcs-key-file":   []byte(g.GCSKeyFile),
	}
}

func (g *GCS) UploadFile(path string, file io.ReadSeeker) error {
	return nil
}

func (g *GCS) Type() string {
	return string(GCSType)
}

func (g *GCS) PodEnvVars() map[string]string {
	return make(map[string]string)
}

func (g *GCS) List(path string) ([]*Object, error) {
	return nil, nil
}

func (g *GCS) Delete(path string) error {
	return nil
}

func newGCS(conf *Config) *GCS {
	gt := &GCS{
		Bucket:     conf.AwsBucket,
		GCSKeyFile: conf.GCSKeyFile,
	}
	return gt
}
