/*
Copyright 2023 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"os"

	"k8s.io/utils/pointer"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	log "github.com/sirupsen/logrus"

	"github.com/deckhouse/deckhouse/go_lib/cloud-data/apis/v1alpha1"
)

type Discoverer struct {
	logger *log.Entry
	region string
}

func NewDiscoverer(logger *log.Entry) *Discoverer {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		logger.Fatal("AWS_REGION not found")
	}

	return &Discoverer{
		logger: logger,
		region: region,
	}
}

func (d *Discoverer) InstanceTypes(_ context.Context) ([]v1alpha1.InstanceType, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(d.region),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize new session: %v", err)
	}

	ec2Client := ec2.New(sess)
	res := make([]v1alpha1.InstanceType, 0)

	var token *string

	for {
		out, err := ec2Client.DescribeInstanceTypes(&ec2.DescribeInstanceTypesInput{
			MaxResults: pointer.Int64(100),
			NextToken:  token,
		})
		if err != nil {
			return nil, err
		}

		for _, ins := range out.InstanceTypes {
			if ins.InstanceType == nil {
				return nil, fmt.Errorf("instance type is nil")
			}

			name := *ins.InstanceType

			if ins.VCpuInfo == nil {
				return nil, fmt.Errorf("VCpuInfo is nil for %s", name)
			}

			if ins.VCpuInfo.DefaultVCpus == nil {
				return nil, fmt.Errorf("VCpuInfo.DefaultVCpus is nil for %s", name)
			}

			if ins.MemoryInfo == nil {
				return nil, fmt.Errorf("MemoryInfo is nil for %s", name)
			}

			if ins.MemoryInfo.SizeInMiB == nil {
				return nil, fmt.Errorf("MemoryInfo.SizeInMiB is nil for %s", name)
			}

			res = append(res, v1alpha1.InstanceType{
				Name:   name,
				CPU:    *ins.VCpuInfo.DefaultVCpus,
				Memory: *ins.MemoryInfo.SizeInMiB,
			})
		}

		if out.NextToken == nil || *out.NextToken == "" {
			break
		} else {
			token = out.NextToken
		}
	}

	return res, nil
}
