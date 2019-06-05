package cloudprovider

import (
	"github.com/luizalabs/teresa/pkg/server/service"
	"github.com/luizalabs/teresa/pkg/server/teresa_errors"
)

const (
	gceSSLCertAnnotation  = "ingress.gce.kubernetes.io/pre-shared-cert"
	gceStaticIPAnnotation = "kubernetes.io/ingress.global-static-ip-name"
)

type gceOperations struct {
	k8s K8sOperations
}

func (ops *gceOperations) CreateOrUpdateSSL(appName, cert string, port int) error {
	hasIngress, err := ops.k8s.HasIngress(appName, appName)
	if err != nil {
		return err
	}
	if !hasIngress {
		return ErrNotImplementedOnLoadBalancer
	}
	anMap := map[string]string{
		gceSSLCertAnnotation: cert,
	}
	if err := ops.k8s.SetIngressAnnotations(appName, appName, anMap); err != nil {
		return teresa_errors.NewInternalServerError(err)
	}
	return nil
}

func (ops *gceOperations) CreateOrUpdateStaticIp(appName, addressName string) error {
	hasIngress, err := ops.k8s.HasIngress(appName, appName)
	if err != nil {
		return err
	}
	if !hasIngress {
		return ErrNotImplementedOnLoadBalancer
	}
	anMap := map[string]string{
		gceStaticIPAnnotation: addressName,
	}
	if err := ops.k8s.SetIngressAnnotations(appName, appName, anMap); err != nil {
		return teresa_errors.NewInternalServerError(err)
	}
	return nil
}

func (ops *gceOperations) SSLInfo(appName string) (*service.SSLInfo, error) {
	an, err := ops.k8s.IngressAnnotations(appName, appName)
	if err != nil {
		if ops.k8s.IsNotFound(err) {
			return nil, ErrIngressNotFound
		}
		return nil, teresa_errors.NewInternalServerError(err)
	}
	info := &service.SSLInfo{
		Cert: an[gceSSLCertAnnotation],
	}
	return info, nil
}

func reserveIP(project string, name string) error {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := compute.New(c)
	addr := &compute.Address{
		Name: name,
	}
	_, err = svc.GlobalAddresses.Insert(project, addr).Context(ctx).Do()
	return err
}

func (ops *gceOperations) Name() string {
	return "gce"
}
