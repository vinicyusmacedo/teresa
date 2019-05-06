package cloudprovider

import (
	"testing"

	"github.com/luizalabs/teresa/pkg/server/teresa_errors"
	"github.com/pkg/errors"
)

func TestgceCreateOrUpdateSSLSuccess(t *testing.T) {
	ops := &gceOperations{&FakeK8sOperations{}}

	if err := ops.CreateOrUpdateSSL("teresa", "cert", 443); err != nil {
		t.Errorf("got %v; want no error", err)
	}
}

func TestgceCreateOrUpdateSSLFail(t *testing.T) {
	k8s := &FakeK8sOperations{SetServiceAnnotationsErr: errors.New("test")}
	ops := &gceOperations{k8s}

	e := teresa_errors.ErrInternalServerError
	if err := ops.CreateOrUpdateSSL("teresa", "cert", 443); teresa_errors.Get(err) != e {
		t.Errorf("got %v; want %v", teresa_errors.Get(err), e)
	}
}

func TestgceSSLInfoSuccess(t *testing.T) {
	ops := &gceOperations{&FakeK8sOperations{}}

	if _, err := ops.SSLInfo("teresa"); err != nil {
		t.Errorf("got %v; want no error", err)
	}
}

func TestgceSSLInfoFail(t *testing.T) {
	k8s := &FakeK8sOperations{ServiceAnnotationsErr: errors.New("test")}
	ops := &gceOperations{k8s}

	e := teresa_errors.ErrInternalServerError
	if _, err := ops.SSLInfo("teresa"); teresa_errors.Get(err) != e {
		t.Errorf("got %v; want %v", teresa_errors.Get(err), e)
	}
}

func TestgceCreateOrUpdateSSLErrNotImplemented(t *testing.T) {
	ops := &gceOperations{&FakeK8sOperations{HasIngressValue: true}}

	if err := ops.CreateOrUpdateSSL("teresa", "cert", 443); err != ErrNotImplementedOnIngress {
		t.Errorf("got %v; want %v", err, ErrNotImplementedOnIngress)
	}
}

func TestgceCreateOrUpdateSSLIngressErr(t *testing.T) {
	want := errors.New("test")
	ops := &gceOperations{&FakeK8sOperations{HasIngressErr: want}}

	if err := ops.CreateOrUpdateSSL("teresa", "cert", 443); err != want {
		t.Errorf("got %v; want %v", err, want)
	}
}

// TODO

func TestgceCreateOrUpdateStaticIPSuccess(t *testing.T) {
	ops := &gceOperations{&FakeK8sOperations{}}

	if err := ops.CreateOrUpdateStaticIP("teresa", "teresa-ingress"); err != nil {
		t.Errorf("got %v; want no error", err)
	}
}

func TestgceCreateOrUpdateStaticIPFail(t *testing.T) {
	k8s := &FakeK8sOperations{SetServiceAnnotationsErr: errors.New("test")}
	ops := &gceOperations{k8s}

	e := teresa_errors.ErrInternalServerError
	if err := ops.CreateOrUpdateStaticIP("teresa", "teresa-ingress"); teresa_errors.Get(err) != e {
		t.Errorf("got %v; want %v", teresa_errors.Get(err), e)
	}
}

func TestgceStaticIPInfoSuccess(t *testing.T) {
	ops := &gceOperations{&FakeK8sOperations{}}

	if _, err := ops.StaticIPInfo("teresa"); err != nil {
		t.Errorf("got %v; want no error", err)
	}
}

func TestgceStaticIPInfoFail(t *testing.T) {
	k8s := &FakeK8sOperations{ServiceAnnotationsErr: errors.New("test")}
	ops := &gceOperations{k8s}

	e := teresa_errors.ErrInternalServerError
	if _, err := ops.StaticIPInfo("teresa"); teresa_errors.Get(err) != e {
		t.Errorf("got %v; want %v", teresa_errors.Get(err), e)
	}
}

func TestgceCreateOrUpdateStaticIPErrNotImplemented(t *testing.T) {
	ops := &gceOperations{&FakeK8sOperations{HasIngressValue: true}}

	if err := ops.CreateOrUpdateStaticIP("teresa", "teresa-ingress"); err != ErrNotImplementedOnIngress {
		t.Errorf("got %v; want %v", err, ErrNotImplementedOnIngress)
	}
}

func TestgceCreateOrUpdateStaticIPIngressErr(t *testing.T) {
	want := errors.New("test")
	ops := &gceOperations{&FakeK8sOperations{HasIngressErr: want}}

	if err := ops.CreateOrUpdateStaticIP("teresa", "teresa-ingress"); err != want {
		t.Errorf("got %v; want %v", err, want)
	}
}
