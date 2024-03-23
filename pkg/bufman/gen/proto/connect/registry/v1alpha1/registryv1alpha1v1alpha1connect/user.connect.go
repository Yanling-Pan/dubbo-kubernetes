// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: registry/v1alpha1/user.proto

package registryv1alpha1v1alpha1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

import (
	connect_go "github.com/bufbuild/connect-go"
)

import (
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/go/registry/v1alpha1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "bufman.dubbo.apache.org.registry.v1alpha1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceCreateUserProcedure is the fully-qualified name of the UserService's CreateUser RPC.
	UserServiceCreateUserProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/CreateUser"
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/GetUser"
	// UserServiceGetUserByUsernameProcedure is the fully-qualified name of the UserService's
	// GetUserByUsername RPC.
	UserServiceGetUserByUsernameProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/GetUserByUsername"
	// UserServiceListUsersProcedure is the fully-qualified name of the UserService's ListUsers RPC.
	UserServiceListUsersProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/ListUsers"
	// UserServiceListOrganizationUsersProcedure is the fully-qualified name of the UserService's
	// ListOrganizationUsers RPC.
	UserServiceListOrganizationUsersProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/ListOrganizationUsers"
	// UserServiceDeleteUserProcedure is the fully-qualified name of the UserService's DeleteUser RPC.
	UserServiceDeleteUserProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/DeleteUser"
	// UserServiceDeactivateUserProcedure is the fully-qualified name of the UserService's
	// DeactivateUser RPC.
	UserServiceDeactivateUserProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/DeactivateUser"
	// UserServiceUpdateUserServerRoleProcedure is the fully-qualified name of the UserService's
	// UpdateUserServerRole RPC.
	UserServiceUpdateUserServerRoleProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/UpdateUserServerRole"
	// UserServiceCountUsersProcedure is the fully-qualified name of the UserService's CountUsers RPC.
	UserServiceCountUsersProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/CountUsers"
	// UserServiceUpdateUserSettingsProcedure is the fully-qualified name of the UserService's
	// UpdateUserSettings RPC.
	UserServiceUpdateUserSettingsProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/UpdateUserSettings"
)

// UserServiceClient is a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.UserService service.
type UserServiceClient interface {
	// CreateUser creates a new user with the given username.
	CreateUser(context.Context, *connect_go.Request[v1alpha1.CreateUserRequest]) (*connect_go.Response[v1alpha1.CreateUserResponse], error)
	// GetUser gets a user by ID.
	GetUser(context.Context, *connect_go.Request[v1alpha1.GetUserRequest]) (*connect_go.Response[v1alpha1.GetUserResponse], error)
	// GetUserByUsername gets a user by username.
	GetUserByUsername(context.Context, *connect_go.Request[v1alpha1.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha1.GetUserByUsernameResponse], error)
	// ListUsers lists all users.
	ListUsers(context.Context, *connect_go.Request[v1alpha1.ListUsersRequest]) (*connect_go.Response[v1alpha1.ListUsersResponse], error)
	// ListOrganizationUsers lists all users for an organization.
	// TODO: #663 move this to organization service
	ListOrganizationUsers(context.Context, *connect_go.Request[v1alpha1.ListOrganizationUsersRequest]) (*connect_go.Response[v1alpha1.ListOrganizationUsersResponse], error)
	// DeleteUser deletes a user.
	DeleteUser(context.Context, *connect_go.Request[v1alpha1.DeleteUserRequest]) (*connect_go.Response[v1alpha1.DeleteUserResponse], error)
	// Deactivate user deactivates a user.
	DeactivateUser(context.Context, *connect_go.Request[v1alpha1.DeactivateUserRequest]) (*connect_go.Response[v1alpha1.DeactivateUserResponse], error)
	// UpdateUserServerRole update the role of an user in the server.
	UpdateUserServerRole(context.Context, *connect_go.Request[v1alpha1.UpdateUserServerRoleRequest]) (*connect_go.Response[v1alpha1.UpdateUserServerRoleResponse], error)
	// CountUsers returns the number of users in the server by the user state provided.
	CountUsers(context.Context, *connect_go.Request[v1alpha1.CountUsersRequest]) (*connect_go.Response[v1alpha1.CountUsersResponse], error)
	// UpdateUserSettings update the user settings including description.
	UpdateUserSettings(context.Context, *connect_go.Request[v1alpha1.UpdateUserSettingsRequest]) (*connect_go.Response[v1alpha1.UpdateUserSettingsResponse], error)
}

