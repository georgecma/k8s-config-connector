// Copyright 2024 Google LLC
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

// +tool:mockgcp-service
// http.host: modelarmor.{location}.rep.googleapis.com
// proto.service: google.cloud.modelarmor.v1.ModelArmor

package mockmodelarmor

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/modelarmor/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked ModelArmor service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	// todo: hardcode regional endpoint for now
	return []string{"modelarmor.us-central1.rep.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterModelArmorServer(grpcServer, &ModelArmorV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{}, pb.RegisterModelArmorHandler)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
