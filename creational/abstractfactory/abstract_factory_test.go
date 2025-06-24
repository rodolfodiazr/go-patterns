package abstractfactory

import (
	"testing"
)

type MockBucket struct {
	name         string
	fileUploaded bool
}

func (m *MockBucket) Upload(filename string) {
	m.fileUploaded = filename != ""
}

type MockComputeInstance struct {
	id              string
	instanceStarted bool
}

func (m *MockComputeInstance) Start() {
	m.instanceStarted = true
}

type MockCloudProvider struct {
	bucket  *MockBucket
	compute *MockComputeInstance
}

func (m *MockCloudProvider) CreateBucket(name string) Bucket {
	m.bucket.name = name
	return m.bucket
}

func (m *MockCloudProvider) CreateComputeInstance(id string) ComputeInstance {
	m.compute.id = id
	return m.compute
}

func Test_ProvisionInfrastructure_(t *testing.T) {
	bucket := &MockBucket{}
	compute := &MockComputeInstance{}
	provisionInfrastructure(&MockCloudProvider{
		bucket:  bucket,
		compute: compute,
	})

	if bucket.name != "logs" {
		t.Errorf("expected bucket name 'logs', got %s", bucket.name)
	}

	if !bucket.fileUploaded {
		t.Errorf("file is expected to be uploaded.")
	}

	if compute.id != "web-1" {
		t.Errorf("expected compute instance ID 'web-1', got %s", compute.id)
	}

	if !compute.instanceStarted {
		t.Errorf("instance is expected to be started.")
	}
}
