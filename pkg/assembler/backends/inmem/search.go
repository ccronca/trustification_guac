//
// Copyright 2023 The GUAC Authors.
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

package inmem

import (
	"context"
	"fmt"

	"github.com/guacsec/guac/internal/testing/ptrfrom"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/guacsec/guac/pkg/assembler/helpers"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/exp/maps"
)

var edgesAllowed = []model.Edge{
	model.EdgeArtifactCertifyVexStatement,
	model.EdgeCertifyVexStatementVulnerability,
	model.EdgeCertifyVexStatementPackage,
	model.EdgeCertifyVexStatementArtifact,
	model.EdgeCertifyVulnPackage,
	model.EdgeIsDependencyPackage,
	model.EdgePackageCertifyVexStatement,
	model.EdgePackageIsDependency,
	model.EdgeVulnerabilityCertifyVexStatement,
	model.EdgeVulnerabilityCertifyVuln,
}

func (c *demoClient) FindSoftware(ctx context.Context, searchText string) ([]model.PackageSourceOrArtifact, error) {
	return []model.PackageSourceOrArtifact{}, fmt.Errorf("not implemented: FindSoftware")
}

func (c *demoClient) FindTopLevelPackagesRelatedToVulnerability(ctx context.Context, vulnerabilityID string) ([][]model.Node, error) {
	hasSBOMs, err := c.HasSBOM(ctx, &model.HasSBOMSpec{})
	if err != nil {
		return nil, gqlerror.Errorf("FindTopLevelPackagesRelatedToVulnerability failed with err: %v", err)
	}
	if len(hasSBOMs) == 0 {
		return nil, gqlerror.Errorf("FindTopLevelPackagesRelatedToVulnerability failed with err: Found no package with an SBOM")
	}

	var result [][]model.Node
	for _, hasSBOM := range hasSBOMs {
		var idProduct *string
		switch v := hasSBOM.Subject.(type) {
		case *model.Artifact:
			idProduct = &v.ID
		case *model.Package:
			idProduct = &v.Namespaces[0].Names[0].Versions[0].ID
		default:
			idProduct = nil
		}
		if idProduct != nil {
			vexStatements, err := c.CertifyVEXStatement(ctx, &model.CertifyVEXStatementSpec{
				Vulnerability: &model.VulnerabilitySpec{
					VulnerabilityID: &vulnerabilityID,
				},
			})
			if err != nil {
				return nil, gqlerror.Errorf("FindTopLevelPackagesRelatedToVulnerability failed with err: %v", err)
			}
			found := false
			for _, vexStatement := range vexStatements {
				path, err := c.Path(ctx, vexStatement.ID, *idProduct, 10, edgesAllowed)
				if err == nil {
					result = append(result, path)
					found = true
				}
			}
			// if no VEX Statements have been found or no path from any VEX statement to product has been found
			// then let's check also for CertifyVuln
			if len(vexStatements) == 0 || !found {
				vulnStatements, err := c.CertifyVuln(ctx, &model.CertifyVulnSpec{
					Vulnerability: &model.VulnerabilitySpec{
						VulnerabilityID: &vulnerabilityID,
					},
				})
				if err != nil {
					return nil, gqlerror.Errorf("FindTopLevelPackagesRelatedToVulnerability failed with err: %v", err)
				}
				for _, vuln := range vulnStatements {
					path, err := c.Path(ctx, vuln.ID, *idProduct, 10, edgesAllowed)
					if err == nil {
						result = append(result, path)
						found = true
					}
				}
			}
		}
	}
	return result, nil
}

// FindVulnerability returns all vulnerabilities related to a package
func (c *demoClient) FindVulnerability(ctx context.Context, purl string) ([]model.CertifyVulnOrCertifyVEXStatement, error) {

	pkgInput, err := helpers.PurlToPkg(purl)
	if err != nil {
		return nil, gqlerror.Errorf("failed to parse PURL: %v", err)
	}

	pkgQualifierFilter := []*model.PackageQualifierSpec{}
	for _, qualifier := range pkgInput.Qualifiers {
		pkgQualifierFilter = append(pkgQualifierFilter, &model.PackageQualifierSpec{
			Key:   qualifier.Key,
			Value: &qualifier.Value,
		})
	}

	pkgFilter := &model.PkgSpec{
		Type:       &pkgInput.Type,
		Namespace:  pkgInput.Namespace,
		Name:       &pkgInput.Name,
		Version:    pkgInput.Version,
		Subpath:    pkgInput.Subpath,
		Qualifiers: pkgQualifierFilter,
	}
	return c.findVulnerabilities(ctx, pkgFilter)
}

