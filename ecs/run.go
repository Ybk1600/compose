/*
   Copyright 2020 Docker Compose CLI authors

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

package ecs

import (
	"context"

	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose-cli/api/compose"
	"github.com/docker/compose-cli/errdefs"
)

func (b *ecsAPIService) CreateOneOffContainer(ctx context.Context, project *types.Project, opts compose.RunOptions) (string, error) {
	return "", errdefs.ErrNotImplemented
}

func (b *ecsAPIService) Run(ctx context.Context, container string, detach bool) error {
	return errdefs.ErrNotImplemented
}
