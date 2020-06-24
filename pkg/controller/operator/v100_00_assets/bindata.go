// Code generated for package v100_0_assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// bindata/v1.0.0/clusterrole.yaml
// bindata/v1.0.0/deployment.yaml
// bindata/v1.0.0/serviceaccount.yaml
package v100_0_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v100ClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: openshift-acme
  labels:
    app: openshift-acme
rules:
- apiGroups:
  - "route.openshift.io"
  resources:
  - routes
  verbs:
  - create
  - get
  - list
  - watch
  - update
  - patch
  - delete
  - deletecollection

- apiGroups:
  - "route.openshift.io"
  resources:
  - routes/custom-host
  verbs:
  - create

- apiGroups:
  - ""
  resources:
  - configmaps
  - services
  - secrets
  verbs:
  - create
  - get
  - list
  - watch
  - update
  - patch
  - delete

- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - update
  - patch

- apiGroups:
  - ""
  resources:
  - limitranges
  verbs:
  - get
  - list
  - watch

- apiGroups:
  - "apps"
  resources:
  - replicasets
  verbs:
  - create
  - get
  - list
  - watch
  - update
  - patch
  - delete
`)

func v100ClusterroleYamlBytes() ([]byte, error) {
	return _v100ClusterroleYaml, nil
}

func v100ClusterroleYaml() (*asset, error) {
	bytes, err := v100ClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1.0.0/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v100DeploymentYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: openshift-acme
  labels:
    app: openshift-acme
  annotations:
spec:
  selector:
    matchLabels:
      app: openshift-acme
  replicas: 2
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: openshift-acme
    spec:
      serviceAccountName: openshift-acme
      containers:
      - name: openshift-acme
        image: quay.io/tnozicka/openshift-acme:controller
        imagePullPolicy: Always
        args:
        - --exposer-image=quay.io/tnozicka/openshift-acme:exposer
        - --loglevel=4
`)

func v100DeploymentYamlBytes() ([]byte, error) {
	return _v100DeploymentYaml, nil
}

func v100DeploymentYaml() (*asset, error) {
	bytes, err := v100DeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1.0.0/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v100ServiceaccountYaml = []byte(`kind: ServiceAccount
apiVersion: v1
metadata:
  name: openshift-acme
  labels:
    app: openshift-acme
`)

func v100ServiceaccountYamlBytes() ([]byte, error) {
	return _v100ServiceaccountYaml, nil
}

func v100ServiceaccountYaml() (*asset, error) {
	bytes, err := v100ServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1.0.0/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v1.0.0/clusterrole.yaml":    v100ClusterroleYaml,
	"v1.0.0/deployment.yaml":     v100DeploymentYaml,
	"v1.0.0/serviceaccount.yaml": v100ServiceaccountYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v1.0.0": {nil, map[string]*bintree{
		"clusterrole.yaml":    {v100ClusterroleYaml, map[string]*bintree{}},
		"deployment.yaml":     {v100DeploymentYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {v100ServiceaccountYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
