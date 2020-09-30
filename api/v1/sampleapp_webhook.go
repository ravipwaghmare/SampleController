/*

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

package v1

import (
	"errors"
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var sampleapplog = logf.Log.WithName("sampleapp-resource")

func (r *SampleApp) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-app-mydomain-v1-sampleapp,mutating=true,failurePolicy=fail,groups=app.mydomain,resources=sampleapps,verbs=create;update,versions=v1,name=msampleapp.kb.io

var _ webhook.Defaulter = &SampleApp{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SampleApp) Default() {
	sampleapplog.Info("default", "name", r.Name)
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-app-mydomain-v1-sampleapp,mutating=false,failurePolicy=fail,groups=app.mydomain,resources=sampleapps,versions=v1,name=vsampleapp.kb.io

var _ webhook.Validator = &SampleApp{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SampleApp) ValidateCreate() error {
	sampleapplog.Info("validate create", "name", r.Name)
	return r.validateUserInfo()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SampleApp) ValidateUpdate(old runtime.Object) error {
	sampleapplog.Info("validate update", "name", r.Name)
	return r.validateUserInfo()
}

func (r *SampleApp) validateUserInfo() error {
	re := regexp.MustCompile("^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$")
	if !re.MatchString(r.Spec.UserInfo.Email) {
		return errors.New("Invalid Email Address")
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SampleApp) ValidateDelete() error {
	sampleapplog.Info("validate delete", "name", r.Name)
	return nil
}
