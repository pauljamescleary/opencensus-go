// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trace

import (
	"context"
	"testing"
)

func BenchmarkStartEndSpan(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, span := StartSpan(ctx, "/foo")
		span.End()
	}
}

func BenchmarkSpanWithAnnotations_3(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, span := StartSpan(ctx, "/foo")
		span.SetAttributes(
			BoolAttribute{Key: "key1", Value: false},
			StringAttribute{Key: "key2", Value: "hello"},
			Int64Attribute{Key: "key3", Value: 123},
		)
		span.End()
	}
}

func BenchmarkSpanWithAnnotations_6(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, span := StartSpan(ctx, "/foo")
		span.SetAttributes(
			BoolAttribute{Key: "key1", Value: false},
			BoolAttribute{Key: "key2", Value: true},
			StringAttribute{Key: "key3", Value: "hello"},
			StringAttribute{Key: "key4", Value: "hello"},
			Int64Attribute{Key: "key5", Value: 123},
			Int64Attribute{Key: "key6", Value: 456},
		)
		span.End()
	}
}