// FindVulnerabilityCPE returns all vulnerabilities related to the package identified by the CPE
func (c *demoClient) FindVulnerabilityCPE(ctx context.Context, cpe string) ([]model.CertifyVulnOrCertifyVEXStatement, error) {

	metadatas, err := c.HasMetadata(ctx, &model.HasMetadataSpec{Key: ptrfrom.String("cpe"), Value: &cpe})
	packagesFound := map[string]*model.Package{}
	if err != nil {
		return nil, gqlerror.Errorf("error querying for HasMetadata: %v", err)
	}
	// if multiple times the same key-value metadata has been attached to the same package,
	// it means the referenced package is just only the same one.
	for i := range metadatas {
		pkg, ok := metadatas[i].Subject.(*model.Package)
		if ok {
			id := pkg.Namespaces[0].Names[0].Versions[0].ID
			if _, found := packagesFound[id]; !found {
				packagesFound[id] = pkg
			}
		}
	}
	if len(maps.Values(packagesFound)) != 1 {
		return nil, gqlerror.Errorf("failed to locate a single package based on the provided CPE")
	}

	pkg := maps.Values(packagesFound)[0]
	pkgQualifierFilter := []*model.PackageQualifierSpec{}
	for _, qualifier := range pkg.Namespaces[0].Names[0].Versions[0].Qualifiers {
		pkgQualifierFilter = append(pkgQualifierFilter, &model.PackageQualifierSpec{
			Key:   qualifier.Key,
			Value: &qualifier.Value,
		})
	}
	pkgFilter := &model.PkgSpec{
		Type:       &pkg.Type,
		Namespace:  &pkg.Namespaces[0].Namespace,
		Name:       &pkg.Namespaces[0].Names[0].Name,
		Version:    &pkg.Namespaces[0].Names[0].Versions[0].Version,
		Subpath:    &pkg.Namespaces[0].Names[0].Versions[0].Subpath,
		Qualifiers: pkgQualifierFilter,
	}
	return c.findVulnerabilities(ctx, pkgFilter)
}

func (c *demoClient) findVulnerabilities(ctx context.Context, pkgFilter *model.PkgSpec) ([]model.CertifyVulnOrCertifyVEXStatement, error) {

	pkgResponse, err := c.Packages(ctx, pkgFilter)
	if err != nil {
		return nil, gqlerror.Errorf("error querying for package: %v", err)
	}
	if len(pkgResponse) != 1 {
		return nil, gqlerror.Errorf("failed to locate package based on purl")
	}

	vulnerabilities := []model.CertifyVulnOrCertifyVEXStatement{}

	vexStatements, err := c.CertifyVEXStatement(ctx, &model.CertifyVEXStatementSpec{})
	if err != nil {
		return nil, gqlerror.Errorf("findVulnerabilities failed with err: %v", err)
	}
	idProduct := pkgResponse[0].Namespaces[0].Names[0].Versions[0].ID
	for _, vexStatement := range vexStatements {
		path, err := c.Path(ctx, vexStatement.ID, idProduct, 10, edgesAllowed)
		if err == nil {
			vulnerabilities = append(vulnerabilities, path[0].(*model.CertifyVEXStatement))
		}
	}
	vulnStatements, err := c.CertifyVuln(ctx, &model.CertifyVulnSpec{})
	if err != nil {
		return nil, gqlerror.Errorf("findVulnerabilities failed with err: %v", err)
	}
	for _, vuln := range vulnStatements {
		path, err := c.Path(ctx, vuln.ID, idProduct, 10, edgesAllowed)
		if err == nil {
			certifyVuln := path[0].(*model.CertifyVuln)
			if certifyVuln.Vulnerability.Type != noVulnType {
				vulnerabilities = append(vulnerabilities, certifyVuln)
			}
		}
	}

	return vulnerabilities, nil
}
