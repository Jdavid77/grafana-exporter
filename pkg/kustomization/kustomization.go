package kustomization

import (
	"fmt"	
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"sigs.k8s.io/kustomize/pkg/types"
)

func GenerateKustomizations(path string){

	folders, _ := ioutil.ReadDir(path)
	resources := make([]string,len(folders))

	for _ , folder := range folders {
			
		folderName := folder.Name()
		resources = append(resources, folderName)

		filesInFolder, _ := ioutil.ReadDir(fmt.Sprintf("%s/%s", path, folderName))

		k := &types.Kustomization{
			TypeMeta: types.TypeMeta{
				Kind: "Kustomization",
				APIVersion: "kustomize.config.k8s.io/v1beta1",
			},
			Namespace: "monitoring",
			ConfigMapGenerator: func() []types.ConfigMapArgs {

				var config []types.ConfigMapArgs

				for _, file := range filesInFolder{
					config = append(config, types.ConfigMapArgs{
						GeneratorArgs: types.GeneratorArgs{
							Name: file.Name(),
							DataSources: types.DataSources{
								FileSources: []string{
									fmt.Sprintf("%s.json=./%s.json",file.Name(),file.Name()),
								},
							},
						},
					})
				}

				return config
			}(),
			GeneratorOptions: &types.GeneratorOptions{
				DisableNameSuffixHash: true,
				Annotations: map[string]string{
					"kustomize.toolkit.fluxcd.io/substitute": "disabled",
					"grafana_fodler": folderName,
				},
				Labels: map[string]string{
					"grafana_dashboard": "true",
				},
			},
		}

		outputPath := fmt.Sprintf("%s/%s/kustomization.yaml",path,folderName) 
		
		if err := createKustomizationFile(k, outputPath) ; err != nil {
			fmt.Printf("Error creating kustomization for folder %s: %s", folderName, err.Error())
		}
		
	}

	k := &types.Kustomization{
		TypeMeta: types.TypeMeta{
			Kind: "Kustomization",
			APIVersion: "kustomize.config.k8s.io/v1beta1",
		},
		Namespace: "monitoring",
		Resources: resources,
	}

	outputPath := fmt.Sprintf("%s/kustomization.yaml",path)

	if err := createKustomizationFile(k, outputPath) ; err != nil {
		fmt.Printf("Error creating kustomization for root folder: %s", err.Error())
	}
	

}

func createKustomizationFile(kustomization *types.Kustomization, outputPath string) error {
	
	kustomizationYAML, err := yaml.Marshal(kustomization)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(kustomizationYAML)
	if err != nil {
		return err
	}

	return nil
}