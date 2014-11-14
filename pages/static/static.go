// Copyright 2014 Google Inc. All Rights Reserved.
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

// Handler for /static content.

package static

import (
	"fmt"
	"net/http"
	"net/url"
)

const StaticResource = "/static/"

var staticFiles = map[string]string{
	"containers.css":          containersCss,
	"containers.js":           containersJs,
	"bootstrap.min.css":       bootstrapCss,
	"bootstrap-theme.min.css": bootstrapThemeCss,
	"jquery.min.js":           jqueryJs,
	"bootstrap.min.js":        bootstrapJs,
	"google-jsapi.js":         googleJsapiJs,
}

func HandleRequest(w http.ResponseWriter, u *url.URL) error {
	if len(u.Path) <= len(StaticResource) {
		return fmt.Errorf("unknown static resource %q", u.Path)
	}

	// Get the static content if it exists.
	resource := u.Path[len(StaticResource):]
	content, ok := staticFiles[resource]
	if !ok {
		return fmt.Errorf("unknown static resource %q", resource)
	}

	_, err := w.Write([]byte(content))
	return err
}
