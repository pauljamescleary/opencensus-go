// Copyright 2017, OpenCensus Authors
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

package exporter

import (
	"log"

	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

// Exporter is a stats and trace exporter that logs
// the exported data to the console.
type Exporter struct{}

// ExportView logs the view data.
func (e *Exporter) ExportView(vd *view.Data) {
	log.Println(vd)
}

// ExportSpan logs the trace span.
func (e *Exporter) ExportSpan(vd *trace.SpanData) {
	log.Println(vd)
}
