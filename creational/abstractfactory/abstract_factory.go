package abstractfactory

import "fmt"

// Bucket is a product abstract.
type Bucket interface {
	Upload(file string)
}

// ComputeInstance is a product abstract.
type ComputeInstance interface {
	Start()
}

// CloudProvider is an abstract factory.
type CloudProvider interface {
	CreateBucket(name string) Bucket
	CreateComputeInstance(id string) ComputeInstance
}

// AWSBucket is an AWS concrete product.
type AWSBucket struct {
	name string
}

func (b *AWSBucket) Upload(file string) {
	fmt.Printf("Uploading %s to AWS S3 bucket %s\n", file, b.name)
}

// AWSComputeInstance is an AWS concrete product.
type AWSComputeInstance struct {
	id string
}

func (c *AWSComputeInstance) Start() {
	fmt.Printf("Starting AWS EC2 instance %s\n", c.id)
}

// AWSProvider is an AWS concrete abstract factory.
type AWSProvider struct{}

func (f *AWSProvider) CreateBucket(name string) Bucket {
	return &AWSBucket{name: name}
}

func (f *AWSProvider) CreateComputeInstance(id string) ComputeInstance {
	return &AWSComputeInstance{id: id}
}

// GCPBucket is a GPC concrete product.
type GCPBucket struct {
	name string
}

func (b *GCPBucket) Upload(file string) {
	fmt.Printf("Uploading %s to GCP Storage bucket %s\n", file, b.name)
}

// GCPBucket is a GPC concrete product.
type GCPComputeInstance struct {
	id string
}

func (c *GCPComputeInstance) Start() {
	fmt.Printf("Starting GCP Compute Engine instance %s\n", c.id)
}

// GCPProvider is a GPC concrete abstract factory.
type GCPProvider struct{}

func (f *GCPProvider) CreateBucket(name string) Bucket {
	return &GCPBucket{name: name}
}

func (f *GCPProvider) CreateComputeInstance(id string) ComputeInstance {
	return &GCPComputeInstance{id: id}
}

func provisionInfrastructure(factory CloudProvider) {
	bucket := factory.CreateBucket("logs")
	bucket.Upload("app.log")

	instance := factory.CreateComputeInstance("web-1")
	instance.Start()
}

func Run() {
	fmt.Println("Provisioning on AWS:")
	provisionInfrastructure(&AWSProvider{})

	fmt.Println("\nProvisioning on GCP:")
	provisionInfrastructure(&GCPProvider{})
}
