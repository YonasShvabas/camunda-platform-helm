// Copyright 2022 Camunda Services GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identity

import (
	"path/filepath"
	"strings"
	"testing"

	"camunda-platform-helm/charts/ccsm-helm/test/golden"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestGoldenIngressDefaultTemplate(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "ccsm-helm-test",
		Namespace:      "ccsm-helm-" + strings.ToLower(random.UniqueId()),
		GoldenFileName: "ingress",
		Templates:      []string{"charts/identity/templates/ingress.yaml"},
		SetValues:      map[string]string{"identity.ingress.enabled": "true"},
	})
}

func TestGoldenIngressAllEnabledTemplate(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../")
	require.NoError(t, err)

	suite.Run(t, &golden.TemplateGoldenTest{
		ChartPath:      chartPath,
		Release:        "ccsm-helm-test",
		Namespace:      "ccsm-helm-" + strings.ToLower(random.UniqueId()),
		GoldenFileName: "ingress-all-enabled",
		Templates:      []string{"charts/identity/templates/ingress.yaml"},
		SetValues: map[string]string{
			"identity.ingress.enabled":        "true",
			"identity.ingress.host":           "local",
			"identity.ingress.tls.enabled":    "true",
			"identity.ingress.tls.secretName": "my-secret",
		},
	})
}