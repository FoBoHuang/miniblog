// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package grpc

import (
	apiv1 "github.com/onexstack/miniblog/pkg/api/apiserver/v1"
)

// Handler 负责处理博客模块的请求.
type Handler struct {
	// 必须内嵌 apiv1.UnimplementedMiniBlogServer 类型
	// 为了提供默认实现，确保未实现的 gRPC 方法返回“未实现”错误，同时满足接口要求，简化服务端开发和向后兼容性。
	apiv1.UnimplementedMiniBlogServer
}

// NewHandler 创建一个新的 Handler 实例.
func NewHandler() *Handler {
	return &Handler{}
}
