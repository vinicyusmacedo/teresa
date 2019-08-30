package storage

import (
	"context"
	"fmt"
	"io"
	"strings"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

type Blob interface {
	OpenBucket(context.Context, string) (*blob.Bucket, error)
}

type BlobHandler struct{}

func (b *BlobHandler) OpenBucket(ctx context.Context, urlstr string) (*blob.Bucket, error) {
	return blob.OpenBucket(ctx, urlstr)
}

type GCSClient interface {
	NewWriter(context.Context, string, *blob.WriterOptions) (*blob.Writer, error)
	List(*blob.ListOptions) *blob.ListIterator
	Delete(context.Context, string) error
}

type GCS struct {
	client     GCSClient
	blob       Blob
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
	ctx := context.Background()
	w, err := g.client.NewWriter(ctx, path, nil)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = io.Copy(w, file)
	if err != nil {
		return err
	}
	return nil
}

func (g *GCS) Type() string {
	return string(GCSType)
}

func (g *GCS) PodEnvVars() map[string]string {
	return make(map[string]string)
}

func (g *GCS) List(path string) ([]*Object, error) {
	ctx := context.Background()
	iter, err := g.listBucket(ctx, path)
	if err != nil {
		return nil, err
	}
	out := []*Object{}
	m := make(map[string]bool)
	for {
		obj, err := iter.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		name := strings.TrimPrefix(obj.Key, path)
		name = strings.Split(name, "/")[0]
		if _, found := m[name]; !found {
			m[name] = true
			out = append(out, &Object{Name: name})
		}
	}
	return out, nil
}

func (g *GCS) Delete(path string) error {
	ctx := context.Background()
	iter, err := g.listBucket(ctx, path)
	if err != nil {
		return err
	}
	for {
		obj, err := iter.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		g.client.Delete(ctx, obj.Key)
	}
	return nil
}

func (g *GCS) listBucket(ctx context.Context, path string) (*blob.ListIterator, error) {
	lo := &blob.ListOptions{
		Prefix: path,
	}
	return g.client.List(lo), nil
}

func newGCS(conf *Config) (*GCS, error) {
	gt := &GCS{
		Bucket:     conf.AwsBucket,
		GCSKeyFile: conf.GCSKeyFile,
		blob:       &BlobHandler{},
	}
	if conf.openBucket {
		ctx := context.Background()
		client, err := gt.blob.OpenBucket(ctx, fmt.Sprintf("gs://%s", gt.Bucket))
		if err != nil {
			return nil, err
		}
		gt.client = client
	}
	return gt, nil
}
