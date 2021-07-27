/*
 * Copyright 2020 Huawei Technologies Co., Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kie

import (
	"context"
	"net/http"
)

func WithAuthContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, HeaderAuth, token)
}

func Headers(ctx context.Context) http.Header {
	h := http.Header{}
	h.Set(HeaderContentType, ContentTypeJSON)
	token := ctx.Value(HeaderAuth)
	if token != nil {
		h.Set(HeaderAuth, token.(string))
	}
	return h
}
