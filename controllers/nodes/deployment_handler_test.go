/*
Copyright 2022 Mondoo, Inc.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package nodes

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	mondoov1alpha2 "go.mondoo.com/mondoo-operator/api/v1alpha2"
	"go.mondoo.com/mondoo-operator/pkg/constants"
	"go.mondoo.com/mondoo-operator/pkg/mondooclient"
	"go.mondoo.com/mondoo-operator/pkg/utils/mondoo"
	fakeMondoo "go.mondoo.com/mondoo-operator/pkg/utils/mondoo/fake"
	"go.mondoo.com/mondoo-operator/tests/framework/utils"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const testNamespace = "mondoo-operator"

type DeploymentHandlerSuite struct {
	suite.Suite
	ctx                    context.Context
	scheme                 *runtime.Scheme
	containerImageResolver mondoo.ContainerImageResolver

	auditConfig       mondoov1alpha2.MondooAuditConfig
	fakeClientBuilder *fake.ClientBuilder
}

func (s *DeploymentHandlerSuite) SetupSuite() {
	s.ctx = context.Background()
	s.scheme = clientgoscheme.Scheme
	s.Require().NoError(mondoov1alpha2.AddToScheme(s.scheme))
	s.containerImageResolver = fakeMondoo.NewNoOpContainerImageResolver()
}

func (s *DeploymentHandlerSuite) BeforeTest(suiteName, testName string) {
	s.auditConfig = utils.DefaultAuditConfig(testNamespace, false, false, true, false)
	s.fakeClientBuilder = fake.NewClientBuilder()
	s.seedKubeSystemNamespace()
}

func (s *DeploymentHandlerSuite) TestReconcile_CreateConfigMap() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	for _, node := range nodes.Items {
		expected, err := ConfigMap(node, "", testClusterUID, s.auditConfig)
		s.Require().NoError(err, "unexpected error while generating ConfigMap")
		s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

		// Set some fields that the kube client sets
		gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
		s.NoError(err)
		expected.SetGroupVersionKind(gvk)
		expected.ResourceVersion = "1"

		created := &corev1.ConfigMap{}
		created.Name = expected.Name
		created.Namespace = expected.Namespace
		s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

		s.Equal(expected, created)
	}
}

func (s *DeploymentHandlerSuite) TestReconcile_CreateConfigMapWithIntegrationMRN() {
	const testIntegrationMRN = "//test-integration-MRN"

	s.seedNodes()

	sa, err := json.Marshal(mondooclient.ServiceAccountCredentials{Mrn: "test-mrn"})
	s.NoError(err)

	s.auditConfig.Spec.ConsoleIntegration.Enable = true
	credsWithIntegration := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      utils.MondooClientSecret,
			Namespace: testNamespace,
		},
		Data: map[string][]byte{
			constants.MondooCredsSecretIntegrationMRNKey: []byte(testIntegrationMRN),
			constants.MondooCredsSecretServiceAccountKey: sa,
		},
	}
	s.fakeClientBuilder = s.fakeClientBuilder.WithObjects(credsWithIntegration)

	d := s.createDeploymentHandler()

	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	for _, node := range nodes.Items {
		expected, err := ConfigMap(node, testIntegrationMRN, testClusterUID, s.auditConfig)
		s.Require().NoError(err, "unexpected error while generating ConfigMap")
		s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

		// Set some fields that the kube client sets
		gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
		s.NoError(err)
		expected.SetGroupVersionKind(gvk)
		expected.ResourceVersion = "1"

		created := &corev1.ConfigMap{}
		created.Name = expected.Name
		created.Namespace = expected.Namespace
		s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

		s.Equal(expected, created)
	}
}

func (s *DeploymentHandlerSuite) TestReconcile_UpdateConfigMap() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	for _, node := range nodes.Items {
		existing, err := ConfigMap(node, "", testClusterUID, s.auditConfig)
		s.Require().NoError(err, "unexpected error while generating ConfigMap")
		existing.Data["inventory"] = ""
		s.NoError(d.KubeClient.Create(s.ctx, existing))
	}

	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	for _, node := range nodes.Items {
		expected, err := ConfigMap(node, "", testClusterUID, s.auditConfig)
		s.Require().NoError(err, "unexpected error while generating ConfigMap")
		s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

		// Set some fields that the kube client sets
		gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
		s.NoError(err)
		expected.SetGroupVersionKind(gvk)
		expected.ResourceVersion = "2"

		created := &corev1.ConfigMap{}
		created.Name = expected.Name
		created.Namespace = expected.Namespace
		s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

		s.Equal(expected, created)
	}
}

func (s *DeploymentHandlerSuite) TestReconcile_CleanConfigMapsForDeletedNodes() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	// Reconcile to create the initial cron jobs
	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	// Delete one node
	s.NoError(d.KubeClient.Delete(s.ctx, &nodes.Items[1]))

	// Reconcile again to delete the cron job for the deleted node
	result, err = d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	configMaps := &corev1.ConfigMapList{}
	s.NoError(d.KubeClient.List(s.ctx, configMaps))

	s.Equal(1, len(configMaps.Items))

	expected, err := ConfigMap(nodes.Items[0], "", testClusterUID, s.auditConfig)
	s.Require().NoError(err, "unexpected error while generating ConfigMap")
	s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

	// Set some fields that the kube client sets
	gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
	s.NoError(err)
	expected.SetGroupVersionKind(gvk)
	expected.ResourceVersion = "1"

	created := &corev1.ConfigMap{}
	created.Name = expected.Name
	created.Namespace = expected.Namespace
	s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

	s.Equal(expected, created)
}

func (s *DeploymentHandlerSuite) TestReconcile_CreateCronJobs() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	image, err := s.containerImageResolver.CnspecImage(
		s.auditConfig.Spec.Scanner.Image.Name, s.auditConfig.Spec.Scanner.Image.Tag, false)
	s.NoError(err)

	for _, n := range nodes.Items {
		expected := CronJob(image, n, s.auditConfig, false)
		s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

		// Set some fields that the kube client sets
		gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
		s.NoError(err)
		expected.SetGroupVersionKind(gvk)
		expected.ResourceVersion = "1"

		created := &batchv1.CronJob{}
		created.Name = expected.Name
		created.Namespace = expected.Namespace
		s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

		s.Equal(expected, created)
	}

	operatorImage, err := s.containerImageResolver.MondooOperatorImage(
		s.auditConfig.Spec.Scanner.Image.Name, s.auditConfig.Spec.Scanner.Image.Tag, false)
	s.NoError(err)

	// Verify node garbage collection cronjob
	expected := GarbageCollectCronJob(operatorImage, "abcdefg", s.auditConfig)
	s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

	// Set some fields that the kube client sets
	gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
	s.NoError(err)
	expected.SetGroupVersionKind(gvk)
	expected.ResourceVersion = "1"

	created := &batchv1.CronJob{}
	created.Name = expected.Name
	created.Namespace = expected.Namespace
	s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))
	s.Equal(expected, created)
}

func (s *DeploymentHandlerSuite) TestReconcile_UpdateCronJobs() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	image, err := s.containerImageResolver.CnspecImage(
		s.auditConfig.Spec.Scanner.Image.Name, s.auditConfig.Spec.Scanner.Image.Tag, false)
	s.NoError(err)

	// Make sure a cron job exists for one of the nodes
	cronJob := CronJob(image, nodes.Items[1], s.auditConfig, false)
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Command = []string{"test-command"}
	s.NoError(d.KubeClient.Create(s.ctx, cronJob))

	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	for i, n := range nodes.Items {
		expected := CronJob(image, n, s.auditConfig, false)
		s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

		// Set some fields that the kube client sets
		gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
		s.NoError(err)
		expected.SetGroupVersionKind(gvk)

		// The second node has an updated cron job so resource version is +1
		expected.ResourceVersion = fmt.Sprintf("%d", 1+i)

		created := &batchv1.CronJob{}
		created.Name = expected.Name
		created.Namespace = expected.Namespace
		s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

		s.Equal(expected, created)
	}
}

func (s *DeploymentHandlerSuite) TestReconcile_CleanCronJobsForDeletedNodes() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	// Reconcile to create the initial cron jobs
	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	nodes := &corev1.NodeList{}
	s.NoError(d.KubeClient.List(s.ctx, nodes))

	// Delete one node
	s.NoError(d.KubeClient.Delete(s.ctx, &nodes.Items[1]))

	// Reconcile again to delete the cron job for the deleted node
	result, err = d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	image, err := s.containerImageResolver.CnspecImage(
		s.auditConfig.Spec.Scanner.Image.Name, s.auditConfig.Spec.Scanner.Image.Tag, false)
	s.NoError(err)

	listOpts := &client.ListOptions{
		Namespace:     s.auditConfig.Namespace,
		LabelSelector: labels.SelectorFromSet(CronJobLabels(s.auditConfig)),
	}
	cronJobs := &batchv1.CronJobList{}
	s.NoError(d.KubeClient.List(s.ctx, cronJobs, listOpts))

	s.Equal(1, len(cronJobs.Items))

	expected := CronJob(image, nodes.Items[0], s.auditConfig, false)
	s.NoError(ctrl.SetControllerReference(&s.auditConfig, expected, d.KubeClient.Scheme()))

	// Set some fields that the kube client sets
	gvk, err := apiutil.GVKForObject(expected, d.KubeClient.Scheme())
	s.NoError(err)
	expected.SetGroupVersionKind(gvk)
	expected.ResourceVersion = "1"

	created := &batchv1.CronJob{}
	created.Name = expected.Name
	created.Namespace = expected.Namespace
	s.NoError(d.KubeClient.Get(s.ctx, client.ObjectKeyFromObject(created), created))

	s.Equal(expected, created)
}

func (s *DeploymentHandlerSuite) TestReconcile_NodeScanningStatus() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	// Reconcile to create all resources
	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	// Verify the node scanning status is set to available
	s.Equal(1, len(d.Mondoo.Status.Conditions))
	condition := d.Mondoo.Status.Conditions[0]
	s.Equal("Node Scanning is available", condition.Message)
	s.Equal("NodeScanningAvailable", condition.Reason)
	s.Equal(corev1.ConditionFalse, condition.Status)

	listOpts := &client.ListOptions{
		Namespace:     s.auditConfig.Namespace,
		LabelSelector: labels.SelectorFromSet(CronJobLabels(s.auditConfig)),
	}
	cronJobs := &batchv1.CronJobList{}
	s.NoError(d.KubeClient.List(s.ctx, cronJobs, listOpts))

	now := time.Now()
	metaNow := metav1.NewTime(now)
	metaHourAgo := metav1.NewTime(now.Add(-1 * time.Hour))
	cronJobs.Items[0].Status.LastScheduleTime = &metaNow
	cronJobs.Items[0].Status.LastSuccessfulTime = &metaHourAgo

	s.NoError(d.KubeClient.Update(s.ctx, &cronJobs.Items[0]))

	// Reconcile to update the audit config status
	result, err = d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	// Verify the node scanning status is set to unavailable
	condition = d.Mondoo.Status.Conditions[0]
	s.Equal("Node Scanning is unavailable", condition.Message)
	s.Equal("NodeScanningUnavailable", condition.Reason)
	s.Equal(corev1.ConditionTrue, condition.Status)

	// Make the jobs successful again
	cronJobs.Items[0].Status.LastScheduleTime = nil
	cronJobs.Items[0].Status.LastSuccessfulTime = nil
	s.NoError(d.KubeClient.Update(s.ctx, &cronJobs.Items[0]))

	// Reconcile to update the audit config status
	result, err = d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	// Verify the node scanning status is set to available
	condition = d.Mondoo.Status.Conditions[0]
	s.Equal("Node Scanning is available", condition.Message)
	s.Equal("NodeScanningAvailable", condition.Reason)
	s.Equal(corev1.ConditionFalse, condition.Status)

	d.Mondoo.Spec.Nodes.Enable = false

	// Reconcile to update the audit config status
	result, err = d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	// Verify the node scanning status is set to disabled
	condition = d.Mondoo.Status.Conditions[0]
	s.Equal("Node Scanning is disabled", condition.Message)
	s.Equal("NodeScanningDisabled", condition.Reason)
	s.Equal(corev1.ConditionFalse, condition.Status)
}

func (s *DeploymentHandlerSuite) TestReconcile_DisableNodeScanning() {
	s.seedNodes()
	d := s.createDeploymentHandler()

	// Reconcile to create all resources
	result, err := d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	// Reconcile again to delete the resources
	d.Mondoo.Spec.Nodes.Enable = false
	result, err = d.Reconcile(s.ctx)
	s.NoError(err)
	s.True(result.IsZero())

	configMaps := &corev1.ConfigMapList{}
	s.NoError(d.KubeClient.List(s.ctx, configMaps))
	s.Equal(0, len(configMaps.Items))

	cronJobs := &batchv1.CronJobList{}
	s.NoError(d.KubeClient.List(s.ctx, cronJobs))
	s.Equal(0, len(cronJobs.Items))
}

func (s *DeploymentHandlerSuite) createDeploymentHandler() DeploymentHandler {
	return DeploymentHandler{
		KubeClient:             s.fakeClientBuilder.Build(),
		Mondoo:                 &s.auditConfig,
		ContainerImageResolver: s.containerImageResolver,
		MondooOperatorConfig:   &mondoov1alpha2.MondooOperatorConfig{},
	}
}

func (s *DeploymentHandlerSuite) seedNodes() {
	master := &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{Name: "node01"},
		Spec: corev1.NodeSpec{
			Taints: []corev1.Taint{
				{Key: "node-role.kubernetes.io/master", Value: "true", Effect: corev1.TaintEffectNoExecute},
			},
		},
	}
	worker := &corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node02"}}
	s.fakeClientBuilder = s.fakeClientBuilder.WithObjects(master, worker)
}

func (s *DeploymentHandlerSuite) seedKubeSystemNamespace() {
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kube-system",
			UID:  testClusterUID,
		},
	}
	s.fakeClientBuilder = s.fakeClientBuilder.WithObjects(namespace)
}

func TestDeploymentHandlerSuite(t *testing.T) {
	suite.Run(t, new(DeploymentHandlerSuite))
}
