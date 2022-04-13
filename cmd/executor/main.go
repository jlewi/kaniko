/*
Copyright 2018 Google LLC

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
	"fmt"
	"os"

	"github.com/GoogleContainerTools/kaniko/cmd/executor/cmd"
	//"github.com/google/slowjam/pkg/stacklog"
)

func main() {
	// s := stacklog.MustStartFromEnv("STACKLOG_PATH")
	//defer s.Stop()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Printf("calling os.exit")
		os.Exit(1)
	}
	fmt.Printf("main is exiting")
}
