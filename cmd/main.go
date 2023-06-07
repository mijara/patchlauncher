package main

import (
	"github.com/gosimple/slug"
	"github.com/manifoldco/promptui"
	"log"
	"path/filepath"
	"smwlauncher/adapter"
	"smwlauncher/interactor"
	"smwlauncher/model"
	"smwlauncher/port"
)

// This is a sample application built on the terminal.
func main() {
	patcherService := adapter.NewPatcherService("multipatch")
	openerService := adapter.NewOpenerService()
	scrapperService := adapter.NewScrapperService("https://www.smwcentral.net", "/?p=section&s=smwhacks")
	downloaderService := adapter.NewDownloaderService("downloads")
	zipCompressionService := adapter.NewZipCompressionService()
	logger := adapter.NewLogger()

	hacks := getHackList(logger, scrapperService)
	selectedHack := promptHackFromList(hacks)
	compressedHackPath := downloadCompressedHack(logger, downloaderService, selectedHack)
	patches := getCompressedHackPatches(logger, zipCompressionService, compressedHackPath)

	compressedPatchPath := ""
	if len(patches) == 1 {
		compressedPatchPath = patches[0]
	} else {
		compressedPatchPath = promptPatchFromList(patches)
	}

	patchPath := extractPatch(logger, zipCompressionService, compressedHackPath, compressedPatchPath)
	patchedROMPath := patchRomFromFile(logger, patcherService, selectedHack, patchPath)
	openROM(logger, openerService, patchedROMPath)
}

func getHackList(
	logger port.Logger,
	scrapperService port.ScrapperService,
) []model.Hack {
	theInteractor := interactor.NewGetHackList(logger, scrapperService)
	output, err := theInteractor.Execute()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return output.Hacks
}

func patchRomFromFile(
	logger port.Logger,
	patcherService port.PatcherService,
	hack model.Hack,
	patchPath string,
) string {
	patchedROMPath := filepath.Join("patched", slug.Make(hack.Title)+".sfc")

	theInteractor := interactor.NewPatchRomFromFile(patcherService, logger)
	if err := theInteractor.Execute(interactor.PatchRomFromFileInput{
		ROMPath:    "smw.sfc",
		PatchPath:  patchPath,
		OutputPath: patchedROMPath,
	}); err != nil {
		log.Fatalln(err.Error())
	}

	return patchedROMPath
}

func openROM(
	logger port.Logger,
	openerService port.OpenerService,
	patchedROMPath string,
) {
	theInteractor := interactor.NewOpenROM(openerService, logger)
	if err := theInteractor.Execute(interactor.OpenROMInput{
		ROMPath: patchedROMPath,
	}); err != nil {
		log.Default().Fatalln(err)
	}
}

func downloadCompressedHack(
	logger port.Logger,
	downloaderService port.DownloaderService,
	hack model.Hack,
) string {
	theInteractor := interactor.NewDownloadCompressedHack(logger, downloaderService)
	path, err := theInteractor.Execute(interactor.DownloadCompressedHackInput{
		URL: hack.DownloadURL,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return path
}

func getCompressedHackPatches(
	logger port.Logger,
	compressionService port.CompressionService,
	compressedHackPath string,
) []string {
	theInteractor := interactor.NewGetCompressedHackPatches(logger, compressionService)
	patches, err := theInteractor.Execute(interactor.GetCompressedHackPatchesInput{
		CompressedHackPath: compressedHackPath,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return patches
}

func promptHackFromList(hacks []model.Hack) model.Hack {
	prompt := promptui.Select{
		Label: "Select ROM Hack",
		Items: hacks,
		Size:  10,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Title }} ({{ .Rating }})",
			Active:   "{{ .Title | cyan }} ({{ .Rating }})",
			Inactive: "{{ .Title }} ({{ .Rating }})",
			Selected: "{{ .Title | cyan }} ({{ .Rating }})",
			Details: `
-------- Details --------
Authors: {{ .Authors }}
Type: {{ .Type }}
Downloads: {{ .Downloads }}
`,
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return hacks[index]
}

func promptPatchFromList(patches []string) string {
	prompt := promptui.Select{
		Label: "Select Patch File",
		Items: patches,
		Size:  10,
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return patches[index]
}

func extractPatch(
	logger port.Logger,
	compressionService port.CompressionService,
	compressedHackPath, compressedPatchPath string,
) string {
	theInteractor := interactor.NewExtractPatch(logger, compressionService)
	patchPath, err := theInteractor.Execute(interactor.ExtractPatchInput{
		CompressedHackPath: compressedHackPath,
		PatchPath:          compressedPatchPath,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return patchPath
}
