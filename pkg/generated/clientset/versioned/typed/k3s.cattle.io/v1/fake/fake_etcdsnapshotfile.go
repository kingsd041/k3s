/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1 "github.com/k3s-io/k3s/pkg/apis/k3s.cattle.io/v1"
	k3scattleiov1 "github.com/k3s-io/k3s/pkg/generated/clientset/versioned/typed/k3s.cattle.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeETCDSnapshotFiles implements ETCDSnapshotFileInterface
type fakeETCDSnapshotFiles struct {
	*gentype.FakeClientWithList[*v1.ETCDSnapshotFile, *v1.ETCDSnapshotFileList]
	Fake *FakeK3sV1
}

func newFakeETCDSnapshotFiles(fake *FakeK3sV1) k3scattleiov1.ETCDSnapshotFileInterface {
	return &fakeETCDSnapshotFiles{
		gentype.NewFakeClientWithList[*v1.ETCDSnapshotFile, *v1.ETCDSnapshotFileList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("etcdsnapshotfiles"),
			v1.SchemeGroupVersion.WithKind("ETCDSnapshotFile"),
			func() *v1.ETCDSnapshotFile { return &v1.ETCDSnapshotFile{} },
			func() *v1.ETCDSnapshotFileList { return &v1.ETCDSnapshotFileList{} },
			func(dst, src *v1.ETCDSnapshotFileList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ETCDSnapshotFileList) []*v1.ETCDSnapshotFile { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ETCDSnapshotFileList, items []*v1.ETCDSnapshotFile) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