// NewUserServiceClient constructs a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.UserService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		createUser: connect_go.NewClient[v1alpha1.CreateUserRequest, v1alpha1.CreateUserResponse](
			httpClient,
			baseURL+UserServiceCreateUserProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
			connect_go.WithClientOptions(opts...),
		),
		getUser: connect_go.NewClient[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		getUserByUsername: connect_go.NewClient[v1alpha1.GetUserByUsernameRequest, v1alpha1.GetUserByUsernameResponse](
			httpClient,
			baseURL+UserServiceGetUserByUsernameProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listUsers: connect_go.NewClient[v1alpha1.ListUsersRequest, v1alpha1.ListUsersResponse](
			httpClient,
			baseURL+UserServiceListUsersProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listOrganizationUsers: connect_go.NewClient[v1alpha1.ListOrganizationUsersRequest, v1alpha1.ListOrganizationUsersResponse](
			httpClient,
			baseURL+UserServiceListOrganizationUsersProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		deleteUser: connect_go.NewClient[v1alpha1.DeleteUserRequest, v1alpha1.DeleteUserResponse](
			httpClient,
			baseURL+UserServiceDeleteUserProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
			connect_go.WithClientOptions(opts...),
		),
		deactivateUser: connect_go.NewClient[v1alpha1.DeactivateUserRequest, v1alpha1.DeactivateUserResponse](
			httpClient,
			baseURL+UserServiceDeactivateUserProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
			connect_go.WithClientOptions(opts...),
		),
		updateUserServerRole: connect_go.NewClient[v1alpha1.UpdateUserServerRoleRequest, v1alpha1.UpdateUserServerRoleResponse](
			httpClient,
			baseURL+UserServiceUpdateUserServerRoleProcedure,
			opts...,
		),
		countUsers: connect_go.NewClient[v1alpha1.CountUsersRequest, v1alpha1.CountUsersResponse](
			httpClient,
			baseURL+UserServiceCountUsersProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		updateUserSettings: connect_go.NewClient[v1alpha1.UpdateUserSettingsRequest, v1alpha1.UpdateUserSettingsResponse](
			httpClient,
			baseURL+UserServiceUpdateUserSettingsProcedure,
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	createUser            *connect_go.Client[v1alpha1.CreateUserRequest, v1alpha1.CreateUserResponse]
	getUser               *connect_go.Client[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse]
	getUserByUsername     *connect_go.Client[v1alpha1.GetUserByUsernameRequest, v1alpha1.GetUserByUsernameResponse]
	listUsers             *connect_go.Client[v1alpha1.ListUsersRequest, v1alpha1.ListUsersResponse]
	listOrganizationUsers *connect_go.Client[v1alpha1.ListOrganizationUsersRequest, v1alpha1.ListOrganizationUsersResponse]
	deleteUser            *connect_go.Client[v1alpha1.DeleteUserRequest, v1alpha1.DeleteUserResponse]
	deactivateUser        *connect_go.Client[v1alpha1.DeactivateUserRequest, v1alpha1.DeactivateUserResponse]
	updateUserServerRole  *connect_go.Client[v1alpha1.UpdateUserServerRoleRequest, v1alpha1.UpdateUserServerRoleResponse]
	countUsers            *connect_go.Client[v1alpha1.CountUsersRequest, v1alpha1.CountUsersResponse]
	updateUserSettings    *connect_go.Client[v1alpha1.UpdateUserSettingsRequest, v1alpha1.UpdateUserSettingsResponse]
}

// CreateUser calls bufman.dubbo.apache.org.registry.v1alpha1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1alpha1.CreateUserRequest]) (*connect_go.Response[v1alpha1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// GetUser calls bufman.dubbo.apache.org.registry.v1alpha1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1alpha1.GetUserRequest]) (*connect_go.Response[v1alpha1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// GetUserByUsername calls
// bufman.dubbo.apache.org.registry.v1alpha1.UserService.GetUserByUsername.
func (c *userServiceClient) GetUserByUsername(ctx context.Context, req *connect_go.Request[v1alpha1.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha1.GetUserByUsernameResponse], error) {
	return c.getUserByUsername.CallUnary(ctx, req)
}

// ListUsers calls bufman.dubbo.apache.org.registry.v1alpha1.UserService.ListUsers.
func (c *userServiceClient) ListUsers(ctx context.Context, req *connect_go.Request[v1alpha1.ListUsersRequest]) (*connect_go.Response[v1alpha1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// ListOrganizationUsers calls
// bufman.dubbo.apache.org.registry.v1alpha1.UserService.ListOrganizationUsers.
func (c *userServiceClient) ListOrganizationUsers(ctx context.Context, req *connect_go.Request[v1alpha1.ListOrganizationUsersRequest]) (*connect_go.Response[v1alpha1.ListOrganizationUsersResponse], error) {
	return c.listOrganizationUsers.CallUnary(ctx, req)
}

// DeleteUser calls bufman.dubbo.apache.org.registry.v1alpha1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteUserRequest]) (*connect_go.Response[v1alpha1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// DeactivateUser calls
// bufman.dubbo.apache.org.registry.v1alpha1.UserService.DeactivateUser.
func (c *userServiceClient) DeactivateUser(ctx context.Context, req *connect_go.Request[v1alpha1.DeactivateUserRequest]) (*connect_go.Response[v1alpha1.DeactivateUserResponse], error) {
	return c.deactivateUser.CallUnary(ctx, req)
}

// UpdateUserServerRole calls
// bufman.dubbo.apache.org.registry.v1alpha1.UserService.UpdateUserServerRole.
func (c *userServiceClient) UpdateUserServerRole(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateUserServerRoleRequest]) (*connect_go.Response[v1alpha1.UpdateUserServerRoleResponse], error) {
	return c.updateUserServerRole.CallUnary(ctx, req)
}

// CountUsers calls bufman.dubbo.apache.org.registry.v1alpha1.UserService.CountUsers.
func (c *userServiceClient) CountUsers(ctx context.Context, req *connect_go.Request[v1alpha1.CountUsersRequest]) (*connect_go.Response[v1alpha1.CountUsersResponse], error) {
	return c.countUsers.CallUnary(ctx, req)
}

// UpdateUserSettings calls
// bufman.dubbo.apache.org.registry.v1alpha1.UserService.UpdateUserSettings.
func (c *userServiceClient) UpdateUserSettings(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateUserSettingsRequest]) (*connect_go.Response[v1alpha1.UpdateUserSettingsResponse], error) {
	return c.updateUserSettings.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the
// bufman.dubbo.apache.org.registry.v1alpha1.UserService service.
type UserServiceHandler interface {
	// CreateUser creates a new user with the given username.
	CreateUser(context.Context, *connect_go.Request[v1alpha1.CreateUserRequest]) (*connect_go.Response[v1alpha1.CreateUserResponse], error)
	// GetUser gets a user by ID.
	GetUser(context.Context, *connect_go.Request[v1alpha1.GetUserRequest]) (*connect_go.Response[v1alpha1.GetUserResponse], error)
	// GetUserByUsername gets a user by username.
	GetUserByUsername(context.Context, *connect_go.Request[v1alpha1.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha1.GetUserByUsernameResponse], error)
	// ListUsers lists all users.
	ListUsers(context.Context, *connect_go.Request[v1alpha1.ListUsersRequest]) (*connect_go.Response[v1alpha1.ListUsersResponse], error)
	// ListOrganizationUsers lists all users for an organization.
	// TODO: #663 move this to organization service
	ListOrganizationUsers(context.Context, *connect_go.Request[v1alpha1.ListOrganizationUsersRequest]) (*connect_go.Response[v1alpha1.ListOrganizationUsersResponse], error)
	// DeleteUser deletes a user.
	DeleteUser(context.Context, *connect_go.Request[v1alpha1.DeleteUserRequest]) (*connect_go.Response[v1alpha1.DeleteUserResponse], error)
	// Deactivate user deactivates a user.
	DeactivateUser(context.Context, *connect_go.Request[v1alpha1.DeactivateUserRequest]) (*connect_go.Response[v1alpha1.DeactivateUserResponse], error)
	// UpdateUserServerRole update the role of an user in the server.
	UpdateUserServerRole(context.Context, *connect_go.Request[v1alpha1.UpdateUserServerRoleRequest]) (*connect_go.Response[v1alpha1.UpdateUserServerRoleResponse], error)
	// CountUsers returns the number of users in the server by the user state provided.
	CountUsers(context.Context, *connect_go.Request[v1alpha1.CountUsersRequest]) (*connect_go.Response[v1alpha1.CountUsersResponse], error)
	// UpdateUserSettings update the user settings including description.
	UpdateUserSettings(context.Context, *connect_go.Request[v1alpha1.UpdateUserSettingsRequest]) (*connect_go.Response[v1alpha1.UpdateUserSettingsResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userServiceCreateUserHandler := connect_go.NewUnaryHandler(
		UserServiceCreateUserProcedure,
		svc.CreateUser,
		connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceGetUserHandler := connect_go.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceGetUserByUsernameHandler := connect_go.NewUnaryHandler(
		UserServiceGetUserByUsernameProcedure,
		svc.GetUserByUsername,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceListUsersHandler := connect_go.NewUnaryHandler(
		UserServiceListUsersProcedure,
		svc.ListUsers,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceListOrganizationUsersHandler := connect_go.NewUnaryHandler(
		UserServiceListOrganizationUsersProcedure,
		svc.ListOrganizationUsers,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceDeleteUserHandler := connect_go.NewUnaryHandler(
		UserServiceDeleteUserProcedure,
		svc.DeleteUser,
		connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceDeactivateUserHandler := connect_go.NewUnaryHandler(
		UserServiceDeactivateUserProcedure,
		svc.DeactivateUser,
		connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceUpdateUserServerRoleHandler := connect_go.NewUnaryHandler(
		UserServiceUpdateUserServerRoleProcedure,
		svc.UpdateUserServerRole,
		opts...,
	)
	userServiceCountUsersHandler := connect_go.NewUnaryHandler(
		UserServiceCountUsersProcedure,
		svc.CountUsers,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceUpdateUserSettingsHandler := connect_go.NewUnaryHandler(
		UserServiceUpdateUserSettingsProcedure,
		svc.UpdateUserSettings,
		opts...,
	)
	return "/bufman.dubbo.apache.org.registry.v1alpha1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceCreateUserProcedure:
			userServiceCreateUserHandler.ServeHTTP(w, r)
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceGetUserByUsernameProcedure:
			userServiceGetUserByUsernameHandler.ServeHTTP(w, r)
		case UserServiceListUsersProcedure:
			userServiceListUsersHandler.ServeHTTP(w, r)
		case UserServiceListOrganizationUsersProcedure:
			userServiceListOrganizationUsersHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserProcedure:
			userServiceDeleteUserHandler.ServeHTTP(w, r)
		case UserServiceDeactivateUserProcedure:
			userServiceDeactivateUserHandler.ServeHTTP(w, r)
		case UserServiceUpdateUserServerRoleProcedure:
			userServiceUpdateUserServerRoleHandler.ServeHTTP(w, r)
		case UserServiceCountUsersProcedure:
			userServiceCountUsersHandler.ServeHTTP(w, r)
		case UserServiceUpdateUserSettingsProcedure:
			userServiceUpdateUserSettingsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect_go.Request[v1alpha1.CreateUserRequest]) (*connect_go.Response[v1alpha1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect_go.Request[v1alpha1.GetUserRequest]) (*connect_go.Response[v1alpha1.GetUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUserByUsername(context.Context, *connect_go.Request[v1alpha1.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha1.GetUserByUsernameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.GetUserByUsername is not implemented"))
}

func (UnimplementedUserServiceHandler) ListUsers(context.Context, *connect_go.Request[v1alpha1.ListUsersRequest]) (*connect_go.Response[v1alpha1.ListUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.ListUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) ListOrganizationUsers(context.Context, *connect_go.Request[v1alpha1.ListOrganizationUsersRequest]) (*connect_go.Response[v1alpha1.ListOrganizationUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.ListOrganizationUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect_go.Request[v1alpha1.DeleteUserRequest]) (*connect_go.Response[v1alpha1.DeleteUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.DeleteUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeactivateUser(context.Context, *connect_go.Request[v1alpha1.DeactivateUserRequest]) (*connect_go.Response[v1alpha1.DeactivateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.DeactivateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUserServerRole(context.Context, *connect_go.Request[v1alpha1.UpdateUserServerRoleRequest]) (*connect_go.Response[v1alpha1.UpdateUserServerRoleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.UpdateUserServerRole is not implemented"))
}

func (UnimplementedUserServiceHandler) CountUsers(context.Context, *connect_go.Request[v1alpha1.CountUsersRequest]) (*connect_go.Response[v1alpha1.CountUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.CountUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUserSettings(context.Context, *connect_go.Request[v1alpha1.UpdateUserSettingsRequest]) (*connect_go.Response[v1alpha1.UpdateUserSettingsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.UserService.UpdateUserSettings is not implemented"))
}
