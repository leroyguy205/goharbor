// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/vmware/harbor/src/replication"
)

//Trigger is replication launching approach definition
type Trigger struct {
	//The name of the trigger
	Kind string `json:"kind"`

	//The parameters with json text format required by the trigger
	Param string `json:"param"`
}

// Valid ...
func (t *Trigger) Valid(v *validation.Validation) {
	if !(t.Kind == replication.TriggerKindImmediate ||
		t.Kind == replication.TriggerKindManual ||
		t.Kind == replication.TriggerKindSchedule) {
		v.SetError("kind", fmt.Sprintf("invalid trigger kind: %s", t.Kind))
	}
}
