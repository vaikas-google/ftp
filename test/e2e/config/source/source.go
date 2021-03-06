// +build e2e

package source

import (
  "context"
  "testing"

  "knative.dev/reconciler-test/pkg/environment"
  "knative.dev/reconciler-test/pkg/feature"
  "knative.dev/reconciler-test/pkg/manifest"
)

func init() {
  environment.RegisterPackage(manifest.ImagesLocalYaml()...)
}

func Install() feature.StepFn {
  return func(ctx context.Context, t *testing.T) {
    if _, err := manifest.InstallLocalYaml(ctx, map[string]interface{}{"producerCount": 5}); err != nil {
      t.Fatal(err)
    }
  }
}
